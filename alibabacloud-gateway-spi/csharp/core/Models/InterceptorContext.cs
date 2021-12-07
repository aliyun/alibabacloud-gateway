// This file is auto-generated, don't edit it. Thanks.

using System;
using System.Collections.Generic;
using System.IO;

using Tea;

namespace AlibabaCloud.GatewaySpi.Models
{
    public class InterceptorContext : TeaModel {
        [NameInMap("request")]
        [Validation(Required=true)]
        public InterceptorContextRequest Request { get; set; }
        public class InterceptorContextRequest : TeaModel {
            [NameInMap("headers")]
            [Validation(Required=false)]
            public Dictionary<string, string> Headers { get; set; }
            [NameInMap("query")]
            [Validation(Required=false)]
            public Dictionary<string, string> Query { get; set; }
            [NameInMap("body")]
            [Validation(Required=false)]
            public object Body { get; set; }
            [NameInMap("stream")]
            [Validation(Required=false)]
            public Stream Stream { get; set; }
            [NameInMap("hostMap")]
            [Validation(Required=false)]
            public Dictionary<string, string> HostMap { get; set; }
            [NameInMap("pathname")]
            [Validation(Required=true)]
            public string Pathname { get; set; }
            [NameInMap("productId")]
            [Validation(Required=true)]
            public string ProductId { get; set; }
            [NameInMap("action")]
            [Validation(Required=true)]
            public string Action { get; set; }
            [NameInMap("version")]
            [Validation(Required=true)]
            public string Version { get; set; }
            [NameInMap("protocol")]
            [Validation(Required=true)]
            public string Protocol { get; set; }
            [NameInMap("method")]
            [Validation(Required=true)]
            public string Method { get; set; }
            [NameInMap("authType")]
            [Validation(Required=true)]
            public string AuthType { get; set; }
            [NameInMap("bodyType")]
            [Validation(Required=true)]
            public string BodyType { get; set; }
            [NameInMap("reqBodyType")]
            [Validation(Required=true)]
            public string ReqBodyType { get; set; }
            [NameInMap("style")]
            [Validation(Required=false)]
            public string Style { get; set; }
            [NameInMap("credential")]
            [Validation(Required=true)]
            public Aliyun.Credentials.Client Credential { get; set; }
            [NameInMap("signatureVersion")]
            [Validation(Required=false)]
            public string SignatureVersion { get; set; }
            [NameInMap("signatureAlgorithm")]
            [Validation(Required=false)]
            public string SignatureAlgorithm { get; set; }
            [NameInMap("userAgent")]
            [Validation(Required=true)]
            public string UserAgent { get; set; }
        };

        [NameInMap("configuration")]
        [Validation(Required=true)]
        public InterceptorContextConfiguration Configuration { get; set; }
        public class InterceptorContextConfiguration : TeaModel {
            [NameInMap("regionId")]
            [Validation(Required=true)]
            public string RegionId { get; set; }
            [NameInMap("endpoint")]
            [Validation(Required=false)]
            public string Endpoint { get; set; }
            [NameInMap("endpointRule")]
            [Validation(Required=false)]
            public string EndpointRule { get; set; }
            [NameInMap("endpointMap")]
            [Validation(Required=false)]
            public Dictionary<string, string> EndpointMap { get; set; }
            [NameInMap("endpointType")]
            [Validation(Required=false)]
            public string EndpointType { get; set; }
            [NameInMap("network")]
            [Validation(Required=false)]
            public string Network { get; set; }
            [NameInMap("suffix")]
            [Validation(Required=false)]
            public string Suffix { get; set; }
        };

        [NameInMap("response")]
        [Validation(Required=true)]
        public InterceptorContextResponse Response { get; set; }
        public class InterceptorContextResponse : TeaModel {
            [NameInMap("statusCode")]
            [Validation(Required=false)]
            public int? StatusCode { get; set; }
            [NameInMap("headers")]
            [Validation(Required=false)]
            public Dictionary<string, string> Headers { get; set; }
            [NameInMap("body")]
            [Validation(Required=false)]
            public Stream Body { get; set; }
            [NameInMap("deserializedBody")]
            [Validation(Required=false)]
            public object DeserializedBody { get; set; }
        };

    }

}
