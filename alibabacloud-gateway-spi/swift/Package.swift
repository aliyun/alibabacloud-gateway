// swift-tools-version:5.6
// The swift-tools-version declares the minimum version of Swift required to build this package.
import PackageDescription

let package = Package(
        name: "AlibabacloudGatewaySPI",
        platforms: [.macOS(.v10_15),
                    .iOS(.v13),
                    .tvOS(.v13),
                    .watchOS(.v6)],
        products: [
            .library(
                    name: "AlibabacloudGatewaySPI",
                    targets: ["AlibabacloudGatewaySPI"])
        ],
        dependencies: [
            // Dependencies declare other packages that this package depends on.
            .package(url: "https://github.com/aliyun/tea-swift.git", from: "1.0.3"),
            .package(url: "https://github.com/aliyun/credentials-swift.git", from: "1.0.1"),
        ],
        targets: [
            .target(
                    name: "AlibabacloudGatewaySPI",
                    dependencies: [
                        .product(name: "Tea", package: "tea-swift"),
                        .product(name: "AlibabaCloudCredentials", package: "credentials-swift")
                    ]),
        ],
        swiftLanguageVersions: [.v5]
)
