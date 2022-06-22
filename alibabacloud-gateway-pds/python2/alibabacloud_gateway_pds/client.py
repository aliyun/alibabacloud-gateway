# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import unicode_literals

from Tea.exceptions import TeaException
from alibabacloud_darabonba_signature_util.signer import Signer
from alibabacloud_darabonba_encode_util.encoder import Encoder
from Tea.core import TeaCore
from Tea.converter import TeaConverter

from alibabacloud_gateway_spi.client import Client as SPIClient
from alibabacloud_tea_util.client import Client as UtilClient
from alibabacloud_openapi_util.client import Client as OpenApiUtilClient
from alibabacloud_darabonba_string.client import Client as StringClient
from alibabacloud_darabonba_map.client import Client as MapClient
from alibabacloud_darabonba_array.client import Client as ArrayClient


class Client(SPIClient):
    def __init__(self):
        super(Client, self).__init__()

    def modify_configuration(self, context, attribute_map):
        config = context.configuration
        config.endpoint = self.get_endpoint(config.network, config.endpoint)

    def modify_request(self, context, attribute_map):
        request = context.request
        config = context.configuration
        request.headers = TeaCore.merge({
            'date': UtilClient.get_date_utcstring(),
            'host': config.endpoint,
            'x-acs-version': request.version,
            'x-acs-action': request.action,
            'user-agent': request.user_agent,
            'x-acs-signature-nonce': UtilClient.get_nonce(),
            'x-acs-signature-method': 'HMAC-SHA1',
            'x-acs-signature-version': '1.0',
            'accept': 'application/json'
        }, request.headers)
        if not UtilClient.is_unset(request.stream):
            tmp = UtilClient.read_as_bytes(request.stream)
            request.stream = tmp
            request.headers['content-type'] = 'application/octet-stream'
        else:
            if not UtilClient.is_unset(request.body):
                if UtilClient.equal_string(request.req_body_type, 'json'):
                    json_obj = UtilClient.to_jsonstring(request.body)
                    request.stream = json_obj
                    request.headers['content-type'] = 'application/json; charset=utf-8'
                else:
                    m = UtilClient.assert_as_map(request.body)
                    form_obj = OpenApiUtilClient.to_form(m)
                    request.stream = form_obj
                    request.headers['content-type'] = 'application/x-www-form-urlencoded'
        if not UtilClient.equal_string(request.auth_type, 'Anonymous'):
            credential = request.credential
            access_key_id = credential.get_access_key_id()
            access_key_secret = credential.get_access_key_secret()
            security_token = credential.get_security_token()
            bearer_token = credential.get_bearer_token()
            if not UtilClient.empty(bearer_token):
                request.headers['x-acs-bearer-token'] = bearer_token
                request.headers['Authorization'] = 'Bearer %s' % TeaConverter.to_unicode(bearer_token)
            else:
                if not UtilClient.empty(security_token):
                    request.headers['x-acs-security-token'] = security_token
                request.headers['Authorization'] = self.get_authorization(request.pathname, request.method, request.query, request.headers, access_key_id, access_key_secret)

    def modify_response(self, context, attribute_map):
        request = context.request
        response = context.response
        if UtilClient.is_4xx(response.status_code) or UtilClient.is_5xx(response.status_code):
            _res = UtilClient.read_as_json(response.body)
            err = UtilClient.assert_as_map(_res)
            request_id = self.default_any(err.get('RequestId'), err.get('requestId'))
            err['statusCode'] = response.status_code
            raise TeaException({
                'code': '%s' % TeaConverter.to_unicode(self.default_any(err.get('Code'), err.get('code'))),
                'message': 'code: %s, %s request id: %s' % (TeaConverter.to_unicode(response.status_code), TeaConverter.to_unicode(self.default_any(err.get('Message'), err.get('message'))), TeaConverter.to_unicode(request_id)),
                'data': err
            })
        if UtilClient.equal_string(request.body_type, 'binary'):
            response.deserialized_body = response.body
        elif UtilClient.equal_string(request.body_type, 'byte'):
            byt = UtilClient.read_as_bytes(response.body)
            response.deserialized_body = byt
        elif UtilClient.equal_string(request.body_type, 'string'):
            str = UtilClient.read_as_string(response.body)
            response.deserialized_body = str
        elif UtilClient.equal_string(request.body_type, 'json'):
            response.deserialized_body = UtilClient.read_as_json(response.body)
        else:
            response.deserialized_body = UtilClient.read_as_string(response.body)

    def get_endpoint(self, network, endpoint):
        real_endpoint = 'api.aliyunpds.com'
        if not UtilClient.empty(endpoint):
            real_endpoint = endpoint
        if not UtilClient.empty(network) and StringClient.equals(network, 'vpc'):
            real_endpoint = StringClient.replace(real_endpoint, 'api.aliyunpds.com', 'api-vpc.aliyunpds.com', None)
            real_endpoint = StringClient.replace(real_endpoint, 'admin.aliyunpds.com', 'admin-vpc.aliyunpds.com', None)
        return real_endpoint

    def default_any(self, input_value, default_value):
        if UtilClient.is_unset(input_value):
            return default_value
        return input_value

    def get_authorization(self, pathname, method, query, headers, ak, secret):
        signature = self.get_signature(pathname, method, query, headers, secret)
        return 'acs %s:%s' % (TeaConverter.to_unicode(ak), TeaConverter.to_unicode(signature))

    def get_signature(self, pathname, method, query, headers, secret):
        string_to_sign = ''
        canonicalized_resource = self.build_canonicalized_resource(pathname, query)
        canonicalized_headers = self.build_canonicalized_headers(headers)
        string_to_sign = '%s\n%s%s' % (TeaConverter.to_unicode(method), TeaConverter.to_unicode(canonicalized_headers), TeaConverter.to_unicode(canonicalized_resource))
        signature = Signer.hmac_sha1sign(string_to_sign, secret)
        return Encoder.base_64encode_to_string(signature)

    def build_canonicalized_resource(self, pathname, query):
        canonicalized_resource = pathname
        if not UtilClient.is_unset(query):
            query_array = MapClient.key_set(query)
            sorted_query_array = ArrayClient.asc_sort(query_array)
            separator = '?'
            for key in sorted_query_array:
                canonicalized_resource = '%s%s%s' % (TeaConverter.to_unicode(canonicalized_resource), TeaConverter.to_unicode(separator), TeaConverter.to_unicode(key))
                if not UtilClient.empty(query.get(key)):
                    canonicalized_resource = '%s=%s' % (TeaConverter.to_unicode(canonicalized_resource), TeaConverter.to_unicode(query.get(key)))
                separator = '&'
        return canonicalized_resource

    def build_canonicalized_headers(self, headers):
        accept = headers.get('accept')
        if UtilClient.is_unset(accept):
            accept = ''
        content_md_5 = headers.get('content-md5')
        if UtilClient.is_unset(content_md_5):
            content_md_5 = ''
        content_type = headers.get('content-type')
        if UtilClient.is_unset(content_type):
            content_type = ''
        date = headers.get('date')
        if UtilClient.is_unset(date):
            date = ''
        canonicalized_headers = '%s\n%s\n%s\n%s\n' % (TeaConverter.to_unicode(accept), TeaConverter.to_unicode(content_md_5), TeaConverter.to_unicode(content_type), TeaConverter.to_unicode(date))
        sorted_headers = self.get_signed_headers(headers)
        for header in sorted_headers:
            value = headers.get(header)
            value_trim = StringClient.trim(value)
            canonicalized_headers = '%s%s:%s\n' % (TeaConverter.to_unicode(canonicalized_headers), TeaConverter.to_unicode(header), TeaConverter.to_unicode(value_trim))
        return canonicalized_headers

    def get_signed_headers(self, headers):
        headers_array = MapClient.key_set(headers)
        sorted_headers_array = ArrayClient.asc_sort(headers_array)
        tmp = ''
        separator = ''
        for key in sorted_headers_array:
            lower_key = StringClient.to_lower(key)
            if StringClient.has_prefix(lower_key, 'x-acs-'):
                if not StringClient.contains(tmp, lower_key):
                    tmp = '%s%s%s' % (TeaConverter.to_unicode(tmp), TeaConverter.to_unicode(separator), TeaConverter.to_unicode(lower_key))
                    separator = ';'
        return StringClient.split(tmp, ';', 0)
