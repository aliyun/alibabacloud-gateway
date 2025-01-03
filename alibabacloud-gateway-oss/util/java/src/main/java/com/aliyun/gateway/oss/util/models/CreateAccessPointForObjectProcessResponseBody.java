// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CreateAccessPointForObjectProcessResponseBody extends TeaModel {
    /**
     * <p>The container that stores information about the Object FC Access Point.</p>
     */
    @NameInMap("CreateAccessPointForObjectProcessResult")
    public CreateAccessPointForObjectProcessResult createAccessPointForObjectProcessResult;

    public static CreateAccessPointForObjectProcessResponseBody build(java.util.Map<String, ?> map) throws Exception {
        CreateAccessPointForObjectProcessResponseBody self = new CreateAccessPointForObjectProcessResponseBody();
        return TeaModel.build(map, self);
    }

    public CreateAccessPointForObjectProcessResponseBody setCreateAccessPointForObjectProcessResult(CreateAccessPointForObjectProcessResult createAccessPointForObjectProcessResult) {
        this.createAccessPointForObjectProcessResult = createAccessPointForObjectProcessResult;
        return this;
    }
    public CreateAccessPointForObjectProcessResult getCreateAccessPointForObjectProcessResult() {
        return this.createAccessPointForObjectProcessResult;
    }

    public static class CreateAccessPointForObjectProcessResult extends TeaModel {
        /**
         * <p>The alias of the Object FC Access Point.</p>
         * 
         * <strong>example:</strong>
         * <p>fc-ap-01-3b00521f653d2b3223680ec39dbbe2****-opapalias</p>
         */
        @NameInMap("AccessPointForObjectProcessAlias")
        public String accessPointForObjectProcessAlias;

        /**
         * <p>The ARN of the Object FC Access Point.</p>
         * 
         * <strong>example:</strong>
         * <p>acs:oss:cn-qingdao:119335441657143:accesspointforobjectprocess/fc-ap-01</p>
         */
        @NameInMap("AccessPointForObjectProcessArn")
        public String accessPointForObjectProcessArn;

        public static CreateAccessPointForObjectProcessResult build(java.util.Map<String, ?> map) throws Exception {
            CreateAccessPointForObjectProcessResult self = new CreateAccessPointForObjectProcessResult();
            return TeaModel.build(map, self);
        }

        public CreateAccessPointForObjectProcessResult setAccessPointForObjectProcessAlias(String accessPointForObjectProcessAlias) {
            this.accessPointForObjectProcessAlias = accessPointForObjectProcessAlias;
            return this;
        }
        public String getAccessPointForObjectProcessAlias() {
            return this.accessPointForObjectProcessAlias;
        }

        public CreateAccessPointForObjectProcessResult setAccessPointForObjectProcessArn(String accessPointForObjectProcessArn) {
            this.accessPointForObjectProcessArn = accessPointForObjectProcessArn;
            return this;
        }
        public String getAccessPointForObjectProcessArn() {
            return this.accessPointForObjectProcessArn;
        }

    }

}
