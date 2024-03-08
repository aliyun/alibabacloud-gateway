// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class GetBucketLifecycleResponseBody extends TeaModel {
    /**
     * <p>The container that stores the lifecycle rules configured for the bucket.</p>
     */
    @NameInMap("LifecycleConfiguration")
    public LifecycleConfiguration lifecycleConfiguration;

    public static GetBucketLifecycleResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketLifecycleResponseBody self = new GetBucketLifecycleResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketLifecycleResponseBody setLifecycleConfiguration(LifecycleConfiguration lifecycleConfiguration) {
        this.lifecycleConfiguration = lifecycleConfiguration;
        return this;
    }
    public LifecycleConfiguration getLifecycleConfiguration() {
        return this.lifecycleConfiguration;
    }

    public static class LifecycleConfiguration extends TeaModel {
        /**
         * <p>The container that stores the lifecycle rules.</p>
         */
        @NameInMap("Rule")
        public java.util.List<LifecycleRule> rules;

        public static LifecycleConfiguration build(java.util.Map<String, ?> map) throws Exception {
            LifecycleConfiguration self = new LifecycleConfiguration();
            return TeaModel.build(map, self);
        }

        public LifecycleConfiguration setRules(java.util.List<LifecycleRule> rules) {
            this.rules = rules;
            return this;
        }
        public java.util.List<LifecycleRule> getRules() {
            return this.rules;
        }

    }

}
