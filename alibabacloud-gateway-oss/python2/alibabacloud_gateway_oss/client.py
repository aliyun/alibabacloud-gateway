# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import unicode_literals

from alibabacloud_darabonba_encode_util.encoder import Encoder
from alibabacloud_darabonba_signature_util.signer import Signer
from Tea.exceptions import TeaException
from Tea.converter import TeaConverter
from Tea.core import TeaCore

from alibabacloud_gateway_spi.client import Client as SPIClient
from alibabacloud_tea_util.client import Client as UtilClient
from alibabacloud_darabonba_map.client import Client as MapClient
from alibabacloud_darabonba_string.client import Client as StringClient
from alibabacloud_tea_xml.client import Client as XMLClient
from alibabacloud_openapi_util.client import Client as OpenApiUtilClient
from alibabacloud_oss_util.client import Client as OSSUtilClient
from alibabacloud_darabonba_array.client import Client as ArrayClient


class Client(SPIClient):
    _default_signed_params = None  # type: list[str]
    _except_signed_params = None  # type: list[str]

    def __init__(self):
        super(Client, self).__init__()
        self._default_signed_params = [
            'response-content-type',
            'response-content-language',
            'response-cache-control',
            'logging',
            'response-content-encoding',
            'acl',
            'uploadId',
            'uploads',
            'partNumber',
            'group',
            'link',
            'delete',
            'website',
            'location',
            'objectInfo',
            'objectMeta',
            'response-expires',
            'response-content-disposition',
            'cors',
            'lifecycle',
            'restore',
            'qos',
            'referer',
            'stat',
            'bucketInfo',
            'append',
            'position',
            'security-token',
            'live',
            'comp',
            'status',
            'vod',
            'startTime',
            'endTime',
            'x-oss-process',
            'symlink',
            'callback',
            'callback-var',
            'tagging',
            'encryption',
            'versions',
            'versioning',
            'versionId',
            'policy',
            'requestPayment',
            'x-oss-traffic-limit',
            'qosInfo',
            'asyncFetch',
            'x-oss-request-payer',
            'sequential',
            'inventory',
            'inventoryId',
            'continuation-token',
            'callback',
            'callback-var',
            'worm',
            'wormId',
            'wormExtend',
            'replication',
            'replicationLocation',
            'replicationProgress',
            'transferAcceleration',
            'cname',
            'metaQuery',
            'x-oss-ac-source-ip',
            'x-oss-ac-subnet-mask',
            'x-oss-ac-vpc-id',
            'x-oss-ac-forward-allow',
            'resourceGroup',
            'style',
            'styleName',
            'x-oss-async-process',
            'rtc',
            'accessPoint',
            'accessPointPolicy',
            'httpsConfig',
            'regionsV2',
            'publicAccessBlock',
            'policyStatus',
            'redundancyTransition',
            'redundancyType',
            'redundancyProgress',
            'dataAccelerator',
            'verbose',
            'accessPointForObjectProcess',
            'accessPointConfigForObjectProcess',
            'accessPointPolicyForObjectProcess',
            'bucketArchiveDirectRead',
            'responseHeader',
            'userDefinedLogFieldsConfig'
        ]
        self._except_signed_params = [
            'list-type',
            'regions'
        ]

    def modify_configuration(self, context, attribute_map):
        config = context.configuration
        config.endpoint = self.get_endpoint(config.region_id, config.network, config.endpoint)

    def modify_request(self, context, attribute_map):
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
                new_key = 'x-oss-meta-%s' % TeaConverter.to_unicode(key)
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
        signature_version = UtilClient.default_string(request.signature_version, 'v1')
        request.headers['authorization'] = self.get_authorization(signature_version, bucket_name, request.pathname, request.method, request.query, request.headers, access_key_id, access_key_secret, region_id)

    def modify_response(self, context, attribute_map):
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
                        'requestId': '%s' % TeaConverter.to_unicode(request_id),
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
                    result = XMLClient.parse_xml(body_str, None)
                    list = MapClient.key_set(result)
                    if UtilClient.equal_number(ArrayClient.size(list), 1):
                        tmp = list[0]
                        try:
                            response.deserialized_body = UtilClient.assert_as_map(result.get(tmp))
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

    def get_endpoint(self, region_id, network, endpoint):
        if not UtilClient.empty(endpoint):
            return endpoint
        if UtilClient.empty(region_id):
            region_id = 'cn-hangzhou'
        if not UtilClient.empty(network):
            if StringClient.contains(network, 'internal'):
                return 'oss-%s-internal.aliyuncs.com' % TeaConverter.to_unicode(region_id)
            elif StringClient.contains(network, 'ipv6'):
                return '%soss.aliyuncs.com' % TeaConverter.to_unicode(region_id)
            elif StringClient.contains(network, 'accelerate'):
                return 'oss-%s.aliyuncs.com' % TeaConverter.to_unicode(network)
        return 'oss-%s.aliyuncs.com' % TeaConverter.to_unicode(region_id)

    def get_host(self, endpoint_type, bucket_name, endpoint):
        if UtilClient.empty(bucket_name):
            return endpoint
        host = '%s.%s' % (TeaConverter.to_unicode(bucket_name), TeaConverter.to_unicode(endpoint))
        if not UtilClient.empty(endpoint_type):
            if StringClient.equals(endpoint_type, 'ip'):
                host = '%s/%s' % (TeaConverter.to_unicode(endpoint), TeaConverter.to_unicode(bucket_name))
            elif StringClient.equals(endpoint_type, 'cname'):
                host = endpoint
        return host

    def get_authorization(self, signature_version, bucket_name, pathname, method, query, headers, ak, secret, region_id):
        sign = ''
        if not UtilClient.is_unset(signature_version):
            if StringClient.equals(signature_version, 'v1'):
                sign = self.get_signature_v1(bucket_name, pathname, method, query, headers, secret)
                return 'OSS %s:%s' % (TeaConverter.to_unicode(ak), TeaConverter.to_unicode(sign))
            if StringClient.equals(signature_version, 'v2'):
                sign = self.get_signature_v2(bucket_name, pathname, method, query, headers, secret)
                return 'OSS2 AccessKeyId:%s,Signature:%s' % (TeaConverter.to_unicode(ak), TeaConverter.to_unicode(sign))
        date_time = OpenApiUtilClient.get_timestamp()
        date_time = StringClient.replace(date_time, '-', '', None)
        date_time = StringClient.replace(date_time, ':', '', None)
        headers['x-oss-date'] = date_time
        headers['x-oss-content-sha256'] = 'UNSIGNED-PAYLOAD'
        only_date = StringClient.sub_string(date_time, 0, 8)
        cred = '%s/%s/%s/oss/aliyun_v4_request' % (TeaConverter.to_unicode(ak), TeaConverter.to_unicode(only_date), TeaConverter.to_unicode(region_id))
        sign = self.get_signature_v4(bucket_name, pathname, method, query, headers, only_date, region_id, secret)
        return 'OSS4-HMAC-SHA256 Credential=%s, Signature=%s' % (TeaConverter.to_unicode(cred), TeaConverter.to_unicode(sign))

    def get_sign_key(self, secret, only_date, region_id):
        temp = 'aliyun_v4%s' % TeaConverter.to_unicode(secret)
        res = Signer.hmac_sha256sign(only_date, temp)
        res = Signer.hmac_sha256sign_by_bytes(region_id, res)
        res = Signer.hmac_sha256sign_by_bytes('oss', res)
        res = Signer.hmac_sha256sign_by_bytes('aliyun_v4_request', res)
        return res

    def get_signature_v4(self, bucket_name, pathname, method, query, headers, only_date, region_id, secret):
        signingkey = self.get_sign_key(secret, only_date, region_id)
        object_name = '/'
        query_map = {}
        if not UtilClient.empty(pathname):
            paths = StringClient.split(pathname, '?', 2)
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
            canonicalized_uri = '/%s%s' % (TeaConverter.to_unicode(bucket_name), TeaConverter.to_unicode(object_name))
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
        canonical_request = '%s\n%s\n%s\n%s\n\n%s' % (TeaConverter.to_unicode(method), TeaConverter.to_unicode(canonicalized_uri), TeaConverter.to_unicode(canonicalized_query_string), TeaConverter.to_unicode(canonicalized_headers), TeaConverter.to_unicode(payload))
        hex = Encoder.hex_encode(Encoder.hash(UtilClient.to_bytes(canonical_request), 'ACS4-HMAC-SHA256'))
        scope = '%s/%s/oss/aliyun_v4_request' % (TeaConverter.to_unicode(only_date), TeaConverter.to_unicode(region_id))
        string_to_sign = 'OSS4-HMAC-SHA256\n%s\n%s\n%s' % (TeaConverter.to_unicode(headers.get('x-oss-date')), TeaConverter.to_unicode(scope), TeaConverter.to_unicode(hex))
        signature = Signer.hmac_sha256sign_by_bytes(string_to_sign, signingkey)
        return Encoder.hex_encode(signature)

    def build_canonicalized_query_string_v4(self, query_map):
        canonicalized_query_string = ''
        if not UtilClient.is_unset(query_map):
            query_array = MapClient.key_set(query_map)
            sorted_query_array = ArrayClient.asc_sort(query_array)
            separator = ''
            for key in sorted_query_array:
                canonicalized_query_string = '%s%s%s' % (TeaConverter.to_unicode(canonicalized_query_string), TeaConverter.to_unicode(separator), TeaConverter.to_unicode(key))
                if not UtilClient.empty(query_map.get(key)):
                    canonicalized_query_string = '%s=%s' % (TeaConverter.to_unicode(canonicalized_query_string), TeaConverter.to_unicode(query_map.get(key)))
                separator = '&'
        return canonicalized_query_string

    def build_canonicalized_headers_v4(self, headers):
        canonicalized_headers = ''
        headers_array = MapClient.key_set(headers)
        sorted_headers_array = ArrayClient.asc_sort(headers_array)
        for key in sorted_headers_array:
            lower_key = StringClient.to_lower(key)
            if StringClient.has_prefix(lower_key, 'x-oss-') or StringClient.equals(lower_key, 'content-type') or StringClient.equals(lower_key, 'content-md5'):
                canonicalized_headers = '%s%s:%s\n' % (TeaConverter.to_unicode(canonicalized_headers), TeaConverter.to_unicode(key), TeaConverter.to_unicode(StringClient.trim(headers.get(key))))
        return canonicalized_headers

    def get_signature_v1(self, bucket_name, pathname, method, query, headers, secret):
        resource = ''
        string_to_sign = ''
        if not UtilClient.empty(bucket_name):
            resource = '/%s' % TeaConverter.to_unicode(bucket_name)
        resource = '%s%s' % (TeaConverter.to_unicode(resource), TeaConverter.to_unicode(pathname))
        canonicalized_resource = self.build_canonicalized_resource(resource, query)
        canonicalized_headers = self.build_canonicalized_headers(headers)
        string_to_sign = '%s\n%s%s' % (TeaConverter.to_unicode(method), TeaConverter.to_unicode(canonicalized_headers), TeaConverter.to_unicode(canonicalized_resource))
        return Encoder.base_64encode_to_string(Signer.hmac_sha1sign(string_to_sign, secret))

    def build_canonicalized_resource(self, pathname, query):
        sub_resources_map = {}
        canonicalized_resource = pathname
        if not UtilClient.empty(pathname):
            paths = StringClient.split(pathname, '?', 2)
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
                canonicalized_resource = '%s%s%s' % (TeaConverter.to_unicode(canonicalized_resource), TeaConverter.to_unicode(separator), TeaConverter.to_unicode(param_name))
                if not UtilClient.is_unset(query) and not UtilClient.is_unset(query.get(param_name)):
                    canonicalized_resource = '%s=%s' % (TeaConverter.to_unicode(canonicalized_resource), TeaConverter.to_unicode(query.get(param_name)))
                elif not UtilClient.is_unset(sub_resources_map.get(param_name)):
                    canonicalized_resource = '%s=%s' % (TeaConverter.to_unicode(canonicalized_resource), TeaConverter.to_unicode(sub_resources_map.get(param_name)))
            elif ArrayClient.contains(sub_resources_array, param_name):
                canonicalized_resource = '%s%s%s' % (TeaConverter.to_unicode(canonicalized_resource), TeaConverter.to_unicode(separator), TeaConverter.to_unicode(param_name))
                if not UtilClient.is_unset(sub_resources_map.get(param_name)):
                    canonicalized_resource = '%s=%s' % (TeaConverter.to_unicode(canonicalized_resource), TeaConverter.to_unicode(sub_resources_map.get(param_name)))
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
            if StringClient.contains(StringClient.to_lower(header), 'x-oss-') and not UtilClient.is_unset(headers.get(header)):
                canonicalized_headers = '%s%s:%s\n' % (TeaConverter.to_unicode(canonicalized_headers), TeaConverter.to_unicode(header), TeaConverter.to_unicode(headers.get(header)))
        return canonicalized_headers

    def get_signature_v2(self, bucket_name, pathname, method, query, headers, secret):
        return ''
