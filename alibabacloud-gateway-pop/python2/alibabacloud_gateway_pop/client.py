# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import unicode_literals

from alibabacloud_darabonba_encode_util.encoder import Encoder
from Tea.exceptions import TeaException
from alibabacloud_darabonba_signature_util.signer import Signer
from Tea.core import TeaCore
from Tea.converter import TeaConverter

from alibabacloud_gateway_spi.client import Client as SPIClient
from alibabacloud_openapi_util.client import Client as OpenApiUtilClient
from alibabacloud_tea_util.client import Client as UtilClient
from alibabacloud_endpoint_util.client import Client as EndpointUtilClient
from alibabacloud_darabonba_array.client import Client as ArrayClient
from alibabacloud_darabonba_string.client import Client as StringClient
from alibabacloud_darabonba_map.client import Client as MapClient


class Client(SPIClient):
    def __init__(self):
        super(Client, self).__init__()

    def modify_configuration(self, context, attribute_map):
        request = context.request
        config = context.configuration
        config.endpoint = self.get_endpoint(request.product_id, config.region_id, config.endpoint_rule, config.network, config.suffix, config.endpoint_map, config.endpoint)

    def modify_request(self, context, attribute_map):
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

    def modify_response(self, context, attribute_map):
        request = context.request
        config = context.configuration
        response = context.response
        if UtilClient.is_4xx(response.status_code) or UtilClient.is_5xx(response.status_code):
            _res = UtilClient.read_as_json(response.body)
            err = UtilClient.assert_as_map(_res)
            request_id = self.default_any(err.get('RequestId'), err.get('requestId'))
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
            obj = UtilClient.read_as_json(response.body)
            res = UtilClient.assert_as_map(obj)
            response.deserialized_body = res
        elif UtilClient.equal_string(request.body_type, 'array'):
            arr = UtilClient.read_as_json(response.body)
            response.deserialized_body = arr
        else:
            response.deserialized_body = UtilClient.read_as_string(response.body)

    def get_endpoint(self, product_id, region_id, endpoint_rule, network, suffix, endpoint_map, endpoint):
        if not UtilClient.empty(endpoint):
            return endpoint
        if not UtilClient.is_unset(endpoint_map) and not UtilClient.empty(endpoint_map.get(region_id)):
            return endpoint_map.get(region_id)
        return EndpointUtilClient.get_endpoint_rules(product_id, region_id, endpoint_rule, network, suffix)

    def default_any(self, input_value, default_value):
        if UtilClient.is_unset(input_value):
            return default_value
        return input_value

    def get_authorization(self, pathname, method, query, headers, signature_algorithm, payload, ak, secret):
        signature = self.get_signature(pathname, method, query, headers, signature_algorithm, payload, secret)
        return '%s  Credential=%s,SignedHeaders=%s,Signature=%s' % (TeaConverter.to_unicode(signature_algorithm), TeaConverter.to_unicode(ak), TeaConverter.to_unicode(ArrayClient.join(self.get_signed_headers(headers), ';')), TeaConverter.to_unicode(signature))

    def get_signature(self, pathname, method, query, headers, signature_algorithm, payload, secret):
        canonical_uri = '/'
        if not UtilClient.empty(pathname):
            canonical_uri = pathname
        string_to_sign = ''
        canonicalized_resource = self.build_canonicalized_resource(query)
        canonicalized_headers = self.build_canonicalized_headers(headers)
        signed_headers = self.get_signed_headers(headers)
        string_to_sign = '%s\n%s\n%s\n%s\n%s\n%s' % (TeaConverter.to_unicode(method), TeaConverter.to_unicode(canonical_uri), TeaConverter.to_unicode(canonicalized_resource), TeaConverter.to_unicode(canonicalized_headers), TeaConverter.to_unicode(ArrayClient.join(signed_headers, ';')), TeaConverter.to_unicode(payload))
        hex = Encoder.hex_encode(Encoder.hash(UtilClient.to_bytes(string_to_sign), signature_algorithm))
        string_to_sign = '%s\n%s' % (TeaConverter.to_unicode(signature_algorithm), TeaConverter.to_unicode(hex))
        signature = UtilClient.to_bytes('')
        if StringClient.equals(signature_algorithm, 'ACS3-HMAC-SHA256'):
            signature = Signer.hmac_sha256sign(string_to_sign, secret)
        elif StringClient.equals(signature_algorithm, 'ACS3-HMAC-SM3'):
            signature = Signer.hmac_sm3sign(string_to_sign, secret)
        elif StringClient.equals(signature_algorithm, 'ACS3-RSA-SHA256'):
            signature = Signer.sha256with_rsasign(string_to_sign, secret)
        return Encoder.hex_encode(signature)

    def build_canonicalized_resource(self, query):
        canonicalized_resource = ''
        if not UtilClient.is_unset(query):
            query_array = MapClient.key_set(query)
            sorted_query_array = ArrayClient.asc_sort(query_array)
            for key in sorted_query_array:
                canonicalized_resource = '%s&%s' % (TeaConverter.to_unicode(canonicalized_resource), TeaConverter.to_unicode(Encoder.percent_encode(key)))
                if not UtilClient.empty(query.get(key)):
                    canonicalized_resource = '%s=%s' % (TeaConverter.to_unicode(canonicalized_resource), TeaConverter.to_unicode(Encoder.percent_encode(query.get(key))))
        return canonicalized_resource

    def build_canonicalized_headers(self, headers):
        canonicalized_headers = ''
        sorted_headers = self.get_signed_headers(headers)
        for header in sorted_headers:
            canonicalized_headers = '%s%s:%s\n' % (TeaConverter.to_unicode(canonicalized_headers), TeaConverter.to_unicode(header), TeaConverter.to_unicode(StringClient.trim(headers.get(header))))
        return canonicalized_headers

    def get_signed_headers(self, headers):
        headers_array = MapClient.key_set(headers)
        sorted_headers_array = ArrayClient.asc_sort(headers_array)
        tmp = ''
        separator = ''
        for key in sorted_headers_array:
            lower_key = StringClient.to_lower(key)
            if StringClient.has_prefix(lower_key, 'x-acs-') or StringClient.equals(lower_key, 'host') or StringClient.equals(lower_key, 'content-type'):
                if not StringClient.contains(tmp, lower_key):
                    tmp = '%s%s%s' % (TeaConverter.to_unicode(tmp), TeaConverter.to_unicode(separator), TeaConverter.to_unicode(lower_key))
                    separator = ';'
        return StringClient.split(tmp, ';', 0)
