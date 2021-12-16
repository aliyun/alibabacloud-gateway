// This file is auto-generated, don't edit it
import SPI, * as $SPI from '@alicloud/gateway-spi';
import Credential from '@alicloud/credentials';
import Util from '@alicloud/tea-util';
import OpenApiUtil from '@alicloud/openapi-util';
import String from '@alicloud/darabonba-string';
import Map from '@alicloud/darabonba-map';
import Array from '@alicloud/darabonba-array';
import EncodeUtil from '@alicloud/darabonba-encode-util';
import SignatureUtil from '@alicloud/darabonba-signature-util';
import * as $tea from '@alicloud/tea-typescript';


export default class Client extends SPI {

  constructor() {
    super();
  }


  async modifyConfiguration(context: $SPI.InterceptorContext, attributeMap: $SPI.AttributeMap): Promise<void> {
    let config = context.configuration;
    config.endpoint = await this.getEndpoint(config.regionId, config.network, config.endpoint);
  }

  async modifyRequest(context: $SPI.InterceptorContext, attributeMap: $SPI.AttributeMap): Promise<void> {
    let request = context.request;
    let hostMap : {[key: string ]: string} = { };
    if (!Util.isUnset(request.hostMap)) {
      hostMap = request.hostMap;
    }

    let project = hostMap["project"];
    let config = context.configuration;
    let credential : Credential = request.credential;
    let accessKeyId = await credential.getAccessKeyId();
    let accessKeySecret = await credential.getAccessKeySecret();
    let securityToken = await credential.getSecurityToken();
    if (!Util.empty(accessKeyId)) {
      request.headers["x-log-signaturemethod"] = "hmac-sha1";
    }

    if (!Util.empty(securityToken)) {
      request.headers["x-acs-security-token"] = securityToken;
    }

    if (!Util.isUnset(request.body)) {
      if (String.equals(request.reqBodyType, "protobuf")) {
        // var bodyMap = Util.assertAsMap(request.body);
        // 缺少body的Content-MD5计算，以及protobuf处理
        request.headers["content-type"] = "application/x-protobuf";
      } else if (String.equals(request.reqBodyType, "json")) {
        let bodyStr = Util.toJSONString(request.body);
        request.headers["content-md5"] = String.toUpper(EncodeUtil.hexEncode(SignatureUtil.MD5Sign(bodyStr)));
        request.stream = new $tea.BytesReadable(bodyStr);
        request.headers["content-type"] = "application/json";
      } else if (String.equals(request.reqBodyType, "formData")) {
        let str = Util.toJSONString(request.body);
        request.headers["content-md5"] = String.toUpper(EncodeUtil.hexEncode(SignatureUtil.MD5Sign(str)));
        request.stream = new $tea.BytesReadable(str);
        request.headers["content-type"] = "application/json";
      }

    }

    request.headers = {
      accept: "application/json",
      host: await this.getHost(config.network, project, config.endpoint),
      date: Util.getDateUTCString(),
      'user-agent': request.userAgent,
      'x-log-apiversion': "0.6.0",
      'x-log-bodyrawsize': "0",
      ...request.headers,
    };
    request.headers["authorization"] = await this.getAuthorization(request.pathname, request.method, request.query, request.headers, accessKeyId, accessKeySecret);
  }

  async modifyResponse(context: $SPI.InterceptorContext, attributeMap: $SPI.AttributeMap): Promise<void> {
    let request = context.request;
    let response = context.response;
    if (Util.is4xx(response.statusCode) || Util.is5xx(response.statusCode)) {
      let error = await Util.readAsJSON(response.body);
      let resMap = Util.assertAsMap(error);
      throw $tea.newError({
        code: resMap["errorCode"],
        message: resMap["errorMessage"],
        data: {
          httpCode: response.statusCode,
          requestId: response.headers["x-log-requestid"],
        },
      });
    }

    if (!Util.isUnset(response.body)) {
      if (Util.equalString(request.bodyType, "binary")) {
        response.deserializedBody = response.body;
      } else if (Util.equalString(request.bodyType, "byte")) {
        let byt = await Util.readAsBytes(response.body);
        response.deserializedBody = byt;
      } else if (Util.equalString(request.bodyType, "string")) {
        response.deserializedBody = await Util.readAsString(response.body);
      } else if (Util.equalString(request.bodyType, "json")) {
        let obj = await Util.readAsJSON(response.body);
        let res = Util.assertAsMap(obj);
        response.deserializedBody = res;
      } else if (Util.equalString(request.bodyType, "array")) {
        response.deserializedBody = await Util.readAsJSON(response.body);
      } else {
        response.deserializedBody = await Util.readAsString(response.body);
      }

    }

  }

  async getEndpoint(regionId: string, network: string, endpoint: string): Promise<string> {
    if (!Util.empty(endpoint)) {
      return endpoint;
    }

    if (Util.empty(regionId)) {
      regionId = "cn-hangzhou";
    }

    if (!Util.empty(network)) {
      if (String.equals(network, "intranet")) {
        return `${regionId}-intranet.log.aliyuncs.com`;
      } else if (String.equals(network, "accelerate")) {
        return `log-global.aliyuncs.com`;
      } else if (String.equals(network, "share")) {
        if (String.equals(regionId, "cn-hangzhou-corp") || String.equals(regionId, "cn-shanghai-corp")) {
          return `${regionId}.sls.aliyuncs.com`;
        } else if (String.equals(regionId, "cn-zhangbei-corp")) {
          return `zhangbei-corp-share.log.aliyuncs.com`;
        }

        return `${regionId}-share.log.aliyuncs.com`;
      }

    }

    return `${regionId}.log.aliyuncs.com`;
  }

  async getHost(network: string, project: string, endpoint: string): Promise<string> {
    if (Util.isUnset(project)) {
      return endpoint;
    }

    return `${project}.${endpoint}`;
  }

  async getAuthorization(pathname: string, method: string, query: {[key: string ]: string}, headers: {[key: string ]: string}, ak: string, secret: string): Promise<string> {
    return `LOG ${ak}:${await this.getSignature(pathname, method, query, headers, secret)}`;
  }

  async getSignature(pathname: string, method: string, query: {[key: string ]: string}, headers: {[key: string ]: string}, secret: string): Promise<string> {
    let resource : string = pathname;
    let stringToSign : string = "";
    let canonicalizedResource = await this.buildCanonicalizedResource(resource, query);
    let canonicalizedHeaders = await this.buildCanonicalizedHeaders(headers);
    stringToSign = `${method}\n${canonicalizedHeaders}${canonicalizedResource}`;
    return EncodeUtil.base64EncodeToString(SignatureUtil.HmacSHA1Sign(stringToSign, secret));
  }

  async buildCanonicalizedResource(pathname: string, query: {[key: string ]: string}): Promise<string> {
    let canonicalizedResource : string = pathname;
    if (!Util.isUnset(query)) {
      let queryList : string[] = Map.keySet(query);
      let sortedParams = Array.ascSort(queryList);
      let separator : string = "?";

      for (let paramName of sortedParams) {
        canonicalizedResource = `${canonicalizedResource}${separator}${paramName}`;
        if (!Util.isUnset(query[paramName])) {
          canonicalizedResource = `${canonicalizedResource}=${query[paramName]}`;
        }

        separator = "&";
      }
    }

    return canonicalizedResource;
  }

  async buildCanonicalizedHeaders(headers: {[key: string ]: string}): Promise<string> {
    let canonicalizedHeaders : string = "";
    let contentType = headers["content-type"];
    if (Util.isUnset(contentType)) {
      contentType = "";
    }

    let contentMd5 = headers["content-md5"];
    if (Util.isUnset(contentMd5)) {
      contentMd5 = "";
    }

    canonicalizedHeaders = `${canonicalizedHeaders}${contentMd5}\n${contentType}\n${headers["date"]}\n`;
    let keys = Map.keySet(headers);
    let sortedHeaders = Array.ascSort(keys);

    for (let header of sortedHeaders) {
      if (String.contains(String.toLower(header), "x-log-") || String.contains(String.toLower(header), "x-acs-")) {
        canonicalizedHeaders = `${canonicalizedHeaders}${header}:${headers[header]}\n`;
      }

    }
    return canonicalizedHeaders;
  }

}
