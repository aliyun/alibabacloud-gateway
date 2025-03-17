import Foundation
import Tea
import AlibabacloudGatewaySPI
import AlibabaCloudCredentials
import TeaUtils
import AlibabaCloudOpenApiUtil
import AlibabacloudEndpointUtil
import DarabonbaEncodeUtil
import DarabonbaSignatureUtil
import DarabonbaString
import DarabonbaMap
import DarabonbaArray

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
    public func modifyConfiguration(_ context: AlibabacloudGatewaySPI.InterceptorContext, _ attributeMap: AlibabacloudGatewaySPI.AttributeMap) async throws -> Void {
        var request: AlibabacloudGatewaySPI.InterceptorContext.Request = context.request!
        var config: AlibabacloudGatewaySPI.InterceptorContext.Configuration = context.configuration!
        var attributes: String = attributeMap.key ?? [:]
        if (!TeaUtils.Client.isUnset(attributes)) {
            self._signatureTypePrefix = attributes["signatureTypePrefix"]
            self._signPrefix = attributes["signPrefix"]
            self._endpointSuffix = attributes["endpointSuffix"]
            self._sha256 = (self._signatureTypePrefix ?? "") + "HMAC-SHA256"
            self._sm3 = (self._signatureTypePrefix ?? "") + "HMAC-SM3"
        }
        config.endpoint = try getEndpoint(request.productId ?? "", config.regionId ?? "", config.endpointRule ?? "", config.network ?? "", config.suffix ?? "", config.endpointMap ?? [:], config.endpoint ?? "")
    }

    @available(macOS 10.15, iOS 13, tvOS 13, watchOS 6, *)
    public func modifyRequest(_ context: AlibabacloudGatewaySPI.InterceptorContext, _ attributeMap: AlibabacloudGatewaySPI.AttributeMap) async throws -> Void {
        var request: AlibabacloudGatewaySPI.InterceptorContext.Request = context.request!
        var config: AlibabacloudGatewaySPI.InterceptorContext.Configuration = context.configuration!
        var date: String = AlibabaCloudOpenApiUtil.Client.getTimestamp()
        request.headers = Tea.TeaConverter.merge([
            "host": config.endpoint ?? "",
            "x-acs-version": request.version ?? "",
            "x-acs-action": request.action ?? "",
            "user-agent": request.userAgent ?? "",
            "x-acs-date": date as! String,
            "x-acs-signature-nonce": TeaUtils.Client.getNonce(),
            "accept": "application/json"
        ], request.headers ?? [:])
        var signatureAlgorithm: String = TeaUtils.Client.defaultString(request.signatureAlgorithm, self._sha256)
        var hashedRequestPayload: String = DarabonbaEncodeUtil.Client.hexEncode(DarabonbaEncodeUtil.Client.hash(TeaUtils.Client.toBytes(""), signatureAlgorithm))
        if (!TeaUtils.Client.isUnset(request.stream)) {
            var tmp: [UInt8] = try await TeaUtils.Client.readAsBytes(request.stream)
            hashedRequestPayload = DarabonbaEncodeUtil.Client.hexEncode(DarabonbaEncodeUtil.Client.hash(tmp, signatureAlgorithm))
            request.stream = Tea.TeaCore.toReadable(tmp)
            request.headers!["content-type"] = "application/octet-stream";
        }
        else {
            if (!TeaUtils.Client.isUnset(request.body)) {
                if (TeaUtils.Client.equalString(request.reqBodyType, "json")) {
                    var jsonObj: String = TeaUtils.Client.toJSONString(request.body)
                    hashedRequestPayload = DarabonbaEncodeUtil.Client.hexEncode(DarabonbaEncodeUtil.Client.hash(TeaUtils.Client.toBytes(jsonObj), signatureAlgorithm))
                    request.stream = Tea.TeaCore.toReadable(jsonObj)
                    request.headers!["content-type"] = "application/json; charset=utf-8";
                }
                else {
                    var m: [String: Any] = try TeaUtils.Client.assertAsMap(request.body)
                    var formObj: String = AlibabaCloudOpenApiUtil.Client.toForm(m)
                    hashedRequestPayload = DarabonbaEncodeUtil.Client.hexEncode(DarabonbaEncodeUtil.Client.hash(TeaUtils.Client.toBytes(formObj), signatureAlgorithm))
                    request.stream = Tea.TeaCore.toReadable(formObj)
                    request.headers!["content-type"] = "application/x-www-form-urlencoded";
                }
            }
        }
        if (TeaUtils.Client.equalString(signatureAlgorithm, self._sm3)) {
            request.headers!["x-acs-content-sm3"] = hashedRequestPayload;
        }
        else {
            request.headers!["x-acs-content-sha256"] = hashedRequestPayload;
        }
        if (!TeaUtils.Client.equalString(request.authType, "Anonymous")) {
            var credential: AlibabaCloudCredentials.Client = request.credential!
            if (TeaUtils.Client.isUnset(credential)) {
                throw Tea.ReuqestError([
                    "code": "ParameterMissing",
                    "message": "'config.credential' can not be unset"
                ])
            }
            var credentialModel: AlibabaCloudCredentials.CredentialModel = try await credential.getCredential()
            if (!TeaUtils.Client.empty(credentialModel.providerName)) {
                request.headers!["x-acs-credentials-provider"] = credentialModel.providerName;
            }
            var authType: String = credentialModel.type ?? ""
            if (TeaUtils.Client.equalString(authType, "bearer")) {
                var bearerToken: String = credential.getBearerToken()
                request.headers!["x-acs-bearer-token"] = bearerToken;
                request.headers!["x-acs-signature-type"] = "BEARERTOKEN";
                request.headers!["Authorization"] = "Bearer " + (bearerToken as! String);
            }
            else if (TeaUtils.Client.equalString(authType, "id_token")) {
                var idToken: String = credentialModel.securityToken ?? ""
                request.headers!["x-acs-zero-trust-idtoken"] = idToken;
            }
            else {
                var accessKeyId: String = credentialModel.accessKeyId ?? ""
                var accessKeySecret: String = credentialModel.accessKeySecret ?? ""
                var securityToken: String = credentialModel.securityToken ?? ""
                if (!TeaUtils.Client.empty(securityToken)) {
                    request.headers!["x-acs-accesskey-id"] = accessKeyId;
                    request.headers!["x-acs-security-token"] = securityToken;
                }
                var dateNew: String = DarabonbaString.Client.subString(date, 0, 10)
                dateNew = DarabonbaString.Client.replace(dateNew, "-", "", nil)
                var region: String = getRegion(request.productId ?? "", config.endpoint ?? "", config.regionId ?? "")
                var signingkey: [UInt8] = try await getSigningkey(signatureAlgorithm as! String, accessKeySecret as! String, request.productId ?? "", region as! String, dateNew as! String)
                request.headers!["Authorization"] = try await getAuthorization(request.pathname ?? "", request.method ?? "", request.query ?? [:], request.headers ?? [:], signatureAlgorithm as! String, hashedRequestPayload as! String, accessKeyId as! String, signingkey as! [UInt8], request.productId ?? "", region as! String, dateNew as! String);
            }
        }
    }

    @available(macOS 10.15, iOS 13, tvOS 13, watchOS 6, *)
    public func modifyResponse(_ context: AlibabacloudGatewaySPI.InterceptorContext, _ attributeMap: AlibabacloudGatewaySPI.AttributeMap) async throws -> Void {
        var request: AlibabacloudGatewaySPI.InterceptorContext.Request = context.request!
        var response: AlibabacloudGatewaySPI.InterceptorContext.Response = context.response!
        if (TeaUtils.Client.is4xx(response.statusCode) || TeaUtils.Client.is5xx(response.statusCode)) {
            var _res: Any = try await TeaUtils.Client.readAsJSON(response.body)
            var err: [String: Any] = try TeaUtils.Client.assertAsMap(_res)
            var requestId: Any = defaultAny(err["RequestId"]!, err["requestId"]!)
            if (!TeaUtils.Client.isUnset(response.headers!["x-acs-request-id"])) {
                requestId = response.headers!["x-acs-request-id"] ?? ""
            }
            err["statusCode"] = response.statusCode
            throw Tea.ReuqestError([
                "code": String(defaultAny(err["Code"]!, err["code"]!)),
                "message": "code: " + String(response.statusCode!) + ", " + String(defaultAny(err["Message"]!, err["message"]!)) + " request id: " + String(requestId as! Any),
                "data": err as! [String: Any],
                "description": String(defaultAny(err["Description"]!, err["description"]!)),
                "accessDeniedDetail": defaultAny(err["AccessDeniedDetail"]!, err["accessDeniedDetail"]!)
            ])
        }
        if (TeaUtils.Client.equalNumber(response.statusCode, 204)) {
            try await TeaUtils.Client.readAsString(response.body)
        }
        else if (TeaUtils.Client.equalString(request.bodyType, "binary")) {
            response.deserializedBody = response.body
        }
        else if (TeaUtils.Client.equalString(request.bodyType, "byte")) {
            var byt: [UInt8] = try await TeaUtils.Client.readAsBytes(response.body)
            response.deserializedBody = byt
        }
        else if (TeaUtils.Client.equalString(request.bodyType, "string")) {
            var str: String = try await TeaUtils.Client.readAsString(response.body)
            response.deserializedBody = str
        }
        else if (TeaUtils.Client.equalString(request.bodyType, "json")) {
            var obj: Any = try await TeaUtils.Client.readAsJSON(response.body)
            var res: [String: Any] = try TeaUtils.Client.assertAsMap(obj)
            response.deserializedBody = res
        }
        else if (TeaUtils.Client.equalString(request.bodyType, "array")) {
            var arr: Any = try await TeaUtils.Client.readAsJSON(response.body)
            response.deserializedBody = arr
        }
        else {
            response.deserializedBody = try await TeaUtils.Client.readAsString(response.body)
        }
    }

    public func getEndpoint(_ productId: String, _ regionId: String, _ endpointRule: String, _ network: String, _ suffix: String, _ endpointMap: [String: String], _ endpoint: String) throws -> String {
        if (!TeaUtils.Client.empty(endpoint)) {
            return endpoint as! String
        }
        if (!TeaUtils.Client.isUnset(endpointMap) && !TeaUtils.Client.empty(endpointMap[regionId as! String])) {
            return endpointMap[regionId as! String] ?? ""
        }
        return try AlibabacloudEndpointUtil.Client.getEndpointRules(productId, regionId, endpointRule, network, suffix)
    }

    public func defaultAny(_ inputValue: Any, _ defaultValue: Any) -> Any {
        if (TeaUtils.Client.isUnset(inputValue)) {
            return defaultValue as! Any
        }
        return inputValue as! Any
    }

    @available(macOS 10.15, iOS 13, tvOS 13, watchOS 6, *)
    public func getAuthorization(_ pathname: String, _ method: String, _ query: [String: String], _ headers: [String: String], _ signatureAlgorithm: String, _ payload: String, _ ak: String, _ signingkey: [UInt8], _ product: String, _ region: String, _ date: String) async throws -> String {
        var signature: String = try await getSignature(pathname as! String, method as! String, query as! [String: String], headers as! [String: String], signatureAlgorithm as! String, payload as! String, signingkey as! [UInt8])
        var signedHeaders: [String] = try await getSignedHeaders(headers as! [String: String])
        var signedHeadersStr: String = DarabonbaArray.Client.join(signedHeaders, ";")
        return (signatureAlgorithm as! String) + " Credential=" + (ak as! String) + "/" + (date as! String) + "/" + (region as! String) + "/" + (product as! String) + "/" + (self._signPrefix ?? "") + "_request,SignedHeaders=" + (signedHeadersStr as! String) + ",Signature=" + (signature as! String)
    }

    @available(macOS 10.15, iOS 13, tvOS 13, watchOS 6, *)
    public func getSignature(_ pathname: String, _ method: String, _ query: [String: String], _ headers: [String: String], _ signatureAlgorithm: String, _ payload: String, _ signingkey: [UInt8]) async throws -> String {
        var canonicalURI: String = "/"
        if (!TeaUtils.Client.empty(pathname)) {
            canonicalURI = pathname as! String
        }
        var stringToSign: String = ""
        var canonicalizedResource: String = try await buildCanonicalizedResource(query as! [String: String])
        var canonicalizedHeaders: String = try await buildCanonicalizedHeaders(headers as! [String: String])
        var signedHeaders: [String] = try await getSignedHeaders(headers as! [String: String])
        var signedHeadersStr: String = DarabonbaArray.Client.join(signedHeaders, ";")
        stringToSign = (method as! String) + "\n" + (canonicalURI as! String) + "\n" + (canonicalizedResource as! String) + "\n" + (canonicalizedHeaders as! String) + "\n" + (signedHeadersStr as! String) + "\n" + (payload as! String)
        var hex: String = DarabonbaEncodeUtil.Client.hexEncode(DarabonbaEncodeUtil.Client.hash(TeaUtils.Client.toBytes(stringToSign), signatureAlgorithm))
        stringToSign = (signatureAlgorithm as! String) + "\n" + (hex as! String)
        var signature: [UInt8] = TeaUtils.Client.toBytes("")
        if (TeaUtils.Client.equalString(signatureAlgorithm, self._sha256)) {
            signature = DarabonbaSignatureUtil.Client.HmacSHA256SignByBytes(stringToSign, signingkey)
        }
        else if (TeaUtils.Client.equalString(signatureAlgorithm, self._sm3)) {
            signature = DarabonbaSignatureUtil.Client.HmacSM3SignByBytes(stringToSign, signingkey)
        }
        return DarabonbaEncodeUtil.Client.hexEncode(signature)
    }

    @available(macOS 10.15, iOS 13, tvOS 13, watchOS 6, *)
    public func getSigningkey(_ signatureAlgorithm: String, _ secret: String, _ product: String, _ region: String, _ date: String) async throws -> [UInt8] {
        var sc1: String = (self._signPrefix ?? "") + (secret as! String)
        var sc2: [UInt8] = TeaUtils.Client.toBytes("")
        if (TeaUtils.Client.equalString(signatureAlgorithm, self._sha256)) {
            sc2 = DarabonbaSignatureUtil.Client.HmacSHA256Sign(date, sc1)
        }
        else if (TeaUtils.Client.equalString(signatureAlgorithm, self._sm3)) {
            sc2 = DarabonbaSignatureUtil.Client.HmacSM3Sign(date, sc1)
        }
        var sc3: [UInt8] = TeaUtils.Client.toBytes("")
        if (TeaUtils.Client.equalString(signatureAlgorithm, self._sha256)) {
            sc3 = DarabonbaSignatureUtil.Client.HmacSHA256SignByBytes(region, sc2)
        }
        else if (TeaUtils.Client.equalString(signatureAlgorithm, self._sm3)) {
            sc3 = DarabonbaSignatureUtil.Client.HmacSM3SignByBytes(region, sc2)
        }
        var sc4: [UInt8] = TeaUtils.Client.toBytes("")
        if (TeaUtils.Client.equalString(signatureAlgorithm, self._sha256)) {
            sc4 = DarabonbaSignatureUtil.Client.HmacSHA256SignByBytes(product, sc3)
        }
        else if (TeaUtils.Client.equalString(signatureAlgorithm, self._sm3)) {
            sc4 = DarabonbaSignatureUtil.Client.HmacSM3SignByBytes(product, sc3)
        }
        var hmac: [UInt8] = TeaUtils.Client.toBytes("")
        if (TeaUtils.Client.equalString(signatureAlgorithm, self._sha256)) {
            hmac = DarabonbaSignatureUtil.Client.HmacSHA256SignByBytes((self._signPrefix ?? "") + "_request", sc4)
        }
        else if (TeaUtils.Client.equalString(signatureAlgorithm, self._sm3)) {
            hmac = DarabonbaSignatureUtil.Client.HmacSM3SignByBytes((self._signPrefix ?? "") + "_request", sc4)
        }
        return hmac as! [UInt8]
    }

    public func getRegion(_ product: String, _ endpoint: String, _ regionId: String) -> String {
        if (!TeaUtils.Client.empty(regionId)) {
            return regionId as! String
        }
        var region: String = "center"
        if (TeaUtils.Client.empty(product) || TeaUtils.Client.empty(endpoint)) {
            return region as! String
        }
        var strs: [String] = DarabonbaString.Client.split(endpoint, ":", nil)
        var withoutPort: String = strs[0]
        var preRegion: String = DarabonbaString.Client.replace(withoutPort, "." + (self._endpointSuffix ?? ""), "", nil)
        var nodes: [String] = DarabonbaString.Client.split(preRegion, ".", nil)
        if (TeaUtils.Client.equalNumber(DarabonbaArray.Client.size(nodes), 2)) {
            region = nodes[1]
        }
        return region as! String
    }

    @available(macOS 10.15, iOS 13, tvOS 13, watchOS 6, *)
    public func buildCanonicalizedResource(_ query: [String: String]) async throws -> String {
        var canonicalizedResource: String = ""
        if (!TeaUtils.Client.isUnset(query)) {
            var queryArray: [String] = DarabonbaMap.Client.keySet(query)
            var sortedQueryArray: [String] = DarabonbaArray.Client.ascSort(queryArray)
            var separator: String = ""
            for key in sortedQueryArray {
                canonicalizedResource = (canonicalizedResource as! String) + (separator as! String) + (DarabonbaEncodeUtil.Client.percentEncode(key))
                if (!TeaUtils.Client.empty(query[key as! String])) {
                    canonicalizedResource = (canonicalizedResource as! String) + "=" + (DarabonbaEncodeUtil.Client.percentEncode(query[key as! String]))
                }
                separator = "&"
            }
        }
        return canonicalizedResource as! String
    }

    @available(macOS 10.15, iOS 13, tvOS 13, watchOS 6, *)
    public func buildCanonicalizedHeaders(_ headers: [String: String]) async throws -> String {
        var canonicalizedHeaders: String = ""
        var sortedHeaders: [String] = try await getSignedHeaders(headers as! [String: String])
        for header in sortedHeaders {
            canonicalizedHeaders = (canonicalizedHeaders as! String) + (header as! String) + ":" + (DarabonbaString.Client.trim(headers[header as! String])) + "\n"
        }
        return canonicalizedHeaders as! String
    }

    @available(macOS 10.15, iOS 13, tvOS 13, watchOS 6, *)
    public func getSignedHeaders(_ headers: [String: String]) async throws -> [String] {
        var headersArray: [String] = DarabonbaMap.Client.keySet(headers)
        var sortedHeadersArray: [String] = DarabonbaArray.Client.ascSort(headersArray)
        var tmp: String = ""
        var separator: String = ""
        for key in sortedHeadersArray {
            var lowerKey: String = DarabonbaString.Client.toLower(key)
            if (DarabonbaString.Client.hasPrefix(lowerKey, "x-acs-") || DarabonbaString.Client.equals(lowerKey, "host") || DarabonbaString.Client.equals(lowerKey, "content-type")) {
                if (!DarabonbaString.Client.contains(tmp, lowerKey)) {
                    tmp = (tmp as! String) + (separator as! String) + (lowerKey as! String)
                    separator = ";"
                }
            }
        }
        return DarabonbaString.Client.split(tmp, ";", nil)
    }
}
