using System;
using System.Collections.Generic;
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

        [Fact]
        public void Test_GetSignedHeaders()
        {
            Client client = new Client();

            // 测试空headers
            var emptyHeaders = new Dictionary<string, string>();
            var result = client.GetSignedHeaders(emptyHeaders);
            Assert.Equal(1, result.Count);

            // 测试只包含需要签名的headers
            var headers = new Dictionary<string, string>
            {
                { "host", "example.com" },
                { "Host", "example.com" },
                { "content-type", "application/json" },
                { "x-acs-action", "TestAction" }
            };
            result = client.GetSignedHeaders(headers);
            Assert.Equal(3, result.Count);
            Assert.Equal("content-type", result[0]);
            Assert.Equal("host", result[1]);
            Assert.Equal("x-acs-action", result[2]);

            // 测试包含不需要签名的headers
            var headersWithIgnore = new Dictionary<string, string>
            {
                { "host", "example.com" },
                { "content-type", "application/json" },
                { "x-acs-action", "TestAction" },
                { "authorization", "Bearer token" },
                { "user-agent", "C#-client" }
            };
            result = client.GetSignedHeaders(headersWithIgnore);
            Assert.Equal(3, result.Count);
            Assert.Equal("content-type", result[0]);
            Assert.Equal("host", result[1]);
            Assert.Equal("x-acs-action", result[2]);

            // 测试空值header
            var headersWithEmpty = new Dictionary<string, string>
            {
                { "host", "example.com" },
                { "content-type", null },
                { "x-acs-action", "TestAction" }
            };
            result = client.GetSignedHeaders(headersWithEmpty);
            Assert.Equal(2, result.Count);
            Assert.Equal("host", result[0]);
            Assert.Equal("x-acs-action", result[1]);

            // 测试大小写不敏感
            var headersWithCase = new Dictionary<string, string>
            {
                { "HOST", "example.com" },
                { "Content-Type", "application/json" },
                { "X-Acs-Action", "TestAction" }
            };
            result = client.GetSignedHeaders(headersWithCase);
            Assert.Equal(3, result.Count);
            Assert.Equal("content-type", result[0]);
            Assert.Equal("host", result[1]);
            Assert.Equal("x-acs-action", result[2]);
        }
    }
}