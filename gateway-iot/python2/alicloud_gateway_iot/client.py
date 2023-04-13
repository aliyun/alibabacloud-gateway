# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import unicode_literals

import time

from Tea.request import TeaRequest
from Tea.exceptions import TeaException, UnretryableException
from Tea.core import TeaCore
from alibabacloud_darabonba_encode_util.encoder import Encoder
from alibabacloud_darabonba_signature_util.signer import Signer
from Tea.converter import TeaConverter

from alibabacloud_tea_util import models as util_models
from alibabacloud_tea_util.client import Client as UtilClient
from Url.client import Client as UrlClient
from Number.client import Client as NumberClient
from alibabacloud_openapi_util.client import Client as OpenApiUtilClient
from alibabacloud_darabonba_map.client import Client as MapClient
from alibabacloud_darabonba_array.client import Client as ArrayClient


class Client(object):
    _app_key = None  # type: str
    _app_secret = None  # type: str
    _endpoint = None  # type: str

    def __init__(self, config):
        self._app_key = config.app_key
        self._app_secret = config.app_secret
        self._endpoint = config.endpoint

    def execute(self, request, runtime):
        request.validate()
        runtime.validate()
        _runtime = {
            'timeouted': 'retry',
            'readTimeout': runtime.read_timeout,
            'connectTimeout': runtime.connect_timeout,
            'maxIdleConns': runtime.max_idle_conns,
            'retry': {
                'retryable': runtime.autoretry,
                'maxAttempts': runtime.max_attempts
            },
            'ignoreSSL': runtime.ignore_ssl
        }
        _last_request = None
        _last_exception = None
        _now = time.time()
        _retry_times = 0
        while TeaCore.allow_retry(_runtime.get('retry'), _retry_times, _now):
            if _retry_times > 0:
                _backoff_time = TeaCore.get_backoff_time(_runtime.get('backoff'), _retry_times)
                if _backoff_time > 0:
                    TeaCore.sleep(_backoff_time)
            _retry_times = _retry_times + 1
            try:
                _request = TeaRequest()
                method = UtilClient.default_string(request.method, 'POST')
                _request.method = method
                url = UrlClient.parse_url(self._endpoint)
                _request.protocol = url.scheme
                _request.pathname = UtilClient.default_string(request.pathname, url.path.pathname)
                _request.port = NumberClient.parse_int(url.host.port)
                _request.headers = TeaCore.merge({
                    'host': url.host.hostname
                }, request.headers)
                params = {}
                if not UtilClient.is_unset(request.body):
                    tmp = UtilClient.assert_as_map(request.body)
                    params = OpenApiUtilClient.query(tmp)
                params = TeaCore.merge({
                    'AppKey': self._app_key
                }, request.query,
                    params)
                params['Signature'] = self.sign(method, self._app_secret, params)
                if UtilClient.equal_string(method, 'GET'):
                    _request.query = params
                else:
                    form_obj = UtilClient.to_form_string(params)
                    _request.body = form_obj
                    _request.headers['content-type'] = 'application/x-www-form-urlencoded'
                _last_request = _request
                _response = TeaCore.do_action(_request, _runtime)
                if UtilClient.is_4xx(_response.status_code) or UtilClient.is_5xx(_response.status_code):
                    res = UtilClient.read_as_json(_response.body)
                    err = UtilClient.assert_as_map(res)
                    err['statusCode'] = _response.status_code
                    raise TeaException({
                        'code': '%s' % TeaConverter.to_unicode(err.get('Code')),
                        'message': 'statusCode: %s, errorMessage: %s, requestId: %s' % (TeaConverter.to_unicode(err.get('statusCode')), TeaConverter.to_unicode(err.get('ErrorMessage')), TeaConverter.to_unicode(err.get('RequestId'))),
                        'data': err
                    })
                return UtilClient.read_as_string(_response.body)
            except Exception as e:
                if TeaCore.is_retryable(e):
                    _last_exception = e
                    continue
                raise e
        raise UnretryableException(_last_request, _last_exception)

    @staticmethod
    def sign(method, app_secret, params):
        keys = MapClient.key_set(params)
        sorted_keys = ArrayClient.asc_sort(keys)
        canonicalized_resource = ''
        separator = ''
        for key in sorted_keys:
            canonicalized_resource = '%s%s%s' % (TeaConverter.to_unicode(canonicalized_resource), TeaConverter.to_unicode(separator), TeaConverter.to_unicode(Encoder.percent_encode(key)))
            if not UtilClient.empty(params.get(key)):
                canonicalized_resource = '%s=%s' % (TeaConverter.to_unicode(canonicalized_resource), TeaConverter.to_unicode(Encoder.percent_encode(params.get(key))))
            separator = '&'
        sign_to_string = '%s&%s&%s' % (TeaConverter.to_unicode(method), TeaConverter.to_unicode(Encoder.percent_encode('/')), TeaConverter.to_unicode(Encoder.percent_encode(canonicalized_resource)))
        sign_data = Signer.hmac_sha1sign(sign_to_string, '%s&' % TeaConverter.to_unicode(app_secret))
        return Encoder.base_64encode_to_string(sign_data)
