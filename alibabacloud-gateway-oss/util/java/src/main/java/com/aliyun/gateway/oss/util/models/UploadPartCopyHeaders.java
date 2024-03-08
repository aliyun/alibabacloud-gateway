// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class UploadPartCopyHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>The address to access the source object. You must have permissions to read the source object.</p>
     * <p><br>Default value: null</p>
     */
    @NameInMap("x-oss-copy-source")
    public String xOssCopySource;

    /**
     * <p>The copy operation condition. If the ETag value of the source object is the same as the ETag value provided by the user, OSS copies data. Otherwise, OSS returns 412 Precondition Failed.</p>
     * <p><br>Default value: null</p>
     */
    @NameInMap("x-oss-copy-source-if-match")
    public String xOssCopySourceIfMatch;

    /**
     * <p>The object transfer condition. If the specified time is earlier than the actual modified time of the object, the system transfers the object normally and returns 200 OK. Otherwise, OSS returns 304 Not Modified.</p>
     * <p><br>Default value: null</p>
     * <p><br>Time format: ddd, dd MMM yyyy HH:mm:ss GMT. Example: Fri, 13 Nov 2015 14:47:53 GMT.</p>
     */
    @NameInMap("x-oss-copy-source-if-modified-since")
    public String xOssCopySourceIfModifiedSince;

    /**
     * <p>The object transfer condition. If the input ETag value does not match the ETag value of the object, the system transfers the object normally and returns 200 OK. Otherwise, OSS returns 304 Not Modified.</p>
     * <p><br>Default value: null</p>
     */
    @NameInMap("x-oss-copy-source-if-none-match")
    public String xOssCopySourceIfNoneMatch;

    /**
     * <p>The object transfer condition. If the specified time is the same as or later than the actual modified time of the object, OSS transfers the object normally and returns 200 OK. Otherwise, OSS returns 412 Precondition Failed.</p>
     * <p><br>Default value: null</p>
     */
    @NameInMap("x-oss-copy-source-if-unmodified-since")
    public String xOssCopySourceIfUnmodifiedSince;

    /**
     * <p>The range of bytes to copy data from the source object. For example, if you specify bytes to 0 to 9, the system transfers byte 0 to byte 9, a total of 10 bytes.</p>
     * <p><br>Default value: null</p>
     * <br>
     * <p>- If the x-oss-copy-source-range request header is not specified, the entire source object is copied.</p>
     * <p>- If the x-oss-copy-source-range request header is specified, the response contains the length of the entire object and the range of bytes to be copied for this operation. For example, Content-Range: bytes 0~9/44 indicates that the length of the entire object is 44 bytes. The range of bytes to be copied is byte 0 to byte 9.</p>
     * <p>- If the specified range does not conform to the range conventions, OSS copies the entire object and does not include Content-Range in the response.</p>
     */
    @NameInMap("x-oss-copy-source-range")
    public String xOssCopySourceRange;

    public static UploadPartCopyHeaders build(java.util.Map<String, ?> map) throws Exception {
        UploadPartCopyHeaders self = new UploadPartCopyHeaders();
        return TeaModel.build(map, self);
    }

    public UploadPartCopyHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public UploadPartCopyHeaders setXOssCopySource(String xOssCopySource) {
        this.xOssCopySource = xOssCopySource;
        return this;
    }
    public String getXOssCopySource() {
        return this.xOssCopySource;
    }

    public UploadPartCopyHeaders setXOssCopySourceIfMatch(String xOssCopySourceIfMatch) {
        this.xOssCopySourceIfMatch = xOssCopySourceIfMatch;
        return this;
    }
    public String getXOssCopySourceIfMatch() {
        return this.xOssCopySourceIfMatch;
    }

    public UploadPartCopyHeaders setXOssCopySourceIfModifiedSince(String xOssCopySourceIfModifiedSince) {
        this.xOssCopySourceIfModifiedSince = xOssCopySourceIfModifiedSince;
        return this;
    }
    public String getXOssCopySourceIfModifiedSince() {
        return this.xOssCopySourceIfModifiedSince;
    }

    public UploadPartCopyHeaders setXOssCopySourceIfNoneMatch(String xOssCopySourceIfNoneMatch) {
        this.xOssCopySourceIfNoneMatch = xOssCopySourceIfNoneMatch;
        return this;
    }
    public String getXOssCopySourceIfNoneMatch() {
        return this.xOssCopySourceIfNoneMatch;
    }

    public UploadPartCopyHeaders setXOssCopySourceIfUnmodifiedSince(String xOssCopySourceIfUnmodifiedSince) {
        this.xOssCopySourceIfUnmodifiedSince = xOssCopySourceIfUnmodifiedSince;
        return this;
    }
    public String getXOssCopySourceIfUnmodifiedSince() {
        return this.xOssCopySourceIfUnmodifiedSince;
    }

    public UploadPartCopyHeaders setXOssCopySourceRange(String xOssCopySourceRange) {
        this.xOssCopySourceRange = xOssCopySourceRange;
        return this;
    }
    public String getXOssCopySourceRange() {
        return this.xOssCopySourceRange;
    }

}
