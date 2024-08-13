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
            .package(url: "https://github.com/aliyun/tea-swift.git", from: "1.0.2"),
            .package(url: "https://github.com/alibabacloud-sdk-swift/alibabacloud-gateway-spi", from: "0.0.1"),
        ],
        targets: [
            .target(
                    name: "AlibabacloudGatewayPOP",
                    dependencies: [
                        .product(name: "Tea", package: "tea-swift"),
                        .product(name: "AlibabacloudGatewaySPI", package: "alibabacloud-gateway-spi")
                    ]),
        ],
        swiftLanguageVersions: [.v5]
)
