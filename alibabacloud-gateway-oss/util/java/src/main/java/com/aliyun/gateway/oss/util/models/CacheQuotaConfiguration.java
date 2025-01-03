// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CacheQuotaConfiguration extends TeaModel {
    @NameInMap("QuotaDesc")
    public QuotaDesc quotaDesc;

    public static CacheQuotaConfiguration build(java.util.Map<String, ?> map) throws Exception {
        CacheQuotaConfiguration self = new CacheQuotaConfiguration();
        return TeaModel.build(map, self);
    }

    public CacheQuotaConfiguration setQuotaDesc(QuotaDesc quotaDesc) {
        this.quotaDesc = quotaDesc;
        return this;
    }
    public QuotaDesc getQuotaDesc() {
        return this.quotaDesc;
    }

    public static class QuotaDesc extends TeaModel {
        /**
         * <strong>example:</strong>
         * <p>30</p>
         */
        @NameInMap("Quota")
        public Long quota;

        public static QuotaDesc build(java.util.Map<String, ?> map) throws Exception {
            QuotaDesc self = new QuotaDesc();
            return TeaModel.build(map, self);
        }

        public QuotaDesc setQuota(Long quota) {
            this.quota = quota;
            return this;
        }
        public Long getQuota() {
            return this.quota;
        }

    }

}
