# -*- coding: utf-8 -*-
import unittest

from alibabacloud_gateway_spi import models as spi_models
from alibabacloud_gateway_sls.client import Client


class TestClient(unittest.TestCase):

    def _context(self, region_id=None, endpoint=None, signature_version=None):
        return spi_models.InterceptorContext(
            request=spi_models.InterceptorContextRequest(signature_version=signature_version),
            configuration=spi_models.InterceptorContextConfiguration(
                region_id=region_id,
                endpoint=endpoint,
            ),
        )

    def test_parse_region(self):
        client = Client()
        cases = [
            ('', ''),
            ('cn-hangzhou-acdr-ut-1.sls.aliyuncs.com', 'cn-hangzhou-acdr-ut-1'),
            ('cn-hangzhou-acdr-ut-1.log.aliyuncs.com', 'cn-hangzhou-acdr-ut-1'),
            ('https://cn-hangzhou-acdr-ut-1.sls.aliyuncs.com', 'cn-hangzhou-acdr-ut-1'),
            ('http://cn-hangzhou-acdr-ut-1.log.aliyuncs.com', 'cn-hangzhou-acdr-ut-1'),
            ('cn-hangzhou-acdr-ut-1-intranet.sls.aliyuncs.com', 'cn-hangzhou-acdr-ut-1'),
            ('cn-hangzhou-acdr-ut-1-share.log.aliyuncs.com', 'cn-hangzhou-acdr-ut-1'),
            ('cn-hangzhou-acdr-ut-1-vpc.sls.aliyuncs.com', 'cn-hangzhou-acdr-ut-1'),
            ('cn-hangzhou-acdr-ut-1-internal.log.aliyuncs.com', 'cn-hangzhou-acdr-ut-1'),
            ('https://cn-hangzhou-acdr-ut-1-intranet.log.aliyuncs.com', 'cn-hangzhou-acdr-ut-1'),
            ('cn-hangzhou.sls.aliyuncs.com', 'cn-hangzhou'),
            ('cn-hangzhou.log.aliyuncs.com', 'cn-hangzhou'),
            ('ftp://cn-hangzhou-acdr-ut-1.sls.aliyuncs.com', ''),
            ('cn-hangzhou-acdr-ut-1.oss.aliyuncs.com', ''),
            ('cn-hangzhou-acdr-ut-1.sls.aliyuncs.com.cn', ''),
            ('cn-hangzhou-acdr-ut-1.sls.aliyuncs.com/path', ''),
            ('https://https://cn-hangzhou-acdr-ut-1.sls.aliyuncs.com', ''),
            ('CN-HANGZHOU.sls.aliyuncs.com', ''),
            ('cn_hangzhou.sls.aliyuncs.com', ''),
            ('cn-hangzhou-acdr-ut-1-intranet-share.sls.aliyuncs.com', 'cn-hangzhou-acdr-ut-1-intranet'),
        ]
        for endpoint, want in cases:
            self.assertEqual(want, client.parse_region(endpoint), endpoint)
        self.assertEqual('', client.parse_region(None))

    def test_set_sign_v4_if_in_acdr(self):
        client = Client()

        acdr_region = self._context('cn-hangzhou-acdr-ut-1', 'cn-hangzhou.log.aliyuncs.com')
        client.set_sign_v4_if_in_acdr(acdr_region)
        self.assertEqual('v4', acdr_region.request.signature_version)
        self.assertEqual('cn-hangzhou-acdr-ut-1', acdr_region.configuration.region_id)

        ordinary_region = self._context('cn-hangzhou', 'cn-hangzhou.log.aliyuncs.com')
        client.set_sign_v4_if_in_acdr(ordinary_region)
        self.assertIsNone(ordinary_region.request.signature_version)

        from_endpoint = self._context(None, 'cn-hangzhou-acdr-ut-1.sls.aliyuncs.com')
        client.set_sign_v4_if_in_acdr(from_endpoint)
        self.assertEqual('v4', from_endpoint.request.signature_version)
        self.assertEqual('cn-hangzhou-acdr-ut-1', from_endpoint.configuration.region_id)

        ordinary_endpoint = self._context('', 'cn-hangzhou.sls.aliyuncs.com')
        client.set_sign_v4_if_in_acdr(ordinary_endpoint)
        self.assertFalse(ordinary_endpoint.request.signature_version)
        self.assertEqual('', ordinary_endpoint.configuration.region_id)

        empty_endpoint = self._context(None, None)
        client.set_sign_v4_if_in_acdr(empty_endpoint)
        self.assertFalse(empty_endpoint.request.signature_version)

        explicit_v1 = self._context('cn-hangzhou-acdr-ut-1', 'cn-hangzhou-acdr-ut-1.sls.aliyuncs.com', 'v1')
        client.set_sign_v4_if_in_acdr(explicit_v1)
        self.assertEqual('v1', explicit_v1.request.signature_version)

        explicit_v4 = self._context('cn-hangzhou', 'cn-hangzhou.log.aliyuncs.com', 'v4')
        client.set_sign_v4_if_in_acdr(explicit_v4)
        self.assertEqual('v4', explicit_v4.request.signature_version)

        intranet = self._context('', 'https://cn-hangzhou-acdr-ut-1-intranet.log.aliyuncs.com')
        client.set_sign_v4_if_in_acdr(intranet)
        self.assertEqual('v4', intranet.request.signature_version)
        self.assertEqual('cn-hangzhou-acdr-ut-1', intranet.configuration.region_id)

        invalid_acdr_hint = self._context(None, 'not-a-valid-acdr-ut-host')
        client.set_sign_v4_if_in_acdr(invalid_acdr_hint)
        self.assertFalse(invalid_acdr_hint.request.signature_version)
        self.assertIsNone(invalid_acdr_hint.configuration.region_id)

        shanghai = self._context('cn-shanghai-acdr-ut-2', None)
        client.set_sign_v4_if_in_acdr(shanghai)
        self.assertEqual('v4', shanghai.request.signature_version)

        explicit_v1_from_endpoint = self._context('', 'cn-hangzhou-acdr-ut-1.sls.aliyuncs.com', 'v1')
        client.set_sign_v4_if_in_acdr(explicit_v1_from_endpoint)
        self.assertEqual('v1', explicit_v1_from_endpoint.request.signature_version)
        self.assertEqual('', explicit_v1_from_endpoint.configuration.region_id)

        client.set_sign_v4_if_in_acdr(acdr_region)
        self.assertEqual('v4', acdr_region.request.signature_version)
        self.assertEqual('cn-hangzhou-acdr-ut-1', acdr_region.configuration.region_id)

    def test_modify_configuration(self):
        client = Client()

        acdr = self._context('cn-hangzhou-acdr-ut-1', None)
        client.modify_configuration(acdr, None)
        self.assertEqual('cn-hangzhou-acdr-ut-1.log.aliyuncs.com', acdr.configuration.endpoint)
        self.assertEqual('v4', acdr.request.signature_version)
        self.assertEqual('cn-hangzhou-acdr-ut-1', acdr.configuration.region_id)

        ordinary = self._context(None, None)
        client.modify_configuration(ordinary, None)
        self.assertEqual('cn-hangzhou.log.aliyuncs.com', ordinary.configuration.endpoint)
        self.assertFalse(ordinary.request.signature_version)

        from_endpoint = self._context('', 'cn-hangzhou-acdr-ut-1.sls.aliyuncs.com')
        client.modify_configuration(from_endpoint, None)
        self.assertEqual('cn-hangzhou-acdr-ut-1.sls.aliyuncs.com', from_endpoint.configuration.endpoint)
        self.assertEqual('v4', from_endpoint.request.signature_version)
        self.assertEqual('cn-hangzhou-acdr-ut-1', from_endpoint.configuration.region_id)

        explicit = self._context('cn-hangzhou-acdr-ut-1', 'cn-hangzhou.log.aliyuncs.com', 'v1')
        client.modify_configuration(explicit, None)
        self.assertEqual('v1', explicit.request.signature_version)

        client.modify_configuration(acdr, None)
        self.assertEqual('v4', acdr.request.signature_version)

    def test_modify_configuration_async(self):
        client = Client()
        acdr = self._context('cn-hangzhou-acdr-ut-1', None)
        self._run(client.modify_configuration_async(acdr, None))
        self.assertEqual('cn-hangzhou-acdr-ut-1.log.aliyuncs.com', acdr.configuration.endpoint)
        self.assertEqual('v4', acdr.request.signature_version)

        ordinary = self._context(None, None)
        self._run(client.modify_configuration_async(ordinary, None))
        self.assertEqual('cn-hangzhou.log.aliyuncs.com', ordinary.configuration.endpoint)
        self.assertFalse(ordinary.request.signature_version)

        from_endpoint = self._context('', 'cn-hangzhou-acdr-ut-1.sls.aliyuncs.com')
        self._run(client.modify_configuration_async(from_endpoint, None))
        self.assertEqual('v4', from_endpoint.request.signature_version)
        self.assertEqual('cn-hangzhou-acdr-ut-1', from_endpoint.configuration.region_id)

        explicit = self._context('cn-hangzhou-acdr-ut-1', 'cn-hangzhou.log.aliyuncs.com', 'v1')
        self._run(client.modify_configuration_async(explicit, None))
        self.assertEqual('v1', explicit.request.signature_version)

        self._run(client.modify_configuration_async(acdr, None))
        self.assertEqual('v4', acdr.request.signature_version)

    def _run(self, coro):
        import asyncio
        loop = asyncio.new_event_loop()
        try:
            return loop.run_until_complete(coro)
        finally:
            loop.close()


if __name__ == '__main__':
    unittest.main()
