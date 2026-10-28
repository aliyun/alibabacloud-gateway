package com.aliyun.gateway.pop;

import com.aliyun.endpointutil.Client;
import org.junit.Assert;
import org.junit.Test;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

public class SupportParityTest {
    @Test
    public void endpointAndFormMatchPreviousLibraries() throws Exception {
        Assert.assertEquals(
                Client.getEndpointRules("Ecs", "cn-hangzhou", "regional", "", ""),
                Support.getEndpointRules("Ecs", "cn-hangzhou", "regional", "", ""));
        Assert.assertEquals(
                Client.getEndpointRules("Ecs", "cn-hangzhou", "regional", "vpc", "inner"),
                Support.getEndpointRules("Ecs", "cn-hangzhou", "regional", "vpc", "inner"));
        Assert.assertEquals(
                Client.getEndpointRules("Ecs", "", "central", "public", null),
                Support.getEndpointRules("Ecs", "", "central", "public", null));
        try {
            Support.getEndpointRules("Ecs", "", "regional", null, null);
            Assert.fail("regional endpoint without region should fail");
        } catch (RuntimeException e) {
            Assert.assertEquals(
                    "RegionId is empty, please set a valid RegionId",
                    e.getMessage());
        }

        Map<String, Object> filter = new HashMap<String, Object>();
        filter.put("b", "x y");
        filter.put("a", Arrays.asList("m", "n"));
        Map<String, Object> nested = new HashMap<String, Object>();
        nested.put("k", "v");
        filter.put("c", nested);
        Assert.assertEquals(com.aliyun.openapiutil.Client.toForm(filter), Support.toForm(filter));
        Assert.assertEquals(com.aliyun.teautil.Common.empty(""), Support.empty(""));
        Assert.assertEquals(com.aliyun.teautil.Common.equalString("v4", "v4"), Support.equalString("v4", "v4"));
        Assert.assertEquals(com.aliyun.teautil.Common.is4xx(404), Support.is4xx(404));
        Assert.assertEquals(com.aliyun.teautil.Common.isUnset(null), Support.isUnset(null));
        Map<String, String> html = new HashMap<String, String>();
        html.put("a", "<b>");
        Assert.assertEquals(com.aliyun.teautil.Common.toJSONString(html), Support.toJSONString(html));
    }
}
