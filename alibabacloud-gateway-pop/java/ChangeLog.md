### 2026-11-03 Version 0.1.5

* Sign POP V4 without endpoint-util, openapi-util, tea-util, or tea-xml.
* Keep the compiler target at 1.7. The supported runtime remains JDK 8 and later.
* Parse POP error XML inside the gateway. SM3 still uses `org.bouncycastle:bcprov-jdk18on`.
* Depend on `com.aliyun:tea` 1.4.2.
