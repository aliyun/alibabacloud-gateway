// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PostObjectGroupRequest extends TeaModel {
    @NameInMap("CreateFileGroup")
    public CreateFileGroup createFileGroup;

    public static PostObjectGroupRequest build(java.util.Map<String, ?> map) throws Exception {
        PostObjectGroupRequest self = new PostObjectGroupRequest();
        return TeaModel.build(map, self);
    }

    public PostObjectGroupRequest setCreateFileGroup(CreateFileGroup createFileGroup) {
        this.createFileGroup = createFileGroup;
        return this;
    }
    public CreateFileGroup getCreateFileGroup() {
        return this.createFileGroup;
    }

}
