// swift-tools-version:5.6
// The swift-tools-version declares the minimum version of Swift required to build this package.
import PackageDescription

let package = Package(
        name: "AlibabacloudGatewayPOP",
        platforms: [.macOS(.v10_15),
                    .iOS(.v13),
                    .tvOS(.v13),
                    .watchOS(.v6)],
        products: [
            .library(
                    name: "AlibabacloudGatewayPOP",
                    targets: ["AlibabacloudGatewayPOP"])
        ],
        dependencies: [
            // Dependencies declare other packages that this package depends on.
            .package(url: "https://github.com/aliyun/tea-swift.git", from: "1.0.3"),
            .package(url: "https://github.com/alibabacloud-sdk-swift/alibabacloud-gateway-spi", from: "0.0.2"),
            .package(url: "https://github.com/aliyun/credentials-swift", from: "1.0.1"),
            .package(url: "https://github.com/alibabacloud-sdk-swift/tea-utils", from: "1.0.6"),
            .package(url: "https://github.com/alibabacloud-sdk-swift/openapi-util", from: "1.0.1"),
            .package(url: "https://github.com/alibabacloud-sdk-swift/endpoint-util", from: "1.0.0"),
            .package(url: "https://github.com/darabonba/darabonba_EncodeUtil.git", from: "0.0.6"),
            .package(url: "https://github.com/darabonba/darabonba_SignatureUtil.git", from: "0.0.10"),
            .package(url: "https://github.com/darabonba/darabonba_String.git", from: "0.0.13"),
            .package(url: "https://github.com/darabonba/darabonba_Map.git", from: "0.0.5"),
            .package(url: "https://github.com/darabonba/darabonba_Array.git", from: "0.1.0"),
        ],
        targets: [
            .target(
                    name: "AlibabacloudGatewayPOP",
                    dependencies: [
                        .product(name: "Tea", package: "tea-swift"),
                        .product(name: "AlibabacloudGatewaySPI", package: "alibabacloud-gateway-spi"),
                        .product(name: "AlibabaCloudCredentials", package: "credentials-swift"),
                        .product(name: "TeaUtils", package: "tea-utils"),
                        .product(name: "AlibabaCloudOpenApiUtil", package: "openapi-util"),
                        .product(name: "AlibabacloudEndpointUtil", package: "endpoint-util"),
                        .product(name: "darabonba_EncodeUtil", package: "darabonba_EncodeUtil"),
                        .product(name: "darabonba_SignatureUtil", package: "darabonba_SignatureUtil"),
                        .product(name: "darabonba_String", package: "darabonba_String"),
                        .product(name: "darabonba_Map", package: "darabonba_Map"),
                        .product(name: "darabonba_Array", package: "darabonba_Array")
                    ]),
        ],
        swiftLanguageVersions: [.v5]
)
