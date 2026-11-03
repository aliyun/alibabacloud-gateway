import Foundation
import Tea
import AlibabacloudGatewaySPI
import AlibabaCloudCredentials

open class Client : AlibabacloudGatewaySPI.Client {
    public var _endpointSuffix: String?

    public var _signatureTypePrefix: String?

    public var _signPrefix: String?

    public var _sha256: String?

    public var _sm3: String?

    public override init() throws {
        try super.init()
        self._signatureTypePrefix = "ACS4-"
        self._signPrefix = "aliyun_v4"
        self._endpointSuffix = "aliyuncs.com"
        self._sha256 = (self._signatureTypePrefix ?? "") + "HMAC-SHA256"
        self._sm3 = (self._signatureTypePrefix ?? "") + "HMAC-SM3"
    }

    @available(macOS 10.15, iOS 13, tvOS 13, watchOS 6, *)
    public override func modifyConfiguration(_ context: AlibabacloudGatewaySPI.InterceptorContext, _ attributeMap: AlibabacloudGatewaySPI.AttributeMap) async throws -> Void {
        var request: AlibabacloudGatewaySPI.InterceptorContext.Request = context.request!
        var config: AlibabacloudGatewaySPI.InterceptorContext.Configuration = context.configuration!
        var attributes: [String: String] = attributeMap.key ?? [:]
        if (!Runtime.isUnset(attributes)) {
            self._signatureTypePrefix = attributes["signatureTypePrefix"]
            self._signPrefix = attributes["signPrefix"]
            self._endpointSuffix = attributes["endpointSuffix"]
            self._sha256 = (self._signatureTypePrefix ?? "") + "HMAC-SHA256"
            self._sm3 = (self._signatureTypePrefix ?? "") + "HMAC-SM3"
        }
        config.endpoint = try getEndpoint(request.productId ?? "", config.regionId ?? "", config.endpointRule ?? "", config.network ?? "", config.suffix ?? "", config.endpointMap ?? [:], config.endpoint ?? "")
    }

    @available(macOS 10.15, iOS 13, tvOS 13, watchOS 6, *)
    public override func modifyRequest(_ context: AlibabacloudGatewaySPI.InterceptorContext, _ attributeMap: AlibabacloudGatewaySPI.AttributeMap) async throws -> Void {
        var request: AlibabacloudGatewaySPI.InterceptorContext.Request = context.request!
        var config: AlibabacloudGatewaySPI.InterceptorContext.Configuration = context.configuration!
        var date: String = Runtime.getTimestamp()
        request.headers = Tea.TeaConverter.merge([
            "host": config.endpoint ?? "",
            "x-acs-version": request.version ?? "",
            "user-agent": request.userAgent ?? "",
            "x-acs-date": date as! String,
            "x-acs-signature-nonce": Runtime.getNonce(),
            "accept": "application/json"
        ], request.headers ?? [:])
        if (!Runtime.empty(request.action)) {
            request.headers!["x-acs-action"] = request.action;
        }
        var signatureAlgorithm: String = Runtime.defaultString(request.signatureAlgorithm, self._sha256)
        var hashedRequestPayload: String = Runtime.hexEncode(Runtime.hash(Runtime.toBytes(""), signatureAlgorithm))
        if (!Runtime.isUnset(request.stream)) {
            var tmp: [UInt8] = try await Runtime.readAsBytes(request.stream)
            hashedRequestPayload = Runtime.hexEncode(Runtime.hash(tmp, signatureAlgorithm))
            request.stream = Tea.TeaCore.toReadable(tmp)
            request.headers!["content-type"] = "application/octet-stream";
        }
        else {
            if (!Runtime.isUnset(request.body)) {
                if (Runtime.equalString(request.reqBodyType, "byte")) {
                    var byteObj: [UInt8] = try Runtime.assertAsBytes(request.body)
                    hashedRequestPayload = Runtime.hexEncode(Runtime.hash(byteObj, signatureAlgorithm))
                    request.stream = Tea.TeaCore.toReadable(byteObj)
                }
                else if (Runtime.equalString(request.reqBodyType, "json")) {
                    var jsonObj: String = Runtime.toJSONString(request.body)
                    hashedRequestPayload = Runtime.hexEncode(Runtime.hash(Runtime.toBytes(jsonObj), signatureAlgorithm))
                    request.stream = Tea.TeaCore.toReadable(jsonObj)
                    request.headers!["content-type"] = "application/json; charset=utf-8";
                }
                else {
                    var m: [String: Any] = try Runtime.assertAsMap(request.body)
                    var formObj: String = Runtime.toForm(m)
                    hashedRequestPayload = Runtime.hexEncode(Runtime.hash(Runtime.toBytes(formObj), signatureAlgorithm))
                    request.stream = Tea.TeaCore.toReadable(formObj)
                    request.headers!["content-type"] = "application/x-www-form-urlencoded";
                }
            }
        }
        if (Runtime.equalString(signatureAlgorithm, self._sm3)) {
            request.headers!["x-acs-content-sm3"] = hashedRequestPayload;
        }
        else {
            request.headers!["x-acs-content-sha256"] = hashedRequestPayload;
        }
        if (!Runtime.equalString(request.authType, "Anonymous")) {
            var credential: AlibabaCloudCredentials.Client = request.credential!
            if (Runtime.isUnset(credential)) {
                throw Tea.ReuqestError([
                    "code": "ParameterMissing",
                    "message": "'config.credential' can not be unset"
                ])
            }
            var credentialModel: AlibabaCloudCredentials.Credential? = try await credential.getCredential()
            var authType: String = credentialModel?.getType() ?? ""
            if (Runtime.equalString(authType, "bearer")) {
                var bearerToken: String = credentialModel?.getBearerToken() ?? ""
                request.headers!["x-acs-bearer-token"] = bearerToken;
                request.headers!["x-acs-signature-type"] = "BEARERTOKEN";
                request.headers!["Authorization"] = "Bearer " + (bearerToken as! String);
            }
            else if (Runtime.equalString(authType, "id_token")) {
                var idToken: String = credentialModel?.getSecurityToken() ?? ""
                request.headers!["x-acs-zero-trust-idtoken"] = idToken;
            }
            else {
                var accessKeyId: String = credentialModel?.getAccessKeyId() ?? ""
                var accessKeySecret: String = credentialModel?.getAccessKeySecret() ?? ""
                var securityToken: String = credentialModel?.getSecurityToken() ?? ""
                if (!Runtime.empty(securityToken)) {
                    request.headers!["x-acs-accesskey-id"] = accessKeyId;
                    request.headers!["x-acs-security-token"] = securityToken;
                }
                var dateNew: String = Runtime.subString(date, 0, 10)
                dateNew = Runtime.replace(dateNew, "-", "", nil)
                var region: String = getRegion(request.productId ?? "", config.endpoint ?? "", config.regionId ?? "")
                var signingkey: [UInt8] = getSigningkey(signatureAlgorithm as! String, accessKeySecret as! String, request.productId ?? "", region as! String, dateNew as! String)
                request.headers!["Authorization"] = try getAuthorization(request.pathname ?? "", request.method ?? "", request.query ?? [:], request.headers ?? [:], signatureAlgorithm as! String, hashedRequestPayload as! String, accessKeyId as! String, signingkey as! [UInt8], request.productId ?? "", region as! String, dateNew as! String);
            }
        }
    }

    @available(macOS 10.15, iOS 13, tvOS 13, watchOS 6, *)
    public override func modifyResponse(_ context: AlibabacloudGatewaySPI.InterceptorContext, _ attributeMap: AlibabacloudGatewaySPI.AttributeMap) async throws -> Void {
        var request: AlibabacloudGatewaySPI.InterceptorContext.Request = context.request!
        var response: AlibabacloudGatewaySPI.InterceptorContext.Response = context.response!
        if (Runtime.is4xx(response.statusCode) || Runtime.is5xx(response.statusCode)) {
            var _res: Any = try await Runtime.readAsJSON(response.body)
            var err: [String: Any] = try Runtime.assertAsMap(_res)
            var requestId: Any = defaultAny(err["RequestId"]!, err["requestId"]!)
            if (!Runtime.isUnset(response.headers!["x-acs-request-id"])) {
                requestId = response.headers!["x-acs-request-id"] ?? ""
            }
            err["statusCode"] = response.statusCode
            let errorCode = String(describing: defaultAny(err["Code"]!, err["code"]!))
            let errorMessage = String(describing: defaultAny(err["Message"]!, err["message"]!))
            let errorDescription = String(describing: defaultAny(err["Description"]!, err["description"]!))
            let accessDeniedDetail = defaultAny(err["AccessDeniedDetail"]!, err["accessDeniedDetail"]!)
            let requestIdText = String(describing: requestId)
            throw Tea.ReuqestError([
                "code": errorCode,
                "message": "code: " + String(response.statusCode!) + ", " + errorMessage + " request id: " + requestIdText,
                "data": err,
                "description": errorDescription,
                "accessDeniedDetail": accessDeniedDetail
            ])
        }
        if (Runtime.equalNumber(response.statusCode, 204)) {
            try await Runtime.readAsString(response.body)
        }
        else if (Runtime.equalString(request.bodyType, "binary")) {
            response.deserializedBody = response.body
        }
        else if (Runtime.equalString(request.bodyType, "byte")) {
            var byt: [UInt8] = try await Runtime.readAsBytes(response.body)
            response.deserializedBody = byt
        }
        else if (Runtime.equalString(request.bodyType, "string")) {
            var str: String = try await Runtime.readAsString(response.body)
            response.deserializedBody = str
        }
        else if (Runtime.equalString(request.bodyType, "json")) {
            var obj: Any = try await Runtime.readAsJSON(response.body)
            var res: [String: Any] = try Runtime.assertAsMap(obj)
            response.deserializedBody = res
        }
        else if (Runtime.equalString(request.bodyType, "array")) {
            var arr: Any = try await Runtime.readAsJSON(response.body)
            response.deserializedBody = arr
        }
        else {
            response.deserializedBody = try await Runtime.readAsString(response.body)
        }
    }

    public func getEndpoint(_ productId: String, _ regionId: String, _ endpointRule: String, _ network: String, _ suffix: String, _ endpointMap: [String: String], _ endpoint: String) throws -> String {
        if (!Runtime.empty(endpoint)) {
            return endpoint as! String
        }
        if (!Runtime.isUnset(endpointMap) && !Runtime.empty(endpointMap[regionId as! String])) {
            return endpointMap[regionId as! String] ?? ""
        }
        return try Runtime.getEndpointRules(productId, regionId, endpointRule, network, suffix)
    }

    public func defaultAny(_ inputValue: Any, _ defaultValue: Any) -> Any {
        if (Runtime.isUnset(inputValue)) {
            return defaultValue as! Any
        }
        return inputValue as! Any
    }

    public func getAuthorization(_ pathname: String, _ method: String, _ query: [String: String], _ headers: [String: String], _ signatureAlgorithm: String, _ payload: String, _ ak: String, _ signingkey: [UInt8], _ product: String, _ region: String, _ date: String) throws -> String {
        var signature: String = try getSignature(pathname as! String, method as! String, query as! [String: String], headers as! [String: String], signatureAlgorithm as! String, payload as! String, signingkey as! [UInt8])
        var signedHeaders: [String] = getSignedHeaders(headers as! [String: String])
        var signedHeadersStr: String = Runtime.join(signedHeaders, ";")
        return (signatureAlgorithm as! String) + " Credential=" + (ak as! String) + "/" + (date as! String) + "/" + (region as! String) + "/" + (product as! String) + "/" + (self._signPrefix ?? "") + "_request,SignedHeaders=" + (signedHeadersStr as! String) + ",Signature=" + (signature as! String)
    }

    public func getSignature(_ pathname: String, _ method: String, _ query: [String: String], _ headers: [String: String], _ signatureAlgorithm: String, _ payload: String, _ signingkey: [UInt8]) throws -> String {
        var canonicalURI: String = "/"
        if (!Runtime.empty(pathname)) {
            canonicalURI = pathname as! String
        }
        var stringToSign: String = ""
        var canonicalizedResource: String = buildCanonicalizedResource(query as! [String: String])
        var canonicalizedHeaders: String = buildCanonicalizedHeaders(headers as! [String: String])
        var signedHeaders: [String] = getSignedHeaders(headers as! [String: String])
        try validateSignedHeaders(signedHeaders)
        var signedHeadersStr: String = Runtime.join(signedHeaders, ";")
        stringToSign = (method as! String) + "\n" + (canonicalURI as! String) + "\n" + (canonicalizedResource as! String) + "\n" + (canonicalizedHeaders as! String) + "\n" + (signedHeadersStr as! String) + "\n" + (payload as! String)
        var hex: String = Runtime.hexEncode(Runtime.hash(Runtime.toBytes(stringToSign), signatureAlgorithm))
        stringToSign = (signatureAlgorithm as! String) + "\n" + (hex as! String)
        var signature: [UInt8] = Runtime.toBytes("")
        if (Runtime.equalString(signatureAlgorithm, self._sha256)) {
            signature = Runtime.HmacSHA256SignByBytes(stringToSign, signingkey)
        }
        else if (Runtime.equalString(signatureAlgorithm, self._sm3)) {
            signature = Runtime.HmacSM3SignByBytes(stringToSign, signingkey)
        }
        return Runtime.hexEncode(signature)
    }

    public func getSigningkey(_ signatureAlgorithm: String, _ secret: String, _ product: String, _ region: String, _ date: String) -> [UInt8] {
        var sc1: String = (self._signPrefix ?? "") + (secret as! String)
        var sc2: [UInt8] = Runtime.toBytes("")
        if (Runtime.equalString(signatureAlgorithm, self._sha256)) {
            sc2 = Runtime.HmacSHA256Sign(date, sc1)
        }
        else if (Runtime.equalString(signatureAlgorithm, self._sm3)) {
            sc2 = Runtime.HmacSM3Sign(date, sc1)
        }
        var sc3: [UInt8] = Runtime.toBytes("")
        if (Runtime.equalString(signatureAlgorithm, self._sha256)) {
            sc3 = Runtime.HmacSHA256SignByBytes(region, sc2)
        }
        else if (Runtime.equalString(signatureAlgorithm, self._sm3)) {
            sc3 = Runtime.HmacSM3SignByBytes(region, sc2)
        }
        var sc4: [UInt8] = Runtime.toBytes("")
        if (Runtime.equalString(signatureAlgorithm, self._sha256)) {
            sc4 = Runtime.HmacSHA256SignByBytes(product, sc3)
        }
        else if (Runtime.equalString(signatureAlgorithm, self._sm3)) {
            sc4 = Runtime.HmacSM3SignByBytes(product, sc3)
        }
        var hmac: [UInt8] = Runtime.toBytes("")
        if (Runtime.equalString(signatureAlgorithm, self._sha256)) {
            hmac = Runtime.HmacSHA256SignByBytes((self._signPrefix ?? "") + "_request", sc4)
        }
        else if (Runtime.equalString(signatureAlgorithm, self._sm3)) {
            hmac = Runtime.HmacSM3SignByBytes((self._signPrefix ?? "") + "_request", sc4)
        }
        return hmac as! [UInt8]
    }

    public func getRegion(_ product: String, _ endpoint: String, _ regionId: String) -> String {
        if (!Runtime.empty(regionId)) {
            return regionId as! String
        }
        var region: String = "center"
        if (Runtime.empty(product) || Runtime.empty(endpoint)) {
            return region as! String
        }
        var strs: [String] = Runtime.split(endpoint, ":", nil)
        var withoutPort: String = strs[0]
        var preRegion: String = Runtime.replace(withoutPort, "." + (self._endpointSuffix ?? ""), "", nil)
        var nodes: [String] = Runtime.split(preRegion, ".", nil)
        if (Runtime.equalNumber(Runtime.size(nodes), 2)) {
            region = nodes[1]
        }
        return region as! String
    }

    public func buildCanonicalizedResource(_ query: [String: String]) -> String {
        var canonicalizedResource: String = ""
        if (!Runtime.isUnset(query)) {
            var queryArray: [String] = Runtime.keySet(query)
            var sortedQueryArray: [String] = Runtime.ascSort(queryArray)
            var separator: String = ""
            for key in sortedQueryArray {
                canonicalizedResource = (canonicalizedResource as! String) + (separator as! String) + (Runtime.percentEncode(key)) + "="
                if (!Runtime.empty(query[key as! String])) {
                    canonicalizedResource = (canonicalizedResource as! String) + (Runtime.percentEncode(query[key as! String]))
                }
                separator = "&"
            }
        }
        return canonicalizedResource as! String
    }

    public func buildCanonicalizedHeaders(_ headers: [String: String]) -> String {
        var canonicalizedHeaders: String = ""
        var sortedHeaders: [String] = getSignedHeaders(headers as! [String: String])
        for header in sortedHeaders {
            canonicalizedHeaders = (canonicalizedHeaders as! String) + (header as! String) + ":" + (Runtime.trim(headers[header as! String])) + "\n"
        }
        return canonicalizedHeaders as! String
    }

    public func validateSignedHeaders(_ signedHeaders: [String]) throws {
        if (!signedHeaders.contains("host") || !signedHeaders.contains("x-acs-date")) {
            throw Tea.ReuqestError([
                "code": "InvalidSignedHeaders",
                "message": "signed headers must include host and x-acs-date"
            ])
        }
    }

    public func getSignedHeaders(_ headers: [String: String]) -> [String] {
        var headersArray: [String] = Runtime.keySet(headers)
        var sortedHeadersArray: [String] = Runtime.ascSort(headersArray)
        var result: [String] = []
        for key in sortedHeadersArray {
            var lowerKey: String = Runtime.toLower(key)
            if (Runtime.hasPrefix(lowerKey, "x-acs-") || Runtime.equals(lowerKey, "host") || Runtime.equals(lowerKey, "content-type")) {
                if (!result.contains(lowerKey)) {
                    result.append(lowerKey)
                }
            }
        }
        return result
    }
}
