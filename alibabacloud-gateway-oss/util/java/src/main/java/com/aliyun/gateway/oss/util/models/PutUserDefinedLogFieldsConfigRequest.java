// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutUserDefinedLogFieldsConfigRequest extends TeaModel {
    @NameInMap("UserDefinedLogFieldsConfiguration")
    public UserDefinedLogFieldsConfiguration userDefinedLogFieldsConfiguration;

    public static PutUserDefinedLogFieldsConfigRequest build(java.util.Map<String, ?> map) throws Exception {
        PutUserDefinedLogFieldsConfigRequest self = new PutUserDefinedLogFieldsConfigRequest();
        return TeaModel.build(map, self);
    }

    public PutUserDefinedLogFieldsConfigRequest setUserDefinedLogFieldsConfiguration(UserDefinedLogFieldsConfiguration userDefinedLogFieldsConfiguration) {
        this.userDefinedLogFieldsConfiguration = userDefinedLogFieldsConfiguration;
        return this;
    }
    public UserDefinedLogFieldsConfiguration getUserDefinedLogFieldsConfiguration() {
        return this.userDefinedLogFieldsConfiguration;
    }

}
