// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetUserDefinedLogFieldsConfigResponseBody extends TeaModel {
    @NameInMap("UserDefinedLogFieldsConfiguration")
    public UserDefinedLogFieldsConfiguration userDefinedLogFieldsConfiguration;

    public static GetUserDefinedLogFieldsConfigResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetUserDefinedLogFieldsConfigResponseBody self = new GetUserDefinedLogFieldsConfigResponseBody();
        return TeaModel.build(map, self);
    }

    public GetUserDefinedLogFieldsConfigResponseBody setUserDefinedLogFieldsConfiguration(UserDefinedLogFieldsConfiguration userDefinedLogFieldsConfiguration) {
        this.userDefinedLogFieldsConfiguration = userDefinedLogFieldsConfiguration;
        return this;
    }
    public UserDefinedLogFieldsConfiguration getUserDefinedLogFieldsConfiguration() {
        return this.userDefinedLogFieldsConfiguration;
    }

}
