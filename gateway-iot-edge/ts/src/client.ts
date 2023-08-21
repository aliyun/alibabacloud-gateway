// This file is auto-generated, don't edit it
import Url, * as $Url from '@alicloud/tea-url';
import Util, * as $Util from '@alicloud/tea-util';
import Number from '@darabonba/number';
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

export class ApiRequest extends $tea.Model {
  headers?: { [key: string]: string };
  pathname?: string;
  body?: { [key: string]: any };
  action?: string;
  static names(): { [key: string]: string } {
    return {
      headers: 'headers',
      pathname: 'pathname',
      body: 'body',
      action: 'action',
    };
  }

  static types(): { [key: string]: any } {
    return {
      headers: { 'type': 'map', 'keyType': 'string', 'valueType': 'string' },
      pathname: 'string',
      body: { 'type': 'map', 'keyType': 'string', 'valueType': 'any' },
      action: 'string',
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

  async execute(request: ApiRequest, runtime: $Util.RuntimeOptions): Promise<string> {
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
        request_.method = "POST";
        let url = Url.parseUrl(this._endpoint);
        request_.protocol = url.scheme;
        request_.port = Number.parseInt(url.host.port);
        request_.pathname = `${Util.defaultString(request.pathname, url.path.pathname)}/${request.action}`;
        request_.headers = {
          host: url.host.hostname,
          accept: "application/json",
          'content-type': "application/json; charset=utf-8",
          ...request.headers,
        };
        let params = {
          AppKey: this._appKey,
          Action: request.action,
          ...request.body,
        };
        request_.body = new $tea.BytesReadable(Util.toJSONString(params));
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


}
