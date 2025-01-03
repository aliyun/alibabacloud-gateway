// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CopyObjectHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>The path of the source object. By default, this header is left empty.</p>
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>/oss-example/oss.jpg</p>
     */
    @NameInMap("x-oss-copy-source")
    public String copySource;

    /**
     * <p>The object copy condition. If the ETag value of the source object is the same as the ETag value that you specify in the request, OSS copies the object and returns 200 OK. By default, this header is left empty.</p>
     * 
     * <strong>example:</strong>
     * <p>5B3C1A2E053D763E1B002CC607C5****</p>
     */
    @NameInMap("x-oss-copy-source-if-match")
    public String copySourceIfMatch;

    /**
     * <p>If the source object is modified after the time that you specify in the request, OSS copies the object. By default, this header is left empty.</p>
     * 
     * <strong>example:</strong>
     * <p>2019-04-09T07:01:56.000Z</p>
     */
    @NameInMap("x-oss-copy-source-if-modified-since")
    public String copySourceIfModifiedSince;

    /**
     * <p>The object copy condition. If the ETag value of the source object is different from the ETag value that you specify in the request, OSS copies the object and returns 200 OK. By default, this header is left empty.</p>
     * 
     * <strong>example:</strong>
     * <p>5B3C1A2E053D763E1B002CC607C5****</p>
     */
    @NameInMap("x-oss-copy-source-if-none-match")
    public String copySourceIfNoneMatch;

    /**
     * <p>The object copy condition. If the time that you specify in the request is the same as or later than the modification time of the object, OSS copies the object and returns 200 OK. By default, this header is left empty.</p>
     * 
     * <strong>example:</strong>
     * <p>2019-04-09T07:01:56.000Z</p>
     */
    @NameInMap("x-oss-copy-source-if-unmodified-since")
    public String copySourceIfUnmodifiedSince;

    /**
     * <p>Specifies whether the CopyObject operation overwrites objects with the same name. The <strong>x-oss-forbid-overwrite</strong> request header does not take effect when versioning is enabled or suspended for the destination bucket. In this case, the CopyObject operation overwrites the existing object that has the same name as the destination object.</p>
     * <ul>
     * <li>If you do not specify the <strong>x-oss-forbid-overwrite</strong> header or set the header to <strong>false</strong>, an existing object that has the same name as the object that you want to copy is overwritten.****</li>
     * <li>If you set the <strong>x-oss-forbid-overwrite</strong> header to <strong>true</strong>, an existing object that has the same name as the object that you want to copy is not overwritten.</li>
     * </ul>
     * <p>If you specify the <strong>x-oss-forbid-overwrite</strong> header, the queries per second (QPS) performance of OSS may be degraded. If you want to specify the <strong>x-oss-forbid-overwrite</strong> header in a large number of requests (QPS greater than 1,000), contact technical support. Default value: false.</p>
     * 
     * <strong>example:</strong>
     * <p>true</p>
     */
    @NameInMap("x-oss-forbid-overwrite")
    public String forbidOverwrite;

    /**
     * <p>You can add parameters that contain the x-oss-meta- prefix when you create an append object. You cannot include these parameters in the requests when you append objects to an existing append object. Parameters that contain the x-oss-meta-\* prefix are considered the metadata of the object. You can specify multiple parameters that contain the x-oss-meta- prefix for an object. The total size of the metadata cannot exceed 8 KB. The names of parameters that contain the x-oss-meta- prefix can contain hyphens (-), digits, and letters. Uppercase letters are converted into lowercase letters. Other characters such as underscores (_) are not supported.</p>
     */
    @NameInMap("x-oss-meta-*")
    public java.util.Map<String, String> metaData;

    /**
     * <p>The method that is used to configure the metadata of the destination object. Default value: COPY.</p>
     * <ul>
     * <li><strong>COPY</strong>: The metadata of the source object is copied to the destination object. The <strong>x-oss-server-side-encryption</strong> attribute of the source object is not copied to the destination object. The <strong>x-oss-server-side-encryption</strong> header in the CopyObject request specifies the method that is used to encrypt the destination object.</li>
     * <li><strong>REPLACE</strong>: The metadata that you specify in the request is used as the metadata of the destination object.</li>
     * </ul>
     * <blockquote>
     * <p> If the path of the source object is the same as the path of the destination object and versioning is disabled for the bucket in which the source and destination objects are stored, the metadata that you specify in the CopyObject request is used as the metadata of the destination object regardless of the value of the x-oss-metadata-directive header.</p>
     * </blockquote>
     * 
     * <strong>example:</strong>
     * <p>COPY</p>
     */
    @NameInMap("x-oss-metadata-directive")
    public String metadataDirective;

    /**
     * <p>The access control list (ACL) of the destination object when the object is created. Default value: default.</p>
     * <p>Valid values:</p>
     * <ul>
     * <li>default: The ACL of the object is the same as the ACL of the bucket in which the object is stored.</li>
     * <li>private: The ACL of the object is private. Only the owner of the object and authorized users have read and write permissions on the object. Other users do not have permissions on the object.</li>
     * <li>public-read: The ACL of the object is public-read. Only the owner of the object and authorized users have read and write permissions on the object. Other users have only read permissions on the object. Exercise caution when you set the ACL of the bucket to this value.</li>
     * <li>public-read-write: The ACL of the object is public-read-write. All users have read and write permissions on the object. Exercise caution when you set the ACL of the bucket to this value.</li>
     * </ul>
     * <p>For more information about ACLs, see <a href="https://help.aliyun.com/document_detail/100676.html">Object ACL</a>.</p>
     */
    @NameInMap("x-oss-object-acl")
    public String acl;

    /**
     * <p>The server side data encryption algorithm. Invalid value: SM4</p>
     * 
     * <strong>example:</strong>
     * <p>SM4</p>
     */
    @NameInMap("x-oss-server-side-data-encryption")
    public String xOssServerSideDataEncryption;

    /**
     * <p>The entropy coding-based encryption algorithm that OSS uses to encrypt an object when you create the object. The valid values of the header are <strong>AES256</strong> and <strong>KMS</strong>. You must activate Key Management Service (KMS) in the OSS console before you can use the KMS encryption algorithm. Otherwise, the KmsServiceNotEnabled error is returned.</p>
     * <ul>
     * <li>If you do not specify the <strong>x-oss-server-side-encryption</strong> header in the CopyObject request, the destination object is not encrypted on the server regardless of whether the source object is encrypted on the server.</li>
     * <li>If you specify the <strong>x-oss-server-side-encryption</strong> header in the CopyObject request, the destination object is encrypted on the server after the CopyObject operation is performed regardless of whether the source object is encrypted on the server. In addition, the response to a CopyObject request contains the <strong>x-oss-server-side-encryption</strong> header whose value is the encryption algorithm of the destination object. When the destination object is downloaded, the <strong>x-oss-server-side-encryption</strong> header is included in the response. The value of this header is the encryption algorithm of the destination object.</li>
     * </ul>
     * 
     * <strong>example:</strong>
     * <p>AES256</p>
     */
    @NameInMap("x-oss-server-side-encryption")
    public String serverSideEncryption;

    /**
     * <p>The ID of the customer master key (CMK) that is managed by KMS. This parameter is available only if you set <strong>x-oss-server-side-encryption</strong> to KMS.</p>
     * 
     * <strong>example:</strong>
     * <p>9468da86-3509-4f8d-a61e-6eab1eac****</p>
     */
    @NameInMap("x-oss-server-side-encryption-key-id")
    public String sseKeyId;

    /**
     * <p>The storage class of the object that you want to upload. Default value: Standard. If you specify a storage class when you upload the object, the storage class applies regardless of the storage class of the bucket to which you upload the object. For example, if you set <strong>x-oss-storage-class</strong> to Standard when you upload an object to an IA bucket, the storage class of the uploaded object is Standard.</p>
     * <p>Valid values:</p>
     * <ul>
     * <li>Standard</li>
     * <li>IA</li>
     * <li>Archive</li>
     * <li>ColdArchive</li>
     * </ul>
     * <p>For more information about storage classes, see <a href="https://help.aliyun.com/document_detail/51374.html">Overview</a>.</p>
     */
    @NameInMap("x-oss-storage-class")
    public String storageClass;

    /**
     * <p>The tag of the destination object. You can add multiple tags to the destination object. Example: TagA=A\&amp;TagB=B.</p>
     * <blockquote>
     * <p> The tag key and tag value must be URL-encoded. If a key-value pair does not contain an equal sign (=), the tag value is considered an empty string.</p>
     * </blockquote>
     * 
     * <strong>example:</strong>
     * <p>a:1</p>
     */
    @NameInMap("x-oss-tagging")
    public String tagging;

    /**
     * <p>The method that is used to add tags to the destination object. Default value: Copy. Valid values:</p>
     * <ul>
     * <li><strong>Copy</strong>: The tags of the source object are copied to the destination object.</li>
     * <li><strong>Replace</strong>: The tags that you specify in the request are added to the destination object.</li>
     * </ul>
     * 
     * <strong>example:</strong>
     * <p>Copy</p>
     */
    @NameInMap("x-oss-tagging-directive")
    public String taggingDirective;

    public static CopyObjectHeaders build(java.util.Map<String, ?> map) throws Exception {
        CopyObjectHeaders self = new CopyObjectHeaders();
        return TeaModel.build(map, self);
    }

    public CopyObjectHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public CopyObjectHeaders setCopySource(String copySource) {
        this.copySource = copySource;
        return this;
    }
    public String getCopySource() {
        return this.copySource;
    }

    public CopyObjectHeaders setCopySourceIfMatch(String copySourceIfMatch) {
        this.copySourceIfMatch = copySourceIfMatch;
        return this;
    }
    public String getCopySourceIfMatch() {
        return this.copySourceIfMatch;
    }

    public CopyObjectHeaders setCopySourceIfModifiedSince(String copySourceIfModifiedSince) {
        this.copySourceIfModifiedSince = copySourceIfModifiedSince;
        return this;
    }
    public String getCopySourceIfModifiedSince() {
        return this.copySourceIfModifiedSince;
    }

    public CopyObjectHeaders setCopySourceIfNoneMatch(String copySourceIfNoneMatch) {
        this.copySourceIfNoneMatch = copySourceIfNoneMatch;
        return this;
    }
    public String getCopySourceIfNoneMatch() {
        return this.copySourceIfNoneMatch;
    }

    public CopyObjectHeaders setCopySourceIfUnmodifiedSince(String copySourceIfUnmodifiedSince) {
        this.copySourceIfUnmodifiedSince = copySourceIfUnmodifiedSince;
        return this;
    }
    public String getCopySourceIfUnmodifiedSince() {
        return this.copySourceIfUnmodifiedSince;
    }

    public CopyObjectHeaders setForbidOverwrite(String forbidOverwrite) {
        this.forbidOverwrite = forbidOverwrite;
        return this;
    }
    public String getForbidOverwrite() {
        return this.forbidOverwrite;
    }

    public CopyObjectHeaders setMetaData(java.util.Map<String, String> metaData) {
        this.metaData = metaData;
        return this;
    }
    public java.util.Map<String, String> getMetaData() {
        return this.metaData;
    }

    public CopyObjectHeaders setMetadataDirective(String metadataDirective) {
        this.metadataDirective = metadataDirective;
        return this;
    }
    public String getMetadataDirective() {
        return this.metadataDirective;
    }

    public CopyObjectHeaders setAcl(String acl) {
        this.acl = acl;
        return this;
    }
    public String getAcl() {
        return this.acl;
    }

    public CopyObjectHeaders setXOssServerSideDataEncryption(String xOssServerSideDataEncryption) {
        this.xOssServerSideDataEncryption = xOssServerSideDataEncryption;
        return this;
    }
    public String getXOssServerSideDataEncryption() {
        return this.xOssServerSideDataEncryption;
    }

    public CopyObjectHeaders setServerSideEncryption(String serverSideEncryption) {
        this.serverSideEncryption = serverSideEncryption;
        return this;
    }
    public String getServerSideEncryption() {
        return this.serverSideEncryption;
    }

    public CopyObjectHeaders setSseKeyId(String sseKeyId) {
        this.sseKeyId = sseKeyId;
        return this;
    }
    public String getSseKeyId() {
        return this.sseKeyId;
    }

    public CopyObjectHeaders setStorageClass(String storageClass) {
        this.storageClass = storageClass;
        return this;
    }
    public String getStorageClass() {
        return this.storageClass;
    }

    public CopyObjectHeaders setTagging(String tagging) {
        this.tagging = tagging;
        return this;
    }
    public String getTagging() {
        return this.tagging;
    }

    public CopyObjectHeaders setTaggingDirective(String taggingDirective) {
        this.taggingDirective = taggingDirective;
        return this;
    }
    public String getTaggingDirective() {
        return this.taggingDirective;
    }

}
