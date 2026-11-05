English | [简体中文](README-CN.md)

![](https://aliyunsdk-pages.alicdn.com/icons/AlibabaCloud.svg)

# Alibaba Cloud POP Gateway

This package signs POP requests for Alibaba Cloud V2 SDKs. A product client already carries the POP product name and this gateway. To sign with signature V4, set the signature version to `v4`. Leaving the signature version unset keeps the default, signature V3.

SHA-256 signing does not depend on the endpoint, OpenAPI, or Tea utility packages. SM3 still uses that language's SM3 library. The gateway SPI and credentials packages stay.

## Packages

| Language | Package | This release |
| --- | --- | --- |
| Go | `github.com/alibabacloud-go/alibabacloud-gateway-pop` | [0.1.5](golang/ChangeLog.md) |
| Java | `com.aliyun:alibabacloud-gateway-pop` | [0.1.5](java/ChangeLog.md) |
| Python | `alibabacloud_gateway_pop` | [0.1.5](python/ChangeLog.md) |
| TypeScript | `@alicloud/gateway-pop` | [0.1.7](ts/ChangeLog.md) |
| PHP | `alibabacloud/gateway-pop` | [0.1.5](php/ChangeLog.md) |
| C# | `AlibabaCloud.GatewayPop` | [0.1.4](csharp/ChangeLog.md) |
| Swift | `AlibabacloudGatewayPOP` | [0.0.2](swift/ChangeLog.md) |
| Python 2 | `alibabacloud_gateway_pop_py2` | [0.0.4, not in this release](python2/ChangeLog.md) |

C++ has an OpenAPI signature version field, and this repository does not contain a C++ POP gateway.

## Issues

[Opening an Issue](https://github.com/aliyun/alibabacloud-gateway/issues/new). Issues that do not follow the guidelines may be closed immediately.

## License

[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)

Copyright (c) 2009-present, Alibaba Cloud All rights reserved.
