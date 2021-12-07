// This file is auto-generated, don't edit it. Thanks.

using System;
using System.Collections.Generic;
using System.IO;

using Tea;

namespace AlibabaCloud.GatewaySpi.Models
{
    public class AttributeMap : TeaModel {
        [NameInMap("attributes")]
        [Validation(Required=true)]
        public Dictionary<string, object> Attributes { get; set; }

        [NameInMap("key")]
        [Validation(Required=true)]
        public Dictionary<string, string> Key { get; set; }

    }

}
