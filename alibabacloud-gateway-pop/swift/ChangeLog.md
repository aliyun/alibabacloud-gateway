### 2026-11-05 Version 0.0.2

* Sign POP V4 without the endpoint, OpenAPI, and Tea utility packages.
* Depend on Tea 1.1, CryptoSwift 1.10, and the published gateway SPI 0.0.4.
* Read the endpoint suffix from the request, then `ALIBABA_CLOUD_ENDPOINT_SUFFIX`, otherwise `aliyuncs.com`.
* Rewrite only a generated host that ends in `.aliyuncs.com`. A caller-supplied endpoint and an endpointMap entry stay unchanged.
* Reject a suffix other than `aliyuncs.com` or `aliyuncs.alibaba`.
* Keep signature prefixes on the request. When the request omits them, use `ACS4-` and `aliyun_v4`.
