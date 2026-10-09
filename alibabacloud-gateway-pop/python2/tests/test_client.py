# -*- coding: utf-8 -*-
from __future__ import unicode_literals

import unittest

from Tea.exceptions import TeaException
from alibabacloud_gateway_pop.client import Client


class TestClient(unittest.TestCase):

    def test_validate_signed_headers(self):
        client = Client()
        client.validate_signed_headers(['content-type', 'host', 'x-acs-date'])
        with self.assertRaises(TeaException) as ctx:
            client.validate_signed_headers(['host'])
        self.assertEqual('InvalidSignedHeaders', ctx.exception.code)
        with self.assertRaises(TeaException):
            client.validate_signed_headers(['x-acs-date'])
        with self.assertRaises(TeaException):
            client.validate_signed_headers([])

        signingkey = client.get_signingkey(client._sha_256, 'secret', 'ecs', 'cn-hangzhou', '20240101')
        with self.assertRaises(TeaException):
            client.get_signature('/', 'GET', {}, {'host': 'example.com'}, client._sha_256, '', signingkey)
        self.assertTrue(client.get_signature('/', 'GET', {}, {
            'host': 'example.com',
            'x-acs-date': '2024-01-01T00:00:00Z',
        }, client._sha_256, '', signingkey))


if __name__ == '__main__':
    unittest.main()
