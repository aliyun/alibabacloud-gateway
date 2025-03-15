import unittest

from alibabacloud_gateway_pop.client import Client


class TestClient(unittest.TestCase):

    def test_get_region(self):
        client = Client()
        self.assertEqual('center', client.get_region(None, None, None))
        self.assertEqual('center', client.get_region('', '', None))
        self.assertEqual('cn-hangzhou', client.get_region('', '', 'cn-hangzhou'))
        self.assertEqual('center', client.get_region('test', 'test', None))
        self.assertEqual('center', client.get_region('test', 'test.alibaba.api.com', None))
        self.assertEqual('center', client.get_region('test', 'test.aliyuncs.com', None))
        self.assertEqual('center', client.get_region('test', 'test-dualstack.aliyuncs.com', None))
        self.assertEqual('center', client.get_region('test', 'test-inner.aliyuncs.com', None))
        self.assertEqual('center', client.get_region('test', 'test-vpc.aliyuncs.com', None))
        self.assertEqual('center', client.get_region('test', 'test-share.aliyuncs.com', None))
        self.assertEqual('center', client.get_region('test', 'test-cn-hangzhou.aliyuncs.com', None))
        self.assertEqual('cn-hangzhou', client.get_region('test', 'test.cn-hangzhou.aliyuncs.com', None))
        self.assertEqual('cn-hangzhou', client.get_region('test', 'test-inner.cn-hangzhou.aliyuncs.com', None))
        self.assertEqual('cn-hangzhou', client.get_region('test', 'test-vpc.cn-hangzhou.aliyuncs.com', None))
        self.assertEqual('cn-hangzhou', client.get_region('test', 'test-share.cn-hangzhou.aliyuncs.com', None))
        self.assertEqual('cn-hangzhou', client.get_region('test', 'test-dualstack.cn-hangzhou.aliyuncs.com', None))
        self.assertEqual('cn-hangzhou', client.get_region('test', 'test-proxy.cn-hangzhou.aliyuncs.com', None))
        self.assertEqual('cn-hangzhou-acdr-ut-1', client.get_region('test', 'test-inner.cn-hangzhou-acdr-ut-1.aliyuncs.com', None))
        self.assertEqual('cn-edge-1', client.get_region('test', 'test-inner.cn-edge-1.aliyuncs.com', None))
