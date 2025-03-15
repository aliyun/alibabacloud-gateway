using System;
using AlibabaCloud.GatewayPop;

using Xunit;

namespace tests
{
    public class UnitTest
    {
        [Fact]
        public void Test_GetRegion()
        {
            Client client = new Client();
            Assert.Equal("center", client.GetRegion(null, null, null));
            Assert.Equal("cn-hangzhou", client.GetRegion(null, null, "cn-hangzhou"));
            Assert.Equal("center", client.GetRegion("", "", null));
            Assert.Equal("center", client.GetRegion("test", "test", null));
            Assert.Equal("center", client.GetRegion("test", "test.alibaba.api.com", null));
            Assert.Equal("center", client.GetRegion("test", "test.aliyuncs.com", null));
            Assert.Equal("center", client.GetRegion("test", "test-dualstack.aliyuncs.com", null));
            Assert.Equal("center", client.GetRegion("test", "test-inner.aliyuncs.com", null));
            Assert.Equal("center", client.GetRegion("test", "test-vpc.aliyuncs.com", null));
            Assert.Equal("center", client.GetRegion("test", "test-share.aliyuncs.com", null));
            Assert.Equal("center", client.GetRegion("test", "test-cn-hangzhou.aliyuncs.com", null));
            Assert.Equal("cn-hangzhou", client.GetRegion("test", "test.cn-hangzhou.aliyuncs.com", null));
            Assert.Equal("cn-hangzhou", client.GetRegion("test", "test-inner.cn-hangzhou.aliyuncs.com", null));
            Assert.Equal("cn-hangzhou", client.GetRegion("test", "test-vpc.cn-hangzhou.aliyuncs.com", null));
            Assert.Equal("cn-hangzhou", client.GetRegion("test", "test-share.cn-hangzhou.aliyuncs.com", null));
            Assert.Equal("cn-hangzhou", client.GetRegion("test", "test-dualstack.cn-hangzhou.aliyuncs.com", null));
            Assert.Equal("cn-hangzhou", client.GetRegion("test", "test-proxy.cn-hangzhou.aliyuncs.com", null));
            Assert.Equal("cn-hangzhou-acdr-ut-1", client.GetRegion("test", "test-inner.cn-hangzhou-acdr-ut-1.aliyuncs.com", null));
            Assert.Equal("cn-edge-1", client.GetRegion("test", "test-inner.cn-edge-1.aliyuncs.com", null));
        }
    }
}
