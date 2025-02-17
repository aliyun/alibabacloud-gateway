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
import * as $tea from '@alicloud/tea-typescript';


export default class Client extends SPI {

  constructor() {
    super();
  }


  async modifyConfiguration(context: $SPI.InterceptorContext, attributeMap: $SPI.AttributeMap): Promise<void> {
    let config = context.configuration;
    config.endpoint = await this.getEndpoint(config.network, config.endpoint);
  }

  async modifyRequest(context: $SPI.InterceptorContext, attributeMap: $SPI.AttributeMap): Promise<void> {
    let request = context.request;
    let config = context.configuration;
    let date = Util.getDateUTCString();
    request.headers = {
      date: date,
      host: config.endpoint,
      'x-acs-version': request.version,
      'x-acs-action': request.action,
      'user-agent': request.userAgent,
      'x-acs-signature-nonce': Util.getNonce(),
      accept: "application/json",
      ...request.headers,
    };
    let signatureAlgorithm: string = Util.defaultString(request.signatureAlgorithm, "ACS4-HMAC-SHA256");
    let signatureVersion = Util.defaultString(request.signatureVersion, "v1");
    let hashedRequestPayload = EncodeUtil.hexEncode(EncodeUtil.hash(Util.toBytes(""), signatureAlgorithm));
    if (!Util.isUnset(request.stream)) {
      let tmp = await Util.readAsBytes(request.stream);
      hashedRequestPayload = EncodeUtil.hexEncode(EncodeUtil.hash(tmp, signatureAlgorithm));
      request.stream = new $tea.BytesReadable(tmp);
      request.headers["content-type"] = "application/octet-stream";
    } else {
      if (!Util.isUnset(request.body)) {
        if (Util.equalString(request.reqBodyType, "json")) {
          let jsonObj = Util.toJSONString(request.body);
          hashedRequestPayload = EncodeUtil.hexEncode(EncodeUtil.hash(Util.toBytes(jsonObj), signatureAlgorithm));
          request.stream = new $tea.BytesReadable(jsonObj);
          request.headers["content-type"] = "application/json; charset=utf-8";
        } else {
          let m = Util.assertAsMap(request.body);
          let formObj = OpenApiUtil.toForm(m);
          hashedRequestPayload = EncodeUtil.hexEncode(EncodeUtil.hash(Util.toBytes(formObj), signatureAlgorithm));
          request.stream = new $tea.BytesReadable(formObj);
          request.headers["content-type"] = "application/x-www-form-urlencoded";
        }

      }

    }

    if (String.equals(signatureVersion, "v4")) {
      if (Util.equalString(signatureAlgorithm, "ACS4-HMAC-SM3")) {
        request.headers["x-acs-content-sm3"] = hashedRequestPayload;
      } else {
        request.headers["x-acs-content-sha256"] = hashedRequestPayload;
      }

    } else {
      request.headers["x-acs-signature-method"] = "HMAC-SHA1";
      request.headers["x-acs-signature-version"] = "1.0";
    }

    if (!Util.equalString(request.authType, "Anonymous") && !Util.isUnset(request.credential)) {
      let credential: Credential = request.credential;
      let credentialModel = await credential.getCredential();
      let authType = credentialModel.type;
      if (Util.equalString(authType, "bearer")) {
        let bearerToken = credentialModel.bearerToken;
        request.headers["x-acs-bearer-token"] = bearerToken;
        request.headers["Authorization"] = `Bearer ${bearerToken}`;
      } else {
        let accessKeyId = credentialModel.accessKeyId;
        let accessKeySecret = credentialModel.accessKeySecret;
        let securityToken = credentialModel.securityToken;
        if (!Util.empty(securityToken)) {
          request.headers["x-acs-security-token"] = securityToken;
        }

        if (String.equals(signatureVersion, "v4")) {
          let dateNew = String.subString(date, 0, 10);
          let region = this.getRegion(config.endpoint);
          let signingkey = await this.getSigningkey(signatureAlgorithm, accessKeySecret, region, dateNew);
          request.headers["Authorization"] = await this.getAuthorizationV4(request.pathname, request.method, request.query, request.headers, signatureAlgorithm, hashedRequestPayload, accessKeyId, signingkey, request.productId, region, dateNew);
        } else {
          request.headers["Authorization"] = await this.getAuthorization(request.pathname, request.method, request.query, request.headers, accessKeyId, accessKeySecret);
        }

      }

    }

  }

  async modifyResponse(context: $SPI.InterceptorContext, attributeMap: $SPI.AttributeMap): Promise<void> {
    let request = context.request;
    let response = context.response;
    if (Util.is4xx(response.statusCode) || Util.is5xx(response.statusCode)) {
      let _res = await Util.readAsJSON(response.body);
      let err = Util.assertAsMap(_res);
      let headers: { [key: string]: string } = response.headers;
      let requestId = headers["x-ca-request-id"];
      err["statusCode"] = response.statusCode;
      err["requestId"] = requestId;
      throw $tea.newError({
        code: `${this.defaultAny(err["Code"], err["code"])}`,
        message: `${this.defaultAny(err["Message"], err["message"])}`,
        data: err,
      });
    }

    if (!Util.isUnset(response.body)) {
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
        response.deserializedBody = await Util.readAsJSON(response.body);
      } else if (Util.equalString(request.bodyType, "array")) {
        response.deserializedBody = await Util.readAsJSON(response.body);
      } else {
        response.deserializedBody = await Util.readAsString(response.body);
      }

    }

  }

  async getEndpoint(network: string, endpoint: string): Promise<string> {
    let realEndpoint = "api.aliyunpds.com";
    if (!Util.empty(endpoint)) {
      realEndpoint = endpoint;
    }

    if (!Util.empty(network) && String.equals(network, "vpc")) {
      realEndpoint = String.replace(realEndpoint, "api.aliyunpds.com", "api-vpc.aliyunpds.com", null);
      realEndpoint = String.replace(realEndpoint, "admin.aliyunpds.com", "admin-vpc.aliyunpds.com", null);
    }

    return realEndpoint;
  }

  defaultAny(inputValue: any, defaultValue: any): any {
    if (Util.isUnset(inputValue)) {
      return defaultValue;
    }

    return inputValue;
  }

  async getAuthorization(pathname: string, method: string, query: { [key: string]: string }, headers: { [key: string]: string }, ak: string, secret: string): Promise<string> {
    let signature = await this.getSignature(pathname, method, query, headers, secret);
    return `acs ${ak}:${signature}`;
  }

  async getSignature(pathname: string, method: string, query: { [key: string]: string }, headers: { [key: string]: string }, secret: string): Promise<string> {
    let stringToSign: string = "";
    let canonicalizedResource = await this.buildCanonicalizedResource(pathname, query);
    let canonicalizedHeaders = await this.buildCanonicalizedHeaders(headers);
    stringToSign = `${method}\n${canonicalizedHeaders}${canonicalizedResource}`;
    let signature = SignatureUtil.HmacSHA1Sign(stringToSign, secret);
    return EncodeUtil.base64EncodeToString(signature);
  }

  async buildCanonicalizedResource(pathname: string, query: { [key: string]: string }): Promise<string> {
    let canonicalizedResource: string = pathname;
    if (!Util.isUnset(query)) {
      let queryArray: string[] = Map.keySet(query);
      let sortedQueryArray = Array.ascSort(queryArray);
      let separator: string = "?";

      for (let key of sortedQueryArray) {
        canonicalizedResource = `${canonicalizedResource}${separator}${key}`;
        if (!Util.empty(query[key])) {
          canonicalizedResource = `${canonicalizedResource}=${query[key]}`;
        }

        separator = "&";
      }
    }

    return canonicalizedResource;
  }

  async buildCanonicalizedHeaders(headers: { [key: string]: string }): Promise<string> {
    let accept = headers["accept"];
    if (Util.isUnset(accept)) {
      accept = "";
    }

    let contentMd5 = headers["content-md5"];
    if (Util.isUnset(contentMd5)) {
      contentMd5 = "";
    }

    let contentType = headers["content-type"];
    if (Util.isUnset(contentType)) {
      contentType = "";
    }

    let date = headers["date"];
    if (Util.isUnset(date)) {
      date = "";
    }

    let canonicalizedHeaders: string = `${accept}\n${contentMd5}\n${contentType}\n${date}\n`;
    let sortedHeaders: string[] = await this.getSignedHeaders(headers);

    for (let header of sortedHeaders) {
      let value = headers[header];
      let valueTrim = String.trim(value);
      canonicalizedHeaders = `${canonicalizedHeaders}${header}:${valueTrim}\n`;
    }
    return canonicalizedHeaders;
  }

  async getSignedHeaders(headers: { [key: string]: string }): Promise<string[]> {
    let headersArray: string[] = Map.keySet(headers);
    let sortedHeadersArray = Array.ascSort(headersArray);
    let tmp: string = "";
    let separator: string = "";

    for (let key of sortedHeadersArray) {
      let lowerKey = String.toLower(key);
      if (String.hasPrefix(lowerKey, "x-acs-")) {
        if (!String.contains(tmp, lowerKey)) {
          tmp = `${tmp}${separator}${lowerKey}`;
          separator = ";";
        }

      }

    }
    return String.split(tmp, ";", null);
  }

  getRegion(endpoint: string): string {
    let region = "center";
    if (Util.empty(endpoint)) {
      return region;
    }

    if (String.contains(endpoint, ".admin.aliyunpds.com")) {
      region = String.replace(endpoint, ".admin.aliyunpds.com", "", null);
    }

    return region;
  }

  async getSigningkey(signatureAlgorithm: string, secret: string, region: string, date: string): Promise<Buffer> {
    let sc1 = `aliyun_v4${secret}`;
    let sc2 = Util.toBytes("");
    if (Util.equalString(signatureAlgorithm, "ACS4-HMAC-SHA256")) {
      sc2 = SignatureUtil.HmacSHA256Sign(date, sc1);
    } else if (Util.equalString(signatureAlgorithm, "ACS4-HMAC-SM3")) {
      sc2 = SignatureUtil.HmacSM3Sign(date, sc1);
    }

    let sc3 = Util.toBytes("");
    if (Util.equalString(signatureAlgorithm, "ACS4-HMAC-SHA256")) {
      sc3 = SignatureUtil.HmacSHA256SignByBytes(region, sc2);
    } else if (Util.equalString(signatureAlgorithm, "ACS4-HMAC-SM3")) {
      sc3 = SignatureUtil.HmacSM3SignByBytes(region, sc2);
    }

    let sc4 = Util.toBytes("");
    if (Util.equalString(signatureAlgorithm, "ACS4-HMAC-SHA256")) {
      sc4 = SignatureUtil.HmacSHA256SignByBytes("pds", sc3);
    } else if (Util.equalString(signatureAlgorithm, "ACS4-HMAC-SM3")) {
      sc4 = SignatureUtil.HmacSM3SignByBytes("pds", sc3);
    }

    let hmac = Util.toBytes("");
    if (Util.equalString(signatureAlgorithm, "ACS4-HMAC-SHA256")) {
      hmac = SignatureUtil.HmacSHA256SignByBytes("aliyun_v4_request", sc4);
    } else if (Util.equalString(signatureAlgorithm, "ACS4-HMAC-SM3")) {
      hmac = SignatureUtil.HmacSM3SignByBytes("aliyun_v4_request", sc4);
    }

    return hmac;
  }

  async getAuthorizationV4(pathname: string, method: string, query: { [key: string]: string }, headers: { [key: string]: string }, signatureAlgorithm: string, payload: string, ak: string, signingkey: Buffer, product: string, region: string, date: string): Promise<string> {
    let signature = await this.getSignatureV4(pathname, method, query, headers, signatureAlgorithm, payload, signingkey);
    let signedHeaders = await this.getSignedHeaders(headers);
    let signedHeadersStr = Array.join(signedHeaders, ";");
    return `${signatureAlgorithm} Credential=${ak}/${date}/${region}/${product}/aliyun_v4_request,SignedHeaders=${signedHeadersStr},Signature=${signature}`;
  }

  async getSignatureV4(pathname: string, method: string, query: { [key: string]: string }, headers: { [key: string]: string }, signatureAlgorithm: string, payload: string, signingkey: Buffer): Promise<string> {
    let stringToSign: string = "";
    let canonicalizedResource = await this.buildCanonicalizedResource(pathname, query);
    let canonicalizedHeaders = await this.buildCanonicalizedHeaders(headers);
    let signedHeaders = await this.getSignedHeaders(headers);
    let signedHeadersStr = Array.join(signedHeaders, ";");
    stringToSign = `${method}\n${canonicalizedResource}\n${canonicalizedHeaders}\n${signedHeadersStr}\n${payload}`;
    let hex = EncodeUtil.hexEncode(EncodeUtil.hash(Util.toBytes(stringToSign), signatureAlgorithm));
    stringToSign = `${signatureAlgorithm}\n${hex}`;
    let signature = Util.toBytes("");
    if (Util.equalString(signatureAlgorithm, "ACS4-HMAC-SHA256")) {
      signature = SignatureUtil.HmacSHA256SignByBytes(stringToSign, signingkey);
    } else if (Util.equalString(signatureAlgorithm, "ACS4-HMAC-SM3")) {
      signature = SignatureUtil.HmacSM3SignByBytes(stringToSign, signingkey);
    }

    return EncodeUtil.hexEncode(signature);
  }

}
