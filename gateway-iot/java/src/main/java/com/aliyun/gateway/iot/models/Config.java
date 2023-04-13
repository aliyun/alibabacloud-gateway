// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.iot.models;

import com.aliyun.tea.*;

public class Config extends TeaModel {
    @NameInMap("appKey")
    @Validation(required = true)
    public String appKey;

    @NameInMap("appSecret")
    @Validation(required = true)
    public String appSecret;

    @NameInMap("endpoint")
    @Validation(required = true)
    public String endpoint;

    public static Config build(java.util.Map<String, ?> map) throws Exception {
        Config self = new Config();
        return TeaModel.build(map, self);
    }

    public Config setAppKey(String appKey) {
        this.appKey = appKey;
        return this;
    }
    public String getAppKey() {
        return this.appKey;
    }

    public Config setAppSecret(String appSecret) {
        this.appSecret = appSecret;
        return this;
    }
    public String getAppSecret() {
        return this.appSecret;
    }

    public Config setEndpoint(String endpoint) {
        this.endpoint = endpoint;
        return this;
    }
    public String getEndpoint() {
        return this.endpoint;
    }

}
