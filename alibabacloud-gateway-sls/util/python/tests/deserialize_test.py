import unittest
from aliyun_log_fastpb import serialize_log_group, deserialize_log_group_list


def wrap_as_log_group_list(log_group_bytes: bytes) -> bytes:
    length = len(log_group_bytes)
    varint = b''
    v = length
    while v > 0x7f:
        varint += bytes([0x80 | (v & 0x7f)])
        v >>= 7
    varint += bytes([v & 0x7f])
    return b'\x0a' + varint + log_group_bytes


def make_log_group_list_bytes(*log_groups) -> bytes:
    result = b''
    for lg in log_groups:
        group_bytes = serialize_log_group(lg)
        result += wrap_as_log_group_list(group_bytes)
    return result


class TestDeserializeLogGroupList(unittest.TestCase):

    def test_empty_bytes(self):
        result = deserialize_log_group_list(b'')
        self.assertEqual(result, {'logGroupList': []})

    def test_single_log_group_single_log(self):
        log_group = {
            'Topic': 'test-topic',
            'Source': '10.0.0.1',
            'LogTags': [{'Key': 'env', 'Value': 'prod'}],
            'LogItems': [{
                'Time': 1700000000,
                'Contents': [
                    {'Key': 'level', 'Value': 'INFO'},
                    {'Key': 'message', 'Value': 'hello'},
                ],
            }],
        }
        data = make_log_group_list_bytes(log_group)
        result = deserialize_log_group_list(data)

        groups = result['logGroupList']
        self.assertEqual(len(groups), 1)
        g = groups[0]
        self.assertEqual(g['Topic'], 'test-topic')
        self.assertEqual(g['Source'], '10.0.0.1')
        self.assertEqual(g['LogTags'], [{'Key': 'env', 'Value': 'prod'}])
        self.assertEqual(len(g['LogItems']), 1)
        self.assertEqual(g['LogItems'][0]['Time'], 1700000000)
        self.assertEqual(g['LogItems'][0]['Contents'], [
            {'Key': 'level', 'Value': 'INFO'},
            {'Key': 'message', 'Value': 'hello'},
        ])

    def test_multiple_log_groups(self):
        lg1 = {
            'Topic': 'topic-a',
            'Source': '1.1.1.1',
            'LogItems': [{'Time': 100, 'Contents': [{'Key': 'k', 'Value': 'v1'}]}],
        }
        lg2 = {
            'Topic': 'topic-b',
            'Source': '2.2.2.2',
            'LogItems': [{'Time': 200, 'Contents': [{'Key': 'k', 'Value': 'v2'}]}],
        }
        data = make_log_group_list_bytes(lg1, lg2)
        result = deserialize_log_group_list(data)

        groups = result['logGroupList']
        self.assertEqual(len(groups), 2)
        self.assertEqual(groups[0]['Topic'], 'topic-a')
        self.assertEqual(groups[1]['Topic'], 'topic-b')

    def test_multiple_logs_in_one_group(self):
        log_group = {
            'LogItems': [
                {'Time': 1000, 'Contents': [{'Key': 'seq', 'Value': '1'}]},
                {'Time': 1001, 'Contents': [{'Key': 'seq', 'Value': '2'}]},
                {'Time': 1002, 'Contents': [{'Key': 'seq', 'Value': '3'}]},
            ],
        }
        data = make_log_group_list_bytes(log_group)
        result = deserialize_log_group_list(data)

        items = result['logGroupList'][0]['LogItems']
        self.assertEqual(len(items), 3)
        self.assertEqual([i['Time'] for i in items], [1000, 1001, 1002])

    def test_multiple_contents_in_one_log(self):
        log_group = {
            'LogItems': [{
                'Time': 500,
                'Contents': [
                    {'Key': 'a', 'Value': '1'},
                    {'Key': 'b', 'Value': '2'},
                    {'Key': 'c', 'Value': '3'},
                    {'Key': 'd', 'Value': '4'},
                ],
            }],
        }
        data = make_log_group_list_bytes(log_group)
        result = deserialize_log_group_list(data)

        contents = result['logGroupList'][0]['LogItems'][0]['Contents']
        self.assertEqual(len(contents), 4)
        self.assertEqual(
            {c['Key']: c['Value'] for c in contents},
            {'a': '1', 'b': '2', 'c': '3', 'd': '4'},
        )

    def test_time_ns(self):
        log_group = {
            'LogItems': [{
                'Time': 1700000000,
                'TimeNs': 123456789,
                'Contents': [{'Key': 'k', 'Value': 'v'}],
            }],
        }
        data = make_log_group_list_bytes(log_group)
        result = deserialize_log_group_list(data)

        item = result['logGroupList'][0]['LogItems'][0]
        self.assertEqual(item['Time'], 1700000000)
        self.assertEqual(item['TimeNs'], 123456789)

    def test_time_ns_absent(self):
        log_group = {
            'LogItems': [{
                'Time': 1700000000,
                'Contents': [{'Key': 'k', 'Value': 'v'}],
            }],
        }
        data = make_log_group_list_bytes(log_group)
        result = deserialize_log_group_list(data)

        item = result['logGroupList'][0]['LogItems'][0]
        self.assertEqual(item['Time'], 1700000000)
        self.assertNotIn('TimeNs', item)

    def test_unicode_content(self):
        log_group = {
            'LogItems': [{
                'Time': 100,
                'Contents': [
                    {'Key': '中文键', 'Value': '中文值'},
                    {'Key': 'emoji', 'Value': '🚀🎉'},
                    {'Key': 'japanese', 'Value': 'こんにちは'},
                ],
            }],
        }
        data = make_log_group_list_bytes(log_group)
        result = deserialize_log_group_list(data)

        contents = result['logGroupList'][0]['LogItems'][0]['Contents']
        kv = {c['Key']: c['Value'] for c in contents}
        self.assertEqual(kv['中文键'], '中文值')
        self.assertEqual(kv['emoji'], '🚀🎉')
        self.assertEqual(kv['japanese'], 'こんにちは')

    def test_empty_log_group(self):
        log_group = {}
        data = make_log_group_list_bytes(log_group)
        result = deserialize_log_group_list(data)

        g = result['logGroupList'][0]
        self.assertEqual(g['LogItems'], [])
        self.assertNotIn('Topic', g)
        self.assertNotIn('Source', g)

    def test_no_log_tags_key_absent(self):
        log_group = {
            'LogItems': [{'Time': 1, 'Contents': [{'Key': 'k', 'Value': 'v'}]}],
        }
        data = make_log_group_list_bytes(log_group)
        result = deserialize_log_group_list(data)

        self.assertNotIn('LogTags', result['logGroupList'][0])

    def test_multiple_log_tags(self):
        log_group = {
            'LogTags': [
                {'Key': 'host', 'Value': 'server-1'},
                {'Key': 'region', 'Value': 'cn-hangzhou'},
                {'Key': 'app', 'Value': 'gateway'},
            ],
            'LogItems': [{'Time': 1, 'Contents': [{'Key': 'k', 'Value': 'v'}]}],
        }
        data = make_log_group_list_bytes(log_group)
        result = deserialize_log_group_list(data)

        tags = result['logGroupList'][0]['LogTags']
        self.assertEqual(len(tags), 3)
        tag_map = {t['Key']: t['Value'] for t in tags}
        self.assertEqual(tag_map['host'], 'server-1')
        self.assertEqual(tag_map['region'], 'cn-hangzhou')
        self.assertEqual(tag_map['app'], 'gateway')

    def test_round_trip_preserves_data(self):
        log_group = {
            'Topic': 'round-trip',
            'Source': '192.168.1.100',
            'LogTags': [
                {'Key': 'tag1', 'Value': 'val1'},
                {'Key': 'tag2', 'Value': 'val2'},
            ],
            'LogItems': [
                {
                    'Time': 1700000000,
                    'TimeNs': 999999999,
                    'Contents': [
                        {'Key': 'level', 'Value': 'ERROR'},
                        {'Key': 'msg', 'Value': 'something went wrong'},
                    ],
                },
                {
                    'Time': 1700000001,
                    'Contents': [
                        {'Key': 'level', 'Value': 'INFO'},
                        {'Key': 'msg', 'Value': 'recovered'},
                    ],
                },
            ],
        }
        data = make_log_group_list_bytes(log_group)
        result = deserialize_log_group_list(data)

        g = result['logGroupList'][0]
        self.assertEqual(g['Topic'], log_group['Topic'])
        self.assertEqual(g['Source'], log_group['Source'])
        self.assertEqual(len(g['LogTags']), 2)
        self.assertEqual(len(g['LogItems']), 2)
        self.assertEqual(g['LogItems'][0]['TimeNs'], 999999999)
        self.assertNotIn('TimeNs', g['LogItems'][1])

    def test_invalid_protobuf(self):
        with self.assertRaises(ValueError):
            deserialize_log_group_list(b'\xff\xff\xff')

    def test_large_batch(self):
        log_group = {
            'Topic': 'perf',
            'LogItems': [
                {
                    'Time': 1000 + i,
                    'Contents': [
                        {'Key': 'idx', 'Value': str(i)},
                        {'Key': 'data', 'Value': 'x' * 100},
                    ],
                }
                for i in range(1000)
            ],
        }
        data = make_log_group_list_bytes(log_group)
        result = deserialize_log_group_list(data)

        items = result['logGroupList'][0]['LogItems']
        self.assertEqual(len(items), 1000)
        self.assertEqual(items[0]['Time'], 1000)
        self.assertEqual(items[999]['Time'], 1999)

    def test_empty_string_values(self):
        log_group = {
            'Topic': '',
            'Source': '',
            'LogItems': [{
                'Time': 1,
                'Contents': [{'Key': '', 'Value': ''}],
            }],
        }
        data = make_log_group_list_bytes(log_group)
        result = deserialize_log_group_list(data)

        g = result['logGroupList'][0]
        self.assertEqual(g['Topic'], '')
        self.assertEqual(g['Source'], '')
        self.assertEqual(g['LogItems'][0]['Contents'][0]['Key'], '')
        self.assertEqual(g['LogItems'][0]['Contents'][0]['Value'], '')


if __name__ == '__main__':
    unittest.main()
