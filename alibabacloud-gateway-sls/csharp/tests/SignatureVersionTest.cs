using System;
using System.IO;
using System.Threading.Tasks;
using AlibabaCloud.GatewaySls;
using AlibabaCloud.GatewaySpi.Models;
using Newtonsoft.Json.Linq;
using Xunit;

namespace tests
{
    public class SignatureVersionTest
    {
        [Fact]
        public async Task SelectsSignatureVersion()
        {
            var cases = JArray.Parse(File.ReadAllText(Path.Combine(AppContext.BaseDirectory, "signature-version.json")));
            var client = new Client();
            foreach (var test in cases)
            {
                foreach (var asynchronous in new[] { false, true })
                {
                    var context = new InterceptorContext
                    {
                        Request = new InterceptorContext.InterceptorContextRequest { SignatureVersion = (string)test["version"] },
                        Configuration = new InterceptorContext.InterceptorContextConfiguration
                        {
                            RegionId = (string)test["region"],
                            Endpoint = (string)test["endpoint"]
                        }
                    };
                    var version = asynchronous ? await client.GetSignatureVersionAsync(context) : client.GetSignatureVersion(context);
                    Assert.True(version == (string)test["want"], (string)test["name"]);
                    Assert.Equal((string)test["resolved"], context.Configuration.RegionId);
                    Assert.Equal((string)test["version"], context.Request.SignatureVersion);
                }
            }
        }
    }
}
