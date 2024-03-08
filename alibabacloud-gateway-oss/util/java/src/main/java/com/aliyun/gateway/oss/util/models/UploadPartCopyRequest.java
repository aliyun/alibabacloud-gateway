// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class UploadPartCopyRequest extends TeaModel {
    /**
     * <p>The number of parts.</p>
     */
    @NameInMap("partNumber")
    public Long partNumber;

    /**
     * <p>The ID that identifies the object to which the parts to upload belong.</p>
     */
    @NameInMap("uploadId")
    public String uploadId;

    public static UploadPartCopyRequest build(java.util.Map<String, ?> map) throws Exception {
        UploadPartCopyRequest self = new UploadPartCopyRequest();
        return TeaModel.build(map, self);
    }

    public UploadPartCopyRequest setPartNumber(Long partNumber) {
        this.partNumber = partNumber;
        return this;
    }
    public Long getPartNumber() {
        return this.partNumber;
    }

    public UploadPartCopyRequest setUploadId(String uploadId) {
        this.uploadId = uploadId;
        return this;
    }
    public String getUploadId() {
        return this.uploadId;
    }

}
