import XCTest
@testable import AlibabacloudGatewayPOP

final class ClientTests: XCTestCase {
    func testCanonicalQueryAndMixedCaseHeaders() throws {
        let client = try Client()
        XCTAssertEqual(
            client.buildCanonicalizedResource(["a-b": "1", "a/b": "2", "a": "",
                                               "é": "中", "punct": "!+*~", "Signature": "ignored"]),
            "%C3%A9=%E4%B8%AD&a=&a%2Fb=2&a-b=1&punct=%21%2B%2A~"
        )
        XCTAssertEqual(
            client.buildCanonicalizedHeaders(["Host": " example.com ", "x-acs-date": "date", "X-Acs-Foo": " foo "]),
            "host:example.com\nx-acs-date:date\nx-acs-foo:foo\n"
        )
    }

    func testSignedHeadersKeepPrefixPairs() throws {
        let client = try Client()
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

    func testSha256AndSm3MatchKnownDigests() {
        XCTAssertEqual(
            Runtime.hexEncode(Runtime.hash(Array("hello".utf8), "ACS4-HMAC-SHA256")),
            "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"
        )
        XCTAssertEqual(
            Runtime.hexEncode(Runtime.hash(Array("abc".utf8), "ACS4-HMAC-SM3")),
            "66c7f0f462eeedd9d1f2d46bdc10e4e24167c4875cf2f7a2297da02b8f4ba8e0"
        )
        XCTAssertEqual(Runtime.percentEncode("a b*~"), "a%20b%2A~")
    }

    func testValidateSignedHeaders() throws {
        let client = try Client()
        try client.validateSignedHeaders(["content-type", "host", "x-acs-date"])
        XCTAssertThrowsError(try client.validateSignedHeaders(["host"]))
        XCTAssertThrowsError(try client.validateSignedHeaders(["x-acs-date"]))
        XCTAssertThrowsError(try client.validateSignedHeaders([]))
    }
}
