// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class StartPartUploadResponseBody extends TeaModel {
    @NameInMap("StartPartUploadResult")
    public StartPartUploadResult startPartUploadResult;

    public static StartPartUploadResponseBody build(java.util.Map<String, ?> map) throws Exception {
        StartPartUploadResponseBody self = new StartPartUploadResponseBody();
        return TeaModel.build(map, self);
    }

    public StartPartUploadResponseBody setStartPartUploadResult(StartPartUploadResult startPartUploadResult) {
        this.startPartUploadResult = startPartUploadResult;
        return this;
    }
    public StartPartUploadResult getStartPartUploadResult() {
        return this.startPartUploadResult;
    }

}
