### 2026-11-05 Version 0.1.5

* Sign POP V4 without endpoint-util, openapi-util, tea-utils, tea-xml, or the small Darabonba utility packages.
* Keep returned errors as `*tea.SDKError`, so an HTTP 4xx response stays non-retryable for the existing retry check.
* Parse POP error XML with the Darabonba builtin. SM3 still uses `github.com/tjfoc/gmsm`.
* Read the endpoint suffix from the request, then `ALIBABA_CLOUD_ENDPOINT_SUFFIX`, otherwise `aliyuncs.com`.
* Rewrite only a generated host that ends in `.aliyuncs.com`. A caller-supplied endpoint and an endpointMap entry stay unchanged.
* Reject a suffix other than `aliyuncs.com` or `aliyuncs.alibaba`.
* Keep signature prefixes on the request. When the request omits them, use `ACS4-` and `aliyun_v4`.
