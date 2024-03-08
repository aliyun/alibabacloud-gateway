// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class CreateAccessPointForObjectProcessRequest extends TeaModel {
    @NameInMap("CreateAccessPointForObjectProcessConfiguration")
    public CreateAccessPointForObjectProcessConfiguration createAccessPointForObjectProcessConfiguration;

    public static CreateAccessPointForObjectProcessRequest build(java.util.Map<String, ?> map) throws Exception {
        CreateAccessPointForObjectProcessRequest self = new CreateAccessPointForObjectProcessRequest();
        return TeaModel.build(map, self);
    }

    public CreateAccessPointForObjectProcessRequest setCreateAccessPointForObjectProcessConfiguration(CreateAccessPointForObjectProcessConfiguration createAccessPointForObjectProcessConfiguration) {
        this.createAccessPointForObjectProcessConfiguration = createAccessPointForObjectProcessConfiguration;
        return this;
    }
    public CreateAccessPointForObjectProcessConfiguration getCreateAccessPointForObjectProcessConfiguration() {
        return this.createAccessPointForObjectProcessConfiguration;
    }

    public static class CreateAccessPointForObjectProcessConfiguration extends TeaModel {
        @NameInMap("AccessPointName")
        public String accessPointName;

        @NameInMap("ObjectProcessConfiguration")
        public ObjectProcessConfiguration objectProcessConfiguration;

        public static CreateAccessPointForObjectProcessConfiguration build(java.util.Map<String, ?> map) throws Exception {
            CreateAccessPointForObjectProcessConfiguration self = new CreateAccessPointForObjectProcessConfiguration();
            return TeaModel.build(map, self);
        }

        public CreateAccessPointForObjectProcessConfiguration setAccessPointName(String accessPointName) {
            this.accessPointName = accessPointName;
            return this;
        }
        public String getAccessPointName() {
            return this.accessPointName;
        }

        public CreateAccessPointForObjectProcessConfiguration setObjectProcessConfiguration(ObjectProcessConfiguration objectProcessConfiguration) {
            this.objectProcessConfiguration = objectProcessConfiguration;
            return this;
        }
        public ObjectProcessConfiguration getObjectProcessConfiguration() {
            return this.objectProcessConfiguration;
        }

    }

}
