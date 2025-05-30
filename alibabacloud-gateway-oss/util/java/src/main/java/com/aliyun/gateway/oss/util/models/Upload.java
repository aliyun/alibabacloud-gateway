// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class Upload extends TeaModel {
    /**
     * <p>Use the UTC time format: yyyy-MM-ddTHH:mmZ</p>
     */
    @NameInMap("Initiated")
    public String initiated;

    @NameInMap("Key")
    public String key;

    @NameInMap("UploadId")
    public String uploadId;

    public static Upload build(java.util.Map<String, ?> map) throws Exception {
        Upload self = new Upload();
        return TeaModel.build(map, self);
    }

    public Upload setInitiated(String initiated) {
        this.initiated = initiated;
        return this;
    }
    public String getInitiated() {
        return this.initiated;
    }

    public Upload setKey(String key) {
        this.key = key;
        return this;
    }
    public String getKey() {
        return this.key;
    }

    public Upload setUploadId(String uploadId) {
        this.uploadId = uploadId;
        return this;
    }
    public String getUploadId() {
        return this.uploadId;
    }

}
