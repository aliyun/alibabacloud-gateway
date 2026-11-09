# -*- coding: utf-8 -*-
import asyncio
import io
import os
import unittest

from Tea.exceptions import TeaException
from alibabacloud_credentials.client import Client as CredentialClient
from alibabacloud_credentials.models import Config as CredentialConfig
from alibabacloud_gateway_pop.client import Client
from alibabacloud_gateway_spi.models import (
    AttributeMap,
    InterceptorContext,
    InterceptorContextConfiguration,
    InterceptorContextRequest,
    InterceptorContextResponse,
)


class Model(object):
    def __init__(self, type_, provider_name='static_ak', access_key_id='ak', access_key_secret='secret',
                 security_token=None, bearer_token=None):
        self.type = type_
        self.provider_name = provider_name
        self.access_key_id = access_key_id
        self.access_key_secret = access_key_secret
        self.security_token = security_token
        self.bearer_token = bearer_token


class FixedCredential(object):
    def __init__(self, model=None, error=None):
        self._model = model
        self._error = error

    def get_credential(self):
        if self._error is not None:
            raise self._error
        return self._model

    async def get_credential_async(self):
        return self.get_credential()


class BoomStream(object):
    def read(self, size=-1):
        raise IOError('read failed')


class AsyncStream(object):
    def __init__(self, data=b'', error=None):
        self._data = data
        self._error = error

    async def read(self, size=-1):
        if self._error is not None:
            raise self._error
        data = self._data
        self._data = b''
        return data


def access_key():
    return CredentialClient(CredentialConfig(
        type='access_key',
        access_key_id='ak',
        access_key_secret='secret',
    ))


def request(credential):
    return InterceptorContextRequest(
        headers={},
        query={},
        pathname='/',
        product_id='ecs',
        action='DescribeRegions',
        version='2014-05-26',
        protocol='HTTPS',
        method='POST',
        auth_type='AK',
        body_type='json',
        req_body_type='json',
        credential=credential,
        user_agent='test-agent',
    )


def config(endpoint='ecs.cn-hangzhou.aliyuncs.com'):
    return InterceptorContextConfiguration(
        region_id='cn-hangzhou',
        endpoint=endpoint,
    )


def context(req, cfg):
    return InterceptorContext(
        request=req,
        configuration=cfg,
        response=InterceptorContextResponse(),
    )


class TestCoverage(unittest.TestCase):
    def test_explicit_endpoint_suffix(self):
        client = Client()
        for original, expected in [
            ('oss-cn-hangzhou.aliyuncs.com', 'oss-cn-hangzhou.aliyuncs.alibaba'),
            ('ecs.cn-hangzhou-internal.aliyuncs.com', 'ecs.cn-hangzhou-internal.aliyuncs.alibaba'),
            ('ecs.aliyun.com', 'ecs.aliyun.com'),
            ('custom.example.com', 'custom.example.com'),
            ('ecs.aliyuncs.alibaba', 'ecs.aliyuncs.alibaba'),
            ('ecs.aliyuncs.com.example.com', 'ecs.aliyuncs.com.example.com'),
        ]:
            for mapped in (False, True):
                cfg = config(original)
                if mapped:
                    cfg.endpoint = None
                    cfg.endpoint_map = {'cn-hangzhou': original}
                attrs = AttributeMap(key={'endpointSuffix': 'aliyuncs.alibaba'})
                for _ in range(2):
                    client.modify_configuration(context(request(access_key()), cfg), attrs)
                    self.assertEqual(expected, cfg.endpoint)

    def test_endpoint_and_configuration(self):
        client = Client()
        self.assertEqual('ecs.example.com', client.get_endpoint('ecs', 'cn-hangzhou', 'regional', '', '', None, 'ecs.example.com'))
        self.assertEqual('ecs.mapped.example.com', client.get_endpoint('ecs', 'cn-hangzhou', 'regional', '', '', {'cn-hangzhou': 'ecs.mapped.example.com'}, None))
        self.assertEqual('ecs.cn-hangzhou.aliyuncs.com', client.get_endpoint('Ecs', 'cn-hangzhou', 'regional', '', '', None, None))
        self.assertEqual('ecs.aliyuncs.com', client.get_endpoint('Ecs', None, 'central', 'public', None, None, ''))
        with self.assertRaises(Exception):
            client.get_endpoint('Ecs', '', 'regional', '', '', None, None)

        built = InterceptorContextConfiguration(region_id='cn-hangzhou', endpoint_rule='regional', network='', suffix='')
        client.modify_configuration(context(request(access_key()), built), AttributeMap(key={
            'signatureTypePrefix': 'ACS4-',
            'signPrefix': 'aliyun_v4',
            'endpointSuffix': 'aliyuncs.com',
        }))
        self.assertEqual('ecs.cn-hangzhou.aliyuncs.com', built.endpoint)
        self.assertEqual('ACS4-HMAC-SHA256', client._sha_256)

        kept = config('already.example.com')
        client.modify_configuration(context(request(access_key()), kept), AttributeMap())
        self.assertEqual('already.example.com', kept.endpoint)

        broken = InterceptorContextConfiguration(region_id='', endpoint_rule='regional')
        with self.assertRaises(Exception):
            client.modify_configuration(context(request(access_key()), broken), AttributeMap())

    def test_signing_uses_selected_endpoint_after_environment_change(self):
        previous = os.environ.get('ALIBABA_CLOUD_ENDPOINT_SUFFIX')
        try:
            for asynchronous in (False, True):
                for selected in ('aliyuncs.com', 'aliyuncs.alibaba'):
                    for changed in ('aliyuncs.com', 'aliyuncs.alibaba', 'invalid.example'):
                        os.environ['ALIBABA_CLOUD_ENDPOINT_SUFFIX'] = selected
                        client = Client()
                        built = config('ecs.cn-hangzhou.aliyuncs.com')
                        built.region_id = ''
                        ctx = context(request(access_key()), built)
                        attributes = AttributeMap()
                        client.modify_configuration(ctx, attributes)
                        os.environ['ALIBABA_CLOUD_ENDPOINT_SUFFIX'] = changed
                        if asynchronous:
                            loop = asyncio.new_event_loop()
                            try:
                                loop.run_until_complete(client.modify_request_async(ctx, attributes))
                            finally:
                                loop.close()
                        else:
                            client.modify_request(ctx, attributes)
                        self.assertEqual('ecs.cn-hangzhou.' + selected, ctx.request.headers['host'])
                        self.assertIn('/cn-hangzhou/ecs/aliyun_v4_request', ctx.request.headers['Authorization'])
        finally:
            if previous is None:
                os.environ.pop('ALIBABA_CLOUD_ENDPOINT_SUFFIX', None)
            else:
                os.environ['ALIBABA_CLOUD_ENDPOINT_SUFFIX'] = previous

    def test_endpoint_suffix_env_and_attribute(self):
        os.environ['ALIBABA_CLOUD_ENDPOINT_SUFFIX'] = 'aliyuncs.alibaba'
        try:
            client = Client()
            built = InterceptorContextConfiguration(region_id='cn-hangzhou', endpoint_rule='regional', network='', suffix='')
            client.modify_configuration(context(request(access_key()), built), AttributeMap())
            self.assertEqual('ecs.cn-hangzhou.aliyuncs.alibaba', built.endpoint)
            self.assertEqual('cn-hangzhou', client.get_region('ecs', 'ecs.cn-hangzhou.aliyuncs.com', ''))
            self.assertEqual('cn-hangzhou', client._get_region_with_suffix('ecs', built.endpoint, '', 'aliyuncs.alibaba'))
            self.assertEqual('cn-hangzhou', client._get_region_with_suffix('ecs', 'ecs.cn-hangzhou.aliyuncs.com', '', 'example.net'))
            self.assertEqual('aliyuncs.com', client._endpoint_suffix)
            self.assertEqual('ACS4-HMAC-SHA256', client._sha_256)

            overridden = InterceptorContextConfiguration(region_id='cn-hangzhou', endpoint_rule='regional', network='', suffix='')
            client.modify_configuration(context(request(access_key()), overridden), AttributeMap(key={
                'endpointSuffix': ' ALIYUNCS.ALIBABA ',
            }))
            self.assertEqual('ecs.cn-hangzhou.aliyuncs.alibaba', overridden.endpoint)
            self.assertEqual('ACS4-HMAC-SHA256', client._sha_256)
            self.assertEqual('leave.example', client._with_endpoint_suffix('leave.example'))
            self.assertEqual('', client._with_endpoint_suffix(''))
            client._endpoint_suffix = ''
            self.assertEqual('aliyuncs.alibaba', client._resolve_endpoint_suffix({
                'signatureTypePrefix': '',
                'signPrefix': '',
                'endpointSuffix': '',
            }))
            self.assertEqual('', client._endpoint_suffix)
        finally:
            os.environ.pop('ALIBABA_CLOUD_ENDPOINT_SUFFIX', None)
        fallback = Client()
        fallback._endpoint_suffix = ''
        self.assertEqual('aliyuncs.com', fallback._resolve_endpoint_suffix(None))
        self.assertEqual('', fallback._endpoint_suffix)
        kept = InterceptorContextConfiguration(region_id='cn-hangzhou', endpoint_rule='regional', network='', suffix='')
        fallback.modify_configuration(context(request(access_key()), kept), AttributeMap(key={'endpointSuffix': 'aliyuncs.alibaba'}))
        self.assertEqual('ecs.cn-hangzhou.aliyuncs.alibaba', kept.endpoint)
        with self.assertRaises(Exception):
            fallback.modify_configuration(context(request(access_key()), kept), AttributeMap(key={'endpointSuffix': 'example.net'}))
        again = InterceptorContextConfiguration(region_id='cn-hangzhou', endpoint_rule='regional', network='', suffix='')
        fallback.modify_configuration(context(request(access_key()), again), AttributeMap())
        self.assertEqual('ecs.cn-hangzhou.aliyuncs.com', again.endpoint)

    def test_subclass_keeps_original_endpoint_signatures(self):
        class Child(Client):
            def get_endpoint(self, product_id, region_id, endpoint_rule, network, suffix, endpoint_map, endpoint):
                return 'child.example'

            def get_region(self, product, endpoint, region_id):
                return 'child-region'

        child = Child()
        built = InterceptorContextConfiguration(region_id='cn-hangzhou', endpoint_rule='regional', network='', suffix='')
        child.modify_configuration(context(request(access_key()), built), AttributeMap(key={'endpointSuffix': 'aliyuncs.alibaba'}))
        self.assertEqual('ecs.cn-hangzhou.aliyuncs.alibaba', built.endpoint)
        self.assertEqual('child.example', child.get_endpoint('ecs', 'cn-hangzhou', 'regional', '', '', None, ''))
        self.assertEqual('child-region', child.get_region('ecs', 'ecs.cn-hangzhou.aliyuncs.com', ''))
        self.assertEqual('cn-hangzhou', child._get_region_with_suffix('ecs', 'ecs.cn-hangzhou.aliyuncs.com', '', 'example.net'))

    def test_canonical_and_signing(self):
        client = Client()
        self.assertEqual('', client.build_canonicalized_resource(None))
        resource = client.build_canonicalized_resource({'b': 'x y', 'a': 'm*n', 'empty': ''})
        self.assertIn('a=m%2An', resource)
        self.assertIn('b=x%20y', resource)
        self.assertIn('empty=', resource)
        query = {'a-b': '1', 'a/b': '2', 'a': '', 'é': '中', '10': 'x', '2': 'y',
                 'null': None, 'punct': '!+*~', 'Signature': 'ignored'}
        self.assertEqual('%C3%A9=%E4%B8%AD&10=x&2=y&a=&a%2Fb=2&a-b=1&null=&punct=%21%2B%2A~',
                         client.build_canonicalized_resource(query))

        canonical = client.build_canonicalized_headers({
            'Host': '  example.com  ',
            'X-Acs-Action': 'One',
            'x-acs-action': 'Two',
        })
        self.assertIn('host:example.com\n', canonical)
        self.assertIn('x-acs-action:', canonical)
        self.assertEqual('fallback', client.default_any(None, 'fallback'))
        self.assertEqual('kept', client.default_any('kept', 'fallback'))

        signingkey = client.get_signingkey(client._sha_256, 'secret', 'ecs', 'cn-hangzhou', '20240101')
        self.assertTrue(signingkey)
        unknown = client.get_signingkey('OTHER', 'secret', 'ecs', 'cn-hangzhou', '20240101')
        self.assertFalse(unknown)
        headers = {'host': 'example.com', 'x-acs-date': '2024-01-01T00:00:00Z'}
        signature = client.get_signature('', 'GET', {'q': 'a b'}, headers, client._sha_256, 'payload', signingkey)
        self.assertEqual(64, len(signature))
        payload = 'e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855'
        self.assertEqual('d5b8fb7a183e2e8773170a77534441983872ca20110d67d71c98baf342dbeb33',
                         client.get_signature('/', 'get', query, headers, client._sha_256, payload, signingkey))
        authorization = client.get_authorization('/', 'GET', None, headers, client._sha_256, 'payload', 'ak', signingkey, 'ecs', 'cn-hangzhou', '20240101')
        self.assertIn('ACS4-HMAC-SHA256 Credential=ak/20240101/cn-hangzhou/ecs/aliyun_v4_request,SignedHeaders=', authorization)
        self.assertIn('host;x-acs-date', authorization)

        sm3key = client.get_signingkey(client._sm_3, 'secret', 'ecs', 'cn-hangzhou', '20240101')
        sm3_signature = client.get_signature('/', 'POST', None, headers, client._sm_3, 'payload', sm3key)
        self.assertEqual(64, len(sm3_signature))

    def _exercise_request(self, async_mode=False):
        client = Client()
        method = client.modify_request_async if async_mode else client.modify_request

        def caller(ctx):
            if async_mode:
                asyncio.run(method(ctx, AttributeMap()))
            else:
                method(ctx, AttributeMap())

        def stream(data):
            return AsyncStream(data) if async_mode else io.BytesIO(data)

        def boom():
            return AsyncStream(error=IOError('read failed')) if async_mode else BoomStream()

        json_request = request(access_key())
        json_request.body = {'name': 'a b'}
        json_request.query = {'InstanceId': 'i-1'}
        caller(context(json_request, config()))
        self.assertEqual('application/json; charset=utf-8', json_request.headers['content-type'])
        self.assertEqual('DescribeRegions', json_request.headers['x-acs-action'])
        self.assertIn('ACS4-HMAC-SHA256 Credential=ak/', json_request.headers['Authorization'])
        self.assertTrue(json_request.headers['x-acs-content-sha256'])

        byte_request = request(access_key())
        byte_request.req_body_type = 'byte'
        byte_request.body = b'hello'
        byte_request.action = ''
        caller(context(byte_request, config()))
        self.assertNotIn('x-acs-action', byte_request.headers)

        form_request = request(access_key())
        form_request.req_body_type = 'form'
        form_request.body = {'b': 'x y', 'a': ['m', 'n']}
        caller(context(form_request, config()))
        self.assertEqual('application/x-www-form-urlencoded', form_request.headers['content-type'])

        stream_request = request(access_key())
        stream_request.stream = stream(b'stream-body')
        caller(context(stream_request, config()))
        self.assertEqual('application/octet-stream', stream_request.headers['content-type'])

        caller(context(request(access_key()), config()))

        sm3_request = request(access_key())
        sm3_request.signature_algorithm = client._sm_3
        sm3_request.body = {'k': 'v'}
        caller(context(sm3_request, config()))
        self.assertTrue(sm3_request.headers['x-acs-content-sm3'])
        self.assertIn('ACS4-HMAC-SM3 Credential=ak/', sm3_request.headers['Authorization'])

        bearer = request(FixedCredential(Model('bearer', provider_name='bearer', bearer_token='bearer-token')))
        caller(context(bearer, config()))
        self.assertEqual('Bearer bearer-token', bearer.headers['Authorization'])

        sts = request(FixedCredential(Model('sts', provider_name='', access_key_id='sts-ak', security_token='token')))
        caller(context(sts, config()))
        self.assertEqual('token', sts.headers['x-acs-security-token'])

        id_token = request(FixedCredential(Model('id_token', provider_name=None, security_token='id-token')))
        caller(context(id_token, config()))
        self.assertEqual('id-token', id_token.headers['x-acs-zero-trust-idtoken'])

        anonymous = request(None)
        anonymous.auth_type = 'Anonymous'
        caller(context(anonymous, config()))
        self.assertNotIn('Authorization', anonymous.headers)

        missing = request(None)
        with self.assertRaises(TeaException) as ctx:
            caller(context(missing, config()))
        self.assertEqual('ParameterMissing', ctx.exception.code)

        with self.assertRaises(RuntimeError):
            caller(context(request(FixedCredential(error=RuntimeError('credential down'))), config()))

        bad_bytes = request(access_key())
        bad_bytes.req_body_type = 'byte'
        bad_bytes.body = 'not-bytes'
        with self.assertRaises(Exception):
            caller(context(bad_bytes, config()))

        bad_form = request(access_key())
        bad_form.req_body_type = 'form'
        bad_form.body = 'not-a-map'
        with self.assertRaises(Exception):
            caller(context(bad_form, config()))

        bad_stream = request(access_key())
        bad_stream.stream = boom()
        with self.assertRaises(Exception):
            caller(context(bad_stream, config()))

    def test_modify_request(self):
        self._exercise_request(False)

    def test_modify_request_async(self):
        self._exercise_request(True)

    def _response(self, status, content_type, payload, body_type, async_mode=False):
        body = AsyncStream(payload) if async_mode else io.BytesIO(payload)
        ctx = context(request(access_key()), config())
        ctx.request.body_type = body_type
        ctx.response = InterceptorContextResponse(
            status_code=status,
            headers={} if content_type is None else {'content-type': content_type},
            body=body,
        )
        return ctx

    def _exercise_response(self, method_name):
        client = Client()

        def call(ctx):
            method = getattr(client, method_name)
            if method_name.endswith('_async'):
                asyncio.run(method(ctx, AttributeMap()))
            else:
                method(ctx, AttributeMap())

        with self.assertRaises(TeaException) as ctx:
            call(self._response(400, 'application/json', b'{"Code":"Invalid","Message":"no","RequestId":"req-1","Description":"d"}', 'json', method_name.endswith('_async')))
        self.assertEqual('Invalid', ctx.exception.code)
        self.assertIn('request id: req-1', ctx.exception.message)

        denied = self._response(403, 'application/json', b'{"code":"Denied","message":"no","requestId":"body-id"}', 'json', method_name.endswith('_async'))
        denied.response.headers['x-acs-request-id'] = 'header-id'
        with self.assertRaises(TeaException) as ctx:
            call(denied)
        self.assertEqual('Denied', ctx.exception.code)

        xml = b'<Error><Code>XmlCode</Code><Message>xml</Message><RequestId>xml-id</RequestId><Description>xd</Description></Error>'
        with self.assertRaises(TeaException) as ctx:
            call(self._response(500, 'text/xml', xml, 'xml', method_name.endswith('_async')))
        self.assertEqual('XmlCode', ctx.exception.code)

        async_mode = method_name.endswith('_async')
        call(self._response(204, None, b'', 'string', async_mode))
        binary = self._response(200, None, b'raw', 'binary', async_mode)
        call(binary)
        self.assertIsNotNone(binary.response.deserialized_body)
        text = self._response(200, None, b'hello', 'string', async_mode)
        call(text)
        self.assertEqual('hello', text.response.deserialized_body)
        parsed = self._response(200, None, b'{"ok":true}', 'json', async_mode)
        call(parsed)
        self.assertEqual(True, parsed.response.deserialized_body['ok'])
        array = self._response(200, None, b'[1,2]', 'array', async_mode)
        call(array)
        self.assertEqual(2, len(array.response.deserialized_body))
        other = self._response(200, None, b'plain', 'other', async_mode)
        call(other)
        self.assertEqual('plain', other.response.deserialized_body)
        byte_body = self._response(200, None, b'raw', 'byte', async_mode)
        call(byte_body)
        self.assertTrue(byte_body.response.deserialized_body)
        with self.assertRaises(Exception):
            call(self._response(200, None, b'{"ok":', 'json', async_mode))

    def test_modify_response(self):
        self._exercise_response('modify_response')

    def test_modify_response_async(self):
        self._exercise_response('modify_response_async')

    def test_modify_configuration_async(self):
        client = Client()
        built = InterceptorContextConfiguration(region_id='cn-hangzhou', endpoint_rule='regional', network='', suffix='')
        asyncio.run(client.modify_configuration_async(context(request(access_key()), built), AttributeMap(key={
            'signatureTypePrefix': 'ACS4-',
            'signPrefix': 'aliyun_v4',
            'endpointSuffix': 'aliyuncs.com',
        })))
        self.assertEqual('ecs.cn-hangzhou.aliyuncs.com', built.endpoint)
