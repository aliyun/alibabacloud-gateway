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
import SLS_Util from '@alicloud/gateway-sls-util';
import { Readable } from 'stream';
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
    let credentialModel = await credential.getCredential();
    let accessKeyId = credentialModel.accessKeyId;
    let accessKeySecret = credentialModel.accessKeySecret;
    let securityToken = credentialModel.securityToken;
    if (!Util.empty(securityToken)) {
      request.headers["x-acs-security-token"] = securityToken;
    }

    let signatureVersion = Util.defaultString(request.signatureVersion, "v1");
    let contentHash = "";
    if (!Util.isUnset(request.body)) {
      if (String.equals(request.reqBodyType, "protobuf")) {
        // var bodyMap = Util.assertAsMap(request.body);
        // 缺少body的Content-MD5计算，以及protobuf处理
        request.headers["content-type"] = "application/x-protobuf";
      } else if (String.equals(request.reqBodyType, "json")) {
        let bodyStr = Util.toJSONString(request.body);
        contentHash = await this.MakeContentHash(bodyStr, signatureVersion);
        request.stream = new $tea.BytesReadable(bodyStr);
        request.headers["content-type"] = "application/json";
      } else if (String.equals(request.reqBodyType, "formData")) {
        let str = Util.toJSONString(request.body);
        contentHash = await this.MakeContentHash(str, signatureVersion);
        request.stream = new $tea.BytesReadable(str);
        request.headers["content-type"] = "application/json";
      }

    }

    let host = await this.getHost(config.network, project, config.endpoint);
    request.headers = {
      accept: "application/json",
      host: host,
      'user-agent': request.userAgent,
      'x-log-apiversion': "0.6.0",
      'x-log-bodyrawsize': "0",
      ...request.headers,
    };
    await this.buildRequest(context);
    // move param in path to query
    if (String.equals(signatureVersion, "v4")) {
      if (Util.empty(contentHash)) {
        contentHash = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855";
      }

      let date = await this.getDateISO8601();
      request.headers["x-log-date"] = date;
      request.headers["x-log-content-sha256"] = contentHash;
      request.headers["authorization"] = await this.getAuthorizationV4(context, date, contentHash, accessKeyId, accessKeySecret);
      return ;
    }

    if (!Util.empty(accessKeyId)) {
      request.headers["x-log-signaturemethod"] = "hmac-sha256";
    }

    request.headers["date"] = Util.getDateUTCString();
    request.headers["content-md5"] = contentHash;
    request.headers["authorization"] = await this.getAuthorization(request.pathname, request.method, request.query, request.headers, accessKeyId, accessKeySecret);
  }

  async MakeContentHash(content: string, signatureVersion: string): Promise<string> {
    if (String.equals(signatureVersion, "v4")) {
      if (Util.empty(content)) {
        return "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855";
      }

      return String.toLower(EncodeUtil.hexEncode(EncodeUtil.hash(Util.toBytes(content), "SLS4-HMAC-SHA256")));
    }

    return String.toUpper(EncodeUtil.hexEncode(SignatureUtil.MD5Sign(content)));
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
        accessDeniedDetail: resMap["accessDeniedDetail"],
        data: {
          httpCode: response.statusCode,
          requestId: response.headers["x-log-requestid"],
          statusCode: response.statusCode,
        },
      });
    }

    if (!Util.isUnset(response.body)) {
      let bodyrawSize = response.headers["x-log-bodyrawsize"];
      let compressType = response.headers["x-log-compresstype"];
      let uncompressedData : Readable = response.body;
      if (!Util.isUnset(bodyrawSize) && !Util.isUnset(compressType)) {
        uncompressedData = await SLS_Util.readAndUncompressBlock(response.body, compressType, bodyrawSize);
      }

      if (Util.equalString(request.bodyType, "binary")) {
        response.deserializedBody = uncompressedData;
      } else if (Util.equalString(request.bodyType, "byte")) {
        let byt = await Util.readAsBytes(uncompressedData);
        response.deserializedBody = byt;
      } else if (Util.equalString(request.bodyType, "string")) {
        response.deserializedBody = await Util.readAsString(uncompressedData);
      } else if (Util.equalString(request.bodyType, "json")) {
        let obj = await Util.readAsJSON(uncompressedData);
        // var res = Util.assertAsMap(obj);
        response.deserializedBody = obj;
      } else if (Util.equalString(request.bodyType, "array")) {
        response.deserializedBody = await Util.readAsJSON(uncompressedData);
      } else {
        response.deserializedBody = await Util.readAsString(uncompressedData);
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
    let sign = await this.getSignature(pathname, method, query, headers, secret);
    return `LOG ${ak}:${sign}`;
  }

  async getSignature(pathname: string, method: string, query: {[key: string ]: string}, headers: {[key: string ]: string}, secret: string): Promise<string> {
    let resource : string = pathname;
    let stringToSign : string = "";
    let canonicalizedResource = await this.buildCanonicalizedResource(resource, query);
    let canonicalizedHeaders = await this.buildCanonicalizedHeaders(headers);
    stringToSign = `${method}\n${canonicalizedHeaders}${canonicalizedResource}`;
    return EncodeUtil.base64EncodeToString(SignatureUtil.HmacSHA256Sign(stringToSign, secret));
  }

  async buildCanonicalizedResource(pathname: string, query: {[key: string ]: string}): Promise<string> {
    let canonicalizedResource : string = pathname;
    if (!Util.isUnset(query)) {
      let queryList : string[] = Map.keySet(query);
      let sortedParams = Array.ascSort(queryList);
      let separator : string = "?";

      for (let paramName of sortedParams) {
        canonicalizedResource = `${canonicalizedResource}${separator}${paramName}`;
        let paramValue = query[paramName];
        if (!Util.isUnset(paramValue)) {
          canonicalizedResource = `${canonicalizedResource}=${paramValue}`;
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

  async getAuthorizationV4(context: $SPI.InterceptorContext, date: string, contentHash: string, accessKeyId: string, accessKeySecret: string): Promise<string> {
    let region = await this.getRegion(context);
    let headerStr = await this.getSignedHeaderStrV4(context.request.headers);
    let dateShort = String.subString(date, 0, 8);
    dateShort = String.replace(dateShort, "T", "", null);
    // for fix php sdk bug
    let scope = `${dateShort}/${region}/sls/aliyun_v4_request`;
    let signingkey = await this.getSigningkeyV4("SLS4-HMAC-SHA256", accessKeySecret, region, dateShort);
    let signature = await this.getSignatureV4(context, "SLS4-HMAC-SHA256", headerStr, date, scope, contentHash, signingkey);
    return `${"SLS4-HMAC-SHA256"} Credential=${accessKeyId}/${scope},Signature=${signature}`;
  }

  async getSigningkeyV4(signatureAlgorithm: string, secret: string, region: string, date: string): Promise<Buffer> {
    let sc1 = `aliyun_v4${secret}`;
    let sc2 = SignatureUtil.HmacSHA256Sign(date, sc1);
    let sc3 = SignatureUtil.HmacSHA256SignByBytes(region, sc2);
    let sc4 = SignatureUtil.HmacSHA256SignByBytes("sls", sc3);
    return SignatureUtil.HmacSHA256SignByBytes("aliyun_v4_request", sc4);
  }

  async getSignatureV4(context: $SPI.InterceptorContext, signatureAlgorithm: string, signedHeaderStr: string, date: string, scope: string, contentSha256: string, signingkey: Buffer): Promise<string> {
    let request = context.request;
    let canonicalURI = "/";
    if (!Util.empty(request.pathname)) {
      canonicalURI = request.pathname;
    }

    let resources = await this.buildCanonicalizedResourceV4(request.query);
    let headers = await this.buildCanonicalizedHeadersV4(request.headers, signedHeaderStr);
    let stringToHash = `${request.method}\n${canonicalURI}\n${resources}\n${headers}\n${signedHeaderStr}\n${contentSha256}`;
    let hex = EncodeUtil.hexEncode(EncodeUtil.hash(Util.toBytes(stringToHash), signatureAlgorithm));
    let stringToSign = `${signatureAlgorithm}\n${date}\n${scope}\n${hex}`;
    let signature = SignatureUtil.HmacSHA256SignByBytes(stringToSign, signingkey);
    return EncodeUtil.hexEncode(signature);
  }

  async buildCanonicalizedResourceV4(query: {[key: string ]: string}): Promise<string> {
    let canonicalizedResource : string = "";
    if (!Util.isUnset(query)) {
      let queryArray : string[] = Map.keySet(query);
      let sortedQueryArray = Array.ascSort(queryArray);
      let separator : string = "";

      for (let key of sortedQueryArray) {
        canonicalizedResource = `${canonicalizedResource}${separator}${key}`;
        if (!Util.empty(query[key])) {
          canonicalizedResource = `${canonicalizedResource}=${EncodeUtil.percentEncode(query[key])}`;
        }

        separator = "&";
      }
    }

    return canonicalizedResource;
  }

  async buildCanonicalizedHeadersV4(headers: {[key: string ]: string}, signedHeaderStr: string): Promise<string> {
    let canonicalizedHeaders : string = "";
    let signedHeaders : string[] = String.split(signedHeaderStr, ";", null);

    for (let header of signedHeaders) {
      canonicalizedHeaders = `${canonicalizedHeaders}${header}:${String.trim(headers[header])}\n`;
    }
    return canonicalizedHeaders;
  }

  async getSignedHeaderStrV4(headers: {[key: string ]: string}): Promise<string> {
    let headersArray : string[] = Map.keySet(headers);
    let sortedHeadersArray = Array.ascSort(headersArray);
    let tmp : string = "";
    let separator : string = "";

    for (let key of sortedHeadersArray) {
      let lowerKey = String.toLower(key);
      if (String.hasPrefix(lowerKey, "x-log-") || String.hasPrefix(lowerKey, "x-acs-") || String.equals(lowerKey, "host") || String.equals(lowerKey, "content-type")) {
        tmp = `${tmp}${separator}${lowerKey}`;
        separator = ";";
      }

    }
    return tmp;
  }

  async getRegion(context: $SPI.InterceptorContext): Promise<string> {
    let config = context.configuration;
    if (!Util.isUnset(config.regionId) && !Util.empty(config.regionId)) {
      return config.regionId;
    }

    // try parse region from endpoint
    // do not use String.subString, subString has bug in php implementation
    let region : string = String.replace(config.endpoint, ".log.aliyuncs.com", "", null);
    region = String.replace(region, ".sls.aliyuncs.com", "", null);
    if (String.equals(region, config.endpoint)) {
      throw $tea.newError({
        code: "ClientConfigError",
        message: "The regionId configuration of sls client is missing.",
      });
    }

    region = String.replace(region, "-share", "", null);
    region = String.replace(region, "-intranet", "", null);
    region = String.replace(region, "-vpc", "", null);
    return region;
  }

  // format: YYYYMMDDTHHMMSSZ
  async getDateISO8601(): Promise<string> {
    let date = OpenApiUtil.getTimestamp();
    // 2024-02-04T11:31:58Z
    date = String.replace(date, "-", "", null);
    return String.replace(date, ":", "", null);
  }

}
