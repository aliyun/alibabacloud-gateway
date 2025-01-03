// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutCacheRequest extends TeaModel {
    @NameInMap("CreateCacheConfiguration")
    public CreateCacheConfiguration createCacheConfiguration;

    public static PutCacheRequest build(java.util.Map<String, ?> map) throws Exception {
        PutCacheRequest self = new PutCacheRequest();
        return TeaModel.build(map, self);
    }

    public PutCacheRequest setCreateCacheConfiguration(CreateCacheConfiguration createCacheConfiguration) {
        this.createCacheConfiguration = createCacheConfiguration;
        return this;
    }
    public CreateCacheConfiguration getCreateCacheConfiguration() {
        return this.createCacheConfiguration;
    }

}
