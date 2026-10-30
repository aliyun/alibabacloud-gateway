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
            .package(url: "https://github.com/aliyun/tea-swift.git", from: "1.1.0"),
            // Published SPI 0.0.3 marks the interceptor methods public, so another
            // module cannot override them. GatewaySPI is a link to the in-repo SPI,
            // where those methods are open. Switch back to the versioned GitHub
            // package after that SPI release.
            .package(path: "../GatewaySPI"),
            .package(url: "https://github.com/aliyun/credentials-swift", from: "1.0.1"),
            .package(url: "https://github.com/krzyzanowskim/CryptoSwift.git", from: "1.10.0"),
        ],
        targets: [
            .target(
                    name: "AlibabacloudGatewayPOP",
                    dependencies: [
                        .product(name: "Tea", package: "tea-swift"),
                        .product(name: "AlibabacloudGatewaySPI", package: "GatewaySPI"),
                        .product(name: "AlibabaCloudCredentials", package: "credentials-swift"),
                        .product(name: "CryptoSwift", package: "CryptoSwift"),
                    ]),
            .testTarget(
                    name: "AlibabacloudGatewayPOPTests",
                    dependencies: ["AlibabacloudGatewayPOP"]),
        ],
        swiftLanguageVersions: [.v5]
)
