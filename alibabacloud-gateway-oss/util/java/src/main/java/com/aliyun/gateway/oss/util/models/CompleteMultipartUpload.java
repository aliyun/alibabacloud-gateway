// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class CompleteMultipartUpload extends TeaModel {
    @NameInMap("Part")
    public java.util.List<CompleteMultipartUploadPart> part;

    public static CompleteMultipartUpload build(java.util.Map<String, ?> map) throws Exception {
        CompleteMultipartUpload self = new CompleteMultipartUpload();
        return TeaModel.build(map, self);
    }

    public CompleteMultipartUpload setPart(java.util.List<CompleteMultipartUploadPart> part) {
        this.part = part;
        return this;
    }
    public java.util.List<CompleteMultipartUploadPart> getPart() {
        return this.part;
    }

    public static class CompleteMultipartUploadPart extends TeaModel {
        @NameInMap("ETag")
        public String ETag;

        @NameInMap("PartNumber")
        public Long partNumber;

        public static CompleteMultipartUploadPart build(java.util.Map<String, ?> map) throws Exception {
            CompleteMultipartUploadPart self = new CompleteMultipartUploadPart();
            return TeaModel.build(map, self);
        }

        public CompleteMultipartUploadPart setETag(String ETag) {
            this.ETag = ETag;
            return this;
        }
        public String getETag() {
            return this.ETag;
        }

        public CompleteMultipartUploadPart setPartNumber(Long partNumber) {
            this.partNumber = partNumber;
            return this;
        }
        public Long getPartNumber() {
            return this.partNumber;
        }

    }

}
