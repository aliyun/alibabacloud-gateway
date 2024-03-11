// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class CompleteMultipartUploadRequest extends TeaModel {
    /**
     * <p>The container that stores the content of the CompleteMultipartUpload request.</p>
     */
    @NameInMap("CompleteMultipartUpload")
    public CompleteMultipartUpload completeMultipartUpload;

    /**
     * <p>The encodingtype of the object name in the response. Only URL encoding is supported.</p>
     * <p>The object name can contain characters that are encoded in UTF-8. However, the XML 1.0 standard cannot be used to parse control characters, such as characters with an ASCII value from 0 to 10. You can configure this parameter to encode the object name in the response.</p>
     */
    @NameInMap("encoding-type")
    public String encodingType;

    /**
     * <p>The identifier of the multipart upload task.</p>
     */
    @NameInMap("uploadId")
    public String uploadId;

    public static CompleteMultipartUploadRequest build(java.util.Map<String, ?> map) throws Exception {
        CompleteMultipartUploadRequest self = new CompleteMultipartUploadRequest();
        return TeaModel.build(map, self);
    }

    public CompleteMultipartUploadRequest setCompleteMultipartUpload(CompleteMultipartUpload completeMultipartUpload) {
        this.completeMultipartUpload = completeMultipartUpload;
        return this;
    }
    public CompleteMultipartUpload getCompleteMultipartUpload() {
        return this.completeMultipartUpload;
    }

    public CompleteMultipartUploadRequest setEncodingType(String encodingType) {
        this.encodingType = encodingType;
        return this;
    }
    public String getEncodingType() {
        return this.encodingType;
    }

    public CompleteMultipartUploadRequest setUploadId(String uploadId) {
        this.uploadId = uploadId;
        return this;
    }
    public String getUploadId() {
        return this.uploadId;
    }

}
