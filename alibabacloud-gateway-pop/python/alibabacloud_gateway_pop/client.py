# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from alibabacloud_darabonba_encode_util.encoder import Encoder
from Tea.exceptions import TeaException
from alibabacloud_darabonba_signature_util.signer import Signer
from Tea.core import TeaCore
from typing import Dict, Any, List

from alibabacloud_gateway_spi.client import Client as SPIClient
from alibabacloud_gateway_spi import models as spi_models
from alibabacloud_tea_util.client import Client as UtilClient
from alibabacloud_openapi_util.client import Client as OpenApiUtilClient
from alibabacloud_darabonba_string.client import Client as StringClient
from alibabacloud_endpoint_util.client import Client as EndpointUtilClient
from alibabacloud_darabonba_array.client import Client as ArrayClient
from alibabacloud_darabonba_map.client import Client as MapClient


class Client(SPIClient):
    _endpoint_suffix: str = None
    _signature_type_prefix: str = None
    _sign_prefix: str = None
    _sha_256: str = None
    _sm_3: str = None

    def __init__(self):
        super().__init__()
        # CLOUD4-\
        self._signature_type_prefix = 'ACS4-'
        # cloud_v4
        self._sign_prefix = 'aliyun_v4'
        self._endpoint_suffix = 'aliyuncs.com'
        self._sha_256 = f'{self._signature_type_prefix}HMAC-SHA256'
        self._sm_3 = f'{self._signature_type_prefix}HMAC-SM3'

    def modify_configuration(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        request = context.request
        config = context.configuration
        attributes = attribute_map.key
        if not UtilClient.is_unset(attributes):
            self._signature_type_prefix = attributes.get('signatureTypePrefix')
            self._sign_prefix = attributes.get('signPrefix')
            self._endpoint_suffix = attributes.get('endpointSuffix')
            self._sha_256 = f'{self._signature_type_prefix}HMAC-SHA256'
            self._sm_3 = f'{self._signature_type_prefix}HMAC-SM3'
        config.endpoint = self.get_endpoint(request.product_id, config.region_id, config.endpoint_rule, config.network, config.suffix, config.endpoint_map, config.endpoint)

    async def modify_configuration_async(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        request = context.request
        config = context.configuration
        attributes = attribute_map.key
        if not UtilClient.is_unset(attributes):
            self._signature_type_prefix = attributes.get('signatureTypePrefix')
            self._sign_prefix = attributes.get('signPrefix')
            self._endpoint_suffix = attributes.get('endpointSuffix')
            self._sha_256 = f'{self._signature_type_prefix}HMAC-SHA256'
            self._sm_3 = f'{self._signature_type_prefix}HMAC-SM3'
        config.endpoint = self.get_endpoint(request.product_id, config.region_id, config.endpoint_rule, config.network, config.suffix, config.endpoint_map, config.endpoint)

    def modify_request(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        request = context.request
        config = context.configuration
        date = OpenApiUtilClient.get_timestamp()
        request.headers = TeaCore.merge({
            'host': config.endpoint,
            'x-acs-version': request.version,
            'user-agent': request.user_agent,
            'x-acs-date': date,
            'x-acs-signature-nonce': UtilClient.get_nonce(),
            'accept': 'application/json'
        }, request.headers)
        if not UtilClient.empty(request.action):
            request.headers['x-acs-action'] = request.action
        signature_algorithm = UtilClient.default_string(request.signature_algorithm, self._sha_256)
        hashed_request_payload = Encoder.hex_encode(Encoder.hash(UtilClient.to_bytes(''), signature_algorithm))
        if not UtilClient.is_unset(request.stream):
            tmp = UtilClient.read_as_bytes(request.stream)
            hashed_request_payload = Encoder.hex_encode(Encoder.hash(tmp, signature_algorithm))
            request.stream = tmp
            request.headers['content-type'] = 'application/octet-stream'
        else:
            if not UtilClient.is_unset(request.body):
                if UtilClient.equal_string(request.req_body_type, 'byte'):
                    byte_obj = UtilClient.assert_as_bytes(request.body)
                    hashed_request_payload = Encoder.hex_encode(Encoder.hash(byte_obj, signature_algorithm))
                    request.stream = byte_obj
                elif UtilClient.equal_string(request.req_body_type, 'json'):
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
        if UtilClient.equal_string(signature_algorithm, self._sm_3):
            request.headers['x-acs-content-sm3'] = hashed_request_payload
        else:
            request.headers['x-acs-content-sha256'] = hashed_request_payload
        if not UtilClient.equal_string(request.auth_type, 'Anonymous'):
            credential = request.credential
            if UtilClient.is_unset(credential):
                raise TeaException({
                    'code': 'ParameterMissing',
                    'message': "'config.credential' can not be unset"
                })
            credential_model = credential.get_credential()
            if not UtilClient.empty(credential_model.provider_name):
                request.headers['x-acs-credentials-provider'] = credential_model.provider_name
            auth_type = credential_model.type
            if UtilClient.equal_string(auth_type, 'bearer'):
                bearer_token = credential_model.bearer_token
                request.headers['x-acs-bearer-token'] = bearer_token
                request.headers['x-acs-signature-type'] = 'BEARERTOKEN'
                request.headers['Authorization'] = f'Bearer {bearer_token}'
            elif UtilClient.equal_string(auth_type, 'id_token'):
                id_token = credential_model.security_token
                request.headers['x-acs-zero-trust-idtoken'] = id_token
            else:
                access_key_id = credential_model.access_key_id
                access_key_secret = credential_model.access_key_secret
                security_token = credential_model.security_token
                if not UtilClient.empty(security_token):
                    request.headers['x-acs-accesskey-id'] = access_key_id
                    request.headers['x-acs-security-token'] = security_token
                date_new = StringClient.sub_string(date, 0, 10)
                date_new = StringClient.replace(date_new, '-', '', None)
                region = self.get_region(request.product_id, config.endpoint, config.region_id)
                signingkey = self.get_signingkey(signature_algorithm, access_key_secret, request.product_id, region, date_new)
                request.headers['Authorization'] = self.get_authorization(request.pathname, request.method, request.query, request.headers, signature_algorithm, hashed_request_payload, access_key_id, signingkey, request.product_id, region, date_new)

    async def modify_request_async(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        request = context.request
        config = context.configuration
        date = OpenApiUtilClient.get_timestamp()
        request.headers = TeaCore.merge({
            'host': config.endpoint,
            'x-acs-version': request.version,
            'user-agent': request.user_agent,
            'x-acs-date': date,
            'x-acs-signature-nonce': UtilClient.get_nonce(),
            'accept': 'application/json'
        }, request.headers)
        if not UtilClient.empty(request.action):
            request.headers['x-acs-action'] = request.action
        signature_algorithm = UtilClient.default_string(request.signature_algorithm, self._sha_256)
        hashed_request_payload = Encoder.hex_encode(Encoder.hash(UtilClient.to_bytes(''), signature_algorithm))
        if not UtilClient.is_unset(request.stream):
            tmp = await UtilClient.read_as_bytes_async(request.stream)
            hashed_request_payload = Encoder.hex_encode(Encoder.hash(tmp, signature_algorithm))
            request.stream = tmp
            request.headers['content-type'] = 'application/octet-stream'
        else:
            if not UtilClient.is_unset(request.body):
                if UtilClient.equal_string(request.req_body_type, 'byte'):
                    byte_obj = UtilClient.assert_as_bytes(request.body)
                    hashed_request_payload = Encoder.hex_encode(Encoder.hash(byte_obj, signature_algorithm))
                    request.stream = byte_obj
                elif UtilClient.equal_string(request.req_body_type, 'json'):
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
        if UtilClient.equal_string(signature_algorithm, self._sm_3):
            request.headers['x-acs-content-sm3'] = hashed_request_payload
        else:
            request.headers['x-acs-content-sha256'] = hashed_request_payload
        if not UtilClient.equal_string(request.auth_type, 'Anonymous'):
            credential = request.credential
            if UtilClient.is_unset(credential):
                raise TeaException({
                    'code': 'ParameterMissing',
                    'message': "'config.credential' can not be unset"
                })
            credential_model = await credential.get_credential_async()
            if not UtilClient.empty(credential_model.provider_name):
                request.headers['x-acs-credentials-provider'] = credential_model.provider_name
            auth_type = credential_model.type
            if UtilClient.equal_string(auth_type, 'bearer'):
                bearer_token = credential_model.bearer_token
                request.headers['x-acs-bearer-token'] = bearer_token
                request.headers['x-acs-signature-type'] = 'BEARERTOKEN'
                request.headers['Authorization'] = f'Bearer {bearer_token}'
            elif UtilClient.equal_string(auth_type, 'id_token'):
                id_token = credential_model.security_token
                request.headers['x-acs-zero-trust-idtoken'] = id_token
            else:
                access_key_id = credential_model.access_key_id
                access_key_secret = credential_model.access_key_secret
                security_token = credential_model.security_token
                if not UtilClient.empty(security_token):
                    request.headers['x-acs-accesskey-id'] = access_key_id
                    request.headers['x-acs-security-token'] = security_token
                date_new = StringClient.sub_string(date, 0, 10)
                date_new = StringClient.replace(date_new, '-', '', None)
                region = self.get_region(request.product_id, config.endpoint, config.region_id)
                signingkey = self.get_signingkey(signature_algorithm, access_key_secret, request.product_id, region, date_new)
                request.headers['Authorization'] = self.get_authorization(request.pathname, request.method, request.query, request.headers, signature_algorithm, hashed_request_payload, access_key_id, signingkey, request.product_id, region, date_new)

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
            request_id = self.default_any(err.get('RequestId'), err.get('requestId'))
            if not UtilClient.is_unset(response.headers.get('x-acs-request-id')):
                request_id = response.headers.get('x-acs-request-id')
            err['statusCode'] = response.status_code
            raise TeaException({
                'code': f"{self.default_any(err.get('Code'), err.get('code'))}",
                'message': f"code: {response.status_code}, {self.default_any(err.get('Message'), err.get('message'))} request id: {request_id}",
                'data': err,
                'description': f"{self.default_any(err.get('Description'), err.get('description'))}",
                'accessDeniedDetail': self.default_any(err.get('AccessDeniedDetail'), err.get('accessDeniedDetail'))
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
            request_id = self.default_any(err.get('RequestId'), err.get('requestId'))
            if not UtilClient.is_unset(response.headers.get('x-acs-request-id')):
                request_id = response.headers.get('x-acs-request-id')
            err['statusCode'] = response.status_code
            raise TeaException({
                'code': f"{self.default_any(err.get('Code'), err.get('code'))}",
                'message': f"code: {response.status_code}, {self.default_any(err.get('Message'), err.get('message'))} request id: {request_id}",
                'data': err,
                'description': f"{self.default_any(err.get('Description'), err.get('description'))}",
                'accessDeniedDetail': self.default_any(err.get('AccessDeniedDetail'), err.get('accessDeniedDetail'))
            })
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
        signingkey: bytes,
        product: str,
        region: str,
        date: str,
    ) -> str:
        signature = self.get_signature(pathname, method, query, headers, signature_algorithm, payload, signingkey)
        signed_headers = self.get_signed_headers(headers)
        signed_headers_str = ArrayClient.join(signed_headers, ';')
        return f'{signature_algorithm} Credential={ak}/{date}/{region}/{product}/{self._sign_prefix}_request,SignedHeaders={signed_headers_str},Signature={signature}'

    def get_signature(
        self,
        pathname: str,
        method: str,
        query: Dict[str, str],
        headers: Dict[str, str],
        signature_algorithm: str,
        payload: str,
        signingkey: bytes,
    ) -> str:
        canonical_uri = '/'
        if not UtilClient.empty(pathname):
            canonical_uri = pathname
        string_to_sign = ''
        canonicalized_resource = self.build_canonicalized_resource(query)
        canonicalized_headers = self.build_canonicalized_headers(headers)
        signed_headers = self.get_signed_headers(headers)
        signed_headers_str = ArrayClient.join(signed_headers, ';')
        string_to_sign = f'{method}\n{canonical_uri}\n{canonicalized_resource}\n{canonicalized_headers}\n{signed_headers_str}\n{payload}'
        hex = Encoder.hex_encode(Encoder.hash(UtilClient.to_bytes(string_to_sign), signature_algorithm))
        string_to_sign = f'{signature_algorithm}\n{hex}'
        signature = UtilClient.to_bytes('')
        if UtilClient.equal_string(signature_algorithm, self._sha_256):
            signature = Signer.hmac_sha256sign_by_bytes(string_to_sign, signingkey)
        elif UtilClient.equal_string(signature_algorithm, self._sm_3):
            signature = Signer.hmac_sm3sign_by_bytes(string_to_sign, signingkey)
        return Encoder.hex_encode(signature)

    def get_signingkey(
        self,
        signature_algorithm: str,
        secret: str,
        product: str,
        region: str,
        date: str,
    ) -> bytes:
        sc_1 = f'{self._sign_prefix}{secret}'
        sc_2 = UtilClient.to_bytes('')
        if UtilClient.equal_string(signature_algorithm, self._sha_256):
            sc_2 = Signer.hmac_sha256sign(date, sc_1)
        elif UtilClient.equal_string(signature_algorithm, self._sm_3):
            sc_2 = Signer.hmac_sm3sign(date, sc_1)
        sc_3 = UtilClient.to_bytes('')
        if UtilClient.equal_string(signature_algorithm, self._sha_256):
            sc_3 = Signer.hmac_sha256sign_by_bytes(region, sc_2)
        elif UtilClient.equal_string(signature_algorithm, self._sm_3):
            sc_3 = Signer.hmac_sm3sign_by_bytes(region, sc_2)
        sc_4 = UtilClient.to_bytes('')
        if UtilClient.equal_string(signature_algorithm, self._sha_256):
            sc_4 = Signer.hmac_sha256sign_by_bytes(product, sc_3)
        elif UtilClient.equal_string(signature_algorithm, self._sm_3):
            sc_4 = Signer.hmac_sm3sign_by_bytes(product, sc_3)
        hmac = UtilClient.to_bytes('')
        if UtilClient.equal_string(signature_algorithm, self._sha_256):
            hmac = Signer.hmac_sha256sign_by_bytes(f'{self._sign_prefix}_request', sc_4)
        elif UtilClient.equal_string(signature_algorithm, self._sm_3):
            hmac = Signer.hmac_sm3sign_by_bytes(f'{self._sign_prefix}_request', sc_4)
        return hmac

    def get_region(
        self,
        product: str,
        endpoint: str,
        region_id: str,
    ) -> str:
        if not UtilClient.empty(region_id):
            return region_id
        region = 'center'
        if UtilClient.empty(product) or UtilClient.empty(endpoint):
            return region
        strs = StringClient.split(endpoint, ':', None)
        without_port = strs[0]
        pre_region = StringClient.replace(without_port, f'.{self._endpoint_suffix}', '', None)
        nodes = StringClient.split(pre_region, '.', None)
        if UtilClient.equal_number(ArrayClient.size(nodes), 2):
            region = nodes[1]
        return region

    def build_canonicalized_resource(
        self,
        query: Dict[str, str],
    ) -> str:
        canonicalized_resource = ''
        if not UtilClient.is_unset(query):
            query_array = MapClient.key_set(query)
            sorted_query_array = ArrayClient.asc_sort(query_array)
            separator = ''
            for key in sorted_query_array:
                canonicalized_resource = f'{canonicalized_resource}{separator}{Encoder.percent_encode(key)}='
                if not UtilClient.empty(query.get(key)):
                    canonicalized_resource = f'{canonicalized_resource}{Encoder.percent_encode(query.get(key))}'
                separator = '&'
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
        return StringClient.split(tmp, ';', None)
