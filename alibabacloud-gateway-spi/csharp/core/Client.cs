// This file is auto-generated, don't edit it. Thanks.

using System;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;

using Tea;
using Tea.Utils;

using AlibabaCloud.GatewaySpi.Models;

namespace AlibabaCloud.GatewaySpi
{
    public interface Client 
    {
        void ModifyConfiguration(InterceptorContext context, AttributeMap attributeMap);

        Task ModifyConfigurationAsync(InterceptorContext context, AttributeMap attributeMap);

        void ModifyRequest(InterceptorContext context, AttributeMap attributeMap);

        Task ModifyRequestAsync(InterceptorContext context, AttributeMap attributeMap);

        void ModifyResponse(InterceptorContext context, AttributeMap attributeMap);

        Task ModifyResponseAsync(InterceptorContext context, AttributeMap attributeMap);

    }
}
