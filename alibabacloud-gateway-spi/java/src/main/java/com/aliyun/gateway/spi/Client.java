// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.spi;

import com.aliyun.tea.*;
import com.aliyun.gateway.spi.models.*;
import com.aliyun.credentials.*;

public abstract class Client {

    public Client() throws Exception {
    }

    public abstract void modifyConfiguration(InterceptorContext context, AttributeMap attributeMap) throws Exception;

    public abstract void modifyRequest(InterceptorContext context, AttributeMap attributeMap) throws Exception;

    public abstract void modifyResponse(InterceptorContext context, AttributeMap attributeMap) throws Exception;
}
