// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class CompleteMultipartUploadRequest extends TeaModel {
    @NameInMap("body")
    public CompleteMultipartUpload body;

    /**
     * <p>The encoding type of the object name in the response. Only URL encoding is supported.</p>
     * <p><br>The object name can contain characters that are encoded in UTF-8. However, the XML 1.0 standard cannot be used to parse control characters, such as characters with an ASCII value from 0 to 10. You can configure this parameter to encode the object name in the response.</p>
     * <p><br>Default value: null</p>
     */
    @NameInMap("encoding-type")
    public String encodingType;

    /**
     * <p>The upload ID of the multipart upload task.</p>
     */
    @NameInMap("uploadId")
    public String uploadId;

    public static CompleteMultipartUploadRequest build(java.util.Map<String, ?> map) throws Exception {
        CompleteMultipartUploadRequest self = new CompleteMultipartUploadRequest();
        return TeaModel.build(map, self);
    }

    public CompleteMultipartUploadRequest setBody(CompleteMultipartUpload body) {
        this.body = body;
        return this;
    }
    public CompleteMultipartUpload getBody() {
        return this.body;
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
