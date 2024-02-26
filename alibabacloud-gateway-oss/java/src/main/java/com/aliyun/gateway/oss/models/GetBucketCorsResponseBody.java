// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class GetBucketCorsResponseBody extends TeaModel {
    /**
     * <p>The container that stores CORS configuration.</p>
     */
    @NameInMap("CORSConfiguration")
    public CORSConfiguration CORSConfiguration;

    public static GetBucketCorsResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketCorsResponseBody self = new GetBucketCorsResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketCorsResponseBody setCORSConfiguration(CORSConfiguration CORSConfiguration) {
        this.CORSConfiguration = CORSConfiguration;
        return this;
    }
    public CORSConfiguration getCORSConfiguration() {
        return this.CORSConfiguration;
    }

    public static class CORSConfiguration extends TeaModel {
        /**
         * <p>The container that stores CORS rules. Up to 10 rules can be configured for a bucket.</p>
         */
        @NameInMap("CORSRule")
        public java.util.List<CORSRule> CORSRule;

        /**
         * <p>Indicates whether the Vary: Origin header was returned. Default value: false.</p>
         * <p>- true: The Vary: Origin header is returned regardless whether the request is a cross-origin request or whether the cross-origin request succeeds.</p>
         * <p>- false: The Vary: Origin header is not returned.</p>
         */
        @NameInMap("ResponseVary")
        public Boolean responseVary;

        public static CORSConfiguration build(java.util.Map<String, ?> map) throws Exception {
            CORSConfiguration self = new CORSConfiguration();
            return TeaModel.build(map, self);
        }

        public CORSConfiguration setCORSRule(java.util.List<CORSRule> CORSRule) {
            this.CORSRule = CORSRule;
            return this;
        }
        public java.util.List<CORSRule> getCORSRule() {
            return this.CORSRule;
        }

        public CORSConfiguration setResponseVary(Boolean responseVary) {
            this.responseVary = responseVary;
            return this;
        }
        public Boolean getResponseVary() {
            return this.responseVary;
        }

    }

}
