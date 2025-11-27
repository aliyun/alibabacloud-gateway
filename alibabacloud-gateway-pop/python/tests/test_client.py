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

    def test_get_signed_headers(self):
        client = Client()
        
        # 测试基本场景：只有必须的headers
        headers1 = {
            'host': 'example.com',
            'Host': 'example.com',
            'content-type': 'application/json'
        }
        result1 = client.get_signed_headers(headers1)
        self.assertEqual(['content-type', 'host'], result1)
        
        # 测试包含acs相关headers的情况
        headers2 = {
            'host': 'example.com',
            'content-type': 'application/json',
            'x-acs-action': 'TestAction',
            'x-acs-version': '2019-01-01'
        }
        result2 = client.get_signed_headers(headers2)
        self.assertEqual(['content-type', 'host', 'x-acs-action', 'x-acs-version'], result2)
        
        # 测试包含非acs headers的情况
        headers3 = {
            'host': 'example.com',
            'content-type': 'application/json',
            'user-agent': 'TestAgent',
            'x-custom-header': 'CustomValue'
        }
        result3 = client.get_signed_headers(headers3)
        self.assertEqual(['content-type', 'host'], result3)  # 应该只包含规定的headers
        
        # 测试空headers情况
        headers4 = {}
        result4 = client.get_signed_headers(headers4)
        self.assertEqual([''], result4)
        
        # 测试包含空值的headers情况
        headers5 = {
            'host': 'example.com',
            'content-type': 'application/json',
            'x-acs-action': None
        }
        result5 = client.get_signed_headers(headers5)
        self.assertEqual(['content-type', 'host'], result5)  # x-acs-action因为值为空被过滤掉
        
        # 测试大小写不敏感的情况
        headers6 = {
            'Host': 'example.com',
            'Content-Type': 'application/json',
            'X-Acs-Action': 'TestAction'
        }
        result6 = client.get_signed_headers(headers6)
        self.assertEqual(['content-type', 'host', 'x-acs-action'], result6)  # 应该都被转为小写
        
        # 测试排序情况
        headers7 = {
            'x-acs-zzz': 'zzz',
            'host': 'example.com',
            'x-acs-aaa': 'aaa',
            'content-type': 'application/json'
        }
        result7 = client.get_signed_headers(headers7)
        self.assertEqual(['content-type', 'host', 'x-acs-aaa', 'x-acs-zzz'], result7)  # 应该按字典序排列