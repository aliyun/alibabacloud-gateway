# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from alibabacloud_darabonba_encode_util.encoder import Encoder
from alibabacloud_darabonba_signature_util.signer import Signer
from Tea.exceptions import TeaException
from typing import List, Dict
from Tea.core import TeaCore

from alibabacloud_gateway_spi.client import Client as SPIClient
from alibabacloud_gateway_spi import models as spi_models
from alibabacloud_tea_util.client import Client as UtilClient
from alibabacloud_darabonba_map.client import Client as MapClient
from alibabacloud_darabonba_string.client import Client as StringClient
from alibabacloud_tea_xml.client import Client as XMLClient
from alibabacloud_openapi_util.client import Client as OpenApiUtilClient
from alibabacloud_oss_util.client import Client as OSSUtilClient
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
            'versionId',
            'wormId'
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
        if not UtilClient.is_unset(request.headers.get('x-oss-meta-*')):
            tmp = UtilClient.parse_json(request.headers.get('x-oss-meta-*'))
            map_data = UtilClient.assert_as_map(tmp)
            meta_data = UtilClient.stringify_map_value(map_data)
            meta_key_set = MapClient.key_set(meta_data)
            request.headers['x-oss-meta-*'] = None
            for key in meta_key_set:
                new_key = f'x-oss-meta-{key}'
                request.headers[new_key] = meta_data.get(key)
        config = context.configuration
        region_id = config.region_id
        credential = request.credential
        access_key_id = credential.get_access_key_id()
        access_key_secret = credential.get_access_key_secret()
        security_token = credential.get_security_token()
        if not UtilClient.empty(security_token):
            request.headers['x-oss-security-token'] = security_token
        if not UtilClient.is_unset(request.body):
            if StringClient.equals(request.req_body_type, 'xml'):
                req_body_map = UtilClient.assert_as_map(request.body)
                xml_str = XMLClient.to_xml(req_body_map)
                request.stream = xml_str
                request.headers['content-type'] = 'application/xml'
                request.headers['content-md5'] = Encoder.base_64encode_to_string(Signer.md5sign(xml_str))
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
        request.headers['authorization'] = self.get_authorization(request.signature_version, bucket_name, request.pathname, request.method, request.query, request.headers, access_key_id, access_key_secret, region_id)

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
        if not UtilClient.is_unset(request.headers.get('x-oss-meta-*')):
            tmp = UtilClient.parse_json(request.headers.get('x-oss-meta-*'))
            map_data = UtilClient.assert_as_map(tmp)
            meta_data = UtilClient.stringify_map_value(map_data)
            meta_key_set = MapClient.key_set(meta_data)
            request.headers['x-oss-meta-*'] = None
            for key in meta_key_set:
                new_key = f'x-oss-meta-{key}'
                request.headers[new_key] = meta_data.get(key)
        config = context.configuration
        region_id = config.region_id
        credential = request.credential
        access_key_id = await credential.get_access_key_id_async()
        access_key_secret = await credential.get_access_key_secret_async()
        security_token = await credential.get_security_token_async()
        if not UtilClient.empty(security_token):
            request.headers['x-oss-security-token'] = security_token
        if not UtilClient.is_unset(request.body):
            if StringClient.equals(request.req_body_type, 'xml'):
                req_body_map = UtilClient.assert_as_map(request.body)
                xml_str = XMLClient.to_xml(req_body_map)
                request.stream = xml_str
                request.headers['content-type'] = 'application/xml'
                request.headers['content-md5'] = Encoder.base_64encode_to_string(Signer.md5sign(xml_str))
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
        request.headers['authorization'] = await self.get_authorization_async(request.signature_version, bucket_name, request.pathname, request.method, request.query, request.headers, access_key_id, access_key_secret, region_id)

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
            if not UtilClient.empty(body_str):
                resp_map = XMLClient.parse_xml(body_str, None)
                err = UtilClient.assert_as_map(resp_map.get('Error'))
                raise TeaException({
                    'code': err.get('Code'),
                    'message': err.get('Message'),
                    'data': {
                        'statusCode': response.status_code,
                        'requestId': err.get('RequestId'),
                        'ecCode': err.get('EC'),
                        'Recommend': err.get('RecommendDoc'),
                        'hostId': err.get('HostId')
                    }
                })
            else:
                headers = response.headers
                request_id = headers.get('x-oss-request-id')
                ec_code = headers.get('x-oss-ec-code')
                raise TeaException({
                    'code': response.status_code,
                    'message': None,
                    'data': {
                        'statusCode': response.status_code,
                        'requestId': f'{request_id}',
                        'ecCode': ec_code
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
            if UtilClient.equal_number(response.status_code, 204):
                UtilClient.read_as_string(response.body)
            elif StringClient.equals(request.body_type, 'xml'):
                body_str = UtilClient.read_as_string(response.body)
                response.deserialized_body = body_str
                if not UtilClient.empty(body_str):
                    resp_struct = self.get_response_body_schema(request.action)
                    result = XMLClient.parse_xml(body_str, resp_struct)
                    try:
                        response.deserialized_body = UtilClient.assert_as_map(result)
                    except Exception as error:
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
            if not UtilClient.empty(body_str):
                resp_map = XMLClient.parse_xml(body_str, None)
                err = UtilClient.assert_as_map(resp_map.get('Error'))
                raise TeaException({
                    'code': err.get('Code'),
                    'message': err.get('Message'),
                    'data': {
                        'statusCode': response.status_code,
                        'requestId': err.get('RequestId'),
                        'ecCode': err.get('EC'),
                        'Recommend': err.get('RecommendDoc'),
                        'hostId': err.get('HostId')
                    }
                })
            else:
                headers = response.headers
                request_id = headers.get('x-oss-request-id')
                ec_code = headers.get('x-oss-ec-code')
                raise TeaException({
                    'code': response.status_code,
                    'message': None,
                    'data': {
                        'statusCode': response.status_code,
                        'requestId': f'{request_id}',
                        'ecCode': ec_code
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
            if UtilClient.equal_number(response.status_code, 204):
                await UtilClient.read_as_string_async(response.body)
            elif StringClient.equals(request.body_type, 'xml'):
                body_str = await UtilClient.read_as_string_async(response.body)
                response.deserialized_body = body_str
                if not UtilClient.empty(body_str):
                    resp_struct = await self.get_response_body_schema_async(request.action)
                    result = XMLClient.parse_xml(body_str, resp_struct)
                    try:
                        response.deserialized_body = UtilClient.assert_as_map(result)
                    except Exception as error:
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
        region_id: str,
    ) -> str:
        sign = ''
        if not UtilClient.is_unset(signature_version):
            if StringClient.equals(signature_version, 'v1'):
                sign = self.get_signature_v1(bucket_name, pathname, method, query, headers, secret)
                return f'OSS {ak}:{sign}'
            if StringClient.equals(signature_version, 'v2'):
                sign = self.get_signature_v2(bucket_name, pathname, method, query, headers, secret)
                return f'OSS2 AccessKeyId:{ak},Signature:{sign}'
        date_time = OpenApiUtilClient.get_timestamp()
        date_time = StringClient.replace(date_time, '-', '', None)
        date_time = StringClient.replace(date_time, ':', '', None)
        headers['x-oss-date'] = date_time
        headers['x-oss-content-sha256'] = 'UNSIGNED-PAYLOAD'
        only_date = StringClient.sub_string(date_time, 0, 8)
        cred = f'{ak}/{only_date}/{region_id}/oss/aliyun_v4_request'
        sign = self.get_signature_v4(bucket_name, pathname, method, query, headers, only_date, region_id, secret)
        return f'OSS4-HMAC-SHA256 Credential={cred}, Signature={sign}'

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
        region_id: str,
    ) -> str:
        sign = ''
        if not UtilClient.is_unset(signature_version):
            if StringClient.equals(signature_version, 'v1'):
                sign = await self.get_signature_v1_async(bucket_name, pathname, method, query, headers, secret)
                return f'OSS {ak}:{sign}'
            if StringClient.equals(signature_version, 'v2'):
                sign = await self.get_signature_v2_async(bucket_name, pathname, method, query, headers, secret)
                return f'OSS2 AccessKeyId:{ak},Signature:{sign}'
        date_time = OpenApiUtilClient.get_timestamp()
        date_time = StringClient.replace(date_time, '-', '', None)
        date_time = StringClient.replace(date_time, ':', '', None)
        headers['x-oss-date'] = date_time
        headers['x-oss-content-sha256'] = 'UNSIGNED-PAYLOAD'
        only_date = StringClient.sub_string(date_time, 0, 8)
        cred = f'{ak}/{only_date}/{region_id}/oss/aliyun_v4_request'
        sign = await self.get_signature_v4_async(bucket_name, pathname, method, query, headers, only_date, region_id, secret)
        return f'OSS4-HMAC-SHA256 Credential={cred}, Signature={sign}'

    def get_sign_key(
        self,
        secret: str,
        only_date: str,
        region_id: str,
    ) -> bytes:
        temp = f'aliyun_v4{secret}'
        res = Signer.hmac_sha256sign(only_date, temp)
        res = Signer.hmac_sha256sign_by_bytes(region_id, res)
        res = Signer.hmac_sha256sign_by_bytes('oss', res)
        res = Signer.hmac_sha256sign_by_bytes('aliyun_v4_request', res)
        return res

    async def get_sign_key_async(
        self,
        secret: str,
        only_date: str,
        region_id: str,
    ) -> bytes:
        temp = f'aliyun_v4{secret}'
        res = Signer.hmac_sha256sign(only_date, temp)
        res = Signer.hmac_sha256sign_by_bytes(region_id, res)
        res = Signer.hmac_sha256sign_by_bytes('oss', res)
        res = Signer.hmac_sha256sign_by_bytes('aliyun_v4_request', res)
        return res

    def get_signature_v4(
        self,
        bucket_name: str,
        pathname: str,
        method: str,
        query: Dict[str, str],
        headers: Dict[str, str],
        only_date: str,
        region_id: str,
        secret: str,
    ) -> str:
        signingkey = self.get_sign_key(secret, only_date, region_id)
        object_name = '/'
        query_map = {}
        if not UtilClient.empty(pathname):
            paths = StringClient.split(pathname, f'?', 2)
            object_name = paths[0]
            if UtilClient.equal_number(ArrayClient.size(paths), 2):
                sub_resources = StringClient.split(paths[1], '&', None)
                for sub in sub_resources:
                    item = StringClient.split(sub, '=', None)
                    key = item[0]
                    key = Encoder.percent_encode(key)
                    key = StringClient.replace(key, '+', '%20', None)
                    value = None
                    if UtilClient.equal_number(ArrayClient.size(item), 2):
                        value = Encoder.percent_encode(item[1])
                        value = StringClient.replace(value, '+', '%20', None)
                    # for go : queryMap[tea.StringValue(key)] = value
                    query_map[key] = value
        canonicalized_uri = '/'
        if not UtilClient.empty(bucket_name):
            canonicalized_uri = f'/{bucket_name}{object_name}'
        # for java:
        # String suffix = (!canonicalizedUri.equals("/") && canonicalizedUri.endsWith("/"))? "/" : "";
        # canonicalizedUri = com.aliyun.openapiutil.Client.getEncodePath(canonicalizedUri) + suffix;
        canonicalized_uri = OpenApiUtilClient.get_encode_path(canonicalized_uri)
        for query_key in MapClient.key_set(query):
            query_value = None
            if not UtilClient.empty(query.get(query_key)):
                query_value = Encoder.percent_encode(query.get(query_key))
                query_value = StringClient.replace(query_value, '+', '%20', None)
            query_key = Encoder.percent_encode(query_key)
            query_key = StringClient.replace(query_key, '+', '%20', None)
            # for go : queryMap[tea.StringValue(queryKey)] = queryValue
            query_map[query_key] = query_value
        canonicalized_query_string = self.build_canonicalized_query_string_v4(query_map)
        canonicalized_headers = self.build_canonicalized_headers_v4(headers)
        payload = 'UNSIGNED-PAYLOAD'
        canonical_request = f'{method}\n{canonicalized_uri}\n{canonicalized_query_string}\n{canonicalized_headers}\n\n{payload}'
        hex = Encoder.hex_encode(Encoder.hash(UtilClient.to_bytes(canonical_request), 'ACS4-HMAC-SHA256'))
        scope = f'{only_date}/{region_id}/oss/aliyun_v4_request'
        string_to_sign = f"OSS4-HMAC-SHA256\n{headers.get('x-oss-date')}\n{scope}\n{hex}"
        signature = Signer.hmac_sha256sign_by_bytes(string_to_sign, signingkey)
        return Encoder.hex_encode(signature)

    async def get_signature_v4_async(
        self,
        bucket_name: str,
        pathname: str,
        method: str,
        query: Dict[str, str],
        headers: Dict[str, str],
        only_date: str,
        region_id: str,
        secret: str,
    ) -> str:
        signingkey = await self.get_sign_key_async(secret, only_date, region_id)
        object_name = '/'
        query_map = {}
        if not UtilClient.empty(pathname):
            paths = StringClient.split(pathname, f'?', 2)
            object_name = paths[0]
            if UtilClient.equal_number(ArrayClient.size(paths), 2):
                sub_resources = StringClient.split(paths[1], '&', None)
                for sub in sub_resources:
                    item = StringClient.split(sub, '=', None)
                    key = item[0]
                    key = Encoder.percent_encode(key)
                    key = StringClient.replace(key, '+', '%20', None)
                    value = None
                    if UtilClient.equal_number(ArrayClient.size(item), 2):
                        value = Encoder.percent_encode(item[1])
                        value = StringClient.replace(value, '+', '%20', None)
                    # for go : queryMap[tea.StringValue(key)] = value
                    query_map[key] = value
        canonicalized_uri = '/'
        if not UtilClient.empty(bucket_name):
            canonicalized_uri = f'/{bucket_name}{object_name}'
        # for java:
        # String suffix = (!canonicalizedUri.equals("/") && canonicalizedUri.endsWith("/"))? "/" : "";
        # canonicalizedUri = com.aliyun.openapiutil.Client.getEncodePath(canonicalizedUri) + suffix;
        canonicalized_uri = OpenApiUtilClient.get_encode_path(canonicalized_uri)
        for query_key in MapClient.key_set(query):
            query_value = None
            if not UtilClient.empty(query.get(query_key)):
                query_value = Encoder.percent_encode(query.get(query_key))
                query_value = StringClient.replace(query_value, '+', '%20', None)
            query_key = Encoder.percent_encode(query_key)
            query_key = StringClient.replace(query_key, '+', '%20', None)
            # for go : queryMap[tea.StringValue(queryKey)] = queryValue
            query_map[query_key] = query_value
        canonicalized_query_string = await self.build_canonicalized_query_string_v4_async(query_map)
        canonicalized_headers = await self.build_canonicalized_headers_v4_async(headers)
        payload = 'UNSIGNED-PAYLOAD'
        canonical_request = f'{method}\n{canonicalized_uri}\n{canonicalized_query_string}\n{canonicalized_headers}\n\n{payload}'
        hex = Encoder.hex_encode(Encoder.hash(UtilClient.to_bytes(canonical_request), 'ACS4-HMAC-SHA256'))
        scope = f'{only_date}/{region_id}/oss/aliyun_v4_request'
        string_to_sign = f"OSS4-HMAC-SHA256\n{headers.get('x-oss-date')}\n{scope}\n{hex}"
        signature = Signer.hmac_sha256sign_by_bytes(string_to_sign, signingkey)
        return Encoder.hex_encode(signature)

    def build_canonicalized_query_string_v4(
        self,
        query_map: Dict[str, str],
    ) -> str:
        canonicalized_query_string = ''
        if not UtilClient.is_unset(query_map):
            query_array = MapClient.key_set(query_map)
            sorted_query_array = ArrayClient.asc_sort(query_array)
            separator = ''
            for key in sorted_query_array:
                canonicalized_query_string = f'{canonicalized_query_string}{separator}{key}'
                if not UtilClient.empty(query_map.get(key)):
                    canonicalized_query_string = f'{canonicalized_query_string}={query_map.get(key)}'
                separator = '&'
        return canonicalized_query_string

    async def build_canonicalized_query_string_v4_async(
        self,
        query_map: Dict[str, str],
    ) -> str:
        canonicalized_query_string = ''
        if not UtilClient.is_unset(query_map):
            query_array = MapClient.key_set(query_map)
            sorted_query_array = ArrayClient.asc_sort(query_array)
            separator = ''
            for key in sorted_query_array:
                canonicalized_query_string = f'{canonicalized_query_string}{separator}{key}'
                if not UtilClient.empty(query_map.get(key)):
                    canonicalized_query_string = f'{canonicalized_query_string}={query_map.get(key)}'
                separator = '&'
        return canonicalized_query_string

    def build_canonicalized_headers_v4(
        self,
        headers: Dict[str, str],
    ) -> str:
        canonicalized_headers = ''
        headers_array = MapClient.key_set(headers)
        sorted_headers_array = ArrayClient.asc_sort(headers_array)
        for key in sorted_headers_array:
            lower_key = StringClient.to_lower(key)
            if StringClient.has_prefix(lower_key, 'x-oss-') or StringClient.equals(lower_key, 'content-type') or StringClient.equals(lower_key, 'content-md5'):
                canonicalized_headers = f'{canonicalized_headers}{key}:{StringClient.trim(headers.get(key))}\n'
        return canonicalized_headers

    async def build_canonicalized_headers_v4_async(
        self,
        headers: Dict[str, str],
    ) -> str:
        canonicalized_headers = ''
        headers_array = MapClient.key_set(headers)
        sorted_headers_array = ArrayClient.asc_sort(headers_array)
        for key in sorted_headers_array:
            lower_key = StringClient.to_lower(key)
            if StringClient.has_prefix(lower_key, 'x-oss-') or StringClient.equals(lower_key, 'content-type') or StringClient.equals(lower_key, 'content-md5'):
                canonicalized_headers = f'{canonicalized_headers}{key}:{StringClient.trim(headers.get(key))}\n'
        return canonicalized_headers

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
                sub_resources = StringClient.split(paths[1], '&', None)
                for sub in sub_resources:
                    has_excepts = False
                    for excepts in self._except_signed_params:
                        if StringClient.contains(sub, excepts):
                            has_excepts = True
                    if not has_excepts:
                        item = StringClient.split(sub, '=', None)
                        key = item[0]
                        value = None
                        if UtilClient.equal_number(ArrayClient.size(item), 2):
                            value = item[1]
                        # for go : subResourcesMap[tea.StringValue(key)] = value
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
                sub_resources = StringClient.split(paths[1], '&', None)
                for sub in sub_resources:
                    has_excepts = False
                    for excepts in self._except_signed_params:
                        if StringClient.contains(sub, excepts):
                            has_excepts = True
                    if not has_excepts:
                        item = StringClient.split(sub, '=', None)
                        key = item[0]
                        value = None
                        if UtilClient.equal_number(ArrayClient.size(item), 2):
                            value = item[1]
                        # for go : subResourcesMap[tea.StringValue(key)] = value
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
            if StringClient.contains(StringClient.to_lower(header), 'x-oss-') and not UtilClient.is_unset(headers.get(header)):
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
            if StringClient.contains(StringClient.to_lower(header), 'x-oss-') and not UtilClient.is_unset(headers.get(header)):
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

    @staticmethod
    def get_response_body_schema(
        api_name: str,
    ) -> object:
        return None

    @staticmethod
    async def get_response_body_schema_async(
        api_name: str,
    ) -> object:
        return None
