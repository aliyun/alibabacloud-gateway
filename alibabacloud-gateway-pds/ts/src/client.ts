// This file is auto-generated, don't edit it
import SPI, * as $SPI from '@alicloud/gateway-spi';
import Credential from '@alicloud/credentials';
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
    request.headers = {
      date: Util.getDateUTCString(),
      host: config.endpoint,
      'x-acs-version': request.version,
      'x-acs-action': request.action,
      'user-agent': request.userAgent,
      'x-acs-signature-nonce': Util.getNonce(),
      'x-acs-signature-method': "HMAC-SHA1",
      'x-acs-signature-version': "1.0",
      accept: "application/json",
      ...request.headers,
    };
    if (!Util.isUnset(request.stream)) {
      let tmp = await Util.readAsBytes(request.stream);
      request.stream = new $tea.BytesReadable(tmp);
      request.headers["content-type"] = "application/octet-stream";
    } else {
      if (!Util.isUnset(request.body)) {
        if (Util.equalString(request.reqBodyType, "json")) {
          let jsonObj = Util.toJSONString(request.body);
          request.stream = new $tea.BytesReadable(jsonObj);
          request.headers["content-type"] = "application/json; charset=utf-8";
        } else {
          let m = Util.assertAsMap(request.body);
          let formObj = OpenApiUtil.toForm(m);
          request.stream = new $tea.BytesReadable(formObj);
          request.headers["content-type"] = "application/x-www-form-urlencoded";
        }

      }

    }

    if (!Util.equalString(request.authType, "Anonymous")) {
      let credential : Credential = request.credential;
      let accessKeyId = await credential.getAccessKeyId();
      let accessKeySecret = await credential.getAccessKeySecret();
      let securityToken = await credential.getSecurityToken();
      let bearerToken = credential.getBearerToken();
      if (!Util.empty(bearerToken)) {
        request.headers["x-acs-bearer-token"] = bearerToken;
        request.headers["Authorization"] = `Bearer ${bearerToken}`;
      } else {
        if (!Util.empty(securityToken)) {
          request.headers["x-acs-security-token"] = securityToken;
        }

        request.headers["Authorization"] = await this.getAuthorization(request.pathname, request.method, request.query, request.headers, accessKeyId, accessKeySecret);
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
      err["statusCode"] = response.statusCode;
      throw $tea.newError({
        code: `${this.defaultAny(err["Code"], err["code"])}`,
        message: `code: ${response.statusCode}, ${this.defaultAny(err["Message"], err["message"])} request id: ${requestId}`,
        data: err,
      });
    }

    if (!Util.isUnset(response.body) && !Util.equalNumber(response.statusCode, 204)) {
      if (Util.equalString(request.bodyType, "binary")) {
        response.deserializedBody = response.body;
      } else if (Util.equalString(request.bodyType, "byte")) {
        let byt = await Util.readAsBytes(response.body);
        response.deserializedBody = byt;
      } else if (Util.equalString(request.bodyType, "string")) {
        let str = await Util.readAsString(response.body);
        response.deserializedBody = str;
      } else if (Util.equalString(request.bodyType, "json")) {
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

  async getAuthorization(pathname: string, method: string, query: {[key: string ]: string}, headers: {[key: string ]: string}, ak: string, secret: string): Promise<string> {
    let signature = await this.getSignature(pathname, method, query, headers, secret);
    return `acs ${ak}:${signature}`;
  }

  async getSignature(pathname: string, method: string, query: {[key: string ]: string}, headers: {[key: string ]: string}, secret: string): Promise<string> {
    let stringToSign : string = "";
    let canonicalizedResource = await this.buildCanonicalizedResource(pathname, query);
    let canonicalizedHeaders = await this.buildCanonicalizedHeaders(headers);
    stringToSign = `${method}\n${canonicalizedHeaders}${canonicalizedResource}`;
    let signature = SignatureUtil.HmacSHA1Sign(stringToSign, secret);
    return EncodeUtil.base64EncodeToString(signature);
  }

  async buildCanonicalizedResource(pathname: string, query: {[key: string ]: string}): Promise<string> {
    let canonicalizedResource : string = pathname;
    if (!Util.isUnset(query)) {
      let queryArray : string[] = Map.keySet(query);
      let sortedQueryArray = Array.ascSort(queryArray);
      let separator : string = "?";

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

  async buildCanonicalizedHeaders(headers: {[key: string ]: string}): Promise<string> {
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

    let canonicalizedHeaders : string = `${accept}\n${contentMd5}\n${contentType}\n${date}\n`;
    let sortedHeaders : string[] = await this.getSignedHeaders(headers);

    for (let header of sortedHeaders) {
      let value = headers[header];
      let valueTrim = String.trim(value);
      canonicalizedHeaders = `${canonicalizedHeaders}${header}:${valueTrim}\n`;
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
      if (String.hasPrefix(lowerKey, "x-acs-")) {
        if (!String.contains(tmp, lowerKey)) {
          tmp = `${tmp}${separator}${lowerKey}`;
          separator = ";";
        }

      }

    }
    return String.split(tmp, ";", 0);
  }

}
