// This file is auto-generated, don't edit it. Thanks.

using System;
using System.Collections.Generic;
using System.IO;

using Tea;

namespace AlibabaCloud.GatewayFc.Models
{
    public class HttpRequest : TeaModel {
        [NameInMap("method")]
        [Validation(Required=true)]
        public string Method { get; set; }

        [NameInMap("path")]
        [Validation(Required=true)]
        public string Path { get; set; }

        [NameInMap("headers")]
        [Validation(Required=false)]
        public Dictionary<string, object> Headers { get; set; }

        [NameInMap("body")]
        [Validation(Required=false)]
        public byte[] Body { get; set; }

        [NameInMap("reqBodyType")]
        [Validation(Required=false)]
        public string ReqBodyType { get; set; }

    }

}
