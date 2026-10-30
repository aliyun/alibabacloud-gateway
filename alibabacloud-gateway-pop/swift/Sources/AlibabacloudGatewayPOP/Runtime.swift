import Foundation
import CryptoSwift
import Tea

/// Gateway-private helpers. Not part of the supported public API.
enum Runtime {
    static func isUnset(_ value: Any?) -> Bool {
        value == nil
    }

    static func empty(_ val: String?) -> Bool {
        val == nil || val!.isEmpty
    }

    static func equalString(_ left: String?, _ right: String?) -> Bool {
        left == right
    }

    static func equalNumber(_ left: Int?, _ right: Int?) -> Bool {
        left == right
    }

    static func defaultString(_ real: String?, _ fallback: String?) -> String {
        if real != nil && !real!.isEmpty {
            return real!
        }
        return fallback ?? ""
    }

    static func toBytes(_ str: String?) -> [UInt8] {
        guard let str = str else {
            return []
        }
        return [UInt8](str.utf8)
    }

    static func toJSONString(_ obj: Any?) -> String {
        if obj == nil {
            return ""
        }
        if obj is [UInt8] {
            return String(bytes: obj as! [UInt8], encoding: .utf8) ?? ""
        }
        if !JSONSerialization.isValidJSONObject(obj!) {
            return "\(obj!)"
        }
        let data = try! JSONSerialization.data(withJSONObject: obj!, options: [])
        return String(data: data, encoding: .utf8) ?? ""
    }

    static func readAsBytes(_ stream: InputStream?) async throws -> [UInt8] {
        stream?.open()
        defer { stream?.close() }
        let bufferSize = 1024
        var buffer = [UInt8](repeating: 0, count: bufferSize)
        var data = Data()
        while stream != nil && stream!.hasBytesAvailable {
            let read = stream!.read(&buffer, maxLength: bufferSize)
            if read < 0 {
                if let error = stream!.streamError {
                    throw error
                }
                break
            }
            if read == 0 {
                break
            }
            data.append(buffer, count: read)
        }
        return [UInt8](data)
    }

    static func readAsString(_ stream: InputStream?) async throws -> String {
        let bytes = try await readAsBytes(stream)
        return String(bytes: bytes, encoding: .utf8) ?? ""
    }

    static func readAsJSON(_ stream: InputStream?) async throws -> Any {
        let str = try await readAsString(stream)
        guard let data = str.data(using: .utf8),
              let obj = try? JSONSerialization.jsonObject(with: data, options: .mutableContainers) as? [String: Any] else {
            return [String: Any]()
        }
        return obj
    }

    static func assertAsBytes(_ any: Any?) throws -> [UInt8] {
        if any is [UInt8] {
            return any as! [UInt8]
        }
        throw TeaError("\(any ?? "") is not a Bytes")
    }

    static func assertAsMap(_ any: Any?) throws -> [String: Any] {
        if let dict = any as? [String: Any] {
            return dict
        }
        throw TeaError("\(any ?? "") is not a map")
    }

    static func is4xx(_ code: Int?) -> Bool {
        code != nil && 400 <= code! && code! < 500
    }

    static func is5xx(_ code: Int?) -> Bool {
        code != nil && 500 <= code! && code! < 600
    }

    static func getNonce() -> String {
        let timestamp = String(Date().timeIntervalSince1970)
        return (timestamp + UUID().uuidString).md5()
    }

    static func getTimestamp() -> String {
        let formatter = DateFormatter()
        formatter.timeZone = TimeZone(identifier: "GMT")
        formatter.dateFormat = "yyyy-MM-dd'T'HH:mm:ss'Z'"
        formatter.locale = Locale(identifier: "en_US_POSIX")
        return formatter.string(from: Date())
    }

    static func getEndpointRules(_ product: String?, _ regionId: String?, _ endpointType: String?, _ network: String?, _ suffix: String?) throws -> String {
        if product == nil || product!.isEmpty {
            throw TeaError("ProductId is empty, please set a valid ProductId")
        }
        var endpoint: String
        if endpointType == "regional" {
            if regionId == nil || regionId!.isEmpty {
                throw TeaError("RegionId is empty, please set a valid RegionId")
            }
            endpoint = "<product><suffix><network>.<region_id>.aliyuncs.com"
            endpoint = endpoint.replacingOccurrences(of: "<region_id>", with: regionId!)
        } else {
            endpoint = "<product><suffix><network>.aliyuncs.com"
        }
        endpoint = endpoint.replacingOccurrences(of: "<product>", with: product!)
        if network == nil || network!.isEmpty || network == "public" {
            endpoint = endpoint.replacingOccurrences(of: "<network>", with: "")
        } else {
            endpoint = endpoint.replacingOccurrences(of: "<network>", with: "-\(network!)")
        }
        if suffix == nil || suffix!.isEmpty {
            endpoint = endpoint.replacingOccurrences(of: "<suffix>", with: "")
        } else {
            endpoint = endpoint.replacingOccurrences(of: "<suffix>", with: "-\(suffix!)")
        }
        return endpoint
    }

    static func toForm(_ filter: [String: Any]?) -> String {
        let dict = query(filter)
        if dict.isEmpty {
            return ""
        }
        let keys = Array(dict.keys).sorted()
        var parts: [String] = []
        for key in keys {
            let value = dict[key]
            if value == nil || value!.isEmpty {
                continue
            }
            parts.append(formEncode(key) + "=" + formEncode(value!))
        }
        return parts.joined(separator: "&")
    }

    static func hexEncode(_ raw: [UInt8]?) -> String {
        raw?.toHexString() ?? ""
    }

    static func hash(_ raw: [UInt8], _ signatureAlgorithm: String) -> [UInt8] {
        if signatureAlgorithm.contains("HMAC-SHA256") || signatureAlgorithm.contains("RSA-SHA256") {
            return Digest.sha2(raw, variant: .sha256)
        }
        if signatureAlgorithm.contains("HMAC-SM3") {
            return Sm3.digest(raw)
        }
        return []
    }

    static func percentEncode(_ raw: String?) -> String {
        guard let raw = raw else {
            return ""
        }
        var allowed = CharacterSet.alphanumerics
        allowed.insert(charactersIn: "-_.!~*'()")
        var encoded = raw.addingPercentEncoding(withAllowedCharacters: allowed) ?? ""
        encoded = encoded.replacingOccurrences(of: "+", with: "%20")
        encoded = encoded.replacingOccurrences(of: "*", with: "%2A")
        encoded = encoded.replacingOccurrences(of: "%7E", with: "~")
        return encoded
    }

    static func HmacSHA256Sign(_ stringToSign: String, _ secret: String) -> [UInt8] {
        HmacSHA256SignByBytes(stringToSign, [UInt8](secret.utf8))
    }

    static func HmacSHA256SignByBytes(_ stringToSign: String, _ secret: [UInt8]) -> [UInt8] {
        try! HMAC(key: secret, variant: .sha2(.sha256)).authenticate([UInt8](stringToSign.utf8))
    }

    static func HmacSM3Sign(_ stringToSign: String, _ secret: String) -> [UInt8] {
        Sm3.hmac([UInt8](stringToSign.utf8), key: [UInt8](secret.utf8))
    }

    static func HmacSM3SignByBytes(_ stringToSign: String, _ secret: [UInt8]) -> [UInt8] {
        Sm3.hmac([UInt8](stringToSign.utf8), key: secret)
    }

    static func subString(_ raw: String?, _ start: Int, _ end: Int) -> String {
        guard let raw = raw else {
            return ""
        }
        if start >= raw.count || end <= start {
            return ""
        }
        let from = raw.index(raw.startIndex, offsetBy: start)
        let to = raw.index(raw.startIndex, offsetBy: min(end, raw.count))
        return String(raw[from..<to])
    }

    static func replace(_ raw: String?, _ old: String, _ new: String, _ count: Int?) -> String {
        guard let raw = raw else {
            return ""
        }
        if count == nil || count! < 0 {
            return raw.replacingOccurrences(of: old, with: new)
        }
        if count == 0 {
            return raw
        }
        var result = ""
        var rest = raw
        var replaced = 0
        while replaced < count! {
            guard let range = rest.range(of: old) else {
                break
            }
            result += rest[..<range.lowerBound]
            result += new
            rest = String(rest[range.upperBound...])
            replaced += 1
        }
        return result + rest
    }

    static func split(_ raw: String?, _ sep: String, _ limit: Int?) -> [String] {
        guard let raw = raw, !raw.isEmpty else {
            return []
        }
        if limit == nil || limit! < 1 {
            return raw.components(separatedBy: sep)
        }
        return raw.components(separatedBy: sep)
    }

    static func contains(_ raw: String?, _ substr: String?) -> Bool {
        guard let raw = raw, let substr = substr else {
            return false
        }
        return raw.contains(substr)
    }

    static func hasPrefix(_ raw: String?, _ prefix: String?) -> Bool {
        guard let raw = raw, let prefix = prefix else {
            return false
        }
        return raw.hasPrefix(prefix)
    }

    static func toLower(_ raw: String?) -> String {
        raw?.lowercased() ?? ""
    }

    static func equals(_ expect: String?, _ actual: String?) -> Bool {
        expect == actual
    }

    static func trim(_ raw: String?) -> String {
        raw?.trimmingCharacters(in: .whitespacesAndNewlines) ?? ""
    }

    static func keySet(_ raw: [String: String]?) -> [String] {
        guard let raw = raw else {
            return []
        }
        return Array(raw.keys)
    }

    static func join(_ raw: [String]?, _ sep: String) -> String {
        raw?.joined(separator: sep) ?? ""
    }

    static func size(_ raw: [String]?) -> Int {
        raw?.count ?? 0
    }

    static func ascSort(_ raw: [String]?) -> [String] {
        raw?.sorted() ?? []
    }

    private static func formEncode(_ raw: String) -> String {
        var allowed = CharacterSet.alphanumerics
        allowed.insert(charactersIn: "*-._")
        allowed.insert(charactersIn: " ")
        var encoded = raw.addingPercentEncoding(withAllowedCharacters: allowed) ?? ""
        encoded = encoded.replacingOccurrences(of: " ", with: "%20")
        return encoded
    }

    private static func query(_ filter: [String: Any]?) -> [String: String] {
        if filter == nil {
            return [:]
        }
        var result: [String: String] = [:]
        parseObject(result: &result, object: filter!, prefix: "")
        return result
    }

    private static func parseObject(result: inout [String: String], object: Any, prefix: String) {
        if object is [Any] {
            for (index, item) in (object as! [Any]).enumerated() {
                parseObject(result: &result, object: item, prefix: "\(prefix).\(index + 1)")
            }
        } else if object is [String: Any] {
            for (key, value) in object as! [String: Any] {
                parseObject(result: &result, object: value, prefix: "\(prefix).\(key)")
            }
        } else if object is TeaModel {
            let dict = (object as! TeaModel).toMap()
            for (key, value) in dict {
                parseObject(result: &result, object: value, prefix: "\(prefix).\(key)")
            }
        } else {
            var key = prefix
            if key.hasPrefix(".") {
                key = String(key[key.index(key.startIndex, offsetBy: 1)...])
            }
            if object is [UInt8] {
                result[key] = String(bytes: object as! [UInt8], encoding: .utf8)
            } else {
                result[key] = "\(object)"
            }
        }
    }
}

private enum Sm3 {
    private static let iv: [UInt32] = [
        0x7380166f, 0x4914b2b9, 0x172442d7, 0xda8a0600,
        0xa96f30bc, 0x163138aa, 0xe38dee4d, 0xb0fb0e4e
    ]

    static func digest(_ message: [UInt8]) -> [UInt8] {
        var msg = message
        let bitLength = UInt64(message.count) * 8
        msg.append(0x80)
        while (msg.count % 64) != 56 {
            msg.append(0)
        }
        var shift = 56
        while shift >= 0 {
            msg.append(UInt8((bitLength >> UInt64(shift)) & 0xff))
            shift -= 8
        }
        var state = iv
        var offset = 0
        while offset < msg.count {
            state = compress(state, Array(msg[offset..<offset + 64]))
            offset += 64
        }
        var out: [UInt8] = []
        for word in state {
            out.append(UInt8((word >> 24) & 0xff))
            out.append(UInt8((word >> 16) & 0xff))
            out.append(UInt8((word >> 8) & 0xff))
            out.append(UInt8(word & 0xff))
        }
        return out
    }

    static func hmac(_ message: [UInt8], key: [UInt8]) -> [UInt8] {
        let block = 64
        var normalized = key
        if normalized.count > block {
            normalized = digest(normalized)
        }
        if normalized.count < block {
            normalized.append(contentsOf: [UInt8](repeating: 0, count: block - normalized.count))
        }
        var ipad = [UInt8](repeating: 0x36, count: block)
        var opad = [UInt8](repeating: 0x5c, count: block)
        for i in 0..<block {
            ipad[i] ^= normalized[i]
            opad[i] ^= normalized[i]
        }
        return digest(opad + digest(ipad + message))
    }

    private static func rotl(_ x: UInt32, _ k: UInt32) -> UInt32 {
        let n = k % 32
        if n == 0 {
            return x
        }
        return (x << n) | (x >> (32 - n))
    }

    private static func ff(_ x: UInt32, _ y: UInt32, _ z: UInt32, _ j: Int) -> UInt32 {
        if j < 16 {
            return x ^ y ^ z
        }
        return (x & y) | (x & z) | (y & z)
    }

    private static func gg(_ x: UInt32, _ y: UInt32, _ z: UInt32, _ j: Int) -> UInt32 {
        if j < 16 {
            return x ^ y ^ z
        }
        return (x & y) | ((~x) & z)
    }

    private static func p0(_ x: UInt32) -> UInt32 {
        x ^ rotl(x, 9) ^ rotl(x, 17)
    }

    private static func p1(_ x: UInt32) -> UInt32 {
        x ^ rotl(x, 15) ^ rotl(x, 23)
    }

    private static func compress(_ state: [UInt32], _ block: [UInt8]) -> [UInt32] {
        var w = [UInt32](repeating: 0, count: 68)
        for i in 0..<16 {
            let b = i * 4
            w[i] = (UInt32(block[b]) << 24) | (UInt32(block[b + 1]) << 16) | (UInt32(block[b + 2]) << 8) | UInt32(block[b + 3])
        }
        for j in 16..<68 {
            w[j] = p1(w[j - 16] ^ w[j - 9] ^ rotl(w[j - 3], 15)) ^ rotl(w[j - 13], 7) ^ w[j - 6]
        }
        var w1 = [UInt32](repeating: 0, count: 64)
        for j in 0..<64 {
            w1[j] = w[j] ^ w[j + 4]
        }
        var a = state[0]
        var b = state[1]
        var c = state[2]
        var d = state[3]
        var e = state[4]
        var f = state[5]
        var g = state[6]
        var h = state[7]
        for j in 0..<64 {
            let constant: UInt32 = j < 16 ? 0x79cc4519 : 0x7a879d8a
            let ss1 = rotl(rotl(a, 12) &+ e &+ rotl(constant, UInt32(j)), 7)
            let ss2 = ss1 ^ rotl(a, 12)
            let tt1 = ff(a, b, c, j) &+ d &+ ss2 &+ w1[j]
            let tt2 = gg(e, f, g, j) &+ h &+ ss1 &+ w[j]
            d = c
            c = rotl(b, 9)
            b = a
            a = tt1
            h = g
            g = rotl(f, 19)
            f = e
            e = p0(tt2)
        }
        return [
            a ^ state[0], b ^ state[1], c ^ state[2], d ^ state[3],
            e ^ state[4], f ^ state[5], g ^ state[6], h ^ state[7]
        ]
    }
}
