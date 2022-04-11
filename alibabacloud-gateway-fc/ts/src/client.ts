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

export class HttpRequest extends $tea.Model {
  method: string;
  path: string;
  headers?: { [key: string]: any };
  body?: Buffer;
  reqBodyType?: string;
  static names(): { [key: string]: string } {
    return {
      method: 'method',
      path: 'path',
      headers: 'headers',
      body: 'body',
      reqBodyType: 'reqBodyType',
    };
  }

  static types(): { [key: string]: any } {
    return {
      method: 'string',
      path: 'string',
      headers: { 'type': 'map', 'keyType': 'string', 'valueType': 'any' },
      body: 'Buffer',
      reqBodyType: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}


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
    let config = context.configuration;
    if (!String.hasSuffix(config.endpoint, "aliyuncs.com")) {
      await this.signRequestForFc(context);
    } else {
      await this.signRequestForPop(context);
    }

  }

  async modifyResponse(context: $SPI.InterceptorContext, attributeMap: $SPI.AttributeMap): Promise<void> {
    let request = context.request;
    let config = context.configuration;
    let response = context.response;
    if (Util.is4xx(response.statusCode) || Util.is5xx(response.statusCode)) {
      if (String.hasPrefix(config.endpoint, "fc.") && String.hasSuffix(config.endpoint, ".aliyuncs.com")) {
        let popRes = await Util.readAsJSON(response.body);
        let popErr = Util.assertAsMap(popRes);
        throw $tea.newError({
          code: `${this.defaultAny(popErr["Code"], popErr["code"])}`,
          message: `code: ${response.statusCode}, ${this.defaultAny(popErr["Message"], popErr["message"])} request id: ${this.defaultAny(popErr["RequestID"], popErr["RequestId"])}`,
          data: popErr,
        });
      } else {
        let _headers = Util.assertAsMap(response.headers);
        let fcRes = await Util.readAsJSON(response.body);
        let fcErr = Util.assertAsMap(fcRes);
        throw $tea.newError({
          code: fcErr["ErrorCode"],
          message: `code: ${response.statusCode}, ${fcErr["ErrorMessage"]} request id: ${_headers["x-fc-request-id"]}`,
          data: fcErr,
        });
      }

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

  async signRequestForFc(context: $SPI.InterceptorContext): Promise<void> {
    let request = context.request;
    let config = context.configuration;
    request.headers = {
      host: config.endpoint,
      date: Util.getDateUTCString(),
      accept: "application/json",
      'user-agent': request.userAgent,
      ...request.headers,
    };
    request.headers["content-type"] = "application/json";
    if (!Util.isUnset(request.stream)) {
      let tmp = await Util.readAsBytes(request.stream);
      request.stream = new $tea.BytesReadable(tmp);
      request.headers["content-type"] = "application/octet-stream";
      request.headers["content-md5"] = EncodeUtil.base64EncodeToString(SignatureUtil.MD5SignForBytes(tmp));
    } else {
      if (!Util.isUnset(request.body)) {
        if (Util.equalString(request.reqBodyType, "json")) {
          let jsonObj = Util.toJSONString(request.body);
          request.stream = new $tea.BytesReadable(jsonObj);
          request.headers["content-type"] = "application/json";
          request.headers["content-md5"] = EncodeUtil.base64EncodeToString(SignatureUtil.MD5Sign(jsonObj));
        } else {
          let m = Util.assertAsMap(request.body);
          let formObj = OpenApiUtil.toForm(m);
          request.stream = new $tea.BytesReadable(formObj);
          request.headers["content-type"] = "application/x-www-form-urlencoded";
          request.headers["content-md5"] = EncodeUtil.base64EncodeToString(SignatureUtil.MD5Sign(formObj));
        }

      }

    }

    let credential : Credential = request.credential;
    let accessKeyId = await credential.getAccessKeyId();
    let accessKeySecret = await credential.getAccessKeySecret();
    let securityToken = await credential.getSecurityToken();
    if (!Util.empty(securityToken)) {
      request.headers["x-fc-security-token"] = securityToken;
    }

    request.headers["Authorization"] = await this.getAuthorizationForFc(request.pathname, request.method, request.query, request.headers, accessKeyId, accessKeySecret);
  }

  async signRequestForPop(context: $SPI.InterceptorContext): Promise<void> {
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

      request.headers["Authorization"] = await this.getAuthorizationForPop(request.pathname, request.method, request.query, request.headers, signatureAlgorithm, hashedRequestPayload, accessKeyId, accessKeySecret);
    }

  }

  async getAuthorizationForFc(pathname: string, method: string, query: {[key: string ]: string}, headers: {[key: string ]: string}, ak: string, secret: string): Promise<string> {
    let sign = await this.getSignatureForFc(pathname, method, query, headers, secret);
    return `FC ${ak}:${sign}`;
  }

  async getSignatureForFc(pathname: string, method: string, query: {[key: string ]: string}, headers: {[key: string ]: string}, secret: string): Promise<string> {
    let resource : string = pathname;
    let contentMd5 = headers["content-md5"];
    if (Util.isUnset(contentMd5)) {
      contentMd5 = "";
    }

    let contentType = headers["content-type"];
    if (Util.isUnset(contentType)) {
      contentType = "";
    }

    let stringToSign : string = "";
    let canonicalizedResource = await this.buildCanonicalizedResourceForFc(resource, query);
    let canonicalizedHeaders = await this.buildCanonicalizedHeadersForFc(headers);
    stringToSign = `${method}\n${contentMd5}\n${contentType}\n${headers["date"]}\n${canonicalizedHeaders}${canonicalizedResource}`;
    return EncodeUtil.base64EncodeToString(SignatureUtil.HmacSHA256Sign(stringToSign, secret));
  }

  async buildCanonicalizedResourceForFc(pathname: string, query: {[key: string ]: string}): Promise<string> {
    let paths : string[] = String.split(pathname, `?`, 2);
    let canonicalizedResource : string = paths[0];
    let resources : string[] = [ ];
    if (Util.equalNumber(Array.size(paths), 2)) {
      resources = String.split(paths[1], "&", 0);
    }

    let subResources : string[] = [ ];
    let tmp : string = "";
    let separator : string = "";
    if (!Util.isUnset(query)) {
      let queryList : string[] = Map.keySet(query);

      for (let paramName of queryList) {
        tmp = `${tmp}${separator}${paramName}`;
        if (!Util.isUnset(query[paramName])) {
          tmp = `${tmp}=${query[paramName]}`;
        }

        separator = ";";
      }
      subResources = String.split(tmp, ";", 0);
    }

    let result : string[] = Array.concat(subResources, resources);
    let sortedParams = Array.ascSort(result);
    if (Util.equalNumber(Array.size(sortedParams), 0)) {
      return `${canonicalizedResource}\n`;
    }

    let subRes = Array.join(sortedParams, "\n");
    return `${canonicalizedResource}\n${subRes}`;
  }

  async buildCanonicalizedHeadersForFc(headers: {[key: string ]: string}): Promise<string> {
    let canonicalizedHeaders : string = "";
    let keys = Map.keySet(headers);
    let sortedHeaders = Array.ascSort(keys);

    for (let header of sortedHeaders) {
      if (String.contains(String.toLower(header), "x-fc-")) {
        canonicalizedHeaders = `${canonicalizedHeaders}${String.toLower(header)}:${headers[header]}\n`;
      }

    }
    return canonicalizedHeaders;
  }

  async getAuthorizationForPop(pathname: string, method: string, query: {[key: string ]: string}, headers: {[key: string ]: string}, signatureAlgorithm: string, payload: string, ak: string, secret: string): Promise<string> {
    let signature = await this.getSignatureForPop(pathname, method, query, headers, signatureAlgorithm, payload, secret);
    return `${signatureAlgorithm}  Credential=${ak},SignedHeaders=${Array.join(await this.getSignedHeaders(headers), ";")},Signature=${signature}`;
  }

  async getSignatureForPop(pathname: string, method: string, query: {[key: string ]: string}, headers: {[key: string ]: string}, signatureAlgorithm: string, payload: string, secret: string): Promise<string> {
    let canonicalURI : string = "/";
    if (!Util.empty(pathname)) {
      canonicalURI = pathname;
    }

    let stringToSign : string = "";
    let canonicalizedResource = await this.buildCanonicalizedResourceForPop(query);
    let canonicalizedHeaders = await this.buildCanonicalizedHeadersForPop(headers);
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

  async buildCanonicalizedResourceForPop(query: {[key: string ]: string}): Promise<string> {
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

  async buildCanonicalizedHeadersForPop(headers: {[key: string ]: string}): Promise<string> {
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

  async signRequest(request: HttpRequest, credential: Credential): Promise<{[key: string ]: any}> {
    let httpRequest : HttpRequest = new HttpRequest({
      method: request.method,
      path: request.path,
      headers: request.headers,
      body: request.body,
      reqBodyType: request.reqBodyType,
    });
    httpRequest.headers["date"] = Util.getDateUTCString();
    httpRequest.headers["accept"] = "application/json";
    httpRequest.headers["content-type"] = "application/json";
    if (!Util.isUnset(request.body)) {
      if (Util.equalString(request.reqBodyType, "json")) {
        httpRequest.headers["content-type"] = "application/json";
      } else if (Util.equalString(request.reqBodyType, "form")) {
        httpRequest.headers["content-type"] = "application/x-www-form-urlencoded";
      } else if (Util.equalString(request.reqBodyType, "binary")) {
        httpRequest.headers["content-type"] = "application/octet-stream";
      }

    }

    let accessKeyId = await credential.getAccessKeyId();
    let accessKeySecret = await credential.getAccessKeySecret();
    let securityToken = await credential.getSecurityToken();
    if (!Util.empty(securityToken)) {
      httpRequest.headers["x-fc-security-token"] = securityToken;
    }

    let resource : string = request.path;
    let contentMd5 = httpRequest.headers["content-md5"];
    if (Util.isUnset(contentMd5)) {
      contentMd5 = "";
    }

    let contentType = httpRequest.headers["content-type"];
    if (Util.isUnset(contentType)) {
      contentType = "";
    }

    let stringToSign : string = "";
    let canonicalizedResource = await this.buildCanonicalizedResource(resource);
    let canonicalizedHeaders = await this.buildCanonicalizedHeaders(httpRequest.headers);
    stringToSign = `${request.method}\n${contentMd5}\n${contentType}\n${httpRequest.headers["date"]}\n${canonicalizedHeaders}${canonicalizedResource}`;
    let signature = EncodeUtil.base64EncodeToString(SignatureUtil.HmacSHA256Sign(stringToSign, accessKeySecret));
    httpRequest.headers["Authorization"] = `FC ${accessKeyId}:${signature}`;
    return httpRequest.headers;
  }

  async buildCanonicalizedResource(pathname: string): Promise<string> {
    let paths : string[] = String.split(pathname, `?`, 2);
    let canonicalizedResource : string = paths[0];
    let resources : string[] = [ ];
    if (Util.equalNumber(Array.size(paths), 2)) {
      resources = String.split(paths[1], "&", 0);
    }

    let sortedParams = Array.ascSort(resources);
    if (Util.equalNumber(Array.size(sortedParams), 0)) {
      return `${canonicalizedResource}\n`;
    }

    let subResources = Array.join(sortedParams, "\n");
    return `${canonicalizedResource}\n${subResources}`;
  }

  async buildCanonicalizedHeaders(headers: {[key: string ]: any}): Promise<string> {
    let canonicalizedHeaders : string = "";
    let keys = Map.keySet(headers);
    let sortedHeaders = Array.ascSort(keys);

    for (let header of sortedHeaders) {
      if (String.contains(String.toLower(header), "x-fc-")) {
        canonicalizedHeaders = `${canonicalizedHeaders}${String.toLower(header)}:${headers[header]}\n`;
      }

    }
    return canonicalizedHeaders;
  }

}
