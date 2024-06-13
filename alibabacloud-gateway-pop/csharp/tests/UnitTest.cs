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
            Assert.Equal("center", client.GetRegion(null, null));
            Assert.Equal("center", client.GetRegion("", ""));
            Assert.Equal("center", client.GetRegion("test", "test"));
            Assert.Equal("center", client.GetRegion("test", "test.alibaba.api.com"));
            Assert.Equal("center", client.GetRegion("test", "test.aliyuncs.com"));
            Assert.Equal("center", client.GetRegion("test", "test-dualstack.aliyuncs.com"));
            Assert.Equal("center", client.GetRegion("test", "test-inner.aliyuncs.com"));
            Assert.Equal("center", client.GetRegion("test", "test-vpc.aliyuncs.com"));
            Assert.Equal("center", client.GetRegion("test", "test-share.aliyuncs.com"));
            Assert.Equal("center", client.GetRegion("test", "test-cn-hangzhou.aliyuncs.com"));
            Assert.Equal("cn-hangzhou", client.GetRegion("test", "test.cn-hangzhou.aliyuncs.com"));
            Assert.Equal("cn-hangzhou", client.GetRegion("test", "test-inner.cn-hangzhou.aliyuncs.com"));
            Assert.Equal("cn-hangzhou", client.GetRegion("test", "test-vpc.cn-hangzhou.aliyuncs.com"));
            Assert.Equal("cn-hangzhou", client.GetRegion("test", "test-share.cn-hangzhou.aliyuncs.com"));
            Assert.Equal("cn-hangzhou", client.GetRegion("test", "test-dualstack.cn-hangzhou.aliyuncs.com"));
            Assert.Equal("cn-hangzhou", client.GetRegion("test", "test-proxy.cn-hangzhou.aliyuncs.com"));
            Assert.Equal("cn-hangzhou-acdr-ut-1", client.GetRegion("test", "test-inner.cn-hangzhou-acdr-ut-1.aliyuncs.com"));
            Assert.Equal("cn-edge-1", client.GetRegion("test", "test-inner.cn-edge-1.aliyuncs.com"));
        }
    }
}
