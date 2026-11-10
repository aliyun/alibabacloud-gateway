### 2026-11-10 Version 0.0.3

* Fix POP V4 query signing by percent-encoding parameter names and values before sorting by the encoded names.
* Correct query percent-encoding and canonicalize signed header names case-insensitively.
* Use a single entry per header name. Tea Swift can collapse keys that differ only in case when sending the request.

### 2026-11-05 Version 0.0.2

* Sign POP V4 without the endpoint, OpenAPI, and Tea utility packages.
* Depend on Tea 1.1, CryptoSwift 1.10, and the published gateway SPI 0.0.4.
* Read the endpoint suffix from the request, then `ALIBABA_CLOUD_ENDPOINT_SUFFIX`, otherwise `aliyuncs.com`.
* Rewrite only a generated host that ends in `.aliyuncs.com`. A caller-supplied endpoint and an endpointMap entry stay unchanged.
* Reject a suffix other than `aliyuncs.com` or `aliyuncs.alibaba`.
* Keep signature prefixes on the request. When the request omits them, use `ACS4-` and `aliyun_v4`.
