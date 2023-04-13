// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.iot;

import com.aliyun.tea.*;
import com.aliyun.tea.interceptor.InterceptorChain;
import com.aliyun.tea.interceptor.RuntimeOptionsInterceptor;
import com.aliyun.tea.interceptor.RequestInterceptor;
import com.aliyun.tea.interceptor.ResponseInterceptor;
import com.aliyun.gateway.iot.models.*;

public class Client {

    private final static InterceptorChain interceptorChain = InterceptorChain.create();

    public String _appKey;
    public String _appSecret;
    public String _endpoint;
    public Client(Config config) throws Exception {
        this._appKey = config.appKey;
        this._appSecret = config.appSecret;
        this._endpoint = config.endpoint;
    }

    public String execute(OpenApiRequest request, com.aliyun.teautil.models.RuntimeOptions runtime) throws Exception {
        TeaModel.validateParams(request, "request");
        java.util.Map<String, Object> runtime_ = TeaConverter.buildMap(
            new TeaPair("timeouted", "retry"),
            new TeaPair("readTimeout", runtime.readTimeout),
            new TeaPair("connectTimeout", runtime.connectTimeout),
            new TeaPair("maxIdleConns", runtime.maxIdleConns),
            new TeaPair("retry", TeaConverter.buildMap(
                new TeaPair("retryable", runtime.autoretry),
                new TeaPair("maxAttempts", runtime.maxAttempts)
            )),
            new TeaPair("ignoreSSL", runtime.ignoreSSL)
        );

        TeaRequest _lastRequest = null;
        Exception _lastException = null;
        long _now = System.currentTimeMillis();
        int _retryTimes = 0;
        while (Tea.allowRetry((java.util.Map<String, Object>) runtime_.get("retry"), _retryTimes, _now)) {
            if (_retryTimes > 0) {
                int backoffTime = Tea.getBackoffTime(runtime_.get("backoff"), _retryTimes);
                if (backoffTime > 0) {
                    Tea.sleep(backoffTime);
                }
            }
            _retryTimes = _retryTimes + 1;
            try {
                TeaRequest request_ = new TeaRequest();
                String method = com.aliyun.teautil.Common.defaultString(request.method, "POST");
                request_.method = method;
                com.aliyun.teaurl.models.Url url = com.aliyun.teaurl.Client.parseUrl(_endpoint);
                request_.protocol = url.scheme;
                request_.pathname = com.aliyun.teautil.Common.defaultString(request.pathname, url.path.pathname);
                request_.port = com.aliyun.darabonbanumber.Client.parseInt(url.host.port);
                request_.headers = TeaConverter.merge(String.class,
                    TeaConverter.buildMap(
                        new TeaPair("host", url.host.hostname)
                    ),
                    request.headers
                );
                java.util.Map<String, String> params = new java.util.HashMap<>();
                if (!com.aliyun.teautil.Common.isUnset(request.body)) {
                    java.util.Map<String, Object> tmp = com.aliyun.teautil.Common.assertAsMap(request.body);
                    params = com.aliyun.openapiutil.Client.query(tmp);
                }

                params = TeaConverter.merge(String.class,
                    TeaConverter.buildMap(
                        new TeaPair("AppKey", _appKey)
                    ),
                    request.query,
                    params
                );
                params.put("Signature", Client.sign(method, _appSecret, params));
                if (com.aliyun.teautil.Common.equalString(method, "GET")) {
                    request_.query = params;
                } else {
                    String formObj = com.aliyun.teautil.Common.toFormString(params);
                    request_.body = Tea.toReadable(formObj);
                    request_.headers.put("content-type", "application/x-www-form-urlencoded");
                }

                _lastRequest = request_;
                TeaResponse response_ = Tea.doAction(request_, runtime_, interceptorChain);

                if (com.aliyun.teautil.Common.is4xx(response_.statusCode) || com.aliyun.teautil.Common.is5xx(response_.statusCode)) {
                    Object res = com.aliyun.teautil.Common.readAsJSON(response_.body);
                    java.util.Map<String, Object> err = com.aliyun.teautil.Common.assertAsMap(res);
                    err.put("statusCode", response_.statusCode);
                    throw new TeaException(TeaConverter.buildMap(
                        new TeaPair("code", "" + err.get("Code") + ""),
                        new TeaPair("message", "statusCode: " + err.get("statusCode") + ", errorMessage: " + err.get("ErrorMessage") + ", requestId: " + err.get("RequestId") + ""),
                        new TeaPair("data", err)
                    ));
                }

                return com.aliyun.teautil.Common.readAsString(response_.body);
            } catch (Exception e) {
                if (Tea.isRetryable(e)) {
                    _lastException = e;
                    continue;
                }
                throw e;
            }
        }
        throw new TeaUnretryableException(_lastRequest, _lastException);
    }

    public void addRuntimeOptionsInterceptor(RuntimeOptionsInterceptor interceptor) {
        interceptorChain.addRuntimeOptionsInterceptor(interceptor);
    }

    public void addRequestInterceptor(RequestInterceptor interceptor) {
        interceptorChain.addRequestInterceptor(interceptor);
    }

    public void addResponseInterceptor(ResponseInterceptor interceptor) {
        interceptorChain.addResponseInterceptor(interceptor);
    }

    public static String sign(String method, String appSecret, java.util.Map<String, String> params) throws Exception {
        java.util.List<String> keys = com.aliyun.darabonba.map.Client.keySet(params);
        java.util.List<String> sortedKeys = com.aliyun.darabonba.array.Client.ascSort(keys);
        String canonicalizedResource = "";
        String separator = "";
        for (String key : sortedKeys) {
            canonicalizedResource = "" + canonicalizedResource + "" + separator + "" + com.aliyun.darabonba.encode.Encoder.percentEncode(key) + "";
            if (!com.aliyun.teautil.Common.empty(params.get(key))) {
                canonicalizedResource = "" + canonicalizedResource + "=" + com.aliyun.darabonba.encode.Encoder.percentEncode(params.get(key)) + "";
            }

            separator = "&";
        }
        String signToString = "" + method + "&" + com.aliyun.darabonba.encode.Encoder.percentEncode("/") + "&" + com.aliyun.darabonba.encode.Encoder.percentEncode(canonicalizedResource) + "";
        byte[] signData = com.aliyun.darabonba.signature.Signer.HmacSHA1Sign(signToString, "" + appSecret + "&");
        return com.aliyun.darabonba.encode.Encoder.base64EncodeToString(signData);
    }
}
