package com.aliyun.gateway.oss;

import org.junit.Assert;
import org.junit.Test;

public class UnitTest {

    @Test
    public void getEndpointTest() throws Exception {
        Client client = new Client();

        Assert.assertEquals("custom.endpoint.com",
            client.getEndpoint("cn-hangzhou", null, "custom.endpoint.com"));
        Assert.assertEquals("custom.endpoint.com",
            client.getEndpoint("cn-hangzhou", "ipv6", "custom.endpoint.com"));

        Assert.assertEquals("oss-cn-shanghai.aliyuncs.com",
            client.getEndpoint("cn-shanghai", null, null));
        Assert.assertEquals("oss-cn-hangzhou.aliyuncs.com",
            client.getEndpoint(null, null, null));
        Assert.assertEquals("oss-cn-hangzhou.aliyuncs.com",
            client.getEndpoint("", "", ""));

        Assert.assertEquals("oss-cn-hangzhou-internal.aliyuncs.com",
            client.getEndpoint("cn-hangzhou", "internal", null));
        Assert.assertEquals("oss-cn-beijing-internal.aliyuncs.com",
            client.getEndpoint("cn-beijing", "vpc-internal", null));

        Assert.assertEquals("cn-hangzhou.oss.aliyuncs.com",
            client.getEndpoint("cn-hangzhou", "ipv6", null));
        Assert.assertEquals("cn-shanghai.oss.aliyuncs.com",
            client.getEndpoint("cn-shanghai", "dualstack-ipv6", null));
        Assert.assertEquals("cn-hangzhou.oss.aliyuncs.com",
            client.getEndpoint(null, "ipv6", null));

        Assert.assertEquals("oss-accelerate.aliyuncs.com",
            client.getEndpoint("cn-hangzhou", "accelerate", null));
        Assert.assertEquals("oss-accelerate-overseas.aliyuncs.com",
            client.getEndpoint("ap-southeast-1", "accelerate-overseas", null));
    }
}
