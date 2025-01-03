// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetObjectGroupIndexResponseBody extends TeaModel {
    @NameInMap("FileGroup")
    public FileGroupInfo fileGroup;

    public static GetObjectGroupIndexResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetObjectGroupIndexResponseBody self = new GetObjectGroupIndexResponseBody();
        return TeaModel.build(map, self);
    }

    public GetObjectGroupIndexResponseBody setFileGroup(FileGroupInfo fileGroup) {
        this.fileGroup = fileGroup;
        return this;
    }
    public FileGroupInfo getFileGroup() {
        return this.fileGroup;
    }

}
