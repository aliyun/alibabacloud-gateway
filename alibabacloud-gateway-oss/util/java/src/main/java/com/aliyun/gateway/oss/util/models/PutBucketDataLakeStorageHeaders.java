// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutBucketDataLakeStorageHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("x-oss-dls-status")
    public String xOssDlsStatus;

    public static PutBucketDataLakeStorageHeaders build(java.util.Map<String, ?> map) throws Exception {
        PutBucketDataLakeStorageHeaders self = new PutBucketDataLakeStorageHeaders();
        return TeaModel.build(map, self);
    }

    public PutBucketDataLakeStorageHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public PutBucketDataLakeStorageHeaders setXOssDlsStatus(String xOssDlsStatus) {
        this.xOssDlsStatus = xOssDlsStatus;
        return this;
    }
    public String getXOssDlsStatus() {
        return this.xOssDlsStatus;
    }

}
