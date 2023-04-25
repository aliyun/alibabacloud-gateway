# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import unicode_literals

from Tea.exceptions import TeaException
from Tea.core import TeaCore
from Tea.converter import TeaConverter

from alibabacloud_gateway_spi.client import Client as SPIClient
from alibabacloud_tea_util.client import Client as UtilClient


class Client(SPIClient):
    def __init__(self):
        super(Client, self).__init__()

    def modify_configuration(self, context, attribute_map):
        pass

    def modify_request(self, context, attribute_map):
        request = context.request
        config = context.configuration
        request.headers = TeaCore.merge({
            'host': config.endpoint,
            'user-agent': request.user_agent
        }, request.headers)
        if not UtilClient.is_unset(request.body):
            json_obj = UtilClient.to_jsonstring(request.body)
            request.stream = json_obj
            request.headers['content-type'] = 'application/json; charset=utf-8'

    def modify_response(self, context, attribute_map):
        request = context.request
        response = context.response
        if UtilClient.is_4xx(response.status_code) or UtilClient.is_5xx(response.status_code):
            _res = UtilClient.read_as_json(response.body)
            err = UtilClient.assert_as_map(_res)
            err['statusCode'] = response.status_code
            raise TeaException({
                'code': '%s' % TeaConverter.to_unicode(self.default_any(err.get('Code'), err.get('code'))),
                'message': 'code: %s, %s request id: %s' % (TeaConverter.to_unicode(response.status_code), TeaConverter.to_unicode(self.default_any(err.get('Message'), err.get('message'))), TeaConverter.to_unicode(self.default_any(err.get('requestId'), err.get('requestid')))),
                'data': err,
                'description': '%s' % TeaConverter.to_unicode(self.default_any(err.get('Description'), err.get('description'))),
                'accessDeniedDetail': self.default_any(err.get('accessDeniedDetail'), err.get('accessdenieddetail'))
            })
        if UtilClient.equal_number(response.status_code, 204):
            UtilClient.read_as_string(response.body)
        elif UtilClient.equal_string(request.body_type, 'binary'):
            response.deserialized_body = response.body
        elif UtilClient.equal_string(request.body_type, 'byte'):
            byt = UtilClient.read_as_bytes(response.body)
            response.deserialized_body = byt
        elif UtilClient.equal_string(request.body_type, 'string'):
            str = UtilClient.read_as_string(response.body)
            response.deserialized_body = str
        elif UtilClient.equal_string(request.body_type, 'json'):
            obj = UtilClient.read_as_json(response.body)
            res = UtilClient.assert_as_map(obj)
            response.deserialized_body = res
        elif UtilClient.equal_string(request.body_type, 'array'):
            arr = UtilClient.read_as_json(response.body)
            response.deserialized_body = arr
        else:
            response.deserialized_body = UtilClient.read_as_string(response.body)

    def default_any(self, input_value, default_value):
        if UtilClient.is_unset(input_value):
            return default_value
        return input_value
