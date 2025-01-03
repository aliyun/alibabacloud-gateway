// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetObjectLinkResponseBody extends TeaModel {
    @NameInMap("ObjectLink")
    public ObjectLinkInfo objectLink;

    public static GetObjectLinkResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetObjectLinkResponseBody self = new GetObjectLinkResponseBody();
        return TeaModel.build(map, self);
    }

    public GetObjectLinkResponseBody setObjectLink(ObjectLinkInfo objectLink) {
        this.objectLink = objectLink;
        return this;
    }
    public ObjectLinkInfo getObjectLink() {
        return this.objectLink;
    }

}
