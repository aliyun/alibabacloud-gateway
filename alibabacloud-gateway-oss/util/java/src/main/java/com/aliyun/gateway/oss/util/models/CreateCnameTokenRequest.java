// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CreateCnameTokenRequest extends TeaModel {
    /**
     * <p>The container that stores the CNAME record.</p>
     */
    @NameInMap("BucketCnameConfiguration")
    public BucketCnameConfiguration bucketCnameConfiguration;

    public static CreateCnameTokenRequest build(java.util.Map<String, ?> map) throws Exception {
        CreateCnameTokenRequest self = new CreateCnameTokenRequest();
        return TeaModel.build(map, self);
    }

    public CreateCnameTokenRequest setBucketCnameConfiguration(BucketCnameConfiguration bucketCnameConfiguration) {
        this.bucketCnameConfiguration = bucketCnameConfiguration;
        return this;
    }
    public BucketCnameConfiguration getBucketCnameConfiguration() {
        return this.bucketCnameConfiguration;
    }

    public static class Cname extends TeaModel {
        /**
         * <p>The custom domain name.</p>
         */
        @NameInMap("Domain")
        public String domain;

        public static Cname build(java.util.Map<String, ?> map) throws Exception {
            Cname self = new Cname();
            return TeaModel.build(map, self);
        }

        public Cname setDomain(String domain) {
            this.domain = domain;
            return this;
        }
        public String getDomain() {
            return this.domain;
        }

    }

    public static class BucketCnameConfiguration extends TeaModel {
        /**
         * <p>The container that stores the CNAME information.</p>
         */
        @NameInMap("Cname")
        public Cname cname;

        public static BucketCnameConfiguration build(java.util.Map<String, ?> map) throws Exception {
            BucketCnameConfiguration self = new BucketCnameConfiguration();
            return TeaModel.build(map, self);
        }

        public BucketCnameConfiguration setCname(Cname cname) {
            this.cname = cname;
            return this;
        }
        public Cname getCname() {
            return this.cname;
        }

    }

}
