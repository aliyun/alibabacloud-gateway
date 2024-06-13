import unittest

from alibabacloud_gateway_pop.client import Client


class TestClient(unittest.TestCase):

    def test_get_region(self):
        client = Client()
        self.assertEqual('center', client.get_region(None, None))
        self.assertEqual('center', client.get_region('', ''))
        self.assertEqual('center', client.get_region('test', 'test'))
        self.assertEqual('center', client.get_region('test', 'test.alibaba.api.com'))
        self.assertEqual('center', client.get_region('test', 'test.aliyuncs.com'))
        self.assertEqual('center', client.get_region('test', 'test-dualstack.aliyuncs.com'))
        self.assertEqual('center', client.get_region('test', 'test-inner.aliyuncs.com'))
        self.assertEqual('center', client.get_region('test', 'test-vpc.aliyuncs.com'))
        self.assertEqual('center', client.get_region('test', 'test-share.aliyuncs.com'))
        self.assertEqual('center', client.get_region('test', 'test-cn-hangzhou.aliyuncs.com'))
        self.assertEqual('cn-hangzhou', client.get_region('test', 'test.cn-hangzhou.aliyuncs.com'))
        self.assertEqual('cn-hangzhou', client.get_region('test', 'test-inner.cn-hangzhou.aliyuncs.com'))
        self.assertEqual('cn-hangzhou', client.get_region('test', 'test-vpc.cn-hangzhou.aliyuncs.com'))
        self.assertEqual('cn-hangzhou', client.get_region('test', 'test-share.cn-hangzhou.aliyuncs.com'))
        self.assertEqual('cn-hangzhou', client.get_region('test', 'test-dualstack.cn-hangzhou.aliyuncs.com'))
        self.assertEqual('cn-hangzhou', client.get_region('test', 'test-proxy.cn-hangzhou.aliyuncs.com'))
        self.assertEqual('cn-hangzhou-acdr-ut-1', client.get_region('test', 'test-inner.cn-hangzhou-acdr-ut-1.aliyuncs.com'))
        self.assertEqual('cn-edge-1', client.get_region('test', 'test-inner.cn-edge-1.aliyuncs.com'))
