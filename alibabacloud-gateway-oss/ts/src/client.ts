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
      "versionId"
    ];
    this._except_signed_params = [
      "list-type"
    ];
  }


  async modifyConfiguration(context: $SPI.InterceptorContext, attributeMap: $SPI.AttributeMap): Promise<void> {
    let config = context.configuration;
    config.endpoint = await this.getEndpoint(config.regionId, config.network, config.endpoint);
  }

  async modifyRequest(context: $SPI.InterceptorContext, attributeMap: $SPI.AttributeMap): Promise<void> {
    let request = context.request;
    let hostMap : {[key: string ]: string} = { };
    if (Util.isUnset(request.hostMap)) {
      hostMap = request.hostMap;
    }

    let bucketName = hostMap["bucket"];
    if (Util.isUnset(bucketName)) {
      bucketName = "";
    }

    let config = context.configuration;
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
        request.stream = new $tea.BytesReadable(XML.toXML(reqBodyMap));
        request.headers["content-type"] = "application/xml";
      } else if (String.equals(request.reqBodyType, "json")) {
        let reqBodyStr = Util.toJSONString(request.body);
        request.stream = new $tea.BytesReadable(reqBodyStr);
        request.headers["content-type"] = "application/json; charset=utf-8";
      } else if (String.equals(request.reqBodyType, "formData")) {
        let reqBodyForm = Util.assertAsMap(request.body);
        request.stream = new $tea.BytesReadable(OpenApiUtil.toForm(reqBodyForm));
        request.headers["content-type"] = "application/x-www-form-urlencoded";
      } else if (String.equals(request.reqBodyType, "binary")) {
        request.stream = OSSUtil.inject(request.stream, attributeMap.key);
        request.headers["content-type"] = "application/octet-stream";
      }

    }

    request.headers = {
      host: await this.getHost(config.endpointType, bucketName, config.endpoint),
      date: Util.getDateUTCString(),
      'user-agent': request.userAgent,
      ...request.headers,
    };
    request.headers["authorization"] = await this.getAuthorization(request.signatureVersion, bucketName, request.pathname, request.method, request.query, request.headers, accessKeyId, accessKeySecret);
  }

  async modifyResponse(context: $SPI.InterceptorContext, attributeMap: $SPI.AttributeMap): Promise<void> {
    let request = context.request;
    let response = context.response;
    let bodyStr : string = null;
    if (Util.is4xx(response.statusCode) || Util.is5xx(response.statusCode)) {
      bodyStr = await Util.readAsString(response.body);
      let respMap : {[key: string ]: any} = XML.parseXml(bodyStr, null);
      throw $tea.newError({
        code: respMap["Code"],
        message: respMap["Message"],
        data: {
          httpCode: response.statusCode,
          requestId: respMap["RequestId"],
          hostId: respMap["HostId"],
        },
      });
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
      if (String.equals(request.bodyType, "xml")) {
        bodyStr = await Util.readAsString(response.body);
        let result : {[key: string ]: any} = XML.parseXml(bodyStr, null);
        let list : string[] = Map.keySet(result);
        if (Util.equalNumber(Array.size(list), 1)) {
          let tmp = list[0];
          response.deserializedBody = result[tmp];
        } else {
          response.deserializedBody = result;
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

  async getAuthorization(signatureVersion: string, bucketName: string, pathname: string, method: string, query: {[key: string ]: string}, headers: {[key: string ]: string}, ak: string, secret: string): Promise<string> {
    if (Util.isUnset(signatureVersion) || String.equals(signatureVersion, "v1")) {
      return `OSS ${ak}:${await this.getSignatureV1(bucketName, pathname, method, query, headers, secret)}`;
    } else {
      return `OSS2 AccessKeyId:${ak},Signature:${await this.getSignatureV2(bucketName, pathname, method, query, headers, secret)}`;
    }

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
        let subResources : string[] = String.split(paths[1], "&", 0);

        for (let sub of subResources) {

          for (let excepts of this._except_signed_params) {
            if (!String.contains(sub, excepts)) {
              let item : string[] = String.split(sub, "&", 2);
              let key : string = item[0];
              let value : string = null;
              if (Util.equalNumber(Array.size(item), 2)) {
                value = item[1];
              }

              subResourcesMap[key] = value;
            }

          }
        }
      }

    }

    let subResourcesArray : string[] = Map.keySet(subResourcesMap);
    let newQueryList : string[] = subResourcesArray;
    if (!Util.isUnset(query)) {
      let queryList : string[] = Map.keySet(query);
      newQueryList = Array.concat(subResourcesArray, queryList);
    }

    let sortedParams = Array.ascSort(newQueryList);
    let separator : string = "?";

    for (let paramName of sortedParams) {
      if (Array.contains(this._default_signed_params, paramName)) {
        canonicalizedResource = `${canonicalizedResource}${separator}${paramName}`;
        if (!Util.isUnset(query[paramName])) {
          canonicalizedResource = `${canonicalizedResource}=${query[paramName]}`;
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
      if (String.contains(String.toLower(header), "x-oss-")) {
        canonicalizedHeaders = `${canonicalizedHeaders}${header}:${headers[header]}\n`;
      }

    }
    return canonicalizedHeaders;
  }

  async getSignatureV2(bucketName: string, pathname: string, method: string, query: {[key: string ]: string}, headers: {[key: string ]: string}, secret: string): Promise<string> {
    return "";
  }

}
