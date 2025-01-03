// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ObjectLinkInfo extends TeaModel {
    @NameInMap("Part")
    public java.util.List<Part> part;

    public static ObjectLinkInfo build(java.util.Map<String, ?> map) throws Exception {
        ObjectLinkInfo self = new ObjectLinkInfo();
        return TeaModel.build(map, self);
    }

    public ObjectLinkInfo setPart(java.util.List<Part> part) {
        this.part = part;
        return this;
    }
    public java.util.List<Part> getPart() {
        return this.part;
    }

    public static class Part extends TeaModel {
        /**
         * <strong>example:</strong>
         * <p>test.txt</p>
         */
        @NameInMap("PartName")
        public String partName;

        /**
         * <strong>example:</strong>
         * <p>3</p>
         */
        @NameInMap("PartNumber")
        public Long partNumber;

        public static Part build(java.util.Map<String, ?> map) throws Exception {
            Part self = new Part();
            return TeaModel.build(map, self);
        }

        public Part setPartName(String partName) {
            this.partName = partName;
            return this;
        }
        public String getPartName() {
            return this.partName;
        }

        public Part setPartNumber(Long partNumber) {
            this.partNumber = partNumber;
            return this;
        }
        public Long getPartNumber() {
            return this.partNumber;
        }

    }

}
