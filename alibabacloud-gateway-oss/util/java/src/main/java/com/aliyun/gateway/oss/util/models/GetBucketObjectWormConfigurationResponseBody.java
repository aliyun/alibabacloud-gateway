// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.oss20190517.models;

import com.aliyun.tea.*;

public class GetBucketObjectWormConfigurationResponseBody extends TeaModel {
    @NameInMap("ObjectWormConfiguration")
    public ObjectWormConfiguration objectWormConfiguration;

    public static GetBucketObjectWormConfigurationResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketObjectWormConfigurationResponseBody self = new GetBucketObjectWormConfigurationResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketObjectWormConfigurationResponseBody setObjectWormConfiguration(ObjectWormConfiguration objectWormConfiguration) {
        this.objectWormConfiguration = objectWormConfiguration;
        return this;
    }
    public ObjectWormConfiguration getObjectWormConfiguration() {
        return this.objectWormConfiguration;
    }

}
