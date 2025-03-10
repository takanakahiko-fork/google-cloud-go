// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package impersonate

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"cloud.google.com/go/auth"
	"cloud.google.com/go/auth/internal"
)

const (
	defaultTokenLifetime = "3600s"
	authHeaderKey        = "Authorization"
)

// generateAccesstokenReq is used for service account impersonation
type generateAccessTokenReq struct {
	Delegates []string `json:"delegates,omitempty"`
	Lifetime  string   `json:"lifetime,omitempty"`
	Scope     []string `json:"scope,omitempty"`
}

type impersonateTokenResponse struct {
	AccessToken string `json:"accessToken"`
	ExpireTime  string `json:"expireTime"`
}

// NewTokenProvider uses a source credential, stored in Ts, to request an access token to the provided URL.
// Scopes can be defined when the access token is requested.
func NewTokenProvider(opts *Options) (auth.TokenProvider, error) {
	// TODO(codyoss): add validation
	return opts, nil
}

// Options for [NewTokenProvider].
type Options struct {
	// Tp is the source credential used to generate a token on the
	// impersonated service account. Required.
	Tp auth.TokenProvider

	// URL is the endpoint to call to generate a token
	// on behalf of the service account. Required.
	URL string
	// Scopes that the impersonated credential should have. Required.
	Scopes []string
	// Delegates are the service account email addresses in a delegation chain.
	// Each service account must be granted roles/iam.serviceAccountTokenCreator
	// on the next service account in the chain. Optional.
	Delegates []string
	// TokenLifetimeSeconds is the number of seconds the impersonation token will
	// be valid for. Defaults to 1 hour if unset. Optional.
	TokenLifetimeSeconds int
	// Client configures the underlying client used to make network requests
	// when fetching tokens. Required.
	Client *http.Client
}

// Token performs the exchange to get a temporary service account token to allow access to GCP.
func (tp *Options) Token(ctx context.Context) (*auth.Token, error) {
	lifetime := defaultTokenLifetime
	if tp.TokenLifetimeSeconds != 0 {
		lifetime = fmt.Sprintf("%ds", tp.TokenLifetimeSeconds)
	}
	reqBody := generateAccessTokenReq{
		Lifetime:  lifetime,
		Scope:     tp.Scopes,
		Delegates: tp.Delegates,
	}
	b, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("detect: unable to marshal request: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "POST", tp.URL, bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("detect: unable to create impersonation request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	if err := setAuthHeader(ctx, tp.Tp, req); err != nil {
		return nil, err
	}
	resp, err := tp.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("detect: unable to generate access token: %w", err)
	}
	defer resp.Body.Close()
	body, err := internal.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("detect: unable to read body: %w", err)
	}
	if c := resp.StatusCode; c < http.StatusOK || c >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("detect: status code %d: %s", c, body)
	}

	var accessTokenResp impersonateTokenResponse
	if err := json.Unmarshal(body, &accessTokenResp); err != nil {
		return nil, fmt.Errorf("detect: unable to parse response: %w", err)
	}
	expiry, err := time.Parse(time.RFC3339, accessTokenResp.ExpireTime)
	if err != nil {
		return nil, fmt.Errorf("detect: unable to parse expiry: %w", err)
	}
	return &auth.Token{
		Value:  accessTokenResp.AccessToken,
		Expiry: expiry,
		Type:   internal.TokenTypeBearer,
	}, nil
}

func setAuthHeader(ctx context.Context, tp auth.TokenProvider, r *http.Request) error {
	t, err := tp.Token(ctx)
	if err != nil {
		return err
	}
	typ := t.Type
	if typ == "" {
		typ = internal.TokenTypeBearer
	}
	r.Header.Set(authHeaderKey, typ+" "+t.Value)
	return nil
}
