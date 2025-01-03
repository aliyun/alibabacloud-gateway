// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutAccessPointPublicAccessBlockRequest extends TeaModel {
    /**
     * <p>The container in which the Block Public Access configurations are stored.</p>
     */
    @NameInMap("PublicAccessBlockConfiguration")
    public PublicAccessBlockConfiguration publicAccessBlockConfiguration;

    /**
     * <p>The name of the access point.</p>
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>ap-01</p>
     */
    @NameInMap("x-oss-access-point-name")
    public String xOssAccessPointName;

    public static PutAccessPointPublicAccessBlockRequest build(java.util.Map<String, ?> map) throws Exception {
        PutAccessPointPublicAccessBlockRequest self = new PutAccessPointPublicAccessBlockRequest();
        return TeaModel.build(map, self);
    }

    public PutAccessPointPublicAccessBlockRequest setPublicAccessBlockConfiguration(PublicAccessBlockConfiguration publicAccessBlockConfiguration) {
        this.publicAccessBlockConfiguration = publicAccessBlockConfiguration;
        return this;
    }
    public PublicAccessBlockConfiguration getPublicAccessBlockConfiguration() {
        return this.publicAccessBlockConfiguration;
    }

    public PutAccessPointPublicAccessBlockRequest setXOssAccessPointName(String xOssAccessPointName) {
        this.xOssAccessPointName = xOssAccessPointName;
        return this;
    }
    public String getXOssAccessPointName() {
        return this.xOssAccessPointName;
    }

}
