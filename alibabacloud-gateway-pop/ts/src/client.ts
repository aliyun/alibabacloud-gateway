// This file is auto-generated, don't edit it
import SPI, * as $SPI from '@alicloud/gateway-spi';
import Credential, * as $Credential from '@alicloud/credentials';
import Util from '@alicloud/tea-util';
import OpenApiUtil from '@alicloud/openapi-util';
import EndpointUtil from '@alicloud/endpoint-util';
import XML from '@alicloud/tea-xml';
import * as codec from './codec';
import * as $tea from '@alicloud/tea-typescript';


export default class Client extends SPI {
  _endpointSuffix: string;
  _signatureTypePrefix: string;
  _signPrefix: string;
  _sha256: string;
  _sm3: string;

  constructor() {
    super();
    // CLOUD4-
    this._signatureTypePrefix = "ACS4-";
    // cloud_v4
    this._signPrefix = "aliyun_v4";
    this._endpointSuffix = "aliyuncs.com";
    this._sha256 = `${this._signatureTypePrefix}HMAC-SHA256`;
    this._sm3 = `${this._signatureTypePrefix}HMAC-SM3`;
  }


  async modifyConfiguration(context: $SPI.InterceptorContext, attributeMap: $SPI.AttributeMap): Promise<void> {
    let request = context.request;
    let config = context.configuration;
    let attributes = attributeMap.key;
    if (!Util.isUnset(attributes)) {
      this._signatureTypePrefix = attributes["signatureTypePrefix"];
      this._signPrefix = attributes["signPrefix"];
      this._endpointSuffix = attributes["endpointSuffix"];
      this._sha256 = `${this._signatureTypePrefix}HMAC-SHA256`;
      this._sm3 = `${this._signatureTypePrefix}HMAC-SM3`;
    }

    config.endpoint = this.getEndpoint(request.productId, config.regionId, config.endpointRule, config.network, config.suffix, config.endpointMap, config.endpoint);
  }

  async modifyRequest(context: $SPI.InterceptorContext, attributeMap: $SPI.AttributeMap): Promise<void> {
    let request = context.request;
    let config = context.configuration;
    let date = OpenApiUtil.getTimestamp();
    request.headers = {
      host: config.endpoint,
      'x-acs-version': request.version,
      'user-agent': request.userAgent,
      'x-acs-date': date,
      'x-acs-signature-nonce': Util.getNonce(),
      accept: "application/json",
      ...request.headers,
    };
    if (!Util.empty(request.action)) {
      request.headers["x-acs-action"] = request.action;
    }

    let signatureAlgorithm : string = Util.defaultString(request.signatureAlgorithm, this._sha256);
    let hashedRequestPayload = codec.hexEncode(codec.hash(Util.toBytes(""), signatureAlgorithm));
    if (!Util.isUnset(request.stream)) {
      let tmp = await Util.readAsBytes(request.stream);
      hashedRequestPayload = codec.hexEncode(codec.hash(tmp, signatureAlgorithm));
      request.stream = new $tea.BytesReadable(tmp);
      request.headers["content-type"] = "application/octet-stream";
    } else {
      if (!Util.isUnset(request.body)) {
        if (Util.equalString(request.reqBodyType, "byte")) {
          let byteObj = Util.assertAsBytes(request.body);
          hashedRequestPayload = codec.hexEncode(codec.hash(byteObj, signatureAlgorithm));
          request.stream = new $tea.BytesReadable(byteObj);
        } else if (Util.equalString(request.reqBodyType, "json")) {
          let jsonObj = Util.toJSONString(request.body);
          hashedRequestPayload = codec.hexEncode(codec.hash(Util.toBytes(jsonObj), signatureAlgorithm));
          request.stream = new $tea.BytesReadable(jsonObj);
          request.headers["content-type"] = "application/json; charset=utf-8";
        } else {
          let m = Util.assertAsMap(request.body);
          let formObj = OpenApiUtil.toForm(m);
          hashedRequestPayload = codec.hexEncode(codec.hash(Util.toBytes(formObj), signatureAlgorithm));
          request.stream = new $tea.BytesReadable(formObj);
          request.headers["content-type"] = "application/x-www-form-urlencoded";
        }

      }

    }

    if (Util.equalString(signatureAlgorithm, this._sm3)) {
      request.headers["x-acs-content-sm3"] = hashedRequestPayload;
    } else {
      request.headers["x-acs-content-sha256"] = hashedRequestPayload;
    }

    if (!Util.equalString(request.authType, "Anonymous")) {
      let credential : Credential = request.credential;
      if (Util.isUnset(credential)) {
        throw $tea.newError({
          code: "ParameterMissing",
          message: "'config.credential' can not be unset",
        });
      }

      let credentialModel = await credential.getCredential();
      if (!Util.empty(credentialModel.providerName)) {
        request.headers["x-acs-credentials-provider"] = credentialModel.providerName;
      }

      let authType = credentialModel.type;
      if (Util.equalString(authType, "bearer")) {
        let bearerToken = credentialModel.bearerToken;
        request.headers["x-acs-bearer-token"] = bearerToken;
        request.headers["x-acs-signature-type"] = "BEARERTOKEN";
        request.headers["Authorization"] = `Bearer ${bearerToken}`;
      } else if (Util.equalString(authType, "id_token")) {
        let idToken = credentialModel.securityToken;
        request.headers["x-acs-zero-trust-idtoken"] = idToken;
      } else {
        let accessKeyId = credentialModel.accessKeyId;
        let accessKeySecret = credentialModel.accessKeySecret;
        let securityToken = credentialModel.securityToken;
        if (!Util.empty(securityToken)) {
          request.headers["x-acs-accesskey-id"] = accessKeyId;
          request.headers["x-acs-security-token"] = securityToken;
        }

        let dateNew = codec.subString(date, 0, 10);
        dateNew = codec.replace(dateNew, "-", "", null);
        let region = this.getRegion(request.productId, config.endpoint, config.regionId);
        let signingkey = this.getSigningkey(signatureAlgorithm, accessKeySecret, request.productId, region, dateNew);
        request.headers["Authorization"] = this.getAuthorization(request.pathname, request.method, request.query, request.headers, signatureAlgorithm, hashedRequestPayload, accessKeyId, signingkey, request.productId, region, dateNew);
      }

    }

  }

  async modifyResponse(context: $SPI.InterceptorContext, attributeMap: $SPI.AttributeMap): Promise<void> {
    let request = context.request;
    let response = context.response;
    if (Util.is4xx(response.statusCode) || Util.is5xx(response.statusCode)) {
      let err : {[key: string ]: any} = { };
      if (!Util.isUnset(response.headers["content-type"]) && codec.contains(response.headers["content-type"], "text/xml")) {
        let _str = await Util.readAsString(response.body);
        let respMap = XML.parseXml(_str, null);
        err = Util.assertAsMap(respMap["Error"]);
      } else {
        let _res = await Util.readAsJSON(response.body);
        err = Util.assertAsMap(_res);
      }

      let requestId = this.defaultAny(err["RequestId"], err["requestId"]);
      if (!Util.isUnset(response.headers["x-acs-request-id"])) {
        requestId = response.headers["x-acs-request-id"];
      }

      err["statusCode"] = response.statusCode;
      throw $tea.newError({
        code: `${this.defaultAny(err["Code"], err["code"])}`,
        message: `code: ${response.statusCode}, ${this.defaultAny(err["Message"], err["message"])} request id: ${requestId}`,
        data: err,
        description: `${this.defaultAny(err["Description"], err["description"])}`,
        accessDeniedDetail: this.defaultAny(err["AccessDeniedDetail"], err["accessDeniedDetail"]),
      });
    }

    if (Util.equalNumber(response.statusCode, 204)) {
      await Util.readAsString(response.body);
    } else if (Util.equalString(request.bodyType, "binary")) {
      response.deserializedBody = response.body;
    } else if (Util.equalString(request.bodyType, "byte")) {
      let byt = await Util.readAsBytes(response.body);
      response.deserializedBody = byt;
    } else if (Util.equalString(request.bodyType, "string")) {
      let str = await Util.readAsString(response.body);
      response.deserializedBody = str;
    } else if (Util.equalString(request.bodyType, "json")) {
      let obj = await Util.readAsJSON(response.body);
      let res = Util.assertAsMap(obj);
      response.deserializedBody = res;
    } else if (Util.equalString(request.bodyType, "array")) {
      let arr = await Util.readAsJSON(response.body);
      response.deserializedBody = arr;
    } else {
      response.deserializedBody = await Util.readAsString(response.body);
    }

  }

  getEndpoint(productId: string, regionId: string, endpointRule: string, network: string, suffix: string, endpointMap: {[key: string ]: string}, endpoint: string): string {
    if (!Util.empty(endpoint)) {
      return endpoint;
    }

    if (!Util.isUnset(endpointMap) && !Util.empty(endpointMap[regionId])) {
      return endpointMap[regionId];
    }

    return EndpointUtil.getEndpointRules(productId, regionId, endpointRule, network, suffix);
  }

  defaultAny(inputValue: any, defaultValue: any): any {
    if (Util.isUnset(inputValue)) {
      return defaultValue;
    }

    return inputValue;
  }

  getAuthorization(pathname: string, method: string, query: {[key: string ]: string}, headers: {[key: string ]: string}, signatureAlgorithm: string, payload: string, ak: string, signingkey: Buffer, product: string, region: string, date: string): string {
    let signature = this.getSignature(pathname, method, query, headers, signatureAlgorithm, payload, signingkey);
    let signedHeaders = this.getSignedHeaders(headers);
    let signedHeadersStr = codec.join(signedHeaders, ";");
    return `${signatureAlgorithm} Credential=${ak}/${date}/${region}/${product}/${this._signPrefix}_request,SignedHeaders=${signedHeadersStr},Signature=${signature}`;
  }

  getSignature(pathname: string, method: string, query: {[key: string ]: string}, headers: {[key: string ]: string}, signatureAlgorithm: string, payload: string, signingkey: Buffer): string {
    let canonicalURI : string = "/";
    if (!Util.empty(pathname)) {
      canonicalURI = pathname;
    }

    let stringToSign : string = "";
    let canonicalizedResource = this.buildCanonicalizedResource(query);
    let canonicalizedHeaders = this.buildCanonicalizedHeaders(headers);
    let signedHeaders = this.getSignedHeaders(headers);
    this.validateSignedHeaders(signedHeaders);
    let signedHeadersStr = codec.join(signedHeaders, ";");
    stringToSign = `${method}\n${canonicalURI}\n${canonicalizedResource}\n${canonicalizedHeaders}\n${signedHeadersStr}\n${payload}`;
    let hex = codec.hexEncode(codec.hash(Util.toBytes(stringToSign), signatureAlgorithm));
    stringToSign = `${signatureAlgorithm}\n${hex}`;
    let signature = Util.toBytes("");
    if (Util.equalString(signatureAlgorithm, this._sha256)) {
      signature = codec.hmacSha256ByBytes(stringToSign, signingkey);
    } else if (Util.equalString(signatureAlgorithm, this._sm3)) {
      signature = codec.hmacSm3ByBytes(stringToSign, signingkey);
    }

    return codec.hexEncode(signature);
  }

  getSigningkey(signatureAlgorithm: string, secret: string, product: string, region: string, date: string): Buffer {
    let sc1 = `${this._signPrefix}${secret}`;
    let sc2 = Util.toBytes("");
    if (Util.equalString(signatureAlgorithm, this._sha256)) {
      sc2 = codec.hmacSha256(date, sc1);
    } else if (Util.equalString(signatureAlgorithm, this._sm3)) {
      sc2 = codec.hmacSm3(date, sc1);
    }

    let sc3 = Util.toBytes("");
    if (Util.equalString(signatureAlgorithm, this._sha256)) {
      sc3 = codec.hmacSha256ByBytes(region, sc2);
    } else if (Util.equalString(signatureAlgorithm, this._sm3)) {
      sc3 = codec.hmacSm3ByBytes(region, sc2);
    }

    let sc4 = Util.toBytes("");
    if (Util.equalString(signatureAlgorithm, this._sha256)) {
      sc4 = codec.hmacSha256ByBytes(product, sc3);
    } else if (Util.equalString(signatureAlgorithm, this._sm3)) {
      sc4 = codec.hmacSm3ByBytes(product, sc3);
    }

    let hmac = Util.toBytes("");
    if (Util.equalString(signatureAlgorithm, this._sha256)) {
      hmac = codec.hmacSha256ByBytes(`${this._signPrefix}_request`, sc4);
    } else if (Util.equalString(signatureAlgorithm, this._sm3)) {
      hmac = codec.hmacSm3ByBytes(`${this._signPrefix}_request`, sc4);
    }

    return hmac;
  }

  getRegion(product: string, endpoint: string, regionId: string): string {
    if (!Util.empty(regionId)) {
      return regionId;
    }

    let region = "center";
    if (Util.empty(product) || Util.empty(endpoint)) {
      return region;
    }

    let strs : string[] = codec.split(endpoint, ":", null);
    let withoutPort : string = strs[0];
    let preRegion : string = codec.replace(withoutPort, `.${this._endpointSuffix}`, "", null);
    let nodes = codec.split(preRegion, ".", null);
    if (Util.equalNumber(codec.size(nodes), 2)) {
      region = nodes[1];
    }

    return region;
  }

  buildCanonicalizedResource(query: {[key: string ]: string}): string {
    let canonicalizedResource : string = "";
    if (!Util.isUnset(query)) {
      let queryArray : string[] = codec.keySet(query);
      let sortedQueryArray = codec.ascSort(queryArray);
      let separator : string = "";

      for (let key of sortedQueryArray) {
        canonicalizedResource = `${canonicalizedResource}${separator}${codec.percentEncode(key)}=`;
        if (!Util.empty(query[key])) {
          canonicalizedResource = `${canonicalizedResource}${codec.percentEncode(query[key])}`;
        }

        separator = "&";
      }
    }

    return canonicalizedResource;
  }

  buildCanonicalizedHeaders(headers: {[key: string ]: string}): string {
    // lower header key
    let headersArray : string[] = codec.keySet(headers);
    let newHeaders : {[key: string ]: string} = { };

    for (let key of headersArray) {
      let lowerKey = codec.toLower(key);
      let value = headers[key];
      if (!Util.isUnset(value)) {
        if (Util.isUnset(newHeaders[lowerKey]) || codec.equals(lowerKey, "host")) {
          newHeaders[lowerKey] = codec.trim(value);
        } else {
          newHeaders[lowerKey] = `${newHeaders[lowerKey]},${codec.trim(value)}`;
        }

      }

    }
    let canonicalizedHeaders : string = "";
    let sortedHeaders : string[] = this.getSignedHeaders(headers);

    for (let header of sortedHeaders) {
      canonicalizedHeaders = `${canonicalizedHeaders}${header}:${newHeaders[header]}\n`;
    }
    return canonicalizedHeaders;
  }

  validateSignedHeaders(signedHeaders: string[]): void {
    if (!codec.listContains(signedHeaders, "host") || !codec.listContains(signedHeaders, "x-acs-date")) {
      throw $tea.newError({
        code: "InvalidSignedHeaders",
        message: "signed headers must include host and x-acs-date",
      });
    }
  }

  getSignedHeaders(headers: {[key: string ]: string}): string[] {
    let headersArray : string[] = codec.keySet(headers);
    let newHeadersArray : string[] = [ ];

    for (let key of headersArray) {
      let lowerKey = codec.toLower(key);
      let value = headers[key];
      if (!Util.isUnset(value)) {
        codec.append(newHeadersArray, lowerKey);
      }

    }
    let sortedHeadersArray = codec.ascSort(newHeadersArray);
    let result : string[] = [ ];

    for (let key of sortedHeadersArray) {
      if (codec.hasPrefix(key, "x-acs-") || codec.equals(key, "host") || codec.equals(key, "content-type")) {
        if (!codec.listContains(result, key)) {
          codec.append(result, key);
        }

      }

    }
    return result;
  }

}
