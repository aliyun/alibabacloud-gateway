// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CreateAccessPointForObjectProcessRequest extends TeaModel {
    /**
     * <p>The container that stores information about the Object FC Access Point.</p>
     */
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
        /**
         * <p>The name of the access point.</p>
         * 
         * <strong>example:</strong>
         * <p>ap-01</p>
         */
        @NameInMap("AccessPointName")
        public String accessPointName;

        /**
         * <p>Whether allow anonymous user to access this FC Access Point.</p>
         * 
         * <strong>example:</strong>
         * <p>disable</p>
         */
        @NameInMap("AllowAnonymousAccessForObjectProcess")
        public String allowAnonymousAccessForObjectProcess;

        /**
         * <p>The container that stores the processing information about the Object FC Access Point.</p>
         */
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

        public CreateAccessPointForObjectProcessConfiguration setAllowAnonymousAccessForObjectProcess(String allowAnonymousAccessForObjectProcess) {
            this.allowAnonymousAccessForObjectProcess = allowAnonymousAccessForObjectProcess;
            return this;
        }
        public String getAllowAnonymousAccessForObjectProcess() {
            return this.allowAnonymousAccessForObjectProcess;
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
