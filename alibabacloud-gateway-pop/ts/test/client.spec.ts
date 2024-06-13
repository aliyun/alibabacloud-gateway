import Client from '../src/client';
import assert from 'assert';
import 'mocha';

describe('Client', function () {

    it('getRegion should ok', function () {
        var client = new Client()
        assert.deepStrictEqual('center', client.getRegion('', ''));
        assert.deepStrictEqual('center', client.getRegion('test', 'test'));
        assert.deepStrictEqual('center', client.getRegion('test', 'test.alibaba.api.com'));
        assert.deepStrictEqual('center', client.getRegion('test', 'test.aliyuncs.com'));
        assert.deepStrictEqual('center', client.getRegion('test', 'test-dualstack.aliyuncs.com'));
        assert.deepStrictEqual('center', client.getRegion('test', 'test-inner.aliyuncs.com'));
        assert.deepStrictEqual('center', client.getRegion('test', 'test-vpc.aliyuncs.com'));
        assert.deepStrictEqual('center', client.getRegion('test', 'test-share.aliyuncs.com'));
        assert.deepStrictEqual('center', client.getRegion('test', 'test-cn-hangzhou.aliyuncs.com'));
        assert.deepStrictEqual('cn-hangzhou', client.getRegion('test', 'test.cn-hangzhou.aliyuncs.com'));
        assert.deepStrictEqual('cn-hangzhou', client.getRegion('test', 'test-inner.cn-hangzhou.aliyuncs.com'));
        assert.deepStrictEqual('cn-hangzhou', client.getRegion('test', 'test-vpc.cn-hangzhou.aliyuncs.com'));
        assert.deepStrictEqual('cn-hangzhou', client.getRegion('test', 'test-share.cn-hangzhou.aliyuncs.com'));
        assert.deepStrictEqual('cn-hangzhou', client.getRegion('test', 'test-dualstack.cn-hangzhou.aliyuncs.com'));
        assert.deepStrictEqual('cn-hangzhou', client.getRegion('test', 'test-proxy.cn-hangzhou.aliyuncs.com'));
        assert.deepStrictEqual('cn-hangzhou-acdr-ut-1', client.getRegion('test', 'test-inner.cn-hangzhou-acdr-ut-1.aliyuncs.com'));
        assert.deepStrictEqual('cn-edge-1', client.getRegion('test', 'test-inner.cn-edge-1.aliyuncs.com'));
    });

});
