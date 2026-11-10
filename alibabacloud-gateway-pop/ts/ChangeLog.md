### 2026-11-10 Version 0.1.8

* Fix POP V4 query signing by percent-encoding parameter names and values before sorting by the encoded names.
* Percent-encode every occurrence of reserved characters in canonical query names and values.

### 2026-11-05 Version 0.1.7

* Sign POP V4 without endpoint-util, openapi-util, tea-util, or tea-xml.
* Depend on `@alicloud/tea-typescript` 1.8 and `@darabonba/typescript` 1.0.
* Parse POP error XML with the Darabonba builtin.
* Read the endpoint suffix from the request, then `ALIBABA_CLOUD_ENDPOINT_SUFFIX`, otherwise `aliyuncs.com`.
* Rewrite only a generated host that ends in `.aliyuncs.com`. A caller-supplied endpoint and an endpointMap entry stay unchanged.
* Reject a suffix other than `aliyuncs.com` or `aliyuncs.alibaba`.
* Keep signature prefixes on the request. When the request omits them, use `ACS4-` and `aliyun_v4`.
