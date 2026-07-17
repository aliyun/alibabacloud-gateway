import XCTest
@testable import AlibabacloudGatewayPOP

final class ClientTests: XCTestCase {
    func testSignedHeadersKeepPrefixPairs() {
        let client = Client()
        let headers = [
            "host": "example.com",
            "x-acs-foobar": "bar",
            "x-acs-foo": "foo",
        ]

        XCTAssertFalse(["x-acs-foobar"].contains("x-acs-foo"))
        XCTAssertEqual(
            client.getSignedHeaders(headers),
            ["host", "x-acs-foo", "x-acs-foobar"]
        )

        let canonical = client.buildCanonicalizedHeaders(headers)
        XCTAssertTrue(canonical.contains("host:example.com\n"))
        XCTAssertTrue(canonical.contains("x-acs-foo:foo\n"))
        XCTAssertTrue(canonical.contains("x-acs-foobar:bar\n"))
    }
}
