// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.oss20190517.models;

import com.aliyun.tea.*;

public class PutBucketObjectWormConfigurationRequest extends TeaModel {
    @NameInMap("body")
    public Body body;

    public static PutBucketObjectWormConfigurationRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketObjectWormConfigurationRequest self = new PutBucketObjectWormConfigurationRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketObjectWormConfigurationRequest setBody(Body body) {
        this.body = body;
        return this;
    }
    public Body getBody() {
        return this.body;
    }

    public static class Body extends TeaModel {
        @NameInMap("ObjectWormConfiguration")
        public ObjectWormConfiguration objectWormConfiguration;

        public static Body build(java.util.Map<String, ?> map) throws Exception {
            Body self = new Body();
            return TeaModel.build(map, self);
        }

        public Body setObjectWormConfiguration(ObjectWormConfiguration objectWormConfiguration) {
            this.objectWormConfiguration = objectWormConfiguration;
            return this;
        }
        public ObjectWormConfiguration getObjectWormConfiguration() {
            return this.objectWormConfiguration;
        }

    }

}
