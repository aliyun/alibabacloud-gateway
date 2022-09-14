# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import unicode_literals

from alibabacloud_darabonba_encode_util.encoder import Encoder
from alibabacloud_darabonba_signature_util.signer import Signer
from Tea.exceptions import TeaException
from Tea.core import TeaCore
from Tea.converter import TeaConverter

from alibabacloud_gateway_spi.client import Client as SPIClient
from alibabacloud_tea_util.client import Client as UtilClient
from alibabacloud_darabonba_string.client import Client as StringClient
from alibabacloud_darabonba_array.client import Client as ArrayClient
from alibabacloud_darabonba_map.client import Client as MapClient


class Client(SPIClient):
    def __init__(self):
        super(Client, self).__init__()

    def modify_configuration(self, context, attribute_map):
        config = context.configuration
        config.endpoint = self.get_endpoint(config.region_id, config.network, config.endpoint)

    def modify_request(self, context, attribute_map):
        request = context.request
        host_map = {}
        if not UtilClient.is_unset(request.host_map):
            host_map = request.host_map
        project = host_map.get('project')
        config = context.configuration
        credential = request.credential
        access_key_id = credential.get_access_key_id()
        access_key_secret = credential.get_access_key_secret()
        security_token = credential.get_security_token()
        if not UtilClient.empty(access_key_id):
            request.headers['x-log-signaturemethod'] = 'hmac-sha1'
        if not UtilClient.empty(security_token):
            request.headers['x-acs-security-token'] = security_token
        if not UtilClient.is_unset(request.body):
            if StringClient.equals(request.req_body_type, 'protobuf'):
                # var bodyMap = Util.assertAsMap(request.body);
                # 缺少body的Content-MD5计算，以及protobuf处理
                request.headers['content-type'] = 'application/x-protobuf'
            elif StringClient.equals(request.req_body_type, 'json'):
                body_str = UtilClient.to_jsonstring(request.body)
                request.headers['content-md5'] = StringClient.to_upper(Encoder.hex_encode(Signer.md5sign(body_str)))
                request.stream = body_str
                request.headers['content-type'] = 'application/json'
            elif StringClient.equals(request.req_body_type, 'formData'):
                str = UtilClient.to_jsonstring(request.body)
                request.headers['content-md5'] = StringClient.to_upper(Encoder.hex_encode(Signer.md5sign(str)))
                request.stream = str
                request.headers['content-type'] = 'application/json'
        host = self.get_host(config.network, project, config.endpoint)
        request.headers = TeaCore.merge({
            'accept': 'application/json',
            'host': host,
            'date': UtilClient.get_date_utcstring(),
            'user-agent': request.user_agent,
            'x-log-apiversion': '0.6.0',
            'x-log-bodyrawsize': '0'
        }, request.headers)
        request.headers['authorization'] = self.get_authorization(request.pathname, request.method, request.query, request.headers, access_key_id, access_key_secret)

    def modify_response(self, context, attribute_map):
        request = context.request
        response = context.response
        if UtilClient.is_4xx(response.status_code) or UtilClient.is_5xx(response.status_code):
            error = UtilClient.read_as_json(response.body)
            res_map = UtilClient.assert_as_map(error)
            raise TeaException({
                'code': res_map.get('errorCode'),
                'message': res_map.get('errorMessage'),
                'data': {
                    'httpCode': response.status_code,
                    'requestId': response.headers.get('x-log-requestid'),
                    'statusCode': response.status_code
                }
            })
        if not UtilClient.is_unset(response.body):
            if UtilClient.equal_string(request.body_type, 'binary'):
                response.deserialized_body = response.body
            elif UtilClient.equal_string(request.body_type, 'byte'):
                byt = UtilClient.read_as_bytes(response.body)
                response.deserialized_body = byt
            elif UtilClient.equal_string(request.body_type, 'string'):
                response.deserialized_body = UtilClient.read_as_string(response.body)
            elif UtilClient.equal_string(request.body_type, 'json'):
                obj = UtilClient.read_as_json(response.body)
                # var res = Util.assertAsMap(obj);
                response.deserialized_body = obj
            elif UtilClient.equal_string(request.body_type, 'array'):
                response.deserialized_body = UtilClient.read_as_json(response.body)
            else:
                response.deserialized_body = UtilClient.read_as_string(response.body)

    def get_endpoint(self, region_id, network, endpoint):
        if not UtilClient.empty(endpoint):
            return endpoint
        if UtilClient.empty(region_id):
            region_id = 'cn-hangzhou'
        if not UtilClient.empty(network):
            if StringClient.equals(network, 'intranet'):
                return '%s-intranet.log.aliyuncs.com' % TeaConverter.to_unicode(region_id)
            elif StringClient.equals(network, 'accelerate'):
                return 'log-global.aliyuncs.com'
            elif StringClient.equals(network, 'share'):
                if StringClient.equals(region_id, 'cn-hangzhou-corp') or StringClient.equals(region_id, 'cn-shanghai-corp'):
                    return '%s.sls.aliyuncs.com' % TeaConverter.to_unicode(region_id)
                elif StringClient.equals(region_id, 'cn-zhangbei-corp'):
                    return 'zhangbei-corp-share.log.aliyuncs.com'
                return '%s-share.log.aliyuncs.com' % TeaConverter.to_unicode(region_id)
        return '%s.log.aliyuncs.com' % TeaConverter.to_unicode(region_id)

    def get_host(self, network, project, endpoint):
        if UtilClient.is_unset(project):
            return endpoint
        return '%s.%s' % (TeaConverter.to_unicode(project), TeaConverter.to_unicode(endpoint))

    def get_authorization(self, pathname, method, query, headers, ak, secret):
        sign = self.get_signature(pathname, method, query, headers, secret)
        return 'LOG %s:%s' % (TeaConverter.to_unicode(ak), TeaConverter.to_unicode(sign))

    def get_signature(self, pathname, method, query, headers, secret):
        resource = pathname
        string_to_sign = ''
        canonicalized_resource = self.build_canonicalized_resource(resource, query)
        canonicalized_headers = self.build_canonicalized_headers(headers)
        string_to_sign = '%s\n%s%s' % (TeaConverter.to_unicode(method), TeaConverter.to_unicode(canonicalized_headers), TeaConverter.to_unicode(canonicalized_resource))
        return Encoder.base_64encode_to_string(Signer.hmac_sha1sign(string_to_sign, secret))

    def build_canonicalized_resource(self, pathname, query):
        canonicalized_resource = pathname
        params_map = TeaCore.merge(query)
        if not UtilClient.empty(pathname):
            paths = StringClient.split(pathname, '?', 2)
            canonicalized_resource = paths[0]
            if UtilClient.equal_number(ArrayClient.size(paths), 2):
                params = StringClient.split(paths[1], '&', 0)
                for sub in params:
                    item = StringClient.split(sub, '=', 0)
                    key = item[0]
                    value = None
                    if UtilClient.equal_number(ArrayClient.size(item), 2):
                        value = item[1]
                    params_map[key] = value
        if not UtilClient.is_unset(params_map):
            query_list = MapClient.key_set(params_map)
            sorted_params = ArrayClient.asc_sort(query_list)
            separator = '?'
            for param_name in sorted_params:
                canonicalized_resource = '%s%s%s' % (TeaConverter.to_unicode(canonicalized_resource), TeaConverter.to_unicode(separator), TeaConverter.to_unicode(param_name))
                param_value = params_map.get(param_name)
                if not UtilClient.is_unset(param_value):
                    canonicalized_resource = '%s=%s' % (TeaConverter.to_unicode(canonicalized_resource), TeaConverter.to_unicode(param_value))
                separator = '&'
        return canonicalized_resource

    def build_canonicalized_headers(self, headers):
        canonicalized_headers = ''
        content_type = headers.get('content-type')
        if UtilClient.is_unset(content_type):
            content_type = ''
        content_md_5 = headers.get('content-md5')
        if UtilClient.is_unset(content_md_5):
            content_md_5 = ''
        canonicalized_headers = '%s%s\n%s\n%s\n' % (TeaConverter.to_unicode(canonicalized_headers), TeaConverter.to_unicode(content_md_5), TeaConverter.to_unicode(content_type), TeaConverter.to_unicode(headers.get('date')))
        keys = MapClient.key_set(headers)
        sorted_headers = ArrayClient.asc_sort(keys)
        for header in sorted_headers:
            if StringClient.contains(StringClient.to_lower(header), 'x-log-') or StringClient.contains(StringClient.to_lower(header), 'x-acs-'):
                canonicalized_headers = '%s%s:%s\n' % (TeaConverter.to_unicode(canonicalized_headers), TeaConverter.to_unicode(header), TeaConverter.to_unicode(headers.get(header)))
        return canonicalized_headers
