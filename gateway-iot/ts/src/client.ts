// This file is auto-generated, don't edit it
import Url, * as $Url from '@alicloud/tea-url';
import Util, * as $Util from '@alicloud/tea-util';
import OpenApiUtil from '@alicloud/openapi-util';
import Credential from '@alicloud/credentials';
import Number from '@darabonba/number';
import Map from '@alicloud/darabonba-map';
import Array from '@alicloud/darabonba-array';
import EncodeUtil from '@alicloud/darabonba-encode-util';
import SignatureUtil from '@alicloud/darabonba-signature-util';
import * as $tea from '@alicloud/tea-typescript';

export class Config extends $tea.Model {
  appKey: string;
  appSecret: string;
  endpoint: string;
  static names(): { [key: string]: string } {
    return {
      appKey: 'appKey',
      appSecret: 'appSecret',
      endpoint: 'endpoint',
    };
  }

  static types(): { [key: string]: any } {
    return {
      appKey: 'string',
      appSecret: 'string',
      endpoint: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}

export class OpenApiRequest extends $tea.Model {
  headers?: { [key: string]: string };
  query?: { [key: string]: string };
  body?: any;
  pathname?: string;
  method?: string;
  static names(): { [key: string]: string } {
    return {
      headers: 'headers',
      query: 'query',
      body: 'body',
      pathname: 'pathname',
      method: 'method',
    };
  }

  static types(): { [key: string]: any } {
    return {
      headers: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      query: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      body: 'any',
      pathname: 'string',
      method: 'string',
    };
  }

  constructor(map?: { [key: string]: any }) {
    super(map);
  }
}


export default class Client {
  _appKey: string;
  _appSecret: string;
  _endpoint: string;

  constructor(config: Config) {
    this._appKey = config.appKey;
    this._appSecret = config.appSecret;
    this._endpoint = config.endpoint;
  }

  async execute(request: OpenApiRequest, runtime: $Util.RuntimeOptions): Promise<string> {
    let _runtime: { [key: string]: any } = {
      timeouted: "retry",
      readTimeout: runtime.readTimeout,
      connectTimeout: runtime.connectTimeout,
      maxIdleConns: runtime.maxIdleConns,
      retry: {
        retryable: runtime.autoretry,
        maxAttempts: runtime.maxAttempts,
      },
      ignoreSSL: runtime.ignoreSSL,
    }

    let _lastRequest = null;
    let _now = Date.now();
    let _retryTimes = 0;
    while ($tea.allowRetry(_runtime['retry'], _retryTimes, _now)) {
      if (_retryTimes > 0) {
        let _backoffTime = $tea.getBackoffTime(_runtime['backoff'], _retryTimes);
        if (_backoffTime > 0) {
          await $tea.sleep(_backoffTime);
        }
      }

      _retryTimes = _retryTimes + 1;
      try {
        let request_ = new $tea.Request();
        let method : string = Util.defaultString(request.method, "POST");
        request_.method = method;
        let url = Url.parseUrl(this._endpoint);
        request_.protocol = url.scheme;
        request_.pathname = Util.defaultString(request.pathname, url.path.pathname);
        request_.port = Number.parseInt(url.host.port);
        request_.headers = {
          host: url.host.hostname,
          ...request.headers,
        };
        let params : {[key: string ]: string} = { };
        if (!Util.isUnset(request.body)) {
          let tmp = Util.assertAsMap(request.body);
          params = OpenApiUtil.query(tmp);
        }

        params = {
          AppKey: this._appKey,
          ...request.query,
          ...params,
        };
        params["Signature"] = Client.sign(method, this._appSecret, params);
        if (Util.equalString(method, "GET")) {
          request_.query = params;
        } else {
          let formObj = Util.toFormString(params);
          request_.body = new $tea.BytesReadable(formObj);
          request_.headers["content-type"] = "application/x-www-form-urlencoded";
        }

        _lastRequest = request_;
        let response_ = await $tea.doAction(request_, _runtime);

        if (Util.is4xx(response_.statusCode) || Util.is5xx(response_.statusCode)) {
          let res = await Util.readAsJSON(response_.body);
          let err = Util.assertAsMap(res);
          err["statusCode"] = response_.statusCode;
          throw $tea.newError({
            code: `${err["Code"]}`,
            message: `statusCode: ${err["statusCode"]}, errorMessage: ${err["ErrorMessage"]}, requestId: ${err["RequestId"]}`,
            data: err,
          });
        }

        return await Util.readAsString(response_.body);
      } catch (ex) {
        if ($tea.isRetryable(ex)) {
          continue;
        }
        throw ex;
      }
    }

    throw $tea.newUnretryableError(_lastRequest);
  }

  static sign(method: string, appSecret: string, params: {[key: string ]: string}): string {
    let keys : string[] = Map.keySet(params);
    let sortedKeys = Array.ascSort(keys);
    let canonicalizedResource : string = "";
    let separator : string = "";

    for (let key of sortedKeys) {
      canonicalizedResource = `${canonicalizedResource}${separator}${EncodeUtil.percentEncode(key)}`;
      if (!Util.empty(params[key])) {
        canonicalizedResource = `${canonicalizedResource}=${EncodeUtil.percentEncode(params[key])}`;
      }

      separator = "&";
    }
    let signToString = `${method}&${EncodeUtil.percentEncode("/")}&${EncodeUtil.percentEncode(canonicalizedResource)}`;
    let signData = SignatureUtil.HmacSHA1Sign(signToString, `${appSecret}&`);
    return EncodeUtil.base64EncodeToString(signData);
  }

}
