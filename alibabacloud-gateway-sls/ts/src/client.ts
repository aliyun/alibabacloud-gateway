// This file is auto-generated, don't edit it
import SPI, * as $SPI from '@alicloud/gateway-spi';
import Credential, * as $Credential from '@alicloud/credentials';
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
import crypto from 'crypto';


export default class Client extends SPI {
  _respBodyDecompressType: {[key: string ]: string[]};
  _reqBodyCompressType: {[key: string ]: string[]};

  constructor() {
    super();
    this._respBodyDecompressType = {
      PullLogs: [
        "zstd",
        "lz4",
        "gzip"
      ],
      GetLogsV2: [
        "zstd",
        "lz4",
        "gzip"
      ],
      PreviewSPL: [
        "lz4"
      ],
    };
    this._reqBodyCompressType = {
      PutLogs: [
        "zstd",
        "lz4",
        "deflate"
      ],
      PreviewSPL: [
        "lz4"
      ],
    };
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
    let finalCompressType = await this.getFinalRequestCompressType(request.action, request.headers);
    let contentHash = "";
    // get body bytes
    let bodyBytes : Buffer = null;
    if (!Util.isUnset(request.body)) {
      if (String.equals(request.reqBodyType, "json") || String.equals(request.reqBodyType, "formData")) {
        request.headers["content-type"] = "application/json";
        let bodyStr = Util.toJSONString(request.body);
        bodyBytes = Util.toBytes(bodyStr);
      } else if (String.equals(request.reqBodyType, "binary")) {
        // content-type: application/octet-stream
        bodyBytes = Util.assertAsBytes(request.body);
      }

    }

    // get body raw size
    let bodyRawSize : string = "0";
    let rawSizeRef = request.headers["x-log-bodyrawsize"];
    // for php bug, Argument #1 ($value) could not be passed by reference
    if (!Util.isUnset(rawSizeRef)) {
      bodyRawSize = rawSizeRef;
    } else if (!Util.isUnset(request.body)) {
      bodyRawSize = `${await SLS_Util.bytesLength(bodyBytes)}`;
    }

    // compress if needed, and set body and hash
    if (!Util.isUnset(request.body)) {
      if (!Util.empty(finalCompressType)) {
        let compressed = await SLS_Util.compress(bodyBytes, finalCompressType);
        bodyBytes = compressed;
      }

      contentHash = await this.makeContentHash(bodyBytes, signatureVersion);
      request.stream = new $tea.BytesReadable(bodyBytes);
    }

    let host = await this.getHost(config.network, project, config.endpoint);
    request.headers = {
      accept: "application/json",
      host: host,
      'user-agent': request.userAgent,
      'x-log-apiversion': "0.6.0",
      ...request.headers,
    };
    request.headers["x-log-bodyrawsize"] = bodyRawSize;
    await this.setDefaultAcceptEncoding(request.action, request.headers);
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

  async getFinalRequestCompressType(action: string, headers: {[key: string ]: string}): Promise<string> {
    let compressType = headers["x-log-compresstype"];
    let rawSize = headers["x-log-bodyrawsize"];
    // for php bug, Argument #1 ($value) could not be passed by reference
    // 1. already compressed, has x-log-compresstype and x-log-bodyrawsize in header, we dont need compress here
    if (!Util.isUnset(compressType) && !Util.isUnset(rawSize)) {
      return "";
    }

    // 2. not compressed, but has x-log-compresstype in header, we need compress here
    if (!Util.isUnset(compressType)) {
      return compressType;
    }

    // 3. not compressed, in pre-defined api list, try use default supported compress type in order
    let encodings = this._reqBodyCompressType[action];
    if (!Util.isUnset(encodings)) {

      for (let encoding of encodings) {
        if (await SLS_Util.isCompressorAvailable(encoding)) {
          headers["x-log-compresstype"] = encoding;
          // set header x-log-compresstype
          return encoding;
        }

      }
    }

    // 4. otherwise we dont need compress here
    return "";
  }

  async setDefaultAcceptEncoding(action: string, headers: {[key: string ]: string}): Promise<void> {
    let acceptEncoding = headers["Accept-Encoding"];
    // for php warning, dont rm this line
    if (!Util.isUnset(acceptEncoding)) {
      return ;
    }

    let encodings = this._respBodyDecompressType[action];
    if (!Util.isUnset(encodings)) {
      for (let c of encodings) {
        if (await SLS_Util.isDecompressorAvailable(c)) {
          headers["Accept-Encoding"] = c;
          return ;
        }

      }
    }

  }

  async makeContentHash(content: Buffer, signatureVersion: string): Promise<string> {
    if (String.equals(signatureVersion, "v4")) {
      if (Util.isUnset(content) || Util.equalString(`${await SLS_Util.bytesLength(content)}`, "0")) {
        return "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855";
      }

      return String.toLower(EncodeUtil.hexEncode(EncodeUtil.hash(content, "SLS4-HMAC-SHA256")));
    }
    return String.toUpper(EncodeUtil.hexEncode(Client.MD5SignForBytes(content)));
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
    let sc2 = Client.HmacSha256(date, sc1);
    let sc3 = Client.HmacSha256(region, sc2);
    let sc4 = Client.HmacSha256("sls", sc3);
    return Client.HmacSha256("aliyun_v4_request", sc4);
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
    let signature = Client.HmacSha256(stringToSign, signingkey);
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

  /**
    * MD5 Signature, SignatureUtil has bug in js implemention for MD5SignForBytes.
    *
    * @param bytesToSign bytes
    * @return signed bytes
    */
  static MD5SignForBytes(bytesToSign) {
      return crypto.createHash('md5').update(bytesToSign).digest();
  }

  static HmacSha256(data: string, secret: Buffer | string): Buffer {
    const hmac = crypto.createHmac('sha256', secret);
    hmac.update(data);
    return hmac.digest();
  }
}
