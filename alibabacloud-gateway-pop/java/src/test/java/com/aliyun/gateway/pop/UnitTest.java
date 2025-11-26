package com.aliyun.gateway.pop;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

import org.junit.Assert;
import org.junit.Test;

public class UnitTest {
    @Test
    public void getRegionTest() throws Exception {
        Client client = new Client();
        Assert.assertEquals("center", client.getRegion(null, null, null));
        Assert.assertEquals("cn-hangzhou", client.getRegion(null, null, "cn-hangzhou"));
        Assert.assertEquals("center", client.getRegion("", "", null));
        Assert.assertEquals("center", client.getRegion("test", "test", null));
        Assert.assertEquals("center", client.getRegion("test", "test.alibaba.api.com", null));
        Assert.assertEquals("center", client.getRegion("test", "test.aliyuncs.com", null));
        Assert.assertEquals("center", client.getRegion("test", "test-dualstack.aliyuncs.com", null));
        Assert.assertEquals("center", client.getRegion("test", "test-inner.aliyuncs.com", null));
        Assert.assertEquals("center", client.getRegion("test", "test-vpc.aliyuncs.com", null));
        Assert.assertEquals("center", client.getRegion("test", "test-share.aliyuncs.com", null));
        Assert.assertEquals("center", client.getRegion("test", "test-cn-hangzhou.aliyuncs.com", null));
        Assert.assertEquals("cn-hangzhou", client.getRegion("test", "test.cn-hangzhou.aliyuncs.com", null));
        Assert.assertEquals("cn-hangzhou", client.getRegion("test", "test-inner.cn-hangzhou.aliyuncs.com", null));
        Assert.assertEquals("cn-hangzhou", client.getRegion("test", "test-vpc.cn-hangzhou.aliyuncs.com", null));
        Assert.assertEquals("cn-hangzhou", client.getRegion("test", "test-share.cn-hangzhou.aliyuncs.com", null));
        Assert.assertEquals("cn-hangzhou", client.getRegion("test", "test-dualstack.cn-hangzhou.aliyuncs.com", null));
        Assert.assertEquals("cn-hangzhou", client.getRegion("test", "test-proxy.cn-hangzhou.aliyuncs.com", null));
        Assert.assertEquals("cn-hangzhou-acdr-ut-1", client.getRegion("test", "test-inner.cn-hangzhou-acdr-ut-1.aliyuncs.com", null));
        Assert.assertEquals("cn-edge-1", client.getRegion("test", "test-inner.cn-edge-1.aliyuncs.com", null));
    }

    @Test
    public void getSignedHeadersTest() throws Exception {
        Client client = new Client();

        // 测试空headers
        Map<String, String> emptyHeaders = new HashMap<>();
        List<String> result = client.getSignedHeaders(emptyHeaders);
        Assert.assertEquals(1, result.size());

        // 测试只包含需要签名的headers
        Map<String, String> headers = new HashMap<>();
        headers.put("host", "example.com");
        headers.put("Host", "example.com");
        headers.put("content-type", "application/json");
        headers.put("x-acs-action", "TestAction");
        result = client.getSignedHeaders(headers);
        Assert.assertEquals(3, result.size());
        Assert.assertEquals("content-type", result.get(0));
        Assert.assertEquals("host", result.get(1));
        Assert.assertEquals("x-acs-action", result.get(2));

        // 测试包含不需要签名的headers
        Map<String, String> headersWithIgnore = new HashMap<>();
        headersWithIgnore.put("host", "example.com");
        headersWithIgnore.put("content-type", "application/json");
        headersWithIgnore.put("x-acs-action", "TestAction");
        headersWithIgnore.put("authorization", "Bearer token");
        headersWithIgnore.put("user-agent", "Java-client");
        result = client.getSignedHeaders(headersWithIgnore);
        Assert.assertEquals(3, result.size());
        Assert.assertEquals("content-type", result.get(0));
        Assert.assertEquals("host", result.get(1));
        Assert.assertEquals("x-acs-action", result.get(2));

        // 测试空值header
        Map<String, String> headersWithEmpty = new HashMap<>();
        headersWithEmpty.put("host", "example.com");
        headersWithEmpty.put("content-type", null);
        headersWithEmpty.put("x-acs-action", "TestAction");
        result = client.getSignedHeaders(headersWithEmpty);
        Assert.assertEquals(2, result.size());
        Assert.assertEquals("host", result.get(0));
        Assert.assertEquals("x-acs-action", result.get(1));

        // 测试大小写不敏感
        Map<String, String> headersWithCase = new HashMap<>();
        headersWithCase.put("HOST", "example.com");
        headersWithCase.put("Content-Type", "application/json");
        headersWithCase.put("X-Acs-Action", "TestAction");
        result = client.getSignedHeaders(headersWithCase);
        Assert.assertEquals(3, result.size());
        Assert.assertEquals("content-type", result.get(0));
        Assert.assertEquals("host", result.get(1));
        Assert.assertEquals("x-acs-action", result.get(2));
    }
}