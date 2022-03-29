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

  constructor() {
    super();
  }


  async modifyConfiguration(context: $SPI.InterceptorContext, attributeMap: $SPI.AttributeMap): Promise<void> {
    let request = context.request;
    let config = context.configuration;
    config.endpoint = this.getEndpoint(request.productId, config.regionId, config.endpointRule, config.network, config.suffix, config.endpointMap, config.endpoint);
  }

  async modifyRequest(context: $SPI.InterceptorContext, attributeMap: $SPI.AttributeMap): Promise<void> {
    let request = context.request;
    let config = context.configuration;
    request.headers = {
      host: config.endpoint,
      'x-acs-version': request.version,
      'x-acs-action': request.action,
      'user-agent': request.userAgent,
      'x-acs-date': OpenApiUtil.getTimestamp(),
      'x-acs-signature-nonce': Util.getNonce(),
      accept: "application/json",
      ...request.headers,
    };
    let signatureAlgorithm : string = "ACS3-HMAC-SHA256";
    if (!Util.isUnset(request.signatureAlgorithm)) {
      signatureAlgorithm = request.signatureAlgorithm;
    }

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

    request.headers["x-acs-content-sha256"] = hashedRequestPayload;
    if (!Util.equalString(request.authType, "Anonymous")) {
      let credential : Credential = request.credential;
      let accessKeyId = await credential.getAccessKeyId();
      let accessKeySecret = await credential.getAccessKeySecret();
      let securityToken = await credential.getSecurityToken();
      if (!Util.empty(securityToken)) {
        request.headers["x-acs-accesskey-id"] = accessKeyId;
        request.headers["x-acs-security-token"] = securityToken;
      }

      request.headers["Authorization"] = await this.getAuthorization(request.pathname, request.method, request.query, request.headers, signatureAlgorithm, hashedRequestPayload, accessKeyId, accessKeySecret);
    }

  }

  async modifyResponse(context: $SPI.InterceptorContext, attributeMap: $SPI.AttributeMap): Promise<void> {
    let request = context.request;
    let config = context.configuration;
    let response = context.response;
    if (Util.is4xx(response.statusCode) || Util.is5xx(response.statusCode)) {
      let _res = await Util.readAsJSON(response.body);
      let err = Util.assertAsMap(_res);
      let requestId = this.defaultAny(err["RequestId"], err["requestId"]);
      throw $tea.newError({
        code: `${this.defaultAny(err["Code"], err["code"])}`,
        message: `code: ${response.statusCode}, ${this.defaultAny(err["Message"], err["message"])} request id: ${requestId}`,
        data: err,
      });
    }

    if (Util.equalString(request.bodyType, "binary")) {
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

  async getAuthorization(pathname: string, method: string, query: {[key: string ]: string}, headers: {[key: string ]: string}, signatureAlgorithm: string, payload: string, ak: string, secret: string): Promise<string> {
    let signature = await this.getSignature(pathname, method, query, headers, signatureAlgorithm, payload, secret);
    return `${signatureAlgorithm}  Credential=${ak},SignedHeaders=${Array.join(await this.getSignedHeaders(headers), ";")},Signature=${signature}`;
  }

  async getSignature(pathname: string, method: string, query: {[key: string ]: string}, headers: {[key: string ]: string}, signatureAlgorithm: string, payload: string, secret: string): Promise<string> {
    let canonicalURI : string = "/";
    if (!Util.empty(pathname)) {
      canonicalURI = pathname;
    }

    let stringToSign : string = "";
    let canonicalizedResource = await this.buildCanonicalizedResource(query);
    let canonicalizedHeaders = await this.buildCanonicalizedHeaders(headers);
    let signedHeaders = await this.getSignedHeaders(headers);
    stringToSign = `${method}\n${canonicalURI}\n${canonicalizedResource}\n${canonicalizedHeaders}\n${Array.join(signedHeaders, ";")}\n${payload}`;
    let hex = EncodeUtil.hexEncode(EncodeUtil.hash(Util.toBytes(stringToSign), signatureAlgorithm));
    stringToSign = `${signatureAlgorithm}\n${hex}`;
    let signature = Util.toBytes("");
    if (String.equals(signatureAlgorithm, "ACS3-HMAC-SHA256")) {
      signature = SignatureUtil.HmacSHA256Sign(stringToSign, secret);
    } else if (String.equals(signatureAlgorithm, "ACS3-HMAC-SM3")) {
      signature = SignatureUtil.HmacSM3Sign(stringToSign, secret);
    } else if (String.equals(signatureAlgorithm, "ACS3-RSA-SHA256")) {
      signature = SignatureUtil.SHA256withRSASign(stringToSign, secret);
    }

    return EncodeUtil.hexEncode(signature);
  }

  async buildCanonicalizedResource(query: {[key: string ]: string}): Promise<string> {
    let canonicalizedResource : string = "";
    if (!Util.isUnset(query)) {
      let queryArray : string[] = Map.keySet(query);
      let sortedQueryArray = Array.ascSort(queryArray);

      for (let key of sortedQueryArray) {
        canonicalizedResource = `${canonicalizedResource}&${EncodeUtil.percentEncode(key)}`;
        if (!Util.empty(query[key])) {
          canonicalizedResource = `${canonicalizedResource}=${EncodeUtil.percentEncode(query[key])}`;
        }

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
    return String.split(tmp, ";", 0);
  }

}
