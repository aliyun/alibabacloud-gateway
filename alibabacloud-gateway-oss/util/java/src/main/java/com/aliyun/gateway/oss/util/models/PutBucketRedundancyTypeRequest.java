// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutBucketRedundancyTypeRequest extends TeaModel {
    @NameInMap("DataRedundancyTypeConfiguration")
    public DataRedundancyTypeConfiguration dataRedundancyTypeConfiguration;

    public static PutBucketRedundancyTypeRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketRedundancyTypeRequest self = new PutBucketRedundancyTypeRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketRedundancyTypeRequest setDataRedundancyTypeConfiguration(DataRedundancyTypeConfiguration dataRedundancyTypeConfiguration) {
        this.dataRedundancyTypeConfiguration = dataRedundancyTypeConfiguration;
        return this;
    }
    public DataRedundancyTypeConfiguration getDataRedundancyTypeConfiguration() {
        return this.dataRedundancyTypeConfiguration;
    }

    public static class DataRedundancyTypeConfiguration extends TeaModel {
        @NameInMap("Type")
        public String type;

        public static DataRedundancyTypeConfiguration build(java.util.Map<String, ?> map) throws Exception {
            DataRedundancyTypeConfiguration self = new DataRedundancyTypeConfiguration();
            return TeaModel.build(map, self);
        }

        public DataRedundancyTypeConfiguration setType(String type) {
            this.type = type;
            return this;
        }
        public String getType() {
            return this.type;
        }

    }

}
