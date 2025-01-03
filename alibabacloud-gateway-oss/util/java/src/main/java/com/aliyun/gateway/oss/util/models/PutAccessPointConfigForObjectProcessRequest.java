// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutAccessPointConfigForObjectProcessRequest extends TeaModel {
    /**
     * <p>The container that stores information about the Object FC Access Point.</p>
     */
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
        /**
         * <p>Whether allow anonymous user to access this FC Access Point.</p>
         * 
         * <strong>example:</strong>
         * <p>false</p>
         */
        @NameInMap("AllowAnonymousAccessForObjectProcess")
        public String allowAnonymousAccessForObjectProcess;

        /**
         * <p>The container that stores the processing information about the Object FC Access Point.</p>
         */
        @NameInMap("ObjectProcessConfiguration")
        public ObjectProcessConfiguration objectProcessConfiguration;

        /**
         * <p>The container in which the Block Public Access configurations are stored.</p>
         */
        @NameInMap("PublicAccessBlockConfiguration")
        public PublicAccessBlockConfiguration publicAccessBlockConfiguration;

        public static PutAccessPointConfigForObjectProcessConfiguration build(java.util.Map<String, ?> map) throws Exception {
            PutAccessPointConfigForObjectProcessConfiguration self = new PutAccessPointConfigForObjectProcessConfiguration();
            return TeaModel.build(map, self);
        }

        public PutAccessPointConfigForObjectProcessConfiguration setAllowAnonymousAccessForObjectProcess(String allowAnonymousAccessForObjectProcess) {
            this.allowAnonymousAccessForObjectProcess = allowAnonymousAccessForObjectProcess;
            return this;
        }
        public String getAllowAnonymousAccessForObjectProcess() {
            return this.allowAnonymousAccessForObjectProcess;
        }

        public PutAccessPointConfigForObjectProcessConfiguration setObjectProcessConfiguration(ObjectProcessConfiguration objectProcessConfiguration) {
            this.objectProcessConfiguration = objectProcessConfiguration;
            return this;
        }
        public ObjectProcessConfiguration getObjectProcessConfiguration() {
            return this.objectProcessConfiguration;
        }

        public PutAccessPointConfigForObjectProcessConfiguration setPublicAccessBlockConfiguration(PublicAccessBlockConfiguration publicAccessBlockConfiguration) {
            this.publicAccessBlockConfiguration = publicAccessBlockConfiguration;
            return this;
        }
        public PublicAccessBlockConfiguration getPublicAccessBlockConfiguration() {
            return this.publicAccessBlockConfiguration;
        }

    }

}
