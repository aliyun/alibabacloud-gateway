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
import { Readable } from 'stream';
import * as $tea from '@alicloud/tea-typescript';


export default class Client extends SPI {
  _default_signed_params: string[];
  _except_signed_params: string[];

  constructor() {
    super();
    this._default_signed_params = [
      "response-content-type",
      "response-content-language",
      "response-cache-control",
      "logging",
      "response-content-encoding",
      "acl",
      "uploadId",
      "uploads",
      "partNumber",
      "group",
      "link",
      "delete",
      "website",
      "location",
      "objectInfo",
      "objectMeta",
      "response-expires",
      "response-content-disposition",
      "cors",
      "lifecycle",
      "restore",
      "qos",
      "referer",
      "stat",
      "bucketInfo",
      "append",
      "position",
      "security-token",
      "live",
      "comp",
      "status",
      "vod",
      "startTime",
      "endTime",
      "x-oss-process",
      "symlink",
      "callback",
      "callback-var",
      "tagging",
      "encryption",
      "versions",
      "versioning",
      "versionId",
      "policy",
      "requestPayment",
      "x-oss-traffic-limit",
      "qosInfo",
      "asyncFetch",
      "x-oss-request-payer",
      "sequential",
      "inventory",
      "inventoryId",
      "continuation-token",
      "callback",
      "callback-var",
      "worm",
      "wormId",
      "wormExtend",
      "replication",
      "replicationLocation",
      "replicationProgress",
      "transferAcceleration",
      "cname",
      "metaQuery",
      "x-oss-ac-source-ip",
      "x-oss-ac-subnet-mask",
      "x-oss-ac-vpc-id",
      "x-oss-ac-forward-allow",
      "resourceGroup",
      "style",
      "styleName",
      "x-oss-async-process",
      "rtc",
      "accessPoint",
      "accessPointPolicy",
      "httpsConfig",
      "regionsV2",
      "publicAccessBlock",
      "policyStatus",
      "redundancyTransition",
      "redundancyType",
      "redundancyProgress",
      "dataAccelerator",
      "verbose",
      "accessPointForObjectProcess",
      "accessPointConfigForObjectProcess",
      "accessPointPolicyForObjectProcess",
      "bucketArchiveDirectRead",
      "responseHeader",
      "userDefinedLogFieldsConfig"
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
    if (Util.isUnset(regionId) || Util.empty(regionId)) {
      regionId = await this.getRegionIdFromEndpoint(config.endpoint);
    }

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
        // for python:
        // xml_str = OSS_UtilClient.to_xml(req_body_map)
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

    let host = await this.getHost(config.endpointType, bucketName, config.endpoint, context);
    request.headers = {
      host: host,
      date: Util.getDateUTCString(),
      'user-agent': request.userAgent,
      ...request.headers,
    };
    let originPath = request.pathname;
    let originQuery = request.query;
    if (!Util.empty(originPath)) {
      let pathAndQueries : string[] = String.split(originPath, `?`, 2);
      request.pathname = pathAndQueries[0];
      if (Util.equalNumber(Array.size(pathAndQueries), 2)) {
        let pathQueries : string[] = String.split(pathAndQueries[1], "&", null);

        for (let sub of pathQueries) {
          let item : string[] = String.split(sub, "=", null);
          let queryKey : string = item[0];
          let queryValue : string = "";
          if (Util.equalNumber(Array.size(item), 2)) {
            queryValue = item[1];
          }

          if (Util.empty(originQuery[queryKey])) {
            request.query[queryKey] = queryValue;
          }

        }
      }

    }

    let signatureVersion = Util.defaultString(request.signatureVersion, "v1");
    request.headers["authorization"] = await this.getAuthorization(signatureVersion, bucketName, request.pathname, request.method, request.query, request.headers, accessKeyId, accessKeySecret, regionId);
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
            AccessDeniedDetail: err["AccessDeniedDetail"],
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
          // var result : any = OSS_Util.parseXml(bodyStr, request.action);
          let result : any = XML.parseXml(bodyStr, null);
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

  async getRegionIdFromEndpoint(endpoint: string): Promise<string> {
    if (!Util.empty(endpoint)) {
      if (String.hasPrefix(endpoint, "oss-") && String.hasSuffix(endpoint, ".aliyuncs.com")) {
        let idx = String.index(endpoint, ".aliyuncs.com");
        return String.subString(endpoint, 4, idx);
      }

    }

    return "cn-hangzhou";
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

  async getHost(endpointType: string, bucketName: string, endpoint: string, context: $SPI.InterceptorContext): Promise<string> {
    if (String.contains(endpoint, ".mgw.aliyuncs.com") && !Util.isUnset(context.request.hostMap["userid"])) {
      return `${context.request.hostMap["userid"]}.${endpoint}`;
    }

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
      objectName = pathname;
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
    let canonicalizedResource : string = pathname;
    let queryKeys : string[] = Map.keySet(query);
    let sortedParams = Array.ascSort(queryKeys);
    let separator : string = "?";

    for (let paramName of sortedParams) {
      if (Array.contains(this._default_signed_params, paramName) || String.hasPrefix(paramName, "x-oss-")) {
        canonicalizedResource = `${canonicalizedResource}${separator}${paramName}`;
        if (!Util.empty(query[paramName])) {
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
      if (String.contains(String.toLower(header), "x-oss-") && !Util.isUnset(headers[header])) {
        canonicalizedHeaders = `${canonicalizedHeaders}${header}:${headers[header]}\n`;
      }

    }
    return canonicalizedHeaders;
  }

  async getSignatureV2(bucketName: string, pathname: string, method: string, query: {[key: string ]: string}, headers: {[key: string ]: string}, secret: string): Promise<string> {
    return "";
  }

}
