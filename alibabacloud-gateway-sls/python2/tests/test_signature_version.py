import io
import json
import os
import unittest

from alibabacloud_gateway_spi import models
from alibabacloud_gateway_sls.client import Client


class SignatureVersionTest(unittest.TestCase):
    def test_selection(self):
        path = os.path.join(os.path.dirname(__file__), '../../testdata/signature-version.json')
        with io.open(path, encoding='utf-8') as stream:
            cases = json.load(stream)
        client = Client()
        for case in cases:
            context = models.InterceptorContext(
                request=models.InterceptorContextRequest(signature_version=case['version']),
                configuration=models.InterceptorContextConfiguration(
                    region_id=case['region'], endpoint=case['endpoint']))
            self.assertEqual(case['want'], client.get_signature_version(context), case['name'])
            self.assertEqual(case['resolved'], context.configuration.region_id, case['name'])
            self.assertEqual(case['version'], context.request.signature_version, case['name'])


if __name__ == '__main__':
    unittest.main()
