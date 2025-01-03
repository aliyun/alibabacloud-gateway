// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PostObjectGroupResponseBody extends TeaModel {
    @NameInMap("CreateFileGroup")
    public CreateFileGroupResult createFileGroup;

    public static PostObjectGroupResponseBody build(java.util.Map<String, ?> map) throws Exception {
        PostObjectGroupResponseBody self = new PostObjectGroupResponseBody();
        return TeaModel.build(map, self);
    }

    public PostObjectGroupResponseBody setCreateFileGroup(CreateFileGroupResult createFileGroup) {
        this.createFileGroup = createFileGroup;
        return this;
    }
    public CreateFileGroupResult getCreateFileGroup() {
        return this.createFileGroup;
    }

}
