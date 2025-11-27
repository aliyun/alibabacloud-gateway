import Client from '../src/client';
import assert from 'assert';
import 'mocha';

describe('Client', function () {

    it('getRegion should ok', function () {
        var client = new Client()
        assert.deepStrictEqual('center', client.getRegion('', '', ''));
        assert.deepStrictEqual('cn-hangzhou', client.getRegion('', '', 'cn-hangzhou'));
        assert.deepStrictEqual('center', client.getRegion('test', 'test', ''));
        assert.deepStrictEqual('center', client.getRegion('test', 'test.alibaba.api.com', ''));
        assert.deepStrictEqual('center', client.getRegion('test', 'test.aliyuncs.com', ''));
        assert.deepStrictEqual('center', client.getRegion('test', 'test-dualstack.aliyuncs.com', ''));
        assert.deepStrictEqual('center', client.getRegion('test', 'test-inner.aliyuncs.com', ''));
        assert.deepStrictEqual('center', client.getRegion('test', 'test-vpc.aliyuncs.com', ''));
        assert.deepStrictEqual('center', client.getRegion('test', 'test-share.aliyuncs.com', ''));
        assert.deepStrictEqual('center', client.getRegion('test', 'test-cn-hangzhou.aliyuncs.com', ''));
        assert.deepStrictEqual('cn-hangzhou', client.getRegion('test', 'test.cn-hangzhou.aliyuncs.com', ''));
        assert.deepStrictEqual('cn-hangzhou', client.getRegion('test', 'test-inner.cn-hangzhou.aliyuncs.com', ''));
        assert.deepStrictEqual('cn-hangzhou', client.getRegion('test', 'test-vpc.cn-hangzhou.aliyuncs.com', ''));
        assert.deepStrictEqual('cn-hangzhou', client.getRegion('test', 'test-share.cn-hangzhou.aliyuncs.com', ''));
        assert.deepStrictEqual('cn-hangzhou', client.getRegion('test', 'test-dualstack.cn-hangzhou.aliyuncs.com', ''));
        assert.deepStrictEqual('cn-hangzhou', client.getRegion('test', 'test-proxy.cn-hangzhou.aliyuncs.com', ''));
        assert.deepStrictEqual('cn-hangzhou-acdr-ut-1', client.getRegion('test', 'test-inner.cn-hangzhou-acdr-ut-1.aliyuncs.com', ''));
        assert.deepStrictEqual('cn-edge-1', client.getRegion('test', 'test-inner.cn-edge-1.aliyuncs.com', ''));
    });

    it('getSignedHeaders should ok', function () {
        var client = new Client();

        // 测试基本场景：只有必须的headers
        const headers1 = {
            'host': 'example.com',
            'Host': 'example.com',
            'content-type': 'application/json'
        };
        assert.deepStrictEqual(['content-type', 'host'], client.getSignedHeaders(headers1));

        // 测试包含acs相关headers的情况
        const headers2 = {
            'host': 'example.com',
            'content-type': 'application/json',
            'x-acs-action': 'TestAction',
            'x-acs-version': '2019-01-01'
        };
        assert.deepStrictEqual(['content-type', 'host', 'x-acs-action', 'x-acs-version'], client.getSignedHeaders(headers2));

        // 测试包含非acs headers的情况
        const headers3 = {
            'host': 'example.com',
            'content-type': 'application/json',
            'user-agent': 'TestAgent',
            'x-custom-header': 'CustomValue'
        };
        assert.deepStrictEqual(['content-type', 'host'], client.getSignedHeaders(headers3)); // 应该只包含规定的headers

        // 测试空headers情况
        const headers4 = {};
        assert.deepStrictEqual([''], client.getSignedHeaders(headers4));

        // 测试包含空值的headers情况
        const headers5 = {
            'host': 'example.com',
            'content-type': 'application/json',
            'x-acs-action': null
        };
        assert.deepStrictEqual(['content-type', 'host'], client.getSignedHeaders(headers5)); // x-acs-action因为值为空被过滤掉

        // 测试大小写不敏感的情况
        const headers6 = {
            'Host': 'example.com',
            'Content-Type': 'application/json',
            'X-Acs-Action': 'TestAction'
        };
        assert.deepStrictEqual(['content-type', 'host', 'x-acs-action'], client.getSignedHeaders(headers6)); // 应该都被转为小写

        // 测试排序情况
        const headers7 = {
            'x-acs-zzz': 'zzz',
            'host': 'example.com',
            'x-acs-aaa': 'aaa',
            'content-type': 'application/json'
        };
        assert.deepStrictEqual(['content-type', 'host', 'x-acs-aaa', 'x-acs-zzz'], client.getSignedHeaders(headers7)); // 应该按字典序排列
    });

});