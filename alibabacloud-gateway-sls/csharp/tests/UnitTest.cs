using AlibabaCloud.GatewaySls;
using AlibabaCloud.GatewaySpi.Models;
using Xunit;

namespace tests
{
    public class UnitTest
    {
        [Fact]
        public void Test_ParseRegion()
        {
            Client client = new Client();
            Assert.Equal("", client.ParseRegion(""));
            Assert.Equal("", client.ParseRegion(null));
            Assert.Equal("cn-hangzhou-acdr-ut-1", client.ParseRegion("cn-hangzhou-acdr-ut-1.sls.aliyuncs.com"));
            Assert.Equal("cn-hangzhou-acdr-ut-1", client.ParseRegion("cn-hangzhou-acdr-ut-1.log.aliyuncs.com"));
            Assert.Equal("cn-hangzhou-acdr-ut-1", client.ParseRegion("https://cn-hangzhou-acdr-ut-1.sls.aliyuncs.com"));
            Assert.Equal("cn-hangzhou-acdr-ut-1", client.ParseRegion("http://cn-hangzhou-acdr-ut-1.log.aliyuncs.com"));
            Assert.Equal("cn-hangzhou-acdr-ut-1", client.ParseRegion("cn-hangzhou-acdr-ut-1-intranet.sls.aliyuncs.com"));
            Assert.Equal("cn-hangzhou-acdr-ut-1", client.ParseRegion("cn-hangzhou-acdr-ut-1-share.log.aliyuncs.com"));
            Assert.Equal("cn-hangzhou-acdr-ut-1", client.ParseRegion("cn-hangzhou-acdr-ut-1-vpc.sls.aliyuncs.com"));
            Assert.Equal("cn-hangzhou-acdr-ut-1", client.ParseRegion("cn-hangzhou-acdr-ut-1-internal.log.aliyuncs.com"));
            Assert.Equal("cn-hangzhou-acdr-ut-1", client.ParseRegion("https://cn-hangzhou-acdr-ut-1-intranet.log.aliyuncs.com"));
            Assert.Equal("cn-hangzhou", client.ParseRegion("cn-hangzhou.sls.aliyuncs.com"));
            Assert.Equal("cn-hangzhou", client.ParseRegion("cn-hangzhou.log.aliyuncs.com"));
            Assert.Equal("", client.ParseRegion("ftp://cn-hangzhou-acdr-ut-1.sls.aliyuncs.com"));
            Assert.Equal("", client.ParseRegion("cn-hangzhou-acdr-ut-1.oss.aliyuncs.com"));
            Assert.Equal("", client.ParseRegion("cn-hangzhou-acdr-ut-1.sls.aliyuncs.com.cn"));
            Assert.Equal("", client.ParseRegion("cn-hangzhou-acdr-ut-1.sls.aliyuncs.com/path"));
            Assert.Equal("", client.ParseRegion("https://https://cn-hangzhou-acdr-ut-1.sls.aliyuncs.com"));
            Assert.Equal("", client.ParseRegion("CN-HANGZHOU.sls.aliyuncs.com"));
            Assert.Equal("", client.ParseRegion("cn_hangzhou.sls.aliyuncs.com"));
            Assert.Equal("cn-hangzhou-acdr-ut-1-intranet", client.ParseRegion("cn-hangzhou-acdr-ut-1-intranet-share.sls.aliyuncs.com"));
        }

        [Fact]
        public void Test_SetSignV4IfInAcdr()
        {
            Client client = new Client();

            var acdrRegion = Context("cn-hangzhou-acdr-ut-1", "cn-hangzhou.log.aliyuncs.com", null);
            client.SetSignV4IfInAcdr(acdrRegion);
            Assert.Equal("v4", acdrRegion.Request.SignatureVersion);
            Assert.Equal("cn-hangzhou-acdr-ut-1", acdrRegion.Configuration.RegionId);

            var ordinaryRegion = Context("cn-hangzhou", "cn-hangzhou.log.aliyuncs.com", null);
            client.SetSignV4IfInAcdr(ordinaryRegion);
            Assert.Null(ordinaryRegion.Request.SignatureVersion);

            var fromEndpoint = Context(null, "cn-hangzhou-acdr-ut-1.sls.aliyuncs.com", null);
            client.SetSignV4IfInAcdr(fromEndpoint);
            Assert.Equal("v4", fromEndpoint.Request.SignatureVersion);
            Assert.Equal("cn-hangzhou-acdr-ut-1", fromEndpoint.Configuration.RegionId);

            var ordinaryEndpoint = Context("", "cn-hangzhou.sls.aliyuncs.com", null);
            client.SetSignV4IfInAcdr(ordinaryEndpoint);
            Assert.Null(ordinaryEndpoint.Request.SignatureVersion);
            Assert.Equal("", ordinaryEndpoint.Configuration.RegionId);

            var emptyEndpoint = Context(null, null, null);
            client.SetSignV4IfInAcdr(emptyEndpoint);
            Assert.Null(emptyEndpoint.Request.SignatureVersion);

            var explicitV1 = Context("cn-hangzhou-acdr-ut-1", "cn-hangzhou-acdr-ut-1.sls.aliyuncs.com", "v1");
            client.SetSignV4IfInAcdr(explicitV1);
            Assert.Equal("v1", explicitV1.Request.SignatureVersion);

            var explicitV4 = Context("cn-hangzhou", "cn-hangzhou.log.aliyuncs.com", "v4");
            client.SetSignV4IfInAcdr(explicitV4);
            Assert.Equal("v4", explicitV4.Request.SignatureVersion);

            var intranet = Context("", "https://cn-hangzhou-acdr-ut-1-intranet.log.aliyuncs.com", null);
            client.SetSignV4IfInAcdr(intranet);
            Assert.Equal("v4", intranet.Request.SignatureVersion);
            Assert.Equal("cn-hangzhou-acdr-ut-1", intranet.Configuration.RegionId);

            var invalidAcdrHint = Context(null, "not-a-valid-acdr-ut-host", null);
            client.SetSignV4IfInAcdr(invalidAcdrHint);
            Assert.Null(invalidAcdrHint.Request.SignatureVersion);
            Assert.Null(invalidAcdrHint.Configuration.RegionId);

            var shanghai = Context("cn-shanghai-acdr-ut-2", null, null);
            client.SetSignV4IfInAcdr(shanghai);
            Assert.Equal("v4", shanghai.Request.SignatureVersion);

            var explicitV1FromEndpoint = Context("", "cn-hangzhou-acdr-ut-1.sls.aliyuncs.com", "v1");
            client.SetSignV4IfInAcdr(explicitV1FromEndpoint);
            Assert.Equal("v1", explicitV1FromEndpoint.Request.SignatureVersion);
            Assert.Equal("", explicitV1FromEndpoint.Configuration.RegionId);

            client.SetSignV4IfInAcdr(acdrRegion);
            Assert.Equal("v4", acdrRegion.Request.SignatureVersion);
            Assert.Equal("cn-hangzhou-acdr-ut-1", acdrRegion.Configuration.RegionId);
        }

        [Fact]
        public void Test_ModifyConfiguration()
        {
            Client client = new Client();

            var acdr = Context("cn-hangzhou-acdr-ut-1", null, null);
            client.ModifyConfiguration(acdr, null);
            Assert.Equal("cn-hangzhou-acdr-ut-1.log.aliyuncs.com", acdr.Configuration.Endpoint);
            Assert.Equal("v4", acdr.Request.SignatureVersion);
            Assert.Equal("cn-hangzhou-acdr-ut-1", acdr.Configuration.RegionId);

            var ordinary = Context(null, null, null);
            client.ModifyConfiguration(ordinary, null);
            Assert.Equal("cn-hangzhou.log.aliyuncs.com", ordinary.Configuration.Endpoint);
            Assert.Null(ordinary.Request.SignatureVersion);

            var fromEndpoint = Context("", "cn-hangzhou-acdr-ut-1.sls.aliyuncs.com", null);
            client.ModifyConfiguration(fromEndpoint, null);
            Assert.Equal("cn-hangzhou-acdr-ut-1.sls.aliyuncs.com", fromEndpoint.Configuration.Endpoint);
            Assert.Equal("v4", fromEndpoint.Request.SignatureVersion);
            Assert.Equal("cn-hangzhou-acdr-ut-1", fromEndpoint.Configuration.RegionId);

            var explicitV1 = Context("cn-hangzhou-acdr-ut-1", "cn-hangzhou.log.aliyuncs.com", "v1");
            client.ModifyConfiguration(explicitV1, null);
            Assert.Equal("v1", explicitV1.Request.SignatureVersion);

            client.ModifyConfiguration(acdr, null);
            Assert.Equal("v4", acdr.Request.SignatureVersion);
        }

        [Fact]
        public async System.Threading.Tasks.Task Test_ModifyConfigurationAsync()
        {
            Client client = new Client();

            var acdr = Context("cn-hangzhou-acdr-ut-1", null, null);
            await client.ModifyConfigurationAsync(acdr, null);
            Assert.Equal("cn-hangzhou-acdr-ut-1.log.aliyuncs.com", acdr.Configuration.Endpoint);
            Assert.Equal("v4", acdr.Request.SignatureVersion);

            var ordinary = Context(null, null, null);
            await client.ModifyConfigurationAsync(ordinary, null);
            Assert.Equal("cn-hangzhou.log.aliyuncs.com", ordinary.Configuration.Endpoint);
            Assert.Null(ordinary.Request.SignatureVersion);

            var fromEndpoint = Context("", "cn-hangzhou-acdr-ut-1.sls.aliyuncs.com", null);
            await client.ModifyConfigurationAsync(fromEndpoint, null);
            Assert.Equal("v4", fromEndpoint.Request.SignatureVersion);
            Assert.Equal("cn-hangzhou-acdr-ut-1", fromEndpoint.Configuration.RegionId);

            var explicitV1 = Context("cn-hangzhou-acdr-ut-1", "cn-hangzhou.log.aliyuncs.com", "v1");
            await client.ModifyConfigurationAsync(explicitV1, null);
            Assert.Equal("v1", explicitV1.Request.SignatureVersion);

            await client.ModifyConfigurationAsync(acdr, null);
            Assert.Equal("v4", acdr.Request.SignatureVersion);
        }

        private static InterceptorContext Context(string regionId, string endpoint, string signatureVersion)
        {
            return new InterceptorContext
            {
                Request = new InterceptorContext.InterceptorContextRequest
                {
                    SignatureVersion = signatureVersion
                },
                Configuration = new InterceptorContext.InterceptorContextConfiguration
                {
                    RegionId = regionId,
                    Endpoint = endpoint
                }
            };
        }
    }
}
