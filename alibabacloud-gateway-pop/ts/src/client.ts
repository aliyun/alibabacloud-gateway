// This file is auto-generated, don't edit it
import SPI, * as $SPI from '@alicloud/gateway-spi';
import Credential from '@alicloud/credentials';
import Util from '@alicloud/tea-util';
import OpenApiUtil from '@alicloud/openapi-util';
import EndpointUtil from '@alicloud/endpoint-util';
import EncodeUtil from '@alicloud/darabonba-encode-util';
import SignatureUtil from '@alicloud/darabonba-signature-util';
import String from '@alicloud/darabonba-string';
import Map from '@alicloud/darabonba-map';
import Array from '@alicloud/darabonba-array';
import * as $tea from '@alicloud/tea-typescript';


export default class Client extends SPI {
  _sha256: string;
  _sm3: string;

  constructor() {
    super();
    this._sha256 = "ACS4-HMAC-SHA256";
    this._sm3 = "ACS4-HMAC-SM3";
  }


  async modifyConfiguration(context: $SPI.InterceptorContext, attributeMap: $SPI.AttributeMap): Promise<void> {
    let request = context.request;
    let config = context.configuration;
    config.endpoint = this.getEndpoint(request.productId, config.regionId, config.endpointRule, config.network, config.suffix, config.endpointMap, config.endpoint);
  }

  async modifyRequest(context: $SPI.InterceptorContext, attributeMap: $SPI.AttributeMap): Promise<void> {
    let request = context.request;
    let config = context.configuration;
    let date = OpenApiUtil.getTimestamp();
    request.headers = {
      host: config.endpoint,
      'x-acs-version': request.version,
      'x-acs-action': request.action,
      'user-agent': request.userAgent,
      'x-acs-date': date,
      'x-acs-signature-nonce': Util.getNonce(),
      accept: "application/json",
      ...request.headers,
    };
    let signatureAlgorithm : string = Util.defaultString(request.signatureAlgorithm, this._sha256);
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

      let authType = credential.getType();
      if (Util.equalString(authType, "bearer")) {
        let bearerToken = credential.getBearerToken();
        request.headers["x-acs-bearer-token"] = bearerToken;
        request.headers["Authorization"] = `Bearer ${bearerToken}`;
      } else {
        let accessKeyId = await credential.getAccessKeyId();
        let accessKeySecret = await credential.getAccessKeySecret();
        let securityToken = await credential.getSecurityToken();
        if (!Util.empty(securityToken)) {
          request.headers["x-acs-accesskey-id"] = accessKeyId;
          request.headers["x-acs-security-token"] = securityToken;
        }

        let dateNew = String.subString(date, 0, 10);
        dateNew = String.replace(dateNew, "-", "", null);
        let region = this.getRegion(request.productId, config.endpoint);
        let signingkey = await this.getSigningkey(signatureAlgorithm, accessKeySecret, request.productId, region, dateNew);
        request.headers["Authorization"] = await this.getAuthorization(request.pathname, request.method, request.query, request.headers, signatureAlgorithm, hashedRequestPayload, accessKeyId, signingkey, request.productId, region, dateNew);
      }

    }

  }

  async modifyResponse(context: $SPI.InterceptorContext, attributeMap: $SPI.AttributeMap): Promise<void> {
    let request = context.request;
    let response = context.response;
    if (Util.is4xx(response.statusCode) || Util.is5xx(response.statusCode)) {
      let _res = await Util.readAsJSON(response.body);
      let err = Util.assertAsMap(_res);
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

  async getAuthorization(pathname: string, method: string, query: {[key: string ]: string}, headers: {[key: string ]: string}, signatureAlgorithm: string, payload: string, ak: string, signingkey: Buffer, product: string, region: string, date: string): Promise<string> {
    let signature = await this.getSignature(pathname, method, query, headers, signatureAlgorithm, payload, signingkey);
    let signedHeaders = await this.getSignedHeaders(headers);
    let signedHeadersStr = Array.join(signedHeaders, ";");
    return `${signatureAlgorithm} Credential=${ak}/${date}/${region}/${product}/aliyun_v4_request,SignedHeaders=${signedHeadersStr},Signature=${signature}`;
  }

  async getSignature(pathname: string, method: string, query: {[key: string ]: string}, headers: {[key: string ]: string}, signatureAlgorithm: string, payload: string, signingkey: Buffer): Promise<string> {
    let canonicalURI : string = "/";
    if (!Util.empty(pathname)) {
      canonicalURI = pathname;
    }

    let stringToSign : string = "";
    let canonicalizedResource = await this.buildCanonicalizedResource(query);
    let canonicalizedHeaders = await this.buildCanonicalizedHeaders(headers);
    let signedHeaders = await this.getSignedHeaders(headers);
    let signedHeadersStr = Array.join(signedHeaders, ";");
    stringToSign = `${method}\n${canonicalURI}\n${canonicalizedResource}\n${canonicalizedHeaders}\n${signedHeadersStr}\n${payload}`;
    let hex = EncodeUtil.hexEncode(EncodeUtil.hash(Util.toBytes(stringToSign), signatureAlgorithm));
    stringToSign = `${signatureAlgorithm}\n${hex}`;
    let signature = Util.toBytes("");
    if (Util.equalString(signatureAlgorithm, this._sha256)) {
      signature = SignatureUtil.HmacSHA256SignByBytes(stringToSign, signingkey);
    } else if (Util.equalString(signatureAlgorithm, this._sm3)) {
      signature = SignatureUtil.HmacSM3SignByBytes(stringToSign, signingkey);
    }

    return EncodeUtil.hexEncode(signature);
  }

  async getSigningkey(signatureAlgorithm: string, secret: string, product: string, region: string, date: string): Promise<Buffer> {
    let sc1 = `aliyun_v4${secret}`;
    let sc2 = Util.toBytes("");
    if (Util.equalString(signatureAlgorithm, this._sha256)) {
      sc2 = SignatureUtil.HmacSHA256Sign(date, sc1);
    } else if (Util.equalString(signatureAlgorithm, this._sm3)) {
      sc2 = SignatureUtil.HmacSM3Sign(date, sc1);
    }

    let sc3 = Util.toBytes("");
    if (Util.equalString(signatureAlgorithm, this._sha256)) {
      sc3 = SignatureUtil.HmacSHA256SignByBytes(region, sc2);
    } else if (Util.equalString(signatureAlgorithm, this._sm3)) {
      sc3 = SignatureUtil.HmacSM3SignByBytes(region, sc2);
    }

    let sc4 = Util.toBytes("");
    if (Util.equalString(signatureAlgorithm, this._sha256)) {
      sc4 = SignatureUtil.HmacSHA256SignByBytes(product, sc3);
    } else if (Util.equalString(signatureAlgorithm, this._sm3)) {
      sc4 = SignatureUtil.HmacSM3SignByBytes(product, sc3);
    }

    let hmac = Util.toBytes("");
    if (Util.equalString(signatureAlgorithm, this._sha256)) {
      hmac = SignatureUtil.HmacSHA256SignByBytes("aliyun_v4_request", sc4);
    } else if (Util.equalString(signatureAlgorithm, this._sm3)) {
      hmac = SignatureUtil.HmacSM3SignByBytes("aliyun_v4_request", sc4);
    }

    return hmac;
  }

  getRegion(product: string, endpoint: string): string {
    let region = "center";
    if (Util.empty(product) || Util.empty(endpoint)) {
      return region;
    }

    let preRegion : string = String.replace(endpoint, ".aliyuncs.com", "", null);
    let nodes = String.split(preRegion, ".", null);
    if (Util.equalNumber(Array.size(nodes), 2)) {
      region = nodes[1];
    }

    return region;
  }

  async buildCanonicalizedResource(query: {[key: string ]: string}): Promise<string> {
    let canonicalizedResource : string = "";
    if (!Util.isUnset(query)) {
      let queryArray : string[] = Map.keySet(query);
      let sortedQueryArray = Array.ascSort(queryArray);
      let separator : string = "";

      for (let key of sortedQueryArray) {
        canonicalizedResource = `${canonicalizedResource}${separator}${EncodeUtil.percentEncode(key)}`;
        if (!Util.empty(query[key])) {
          canonicalizedResource = `${canonicalizedResource}=${EncodeUtil.percentEncode(query[key])}`;
        }

        separator = "&";
      }
    }

    return canonicalizedResource;
  }

  async buildCanonicalizedHeaders(headers: {[key: string ]: string}): Promise<string> {
    let canonicalizedHeaders : string = "";
    let sortedHeaders : string[] = await this.getSignedHeaders(headers);

    for (let header of sortedHeaders) {
      canonicalizedHeaders = `${canonicalizedHeaders}${header}:${String.trim(headers[header])}\n`;
    }
    return canonicalizedHeaders;
  }

  async getSignedHeaders(headers: {[key: string ]: string}): Promise<string[]> {
    let headersArray : string[] = Map.keySet(headers);
    let sortedHeadersArray = Array.ascSort(headersArray);
    let tmp : string = "";
    let separator : string = "";

    for (let key of sortedHeadersArray) {
      let lowerKey = String.toLower(key);
      if (String.hasPrefix(lowerKey, "x-acs-") || String.equals(lowerKey, "host") || String.equals(lowerKey, "content-type")) {
        if (!String.contains(tmp, lowerKey)) {
          tmp = `${tmp}${separator}${lowerKey}`;
          separator = ";";
        }

      }

    }
    return String.split(tmp, ";", null);
  }

}
