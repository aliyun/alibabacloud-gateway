[English](README.md) | 简体中文

![](https://aliyunsdk-pages.alicdn.com/icons/AlibabaCloud.svg)

# 阿里云 POP Gateway

这个包给阿里云 V2 SDK 的 POP 请求签名。产品 Client 里已经带上 POP 产品名和这个 gateway。签名要换成 V4 时，把签名版本设为 `v4`。不设置签名版本时，默认仍是签名 V3。

SHA-256 签名不再依赖 endpoint、OpenAPI 和 Tea 的工具包。SM3 仍使用该语言自己的 SM3 库。Gateway SPI 和 Credentials 保留。

## 包坐标

| 语言 | 包名 | 本次版本 |
| --- | --- | --- |
| Go | `github.com/alibabacloud-go/alibabacloud-gateway-pop` | [0.1.5](golang/ChangeLog.md) |
| Java | `com.aliyun:alibabacloud-gateway-pop` | [0.1.5](java/ChangeLog.md) |
| Python | `alibabacloud_gateway_pop` | [0.1.5](python/ChangeLog.md) |
| TypeScript | `@alicloud/gateway-pop` | [0.1.7](ts/ChangeLog.md) |
| PHP | `alibabacloud/gateway-pop` | [0.1.5](php/ChangeLog.md) |
| C# | `AlibabaCloud.GatewayPop` | [0.1.4](csharp/ChangeLog.md) |
| Swift | `AlibabacloudGatewayPOP` | [0.0.2](swift/ChangeLog.md) |
| Python 2 | `alibabacloud_gateway_pop_py2` | [0.0.4，本次不发布](python2/ChangeLog.md) |

C++ 的 OpenAPI 有签名版本字段，本仓库没有 C++ POP gateway。

## 问题

[提交 Issue](https://github.com/aliyun/alibabacloud-gateway/issues/new)。不符合要求的 Issue 可能会被直接关闭。

## 许可证

[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)

Copyright (c) 2009-present, Alibaba Cloud All rights reserved.
