# -*- coding: utf-8 -*-
import json
import os
import unittest

import requests
from alibabacloud_credentials import credentials
from alibabacloud_gateway_fc_util.client import Client


class TestClient(unittest.TestCase):

    def test_invoke_httptrigger(self):
        ak = os.getenv('ak')
        sk = os.getenv('sk')
        url = os.getenv('url')
        d = {
            'type': 'access_key',
            'access_key_id': ak,
            'access_key_secret': sk,
        }
        cred = credentials.AccessKeyCredential(access_key_id=ak, access_key_secret=sk)
        client = Client(cred=cred)
        resp = client.invoke_httptrigger(url=url,
                                         method="GET", body="anything".encode(encoding='utf-8'),
                                         headers={"k1": "v1", "k2": "v2"})
        TestClient.pretty_print(resp)
        self.assertEqual(resp.status_code, 200)

    @staticmethod
    def pretty_print(resp: requests.Response):
        """
        At this point it is completely built and ready
        to be fired; it is "prepared".

        However pay attention at the formatting used in
        this function because it is programmed to be pretty
        printed and may differ from the actual request.
        """
        print('{}\n{}\r\n{}\r\n\r\n{}'.format(
            '-----------START-----------',
            resp.url,
            '\r\n'.join('{}: {}'.format(k, v) for k, v in resp.headers.items()),
            json.dumps(resp.json(), indent=2),
        ))


if __name__ == '__main__':
    unittest.main()
