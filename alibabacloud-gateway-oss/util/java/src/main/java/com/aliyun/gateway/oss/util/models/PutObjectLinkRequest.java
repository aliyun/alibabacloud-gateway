// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutObjectLinkRequest extends TeaModel {
    @NameInMap("CreateObjectLink")
    public ObjectLinkInfo createObjectLink;

    public static PutObjectLinkRequest build(java.util.Map<String, ?> map) throws Exception {
        PutObjectLinkRequest self = new PutObjectLinkRequest();
        return TeaModel.build(map, self);
    }

    public PutObjectLinkRequest setCreateObjectLink(ObjectLinkInfo createObjectLink) {
        this.createObjectLink = createObjectLink;
        return this;
    }
    public ObjectLinkInfo getCreateObjectLink() {
        return this.createObjectLink;
    }

}
