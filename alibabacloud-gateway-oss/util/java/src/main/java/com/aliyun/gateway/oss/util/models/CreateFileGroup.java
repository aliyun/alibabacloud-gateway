// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CreateFileGroup extends TeaModel {
    @NameInMap("Part")
    public java.util.List<Part> part;

    public static CreateFileGroup build(java.util.Map<String, ?> map) throws Exception {
        CreateFileGroup self = new CreateFileGroup();
        return TeaModel.build(map, self);
    }

    public CreateFileGroup setPart(java.util.List<Part> part) {
        this.part = part;
        return this;
    }
    public java.util.List<Part> getPart() {
        return this.part;
    }

    public static class Part extends TeaModel {
        /**
         * <strong>example:</strong>
         * <p>&quot;EB327B57BB20D17C293A966115FE27BD&quot;</p>
         */
        @NameInMap("ETag")
        public String ETag;

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

        public Part setETag(String ETag) {
            this.ETag = ETag;
            return this;
        }
        public String getETag() {
            return this.ETag;
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
