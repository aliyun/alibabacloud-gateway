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
from alibabacloud_gateway_sls_util.client import Client as SLS_UtilClient
from alibabacloud_darabonba_map.client import Client as MapClient
from alibabacloud_darabonba_array.client import Client as ArrayClient
from alibabacloud_openapi_util.client import Client as OpenApiUtilClient


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
        if not UtilClient.empty(security_token):
            request.headers['x-acs-security-token'] = security_token
        signature_version = UtilClient.default_string(request.signature_version, 'v1')
        content_hash = ''
        if not UtilClient.is_unset(request.body):
            if StringClient.equals(request.req_body_type, 'protobuf'):
                # var bodyMap = Util.assertAsMap(request.body);
                # 缺少body的Content-MD5计算，以及protobuf处理
                request.headers['content-type'] = 'application/x-protobuf'
            elif StringClient.equals(request.req_body_type, 'json'):
                body_str = UtilClient.to_jsonstring(request.body)
                content_hash = self.make_content_hash(body_str, signature_version)
                request.stream = body_str
                request.headers['content-type'] = 'application/json'
            elif StringClient.equals(request.req_body_type, 'formData'):
                str = UtilClient.to_jsonstring(request.body)
                content_hash = self.make_content_hash(str, signature_version)
                request.stream = str
                request.headers['content-type'] = 'application/json'
        host = self.get_host(config.network, project, config.endpoint)
        request.headers = TeaCore.merge({
            'accept': 'application/json',
            'host': host,
            'user-agent': request.user_agent,
            'x-log-apiversion': '0.6.0',
            'x-log-bodyrawsize': '0'
        }, request.headers)
        self.build_request(context)
        # move param in path to query
        if StringClient.equals(signature_version, 'v4'):
            if UtilClient.empty(content_hash):
                content_hash = 'e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855'
            date = self.get_date_iso8601()
            request.headers['x-log-date'] = date
            request.headers['x-log-content-sha256'] = content_hash
            request.headers['authorization'] = self.get_authorization_v4(context, date, content_hash, access_key_id, access_key_secret)
            return
        if not UtilClient.empty(access_key_id):
            request.headers['x-log-signaturemethod'] = 'hmac-sha256'
        request.headers['date'] = UtilClient.get_date_utcstring()
        request.headers['content-md5'] = content_hash
        request.headers['authorization'] = self.get_authorization(request.pathname, request.method, request.query, request.headers, access_key_id, access_key_secret)

    def make_content_hash(self, content, signature_version):
        if StringClient.equals(signature_version, 'v4'):
            if UtilClient.empty(content):
                return ''
            return StringClient.to_lower(Encoder.hex_encode(Encoder.hash(UtilClient.to_bytes(content), 'SLS4-HMAC-SHA256')))
        return StringClient.to_upper(Encoder.hex_encode(Signer.md5sign(content)))

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
            bodyraw_size = response.headers.get('x-log-bodyrawsize')
            compress_type = response.headers.get('x-log-compresstype')
            uncompressed_data = response.body
            if not UtilClient.is_unset(bodyraw_size) and not UtilClient.is_unset(compress_type):
                uncompressed_data = SLS_UtilClient.read_and_uncompress_block(response.body, compress_type, bodyraw_size)
            if UtilClient.equal_string(request.body_type, 'binary'):
                response.deserialized_body = uncompressed_data
            elif UtilClient.equal_string(request.body_type, 'byte'):
                byt = UtilClient.read_as_bytes(uncompressed_data)
                response.deserialized_body = byt
            elif UtilClient.equal_string(request.body_type, 'string'):
                response.deserialized_body = UtilClient.read_as_string(uncompressed_data)
            elif UtilClient.equal_string(request.body_type, 'json'):
                obj = UtilClient.read_as_json(uncompressed_data)
                # var res = Util.assertAsMap(obj);
                response.deserialized_body = obj
            elif UtilClient.equal_string(request.body_type, 'array'):
                response.deserialized_body = UtilClient.read_as_json(uncompressed_data)
            else:
                response.deserialized_body = UtilClient.read_as_string(uncompressed_data)

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
        return Encoder.base_64encode_to_string(Signer.hmac_sha256sign(string_to_sign, secret))

    def build_canonicalized_resource(self, pathname, query):
        canonicalized_resource = pathname
        if not UtilClient.is_unset(query):
            query_list = MapClient.key_set(query)
            sorted_params = ArrayClient.asc_sort(query_list)
            separator = '?'
            for param_name in sorted_params:
                canonicalized_resource = '%s%s%s' % (TeaConverter.to_unicode(canonicalized_resource), TeaConverter.to_unicode(separator), TeaConverter.to_unicode(param_name))
                param_value = query.get(param_name)
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

    def build_request(self, context):
        request = context.request
        resource = request.pathname
        if not UtilClient.empty(resource):
            paths = StringClient.split(resource, '?', 2)
            resource = paths[0]
            if UtilClient.equal_number(ArrayClient.size(paths), 2):
                params = StringClient.split(paths[1], '&', None)
                for sub in params:
                    item = StringClient.split(sub, '=', None)
                    key = item[0]
                    value = None
                    if UtilClient.equal_number(ArrayClient.size(item), 2):
                        value = item[1]
                    request.query[key] = value
        request.pathname = resource

    def get_authorization_v4(self, context, date, content_hash, access_key_id, access_key_secret):
        region = self.get_region(context)
        header_str = self.get_signed_header_str_v4(context.request.headers)
        date_short = StringClient.sub_string(date, 0, 8)
        date_short = StringClient.replace(date_short, 'T', '', None)
        # for fix php sdk bug
        scope = '%s/%s/sls/aliyun_v4_request' % (TeaConverter.to_unicode(date_short), TeaConverter.to_unicode(region))
        signingkey = self.get_signingkey_v4('SLS4-HMAC-SHA256', access_key_secret, region, date_short)
        signature = self.get_signature_v4(context, 'SLS4-HMAC-SHA256', header_str, date, scope, content_hash, signingkey)
        return '%s Credential=%s/%s,Signature=%s' % (TeaConverter.to_unicode('SLS4-HMAC-SHA256'), TeaConverter.to_unicode(access_key_id), TeaConverter.to_unicode(scope), TeaConverter.to_unicode(signature))

    def get_signingkey_v4(self, signature_algorithm, secret, region, date):
        sc_1 = 'aliyun_v4%s' % TeaConverter.to_unicode(secret)
        sc_2 = Signer.hmac_sha256sign(date, sc_1)
        sc_3 = Signer.hmac_sha256sign_by_bytes(region, sc_2)
        sc_4 = Signer.hmac_sha256sign_by_bytes('sls', sc_3)
        return Signer.hmac_sha256sign_by_bytes('aliyun_v4_request', sc_4)

    def get_signature_v4(self, context, signature_algorithm, signed_header_str, date, scope, content_sha_256, signingkey):
        request = context.request
        canonical_uri = '/'
        if not UtilClient.empty(request.pathname):
            canonical_uri = request.pathname
        resources = self.build_canonicalized_resource_v4(request.query)
        headers = self.build_canonicalized_headers_v4(request.headers, signed_header_str)
        string_to_hash = '%s\n%s\n%s\n%s\n%s\n%s' % (TeaConverter.to_unicode(request.method), TeaConverter.to_unicode(canonical_uri), TeaConverter.to_unicode(resources), TeaConverter.to_unicode(headers), TeaConverter.to_unicode(signed_header_str), TeaConverter.to_unicode(content_sha_256))
        hex = Encoder.hex_encode(Encoder.hash(UtilClient.to_bytes(string_to_hash), signature_algorithm))
        string_to_sign = '%s\n%s\n%s\n%s' % (TeaConverter.to_unicode(signature_algorithm), TeaConverter.to_unicode(date), TeaConverter.to_unicode(scope), TeaConverter.to_unicode(hex))
        signature = Signer.hmac_sha256sign_by_bytes(string_to_sign, signingkey)
        return Encoder.hex_encode(signature)

    def build_canonicalized_resource_v4(self, query):
        canonicalized_resource = ''
        if not UtilClient.is_unset(query):
            query_array = MapClient.key_set(query)
            sorted_query_array = ArrayClient.asc_sort(query_array)
            separator = ''
            for key in sorted_query_array:
                canonicalized_resource = '%s%s%s' % (TeaConverter.to_unicode(canonicalized_resource), TeaConverter.to_unicode(separator), TeaConverter.to_unicode(key))
                if not UtilClient.empty(query.get(key)):
                    canonicalized_resource = '%s=%s' % (TeaConverter.to_unicode(canonicalized_resource), TeaConverter.to_unicode(Encoder.percent_encode(query.get(key))))
                separator = '&'
        return canonicalized_resource

    def build_canonicalized_headers_v4(self, headers, signed_header_str):
        canonicalized_headers = ''
        signed_headers = StringClient.split(signed_header_str, ';', None)
        for header in signed_headers:
            canonicalized_headers = '%s%s:%s\n' % (TeaConverter.to_unicode(canonicalized_headers), TeaConverter.to_unicode(header), TeaConverter.to_unicode(StringClient.trim(headers.get(header))))
        return canonicalized_headers

    def get_signed_header_str_v4(self, headers):
        headers_array = MapClient.key_set(headers)
        sorted_headers_array = ArrayClient.asc_sort(headers_array)
        tmp = ''
        separator = ''
        for key in sorted_headers_array:
            lower_key = StringClient.to_lower(key)
            if StringClient.has_prefix(lower_key, 'x-log-') or StringClient.equals(lower_key, 'host') or StringClient.equals(lower_key, 'content-type'):
                if not StringClient.contains(tmp, lower_key):
                    tmp = '%s%s%s' % (TeaConverter.to_unicode(tmp), TeaConverter.to_unicode(separator), TeaConverter.to_unicode(lower_key))
                    separator = ';'
        return tmp

    def get_region(self, context):
        config = context.configuration
        if not UtilClient.is_unset(config.region_id) and not UtilClient.empty(config.region_id):
            return config.region_id
        # try parse region from endpoint
        # do not use String.subString, subString has bug in php implementation
        region = StringClient.replace(config.endpoint, '.log.aliyuncs.com', '', None)
        region = StringClient.replace(region, '.sls.aliyuncs.com', '', None)
        if StringClient.equals(region, config.endpoint):
            raise TeaException({
                'code': 'ClientConfigError',
                'message': 'The regionId configuration of sls client is missing.'
            })
        region = StringClient.replace(region, '-share', '', None)
        region = StringClient.replace(region, '-intranet', '', None)
        region = StringClient.replace(region, '-vpc', '', None)
        return region

    def get_date_iso8601(self):
        """
        format: YYYYMMDDTHHMMSSZ
        """
        date = OpenApiUtilClient.get_timestamp()
        # 2024-02-04T11:31:58Z
        date = StringClient.replace(date, '-', '', None)
        return StringClient.replace(date, ':', '', None)
