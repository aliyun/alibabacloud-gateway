// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class CompleteMultipartUpload extends TeaModel {
    @NameInMap("Part")
    public java.util.List<Part> part;

    public static CompleteMultipartUpload build(java.util.Map<String, ?> map) throws Exception {
        CompleteMultipartUpload self = new CompleteMultipartUpload();
        return TeaModel.build(map, self);
    }

    public CompleteMultipartUpload setPart(java.util.List<Part> part) {
        this.part = part;
        return this;
    }
    public java.util.List<Part> getPart() {
        return this.part;
    }

    public static class Part extends TeaModel {
        @NameInMap("ETag")
        public String ETag;

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

        public Part setPartNumber(Long partNumber) {
            this.partNumber = partNumber;
            return this;
        }
        public Long getPartNumber() {
            return this.partNumber;
        }

    }

}
