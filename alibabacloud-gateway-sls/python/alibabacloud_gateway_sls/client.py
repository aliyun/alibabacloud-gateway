# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from alibabacloud_darabonba_encode_util.encoder import Encoder
from alibabacloud_darabonba_signature_util.signer import Signer
from Tea.exceptions import TeaException
from typing import Dict, List
from Tea.core import TeaCore

from alibabacloud_gateway_spi.client import Client as SPIClient
from alibabacloud_gateway_spi import models as spi_models
from alibabacloud_tea_util.client import Client as UtilClient
from alibabacloud_darabonba_string.client import Client as StringClient
from alibabacloud_gateway_sls_util.client import Client as SLS_UtilClient
from alibabacloud_darabonba_map.client import Client as MapClient
from alibabacloud_darabonba_array.client import Client as ArrayClient
from alibabacloud_openapi_util.client import Client as OpenApiUtilClient


class Client(SPIClient):

    def __init__(self):
        super().__init__()
        self._resp_body_decompress_type = {
            'PullLogs': [
                'zstd',
                'lz4',
                'gzip'
            ],
            'GetLogsV2': [
                'zstd',
                'lz4',
                'gzip'
            ],
            'PreviewSPL': [
                'lz4'
            ]
        }
        self._req_body_compress_type = {
            'PutLogs': [
                'zstd',
                'lz4',
                'deflate'
            ],
            'PreviewSPL': [
                'lz4'
            ]
        }

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
        project = host_map.get('project')
        config = context.configuration
        credential = request.credential
        credential_model = credential.get_credential()
        access_key_id = credential_model.access_key_id
        access_key_secret = credential_model.access_key_secret
        security_token = credential_model.security_token
        if not UtilClient.empty(security_token):
            request.headers['x-acs-security-token'] = security_token
        signature_version = UtilClient.default_string(request.signature_version, 'v1')
        final_compress_type = self.get_final_request_compress_type(request.action, request.headers)
        content_hash = ''
        # get body bytes
        body_bytes = None
        if not UtilClient.is_unset(request.body):
            # PutLogs
            if StringClient.equals(request.action, 'PutLogs'):
                body_bytes = SLS_UtilClient.serialize_log_group_to_pb(request.body)
                request.headers['content-type'] = 'application/x-protobuf'
            elif StringClient.equals(request.req_body_type, 'json') or StringClient.equals(request.req_body_type, 'formData'):
                request.headers['content-type'] = 'application/json'
                body_str = UtilClient.to_jsonstring(request.body)
                body_bytes = UtilClient.to_bytes(body_str)
            elif StringClient.equals(request.req_body_type, 'binary'):
                # content-type: application/octet-stream
                body_bytes = UtilClient.assert_as_bytes(request.body)
        # get body raw size
        body_raw_size = '0'
        raw_size_ref = request.headers.get('x-log-bodyrawsize')
        # for php bug, Argument #1 ($value) could not be passed by reference
        if not UtilClient.is_unset(raw_size_ref):
            body_raw_size = raw_size_ref
        elif not UtilClient.is_unset(request.body):
            body_raw_size = f'{SLS_UtilClient.bytes_length(body_bytes)}'
        # compress if needed, and set body and hash
        if not UtilClient.is_unset(request.body):
            if not UtilClient.empty(final_compress_type):
                compressed = SLS_UtilClient.compress(body_bytes, final_compress_type)
                body_bytes = compressed
            content_hash = self.make_content_hash(body_bytes, signature_version)
            request.stream = body_bytes
        host = self.get_host(config.network, project, config.endpoint)
        request.headers = TeaCore.merge({
            'accept': 'application/json',
            'host': host,
            'user-agent': request.user_agent,
            'x-log-apiversion': '0.6.0'
        }, request.headers)
        request.headers['x-log-bodyrawsize'] = body_raw_size
        self.set_default_accept_encoding(request.action, request.headers)
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

    async def modify_request_async(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        request = context.request
        host_map = {}
        if not UtilClient.is_unset(request.host_map):
            host_map = request.host_map
        project = host_map.get('project')
        config = context.configuration
        credential = request.credential
        credential_model = await credential.get_credential_async()
        access_key_id = credential_model.access_key_id
        access_key_secret = credential_model.access_key_secret
        security_token = credential_model.security_token
        if not UtilClient.empty(security_token):
            request.headers['x-acs-security-token'] = security_token
        signature_version = UtilClient.default_string(request.signature_version, 'v1')
        final_compress_type = await self.get_final_request_compress_type_async(request.action, request.headers)
        content_hash = ''
        # get body bytes
        body_bytes = None
        if not UtilClient.is_unset(request.body):
            # PutLogs
            if StringClient.equals(request.action, 'PutLogs'):
                body_bytes = await SLS_UtilClient.serialize_log_group_to_pb_async(request.body)
                request.headers['content-type'] = 'application/x-protobuf'
            elif StringClient.equals(request.req_body_type, 'json') or StringClient.equals(request.req_body_type, 'formData'):
                request.headers['content-type'] = 'application/json'
                body_str = UtilClient.to_jsonstring(request.body)
                body_bytes = UtilClient.to_bytes(body_str)
            elif StringClient.equals(request.req_body_type, 'binary'):
                # content-type: application/octet-stream
                body_bytes = UtilClient.assert_as_bytes(request.body)
        # get body raw size
        body_raw_size = '0'
        raw_size_ref = request.headers.get('x-log-bodyrawsize')
        # for php bug, Argument #1 ($value) could not be passed by reference
        if not UtilClient.is_unset(raw_size_ref):
            body_raw_size = raw_size_ref
        elif not UtilClient.is_unset(request.body):
            body_raw_size = f'{SLS_UtilClient.bytes_length(body_bytes)}'
        # compress if needed, and set body and hash
        if not UtilClient.is_unset(request.body):
            if not UtilClient.empty(final_compress_type):
                compressed = await SLS_UtilClient.compress_async(body_bytes, final_compress_type)
                body_bytes = compressed
            content_hash = await self.make_content_hash_async(body_bytes, signature_version)
            request.stream = body_bytes
        host = await self.get_host_async(config.network, project, config.endpoint)
        request.headers = TeaCore.merge({
            'accept': 'application/json',
            'host': host,
            'user-agent': request.user_agent,
            'x-log-apiversion': '0.6.0'
        }, request.headers)
        request.headers['x-log-bodyrawsize'] = body_raw_size
        await self.set_default_accept_encoding_async(request.action, request.headers)
        await self.build_request_async(context)
        # move param in path to query
        if StringClient.equals(signature_version, 'v4'):
            if UtilClient.empty(content_hash):
                content_hash = 'e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855'
            date = await self.get_date_iso8601_async()
            request.headers['x-log-date'] = date
            request.headers['x-log-content-sha256'] = content_hash
            request.headers['authorization'] = await self.get_authorization_v4_async(context, date, content_hash, access_key_id, access_key_secret)
            return
        if not UtilClient.empty(access_key_id):
            request.headers['x-log-signaturemethod'] = 'hmac-sha256'
        request.headers['date'] = UtilClient.get_date_utcstring()
        request.headers['content-md5'] = content_hash
        request.headers['authorization'] = await self.get_authorization_async(request.pathname, request.method, request.query, request.headers, access_key_id, access_key_secret)

    def get_final_request_compress_type(
        self,
        action: str,
        headers: Dict[str, str],
    ) -> str:
        compress_type = headers.get('x-log-compresstype')
        raw_size = headers.get('x-log-bodyrawsize')
        # for php bug, Argument #1 ($value) could not be passed by reference
        # 1. already compressed, has x-log-compresstype and x-log-bodyrawsize in header, we dont need compress here
        if not UtilClient.is_unset(compress_type) and not UtilClient.is_unset(raw_size):
            return ''
        # 2. not compressed, but has x-log-compresstype in header, we need compress here
        if not UtilClient.is_unset(compress_type):
            return compress_type
        # 3. not compressed, in pre-defined api list, try use default supported compress type in order
        encodings = self._req_body_compress_type.get(action)
        if not UtilClient.is_unset(encodings):
            for encoding in encodings:
                if SLS_UtilClient.is_compressor_available(encoding):
                    headers['x-log-compresstype'] = encoding
                    # set header x-log-compresstype
                    return encoding
        # 4. otherwise we dont need compress here
        return ''

    async def get_final_request_compress_type_async(
        self,
        action: str,
        headers: Dict[str, str],
    ) -> str:
        compress_type = headers.get('x-log-compresstype')
        raw_size = headers.get('x-log-bodyrawsize')
        # for php bug, Argument #1 ($value) could not be passed by reference
        # 1. already compressed, has x-log-compresstype and x-log-bodyrawsize in header, we dont need compress here
        if not UtilClient.is_unset(compress_type) and not UtilClient.is_unset(raw_size):
            return ''
        # 2. not compressed, but has x-log-compresstype in header, we need compress here
        if not UtilClient.is_unset(compress_type):
            return compress_type
        # 3. not compressed, in pre-defined api list, try use default supported compress type in order
        encodings = self._req_body_compress_type.get(action)
        if not UtilClient.is_unset(encodings):
            for encoding in encodings:
                if await SLS_UtilClient.is_compressor_available_async(encoding):
                    headers['x-log-compresstype'] = encoding
                    # set header x-log-compresstype
                    return encoding
        # 4. otherwise we dont need compress here
        return ''

    def set_default_accept_encoding(
        self,
        action: str,
        headers: Dict[str, str],
    ) -> None:
        accept_encoding = headers.get('Accept-Encoding')
        # for php warning, dont rm this line
        if not UtilClient.is_unset(accept_encoding):
            return
        encodings = self._resp_body_decompress_type.get(action)
        if not UtilClient.is_unset(encodings):
            for c in encodings:
                if SLS_UtilClient.is_decompressor_available(c):
                    headers['Accept-Encoding'] = c
                    return

    async def set_default_accept_encoding_async(
        self,
        action: str,
        headers: Dict[str, str],
    ) -> None:
        accept_encoding = headers.get('Accept-Encoding')
        # for php warning, dont rm this line
        if not UtilClient.is_unset(accept_encoding):
            return
        encodings = self._resp_body_decompress_type.get(action)
        if not UtilClient.is_unset(encodings):
            for c in encodings:
                if await SLS_UtilClient.is_decompressor_available_async(c):
                    headers['Accept-Encoding'] = c
                    return

    def make_content_hash(
        self,
        content: bytes,
        signature_version: str,
    ) -> str:
        if StringClient.equals(signature_version, 'v4'):
            if UtilClient.is_unset(content) or UtilClient.equal_string(f'{SLS_UtilClient.bytes_length(content)}', '0'):
                return ''
            return StringClient.to_lower(Encoder.hex_encode(Encoder.hash(content, 'SLS4-HMAC-SHA256')))
        return StringClient.to_upper(Encoder.hex_encode(Signer.md5sign_for_bytes(content)))

    async def make_content_hash_async(
        self,
        content: bytes,
        signature_version: str,
    ) -> str:
        if StringClient.equals(signature_version, 'v4'):
            if UtilClient.is_unset(content) or UtilClient.equal_string(f'{SLS_UtilClient.bytes_length(content)}', '0'):
                return ''
            return StringClient.to_lower(Encoder.hex_encode(Encoder.hash(content, 'SLS4-HMAC-SHA256')))
        return StringClient.to_upper(Encoder.hex_encode(Signer.md5sign_for_bytes(content)))

    def modify_response(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        request = context.request
        response = context.response
        if UtilClient.is_4xx(response.status_code) or UtilClient.is_5xx(response.status_code):
            error = UtilClient.read_as_json(response.body)
            res_map = UtilClient.assert_as_map(error)
            raise TeaException({
                'code': res_map.get('errorCode'),
                'message': res_map.get('errorMessage'),
                'accessDeniedDetail': res_map.get('accessDeniedDetail'),
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

    async def modify_response_async(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        request = context.request
        response = context.response
        if UtilClient.is_4xx(response.status_code) or UtilClient.is_5xx(response.status_code):
            error = await UtilClient.read_as_json_async(response.body)
            res_map = UtilClient.assert_as_map(error)
            raise TeaException({
                'code': res_map.get('errorCode'),
                'message': res_map.get('errorMessage'),
                'accessDeniedDetail': res_map.get('accessDeniedDetail'),
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
                uncompressed_data = await SLS_UtilClient.read_and_uncompress_block_async(response.body, compress_type, bodyraw_size)
            if UtilClient.equal_string(request.body_type, 'binary'):
                response.deserialized_body = uncompressed_data
            elif UtilClient.equal_string(request.body_type, 'byte'):
                byt = await UtilClient.read_as_bytes_async(uncompressed_data)
                response.deserialized_body = byt
            elif UtilClient.equal_string(request.body_type, 'string'):
                response.deserialized_body = await UtilClient.read_as_string_async(uncompressed_data)
            elif UtilClient.equal_string(request.body_type, 'json'):
                obj = await UtilClient.read_as_json_async(uncompressed_data)
                # var res = Util.assertAsMap(obj);
                response.deserialized_body = obj
            elif UtilClient.equal_string(request.body_type, 'array'):
                response.deserialized_body = await UtilClient.read_as_json_async(uncompressed_data)
            else:
                response.deserialized_body = await UtilClient.read_as_string_async(uncompressed_data)

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
            if StringClient.equals(network, 'intranet'):
                return f'{region_id}-intranet.log.aliyuncs.com'
            elif StringClient.equals(network, 'accelerate'):
                return f'log-global.aliyuncs.com'
            elif StringClient.equals(network, 'share'):
                if StringClient.equals(region_id, 'cn-hangzhou-corp') or StringClient.equals(region_id, 'cn-shanghai-corp'):
                    return f'{region_id}.sls.aliyuncs.com'
                elif StringClient.equals(region_id, 'cn-zhangbei-corp'):
                    return f'zhangbei-corp-share.log.aliyuncs.com'
                return f'{region_id}-share.log.aliyuncs.com'
        return f'{region_id}.log.aliyuncs.com'

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
            if StringClient.equals(network, 'intranet'):
                return f'{region_id}-intranet.log.aliyuncs.com'
            elif StringClient.equals(network, 'accelerate'):
                return f'log-global.aliyuncs.com'
            elif StringClient.equals(network, 'share'):
                if StringClient.equals(region_id, 'cn-hangzhou-corp') or StringClient.equals(region_id, 'cn-shanghai-corp'):
                    return f'{region_id}.sls.aliyuncs.com'
                elif StringClient.equals(region_id, 'cn-zhangbei-corp'):
                    return f'zhangbei-corp-share.log.aliyuncs.com'
                return f'{region_id}-share.log.aliyuncs.com'
        return f'{region_id}.log.aliyuncs.com'

    def get_host(
        self,
        network: str,
        project: str,
        endpoint: str,
    ) -> str:
        if UtilClient.is_unset(project):
            return endpoint
        return f'{project}.{endpoint}'

    async def get_host_async(
        self,
        network: str,
        project: str,
        endpoint: str,
    ) -> str:
        if UtilClient.is_unset(project):
            return endpoint
        return f'{project}.{endpoint}'

    def get_authorization(
        self,
        pathname: str,
        method: str,
        query: Dict[str, str],
        headers: Dict[str, str],
        ak: str,
        secret: str,
    ) -> str:
        sign = self.get_signature(pathname, method, query, headers, secret)
        return f'LOG {ak}:{sign}'

    async def get_authorization_async(
        self,
        pathname: str,
        method: str,
        query: Dict[str, str],
        headers: Dict[str, str],
        ak: str,
        secret: str,
    ) -> str:
        sign = await self.get_signature_async(pathname, method, query, headers, secret)
        return f'LOG {ak}:{sign}'

    def get_signature(
        self,
        pathname: str,
        method: str,
        query: Dict[str, str],
        headers: Dict[str, str],
        secret: str,
    ) -> str:
        resource = pathname
        string_to_sign = ''
        canonicalized_resource = self.build_canonicalized_resource(resource, query)
        canonicalized_headers = self.build_canonicalized_headers(headers)
        string_to_sign = f'{method}\n{canonicalized_headers}{canonicalized_resource}'
        return Encoder.base_64encode_to_string(Signer.hmac_sha256sign(string_to_sign, secret))

    async def get_signature_async(
        self,
        pathname: str,
        method: str,
        query: Dict[str, str],
        headers: Dict[str, str],
        secret: str,
    ) -> str:
        resource = pathname
        string_to_sign = ''
        canonicalized_resource = await self.build_canonicalized_resource_async(resource, query)
        canonicalized_headers = await self.build_canonicalized_headers_async(headers)
        string_to_sign = f'{method}\n{canonicalized_headers}{canonicalized_resource}'
        return Encoder.base_64encode_to_string(Signer.hmac_sha256sign(string_to_sign, secret))

    def build_canonicalized_resource(
        self,
        pathname: str,
        query: Dict[str, str],
    ) -> str:
        canonicalized_resource = pathname
        if not UtilClient.is_unset(query):
            query_list = MapClient.key_set(query)
            sorted_params = ArrayClient.asc_sort(query_list)
            separator = '?'
            for param_name in sorted_params:
                canonicalized_resource = f'{canonicalized_resource}{separator}{param_name}'
                param_value = query.get(param_name)
                if not UtilClient.is_unset(param_value):
                    canonicalized_resource = f'{canonicalized_resource}={param_value}'
                separator = '&'
        return canonicalized_resource

    async def build_canonicalized_resource_async(
        self,
        pathname: str,
        query: Dict[str, str],
    ) -> str:
        canonicalized_resource = pathname
        if not UtilClient.is_unset(query):
            query_list = MapClient.key_set(query)
            sorted_params = ArrayClient.asc_sort(query_list)
            separator = '?'
            for param_name in sorted_params:
                canonicalized_resource = f'{canonicalized_resource}{separator}{param_name}'
                param_value = query.get(param_name)
                if not UtilClient.is_unset(param_value):
                    canonicalized_resource = f'{canonicalized_resource}={param_value}'
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
            if StringClient.contains(StringClient.to_lower(header), 'x-log-') or StringClient.contains(StringClient.to_lower(header), 'x-acs-'):
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
            if StringClient.contains(StringClient.to_lower(header), 'x-log-') or StringClient.contains(StringClient.to_lower(header), 'x-acs-'):
                canonicalized_headers = f'{canonicalized_headers}{header}:{headers.get(header)}\n'
        return canonicalized_headers

    def build_request(
        self,
        context: spi_models.InterceptorContext,
    ) -> None:
        request = context.request
        resource = request.pathname
        if not UtilClient.empty(resource):
            paths = StringClient.split(resource, f'?', 2)
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

    async def build_request_async(
        self,
        context: spi_models.InterceptorContext,
    ) -> None:
        request = context.request
        resource = request.pathname
        if not UtilClient.empty(resource):
            paths = StringClient.split(resource, f'?', 2)
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

    def get_authorization_v4(
        self,
        context: spi_models.InterceptorContext,
        date: str,
        content_hash: str,
        access_key_id: str,
        access_key_secret: str,
    ) -> str:
        region = self.get_region(context)
        header_str = self.get_signed_header_str_v4(context.request.headers)
        date_short = StringClient.sub_string(date, 0, 8)
        date_short = StringClient.replace(date_short, 'T', '', None)
        # for fix php sdk bug
        scope = f'{date_short}/{region}/sls/aliyun_v4_request'
        signingkey = self.get_signingkey_v4('SLS4-HMAC-SHA256', access_key_secret, region, date_short)
        signature = self.get_signature_v4(context, 'SLS4-HMAC-SHA256', header_str, date, scope, content_hash, signingkey)
        return f"{'SLS4-HMAC-SHA256'} Credential={access_key_id}/{scope},Signature={signature}"

    async def get_authorization_v4_async(
        self,
        context: spi_models.InterceptorContext,
        date: str,
        content_hash: str,
        access_key_id: str,
        access_key_secret: str,
    ) -> str:
        region = await self.get_region_async(context)
        header_str = await self.get_signed_header_str_v4_async(context.request.headers)
        date_short = StringClient.sub_string(date, 0, 8)
        date_short = StringClient.replace(date_short, 'T', '', None)
        # for fix php sdk bug
        scope = f'{date_short}/{region}/sls/aliyun_v4_request'
        signingkey = await self.get_signingkey_v4_async('SLS4-HMAC-SHA256', access_key_secret, region, date_short)
        signature = await self.get_signature_v4_async(context, 'SLS4-HMAC-SHA256', header_str, date, scope, content_hash, signingkey)
        return f"{'SLS4-HMAC-SHA256'} Credential={access_key_id}/{scope},Signature={signature}"

    def get_signingkey_v4(
        self,
        signature_algorithm: str,
        secret: str,
        region: str,
        date: str,
    ) -> bytes:
        sc_1 = f'aliyun_v4{secret}'
        sc_2 = Signer.hmac_sha256sign(date, sc_1)
        sc_3 = Signer.hmac_sha256sign_by_bytes(region, sc_2)
        sc_4 = Signer.hmac_sha256sign_by_bytes('sls', sc_3)
        return Signer.hmac_sha256sign_by_bytes('aliyun_v4_request', sc_4)

    async def get_signingkey_v4_async(
        self,
        signature_algorithm: str,
        secret: str,
        region: str,
        date: str,
    ) -> bytes:
        sc_1 = f'aliyun_v4{secret}'
        sc_2 = Signer.hmac_sha256sign(date, sc_1)
        sc_3 = Signer.hmac_sha256sign_by_bytes(region, sc_2)
        sc_4 = Signer.hmac_sha256sign_by_bytes('sls', sc_3)
        return Signer.hmac_sha256sign_by_bytes('aliyun_v4_request', sc_4)

    def get_signature_v4(
        self,
        context: spi_models.InterceptorContext,
        signature_algorithm: str,
        signed_header_str: str,
        date: str,
        scope: str,
        content_sha_256: str,
        signingkey: bytes,
    ) -> str:
        request = context.request
        canonical_uri = '/'
        if not UtilClient.empty(request.pathname):
            canonical_uri = request.pathname
        resources = self.build_canonicalized_resource_v4(request.query)
        headers = self.build_canonicalized_headers_v4(request.headers, signed_header_str)
        string_to_hash = f'{request.method}\n{canonical_uri}\n{resources}\n{headers}\n{signed_header_str}\n{content_sha_256}'
        hex = Encoder.hex_encode(Encoder.hash(UtilClient.to_bytes(string_to_hash), signature_algorithm))
        string_to_sign = f'{signature_algorithm}\n{date}\n{scope}\n{hex}'
        signature = Signer.hmac_sha256sign_by_bytes(string_to_sign, signingkey)
        return Encoder.hex_encode(signature)

    async def get_signature_v4_async(
        self,
        context: spi_models.InterceptorContext,
        signature_algorithm: str,
        signed_header_str: str,
        date: str,
        scope: str,
        content_sha_256: str,
        signingkey: bytes,
    ) -> str:
        request = context.request
        canonical_uri = '/'
        if not UtilClient.empty(request.pathname):
            canonical_uri = request.pathname
        resources = await self.build_canonicalized_resource_v4_async(request.query)
        headers = await self.build_canonicalized_headers_v4_async(request.headers, signed_header_str)
        string_to_hash = f'{request.method}\n{canonical_uri}\n{resources}\n{headers}\n{signed_header_str}\n{content_sha_256}'
        hex = Encoder.hex_encode(Encoder.hash(UtilClient.to_bytes(string_to_hash), signature_algorithm))
        string_to_sign = f'{signature_algorithm}\n{date}\n{scope}\n{hex}'
        signature = Signer.hmac_sha256sign_by_bytes(string_to_sign, signingkey)
        return Encoder.hex_encode(signature)

    def build_canonicalized_resource_v4(
        self,
        query: Dict[str, str],
    ) -> str:
        canonicalized_resource = ''
        if not UtilClient.is_unset(query):
            query_array = MapClient.key_set(query)
            sorted_query_array = ArrayClient.asc_sort(query_array)
            separator = ''
            for key in sorted_query_array:
                canonicalized_resource = f'{canonicalized_resource}{separator}{key}'
                if not UtilClient.empty(query.get(key)):
                    canonicalized_resource = f'{canonicalized_resource}={Encoder.percent_encode(query.get(key))}'
                separator = '&'
        return canonicalized_resource

    async def build_canonicalized_resource_v4_async(
        self,
        query: Dict[str, str],
    ) -> str:
        canonicalized_resource = ''
        if not UtilClient.is_unset(query):
            query_array = MapClient.key_set(query)
            sorted_query_array = ArrayClient.asc_sort(query_array)
            separator = ''
            for key in sorted_query_array:
                canonicalized_resource = f'{canonicalized_resource}{separator}{key}'
                if not UtilClient.empty(query.get(key)):
                    canonicalized_resource = f'{canonicalized_resource}={Encoder.percent_encode(query.get(key))}'
                separator = '&'
        return canonicalized_resource

    def build_canonicalized_headers_v4(
        self,
        headers: Dict[str, str],
        signed_header_str: str,
    ) -> str:
        canonicalized_headers = ''
        signed_headers = StringClient.split(signed_header_str, ';', None)
        for header in signed_headers:
            canonicalized_headers = f'{canonicalized_headers}{header}:{StringClient.trim(headers.get(header))}\n'
        return canonicalized_headers

    async def build_canonicalized_headers_v4_async(
        self,
        headers: Dict[str, str],
        signed_header_str: str,
    ) -> str:
        canonicalized_headers = ''
        signed_headers = StringClient.split(signed_header_str, ';', None)
        for header in signed_headers:
            canonicalized_headers = f'{canonicalized_headers}{header}:{StringClient.trim(headers.get(header))}\n'
        return canonicalized_headers

    def get_signed_header_str_v4(
        self,
        headers: Dict[str, str],
    ) -> str:
        headers_array = MapClient.key_set(headers)
        sorted_headers_array = ArrayClient.asc_sort(headers_array)
        tmp = ''
        separator = ''
        for key in sorted_headers_array:
            lower_key = StringClient.to_lower(key)
            if StringClient.has_prefix(lower_key, 'x-log-') or StringClient.has_prefix(lower_key, 'x-acs-') or StringClient.equals(lower_key, 'host') or StringClient.equals(lower_key, 'content-type'):
                tmp = f'{tmp}{separator}{lower_key}'
                separator = ';'
        return tmp

    async def get_signed_header_str_v4_async(
        self,
        headers: Dict[str, str],
    ) -> str:
        headers_array = MapClient.key_set(headers)
        sorted_headers_array = ArrayClient.asc_sort(headers_array)
        tmp = ''
        separator = ''
        for key in sorted_headers_array:
            lower_key = StringClient.to_lower(key)
            if StringClient.has_prefix(lower_key, 'x-log-') or StringClient.has_prefix(lower_key, 'x-acs-') or StringClient.equals(lower_key, 'host') or StringClient.equals(lower_key, 'content-type'):
                tmp = f'{tmp}{separator}{lower_key}'
                separator = ';'
        return tmp

    def get_region(
        self,
        context: spi_models.InterceptorContext,
    ) -> str:
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

    async def get_region_async(
        self,
        context: spi_models.InterceptorContext,
    ) -> str:
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

    def get_date_iso8601(self) -> str:
        """
        format: YYYYMMDDTHHMMSSZ
        """
        date = OpenApiUtilClient.get_timestamp()
        # 2024-02-04T11:31:58Z
        date = StringClient.replace(date, '-', '', None)
        return StringClient.replace(date, ':', '', None)

    async def get_date_iso8601_async(self) -> str:
        """
        format: YYYYMMDDTHHMMSSZ
        """
        date = OpenApiUtilClient.get_timestamp()
        # 2024-02-04T11:31:58Z
        date = StringClient.replace(date, '-', '', None)
        return StringClient.replace(date, ':', '', None)
