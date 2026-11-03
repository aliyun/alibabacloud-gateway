[English](README.md) | 简体中文
![](https://aliyunsdk-pages.alicdn.com/icons/AlibabaCloud.svg)

# Alibaba Cloud OpenApi Client

## POP Endpoint 后缀切换

设置 `endpointSuffix=aliyuncs.alibaba` 后，POP Gateway 会将 endpoint 末尾的 `.aliyuncs.com` 替换为 `.aliyuncs.alibaba`。这适用于规则生成的 endpoint、`Config.endpoint`、请求的 `endpointOverride` 和 `endpointMap` 命中值，包括 `-internal` / VPC endpoint。`.aliyun.com`、其他自定义域名及已切换的 endpoint 保持不变。默认后缀 `aliyuncs.com` 不改写 endpoint。

## 问题

[Opening an Issue](https://github.com/aliyun/alibabacloud-gateway/issues/new)，不符合指南的问题可能会立即关闭。

## 相关

- [Latest Release](https://github.com/aliyun/alibabacloud-gateway)

## 许可证

[Apache-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Copyright (c) 2009-present, Alibaba Cloud All rights reserved.
