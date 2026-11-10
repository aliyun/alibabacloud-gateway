### 2026-11-10 Version 0.1.6

* Set the minimum Tea dependency to 1.1.3 so generated SDKs that reference Tea 1.1.3 build without NU1605.

### 2026-11-10 Version 0.1.5

* Fix POP V4 query signing by percent-encoding parameter names and values before sorting by the encoded names.

### 2026-11-05 Version 0.1.4

* Sign POP V4 with Tea. This package no longer depends on the Darabonba NuGet package.
* Parse POP error XML inside the gateway.
* SM3 uses `BouncyCastle.Cryptography` 2.7.0 on `netstandard2.0` and `Portable.BouncyCastle` 1.9.0 on `net45`.
* The assembly version stays `0.1.2.0`. The file version matches this package version.
* Read the endpoint suffix from the request, then `ALIBABA_CLOUD_ENDPOINT_SUFFIX`, otherwise `aliyuncs.com`.
* Rewrite only a generated host that ends in `.aliyuncs.com`. A caller-supplied endpoint and an endpointMap entry stay unchanged.
* Reject a suffix other than `aliyuncs.com` or `aliyuncs.alibaba`.
* Keep signature prefixes on the request. When the request omits them, use `ACS4-` and `aliyun_v4`.
