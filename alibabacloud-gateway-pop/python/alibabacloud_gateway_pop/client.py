# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from alibabacloud_darabonba_encode_util.encoder import Encoder
from Tea.exceptions import TeaException
from alibabacloud_darabonba_signature_util.signer import Signer
from Tea.core import TeaCore
from typing import Dict, Any, List

from alibabacloud_gateway_spi.client import Client as SPIClient
from alibabacloud_gateway_spi import models as spi_models
from alibabacloud_openapi_util.client import Client as OpenApiUtilClient
from alibabacloud_tea_util.client import Client as UtilClient
from alibabacloud_endpoint_util.client import Client as EndpointUtilClient
from alibabacloud_darabonba_array.client import Client as ArrayClient
from alibabacloud_darabonba_string.client import Client as StringClient
from alibabacloud_darabonba_map.client import Client as MapClient


class Client(SPIClient):
    def __init__(self):
        super().__init__()

    def modify_configuration(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        request = context.request
        config = context.configuration
        config.endpoint = self.get_endpoint(request.product_id, config.region_id, config.endpoint_rule, config.network, config.suffix, config.endpoint_map, config.endpoint)

    async def modify_configuration_async(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        request = context.request
        config = context.configuration
        config.endpoint = self.get_endpoint(request.product_id, config.region_id, config.endpoint_rule, config.network, config.suffix, config.endpoint_map, config.endpoint)

    def modify_request(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        request = context.request
        config = context.configuration
        request.headers = TeaCore.merge({
            'host': config.endpoint,
            'x-acs-version': request.version,
            'x-acs-action': request.action,
            'user-agent': request.user_agent,
            'x-acs-date': OpenApiUtilClient.get_timestamp(),
            'x-acs-signature-nonce': UtilClient.get_nonce(),
            'accept': 'application/json'
        }, request.headers)
        signature_algorithm = 'ACS3-HMAC-SHA256'
        if not UtilClient.is_unset(request.signature_algorithm):
            signature_algorithm = request.signature_algorithm
        hashed_request_payload = Encoder.hex_encode(Encoder.hash(UtilClient.to_bytes(''), signature_algorithm))
        if not UtilClient.is_unset(request.stream):
            tmp = UtilClient.read_as_bytes(request.stream)
            hashed_request_payload = Encoder.hex_encode(Encoder.hash(tmp, signature_algorithm))
            request.stream = tmp
            request.headers['content-type'] = 'application/octet-stream'
        else:
            if not UtilClient.is_unset(request.body):
                if UtilClient.equal_string(request.req_body_type, 'json'):
                    json_obj = UtilClient.to_jsonstring(request.body)
                    hashed_request_payload = Encoder.hex_encode(Encoder.hash(UtilClient.to_bytes(json_obj), signature_algorithm))
                    request.stream = json_obj
                    request.headers['content-type'] = 'application/json; charset=utf-8'
                else:
                    m = UtilClient.assert_as_map(request.body)
                    form_obj = OpenApiUtilClient.to_form(m)
                    hashed_request_payload = Encoder.hex_encode(Encoder.hash(UtilClient.to_bytes(form_obj), signature_algorithm))
                    request.stream = form_obj
                    request.headers['content-type'] = 'application/x-www-form-urlencoded'
        request.headers['x-acs-content-sha256'] = hashed_request_payload
        if not UtilClient.equal_string(request.auth_type, 'Anonymous'):
            credential = request.credential
            access_key_id = credential.get_access_key_id()
            access_key_secret = credential.get_access_key_secret()
            security_token = credential.get_security_token()
            if not UtilClient.empty(security_token):
                request.headers['x-acs-accesskey-id'] = access_key_id
                request.headers['x-acs-security-token'] = security_token
            request.headers['Authorization'] = self.get_authorization(request.pathname, request.method, request.query, request.headers, signature_algorithm, hashed_request_payload, access_key_id, access_key_secret)

    async def modify_request_async(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        request = context.request
        config = context.configuration
        request.headers = TeaCore.merge({
            'host': config.endpoint,
            'x-acs-version': request.version,
            'x-acs-action': request.action,
            'user-agent': request.user_agent,
            'x-acs-date': OpenApiUtilClient.get_timestamp(),
            'x-acs-signature-nonce': UtilClient.get_nonce(),
            'accept': 'application/json'
        }, request.headers)
        signature_algorithm = 'ACS3-HMAC-SHA256'
        if not UtilClient.is_unset(request.signature_algorithm):
            signature_algorithm = request.signature_algorithm
        hashed_request_payload = Encoder.hex_encode(Encoder.hash(UtilClient.to_bytes(''), signature_algorithm))
        if not UtilClient.is_unset(request.stream):
            tmp = await UtilClient.read_as_bytes_async(request.stream)
            hashed_request_payload = Encoder.hex_encode(Encoder.hash(tmp, signature_algorithm))
            request.stream = tmp
            request.headers['content-type'] = 'application/octet-stream'
        else:
            if not UtilClient.is_unset(request.body):
                if UtilClient.equal_string(request.req_body_type, 'json'):
                    json_obj = UtilClient.to_jsonstring(request.body)
                    hashed_request_payload = Encoder.hex_encode(Encoder.hash(UtilClient.to_bytes(json_obj), signature_algorithm))
                    request.stream = json_obj
                    request.headers['content-type'] = 'application/json; charset=utf-8'
                else:
                    m = UtilClient.assert_as_map(request.body)
                    form_obj = OpenApiUtilClient.to_form(m)
                    hashed_request_payload = Encoder.hex_encode(Encoder.hash(UtilClient.to_bytes(form_obj), signature_algorithm))
                    request.stream = form_obj
                    request.headers['content-type'] = 'application/x-www-form-urlencoded'
        request.headers['x-acs-content-sha256'] = hashed_request_payload
        if not UtilClient.equal_string(request.auth_type, 'Anonymous'):
            credential = request.credential
            access_key_id = await credential.get_access_key_id_async()
            access_key_secret = await credential.get_access_key_secret_async()
            security_token = await credential.get_security_token_async()
            if not UtilClient.empty(security_token):
                request.headers['x-acs-accesskey-id'] = access_key_id
                request.headers['x-acs-security-token'] = security_token
            request.headers['Authorization'] = await self.get_authorization_async(request.pathname, request.method, request.query, request.headers, signature_algorithm, hashed_request_payload, access_key_id, access_key_secret)

    def modify_response(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        request = context.request
        config = context.configuration
        response = context.response
        if UtilClient.is_4xx(response.status_code) or UtilClient.is_5xx(response.status_code):
            _res = UtilClient.read_as_json(response.body)
            err = UtilClient.assert_as_map(_res)
            request_id = self.default_any(err.get('RequestId'), err.get('requestId'))
            raise TeaException({
                'code': f"{self.default_any(err.get('Code'), err.get('code'))}",
                'message': f"code: {response.status_code}, {self.default_any(err.get('Message'), err.get('message'))} request id: {request_id}",
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
            obj = UtilClient.read_as_json(response.body)
            res = UtilClient.assert_as_map(obj)
            response.deserialized_body = res
        elif UtilClient.equal_string(request.body_type, 'array'):
            arr = UtilClient.read_as_json(response.body)
            response.deserialized_body = arr
        else:
            response.deserialized_body = UtilClient.read_as_string(response.body)

    async def modify_response_async(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        request = context.request
        config = context.configuration
        response = context.response
        if UtilClient.is_4xx(response.status_code) or UtilClient.is_5xx(response.status_code):
            _res = await UtilClient.read_as_json_async(response.body)
            err = UtilClient.assert_as_map(_res)
            request_id = self.default_any(err.get('RequestId'), err.get('requestId'))
            raise TeaException({
                'code': f"{self.default_any(err.get('Code'), err.get('code'))}",
                'message': f"code: {response.status_code}, {self.default_any(err.get('Message'), err.get('message'))} request id: {request_id}",
                'data': err
            })
        if UtilClient.equal_string(request.body_type, 'binary'):
            response.deserialized_body = response.body
        elif UtilClient.equal_string(request.body_type, 'byte'):
            byt = await UtilClient.read_as_bytes_async(response.body)
            response.deserialized_body = byt
        elif UtilClient.equal_string(request.body_type, 'string'):
            str = await UtilClient.read_as_string_async(response.body)
            response.deserialized_body = str
        elif UtilClient.equal_string(request.body_type, 'json'):
            obj = await UtilClient.read_as_json_async(response.body)
            res = UtilClient.assert_as_map(obj)
            response.deserialized_body = res
        elif UtilClient.equal_string(request.body_type, 'array'):
            arr = await UtilClient.read_as_json_async(response.body)
            response.deserialized_body = arr
        else:
            response.deserialized_body = await UtilClient.read_as_string_async(response.body)

    def get_endpoint(
        self,
        product_id: str,
        region_id: str,
        endpoint_rule: str,
        network: str,
        suffix: str,
        endpoint_map: Dict[str, str],
        endpoint: str,
    ) -> str:
        if not UtilClient.empty(endpoint):
            return endpoint
        if not UtilClient.is_unset(endpoint_map) and not UtilClient.empty(endpoint_map.get(region_id)):
            return endpoint_map.get(region_id)
        return EndpointUtilClient.get_endpoint_rules(product_id, region_id, endpoint_rule, network, suffix)

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
        signature_algorithm: str,
        payload: str,
        ak: str,
        secret: str,
    ) -> str:
        signature = self.get_signature(pathname, method, query, headers, signature_algorithm, payload, secret)
        return f"{signature_algorithm}  Credential={ak},SignedHeaders={ArrayClient.join(self.get_signed_headers(headers), ';')},Signature={signature}"

    async def get_authorization_async(
        self,
        pathname: str,
        method: str,
        query: Dict[str, str],
        headers: Dict[str, str],
        signature_algorithm: str,
        payload: str,
        ak: str,
        secret: str,
    ) -> str:
        signature = await self.get_signature_async(pathname, method, query, headers, signature_algorithm, payload, secret)
        return f"{signature_algorithm}  Credential={ak},SignedHeaders={ArrayClient.join(self.get_signed_headers(headers), ';')},Signature={signature}"

    def get_signature(
        self,
        pathname: str,
        method: str,
        query: Dict[str, str],
        headers: Dict[str, str],
        signature_algorithm: str,
        payload: str,
        secret: str,
    ) -> str:
        canonical_uri = '/'
        if not UtilClient.empty(pathname):
            canonical_uri = pathname
        string_to_sign = ''
        canonicalized_resource = self.build_canonicalized_resource(query)
        canonicalized_headers = self.build_canonicalized_headers(headers)
        signed_headers = self.get_signed_headers(headers)
        string_to_sign = f"{method}\n{canonical_uri}\n{canonicalized_resource}\n{canonicalized_headers}\n{ArrayClient.join(signed_headers, ';')}\n{payload}"
        hex = Encoder.hex_encode(Encoder.hash(UtilClient.to_bytes(string_to_sign), signature_algorithm))
        string_to_sign = f'{signature_algorithm}\n{hex}'
        signature = UtilClient.to_bytes('')
        if StringClient.equals(signature_algorithm, 'ACS3-HMAC-SHA256'):
            signature = Signer.hmac_sha256sign(string_to_sign, secret)
        elif StringClient.equals(signature_algorithm, 'ACS3-HMAC-SM3'):
            signature = Signer.hmac_sm3sign(string_to_sign, secret)
        elif StringClient.equals(signature_algorithm, 'ACS3-RSA-SHA256'):
            signature = Signer.sha256with_rsasign(string_to_sign, secret)
        return Encoder.hex_encode(signature)

    async def get_signature_async(
        self,
        pathname: str,
        method: str,
        query: Dict[str, str],
        headers: Dict[str, str],
        signature_algorithm: str,
        payload: str,
        secret: str,
    ) -> str:
        canonical_uri = '/'
        if not UtilClient.empty(pathname):
            canonical_uri = pathname
        string_to_sign = ''
        canonicalized_resource = await self.build_canonicalized_resource_async(query)
        canonicalized_headers = await self.build_canonicalized_headers_async(headers)
        signed_headers = await self.get_signed_headers_async(headers)
        string_to_sign = f"{method}\n{canonical_uri}\n{canonicalized_resource}\n{canonicalized_headers}\n{ArrayClient.join(signed_headers, ';')}\n{payload}"
        hex = Encoder.hex_encode(Encoder.hash(UtilClient.to_bytes(string_to_sign), signature_algorithm))
        string_to_sign = f'{signature_algorithm}\n{hex}'
        signature = UtilClient.to_bytes('')
        if StringClient.equals(signature_algorithm, 'ACS3-HMAC-SHA256'):
            signature = Signer.hmac_sha256sign(string_to_sign, secret)
        elif StringClient.equals(signature_algorithm, 'ACS3-HMAC-SM3'):
            signature = Signer.hmac_sm3sign(string_to_sign, secret)
        elif StringClient.equals(signature_algorithm, 'ACS3-RSA-SHA256'):
            signature = Signer.sha256with_rsasign(string_to_sign, secret)
        return Encoder.hex_encode(signature)

    def build_canonicalized_resource(
        self,
        query: Dict[str, str],
    ) -> str:
        canonicalized_resource = ''
        if not UtilClient.is_unset(query):
            query_array = MapClient.key_set(query)
            sorted_query_array = ArrayClient.asc_sort(query_array)
            for key in sorted_query_array:
                canonicalized_resource = f'{canonicalized_resource}&{Encoder.percent_encode(key)}'
                if not UtilClient.empty(query.get(key)):
                    canonicalized_resource = f'{canonicalized_resource}={Encoder.percent_encode(query.get(key))}'
        return canonicalized_resource

    async def build_canonicalized_resource_async(
        self,
        query: Dict[str, str],
    ) -> str:
        canonicalized_resource = ''
        if not UtilClient.is_unset(query):
            query_array = MapClient.key_set(query)
            sorted_query_array = ArrayClient.asc_sort(query_array)
            for key in sorted_query_array:
                canonicalized_resource = f'{canonicalized_resource}&{Encoder.percent_encode(key)}'
                if not UtilClient.empty(query.get(key)):
                    canonicalized_resource = f'{canonicalized_resource}={Encoder.percent_encode(query.get(key))}'
        return canonicalized_resource

    def build_canonicalized_headers(
        self,
        headers: Dict[str, str],
    ) -> str:
        canonicalized_headers = ''
        sorted_headers = self.get_signed_headers(headers)
        for header in sorted_headers:
            canonicalized_headers = f'{canonicalized_headers}{header}:{StringClient.trim(headers.get(header))}\n'
        return canonicalized_headers

    async def build_canonicalized_headers_async(
        self,
        headers: Dict[str, str],
    ) -> str:
        canonicalized_headers = ''
        sorted_headers = await self.get_signed_headers_async(headers)
        for header in sorted_headers:
            canonicalized_headers = f'{canonicalized_headers}{header}:{StringClient.trim(headers.get(header))}\n'
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
            if StringClient.has_prefix(lower_key, 'x-acs-') or StringClient.equals(lower_key, 'host') or StringClient.equals(lower_key, 'content-type'):
                if not StringClient.contains(tmp, lower_key):
                    tmp = f'{tmp}{separator}{lower_key}'
                    separator = ';'
        return StringClient.split(tmp, ';', 0)

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
            if StringClient.has_prefix(lower_key, 'x-acs-') or StringClient.equals(lower_key, 'host') or StringClient.equals(lower_key, 'content-type'):
                if not StringClient.contains(tmp, lower_key):
                    tmp = f'{tmp}{separator}{lower_key}'
                    separator = ';'
        return StringClient.split(tmp, ';', 0)
