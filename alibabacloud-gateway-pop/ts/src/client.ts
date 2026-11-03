// This file is auto-generated, don't edit it
import SPI, * as $SPI from '@alicloud/gateway-spi';
import Credential, * as $Credential from '@alicloud/credentials';
import * as runtime from './runtime';
import * as codec from './codec';
import * as $tea from '@alicloud/tea-typescript';


export default class Client extends SPI {
  _endpointSuffix: string;
  _signatureTypePrefix: string;
  _signPrefix: string;
  _sha256: string;
  _sm3: string;

  constructor() {
    super();
    // CLOUD4-
    this._signatureTypePrefix = "ACS4-";
    // cloud_v4
    this._signPrefix = "aliyun_v4";
    this._endpointSuffix = "aliyuncs.com";
    this._sha256 = `${this._signatureTypePrefix}HMAC-SHA256`;
    this._sm3 = `${this._signatureTypePrefix}HMAC-SM3`;
  }


  async modifyConfiguration(context: $SPI.InterceptorContext, attributeMap: $SPI.AttributeMap): Promise<void> {
    let request = context.request;
    let config = context.configuration;
    let endpointSuffix = this.resolveEndpointSuffix(attributeMap.key);
    config.endpoint = this.getEndpointWithSuffix(request.productId, config.regionId, config.endpointRule, config.network, config.suffix, config.endpointMap, config.endpoint, endpointSuffix);
  }

  async modifyRequest(context: $SPI.InterceptorContext, attributeMap: $SPI.AttributeMap): Promise<void> {
    let request = context.request;
    let config = context.configuration;
    let date = runtime.getTimestamp();
    request.headers = {
      host: config.endpoint,
      'x-acs-version': request.version,
      'user-agent': request.userAgent,
      'x-acs-date': date,
      'x-acs-signature-nonce': runtime.getNonce(),
      accept: "application/json",
      ...request.headers,
    };
    if (!runtime.empty(request.action)) {
      request.headers["x-acs-action"] = request.action;
    }

    let names = this.signatureNames(attributeMap.key);
    let signatureAlgorithm : string = runtime.defaultString(request.signatureAlgorithm, names.sha256);
    let hashedRequestPayload = codec.hexEncode(codec.hash(runtime.toBytes(""), signatureAlgorithm));
    if (!runtime.isUnset(request.stream)) {
      let tmp = await runtime.readAsBytes(request.stream);
      hashedRequestPayload = codec.hexEncode(codec.hash(tmp, signatureAlgorithm));
      request.stream = new $tea.BytesReadable(tmp);
      request.headers["content-type"] = "application/octet-stream";
    } else {
      if (!runtime.isUnset(request.body)) {
        if (runtime.equalString(request.reqBodyType, "byte")) {
          let byteObj = runtime.assertAsBytes(request.body);
          hashedRequestPayload = codec.hexEncode(codec.hash(byteObj, signatureAlgorithm));
          request.stream = new $tea.BytesReadable(byteObj);
        } else if (runtime.equalString(request.reqBodyType, "json")) {
          let jsonObj = runtime.toJSONString(request.body);
          hashedRequestPayload = codec.hexEncode(codec.hash(runtime.toBytes(jsonObj), signatureAlgorithm));
          request.stream = new $tea.BytesReadable(jsonObj);
          request.headers["content-type"] = "application/json; charset=utf-8";
        } else {
          let m = runtime.assertAsMap(request.body);
          let formObj = runtime.toForm(m);
          hashedRequestPayload = codec.hexEncode(codec.hash(runtime.toBytes(formObj), signatureAlgorithm));
          request.stream = new $tea.BytesReadable(formObj);
          request.headers["content-type"] = "application/x-www-form-urlencoded";
        }

      }

    }

    if (runtime.equalString(signatureAlgorithm, names.sm3)) {
      request.headers["x-acs-content-sm3"] = hashedRequestPayload;
    } else {
      request.headers["x-acs-content-sha256"] = hashedRequestPayload;
    }

    if (!runtime.equalString(request.authType, "Anonymous")) {
      let credential : Credential = request.credential;
      if (runtime.isUnset(credential)) {
        throw $tea.newError({
          code: "ParameterMissing",
          message: "'config.credential' can not be unset",
        });
      }

      let credentialModel = await credential.getCredential();
      if (!runtime.empty(credentialModel.providerName)) {
        request.headers["x-acs-credentials-provider"] = credentialModel.providerName;
      }

      let authType = credentialModel.type;
      if (runtime.equalString(authType, "bearer")) {
        let bearerToken = credentialModel.bearerToken;
        request.headers["x-acs-bearer-token"] = bearerToken;
        request.headers["x-acs-signature-type"] = "BEARERTOKEN";
        request.headers["Authorization"] = `Bearer ${bearerToken}`;
      } else if (runtime.equalString(authType, "id_token")) {
        let idToken = credentialModel.securityToken;
        request.headers["x-acs-zero-trust-idtoken"] = idToken;
      } else {
        let accessKeyId = credentialModel.accessKeyId;
        let accessKeySecret = credentialModel.accessKeySecret;
        let securityToken = credentialModel.securityToken;
        if (!runtime.empty(securityToken)) {
          request.headers["x-acs-accesskey-id"] = accessKeyId;
          request.headers["x-acs-security-token"] = securityToken;
        }

        let dateNew = codec.subString(date, 0, 10);
        dateNew = codec.replace(dateNew, "-", "", null);
        let region = this.getRegionWithSuffix(request.productId, config.endpoint, config.regionId, this.resolveEndpointSuffix(attributeMap.key));
        let signingkey = this.signingKey(signatureAlgorithm, accessKeySecret, request.productId, region, dateNew, names);
        request.headers["Authorization"] = this.authorization(request.pathname, request.method, request.query, request.headers, signatureAlgorithm, hashedRequestPayload, accessKeyId, signingkey, request.productId, region, dateNew, names);
      }

    }

  }

  async modifyResponse(context: $SPI.InterceptorContext, attributeMap: $SPI.AttributeMap): Promise<void> {
    let request = context.request;
    let response = context.response;
    if (runtime.is4xx(response.statusCode) || runtime.is5xx(response.statusCode)) {
      let err : {[key: string ]: any} = { };
      if (!runtime.isUnset(response.headers["content-type"]) && codec.contains(response.headers["content-type"], "text/xml")) {
        let _str = await runtime.readAsString(response.body);
        let respMap = runtime.parseXml(_str);
        err = runtime.assertAsMap(respMap["Error"]);
      } else {
        let _res = await runtime.readAsJSON(response.body);
        err = runtime.assertAsMap(_res);
      }

      let requestId = this.defaultAny(err["RequestId"], err["requestId"]);
      if (!runtime.isUnset(response.headers["x-acs-request-id"])) {
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

    if (runtime.equalNumber(response.statusCode, 204)) {
      await runtime.readAsString(response.body);
    } else if (runtime.equalString(request.bodyType, "binary")) {
      response.deserializedBody = response.body;
    } else if (runtime.equalString(request.bodyType, "byte")) {
      let byt = await runtime.readAsBytes(response.body);
      response.deserializedBody = byt;
    } else if (runtime.equalString(request.bodyType, "string")) {
      let str = await runtime.readAsString(response.body);
      response.deserializedBody = str;
    } else if (runtime.equalString(request.bodyType, "json")) {
      let obj = await runtime.readAsJSON(response.body);
      let res = runtime.assertAsMap(obj);
      response.deserializedBody = res;
    } else if (runtime.equalString(request.bodyType, "array")) {
      let arr = await runtime.readAsJSON(response.body);
      response.deserializedBody = arr;
    } else {
      response.deserializedBody = await runtime.readAsString(response.body);
    }

  }

  getEndpoint(productId: string, regionId: string, endpointRule: string, network: string, suffix: string, endpointMap: {[key: string ]: string}, endpoint: string): string {
    return this.getEndpointWithSuffix(productId, regionId, endpointRule, network, suffix, endpointMap, endpoint, this._endpointSuffix);
  }

  /** @internal */
  private getEndpointWithSuffix(productId: string, regionId: string, endpointRule: string, network: string, suffix: string, endpointMap: {[key: string ]: string}, endpoint: string, endpointSuffix: string): string {
    if (!runtime.empty(endpoint)) {
      return this.withEndpointSuffix(endpoint, endpointSuffix);
    }

    if (!runtime.isUnset(endpointMap) && !runtime.empty(endpointMap[regionId])) {
      return this.withEndpointSuffix(endpointMap[regionId], endpointSuffix);
    }

    return this.withEndpointSuffix(runtime.getEndpointRules(productId, regionId, endpointRule, network, suffix), endpointSuffix);
  }

  /** @internal */
  signatureNames(attributes: {[key: string]: string}): {signPrefix: string, sha256: string, sm3: string} {
    let prefix = "ACS4-";
    let signPrefix = "aliyun_v4";
    if (!runtime.isUnset(attributes)) {
      let requested = attributes["signatureTypePrefix"];
      if (!runtime.empty(requested)) {
        prefix = requested;
      }
      let requestedSign = attributes["signPrefix"];
      if (!runtime.empty(requestedSign)) {
        signPrefix = requestedSign;
      }
    }
    return {
      signPrefix: signPrefix,
      sha256: `${prefix}HMAC-SHA256`,
      sm3: `${prefix}HMAC-SM3`,
    };
  }

  resolveEndpointSuffix(attributes: {[key: string]: string}): string {
    if (!runtime.isUnset(attributes)) {
      let suffix = this.checkedEndpointSuffix(attributes["endpointSuffix"]);
      if (!runtime.empty(suffix)) {
        return suffix;
      }
    }
    let envSuffix = this.checkedEndpointSuffix(process.env.ALIBABA_CLOUD_ENDPOINT_SUFFIX);
    if (!runtime.empty(envSuffix)) {
      return envSuffix;
    }
    return "aliyuncs.com";
  }

  checkedEndpointSuffix(suffix: string): string {
    if (runtime.isUnset(suffix)) {
      return "";
    }
    suffix = suffix.trim().toLowerCase();
    if (!suffix) {
      return "";
    }
    if ((suffix !== "aliyuncs.com" && suffix !== "aliyuncs.alibaba") || !/^[a-z0-9]([a-z0-9-]*[a-z0-9])?(\.[a-z0-9]([a-z0-9-]*[a-z0-9])?)+$/.test(suffix)) {
      let error: any = new Error("endpointSuffix must be aliyuncs.com or aliyuncs.alibaba");
      error.code = "InvalidEndpointSuffix";
      throw error;
    }
    return suffix;
  }

  /** @internal */
  withEndpointSuffix(endpoint: string, endpointSuffix?: string): string {
    let suffix = endpointSuffix === undefined ? this._endpointSuffix : endpointSuffix;
    suffix = this.checkedEndpointSuffix(suffix) || "aliyuncs.com";
    if (runtime.empty(endpoint) || suffix === "aliyuncs.com") {
      return endpoint;
    }
    let old = ".aliyuncs.com";
    if (endpoint.length >= old.length && endpoint.substring(endpoint.length - old.length) === old) {
      return endpoint.substring(0, endpoint.length - old.length) + "." + suffix;
    }
    return endpoint;
  }

  defaultAny(inputValue: any, defaultValue: any): any {
    if (runtime.isUnset(inputValue)) {
      return defaultValue;
    }

    return inputValue;
  }

  getAuthorization(pathname: string, method: string, query: {[key: string ]: string}, headers: {[key: string ]: string}, signatureAlgorithm: string, payload: string, ak: string, signingkey: Buffer, product: string, region: string, date: string): string {
    return this.authorization(pathname, method, query, headers, signatureAlgorithm, payload, ak, signingkey, product, region, date, {
      signPrefix: this._signPrefix,
      sha256: this._sha256,
      sm3: this._sm3,
    });
  }

  private authorization(pathname: string, method: string, query: {[key: string ]: string}, headers: {[key: string ]: string}, signatureAlgorithm: string, payload: string, ak: string, signingkey: Buffer, product: string, region: string, date: string, names: {signPrefix: string, sha256: string, sm3: string}): string {
    let signature = this.signature(pathname, method, query, headers, signatureAlgorithm, payload, signingkey, names);
    let signedHeaders = this.getSignedHeaders(headers);
    let signedHeadersStr = codec.join(signedHeaders, ";");
    return `${signatureAlgorithm} Credential=${ak}/${date}/${region}/${product}/${names.signPrefix}_request,SignedHeaders=${signedHeadersStr},Signature=${signature}`;
  }

  getSignature(pathname: string, method: string, query: {[key: string ]: string}, headers: {[key: string ]: string}, signatureAlgorithm: string, payload: string, signingkey: Buffer): string {
    return this.signature(pathname, method, query, headers, signatureAlgorithm, payload, signingkey, {
      signPrefix: this._signPrefix,
      sha256: this._sha256,
      sm3: this._sm3,
    });
  }

  private signature(pathname: string, method: string, query: {[key: string ]: string}, headers: {[key: string ]: string}, signatureAlgorithm: string, payload: string, signingkey: Buffer, names: {signPrefix: string, sha256: string, sm3: string}): string {
    let canonicalURI : string = "/";
    if (!runtime.empty(pathname)) {
      canonicalURI = pathname;
    }

    let stringToSign : string = "";
    let canonicalizedResource = this.buildCanonicalizedResource(query);
    let canonicalizedHeaders = this.buildCanonicalizedHeaders(headers);
    let signedHeaders = this.getSignedHeaders(headers);
    this.validateSignedHeaders(signedHeaders);
    let signedHeadersStr = codec.join(signedHeaders, ";");
    stringToSign = `${method}\n${canonicalURI}\n${canonicalizedResource}\n${canonicalizedHeaders}\n${signedHeadersStr}\n${payload}`;
    let hex = codec.hexEncode(codec.hash(runtime.toBytes(stringToSign), signatureAlgorithm));
    stringToSign = `${signatureAlgorithm}\n${hex}`;
    let signature = runtime.toBytes("");
    if (runtime.equalString(signatureAlgorithm, names.sha256)) {
      signature = codec.hmacSha256ByBytes(stringToSign, signingkey);
    } else if (runtime.equalString(signatureAlgorithm, names.sm3)) {
      signature = codec.hmacSm3ByBytes(stringToSign, signingkey);
    }

    return codec.hexEncode(signature);
  }

  getSigningkey(signatureAlgorithm: string, secret: string, product: string, region: string, date: string): Buffer {
    return this.signingKey(signatureAlgorithm, secret, product, region, date, {
      signPrefix: this._signPrefix,
      sha256: this._sha256,
      sm3: this._sm3,
    });
  }

  private signingKey(signatureAlgorithm: string, secret: string, product: string, region: string, date: string, names: {signPrefix: string, sha256: string, sm3: string}): Buffer {
    let sc1 = `${names.signPrefix}${secret}`;
    let sc2 = runtime.toBytes("");
    if (runtime.equalString(signatureAlgorithm, names.sha256)) {
      sc2 = codec.hmacSha256(date, sc1);
    } else if (runtime.equalString(signatureAlgorithm, names.sm3)) {
      sc2 = codec.hmacSm3(date, sc1);
    }

    let sc3 = runtime.toBytes("");
    if (runtime.equalString(signatureAlgorithm, names.sha256)) {
      sc3 = codec.hmacSha256ByBytes(region, sc2);
    } else if (runtime.equalString(signatureAlgorithm, names.sm3)) {
      sc3 = codec.hmacSm3ByBytes(region, sc2);
    }

    let sc4 = runtime.toBytes("");
    if (runtime.equalString(signatureAlgorithm, names.sha256)) {
      sc4 = codec.hmacSha256ByBytes(product, sc3);
    } else if (runtime.equalString(signatureAlgorithm, names.sm3)) {
      sc4 = codec.hmacSm3ByBytes(product, sc3);
    }

    let hmac = runtime.toBytes("");
    if (runtime.equalString(signatureAlgorithm, names.sha256)) {
      hmac = codec.hmacSha256ByBytes(`${names.signPrefix}_request`, sc4);
    } else if (runtime.equalString(signatureAlgorithm, names.sm3)) {
      hmac = codec.hmacSm3ByBytes(`${names.signPrefix}_request`, sc4);
    }

    return hmac;
  }

  getRegion(product: string, endpoint: string, regionId: string): string {
    return this.getRegionWithSuffix(product, endpoint, regionId, this._endpointSuffix);
  }

  /** @internal */
  private getRegionWithSuffix(product: string, endpoint: string, regionId: string, endpointSuffix: string): string {
    if (!runtime.empty(regionId)) {
      return regionId;
    }

    let region = "center";
    if (runtime.empty(product) || runtime.empty(endpoint)) {
      return region;
    }

    let strs : string[] = codec.split(endpoint, ":", null);
    let withoutPort : string = strs[0];
    let preRegion : string = this.hostWithoutDomain(withoutPort, endpointSuffix);
    let nodes = codec.split(preRegion, ".", null);
    if (runtime.equalNumber(codec.size(nodes), 2)) {
      region = nodes[1];
    }

    return region;
  }

  /** @internal */
  hostWithoutDomain(host: string, suffix: string): string {
    let domains = [suffix, "aliyuncs.com"];
    let seen: {[key: string]: boolean} = {};
    for (let domain of domains) {
      if (runtime.empty(domain) || seen[domain]) {
        continue;
      }
      seen[domain] = true;
      let needle = `.${domain}`;
      if (host.length >= needle.length && host.substring(host.length - needle.length) === needle) {
        return host.substring(0, host.length - needle.length);
      }
    }
    return host;
  }

  buildCanonicalizedResource(query: {[key: string ]: string}): string {
    let canonicalizedResource : string = "";
    if (!runtime.isUnset(query)) {
      let queryArray : string[] = codec.keySet(query);
      let sortedQueryArray = codec.ascSort(queryArray);
      let separator : string = "";

      for (let key of sortedQueryArray) {
        canonicalizedResource = `${canonicalizedResource}${separator}${codec.percentEncode(key)}=`;
        if (!runtime.empty(query[key])) {
          canonicalizedResource = `${canonicalizedResource}${codec.percentEncode(query[key])}`;
        }

        separator = "&";
      }
    }

    return canonicalizedResource;
  }

  buildCanonicalizedHeaders(headers: {[key: string ]: string}): string {
    // lower header key
    let headersArray : string[] = codec.keySet(headers);
    let newHeaders : {[key: string ]: string} = { };

    for (let key of headersArray) {
      let lowerKey = codec.toLower(key);
      let value = headers[key];
      if (!runtime.isUnset(value)) {
        if (runtime.isUnset(newHeaders[lowerKey]) || codec.equals(lowerKey, "host")) {
          newHeaders[lowerKey] = codec.trim(value);
        } else {
          newHeaders[lowerKey] = `${newHeaders[lowerKey]},${codec.trim(value)}`;
        }

      }

    }
    let canonicalizedHeaders : string = "";
    let sortedHeaders : string[] = this.getSignedHeaders(headers);

    for (let header of sortedHeaders) {
      canonicalizedHeaders = `${canonicalizedHeaders}${header}:${newHeaders[header]}\n`;
    }
    return canonicalizedHeaders;
  }

  validateSignedHeaders(signedHeaders: string[]): void {
    if (!codec.listContains(signedHeaders, "host") || !codec.listContains(signedHeaders, "x-acs-date")) {
      throw $tea.newError({
        code: "InvalidSignedHeaders",
        message: "signed headers must include host and x-acs-date",
      });
    }
  }

  getSignedHeaders(headers: {[key: string ]: string}): string[] {
    let headersArray : string[] = codec.keySet(headers);
    let newHeadersArray : string[] = [ ];

    for (let key of headersArray) {
      let lowerKey = codec.toLower(key);
      let value = headers[key];
      if (!runtime.isUnset(value)) {
        codec.append(newHeadersArray, lowerKey);
      }

    }
    let sortedHeadersArray = codec.ascSort(newHeadersArray);
    let result : string[] = [ ];

    for (let key of sortedHeadersArray) {
      if (codec.hasPrefix(key, "x-acs-") || codec.equals(key, "host") || codec.equals(key, "content-type")) {
        if (!codec.listContains(result, key)) {
          codec.append(result, key);
        }

      }

    }
    return result;
  }

}
