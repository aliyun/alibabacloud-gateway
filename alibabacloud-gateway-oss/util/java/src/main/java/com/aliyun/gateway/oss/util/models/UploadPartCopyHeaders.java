// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class UploadPartCopyHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>The address to access the source object. You must have permissions to read the source object.
     * <br>Default value: null</p>
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>/oss-example/ src-object</p>
     */
    @NameInMap("x-oss-copy-source")
    public String copySource;

    /**
     * <p>The copy operation condition. If the ETag value of the source object is the same as the ETag value provided by the user, OSS copies data. Otherwise, OSS returns 412 Precondition Failed.
     * <br>Default value: null</p>
     * 
     * <strong>example:</strong>
     * <p>5B3C1A2E053D763E1B002CC607C5****</p>
     */
    @NameInMap("x-oss-copy-source-if-match")
    public String copySourceIfMatch;

    /**
     * <p>The object transfer condition. If the specified time is earlier than the actual modified time of the object, the system transfers the object normally and returns 200 OK. Otherwise, OSS returns 304 Not Modified.
     * <br>Default value: null
     * <br>Time format: ddd, dd MMM yyyy HH:mm:ss GMT. Example: Fri, 13 Nov 2015 14:47:53 GMT.</p>
     * 
     * <strong>example:</strong>
     * <p>Fri, 13 Nov 2015 14:47:53 GMT</p>
     */
    @NameInMap("x-oss-copy-source-if-modified-since")
    public String copySourceIfModifiedSince;

    /**
     * <p>The object transfer condition. If the input ETag value does not match the ETag value of the object, the system transfers the object normally and returns 200 OK. Otherwise, OSS returns 304 Not Modified.
     * <br>Default value: null</p>
     * 
     * <strong>example:</strong>
     * <p>5B3C1A2E053D763E1B002CC607C5****</p>
     */
    @NameInMap("x-oss-copy-source-if-none-match")
    public String copySourceIfNoneMatch;

    /**
     * <p>The object transfer condition. If the specified time is the same as or later than the actual modified time of the object, OSS transfers the object normally and returns 200 OK. Otherwise, OSS returns 412 Precondition Failed.
     * <br>Default value: null</p>
     * 
     * <strong>example:</strong>
     * <p>Fri, 13 Oct 2015 14:47:53 GMT</p>
     */
    @NameInMap("x-oss-copy-source-if-unmodified-since")
    public String copySourceIfUnmodifiedSince;

    /**
     * <p>The range of bytes to copy data from the source object. For example, if you specify bytes to 0 to 9, the system transfers byte 0 to byte 9, a total of 10 bytes.
     * <br>Default value: null</p>
     * <ul>
     * <li>If the x-oss-copy-source-range request header is not specified, the entire source object is copied.</li>
     * <li>If the x-oss-copy-source-range request header is specified, the response contains the length of the entire object and the range of bytes to be copied for this operation. For example, Content-Range: bytes 0~9/44 indicates that the length of the entire object is 44 bytes. The range of bytes to be copied is byte 0 to byte 9.</li>
     * <li>If the specified range does not conform to the range conventions, OSS copies the entire object and does not include Content-Range in the response.</li>
     * </ul>
     * 
     * <strong>example:</strong>
     * <p>bytes=100-6291756</p>
     */
    @NameInMap("x-oss-copy-source-range")
    public String copySourceRange;

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

    public UploadPartCopyHeaders setCopySource(String copySource) {
        this.copySource = copySource;
        return this;
    }
    public String getCopySource() {
        return this.copySource;
    }

    public UploadPartCopyHeaders setCopySourceIfMatch(String copySourceIfMatch) {
        this.copySourceIfMatch = copySourceIfMatch;
        return this;
    }
    public String getCopySourceIfMatch() {
        return this.copySourceIfMatch;
    }

    public UploadPartCopyHeaders setCopySourceIfModifiedSince(String copySourceIfModifiedSince) {
        this.copySourceIfModifiedSince = copySourceIfModifiedSince;
        return this;
    }
    public String getCopySourceIfModifiedSince() {
        return this.copySourceIfModifiedSince;
    }

    public UploadPartCopyHeaders setCopySourceIfNoneMatch(String copySourceIfNoneMatch) {
        this.copySourceIfNoneMatch = copySourceIfNoneMatch;
        return this;
    }
    public String getCopySourceIfNoneMatch() {
        return this.copySourceIfNoneMatch;
    }

    public UploadPartCopyHeaders setCopySourceIfUnmodifiedSince(String copySourceIfUnmodifiedSince) {
        this.copySourceIfUnmodifiedSince = copySourceIfUnmodifiedSince;
        return this;
    }
    public String getCopySourceIfUnmodifiedSince() {
        return this.copySourceIfUnmodifiedSince;
    }

    public UploadPartCopyHeaders setCopySourceRange(String copySourceRange) {
        this.copySourceRange = copySourceRange;
        return this;
    }
    public String getCopySourceRange() {
        return this.copySourceRange;
    }

}
