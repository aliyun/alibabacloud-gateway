// This file is auto-generated, don't edit it
import SPI, * as $SPI from '@alicloud/gateway-spi';
import Credential from '@alicloud/credentials';
import Util from '@alicloud/tea-util';
import OpenApiUtil from '@alicloud/openapi-util';
import EndpointUtil from '@alicloud/endpoint-util';
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
    let credential : Credential = request.credential;
    request.query = {
      Action: request.action,
      Format: "json",
      Version: request.version,
      Timestamp: OpenApiUtil.getTimestamp(),
      SignatureNonce: Util.getNonce(),
      ...request.query,
    };
    request.headers = {
      host: config.endpoint,
      'x-acs-version': request.version,
      'x-acs-action': request.action,
      'user-agent': request.userAgent,
    };
    let t : {[key: string ]: any} = null;
    if (!Util.isUnset(request.body)) {
      t = Util.assertAsMap(request.body);
      let tmp = Util.anyifyMapValue(OpenApiUtil.query(t));
      request.stream = new $tea.BytesReadable(Util.toFormString(tmp));
      request.headers["content-type"] = "application/x-www-form-urlencoded";
    }

    if (!Util.equalString(request.authType, "Anonymous")) {
      let accessKeyId = await credential.getAccessKeyId();
      let accessKeySecret = await credential.getAccessKeySecret();
      let securityToken = await credential.getSecurityToken();
      if (!Util.empty(securityToken)) {
        request.query["SecurityToken"] = securityToken;
      }

      request.query["SignatureMethod"] = "HMAC-SHA1";
      request.query["SignatureVersion"] = "1.0";
      request.query["AccessKeyId"] = accessKeyId;
      let signedParam = {
        ...request.query,
        ...OpenApiUtil.query(t),
      };
      request.query["Signature"] = OpenApiUtil.getRPCSignature(signedParam, request.method, accessKeySecret);
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

}
