// This file is auto-generated, don't edit it
import SPI, * as $SPI from '@alicloud/gateway-spi';
import Credential from '@alicloud/credentials';
import Util from '@alicloud/tea-util';
import OSSUtil from '@alicloud/oss-util';
import OpenApiUtil from '@alicloud/openapi-util';
import XML from '@alicloud/tea-xml';
import String from '@alicloud/darabonba-string';
import Map from '@alicloud/darabonba-map';
import Array from '@alicloud/darabonba-array';
import EncodeUtil from '@alicloud/darabonba-encode-util';
import SignatureUtil from '@alicloud/darabonba-signature-util';
import Time from '@darabonba/time';
import * as $tea from '@alicloud/tea-typescript';


export default class Client extends SPI {
  _default_signed_params: string[];
  _except_signed_params: string[];

  constructor() {
    super();
    this._default_signed_params = [
      "location",
      "cors",
      "objectMeta",
      "uploadId",
      "partNumber",
      "security-token",
      "position",
      "img",
      "style",
      "styleName",
      "replication",
      "replicationProgress",
      "replicationLocation",
      "cname",
      "qos",
      "startTime",
      "endTime",
      "symlink",
      "x-oss-process",
      "response-content-type",
      "response-content-language",
      "response-expires",
      "response-cache-control",
      "response-content-disposition",
      "response-content-encoding",
      "udf",
      "udfName",
      "udfImage",
      "udfId",
      "udfImageDesc",
      "udfApplication",
      "udfApplicationLog",
      "restore",
      "callback",
      "callback-var",
      "policy",
      "encryption",
      "versions",
      "versioning",
      "versionId",
      "wormId"
    ];
    this._except_signed_params = [
      "list-type",
      "regions"
    ];
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

    let bucketName = hostMap["bucket"];
    if (Util.isUnset(bucketName)) {
      bucketName = "";
    }

    if (!Util.isUnset(request.headers["x-oss-meta-*"])) {
      let tmp : any = Util.parseJSON(request.headers["x-oss-meta-*"]);
      let mapData : {[key: string ]: any} = Util.assertAsMap(tmp);
      let metaData : {[key: string ]: string} = Util.stringifyMapValue(mapData);
      let metaKeySet : string[] = Map.keySet(metaData);
      request.headers["x-oss-meta-*"] = null;

      for (let key of metaKeySet) {
        let newKey = `x-oss-meta-${key}`;
        request.headers[newKey] = metaData[key];
      }
    }

    let config = context.configuration;
    let regionId = config.regionId;
    let credential : Credential = request.credential;
    let accessKeyId = await credential.getAccessKeyId();
    let accessKeySecret = await credential.getAccessKeySecret();
    let securityToken = await credential.getSecurityToken();
    if (!Util.empty(securityToken)) {
      request.headers["x-oss-security-token"] = securityToken;
    }

    if (!Util.isUnset(request.body)) {
      if (String.equals(request.reqBodyType, "xml")) {
        let reqBodyMap = Util.assertAsMap(request.body);
        let xmlStr = XML.toXML(reqBodyMap);
        request.stream = new $tea.BytesReadable(xmlStr);
        request.headers["content-type"] = "application/xml";
        request.headers["content-md5"] = EncodeUtil.base64EncodeToString(SignatureUtil.MD5Sign(xmlStr));
      } else if (String.equals(request.reqBodyType, "json")) {
        let reqBodyStr = Util.toJSONString(request.body);
        request.stream = new $tea.BytesReadable(reqBodyStr);
        request.headers["content-type"] = "application/json; charset=utf-8";
      } else if (String.equals(request.reqBodyType, "formData")) {
        let reqBodyForm = Util.assertAsMap(request.body);
        request.stream = new $tea.BytesReadable(OpenApiUtil.toForm(reqBodyForm));
        request.headers["content-type"] = "application/x-www-form-urlencoded";
      } else if (String.equals(request.reqBodyType, "binary")) {
        attributeMap.key = {
          crc: "",
          md5: "",
        };
        request.stream = OSSUtil.inject(request.stream, attributeMap.key);
        request.headers["content-type"] = "application/octet-stream";
      }

    }

    let host = await this.getHost(config.endpointType, bucketName, config.endpoint);
    request.headers = {
      host: host,
      date: Util.getDateUTCString(),
      'user-agent': request.userAgent,
      ...request.headers,
    };
    request.headers["authorization"] = await this.getAuthorization(request.signatureVersion, bucketName, request.pathname, request.method, request.query, request.headers, accessKeyId, accessKeySecret, regionId);
  }

  async modifyResponse(context: $SPI.InterceptorContext, attributeMap: $SPI.AttributeMap): Promise<void> {
    let request = context.request;
    let response = context.response;
    let bodyStr : string = null;
    if (Util.is4xx(response.statusCode) || Util.is5xx(response.statusCode)) {
      bodyStr = await Util.readAsString(response.body);
      if (!Util.empty(bodyStr)) {
        let respMap : {[key: string ]: any} = XML.parseXml(bodyStr, null);
        let err : {[key: string ]: any} = Util.assertAsMap(respMap["Error"]);
        throw $tea.newError({
          code: err["Code"],
          message: err["Message"],
          data: {
            statusCode: response.statusCode,
            requestId: err["RequestId"],
            ecCode: err["EC"],
            Recommend: err["RecommendDoc"],
            hostId: err["HostId"],
          },
        });
      } else {
        let headers : {[key: string ]: string} = response.headers;
        let requestId = headers["x-oss-request-id"];
        let ecCode = headers["x-oss-ec-code"];
        throw $tea.newError({
          code: response.statusCode,
          message: null,
          data: {
            statusCode: response.statusCode,
            requestId: `${requestId}`,
            ecCode: ecCode,
          },
        });
      }

    }

    let ctx : {[key: string ]: string} = attributeMap.key;
    if (!Util.isUnset(ctx)) {
      if (!Util.isUnset(ctx["crc"]) && !Util.isUnset(response.headers["x-oss-hash-crc64ecma"]) && !String.equals(ctx["crc"], response.headers["x-oss-hash-crc64ecma"])) {
        throw $tea.newError({
          code: "CrcNotMatched",
          data: {
            clientCrc: ctx["crc"],
            serverCrc: response.headers["x-oss-hash-crc64ecma"],
          },
        });
      }

      if (!Util.isUnset(ctx["md5"]) && !Util.isUnset(response.headers["content-md5"]) && !String.equals(ctx["md5"], response.headers["content-md5"])) {
        throw $tea.newError({
          code: "MD5NotMatched",
          data: {
            clientMD5: ctx["md5"],
            serverMD5: response.headers["content-md5"],
          },
        });
      }

    }

    if (!Util.isUnset(response.body)) {
      if (Util.equalNumber(response.statusCode, 204)) {
        await Util.readAsString(response.body);
      } else if (String.equals(request.bodyType, "xml")) {
        bodyStr = await Util.readAsString(response.body);
        response.deserializedBody = bodyStr;
        if (!Util.empty(bodyStr)) {
          let respStruct = await Client.getResponseBodySchema(request.action);
          let result : {[key: string ]: any} = XML.parseXml(bodyStr, respStruct);
          try {
            response.deserializedBody = Util.assertAsMap(result);
          } catch (error) {
            response.deserializedBody = result;
          }          
        }

      } else if (Util.equalString(request.bodyType, "binary")) {
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
      if (String.contains(network, "internal")) {
        return `oss-${regionId}-internal.aliyuncs.com`;
      } else if (String.contains(network, "ipv6")) {
        return `${regionId}oss.aliyuncs.com`;
      } else if (String.contains(network, "accelerate")) {
        return `oss-${network}.aliyuncs.com`;
      }

    }

    return `oss-${regionId}.aliyuncs.com`;
  }

  async getHost(endpointType: string, bucketName: string, endpoint: string): Promise<string> {
    if (Util.empty(bucketName)) {
      return endpoint;
    }

    let host : string = `${bucketName}.${endpoint}`;
    if (!Util.empty(endpointType)) {
      if (String.equals(endpointType, "ip")) {
        host = `${endpoint}/${bucketName}`;
      } else if (String.equals(endpointType, "cname")) {
        host = endpoint;
      }

    }

    return host;
  }

  async getAuthorization(signatureVersion: string, bucketName: string, pathname: string, method: string, query: {[key: string ]: string}, headers: {[key: string ]: string}, ak: string, secret: string, regionId: string): Promise<string> {
    let sign : string = "";
    if (!Util.isUnset(signatureVersion)) {
      if (String.equals(signatureVersion, "v1")) {
        sign = await this.getSignatureV1(bucketName, pathname, method, query, headers, secret);
        return `OSS ${ak}:${sign}`;
      }

      if (String.equals(signatureVersion, "v2")) {
        sign = await this.getSignatureV2(bucketName, pathname, method, query, headers, secret);
        return `OSS2 AccessKeyId:${ak},Signature:${sign}`;
      }

    }

    let dateTime = OpenApiUtil.getTimestamp();
    dateTime = String.replace(dateTime, "-", "", null);
    dateTime = String.replace(dateTime, ":", "", null);
    headers["x-oss-date"] = dateTime;
    headers["x-oss-content-sha256"] = "UNSIGNED-PAYLOAD";
    let onlyDate : string = String.subString(dateTime, 0, 8);
    let cred : string = `${ak}/${onlyDate}/${regionId}/oss/aliyun_v4_request`;
    sign = await this.getSignatureV4(bucketName, pathname, method, query, headers, onlyDate, regionId, secret);
    return `OSS4-HMAC-SHA256 Credential=${cred}, Signature=${sign}`;
  }

  async getSignKey(secret: string, onlyDate: string, regionId: string): Promise<Buffer> {
    let temp = `aliyun_v4${secret}`;
    let res = SignatureUtil.HmacSHA256Sign(onlyDate, temp);
    res = SignatureUtil.HmacSHA256SignByBytes(regionId, res);
    res = SignatureUtil.HmacSHA256SignByBytes("oss", res);
    res = SignatureUtil.HmacSHA256SignByBytes("aliyun_v4_request", res);
    return res;
  }

  async getSignatureV4(bucketName: string, pathname: string, method: string, query: {[key: string ]: string}, headers: {[key: string ]: string}, onlyDate: string, regionId: string, secret: string): Promise<string> {
    let signingkey : Buffer = await this.getSignKey(secret, onlyDate, regionId);
    let objectName : string = "/";
    let queryMap : {[key: string ]: string} = { };
    if (!Util.empty(pathname)) {
      let paths : string[] = String.split(pathname, `?`, 2);
      objectName = paths[0];
      if (Util.equalNumber(Array.size(paths), 2)) {
        let subResources : string[] = String.split(paths[1], "&", null);

        for (let sub of subResources) {
          let item : string[] = String.split(sub, "=", null);
          let key : string = item[0];
          key = EncodeUtil.percentEncode(key);
          key = String.replace(key, "+", "%20", null);
          let value : string = null;
          if (Util.equalNumber(Array.size(item), 2)) {
            value = EncodeUtil.percentEncode(item[1]);
            value = String.replace(value, "+", "%20", null);
          }

          // for go : queryMap[tea.StringValue(key)] = value
          queryMap[key] = value;
        }
      }

    }

    let canonicalizedUri : string = "/";
    if (!Util.empty(bucketName)) {
      canonicalizedUri = `/${bucketName}${objectName}`;
    }

    // for java:
    // String suffix = (!canonicalizedUri.equals("/") && canonicalizedUri.endsWith("/"))? "/" : "";
    // canonicalizedUri = com.aliyun.openapiutil.Client.getEncodePath(canonicalizedUri) + suffix;
    canonicalizedUri = OpenApiUtil.getEncodePath(canonicalizedUri);

    for (let queryKey of Map.keySet(query)) {
      let queryValue : string = null;
      if (!Util.empty(query[queryKey])) {
        queryValue = EncodeUtil.percentEncode(query[queryKey]);
        queryValue = String.replace(queryValue, "+", "%20", null);
      }

      queryKey = EncodeUtil.percentEncode(queryKey);
      queryKey = String.replace(queryKey, "+", "%20", null);
      // for go : queryMap[tea.StringValue(queryKey)] = queryValue
      queryMap[queryKey] = queryValue;
    }
    let canonicalizedQueryString = await this.buildCanonicalizedQueryStringV4(queryMap);
    let canonicalizedHeaders = await this.buildCanonicalizedHeadersV4(headers);
    let payload = "UNSIGNED-PAYLOAD";
    let canonicalRequest = `${method}\n${canonicalizedUri}\n${canonicalizedQueryString}\n${canonicalizedHeaders}\n\n${payload}`;
    let hex = EncodeUtil.hexEncode(EncodeUtil.hash(Util.toBytes(canonicalRequest), "ACS4-HMAC-SHA256"));
    let scope : string = `${onlyDate}/${regionId}/oss/aliyun_v4_request`;
    let stringToSign = `OSS4-HMAC-SHA256\n${headers["x-oss-date"]}\n${scope}\n${hex}`;
    let signature = SignatureUtil.HmacSHA256SignByBytes(stringToSign, signingkey);
    return EncodeUtil.hexEncode(signature);
  }

  async buildCanonicalizedQueryStringV4(queryMap: {[key: string ]: string}): Promise<string> {
    let canonicalizedQueryString : string = "";
    if (!Util.isUnset(queryMap)) {
      let queryArray : string[] = Map.keySet(queryMap);
      let sortedQueryArray = Array.ascSort(queryArray);
      let separator : string = "";

      for (let key of sortedQueryArray) {
        canonicalizedQueryString = `${canonicalizedQueryString}${separator}${key}`;
        if (!Util.empty(queryMap[key])) {
          canonicalizedQueryString = `${canonicalizedQueryString}=${queryMap[key]}`;
        }

        separator = "&";
      }
    }

    return canonicalizedQueryString;
  }

  async buildCanonicalizedHeadersV4(headers: {[key: string ]: string}): Promise<string> {
    let canonicalizedHeaders : string = "";
    let headersArray : string[] = Map.keySet(headers);
    let sortedHeadersArray = Array.ascSort(headersArray);

    for (let key of sortedHeadersArray) {
      let lowerKey = String.toLower(key);
      if (String.hasPrefix(lowerKey, "x-oss-") || String.equals(lowerKey, "content-type") || String.equals(lowerKey, "content-md5")) {
        canonicalizedHeaders = `${canonicalizedHeaders}${key}:${String.trim(headers[key])}\n`;
      }

    }
    return canonicalizedHeaders;
  }

  async getSignatureV1(bucketName: string, pathname: string, method: string, query: {[key: string ]: string}, headers: {[key: string ]: string}, secret: string): Promise<string> {
    let resource : string = "";
    let stringToSign : string = "";
    if (!Util.empty(bucketName)) {
      resource = `/${bucketName}`;
    }

    resource = `${resource}${pathname}`;
    let canonicalizedResource = await this.buildCanonicalizedResource(resource, query);
    let canonicalizedHeaders = await this.buildCanonicalizedHeaders(headers);
    stringToSign = `${method}\n${canonicalizedHeaders}${canonicalizedResource}`;
    return EncodeUtil.base64EncodeToString(SignatureUtil.HmacSHA1Sign(stringToSign, secret));
  }

  async buildCanonicalizedResource(pathname: string, query: {[key: string ]: string}): Promise<string> {
    let subResourcesMap : {[key: string ]: string} = { };
    let canonicalizedResource : string = pathname;
    if (!Util.empty(pathname)) {
      let paths : string[] = String.split(pathname, `?`, 2);
      canonicalizedResource = paths[0];
      if (Util.equalNumber(Array.size(paths), 2)) {
        let subResources : string[] = String.split(paths[1], "&", null);

        for (let sub of subResources) {
          let hasExcepts : boolean = false;

          for (let excepts of this._except_signed_params) {
            if (String.contains(sub, excepts)) {
              hasExcepts = true;
            }

          }
          if (!hasExcepts) {
            let item : string[] = String.split(sub, "=", null);
            let key : string = item[0];
            let value : string = null;
            if (Util.equalNumber(Array.size(item), 2)) {
              value = item[1];
            }

            // for go : subResourcesMap[tea.StringValue(key)] = value
            subResourcesMap[key] = value;
          }

        }
      }

    }

    let subResourcesArray : string[] = Map.keySet(subResourcesMap);
    let newQueryList : string[] = subResourcesArray;
    if (!Util.isUnset(query)) {
      let queryList : string[] = Map.keySet(query);
      newQueryList = Array.concat(queryList, subResourcesArray);
    }

    let sortedParams = Array.ascSort(newQueryList);
    let separator : string = "?";

    for (let paramName of sortedParams) {
      if (Array.contains(this._default_signed_params, paramName)) {
        canonicalizedResource = `${canonicalizedResource}${separator}${paramName}`;
        if (!Util.isUnset(query) && !Util.isUnset(query[paramName])) {
          canonicalizedResource = `${canonicalizedResource}=${query[paramName]}`;
        } else if (!Util.isUnset(subResourcesMap[paramName])) {
          canonicalizedResource = `${canonicalizedResource}=${subResourcesMap[paramName]}`;
        }

      } else if (Array.contains(subResourcesArray, paramName)) {
        canonicalizedResource = `${canonicalizedResource}${separator}${paramName}`;
        if (!Util.isUnset(subResourcesMap[paramName])) {
          canonicalizedResource = `${canonicalizedResource}=${subResourcesMap[paramName]}`;
        }

      }

      separator = "&";
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
      if (String.contains(String.toLower(header), "x-oss-") && !Util.isUnset(headers[header])) {
        canonicalizedHeaders = `${canonicalizedHeaders}${header}:${headers[header]}\n`;
      }

    }
    return canonicalizedHeaders;
  }

  async getSignatureV2(bucketName: string, pathname: string, method: string, query: {[key: string ]: string}, headers: {[key: string ]: string}, secret: string): Promise<string> {
    return "";
  }

  static async getResponseBodySchema(apiName: string): Promise<Object> {
    return null;
  }

}
