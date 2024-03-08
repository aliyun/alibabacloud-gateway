// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class BucketAntiDDOSConfiguration extends TeaModel {
    @NameInMap("Cnames")
    public BucketAntiDDOSConfigurationCnames cnames;

    public static BucketAntiDDOSConfiguration build(java.util.Map<String, ?> map) throws Exception {
        BucketAntiDDOSConfiguration self = new BucketAntiDDOSConfiguration();
        return TeaModel.build(map, self);
    }

    public BucketAntiDDOSConfiguration setCnames(BucketAntiDDOSConfigurationCnames cnames) {
        this.cnames = cnames;
        return this;
    }
    public BucketAntiDDOSConfigurationCnames getCnames() {
        return this.cnames;
    }

    public static class BucketAntiDDOSConfigurationCnames extends TeaModel {
        @NameInMap("Domain")
        public java.util.List<String> domain;

        public static BucketAntiDDOSConfigurationCnames build(java.util.Map<String, ?> map) throws Exception {
            BucketAntiDDOSConfigurationCnames self = new BucketAntiDDOSConfigurationCnames();
            return TeaModel.build(map, self);
        }

        public BucketAntiDDOSConfigurationCnames setDomain(java.util.List<String> domain) {
            this.domain = domain;
            return this;
        }
        public java.util.List<String> getDomain() {
            return this.domain;
        }

    }

}
