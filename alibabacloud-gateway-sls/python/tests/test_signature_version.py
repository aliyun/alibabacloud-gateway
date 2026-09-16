import asyncio
import hashlib
import json
from pathlib import Path
import unittest

from alibabacloud_credentials.client import Client as Credential
from alibabacloud_credentials.models import Config
from alibabacloud_gateway_spi import models
from alibabacloud_gateway_sls.client import Client


class SignatureVersionTest(unittest.TestCase):
    def test_selection(self):
        path = Path(__file__).resolve().parents[2] / 'testdata/signature-version.json'
        cases = json.loads(path.read_text())
        client = Client()
        for case in cases:
            for asynchronous in (False, True):
                with self.subTest(case=case['name'], asynchronous=asynchronous):
                    context = models.InterceptorContext(
                        request=models.InterceptorContextRequest(signature_version=case['version']),
                        configuration=models.InterceptorContextConfiguration(
                            region_id=case['region'], endpoint=case['endpoint']))
                    if asynchronous:
                        version = asyncio.run(client.get_signature_version_async(context))
                    else:
                        version = client.get_signature_version(context)
                    self.assertEqual(case['want'], version)
                    self.assertEqual(case['resolved'], context.configuration.region_id)
                    self.assertEqual(case['version'], context.request.signature_version)

    def test_signed_request(self):
        for asynchronous in (False, True):
            for version in (None, '', 'v1'):
                for body in (None, b'request body'):
                    with self.subTest(asynchronous=asynchronous, version=version, body=body):
                        client = Client()
                        request = models.InterceptorContextRequest(
                            signature_version=version,
                            credential=Credential(Config(type='access_key', access_key_id='test-id', access_key_secret='test-secret')),
                            headers={}, query={}, method='POST', pathname='/',
                            action='TestAction', req_body_type='binary', user_agent='test', body=body)
                        context = models.InterceptorContext(
                            request=request, configuration=models.InterceptorContextConfiguration(
                                endpoint='cn-acdr-ut-1-internal.log.aliyuncs.com'))
                        if asynchronous:
                            asyncio.run(client.modify_request_async(context, models.AttributeMap()))
                        else:
                            client.modify_request(context, models.AttributeMap())
                        authorization = request.headers['authorization']
                        if version == 'v1':
                            self.assertTrue(authorization.startswith('LOG '))
                            self.assertNotIn('x-log-content-sha256', request.headers)
                        else:
                            self.assertTrue(authorization.startswith('SLS4-HMAC-SHA256 '))
                            self.assertIn('/cn-acdr-ut-1/sls/aliyun_v4_request', authorization)
                            self.assertEqual(hashlib.sha256(body or b'').hexdigest(), request.headers['x-log-content-sha256'])


if __name__ == '__main__':
    unittest.main()
