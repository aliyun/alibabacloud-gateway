// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutAccessPointConfigForObjectProcessRequest extends TeaModel {
    @NameInMap("PutAccessPointConfigForObjectProcessConfiguration")
    public PutAccessPointConfigForObjectProcessConfiguration putAccessPointConfigForObjectProcessConfiguration;

    public static PutAccessPointConfigForObjectProcessRequest build(java.util.Map<String, ?> map) throws Exception {
        PutAccessPointConfigForObjectProcessRequest self = new PutAccessPointConfigForObjectProcessRequest();
        return TeaModel.build(map, self);
    }

    public PutAccessPointConfigForObjectProcessRequest setPutAccessPointConfigForObjectProcessConfiguration(PutAccessPointConfigForObjectProcessConfiguration putAccessPointConfigForObjectProcessConfiguration) {
        this.putAccessPointConfigForObjectProcessConfiguration = putAccessPointConfigForObjectProcessConfiguration;
        return this;
    }
    public PutAccessPointConfigForObjectProcessConfiguration getPutAccessPointConfigForObjectProcessConfiguration() {
        return this.putAccessPointConfigForObjectProcessConfiguration;
    }

    public static class PutAccessPointConfigForObjectProcessConfiguration extends TeaModel {
        @NameInMap("ObjectProcessConfiguration")
        public ObjectProcessConfiguration objectProcessConfiguration;

        public static PutAccessPointConfigForObjectProcessConfiguration build(java.util.Map<String, ?> map) throws Exception {
            PutAccessPointConfigForObjectProcessConfiguration self = new PutAccessPointConfigForObjectProcessConfiguration();
            return TeaModel.build(map, self);
        }

        public PutAccessPointConfigForObjectProcessConfiguration setObjectProcessConfiguration(ObjectProcessConfiguration objectProcessConfiguration) {
            this.objectProcessConfiguration = objectProcessConfiguration;
            return this;
        }
        public ObjectProcessConfiguration getObjectProcessConfiguration() {
            return this.objectProcessConfiguration;
        }

    }

}
