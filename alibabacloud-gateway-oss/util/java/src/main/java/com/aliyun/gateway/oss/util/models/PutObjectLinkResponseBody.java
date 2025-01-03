// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutObjectLinkResponseBody extends TeaModel {
    @NameInMap("CreateObjectLink")
    public CreateObjectLinkResult createObjectLink;

    public static PutObjectLinkResponseBody build(java.util.Map<String, ?> map) throws Exception {
        PutObjectLinkResponseBody self = new PutObjectLinkResponseBody();
        return TeaModel.build(map, self);
    }

    public PutObjectLinkResponseBody setCreateObjectLink(CreateObjectLinkResult createObjectLink) {
        this.createObjectLink = createObjectLink;
        return this;
    }
    public CreateObjectLinkResult getCreateObjectLink() {
        return this.createObjectLink;
    }

}
