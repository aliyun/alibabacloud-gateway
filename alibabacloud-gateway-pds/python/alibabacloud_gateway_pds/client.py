# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from Tea.exceptions import TeaException
from alibabacloud_darabonba_signature_util.signer import Signer
from alibabacloud_darabonba_encode_util.encoder import Encoder
from Tea.core import TeaCore
from typing import Any, Dict, List

from alibabacloud_gateway_spi.client import Client as SPIClient
from alibabacloud_gateway_spi import models as spi_models
from alibabacloud_tea_util.client import Client as UtilClient
from alibabacloud_openapi_util.client import Client as OpenApiUtilClient
from alibabacloud_darabonba_string.client import Client as StringClient
from alibabacloud_darabonba_map.client import Client as MapClient
from alibabacloud_darabonba_array.client import Client as ArrayClient


class Client(SPIClient):
    def __init__(self):
        super().__init__()

    def modify_configuration(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        config = context.configuration
        config.endpoint = self.get_endpoint(config.network, config.endpoint)

    async def modify_configuration_async(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        config = context.configuration
        config.endpoint = await self.get_endpoint_async(config.network, config.endpoint)

    def modify_request(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
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
        if not UtilClient.equal_string(request.auth_type, 'Anonymous') and not UtilClient.is_unset(request.credential):
            credential = request.credential
            auth_type = credential.get_type()
            if UtilClient.equal_string(auth_type, 'bearer'):
                bearer_token = credential.get_bearer_token()
                request.headers['x-acs-bearer-token'] = bearer_token
                request.headers['Authorization'] = f'Bearer {bearer_token}'
            else:
                access_key_id = credential.get_access_key_id()
                access_key_secret = credential.get_access_key_secret()
                security_token = credential.get_security_token()
                if not UtilClient.empty(security_token):
                    request.headers['x-acs-security-token'] = security_token
                request.headers['Authorization'] = self.get_authorization(request.pathname, request.method, request.query, request.headers, access_key_id, access_key_secret)

    async def modify_request_async(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
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
            tmp = await UtilClient.read_as_bytes_async(request.stream)
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
        if not UtilClient.equal_string(request.auth_type, 'Anonymous') and not UtilClient.is_unset(request.credential):
            credential = request.credential
            auth_type = credential.get_type()
            if UtilClient.equal_string(auth_type, 'bearer'):
                bearer_token = credential.get_bearer_token()
                request.headers['x-acs-bearer-token'] = bearer_token
                request.headers['Authorization'] = f'Bearer {bearer_token}'
            else:
                access_key_id = await credential.get_access_key_id_async()
                access_key_secret = await credential.get_access_key_secret_async()
                security_token = await credential.get_security_token_async()
                if not UtilClient.empty(security_token):
                    request.headers['x-acs-security-token'] = security_token
                request.headers['Authorization'] = await self.get_authorization_async(request.pathname, request.method, request.query, request.headers, access_key_id, access_key_secret)

    def modify_response(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        request = context.request
        response = context.response
        if UtilClient.is_4xx(response.status_code) or UtilClient.is_5xx(response.status_code):
            _res = UtilClient.read_as_json(response.body)
            err = UtilClient.assert_as_map(_res)
            headers = response.headers
            request_id = headers.get('x-ca-request-id')
            err['statusCode'] = response.status_code
            err['requestId'] = request_id
            raise TeaException({
                'code': f"{self.default_any(err.get('Code'), err.get('code'))}",
                'message': f"{self.default_any(err.get('Message'), err.get('message'))}",
                'data': err
            })
        if not UtilClient.is_unset(response.body):
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
                response.deserialized_body = UtilClient.read_as_json(response.body)
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
        if UtilClient.is_4xx(response.status_code) or UtilClient.is_5xx(response.status_code):
            _res = await UtilClient.read_as_json_async(response.body)
            err = UtilClient.assert_as_map(_res)
            headers = response.headers
            request_id = headers.get('x-ca-request-id')
            err['statusCode'] = response.status_code
            err['requestId'] = request_id
            raise TeaException({
                'code': f"{self.default_any(err.get('Code'), err.get('code'))}",
                'message': f"{self.default_any(err.get('Message'), err.get('message'))}",
                'data': err
            })
        if not UtilClient.is_unset(response.body):
            if UtilClient.equal_number(response.status_code, 204):
                await UtilClient.read_as_string_async(response.body)
            elif UtilClient.equal_string(request.body_type, 'binary'):
                response.deserialized_body = response.body
            elif UtilClient.equal_string(request.body_type, 'byte'):
                byt = await UtilClient.read_as_bytes_async(response.body)
                response.deserialized_body = byt
            elif UtilClient.equal_string(request.body_type, 'string'):
                str = await UtilClient.read_as_string_async(response.body)
                response.deserialized_body = str
            elif UtilClient.equal_string(request.body_type, 'json'):
                response.deserialized_body = await UtilClient.read_as_json_async(response.body)
            elif UtilClient.equal_string(request.body_type, 'array'):
                response.deserialized_body = await UtilClient.read_as_json_async(response.body)
            else:
                response.deserialized_body = await UtilClient.read_as_string_async(response.body)

    def get_endpoint(
        self,
        network: str,
        endpoint: str,
    ) -> str:
        real_endpoint = 'api.aliyunpds.com'
        if not UtilClient.empty(endpoint):
            real_endpoint = endpoint
        if not UtilClient.empty(network) and StringClient.equals(network, 'vpc'):
            real_endpoint = StringClient.replace(real_endpoint, 'api.aliyunpds.com', 'api-vpc.aliyunpds.com', None)
            real_endpoint = StringClient.replace(real_endpoint, 'admin.aliyunpds.com', 'admin-vpc.aliyunpds.com', None)
        return real_endpoint

    async def get_endpoint_async(
        self,
        network: str,
        endpoint: str,
    ) -> str:
        real_endpoint = 'api.aliyunpds.com'
        if not UtilClient.empty(endpoint):
            real_endpoint = endpoint
        if not UtilClient.empty(network) and StringClient.equals(network, 'vpc'):
            real_endpoint = StringClient.replace(real_endpoint, 'api.aliyunpds.com', 'api-vpc.aliyunpds.com', None)
            real_endpoint = StringClient.replace(real_endpoint, 'admin.aliyunpds.com', 'admin-vpc.aliyunpds.com', None)
        return real_endpoint

    def default_any(
        self,
        input_value: Any,
        default_value: Any,
    ) -> Any:
        if UtilClient.is_unset(input_value):
            return default_value
        return input_value

    def get_authorization(
        self,
        pathname: str,
        method: str,
        query: Dict[str, str],
        headers: Dict[str, str],
        ak: str,
        secret: str,
    ) -> str:
        signature = self.get_signature(pathname, method, query, headers, secret)
        return f'acs {ak}:{signature}'

    async def get_authorization_async(
        self,
        pathname: str,
        method: str,
        query: Dict[str, str],
        headers: Dict[str, str],
        ak: str,
        secret: str,
    ) -> str:
        signature = await self.get_signature_async(pathname, method, query, headers, secret)
        return f'acs {ak}:{signature}'

    def get_signature(
        self,
        pathname: str,
        method: str,
        query: Dict[str, str],
        headers: Dict[str, str],
        secret: str,
    ) -> str:
        string_to_sign = ''
        canonicalized_resource = self.build_canonicalized_resource(pathname, query)
        canonicalized_headers = self.build_canonicalized_headers(headers)
        string_to_sign = f'{method}\n{canonicalized_headers}{canonicalized_resource}'
        signature = Signer.hmac_sha1sign(string_to_sign, secret)
        return Encoder.base_64encode_to_string(signature)

    async def get_signature_async(
        self,
        pathname: str,
        method: str,
        query: Dict[str, str],
        headers: Dict[str, str],
        secret: str,
    ) -> str:
        string_to_sign = ''
        canonicalized_resource = await self.build_canonicalized_resource_async(pathname, query)
        canonicalized_headers = await self.build_canonicalized_headers_async(headers)
        string_to_sign = f'{method}\n{canonicalized_headers}{canonicalized_resource}'
        signature = Signer.hmac_sha1sign(string_to_sign, secret)
        return Encoder.base_64encode_to_string(signature)

    def build_canonicalized_resource(
        self,
        pathname: str,
        query: Dict[str, str],
    ) -> str:
        canonicalized_resource = pathname
        if not UtilClient.is_unset(query):
            query_array = MapClient.key_set(query)
            sorted_query_array = ArrayClient.asc_sort(query_array)
            separator = '?'
            for key in sorted_query_array:
                canonicalized_resource = f'{canonicalized_resource}{separator}{key}'
                if not UtilClient.empty(query.get(key)):
                    canonicalized_resource = f'{canonicalized_resource}={query.get(key)}'
                separator = '&'
        return canonicalized_resource

    async def build_canonicalized_resource_async(
        self,
        pathname: str,
        query: Dict[str, str],
    ) -> str:
        canonicalized_resource = pathname
        if not UtilClient.is_unset(query):
            query_array = MapClient.key_set(query)
            sorted_query_array = ArrayClient.asc_sort(query_array)
            separator = '?'
            for key in sorted_query_array:
                canonicalized_resource = f'{canonicalized_resource}{separator}{key}'
                if not UtilClient.empty(query.get(key)):
                    canonicalized_resource = f'{canonicalized_resource}={query.get(key)}'
                separator = '&'
        return canonicalized_resource

    def build_canonicalized_headers(
        self,
        headers: Dict[str, str],
    ) -> str:
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
        canonicalized_headers = f'{accept}\n{content_md_5}\n{content_type}\n{date}\n'
        sorted_headers = self.get_signed_headers(headers)
        for header in sorted_headers:
            value = headers.get(header)
            value_trim = StringClient.trim(value)
            canonicalized_headers = f'{canonicalized_headers}{header}:{value_trim}\n'
        return canonicalized_headers

    async def build_canonicalized_headers_async(
        self,
        headers: Dict[str, str],
    ) -> str:
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
        canonicalized_headers = f'{accept}\n{content_md_5}\n{content_type}\n{date}\n'
        sorted_headers = await self.get_signed_headers_async(headers)
        for header in sorted_headers:
            value = headers.get(header)
            value_trim = StringClient.trim(value)
            canonicalized_headers = f'{canonicalized_headers}{header}:{value_trim}\n'
        return canonicalized_headers

    def get_signed_headers(
        self,
        headers: Dict[str, str],
    ) -> List[str]:
        headers_array = MapClient.key_set(headers)
        sorted_headers_array = ArrayClient.asc_sort(headers_array)
        tmp = ''
        separator = ''
        for key in sorted_headers_array:
            lower_key = StringClient.to_lower(key)
            if StringClient.has_prefix(lower_key, 'x-acs-'):
                if not StringClient.contains(tmp, lower_key):
                    tmp = f'{tmp}{separator}{lower_key}'
                    separator = ';'
        return StringClient.split(tmp, ';', None)

    async def get_signed_headers_async(
        self,
        headers: Dict[str, str],
    ) -> List[str]:
        headers_array = MapClient.key_set(headers)
        sorted_headers_array = ArrayClient.asc_sort(headers_array)
        tmp = ''
        separator = ''
        for key in sorted_headers_array:
            lower_key = StringClient.to_lower(key)
            if StringClient.has_prefix(lower_key, 'x-acs-'):
                if not StringClient.contains(tmp, lower_key):
                    tmp = f'{tmp}{separator}{lower_key}'
                    separator = ';'
        return StringClient.split(tmp, ';', None)
