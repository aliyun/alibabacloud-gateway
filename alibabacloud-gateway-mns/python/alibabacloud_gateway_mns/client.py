# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from alibabacloud_darabonba_encode_util.encoder import Encoder
from alibabacloud_darabonba_signature_util.signer import Signer
from Tea.exceptions import TeaException
from Tea.core import TeaCore
from typing import Any, Dict

from alibabacloud_gateway_spi.client import Client as SPIClient
from alibabacloud_gateway_spi import models as spi_models
from alibabacloud_tea_util.client import Client as UtilClient
from alibabacloud_darabonba_string.client import Client as StringClient
from alibabacloud_tea_xml.client import Client as XMLClient
from alibabacloud_darabonba_array.client import Client as ArrayClient
from alibabacloud_darabonba_map.client import Client as MapClient
from alibabacloud_openapi_util.client import Client as OpenApiUtilClient


class Client(SPIClient):
    _sign_prefix: str = None
    _sign_suffix: str = None
    _auth_prefix: str = None

    def __init__(self):
        super().__init__()
        undefined._sign_prefix = 'aliyun_v4'
        undefined._sign_suffix = 'aliyun_v4_request'
        undefined._auth_prefix = 'MNS4-HMAC-SHA256'

    def modify_configuration(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        config = context.configuration
        config.endpoint = self.get_endpoint(config.region_id, config.endpoint)

    async def modify_configuration_async(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        config = context.configuration
        config.endpoint = await self.get_endpoint_async(config.region_id, config.endpoint)

    def modify_request(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        request = context.request
        config = context.configuration
        signature_version = UtilClient.default_string(request.signature_version, 'v2')
        if not UtilClient.is_unset(request.body):
            if StringClient.equals(request.req_body_type, 'xml'):
                req_body_map = UtilClient.assert_as_map(request.body)
                xml_str = XMLClient.to_xml(req_body_map)
                request.stream = xml_str
                request.headers['content-type'] = 'text/xml;charset=UTF-8'
                request.headers['content-md5'] = Encoder.base_64encode_to_string(Signer.md5sign(xml_str))
            elif StringClient.equals(request.req_body_type, 'json') or StringClient.equals(request.req_body_type, 'formData'):
                body_str = UtilClient.to_jsonstring(request.body)
                request.stream = body_str
                request.headers['content-type'] = 'application/json'
                request.headers['content-md5'] = Encoder.base_64encode_to_string(Signer.md5sign(body_str))
            elif StringClient.equals(request.req_body_type, 'byte') or StringClient.equals(request.req_body_type, 'binary'):
                body_bytes = UtilClient.assert_as_bytes(request.body)
                request.stream = body_bytes
                request.headers['content-md5'] = Encoder.base_64encode_to_string(Signer.md5sign_for_bytes(body_bytes))
        self.build_request(context)
        request.headers = TeaCore.merge({
            'host': config.endpoint,
            'x-mns-version': request.version,
            'user-agent': request.user_agent,
            'accept': 'application/json'
        }, request.headers)
        if not UtilClient.equal_string(request.auth_type, 'Anonymous'):
            credential = request.credential
            if UtilClient.is_unset(credential):
                raise TeaException({
                    'code': 'ParameterMissing',
                    'message': "'config.credential' can not be unset"
                })
            credential_model = credential.get_credential()
            auth_type = credential_model.type
            if UtilClient.equal_string(auth_type, 'bearer'):
                bearer_token = credential_model.bearer_token
                request.headers['x-acs-bearer-token'] = bearer_token
                request.headers['x-acs-signature-type'] = 'BEARERTOKEN'
                request.headers['Authorization'] = f'Bearer {bearer_token}'
            else:
                access_key_id = credential_model.access_key_id
                access_key_secret = credential_model.access_key_secret
                security_token = credential_model.security_token
                if not UtilClient.empty(security_token):
                    request.headers['security-token'] = security_token
                request.headers['date'] = UtilClient.get_date_utcstring()
                if StringClient.equals(signature_version, 'v4'):
                    date = self.get_date_iso8601()
                    request.headers['authorization'] = self.get_authorization_v4(context, date, access_key_id, access_key_secret)
                else:
                    request.headers['authorization'] = self.get_authorization_v2(request.pathname, request.method, request.headers, access_key_id, access_key_secret)

    async def modify_request_async(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        request = context.request
        config = context.configuration
        signature_version = UtilClient.default_string(request.signature_version, 'v2')
        if not UtilClient.is_unset(request.body):
            if StringClient.equals(request.req_body_type, 'xml'):
                req_body_map = UtilClient.assert_as_map(request.body)
                xml_str = XMLClient.to_xml(req_body_map)
                request.stream = xml_str
                request.headers['content-type'] = 'text/xml;charset=UTF-8'
                request.headers['content-md5'] = Encoder.base_64encode_to_string(Signer.md5sign(xml_str))
            elif StringClient.equals(request.req_body_type, 'json') or StringClient.equals(request.req_body_type, 'formData'):
                body_str = UtilClient.to_jsonstring(request.body)
                request.stream = body_str
                request.headers['content-type'] = 'application/json'
                request.headers['content-md5'] = Encoder.base_64encode_to_string(Signer.md5sign(body_str))
            elif StringClient.equals(request.req_body_type, 'byte') or StringClient.equals(request.req_body_type, 'binary'):
                body_bytes = UtilClient.assert_as_bytes(request.body)
                request.stream = body_bytes
                request.headers['content-md5'] = Encoder.base_64encode_to_string(Signer.md5sign_for_bytes(body_bytes))
        await self.build_request_async(context)
        request.headers = TeaCore.merge({
            'host': config.endpoint,
            'x-mns-version': request.version,
            'user-agent': request.user_agent,
            'accept': 'application/json'
        }, request.headers)
        if not UtilClient.equal_string(request.auth_type, 'Anonymous'):
            credential = request.credential
            if UtilClient.is_unset(credential):
                raise TeaException({
                    'code': 'ParameterMissing',
                    'message': "'config.credential' can not be unset"
                })
            credential_model = await credential.get_credential_async()
            auth_type = credential_model.type
            if UtilClient.equal_string(auth_type, 'bearer'):
                bearer_token = credential_model.bearer_token
                request.headers['x-acs-bearer-token'] = bearer_token
                request.headers['x-acs-signature-type'] = 'BEARERTOKEN'
                request.headers['Authorization'] = f'Bearer {bearer_token}'
            else:
                access_key_id = credential_model.access_key_id
                access_key_secret = credential_model.access_key_secret
                security_token = credential_model.security_token
                if not UtilClient.empty(security_token):
                    request.headers['security-token'] = security_token
                request.headers['date'] = UtilClient.get_date_utcstring()
                if StringClient.equals(signature_version, 'v4'):
                    date = await self.get_date_iso8601_async()
                    request.headers['authorization'] = await self.get_authorization_v4_async(context, date, access_key_id, access_key_secret)
                else:
                    request.headers['authorization'] = await self.get_authorization_v2_async(request.pathname, request.method, request.headers, access_key_id, access_key_secret)

    def modify_response(
        self,
        context: spi_models.InterceptorContext,
        attribute_map: spi_models.AttributeMap,
    ) -> None:
        request = context.request
        response = context.response
        if UtilClient.is_4xx(response.status_code) or UtilClient.is_5xx(response.status_code):
            err = {}
            if not UtilClient.is_unset(response.headers.get('content-type')) and StringClient.contains(response.headers.get('content-type'), 'text/xml'):
                _str = UtilClient.read_as_string(response.body)
                resp_map = XMLClient.parse_xml(_str, None)
                err = UtilClient.assert_as_map(resp_map.get('Error'))
            else:
                _res = UtilClient.read_as_json(response.body)
                err = UtilClient.assert_as_map(_res)
            request_id = self.default_any(err.get('RequestId'), err.get('requestId'))
            if not UtilClient.is_unset(response.headers.get('x-mns-request-id')):
                request_id = response.headers.get('x-mns-request-id')
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
            err = {}
            if not UtilClient.is_unset(response.headers.get('content-type')) and StringClient.contains(response.headers.get('content-type'), 'text/xml'):
                _str = await UtilClient.read_as_string_async(response.body)
                resp_map = XMLClient.parse_xml(_str, None)
                err = UtilClient.assert_as_map(resp_map.get('Error'))
            else:
                _res = await UtilClient.read_as_json_async(response.body)
                err = UtilClient.assert_as_map(_res)
            request_id = self.default_any(err.get('RequestId'), err.get('requestId'))
            if not UtilClient.is_unset(response.headers.get('x-mns-request-id')):
                request_id = response.headers.get('x-mns-request-id')
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
        region_id: str,
        endpoint: str,
    ) -> str:
        if not UtilClient.empty(endpoint):
            return endpoint
        if UtilClient.empty(region_id):
            region_id = 'cn-hangzhou'
        return f'{region_id}.mns.aliyuncs.com'

    async def get_endpoint_async(
        self,
        region_id: str,
        endpoint: str,
    ) -> str:
        if not UtilClient.empty(endpoint):
            return endpoint
        if UtilClient.empty(region_id):
            region_id = 'cn-hangzhou'
        return f'{region_id}.mns.aliyuncs.com'

    def default_any(
        self,
        input_value: Any,
        default_value: Any,
    ) -> Any:
        if UtilClient.is_unset(input_value):
            return default_value
        return input_value

    def default_string(
        self,
        input_value: str,
        default_value: str,
    ) -> str:
        if UtilClient.is_unset(input_value) or UtilClient.empty(input_value):
            return default_value
        return input_value

    def base_64encode(
        self,
        input: str,
    ) -> str:
        if UtilClient.is_unset(input):
            return ''
        return Encoder.base_64encode_to_string(UtilClient.to_bytes(input))

    async def base_64encode_async(
        self,
        input: str,
    ) -> str:
        if UtilClient.is_unset(input):
            return ''
        return Encoder.base_64encode_to_string(UtilClient.to_bytes(input))

    def base_64decode(
        self,
        input: str,
    ) -> str:
        if UtilClient.is_unset(input):
            return ''
        return UtilClient.to_string(Encoder.base_64decode(input))

    async def base_64decode_async(
        self,
        input: str,
    ) -> str:
        if UtilClient.is_unset(input):
            return ''
        return UtilClient.to_string(Encoder.base_64decode(input))

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

    def get_authorization_v2(
        self,
        pathname: str,
        method: str,
        headers: Dict[str, str],
        ak: str,
        secret: str,
    ) -> str:
        sign = self.get_signature_v2(pathname, method, headers, secret)
        return f'MNS {ak}:{sign}'

    async def get_authorization_v2_async(
        self,
        pathname: str,
        method: str,
        headers: Dict[str, str],
        ak: str,
        secret: str,
    ) -> str:
        sign = await self.get_signature_v2_async(pathname, method, headers, secret)
        return f'MNS {ak}:{sign}'

    def get_signature_v2(
        self,
        pathname: str,
        method: str,
        headers: Dict[str, str],
        secret: str,
    ) -> str:
        canonicalized_resource = self.default_string(pathname, '/')
        canonicalized_headers = self.build_canonicalized_headers_v2(headers)
        string_to_sign = f'{method}\n{canonicalized_headers}{canonicalized_resource}'
        return Encoder.base_64encode_to_string(Signer.hmac_sha1sign(string_to_sign, secret))

    async def get_signature_v2_async(
        self,
        pathname: str,
        method: str,
        headers: Dict[str, str],
        secret: str,
    ) -> str:
        canonicalized_resource = self.default_string(pathname, '/')
        canonicalized_headers = await self.build_canonicalized_headers_v2_async(headers)
        string_to_sign = f'{method}\n{canonicalized_headers}{canonicalized_resource}'
        return Encoder.base_64encode_to_string(Signer.hmac_sha1sign(string_to_sign, secret))

    def build_canonicalized_headers_v2(
        self,
        headers: Dict[str, str],
    ) -> str:
        content_md_5 = self.default_string(headers.get('content-md5'), '')
        content_type = self.default_string(headers.get('content-type'), '')
        date = self.default_string(headers.get('date'), '')
        canonicalized_headers = f'{content_md_5}\n{content_type}\n{date}\n'
        mns_headers = {}
        for header in MapClient.key_set(headers):
            lower_header = StringClient.to_lower(header)
            if StringClient.has_prefix(lower_header, 'x-mns-'):
                mns_headers[lower_header] = headers.get(header)
        for header in ArrayClient.asc_sort(MapClient.key_set(mns_headers)):
            canonicalized_headers = f'{canonicalized_headers}{header}:{mns_headers.get(header)}\n'
        return canonicalized_headers

    async def build_canonicalized_headers_v2_async(
        self,
        headers: Dict[str, str],
    ) -> str:
        content_md_5 = self.default_string(headers.get('content-md5'), '')
        content_type = self.default_string(headers.get('content-type'), '')
        date = self.default_string(headers.get('date'), '')
        canonicalized_headers = f'{content_md_5}\n{content_type}\n{date}\n'
        mns_headers = {}
        for header in MapClient.key_set(headers):
            lower_header = StringClient.to_lower(header)
            if StringClient.has_prefix(lower_header, 'x-mns-'):
                mns_headers[lower_header] = headers.get(header)
        for header in ArrayClient.asc_sort(MapClient.key_set(mns_headers)):
            canonicalized_headers = f'{canonicalized_headers}{header}:{mns_headers.get(header)}\n'
        return canonicalized_headers

    def get_authorization_v4(
        self,
        context: spi_models.InterceptorContext,
        date: str,
        access_key_id: str,
        access_key_secret: str,
    ) -> str:
        region = self.get_region(context)
        date_short = StringClient.sub_string(date, 0, 8)
        date_short = StringClient.replace(date_short, 'T', '', None)
        scope = f'{date_short}/{region}/mns/{self._sign_suffix}'
        signingkey = self.get_signingkey_v4(access_key_secret, region, date_short)
        signature = self.get_signature_v4(context, date, scope, signingkey)
        return f'{self._auth_prefix} Credential={access_key_id}/{scope},Signature={signature}'

    async def get_authorization_v4_async(
        self,
        context: spi_models.InterceptorContext,
        date: str,
        access_key_id: str,
        access_key_secret: str,
    ) -> str:
        region = await self.get_region_async(context)
        date_short = StringClient.sub_string(date, 0, 8)
        date_short = StringClient.replace(date_short, 'T', '', None)
        scope = f'{date_short}/{region}/mns/{self._sign_suffix}'
        signingkey = await self.get_signingkey_v4_async(access_key_secret, region, date_short)
        signature = await self.get_signature_v4_async(context, date, scope, signingkey)
        return f'{self._auth_prefix} Credential={access_key_id}/{scope},Signature={signature}'

    def get_signingkey_v4(
        self,
        secret: str,
        region: str,
        date: str,
    ) -> bytes:
        sc_1 = f'{self._sign_prefix}{secret}'
        sc_2 = Signer.hmac_sha256sign(date, sc_1)
        sc_3 = Signer.hmac_sha256sign_by_bytes(region, sc_2)
        sc_4 = Signer.hmac_sha256sign_by_bytes('mns', sc_3)
        return Signer.hmac_sha256sign_by_bytes(self._sign_suffix, sc_4)

    async def get_signingkey_v4_async(
        self,
        secret: str,
        region: str,
        date: str,
    ) -> bytes:
        sc_1 = f'{self._sign_prefix}{secret}'
        sc_2 = Signer.hmac_sha256sign(date, sc_1)
        sc_3 = Signer.hmac_sha256sign_by_bytes(region, sc_2)
        sc_4 = Signer.hmac_sha256sign_by_bytes('mns', sc_3)
        return Signer.hmac_sha256sign_by_bytes(self._sign_suffix, sc_4)

    def get_signature_v4(
        self,
        context: spi_models.InterceptorContext,
        date: str,
        scope: str,
        signingkey: bytes,
    ) -> str:
        request = context.request
        canonical_string = self.build_canonical_request_v4(request.pathname, request.method, request.query, request.headers)
        string_to_sign = f'{self._auth_prefix}\n{date}\n{scope}\n{canonical_string}'
        signature = Signer.hmac_sha256sign_by_bytes(string_to_sign, signingkey)
        return Encoder.hex_encode(signature)

    async def get_signature_v4_async(
        self,
        context: spi_models.InterceptorContext,
        date: str,
        scope: str,
        signingkey: bytes,
    ) -> str:
        request = context.request
        canonical_string = await self.build_canonical_request_v4_async(request.pathname, request.method, request.query, request.headers)
        string_to_sign = f'{self._auth_prefix}\n{date}\n{scope}\n{canonical_string}'
        signature = Signer.hmac_sha256sign_by_bytes(string_to_sign, signingkey)
        return Encoder.hex_encode(signature)

    def build_canonical_request_v4(
        self,
        pathname: str,
        method: str,
        query: Dict[str, str],
        headers: Dict[str, str],
    ) -> str:
        canonical_uri = '/'
        if not UtilClient.empty(pathname):
            canonical_uri = pathname
            if not StringClient.has_prefix(canonical_uri, '/'):
                canonical_uri = f'/{canonical_uri}'
        suffix = ''
        if not StringClient.equals(canonical_uri, '/') and StringClient.has_suffix(canonical_uri, '/'):
            suffix = '/'
        canonical_uri = f'{OpenApiUtilClient.get_encode_path(canonical_uri)}{suffix}'
        resources = self.build_canonicalized_resource_v4(query)
        canonical_headers = self.build_canonicalized_headers_v4(headers)
        return f'{method}\n{canonical_uri}\n{resources}\n{canonical_headers}\n'

    async def build_canonical_request_v4_async(
        self,
        pathname: str,
        method: str,
        query: Dict[str, str],
        headers: Dict[str, str],
    ) -> str:
        canonical_uri = '/'
        if not UtilClient.empty(pathname):
            canonical_uri = pathname
            if not StringClient.has_prefix(canonical_uri, '/'):
                canonical_uri = f'/{canonical_uri}'
        suffix = ''
        if not StringClient.equals(canonical_uri, '/') and StringClient.has_suffix(canonical_uri, '/'):
            suffix = '/'
        canonical_uri = f'{OpenApiUtilClient.get_encode_path(canonical_uri)}{suffix}'
        resources = await self.build_canonicalized_resource_v4_async(query)
        canonical_headers = await self.build_canonicalized_headers_v4_async(headers)
        return f'{method}\n{canonical_uri}\n{resources}\n{canonical_headers}\n'

    def build_canonicalized_resource_v4(
        self,
        query: Dict[str, str],
    ) -> str:
        canonicalized_resource = ''
        if not UtilClient.is_unset(query):
            query_map = {}
            for key in MapClient.key_set(query):
                encoded_key = self.percent_encode_mns(StringClient.to_lower(StringClient.trim(key)))
                encoded_value = ''
                if not UtilClient.is_unset(query.get(key)) and not UtilClient.empty(query.get(key)):
                    encoded_value = self.percent_encode_mns(StringClient.trim(query.get(key)))
                query_map[encoded_key] = encoded_value
            query_array = MapClient.key_set(query_map)
            sorted_query_array = ArrayClient.asc_sort(query_array)
            separator = ''
            for encoded_key in sorted_query_array:
                canonicalized_resource = f'{canonicalized_resource}{separator}{encoded_key}'
                encoded_value = query_map.get(encoded_key)
                if not UtilClient.empty(encoded_value):
                    canonicalized_resource = f'{canonicalized_resource}={encoded_value}'
                separator = '&'
        return canonicalized_resource

    async def build_canonicalized_resource_v4_async(
        self,
        query: Dict[str, str],
    ) -> str:
        canonicalized_resource = ''
        if not UtilClient.is_unset(query):
            query_map = {}
            for key in MapClient.key_set(query):
                encoded_key = self.percent_encode_mns(StringClient.to_lower(StringClient.trim(key)))
                encoded_value = ''
                if not UtilClient.is_unset(query.get(key)) and not UtilClient.empty(query.get(key)):
                    encoded_value = self.percent_encode_mns(StringClient.trim(query.get(key)))
                query_map[encoded_key] = encoded_value
            query_array = MapClient.key_set(query_map)
            sorted_query_array = ArrayClient.asc_sort(query_array)
            separator = ''
            for encoded_key in sorted_query_array:
                canonicalized_resource = f'{canonicalized_resource}{separator}{encoded_key}'
                encoded_value = query_map.get(encoded_key)
                if not UtilClient.empty(encoded_value):
                    canonicalized_resource = f'{canonicalized_resource}={encoded_value}'
                separator = '&'
        return canonicalized_resource

    def build_canonicalized_headers_v4(
        self,
        headers: Dict[str, str],
    ) -> str:
        signed_headers = {}
        for key in MapClient.key_set(headers):
            lower_key = StringClient.to_lower(key)
            if self.has_signed_header_v4(lower_key):
                signed_headers[lower_key] = StringClient.trim(headers.get(key))
        canonicalized_headers = ''
        for lower_key in ArrayClient.asc_sort(MapClient.key_set(signed_headers)):
            canonicalized_headers = f'{canonicalized_headers}{lower_key}:{signed_headers.get(lower_key)}\n'
        return canonicalized_headers

    async def build_canonicalized_headers_v4_async(
        self,
        headers: Dict[str, str],
    ) -> str:
        signed_headers = {}
        for key in MapClient.key_set(headers):
            lower_key = StringClient.to_lower(key)
            if self.has_signed_header_v4(lower_key):
                signed_headers[lower_key] = StringClient.trim(headers.get(key))
        canonicalized_headers = ''
        for lower_key in ArrayClient.asc_sort(MapClient.key_set(signed_headers)):
            canonicalized_headers = f'{canonicalized_headers}{lower_key}:{signed_headers.get(lower_key)}\n'
        return canonicalized_headers

    def has_signed_header_v4(
        self,
        header: str,
    ) -> bool:
        if StringClient.equals(header, 'content-type') or StringClient.equals(header, 'content-md5'):
            return True
        return StringClient.has_prefix(header, 'x-mns-')

    def percent_encode_mns(
        self,
        value: str,
    ) -> str:
        encoded = Encoder.percent_encode(value)
        return StringClient.replace(encoded, '+', '%20', None)

    def get_region(
        self,
        context: spi_models.InterceptorContext,
    ) -> str:
        config = context.configuration
        if not UtilClient.is_unset(config.region_id) and not UtilClient.empty(config.region_id):
            return config.region_id
        region = StringClient.replace(config.endpoint, '.mns.aliyuncs.com', '', None)
        if StringClient.equals(region, config.endpoint):
            raise TeaException({
                'code': 'ClientConfigError',
                'message': 'The regionId configuration of mns client is missing.'
            })
        return region

    async def get_region_async(
        self,
        context: spi_models.InterceptorContext,
    ) -> str:
        config = context.configuration
        if not UtilClient.is_unset(config.region_id) and not UtilClient.empty(config.region_id):
            return config.region_id
        region = StringClient.replace(config.endpoint, '.mns.aliyuncs.com', '', None)
        if StringClient.equals(region, config.endpoint):
            raise TeaException({
                'code': 'ClientConfigError',
                'message': 'The regionId configuration of mns client is missing.'
            })
        return region

    def get_date_iso8601(self) -> str:
        date = OpenApiUtilClient.get_timestamp()
        date = StringClient.replace(date, '-', '', None)
        return StringClient.replace(date, ':', '', None)

    async def get_date_iso8601_async(self) -> str:
        date = OpenApiUtilClient.get_timestamp()
        date = StringClient.replace(date, '-', '', None)
        return StringClient.replace(date, ':', '', None)
