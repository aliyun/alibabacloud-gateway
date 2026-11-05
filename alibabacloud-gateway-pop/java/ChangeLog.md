### 2026-11-05 Version 0.1.5

* Sign POP V4 without endpoint-util, openapi-util, tea-util, or tea-xml.
* Keep the compiler target at 1.7. The supported runtime remains JDK 8 and later.
* Parse POP error XML inside the gateway. SM3 still uses `org.bouncycastle:bcprov-jdk18on`.
* Depend on `com.aliyun:tea` 1.4.2.
* Read the endpoint suffix from the request, then `ALIBABA_CLOUD_ENDPOINT_SUFFIX`, otherwise `aliyuncs.com`.
* Rewrite only a generated host that ends in `.aliyuncs.com`. A caller-supplied endpoint and an endpointMap entry stay unchanged.
* Reject a suffix other than `aliyuncs.com` or `aliyuncs.alibaba`.
* Keep signature prefixes on the request. When the request omits them, use `ACS4-` and `aliyun_v4`.
