// This file is auto-generated, don't edit it
import SPI, * as $SPI from '@alicloud/gateway-spi';
import Credential, * as $Credential from '@alicloud/credentials';
import Util from '@alicloud/tea-util';
import OpenApiUtil from '@alicloud/openapi-util';
import EncodeUtil from '@alicloud/darabonba-encode-util';
import SignatureUtil from '@alicloud/darabonba-signature-util';
import String from '@alicloud/darabonba-string';
import Map from '@alicloud/darabonba-map';
import Array from '@alicloud/darabonba-array';
import XML from '@alicloud/tea-xml';
import * as $tea from '@alicloud/tea-typescript';


export default class Client extends SPI {
  _signPrefix: string;
  _signSuffix: string;
  _authPrefix: string;

  constructor() {
    super();
    this._signPrefix = "aliyun_v4";
    this._signSuffix = "aliyun_v4_request";
    this._authPrefix = "MNS4-HMAC-SHA256";
  }


  async modifyConfiguration(context: $SPI.InterceptorContext, attributeMap: $SPI.AttributeMap): Promise<void> {
    let config = context.configuration;
    config.endpoint = await this.getEndpoint(config.regionId, config.endpoint);
  }

  async modifyRequest(context: $SPI.InterceptorContext, attributeMap: $SPI.AttributeMap): Promise<void> {
    let request = context.request;
    let config = context.configuration;
    let signatureVersion = Util.defaultString(request.signatureVersion, "v2");
    if (!Util.isUnset(request.body)) {
      if (String.equals(request.reqBodyType, "xml")) {
        let reqBodyMap = Util.assertAsMap(request.body);
        let xmlStr = XML.toXML(reqBodyMap);
        request.stream = new $tea.BytesReadable(xmlStr);
        request.headers["content-type"] = "text/xml;charset=UTF-8";
        request.headers["content-md5"] = EncodeUtil.base64EncodeToString(SignatureUtil.MD5Sign(xmlStr));
      } else if (String.equals(request.reqBodyType, "json") || String.equals(request.reqBodyType, "formData")) {
        let bodyStr = Util.toJSONString(request.body);
        request.stream = new $tea.BytesReadable(bodyStr);
        request.headers["content-type"] = "application/json";
        request.headers["content-md5"] = EncodeUtil.base64EncodeToString(SignatureUtil.MD5Sign(bodyStr));
      } else if (String.equals(request.reqBodyType, "byte") || String.equals(request.reqBodyType, "binary")) {
        let bodyBytes = Util.assertAsBytes(request.body);
        request.stream = new $tea.BytesReadable(bodyBytes);
        request.headers["content-md5"] = EncodeUtil.base64EncodeToString(SignatureUtil.MD5SignForBytes(bodyBytes));
      }

    }

    await this.buildRequest(context);
    request.headers = {
      host: config.endpoint,
      'x-mns-version': request.version,
      'user-agent': request.userAgent,
      accept: "application/json",
      ...request.headers,
    };
    if (!Util.equalString(request.authType, "Anonymous")) {
      let credential : Credential = request.credential;
      if (Util.isUnset(credential)) {
        throw $tea.newError({
          code: "ParameterMissing",
          message: "'config.credential' can not be unset",
        });
      }

      let credentialModel = await credential.getCredential();
      let authType = credentialModel.type;
      if (Util.equalString(authType, "bearer")) {
        let bearerToken = credentialModel.bearerToken;
        request.headers["x-acs-bearer-token"] = bearerToken;
        request.headers["x-acs-signature-type"] = "BEARERTOKEN";
        request.headers["Authorization"] = `Bearer ${bearerToken}`;
      } else {
        let accessKeyId = credentialModel.accessKeyId;
        let accessKeySecret = credentialModel.accessKeySecret;
        let securityToken = credentialModel.securityToken;
        if (!Util.empty(securityToken)) {
          request.headers["security-token"] = securityToken;
        }

        request.headers["date"] = Util.getDateUTCString();
        if (String.equals(signatureVersion, "v4")) {
          let date = await this.getDateISO8601();
          request.headers["authorization"] = await this.getAuthorizationV4(context, date, accessKeyId, accessKeySecret);
        } else {
          request.headers["authorization"] = await this.getAuthorizationV2(request.pathname, request.method, request.headers, accessKeyId, accessKeySecret);
        }

      }

    }

  }

  async modifyResponse(context: $SPI.InterceptorContext, attributeMap: $SPI.AttributeMap): Promise<void> {
    let request = context.request;
    let response = context.response;
    if (Util.is4xx(response.statusCode) || Util.is5xx(response.statusCode)) {
      let err : {[key: string ]: any} = { };
      if (!Util.isUnset(response.headers["content-type"]) && String.contains(response.headers["content-type"], "text/xml")) {
        let _str = await Util.readAsString(response.body);
        let respMap = XML.parseXml(_str, null);
        err = Util.assertAsMap(respMap["Error"]);
      } else {
        let _res = await Util.readAsJSON(response.body);
        err = Util.assertAsMap(_res);
      }

      let requestId = this.defaultAny(err["RequestId"], err["requestId"]);
      if (!Util.isUnset(response.headers["x-mns-request-id"])) {
        requestId = response.headers["x-mns-request-id"];
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

  async getEndpoint(regionId: string, endpoint: string): Promise<string> {
    if (!Util.empty(endpoint)) {
      return endpoint;
    }

    if (Util.empty(regionId)) {
      regionId = "cn-hangzhou";
    }

    return `${regionId}.mns.aliyuncs.com`;
  }

  defaultAny(inputValue: any, defaultValue: any): any {
    if (Util.isUnset(inputValue)) {
      return defaultValue;
    }

    return inputValue;
  }

  defaultString(inputValue: string, defaultValue: string): string {
    if (Util.isUnset(inputValue) || Util.empty(inputValue)) {
      return defaultValue;
    }

    return inputValue;
  }

  static base64Encode(input: string): string {
    if (Util.isUnset(input)) {
      return "";
    }

    return EncodeUtil.base64EncodeToString(Util.toBytes(input));
  }

  static base64Decode(input: string): string {
    if (Util.isUnset(input)) {
      return "";
    }

    return Util.toString(EncodeUtil.base64Decode(input));
  }

  async buildRequest(context: $SPI.InterceptorContext): Promise<void> {
    let request = context.request;
    let resource : string = request.pathname;
    if (!Util.empty(resource)) {
      let paths : string[] = String.split(resource, `?`, 2);
      resource = paths[0];
      if (Util.equalNumber(Array.size(paths), 2)) {
        let params : string[] = String.split(paths[1], "&", null);

        for (let sub of params) {
          let item : string[] = String.split(sub, "=", null);
          let key : string = item[0];
          let value : string = null;
          if (Util.equalNumber(Array.size(item), 2)) {
            value = item[1];
          }

          request.query[key] = value;
        }
      }

    }

    request.pathname = resource;
  }

  async getAuthorizationV2(pathname: string, method: string, headers: {[key: string ]: string}, ak: string, secret: string): Promise<string> {
    let sign = await this.getSignatureV2(pathname, method, headers, secret);
    return `MNS ${ak}:${sign}`;
  }

  async getSignatureV2(pathname: string, method: string, headers: {[key: string ]: string}, secret: string): Promise<string> {
    let canonicalizedResource = this.defaultString(pathname, "/");
    let canonicalizedHeaders = await this.buildCanonicalizedHeadersV2(headers);
    let stringToSign = `${method}\n${canonicalizedHeaders}${canonicalizedResource}`;
    return EncodeUtil.base64EncodeToString(SignatureUtil.HmacSHA1Sign(stringToSign, secret));
  }

  async buildCanonicalizedHeadersV2(headers: {[key: string ]: string}): Promise<string> {
    let contentMd5 = this.defaultString(headers["content-md5"], "");
    let contentType = this.defaultString(headers["content-type"], "");
    let date = this.defaultString(headers["date"], "");
    let canonicalizedHeaders : string = `${contentMd5}\n${contentType}\n${date}\n`;
    let mnsHeaders : {[key: string ]: string} = { };

    for (let header of Map.keySet(headers)) {
      let lowerHeader = String.toLower(header);
      if (String.hasPrefix(lowerHeader, "x-mns-")) {
        mnsHeaders[lowerHeader] = headers[header];
      }

    }

    for (let header of Array.ascSort(Map.keySet(mnsHeaders))) {
      canonicalizedHeaders = `${canonicalizedHeaders}${header}:${mnsHeaders[header]}\n`;
    }
    return canonicalizedHeaders;
  }

  async getAuthorizationV4(context: $SPI.InterceptorContext, date: string, accessKeyId: string, accessKeySecret: string): Promise<string> {
    let region = await this.getRegion(context);
    let dateShort = String.subString(date, 0, 8);
    dateShort = String.replace(dateShort, "T", "", null);
    let scope = `${dateShort}/${region}/mns/${this._signSuffix}`;
    let signingkey = await this.getSigningkeyV4(accessKeySecret, region, dateShort);
    let signature = await this.getSignatureV4(context, date, scope, signingkey);
    return `${this._authPrefix} Credential=${accessKeyId}/${scope},Signature=${signature}`;
  }

  async getSigningkeyV4(secret: string, region: string, date: string): Promise<Buffer> {
    let sc1 = `${this._signPrefix}${secret}`;
    let sc2 = SignatureUtil.HmacSHA256Sign(date, sc1);
    let sc3 = SignatureUtil.HmacSHA256SignByBytes(region, sc2);
    let sc4 = SignatureUtil.HmacSHA256SignByBytes("mns", sc3);
    return SignatureUtil.HmacSHA256SignByBytes(this._signSuffix, sc4);
  }

  async getSignatureV4(context: $SPI.InterceptorContext, date: string, scope: string, signingkey: Buffer): Promise<string> {
    let request = context.request;
    let canonicalString = await this.buildCanonicalRequestV4(request.pathname, request.method, request.query, request.headers);
    let stringToSign = `${this._authPrefix}\n${date}\n${scope}\n${canonicalString}`;
    let signature = SignatureUtil.HmacSHA256SignByBytes(stringToSign, signingkey);
    return EncodeUtil.hexEncode(signature);
  }

  async buildCanonicalRequestV4(pathname: string, method: string, query: {[key: string ]: string}, headers: {[key: string ]: string}): Promise<string> {
    let canonicalURI = "/";
    if (!Util.empty(pathname)) {
      canonicalURI = pathname;
      if (!String.hasPrefix(canonicalURI, "/")) {
        canonicalURI = `/${canonicalURI}`;
      }

    }

    let suffix : string = "";
    if (!String.equals(canonicalURI, "/") && String.hasSuffix(canonicalURI, "/")) {
      suffix = "/";
    }

    canonicalURI = `${OpenApiUtil.getEncodePath(canonicalURI)}${suffix}`;
    let resources = await this.buildCanonicalizedResourceV4(query);
    let canonicalHeaders = await this.buildCanonicalizedHeadersV4(headers);
    return `${method}\n${canonicalURI}\n${resources}\n${canonicalHeaders}\n`;
  }

  async buildCanonicalizedResourceV4(query: {[key: string ]: string}): Promise<string> {
    let canonicalizedResource : string = "";
    if (!Util.isUnset(query)) {
      let queryMap : {[key: string ]: string} = { };

      for (let key of Map.keySet(query)) {
        let encodedKey = this.percentEncodeMns(String.toLower(String.trim(key)));
        let encodedValue = "";
        if (!Util.isUnset(query[key]) && !Util.empty(query[key])) {
          encodedValue = this.percentEncodeMns(String.trim(query[key]));
        }

        queryMap[encodedKey] = encodedValue;
      }
      let queryArray : string[] = Map.keySet(queryMap);
      let sortedQueryArray = Array.ascSort(queryArray);
      let separator : string = "";

      for (let encodedKey of sortedQueryArray) {
        canonicalizedResource = `${canonicalizedResource}${separator}${encodedKey}`;
        let encodedValue = queryMap[encodedKey];
        if (!Util.empty(encodedValue)) {
          canonicalizedResource = `${canonicalizedResource}=${encodedValue}`;
        }

        separator = "&";
      }
    }

    return canonicalizedResource;
  }

  async buildCanonicalizedHeadersV4(headers: {[key: string ]: string}): Promise<string> {
    let signedHeaders : {[key: string ]: string} = { };

    for (let key of Map.keySet(headers)) {
      let lowerKey = String.toLower(key);
      if (this.hasSignedHeaderV4(lowerKey)) {
        signedHeaders[lowerKey] = String.trim(headers[key]);
      }

    }
    let canonicalizedHeaders : string = "";

    for (let lowerKey of Array.ascSort(Map.keySet(signedHeaders))) {
      canonicalizedHeaders = `${canonicalizedHeaders}${lowerKey}:${signedHeaders[lowerKey]}\n`;
    }
    return canonicalizedHeaders;
  }

  hasSignedHeaderV4(header: string): boolean {
    if (String.equals(header, "content-type") || String.equals(header, "content-md5")) {
      return true;
    }

    return String.hasPrefix(header, "x-mns-");
  }

  percentEncodeMns(value: string): string {
    let encoded = EncodeUtil.percentEncode(value);
    return String.replace(encoded, "+", "%20", null);
  }

  async getRegion(context: $SPI.InterceptorContext): Promise<string> {
    let config = context.configuration;
    if (!Util.isUnset(config.regionId) && !Util.empty(config.regionId)) {
      return config.regionId;
    }

    let region : string = String.replace(config.endpoint, ".mns.aliyuncs.com", "", null);
    if (String.equals(region, config.endpoint)) {
      throw $tea.newError({
        code: "ClientConfigError",
        message: "The regionId configuration of mns client is missing.",
      });
    }

    return region;
  }

  async getDateISO8601(): Promise<string> {
    let date = OpenApiUtil.getTimestamp();
    date = String.replace(date, "-", "", null);
    return String.replace(date, ":", "", null);
  }

}
