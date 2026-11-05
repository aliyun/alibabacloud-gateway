### 2026-11-05 Version 0.1.5

* Sign POP V4 without endpoint-util, openapi-util, tea-util, or tea-xml.
* Keep JSON encoding for `TeaModel`, including a model nested in a dict, by calling `to_map()`.
* Parse POP error XML with the Darabonba builtin.
* Read the endpoint suffix from the request, then `ALIBABA_CLOUD_ENDPOINT_SUFFIX`, otherwise `aliyuncs.com`.
* Rewrite only a generated host that ends in `.aliyuncs.com`. A caller-supplied endpoint and an endpointMap entry stay unchanged.
* Reject a suffix other than `aliyuncs.com` or `aliyuncs.alibaba`.
* Keep signature prefixes on the request. When the request omits them, use `ACS4-` and `aliyun_v4`.
