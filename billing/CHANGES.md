# Changes

## [1.17.0](https://github.com/googleapis/google-cloud-go/compare/billing/v1.16.0...billing/v1.17.0) (2023-09-11)


### Features

* **billing/budgets:** Supported project-level-budgets in Public Budget API V1 ([20725c8](https://github.com/googleapis/google-cloud-go/commit/20725c86c970ad24efa18c056fc3aa71dc3a4f03))
* **billing/budgets:** Supported project-level-budgets in Public Budget API V1Beta ([20725c8](https://github.com/googleapis/google-cloud-go/commit/20725c86c970ad24efa18c056fc3aa71dc3a4f03))


### Bug Fixes

* **billing:** Fixed resource_reference for name in GetProjectBillingInfo ([ac10224](https://github.com/googleapis/google-cloud-go/commit/ac102249403e6c1604bff7c537343645c950ae13))


### Documentation

* **billing:** Update comments ([20725c8](https://github.com/googleapis/google-cloud-go/commit/20725c86c970ad24efa18c056fc3aa71dc3a4f03))

## [1.16.0](https://github.com/googleapis/google-cloud-go/compare/billing/v1.15.0...billing/v1.16.0) (2023-06-20)


### Features

* **billing:** Added resource_reference for name in GetProjectBillingInfoRequest message ([b726d41](https://github.com/googleapis/google-cloud-go/commit/b726d413166faa8c84c0a09c6019ff50f3249b9d))


### Bug Fixes

* **billing:** REST query UpdateMask bug ([df52820](https://github.com/googleapis/google-cloud-go/commit/df52820b0e7721954809a8aa8700b93c5662dc9b))

## [1.15.0](https://github.com/googleapis/google-cloud-go/compare/billing/v1.14.0...billing/v1.15.0) (2023-06-07)


### Features

* **billing/budgets:** Add resource_ancestors field to support filtering by folders & organizations ([79eac77](https://github.com/googleapis/google-cloud-go/commit/79eac771ecf99172157cc4499ba95536778354e6))
* **billing/budgets:** Add resource_ancestors field to support filtering by folders & organizations ([79eac77](https://github.com/googleapis/google-cloud-go/commit/79eac771ecf99172157cc4499ba95536778354e6))

## [1.14.0](https://github.com/googleapis/google-cloud-go/compare/billing/v1.13.1...billing/v1.14.0) (2023-05-30)


### Features

* **billing:** Update all direct dependencies ([b340d03](https://github.com/googleapis/google-cloud-go/commit/b340d030f2b52a4ce48846ce63984b28583abde6))

## [1.13.1](https://github.com/googleapis/google-cloud-go/compare/billing/v1.13.0...billing/v1.13.1) (2023-05-08)


### Bug Fixes

* **billing:** Update grpc to v1.55.0 ([1147ce0](https://github.com/googleapis/google-cloud-go/commit/1147ce02a990276ca4f8ab7a1ab65c14da4450ef))

## [1.13.0](https://github.com/googleapis/google-cloud-go/compare/billing/v1.12.0...billing/v1.13.0) (2023-03-15)


### Features

* **billing:** Update iam and longrunning deps ([91a1f78](https://github.com/googleapis/google-cloud-go/commit/91a1f784a109da70f63b96414bba8a9b4254cddd))

## [1.12.0](https://github.com/googleapis/google-cloud-go/compare/billing-v1.11.0...billing/v1.12.0) (2023-01-26)


### Features

* **billing/budgets/apiv1beta1:** Add REST transport ([f7b0822](https://github.com/googleapis/google-cloud-go/commit/f7b082212b1e46ff2f4126b52d49618785c2e8ca))
* **billing:** Add REST client ([06a54a1](https://github.com/googleapis/google-cloud-go/commit/06a54a16a5866cce966547c51e203b9e09a25bc0))
* **billing:** Added Sku.geo_taxonomy fix: more oauth scopes ([2a0b1ae](https://github.com/googleapis/google-cloud-go/commit/2a0b1aeb1683222e6aa5c876cb945845c00cef79))
* **billing:** Enable REST transport in C# ([447afdd](https://github.com/googleapis/google-cloud-go/commit/447afddf34d59c599cabe5415b4f9265b228bb9a))
* **billing:** Rewrite signatures in terms of new location ([3c4b2b3](https://github.com/googleapis/google-cloud-go/commit/3c4b2b34565795537aac1661e6af2442437e34ad))
* **billing:** Rewrite signatures in terms of new types for betas ([9f303f9](https://github.com/googleapis/google-cloud-go/commit/9f303f9efc2e919a9a6bd828f3cdb1fcb3b8b390))
* **billing:** Start generating proto message types ([563f546](https://github.com/googleapis/google-cloud-go/commit/563f546262e68102644db64134d1071fc8caa383))
* **billing:** Start generating stubs dir ([de2d180](https://github.com/googleapis/google-cloud-go/commit/de2d18066dc613b72f6f8db93ca60146dabcfdcc))

## [1.11.0](https://github.com/googleapis/google-cloud-go/compare/billing-v1.10.0...billing/v1.11.0) (2023-01-26)


### Features

* **billing/budgets/apiv1beta1:** Add REST transport ([f7b0822](https://github.com/googleapis/google-cloud-go/commit/f7b082212b1e46ff2f4126b52d49618785c2e8ca))
* **billing:** Add REST client ([06a54a1](https://github.com/googleapis/google-cloud-go/commit/06a54a16a5866cce966547c51e203b9e09a25bc0))
* **billing:** Added Sku.geo_taxonomy fix: more oauth scopes ([2a0b1ae](https://github.com/googleapis/google-cloud-go/commit/2a0b1aeb1683222e6aa5c876cb945845c00cef79))
* **billing:** Enable REST transport in C# ([447afdd](https://github.com/googleapis/google-cloud-go/commit/447afddf34d59c599cabe5415b4f9265b228bb9a))
* **billing:** Rewrite signatures in terms of new location ([3c4b2b3](https://github.com/googleapis/google-cloud-go/commit/3c4b2b34565795537aac1661e6af2442437e34ad))
* **billing:** Rewrite signatures in terms of new types for betas ([9f303f9](https://github.com/googleapis/google-cloud-go/commit/9f303f9efc2e919a9a6bd828f3cdb1fcb3b8b390))
* **billing:** Start generating proto message types ([563f546](https://github.com/googleapis/google-cloud-go/commit/563f546262e68102644db64134d1071fc8caa383))
* **billing:** Start generating stubs dir ([de2d180](https://github.com/googleapis/google-cloud-go/commit/de2d18066dc613b72f6f8db93ca60146dabcfdcc))

## [1.10.0](https://github.com/googleapis/google-cloud-go/compare/billing/v1.9.0...billing/v1.10.0) (2023-01-26)


### Features

* **billing:** Enable REST transport in C# ([447afdd](https://github.com/googleapis/google-cloud-go/commit/447afddf34d59c599cabe5415b4f9265b228bb9a))

## [1.9.0](https://github.com/googleapis/google-cloud-go/compare/billing/v1.8.0...billing/v1.9.0) (2023-01-04)


### Features

* **billing:** Add REST client ([06a54a1](https://github.com/googleapis/google-cloud-go/commit/06a54a16a5866cce966547c51e203b9e09a25bc0))

## [1.8.0](https://github.com/googleapis/google-cloud-go/compare/billing/v1.7.0...billing/v1.8.0) (2022-12-01)


### Features

* **billing:** added Sku.geo_taxonomy fix: more oauth scopes ([2a0b1ae](https://github.com/googleapis/google-cloud-go/commit/2a0b1aeb1683222e6aa5c876cb945845c00cef79))

## [1.7.0](https://github.com/googleapis/google-cloud-go/compare/billing/v1.6.0...billing/v1.7.0) (2022-11-03)


### Features

* **billing:** rewrite signatures in terms of new location ([3c4b2b3](https://github.com/googleapis/google-cloud-go/commit/3c4b2b34565795537aac1661e6af2442437e34ad))

## [1.6.0](https://github.com/googleapis/google-cloud-go/compare/billing/v1.5.0...billing/v1.6.0) (2022-10-25)


### Features

* **billing:** start generating stubs dir ([de2d180](https://github.com/googleapis/google-cloud-go/commit/de2d18066dc613b72f6f8db93ca60146dabcfdcc))

## [1.5.0](https://github.com/googleapis/google-cloud-go/compare/billing/v1.4.0...billing/v1.5.0) (2022-09-21)


### Features

* **billing:** rewrite signatures in terms of new types for betas ([9f303f9](https://github.com/googleapis/google-cloud-go/commit/9f303f9efc2e919a9a6bd828f3cdb1fcb3b8b390))

## [1.4.0](https://github.com/googleapis/google-cloud-go/compare/billing/v1.3.0...billing/v1.4.0) (2022-09-19)


### Features

* **billing:** start generating proto message types ([563f546](https://github.com/googleapis/google-cloud-go/commit/563f546262e68102644db64134d1071fc8caa383))

## [1.3.0](https://github.com/googleapis/google-cloud-go/compare/billing/v1.2.0...billing/v1.3.0) (2022-09-15)


### Features

* **billing/budgets/apiv1beta1:** add REST transport ([f7b0822](https://github.com/googleapis/google-cloud-go/commit/f7b082212b1e46ff2f4126b52d49618785c2e8ca))

## [1.2.0](https://github.com/googleapis/google-cloud-go/compare/billing/v1.1.0...billing/v1.2.0) (2022-02-23)


### Features

* **billing:** set versionClient to module version ([55f0d92](https://github.com/googleapis/google-cloud-go/commit/55f0d92bf112f14b024b4ab0076c9875a17423c9))

## [1.1.0](https://github.com/googleapis/google-cloud-go/compare/billing/v1.0.0...billing/v1.1.0) (2022-02-14)


### Features

* **billing:** add file for tracking version ([17b36ea](https://github.com/googleapis/google-cloud-go/commit/17b36ead42a96b1a01105122074e65164357519e))

## 1.0.0

Stabilize GA surface.

## v0.1.0

This is the first tag to carve out billing as its own module. See
[Add a module to a multi-module repository](https://github.com/golang/go/wiki/Modules#is-it-possible-to-add-a-module-to-a-multi-module-repository).
