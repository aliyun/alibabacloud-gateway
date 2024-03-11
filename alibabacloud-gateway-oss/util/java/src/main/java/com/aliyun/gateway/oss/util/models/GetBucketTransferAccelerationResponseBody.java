// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class GetBucketTransferAccelerationResponseBody extends TeaModel {
    /**
     * <p>The container that stores the transfer acceleration configurations.</p>
     */
    @NameInMap("TransferAccelerationConfiguration")
    public TransferAccelerationConfiguration transferAccelerationConfiguration;

    public static GetBucketTransferAccelerationResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketTransferAccelerationResponseBody self = new GetBucketTransferAccelerationResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketTransferAccelerationResponseBody setTransferAccelerationConfiguration(TransferAccelerationConfiguration transferAccelerationConfiguration) {
        this.transferAccelerationConfiguration = transferAccelerationConfiguration;
        return this;
    }
    public TransferAccelerationConfiguration getTransferAccelerationConfiguration() {
        return this.transferAccelerationConfiguration;
    }

    public static class TransferAccelerationConfiguration extends TeaModel {
        /**
         * <p>Whether the transfer acceleration is enabled for this bucket.</p>
         */
        @NameInMap("Enabled")
        public Boolean enabled;

        public static TransferAccelerationConfiguration build(java.util.Map<String, ?> map) throws Exception {
            TransferAccelerationConfiguration self = new TransferAccelerationConfiguration();
            return TeaModel.build(map, self);
        }

        public TransferAccelerationConfiguration setEnabled(Boolean enabled) {
            this.enabled = enabled;
            return this;
        }
        public Boolean getEnabled() {
            return this.enabled;
        }

    }

}
