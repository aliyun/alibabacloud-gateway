# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import unicode_literals

import time

from Tea.request import TeaRequest
from Tea.exceptions import TeaException, UnretryableException
from Tea.core import TeaCore
from Tea.converter import TeaConverter

from alibabacloud_tea_util import models as util_models
from Url.client import Client as UrlClient
from Number.client import Client as NumberClient
from alibabacloud_tea_util.client import Client as UtilClient


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
                _request.method = 'POST'
                url = UrlClient.parse_url(self._endpoint)
                _request.protocol = url.scheme
                _request.port = NumberClient.parse_int(url.host.port)
                _request.pathname = '%s/%s' % (TeaConverter.to_unicode(UtilClient.default_string(request.pathname, url.path.pathname)), TeaConverter.to_unicode(request.action))
                _request.headers = TeaCore.merge({
                    'host': url.host.hostname,
                    'accept': 'application/json',
                    'content-type': 'application/json; charset=utf-8'
                }, request.headers)
                params = TeaCore.merge({
                    'AppKey': self._app_key,
                    'Action': request.action
                }, request.body)
                _request.body = UtilClient.to_jsonstring(params)
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
