# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from Tea.exceptions import TeaException
from alibabacloud_darabonba_encode_util.encoder import Encoder
from alibabacloud_darabonba_signature_util.signer import Signer
from typing import List, Dict
from Tea.core import TeaCore

from alibabacloud_gateway_spi.client import Client as SPIClient
from alibabacloud_gateway_spi import models as spi_models
from alibabacloud_tea_util.client import Client as UtilClient
from alibabacloud_darabonba_string.client import Client as StringClient
from alibabacloud_tea_xml.client import Client as XMLClient
from alibabacloud_openapi_util.client import Client as OpenApiUtilClient
from alibabacloud_oss_util.client import Client as OSSUtilClient
from alibabacloud_darabonba_map.client import Client as MapClient
from alibabacloud_darabonba_array.client import Client as ArrayClient


class Client(SPIClient):
    _default_signed_params: List[str] = None
    _except_signed_params: List[str] = None

    def __init__(self):
        super().__init__()
        self._default_signed_params = [
            'location',
            'cors',
            'objectMeta',
            'uploadId',
            'partNumber',
            'security-token',
            'position',
            'img',
            'style',
            'styleName',
            'replication',
            'replicationProgress',
            'replicationLocation',
            'cname',
            'qos',
            'startTime',
            'endTime',
            'symlink',
            'x-oss-process',
            'response-content-type',
            'response-content-language',
            'response-expires',
            'response-cache-control',
            'response-content-disposition',
            'response-content-encoding',
            'udf',
            'udfName',
            'udfImage',
            'udfId',
            'udfImageDesc',
            'udfApplication',
            'udfApplicationLog',
            'restore',
            'callback',
            'callback-var',
            'policy',
            'encryption',
            'versions',
            'versioning',
            'versionId'
        ]
        self._except_signed_params = [
            'list-type',
            'regions'
        ]

    def modify_configuration(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        config = context.configuration
        config.endpoint = self.get_endpoint(config.region_id, config.network, config.endpoint)

    async def modify_configuration_async(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        config = context.configuration
        config.endpoint = await self.get_endpoint_async(config.region_id, config.network, config.endpoint)

    def modify_request(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        request = context.request
        host_map = {}
        if not UtilClient.is_unset(request.host_map):
            host_map = request.host_map
        bucket_name = host_map.get('bucket')
        if UtilClient.is_unset(bucket_name):
            bucket_name = ''
        config = context.configuration
        credential = request.credential
        access_key_id = credential.get_access_key_id()
        access_key_secret = credential.get_access_key_secret()
        security_token = credential.get_security_token()
        if not UtilClient.empty(security_token):
            request.headers['x-oss-security-token'] = security_token
        if not UtilClient.is_unset(request.body):
            if StringClient.equals(request.req_body_type, 'xml'):
                req_body_map = UtilClient.assert_as_map(request.body)
                request.stream = XMLClient.to_xml(req_body_map)
                request.headers['content-type'] = 'application/xml'
            elif StringClient.equals(request.req_body_type, 'json'):
                req_body_str = UtilClient.to_jsonstring(request.body)
                request.stream = req_body_str
                request.headers['content-type'] = 'application/json; charset=utf-8'
            elif StringClient.equals(request.req_body_type, 'formData'):
                req_body_form = UtilClient.assert_as_map(request.body)
                request.stream = OpenApiUtilClient.to_form(req_body_form)
                request.headers['content-type'] = 'application/x-www-form-urlencoded'
            elif StringClient.equals(request.req_body_type, 'binary'):
                attribute_map.key = {
                    'crc': '',
                    'md5': ''
                }
                request.stream = OSSUtilClient.inject(request.stream, attribute_map.key)
                request.headers['content-type'] = 'application/octet-stream'
        host = self.get_host(config.endpoint_type, bucket_name, config.endpoint)
        request.headers = TeaCore.merge({
            'host': host,
            'date': UtilClient.get_date_utcstring(),
            'user-agent': request.user_agent
        }, request.headers)
        request.headers['authorization'] = self.get_authorization(request.signature_version, bucket_name, request.pathname, request.method, request.query, request.headers, access_key_id, access_key_secret)

    async def modify_request_async(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        request = context.request
        host_map = {}
        if not UtilClient.is_unset(request.host_map):
            host_map = request.host_map
        bucket_name = host_map.get('bucket')
        if UtilClient.is_unset(bucket_name):
            bucket_name = ''
        config = context.configuration
        credential = request.credential
        access_key_id = await credential.get_access_key_id_async()
        access_key_secret = await credential.get_access_key_secret_async()
        security_token = await credential.get_security_token_async()
        if not UtilClient.empty(security_token):
            request.headers['x-oss-security-token'] = security_token
        if not UtilClient.is_unset(request.body):
            if StringClient.equals(request.req_body_type, 'xml'):
                req_body_map = UtilClient.assert_as_map(request.body)
                request.stream = XMLClient.to_xml(req_body_map)
                request.headers['content-type'] = 'application/xml'
            elif StringClient.equals(request.req_body_type, 'json'):
                req_body_str = UtilClient.to_jsonstring(request.body)
                request.stream = req_body_str
                request.headers['content-type'] = 'application/json; charset=utf-8'
            elif StringClient.equals(request.req_body_type, 'formData'):
                req_body_form = UtilClient.assert_as_map(request.body)
                request.stream = OpenApiUtilClient.to_form(req_body_form)
                request.headers['content-type'] = 'application/x-www-form-urlencoded'
            elif StringClient.equals(request.req_body_type, 'binary'):
                attribute_map.key = {
                    'crc': '',
                    'md5': ''
                }
                request.stream = OSSUtilClient.inject(request.stream, attribute_map.key)
                request.headers['content-type'] = 'application/octet-stream'
        host = await self.get_host_async(config.endpoint_type, bucket_name, config.endpoint)
        request.headers = TeaCore.merge({
            'host': host,
            'date': UtilClient.get_date_utcstring(),
            'user-agent': request.user_agent
        }, request.headers)
        request.headers['authorization'] = await self.get_authorization_async(request.signature_version, bucket_name, request.pathname, request.method, request.query, request.headers, access_key_id, access_key_secret)

    def modify_response(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        request = context.request
        response = context.response
        body_str = None
        if UtilClient.is_4xx(response.status_code) or UtilClient.is_5xx(response.status_code):
            body_str = UtilClient.read_as_string(response.body)
            resp_map = XMLClient.parse_xml(body_str, None)
            err = UtilClient.assert_as_map(resp_map.get('Error'))
            raise TeaException({
                'code': err.get('Code'),
                'message': err.get('Message'),
                'data': {
                    'statusCode': response.status_code,
                    'requestId': err.get('RequestId'),
                    'hostId': err.get('HostId')
                }
            })
        ctx = attribute_map.key
        if not UtilClient.is_unset(ctx):
            if not UtilClient.is_unset(ctx.get('crc')) and not UtilClient.is_unset(response.headers.get('x-oss-hash-crc64ecma')) and not StringClient.equals(ctx.get('crc'), response.headers.get('x-oss-hash-crc64ecma')):
                raise TeaException({
                    'code': 'CrcNotMatched',
                    'data': {
                        'clientCrc': ctx.get('crc'),
                        'serverCrc': response.headers.get('x-oss-hash-crc64ecma')
                    }
                })
            if not UtilClient.is_unset(ctx.get('md5')) and not UtilClient.is_unset(response.headers.get('content-md5')) and not StringClient.equals(ctx.get('md5'), response.headers.get('content-md5')):
                raise TeaException({
                    'code': 'MD5NotMatched',
                    'data': {
                        'clientMD5': ctx.get('md5'),
                        'serverMD5': response.headers.get('content-md5')
                    }
                })
        if not UtilClient.is_unset(response.body):
            if StringClient.equals(request.body_type, 'xml'):
                body_str = UtilClient.read_as_string(response.body)
                result = XMLClient.parse_xml(body_str, None)
                list = MapClient.key_set(result)
                if UtilClient.equal_number(ArrayClient.size(list), 1):
                    tmp = list[0]
                    try:
                        response.deserialized_body = UtilClient.assert_as_map(result.get(tmp))
                    except Exception as error:
                        response.deserialized_body = result
                else:
                    response.deserialized_body = result
            elif UtilClient.equal_string(request.body_type, 'binary'):
                response.deserialized_body = response.body
            elif UtilClient.equal_string(request.body_type, 'byte'):
                byt = UtilClient.read_as_bytes(response.body)
                response.deserialized_body = byt
            elif UtilClient.equal_string(request.body_type, 'string'):
                response.deserialized_body = UtilClient.read_as_string(response.body)
            elif UtilClient.equal_string(request.body_type, 'json'):
                obj = UtilClient.read_as_json(response.body)
                res = UtilClient.assert_as_map(obj)
                response.deserialized_body = res
            elif UtilClient.equal_string(request.body_type, 'array'):
                response.deserialized_body = UtilClient.read_as_json(response.body)
            else:
                response.deserialized_body = UtilClient.read_as_string(response.body)

    async def modify_response_async(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        request = context.request
        response = context.response
        body_str = None
        if UtilClient.is_4xx(response.status_code) or UtilClient.is_5xx(response.status_code):
            body_str = await UtilClient.read_as_string_async(response.body)
            resp_map = XMLClient.parse_xml(body_str, None)
            err = UtilClient.assert_as_map(resp_map.get('Error'))
            raise TeaException({
                'code': err.get('Code'),
                'message': err.get('Message'),
                'data': {
                    'statusCode': response.status_code,
                    'requestId': err.get('RequestId'),
                    'hostId': err.get('HostId')
                }
            })
        ctx = attribute_map.key
        if not UtilClient.is_unset(ctx):
            if not UtilClient.is_unset(ctx.get('crc')) and not UtilClient.is_unset(response.headers.get('x-oss-hash-crc64ecma')) and not StringClient.equals(ctx.get('crc'), response.headers.get('x-oss-hash-crc64ecma')):
                raise TeaException({
                    'code': 'CrcNotMatched',
                    'data': {
                        'clientCrc': ctx.get('crc'),
                        'serverCrc': response.headers.get('x-oss-hash-crc64ecma')
                    }
                })
            if not UtilClient.is_unset(ctx.get('md5')) and not UtilClient.is_unset(response.headers.get('content-md5')) and not StringClient.equals(ctx.get('md5'), response.headers.get('content-md5')):
                raise TeaException({
                    'code': 'MD5NotMatched',
                    'data': {
                        'clientMD5': ctx.get('md5'),
                        'serverMD5': response.headers.get('content-md5')
                    }
                })
        if not UtilClient.is_unset(response.body):
            if StringClient.equals(request.body_type, 'xml'):
                body_str = await UtilClient.read_as_string_async(response.body)
                result = XMLClient.parse_xml(body_str, None)
                list = MapClient.key_set(result)
                if UtilClient.equal_number(ArrayClient.size(list), 1):
                    tmp = list[0]
                    try:
                        response.deserialized_body = UtilClient.assert_as_map(result.get(tmp))
                    except Exception as error:
                        response.deserialized_body = result
                else:
                    response.deserialized_body = result
            elif UtilClient.equal_string(request.body_type, 'binary'):
                response.deserialized_body = response.body
            elif UtilClient.equal_string(request.body_type, 'byte'):
                byt = await UtilClient.read_as_bytes_async(response.body)
                response.deserialized_body = byt
            elif UtilClient.equal_string(request.body_type, 'string'):
                response.deserialized_body = await UtilClient.read_as_string_async(response.body)
            elif UtilClient.equal_string(request.body_type, 'json'):
                obj = await UtilClient.read_as_json_async(response.body)
                res = UtilClient.assert_as_map(obj)
                response.deserialized_body = res
            elif UtilClient.equal_string(request.body_type, 'array'):
                response.deserialized_body = await UtilClient.read_as_json_async(response.body)
            else:
                response.deserialized_body = await UtilClient.read_as_string_async(response.body)

    def get_endpoint(
        self,
        region_id: str,
        network: str,
        endpoint: str,
    ) -> str:
        if not UtilClient.empty(endpoint):
            return endpoint
        if UtilClient.empty(region_id):
            region_id = 'cn-hangzhou'
        if not UtilClient.empty(network):
            if StringClient.contains(network, 'internal'):
                return f'oss-{region_id}-internal.aliyuncs.com'
            elif StringClient.contains(network, 'ipv6'):
                return f'{region_id}oss.aliyuncs.com'
            elif StringClient.contains(network, 'accelerate'):
                return f'oss-{network}.aliyuncs.com'
        return f'oss-{region_id}.aliyuncs.com'

    async def get_endpoint_async(
        self,
        region_id: str,
        network: str,
        endpoint: str,
    ) -> str:
        if not UtilClient.empty(endpoint):
            return endpoint
        if UtilClient.empty(region_id):
            region_id = 'cn-hangzhou'
        if not UtilClient.empty(network):
            if StringClient.contains(network, 'internal'):
                return f'oss-{region_id}-internal.aliyuncs.com'
            elif StringClient.contains(network, 'ipv6'):
                return f'{region_id}oss.aliyuncs.com'
            elif StringClient.contains(network, 'accelerate'):
                return f'oss-{network}.aliyuncs.com'
        return f'oss-{region_id}.aliyuncs.com'

    def get_host(
        self,
        endpoint_type: str,
        bucket_name: str,
        endpoint: str,
    ) -> str:
        if UtilClient.empty(bucket_name):
            return endpoint
        host = f'{bucket_name}.{endpoint}'
        if not UtilClient.empty(endpoint_type):
            if StringClient.equals(endpoint_type, 'ip'):
                host = f'{endpoint}/{bucket_name}'
            elif StringClient.equals(endpoint_type, 'cname'):
                host = endpoint
        return host

    async def get_host_async(
        self,
        endpoint_type: str,
        bucket_name: str,
        endpoint: str,
    ) -> str:
        if UtilClient.empty(bucket_name):
            return endpoint
        host = f'{bucket_name}.{endpoint}'
        if not UtilClient.empty(endpoint_type):
            if StringClient.equals(endpoint_type, 'ip'):
                host = f'{endpoint}/{bucket_name}'
            elif StringClient.equals(endpoint_type, 'cname'):
                host = endpoint
        return host

    def get_authorization(
        self,
        signature_version: str,
        bucket_name: str,
        pathname: str,
        method: str,
        query: Dict[str, str],
        headers: Dict[str, str],
        ak: str,
        secret: str,
    ) -> str:
        sign = ''
        if UtilClient.is_unset(signature_version) or StringClient.equals(signature_version, 'v1'):
            sign = self.get_signature_v1(bucket_name, pathname, method, query, headers, secret)
            return f'OSS {ak}:{sign}'
        else:
            sign = self.get_signature_v2(bucket_name, pathname, method, query, headers, secret)
            return f'OSS2 AccessKeyId:{ak},Signature:{sign}'

    async def get_authorization_async(
        self,
        signature_version: str,
        bucket_name: str,
        pathname: str,
        method: str,
        query: Dict[str, str],
        headers: Dict[str, str],
        ak: str,
        secret: str,
    ) -> str:
        sign = ''
        if UtilClient.is_unset(signature_version) or StringClient.equals(signature_version, 'v1'):
            sign = await self.get_signature_v1_async(bucket_name, pathname, method, query, headers, secret)
            return f'OSS {ak}:{sign}'
        else:
            sign = await self.get_signature_v2_async(bucket_name, pathname, method, query, headers, secret)
            return f'OSS2 AccessKeyId:{ak},Signature:{sign}'

    def get_signature_v1(
        self,
        bucket_name: str,
        pathname: str,
        method: str,
        query: Dict[str, str],
        headers: Dict[str, str],
        secret: str,
    ) -> str:
        resource = ''
        string_to_sign = ''
        if not UtilClient.empty(bucket_name):
            resource = f'/{bucket_name}'
        resource = f'{resource}{pathname}'
        canonicalized_resource = self.build_canonicalized_resource(resource, query)
        canonicalized_headers = self.build_canonicalized_headers(headers)
        string_to_sign = f'{method}\n{canonicalized_headers}{canonicalized_resource}'
        return Encoder.base_64encode_to_string(Signer.hmac_sha1sign(string_to_sign, secret))

    async def get_signature_v1_async(
        self,
        bucket_name: str,
        pathname: str,
        method: str,
        query: Dict[str, str],
        headers: Dict[str, str],
        secret: str,
    ) -> str:
        resource = ''
        string_to_sign = ''
        if not UtilClient.empty(bucket_name):
            resource = f'/{bucket_name}'
        resource = f'{resource}{pathname}'
        canonicalized_resource = await self.build_canonicalized_resource_async(resource, query)
        canonicalized_headers = await self.build_canonicalized_headers_async(headers)
        string_to_sign = f'{method}\n{canonicalized_headers}{canonicalized_resource}'
        return Encoder.base_64encode_to_string(Signer.hmac_sha1sign(string_to_sign, secret))

    def build_canonicalized_resource(
        self,
        pathname: str,
        query: Dict[str, str],
    ) -> str:
        sub_resources_map = {}
        canonicalized_resource = pathname
        if not UtilClient.empty(pathname):
            paths = StringClient.split(pathname, f'?', 2)
            canonicalized_resource = paths[0]
            if UtilClient.equal_number(ArrayClient.size(paths), 2):
                sub_resources = StringClient.split(paths[1], '&', 0)
                for sub in sub_resources:
                    has_excepts = False
                    for excepts in self._except_signed_params:
                        if StringClient.contains(sub, excepts):
                            has_excepts = True
                    if not has_excepts:
                        item = StringClient.split(sub, '=', 0)
                        key = item[0]
                        value = None
                        if UtilClient.equal_number(ArrayClient.size(item), 2):
                            value = item[1]
                        sub_resources_map[key] = value
        sub_resources_array = MapClient.key_set(sub_resources_map)
        new_query_list = sub_resources_array
        if not UtilClient.is_unset(query):
            query_list = MapClient.key_set(query)
            new_query_list = ArrayClient.concat(query_list, sub_resources_array)
        sorted_params = ArrayClient.asc_sort(new_query_list)
        separator = '?'
        for param_name in sorted_params:
            if ArrayClient.contains(self._default_signed_params, param_name):
                canonicalized_resource = f'{canonicalized_resource}{separator}{param_name}'
                if not UtilClient.is_unset(query) and not UtilClient.is_unset(query.get(param_name)):
                    canonicalized_resource = f'{canonicalized_resource}={query.get(param_name)}'
                elif not UtilClient.is_unset(sub_resources_map.get(param_name)):
                    canonicalized_resource = f'{canonicalized_resource}={sub_resources_map.get(param_name)}'
            elif ArrayClient.contains(sub_resources_array, param_name):
                canonicalized_resource = f'{canonicalized_resource}{separator}{param_name}'
                if not UtilClient.is_unset(sub_resources_map.get(param_name)):
                    canonicalized_resource = f'{canonicalized_resource}={sub_resources_map.get(param_name)}'
            separator = '&'
        return canonicalized_resource

    async def build_canonicalized_resource_async(
        self,
        pathname: str,
        query: Dict[str, str],
    ) -> str:
        sub_resources_map = {}
        canonicalized_resource = pathname
        if not UtilClient.empty(pathname):
            paths = StringClient.split(pathname, f'?', 2)
            canonicalized_resource = paths[0]
            if UtilClient.equal_number(ArrayClient.size(paths), 2):
                sub_resources = StringClient.split(paths[1], '&', 0)
                for sub in sub_resources:
                    has_excepts = False
                    for excepts in self._except_signed_params:
                        if StringClient.contains(sub, excepts):
                            has_excepts = True
                    if not has_excepts:
                        item = StringClient.split(sub, '=', 0)
                        key = item[0]
                        value = None
                        if UtilClient.equal_number(ArrayClient.size(item), 2):
                            value = item[1]
                        sub_resources_map[key] = value
        sub_resources_array = MapClient.key_set(sub_resources_map)
        new_query_list = sub_resources_array
        if not UtilClient.is_unset(query):
            query_list = MapClient.key_set(query)
            new_query_list = ArrayClient.concat(query_list, sub_resources_array)
        sorted_params = ArrayClient.asc_sort(new_query_list)
        separator = '?'
        for param_name in sorted_params:
            if ArrayClient.contains(self._default_signed_params, param_name):
                canonicalized_resource = f'{canonicalized_resource}{separator}{param_name}'
                if not UtilClient.is_unset(query) and not UtilClient.is_unset(query.get(param_name)):
                    canonicalized_resource = f'{canonicalized_resource}={query.get(param_name)}'
                elif not UtilClient.is_unset(sub_resources_map.get(param_name)):
                    canonicalized_resource = f'{canonicalized_resource}={sub_resources_map.get(param_name)}'
            elif ArrayClient.contains(sub_resources_array, param_name):
                canonicalized_resource = f'{canonicalized_resource}{separator}{param_name}'
                if not UtilClient.is_unset(sub_resources_map.get(param_name)):
                    canonicalized_resource = f'{canonicalized_resource}={sub_resources_map.get(param_name)}'
            separator = '&'
        return canonicalized_resource

    def build_canonicalized_headers(
        self,
        headers: Dict[str, str],
    ) -> str:
        canonicalized_headers = ''
        content_type = headers.get('content-type')
        if UtilClient.is_unset(content_type):
            content_type = ''
        content_md_5 = headers.get('content-md5')
        if UtilClient.is_unset(content_md_5):
            content_md_5 = ''
        canonicalized_headers = f"{canonicalized_headers}{content_md_5}\n{content_type}\n{headers.get('date')}\n"
        keys = MapClient.key_set(headers)
        sorted_headers = ArrayClient.asc_sort(keys)
        for header in sorted_headers:
            if StringClient.contains(StringClient.to_lower(header), 'x-oss-'):
                canonicalized_headers = f'{canonicalized_headers}{header}:{headers.get(header)}\n'
        return canonicalized_headers

    async def build_canonicalized_headers_async(
        self,
        headers: Dict[str, str],
    ) -> str:
        canonicalized_headers = ''
        content_type = headers.get('content-type')
        if UtilClient.is_unset(content_type):
            content_type = ''
        content_md_5 = headers.get('content-md5')
        if UtilClient.is_unset(content_md_5):
            content_md_5 = ''
        canonicalized_headers = f"{canonicalized_headers}{content_md_5}\n{content_type}\n{headers.get('date')}\n"
        keys = MapClient.key_set(headers)
        sorted_headers = ArrayClient.asc_sort(keys)
        for header in sorted_headers:
            if StringClient.contains(StringClient.to_lower(header), 'x-oss-'):
                canonicalized_headers = f'{canonicalized_headers}{header}:{headers.get(header)}\n'
        return canonicalized_headers

    def get_signature_v2(
        self,
        bucket_name: str,
        pathname: str,
        method: str,
        query: Dict[str, str],
        headers: Dict[str, str],
        secret: str,
    ) -> str:
        return ''

    async def get_signature_v2_async(
        self,
        bucket_name: str,
        pathname: str,
        method: str,
        query: Dict[str, str],
        headers: Dict[str, str],
        secret: str,
    ) -> str:
        return ''
