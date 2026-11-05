package com.aliyun.gateway.pop;

import java.util.HashMap;
import java.util.Map;
import org.junit.Assert;
import org.junit.Test;

public class ExplicitEndpointTest {
    @Test
    public void rewritesExplicitAndMappedEndpoints() throws Exception {
        Client client = new Client();
        String[][] cases = {
            {"oss-cn-hangzhou.aliyuncs.com", "oss-cn-hangzhou.aliyuncs.alibaba"},
            {"ecs.cn-hangzhou-internal.aliyuncs.com", "ecs.cn-hangzhou-internal.aliyuncs.alibaba"},
            {"ecs.aliyun.com", "ecs.aliyun.com"},
            {"custom.example.com", "custom.example.com"},
            {"ecs.aliyuncs.alibaba", "ecs.aliyuncs.alibaba"},
            {"ecs.aliyuncs.com.example.com", "ecs.aliyuncs.com.example.com"}
        };
        for (String[] item : cases) {
            Map<String, String> endpoints = new HashMap<String, String>();
            endpoints.put("cn-hangzhou", item[0]);
            Assert.assertEquals(item[1], client.getEndpoint("ecs", "cn-hangzhou", "regional", "", "", null, item[0], "aliyuncs.alibaba"));
            Assert.assertEquals(item[1], client.getEndpoint("ecs", "cn-hangzhou", "regional", "", "", endpoints, "", "aliyuncs.alibaba"));
            Assert.assertEquals(item[1], client.getEndpoint("ecs", "cn-hangzhou", "regional", "", "", null, item[1], "aliyuncs.alibaba"));
            Assert.assertEquals(item[0], client.getEndpoint("ecs", "cn-hangzhou", "regional", "", "", null, item[0], "aliyuncs.com"));
        }
    }
}
