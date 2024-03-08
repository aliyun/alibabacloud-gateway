// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class BucketCnameConfiguration extends TeaModel {
    @NameInMap("Cname")
    public BucketCnameConfigurationCname cname;

    public static BucketCnameConfiguration build(java.util.Map<String, ?> map) throws Exception {
        BucketCnameConfiguration self = new BucketCnameConfiguration();
        return TeaModel.build(map, self);
    }

    public BucketCnameConfiguration setCname(BucketCnameConfigurationCname cname) {
        this.cname = cname;
        return this;
    }
    public BucketCnameConfigurationCname getCname() {
        return this.cname;
    }

    public static class BucketCnameConfigurationCname extends TeaModel {
        @NameInMap("Domain")
        public String domain;

        public static BucketCnameConfigurationCname build(java.util.Map<String, ?> map) throws Exception {
            BucketCnameConfigurationCname self = new BucketCnameConfigurationCname();
            return TeaModel.build(map, self);
        }

        public BucketCnameConfigurationCname setDomain(String domain) {
            this.domain = domain;
            return this;
        }
        public String getDomain() {
            return this.domain;
        }

    }

}
