import unittest
from urllib.parse import quote

from Tea.model import TeaModel

from alibabacloud_gateway_pop import _codec as codec, _support as support


class Body(TeaModel):
    def to_map(self):
        return {'Name': 'demo'}


class Outer(TeaModel):
    def to_map(self):
        return {'Item': Body()}


class TestSupport(unittest.TestCase):
    def test_percent_encode(self):
        for raw in ['', 'a b*~', '中文', 'a+b', '100%', '/', '~']:
            self.assertEqual(quote(raw, safe='~', encoding='utf-8'), codec.percent_encode(raw))

    def test_endpoint_and_form_match_previous_libraries(self):
        self.assertEqual(
            'ecs.cn-hangzhou.aliyuncs.com',
            support.get_endpoint_rules('Ecs', 'cn-hangzhou', 'regional', '', ''))
        self.assertEqual(
            'ecs.aliyuncs.com',
            support.get_endpoint_rules('Ecs', '', 'central', 'public', None))
        self.assertEqual('a.1=m&a.2=n&b=x+y&c.k=v', support.to_form({'b': 'x y', 'a': ['m', 'n'], 'c': {'k': 'v'}}))
        self.assertEqual('{"a":"<b>"}', support.to_jsonstring({'a': '<b>'}))
        self.assertEqual('{"Item":{"Name":"demo"}}', support.to_jsonstring({'Item': Body()}))
        self.assertEqual('{"Item":{"Name":"demo"}}', support.to_jsonstring(Outer()))


if __name__ == '__main__':
    unittest.main()
