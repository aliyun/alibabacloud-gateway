// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class UploadPartCopyResponseBody extends TeaModel {
    /**
     * <p>The ETag that is generated when an object is created. ETags are used to identify the content of objects.</p>
     * <p><br>If an object is created by calling the CompleteMultipartUpload operation, the ETag value is not the MD5 hash of the object content but a unique value calculated based on a specific rule.</p>
     * <p>>The ETag value of the object can be used to check whether the object content is modified. However, we recommend that you use the MD5 hash of an object rather than the ETag value of the object to verify data integrity.</p>
     */
    @NameInMap("ETag")
    public String ETag;

    /**
     * <p>The time when the object was last modified.</p>
     */
    @NameInMap("LastModified")
    public String lastModified;

    public static UploadPartCopyResponseBody build(java.util.Map<String, ?> map) throws Exception {
        UploadPartCopyResponseBody self = new UploadPartCopyResponseBody();
        return TeaModel.build(map, self);
    }

    public UploadPartCopyResponseBody setETag(String ETag) {
        this.ETag = ETag;
        return this;
    }
    public String getETag() {
        return this.ETag;
    }

    public UploadPartCopyResponseBody setLastModified(String lastModified) {
        this.lastModified = lastModified;
        return this;
    }
    public String getLastModified() {
        return this.lastModified;
    }

}
