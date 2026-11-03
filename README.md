English | [简体中文](README-CN.md)

![](https://aliyunsdk-pages.alicdn.com/icons/AlibabaCloud.svg)

# Alibaba Cloud OpenApi Client

## POP endpoint suffix switching

With `endpointSuffix=aliyuncs.alibaba`, the POP gateway replaces a trailing `.aliyuncs.com` with `.aliyuncs.alibaba`. This applies to generated endpoints, `Config.endpoint`, request `endpointOverride`, and matching `endpointMap` entries, including internal and VPC endpoints. Endpoints ending in `.aliyun.com`, custom domains, and already switched endpoints are preserved. The default suffix, `aliyuncs.com`, leaves endpoints unchanged.

## Issues

[Opening an Issue](https://github.com/aliyun/alibabacloud-gateway/issues/new), Issues not conforming to the guidelines may be closed immediately.

## References

- [Latest Release](https://github.com/aliyun/alibabacloud-gateway)

## License

[Apache-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Copyright (c) 2009-present, Alibaba Cloud All rights reserved.
