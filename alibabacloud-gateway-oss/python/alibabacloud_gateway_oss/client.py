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
from alibabacloud_gateway_oss_util.client import Client as OSS_UtilClient


class Client(SPIClient):
    _default_signed_params: List[str] = None
    _except_signed_params: List[str] = None

    def __init__(self):
        super().__init__()
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
            'userDefinedLogFieldsConfig',
            'reservedcapacity',
            'requesterQosInfo',
            'qosRequester',
            'resourcePool',
            'resourcePoolInfo',
            'resourcePoolBuckets',
            'processConfiguration',
            'img',
            'asyncFetch',
            'virtualBucket',
            'copy',
            'userRegion',
            'partSize',
            'chunkSize',
            'partUploadId',
            'chunkNumber',
            'userRegion',
            'regionList',
            'eventnotification',
            'cacheConfiguration',
            'dfs',
            'dfsadmin',
            'dfssecurity'
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
        if UtilClient.is_unset(region_id) or UtilClient.empty(region_id):
            region_id = self.get_region_id_from_endpoint(config.endpoint)
        credential = request.credential
        access_key_id = credential.get_access_key_id()
        access_key_secret = credential.get_access_key_secret()
        security_token = credential.get_security_token()
        if not UtilClient.empty(security_token):
            request.headers['x-oss-security-token'] = security_token
        if not UtilClient.is_unset(request.body):
            if StringClient.equals(request.req_body_type, 'xml'):
                req_body_map = UtilClient.assert_as_map(request.body)
                xml_str = OSS_UtilClient.to_xml(req_body_map)
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
        host = self.get_host(config.endpoint_type, bucket_name, config.endpoint, context)
        request.headers = TeaCore.merge({
            'host': host,
            'date': UtilClient.get_date_utcstring(),
            'user-agent': request.user_agent
        }, request.headers)
        origin_path = request.pathname
        origin_query = request.query
        if not UtilClient.empty(origin_path):
            path_and_queries = StringClient.split(origin_path, f'?', 2)
            request.pathname = path_and_queries[0]
            if UtilClient.equal_number(ArrayClient.size(path_and_queries), 2):
                path_queries = StringClient.split(path_and_queries[1], '&', None)
                for sub in path_queries:
                    item = StringClient.split(sub, '=', None)
                    query_key = item[0]
                    query_value = ''
                    if UtilClient.equal_number(ArrayClient.size(item), 2):
                        query_value = item[1]
                    if UtilClient.empty(origin_query.get(query_key)):
                        request.query[query_key] = query_value
        signature_version = UtilClient.default_string(request.signature_version, 'v1')
        request.headers['authorization'] = self.get_authorization(signature_version, bucket_name, request.pathname, request.method, request.query, request.headers, access_key_id, access_key_secret, region_id)

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
        if UtilClient.is_unset(region_id) or UtilClient.empty(region_id):
            region_id = await self.get_region_id_from_endpoint_async(config.endpoint)
        credential = request.credential
        access_key_id = await credential.get_access_key_id_async()
        access_key_secret = await credential.get_access_key_secret_async()
        security_token = await credential.get_security_token_async()
        if not UtilClient.empty(security_token):
            request.headers['x-oss-security-token'] = security_token
        if not UtilClient.is_unset(request.body):
            if StringClient.equals(request.req_body_type, 'xml'):
                req_body_map = UtilClient.assert_as_map(request.body)
                xml_str = OSS_UtilClient.to_xml(req_body_map)
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
        host = await self.get_host_async(config.endpoint_type, bucket_name, config.endpoint, context)
        request.headers = TeaCore.merge({
            'host': host,
            'date': UtilClient.get_date_utcstring(),
            'user-agent': request.user_agent
        }, request.headers)
        origin_path = request.pathname
        origin_query = request.query
        if not UtilClient.empty(origin_path):
            path_and_queries = StringClient.split(origin_path, f'?', 2)
            request.pathname = path_and_queries[0]
            if UtilClient.equal_number(ArrayClient.size(path_and_queries), 2):
                path_queries = StringClient.split(path_and_queries[1], '&', None)
                for sub in path_queries:
                    item = StringClient.split(sub, '=', None)
                    query_key = item[0]
                    query_value = ''
                    if UtilClient.equal_number(ArrayClient.size(item), 2):
                        query_value = item[1]
                    if UtilClient.empty(origin_query.get(query_key)):
                        request.query[query_key] = query_value
        signature_version = UtilClient.default_string(request.signature_version, 'v1')
        request.headers['authorization'] = await self.get_authorization_async(signature_version, bucket_name, request.pathname, request.method, request.query, request.headers, access_key_id, access_key_secret, region_id)

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
                        'hostId': err.get('HostId'),
                        'AccessDeniedDetail': err.get('AccessDeniedDetail')
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
                    result = OSS_UtilClient.parse_xml(body_str, request.action)
                    # for no util language
                    # var result : any = XML.parseXml(bodyStr, null);
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
                        'hostId': err.get('HostId'),
                        'AccessDeniedDetail': err.get('AccessDeniedDetail')
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
                    result = await OSS_UtilClient.parse_xml_async(body_str, request.action)
                    # for no util language
                    # var result : any = XML.parseXml(bodyStr, null);
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

    def get_region_id_from_endpoint(
        self,
        endpoint: str,
    ) -> str:
        if not UtilClient.empty(endpoint):
            idx = -1
            if StringClient.has_prefix(endpoint, 'oss-') and StringClient.has_suffix(endpoint, '.aliyuncs.com'):
                idx = StringClient.index(endpoint, '.aliyuncs.com')
                return StringClient.sub_string(endpoint, 4, idx)
            if StringClient.has_suffix(endpoint, '.mgw.aliyuncs.com'):
                idx = StringClient.index(endpoint, '.mgw.aliyuncs.com')
                return StringClient.sub_string(endpoint, 0, idx)
            if StringClient.has_suffix(endpoint, '-internal.oss-data-acc.aliyuncs.com'):
                idx = StringClient.index(endpoint, '-internal.oss-data-acc.aliyuncs.com')
                return StringClient.sub_string(endpoint, 0, idx)
            if StringClient.has_suffix(endpoint, '.oss-dls.aliyuncs.com'):
                idx = StringClient.index(endpoint, '.oss-dls.aliyuncs.com')
                return StringClient.sub_string(endpoint, 0, idx)
        return ''

    async def get_region_id_from_endpoint_async(
        self,
        endpoint: str,
    ) -> str:
        if not UtilClient.empty(endpoint):
            idx = -1
            if StringClient.has_prefix(endpoint, 'oss-') and StringClient.has_suffix(endpoint, '.aliyuncs.com'):
                idx = StringClient.index(endpoint, '.aliyuncs.com')
                return StringClient.sub_string(endpoint, 4, idx)
            if StringClient.has_suffix(endpoint, '.mgw.aliyuncs.com'):
                idx = StringClient.index(endpoint, '.mgw.aliyuncs.com')
                return StringClient.sub_string(endpoint, 0, idx)
            if StringClient.has_suffix(endpoint, '-internal.oss-data-acc.aliyuncs.com'):
                idx = StringClient.index(endpoint, '-internal.oss-data-acc.aliyuncs.com')
                return StringClient.sub_string(endpoint, 0, idx)
            if StringClient.has_suffix(endpoint, '.oss-dls.aliyuncs.com'):
                idx = StringClient.index(endpoint, '.oss-dls.aliyuncs.com')
                return StringClient.sub_string(endpoint, 0, idx)
        return ''

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
        context: spi_models.InterceptorContext,
    ) -> str:
        if StringClient.contains(endpoint, '.mgw.aliyuncs.com') and not UtilClient.is_unset(context.request.host_map.get('userid')):
            return f"{context.request.host_map.get('userid')}.{endpoint}"
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
        context: spi_models.InterceptorContext,
    ) -> str:
        if StringClient.contains(endpoint, '.mgw.aliyuncs.com') and not UtilClient.is_unset(context.request.host_map.get('userid')):
            return f"{context.request.host_map.get('userid')}.{endpoint}"
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
            object_name = pathname
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
            object_name = pathname
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
        canonicalized_resource = pathname
        query_keys = MapClient.key_set(query)
        sorted_params = ArrayClient.asc_sort(query_keys)
        separator = '?'
        for param_name in sorted_params:
            if ArrayClient.contains(self._default_signed_params, param_name) or StringClient.has_prefix(param_name, 'x-oss-'):
                canonicalized_resource = f'{canonicalized_resource}{separator}{param_name}'
                if not UtilClient.empty(query.get(param_name)):
                    canonicalized_resource = f'{canonicalized_resource}={query.get(param_name)}'
                separator = '&'
        return canonicalized_resource

    async def build_canonicalized_resource_async(
        self,
        pathname: str,
        query: Dict[str, str],
    ) -> str:
        canonicalized_resource = pathname
        query_keys = MapClient.key_set(query)
        sorted_params = ArrayClient.asc_sort(query_keys)
        separator = '?'
        for param_name in sorted_params:
            if ArrayClient.contains(self._default_signed_params, param_name) or StringClient.has_prefix(param_name, 'x-oss-'):
                canonicalized_resource = f'{canonicalized_resource}{separator}{param_name}'
                if not UtilClient.empty(query.get(param_name)):
                    canonicalized_resource = f'{canonicalized_resource}={query.get(param_name)}'
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
