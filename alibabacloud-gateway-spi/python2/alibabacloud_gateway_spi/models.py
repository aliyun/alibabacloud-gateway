# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from Tea.model import TeaModel

from alibabacloud_credentials.client import Client as CredentialClient


class InterceptorContextRequest(TeaModel):
    def __init__(self, headers=None, query=None, body=None, stream=None, host_map=None, pathname=None,
                 product_id=None, action=None, version=None, protocol=None, method=None, auth_type=None, body_type=None,
                 req_body_type=None, style=None, credential=None, signature_version=None, signature_algorithm=None,
                 user_agent=None):
        self.headers = headers  # type: dict[str, str]
        self.query = query  # type: dict[str, str]
        self.body = body  # type: any
        self.stream = stream  # type: READABLE
        self.host_map = host_map  # type: dict[str, str]
        self.pathname = pathname  # type: str
        self.product_id = product_id  # type: str
        self.action = action  # type: str
        self.version = version  # type: str
        self.protocol = protocol  # type: str
        self.method = method  # type: str
        self.auth_type = auth_type  # type: str
        self.body_type = body_type  # type: str
        self.req_body_type = req_body_type  # type: str
        self.style = style  # type: str
        self.credential = credential  # type: CredentialClient
        self.signature_version = signature_version  # type: str
        self.signature_algorithm = signature_algorithm  # type: str
        self.user_agent = user_agent  # type: str

    def validate(self):
        self.validate_required(self.pathname, 'pathname')
        self.validate_required(self.product_id, 'product_id')
        self.validate_required(self.action, 'action')
        self.validate_required(self.version, 'version')
        self.validate_required(self.protocol, 'protocol')
        self.validate_required(self.method, 'method')
        self.validate_required(self.auth_type, 'auth_type')
        self.validate_required(self.body_type, 'body_type')
        self.validate_required(self.req_body_type, 'req_body_type')
        self.validate_required(self.credential, 'credential')
        self.validate_required(self.user_agent, 'user_agent')

    def to_map(self):
        _map = super(InterceptorContextRequest, self).to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.headers is not None:
            result['headers'] = self.headers
        if self.query is not None:
            result['query'] = self.query
        if self.body is not None:
            result['body'] = self.body
        if self.stream is not None:
            result['stream'] = self.stream
        if self.host_map is not None:
            result['hostMap'] = self.host_map
        if self.pathname is not None:
            result['pathname'] = self.pathname
        if self.product_id is not None:
            result['productId'] = self.product_id
        if self.action is not None:
            result['action'] = self.action
        if self.version is not None:
            result['version'] = self.version
        if self.protocol is not None:
            result['protocol'] = self.protocol
        if self.method is not None:
            result['method'] = self.method
        if self.auth_type is not None:
            result['authType'] = self.auth_type
        if self.body_type is not None:
            result['bodyType'] = self.body_type
        if self.req_body_type is not None:
            result['reqBodyType'] = self.req_body_type
        if self.style is not None:
            result['style'] = self.style
        if self.credential is not None:
            result['credential'] = self.credential
        if self.signature_version is not None:
            result['signatureVersion'] = self.signature_version
        if self.signature_algorithm is not None:
            result['signatureAlgorithm'] = self.signature_algorithm
        if self.user_agent is not None:
            result['userAgent'] = self.user_agent
        return result

    def from_map(self, m=None):
        m = m or dict()
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('query') is not None:
            self.query = m.get('query')
        if m.get('body') is not None:
            self.body = m.get('body')
        if m.get('stream') is not None:
            self.stream = m.get('stream')
        if m.get('hostMap') is not None:
            self.host_map = m.get('hostMap')
        if m.get('pathname') is not None:
            self.pathname = m.get('pathname')
        if m.get('productId') is not None:
            self.product_id = m.get('productId')
        if m.get('action') is not None:
            self.action = m.get('action')
        if m.get('version') is not None:
            self.version = m.get('version')
        if m.get('protocol') is not None:
            self.protocol = m.get('protocol')
        if m.get('method') is not None:
            self.method = m.get('method')
        if m.get('authType') is not None:
            self.auth_type = m.get('authType')
        if m.get('bodyType') is not None:
            self.body_type = m.get('bodyType')
        if m.get('reqBodyType') is not None:
            self.req_body_type = m.get('reqBodyType')
        if m.get('style') is not None:
            self.style = m.get('style')
        if m.get('credential') is not None:
            self.credential = m.get('credential')
        if m.get('signatureVersion') is not None:
            self.signature_version = m.get('signatureVersion')
        if m.get('signatureAlgorithm') is not None:
            self.signature_algorithm = m.get('signatureAlgorithm')
        if m.get('userAgent') is not None:
            self.user_agent = m.get('userAgent')
        return self


class InterceptorContextConfiguration(TeaModel):
    def __init__(self, region_id=None, endpoint=None, endpoint_rule=None, endpoint_map=None, endpoint_type=None,
                 network=None, suffix=None):
        self.region_id = region_id  # type: str
        self.endpoint = endpoint  # type: str
        self.endpoint_rule = endpoint_rule  # type: str
        self.endpoint_map = endpoint_map  # type: dict[str, str]
        self.endpoint_type = endpoint_type  # type: str
        self.network = network  # type: str
        self.suffix = suffix  # type: str

    def validate(self):
        self.validate_required(self.region_id, 'region_id')

    def to_map(self):
        _map = super(InterceptorContextConfiguration, self).to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.region_id is not None:
            result['regionId'] = self.region_id
        if self.endpoint is not None:
            result['endpoint'] = self.endpoint
        if self.endpoint_rule is not None:
            result['endpointRule'] = self.endpoint_rule
        if self.endpoint_map is not None:
            result['endpointMap'] = self.endpoint_map
        if self.endpoint_type is not None:
            result['endpointType'] = self.endpoint_type
        if self.network is not None:
            result['network'] = self.network
        if self.suffix is not None:
            result['suffix'] = self.suffix
        return result

    def from_map(self, m=None):
        m = m or dict()
        if m.get('regionId') is not None:
            self.region_id = m.get('regionId')
        if m.get('endpoint') is not None:
            self.endpoint = m.get('endpoint')
        if m.get('endpointRule') is not None:
            self.endpoint_rule = m.get('endpointRule')
        if m.get('endpointMap') is not None:
            self.endpoint_map = m.get('endpointMap')
        if m.get('endpointType') is not None:
            self.endpoint_type = m.get('endpointType')
        if m.get('network') is not None:
            self.network = m.get('network')
        if m.get('suffix') is not None:
            self.suffix = m.get('suffix')
        return self


class InterceptorContextResponse(TeaModel):
    def __init__(self, status_code=None, headers=None, body=None, deserialized_body=None):
        self.status_code = status_code  # type: int
        self.headers = headers  # type: dict[str, str]
        self.body = body  # type: READABLE
        self.deserialized_body = deserialized_body  # type: any

    def validate(self):
        pass

    def to_map(self):
        _map = super(InterceptorContextResponse, self).to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.status_code is not None:
            result['statusCode'] = self.status_code
        if self.headers is not None:
            result['headers'] = self.headers
        if self.body is not None:
            result['body'] = self.body
        if self.deserialized_body is not None:
            result['deserializedBody'] = self.deserialized_body
        return result

    def from_map(self, m=None):
        m = m or dict()
        if m.get('statusCode') is not None:
            self.status_code = m.get('statusCode')
        if m.get('headers') is not None:
            self.headers = m.get('headers')
        if m.get('body') is not None:
            self.body = m.get('body')
        if m.get('deserializedBody') is not None:
            self.deserialized_body = m.get('deserializedBody')
        return self


class InterceptorContext(TeaModel):
    def __init__(self, request=None, configuration=None, response=None):
        self.request = request  # type: InterceptorContextRequest
        self.configuration = configuration  # type: InterceptorContextConfiguration
        self.response = response  # type: InterceptorContextResponse

    def validate(self):
        self.validate_required(self.request, 'request')
        if self.request:
            self.request.validate()
        self.validate_required(self.configuration, 'configuration')
        if self.configuration:
            self.configuration.validate()
        self.validate_required(self.response, 'response')
        if self.response:
            self.response.validate()

    def to_map(self):
        _map = super(InterceptorContext, self).to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.request is not None:
            result['request'] = self.request.to_map()
        if self.configuration is not None:
            result['configuration'] = self.configuration.to_map()
        if self.response is not None:
            result['response'] = self.response.to_map()
        return result

    def from_map(self, m=None):
        m = m or dict()
        if m.get('request') is not None:
            temp_model = InterceptorContextRequest()
            self.request = temp_model.from_map(m['request'])
        if m.get('configuration') is not None:
            temp_model = InterceptorContextConfiguration()
            self.configuration = temp_model.from_map(m['configuration'])
        if m.get('response') is not None:
            temp_model = InterceptorContextResponse()
            self.response = temp_model.from_map(m['response'])
        return self


class AttributeMap(TeaModel):
    def __init__(self, attributes=None, key=None):
        self.attributes = attributes  # type: dict[str, any]
        self.key = key  # type: dict[str, str]

    def validate(self):
        self.validate_required(self.attributes, 'attributes')
        self.validate_required(self.key, 'key')

    def to_map(self):
        _map = super(AttributeMap, self).to_map()
        if _map is not None:
            return _map

        result = dict()
        if self.attributes is not None:
            result['attributes'] = self.attributes
        if self.key is not None:
            result['key'] = self.key
        return result

    def from_map(self, m=None):
        m = m or dict()
        if m.get('attributes') is not None:
            self.attributes = m.get('attributes')
        if m.get('key') is not None:
            self.key = m.get('key')
        return self


