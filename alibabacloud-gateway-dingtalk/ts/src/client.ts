// This file is auto-generated, don't edit it
import SPI, * as $SPI from '@alicloud/gateway-spi';
import Util from '@alicloud/tea-util';
import * as $tea from '@alicloud/tea-typescript';


export default class Client extends SPI {

  constructor() {
    super();
  }


  async modifyConfiguration(context: $SPI.InterceptorContext, attributeMap: $SPI.AttributeMap): Promise<void> {
  }

  async modifyRequest(context: $SPI.InterceptorContext, attributeMap: $SPI.AttributeMap): Promise<void> {
    let request = context.request;
    let config = context.configuration;
    request.headers = {
      host: config.endpoint,
      'user-agent': request.userAgent,
      accept: "application/json",
      ...request.headers,
    };
    if (!Util.isUnset(request.body)) {
      let jsonObj = Util.toJSONString(request.body);
      request.stream = new $tea.BytesReadable(jsonObj);
      request.headers["content-type"] = "application/json; charset=utf-8";
    }

  }

  async modifyResponse(context: $SPI.InterceptorContext, attributeMap: $SPI.AttributeMap): Promise<void> {
    let request = context.request;
    let response = context.response;
    if (Util.is4xx(response.statusCode) || Util.is5xx(response.statusCode)) {
      let _res = await Util.readAsJSON(response.body);
      let err = Util.assertAsMap(_res);
      err["statusCode"] = response.statusCode;
      throw $tea.newError({
        code: `${this.defaultAny(err["Code"], err["code"])}`,
        message: `code: ${response.statusCode}, ${this.defaultAny(err["Message"], err["message"])} request id: ${this.defaultAny(err["RequestId"], err["requestid"])}`,
        data: err,
        description: `${this.defaultAny(err["Description"], err["description"])}`,
        accessDeniedDetail: this.defaultAny(err["AccessDeniedDetail"], err["accessdenieddetail"]),
      });
    }

    if (Util.equalNumber(response.statusCode, 204)) {
      await Util.readAsString(response.body);
    } else if (Util.equalString(request.bodyType, "binary")) {
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

  defaultAny(inputValue: any, defaultValue: any): any {
    if (Util.isUnset(inputValue)) {
      return defaultValue;
    }

    return inputValue;
  }

}
