// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class CopyObjectHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>The path of the source object. By default, this header is left empty.</p>
     */
    @NameInMap("x-oss-copy-source")
    public String xOssCopySource;

    /**
     * <p>The object copy condition. If the ETag value of the source object is the same as the ETag value that you specify in the request, OSS copies the object and returns 200 OK. By default, this header is left empty.</p>
     */
    @NameInMap("x-oss-copy-source-if-match")
    public String xOssCopySourceIfMatch;

    /**
     * <p>If the source object is modified after the time that you specify in the request, OSS copies the object. By default, this header is left empty.</p>
     */
    @NameInMap("x-oss-copy-source-if-modified-since")
    public String xOssCopySourceIfModifiedSince;

    /**
     * <p>The object copy condition. If the ETag value of the source object is different from the ETag value that you specify in the request, OSS copies the object and returns 200 OK. By default, this header is left empty.</p>
     */
    @NameInMap("x-oss-copy-source-if-none-match")
    public String xOssCopySourceIfNoneMatch;

    /**
     * <p>The object copy condition. If the time that you specify in the request is the same as or later than the modification time of the object, OSS copies the object and returns 200 OK. By default, this header is left empty.</p>
     */
    @NameInMap("x-oss-copy-source-if-unmodified-since")
    public String xOssCopySourceIfUnmodifiedSince;

    /**
     * <p>Specifies whether the CopyObject operation overwrites objects with the same name. The **x-oss-forbid-overwrite** request header does not take effect when versioning is enabled or suspended for the destination bucket. In this case, the CopyObject operation overwrites the existing object that has the same name as the destination object.</p>
     * <br>
     * <p>*   If you do not specify the **x-oss-forbid-overwrite** header or set the header to **false**, an existing object that has the same name as the object that you want to copy is overwritten.****</p>
     * <p>*   If you set the **x-oss-forbid-overwrite** header to **true**, an existing object that has the same name as the object that you want to copy is not overwritten.</p>
     * <br>
     * <p>If you specify the **x-oss-forbid-overwrite** header, the queries per second (QPS) performance of OSS may be degraded. If you want to specify the **x-oss-forbid-overwrite** header in a large number of requests (QPS greater than 1,000), contact technical support. Default value: false.</p>
     */
    @NameInMap("x-oss-forbid-overwrite")
    public String xOssForbidOverwrite;

    /**
     * <p>You can add parameters that contain the x-oss-meta- prefix when you create an append object. You cannot include these parameters in the requests when you append objects to an existing append object. Parameters that contain the x-oss-meta-\* prefix are considered the metadata of the object. You can specify multiple parameters that contain the x-oss-meta- prefix for an object. The total size of the metadata cannot exceed 8 KB. The names of parameters that contain the x-oss-meta- prefix can contain hyphens (-), digits, and letters. Uppercase letters are converted into lowercase letters. Other characters such as underscores (\_) are not supported.</p>
     */
    @NameInMap("x-oss-meta-*")
    public java.util.Map<String, String> xOssMeta;

    /**
     * <p>The method that is used to configure the metadata of the destination object. Default value: COPY.</p>
     * <br>
     * <p>*   **COPY**: The metadata of the source object is copied to the destination object. The **x-oss-server-side-encryption** attribute of the source object is not copied to the destination object. The **x-oss-server-side-encryption** header in the CopyObject request specifies the method that is used to encrypt the destination object.</p>
     * <p>*   **REPLACE**: The metadata that you specify in the request is used as the metadata of the destination object.</p>
     * <br>
     * <p>>  If the path of the source object is the same as the path of the destination object and versioning is disabled for the bucket in which the source and destination objects are stored, the metadata that you specify in the CopyObject request is used as the metadata of the destination object regardless of the value of the x-oss-metadata-directive header.</p>
     */
    @NameInMap("x-oss-metadata-directive")
    public String xOssMetadataDirective;

    /**
     * <p>The access control list (ACL) of the destination object when the object is created. Default value: default.</p>
     * <br>
     * <p>Valid values:</p>
     * <br>
     * <p>*   default: The ACL of the object is the same as the ACL of the bucket in which the object is stored.</p>
     * <p>*   private: The ACL of the object is private. Only the owner of the object and authorized users have read and write permissions on the object. Other users do not have permissions on the object.</p>
     * <p>*   public-read: The ACL of the object is public-read. Only the owner of the object and authorized users have read and write permissions on the object. Other users have only read permissions on the object. Exercise caution when you set the ACL of the bucket to this value.</p>
     * <p>*   public-read-write: The ACL of the object is public-read-write. All users have read and write permissions on the object. Exercise caution when you set the ACL of the bucket to this value.</p>
     * <br>
     * <p>For more information about ACLs, see [Object ACL](~~100676~~).</p>
     */
    @NameInMap("x-oss-object-acl")
    public String xOssObjectAcl;

    /**
     * <p>The entropy coding-based encryption algorithm that OSS uses to encrypt an object when you create the object. The valid values of the header are **AES256** and **KMS**. You must activate Key Management Service (KMS) in the OSS console before you can use the KMS encryption algorithm. Otherwise, the KmsServiceNotEnabled error is returned.</p>
     * <br>
     * <p>*   If you do not specify the **x-oss-server-side-encryption** header in the CopyObject request, the destination object is not encrypted on the server regardless of whether the source object is encrypted on the server.</p>
     * <p>*   If you specify the **x-oss-server-side-encryption** header in the CopyObject request, the destination object is encrypted on the server after the CopyObject operation is performed regardless of whether the source object is encrypted on the server. In addition, the response to a CopyObject request contains the **x-oss-server-side-encryption** header whose value is the encryption algorithm of the destination object. When the destination object is downloaded, the **x-oss-server-side-encryption** header is included in the response. The value of this header is the encryption algorithm of the destination object.</p>
     */
    @NameInMap("x-oss-server-side-encryption")
    public String xOssServerSideEncryption;

    /**
     * <p>The ID of the customer master key (CMK) that is managed by KMS. This parameter is available only if you set **x-oss-server-side-encryption** to KMS.</p>
     */
    @NameInMap("x-oss-server-side-encryption-key-id")
    public String xOssServerSideEncryptionKeyId;

    /**
     * <p>The storage class of the object that you want to upload. Default value: Standard. If you specify a storage class when you upload the object, the storage class applies regardless of the storage class of the bucket to which you upload the object. For example, if you set **x-oss-storage-class** to Standard when you upload an object to an IA bucket, the storage class of the uploaded object is Standard.</p>
     * <br>
     * <p>Valid values:</p>
     * <br>
     * <p>*   Standard</p>
     * <p>*   IA</p>
     * <p>*   Archive</p>
     * <p>*   ColdArchive</p>
     * <br>
     * <p>For more information about storage classes, see [Overview](~~51374~~).</p>
     */
    @NameInMap("x-oss-storage-class")
    public String xOssStorageClass;

    /**
     * <p>The tag of the destination object. You can add multiple tags to the destination object. Example: TagA=A\&TagB=B.</p>
     * <br>
     * <p>>  The tag key and tag value must be URL-encoded. If a key-value pair does not contain an equal sign (=), the tag value is considered an empty string.</p>
     */
    @NameInMap("x-oss-tagging")
    public String xOssTagging;

    /**
     * <p>The method that is used to add tags to the destination object. Default value: Copy. Valid values:</p>
     * <br>
     * <p>*   **Copy**: The tags of the source object are copied to the destination object.</p>
     * <p>*   **Replace**: The tags that you specify in the request are added to the destination object.</p>
     */
    @NameInMap("x-oss-tagging-directive")
    public String xOssTaggingDirective;

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

    public CopyObjectHeaders setXOssCopySource(String xOssCopySource) {
        this.xOssCopySource = xOssCopySource;
        return this;
    }
    public String getXOssCopySource() {
        return this.xOssCopySource;
    }

    public CopyObjectHeaders setXOssCopySourceIfMatch(String xOssCopySourceIfMatch) {
        this.xOssCopySourceIfMatch = xOssCopySourceIfMatch;
        return this;
    }
    public String getXOssCopySourceIfMatch() {
        return this.xOssCopySourceIfMatch;
    }

    public CopyObjectHeaders setXOssCopySourceIfModifiedSince(String xOssCopySourceIfModifiedSince) {
        this.xOssCopySourceIfModifiedSince = xOssCopySourceIfModifiedSince;
        return this;
    }
    public String getXOssCopySourceIfModifiedSince() {
        return this.xOssCopySourceIfModifiedSince;
    }

    public CopyObjectHeaders setXOssCopySourceIfNoneMatch(String xOssCopySourceIfNoneMatch) {
        this.xOssCopySourceIfNoneMatch = xOssCopySourceIfNoneMatch;
        return this;
    }
    public String getXOssCopySourceIfNoneMatch() {
        return this.xOssCopySourceIfNoneMatch;
    }

    public CopyObjectHeaders setXOssCopySourceIfUnmodifiedSince(String xOssCopySourceIfUnmodifiedSince) {
        this.xOssCopySourceIfUnmodifiedSince = xOssCopySourceIfUnmodifiedSince;
        return this;
    }
    public String getXOssCopySourceIfUnmodifiedSince() {
        return this.xOssCopySourceIfUnmodifiedSince;
    }

    public CopyObjectHeaders setXOssForbidOverwrite(String xOssForbidOverwrite) {
        this.xOssForbidOverwrite = xOssForbidOverwrite;
        return this;
    }
    public String getXOssForbidOverwrite() {
        return this.xOssForbidOverwrite;
    }

    public CopyObjectHeaders setXOssMeta(java.util.Map<String, String> xOssMeta) {
        this.xOssMeta = xOssMeta;
        return this;
    }
    public java.util.Map<String, String> getXOssMeta() {
        return this.xOssMeta;
    }

    public CopyObjectHeaders setXOssMetadataDirective(String xOssMetadataDirective) {
        this.xOssMetadataDirective = xOssMetadataDirective;
        return this;
    }
    public String getXOssMetadataDirective() {
        return this.xOssMetadataDirective;
    }

    public CopyObjectHeaders setXOssObjectAcl(String xOssObjectAcl) {
        this.xOssObjectAcl = xOssObjectAcl;
        return this;
    }
    public String getXOssObjectAcl() {
        return this.xOssObjectAcl;
    }

    public CopyObjectHeaders setXOssServerSideEncryption(String xOssServerSideEncryption) {
        this.xOssServerSideEncryption = xOssServerSideEncryption;
        return this;
    }
    public String getXOssServerSideEncryption() {
        return this.xOssServerSideEncryption;
    }

    public CopyObjectHeaders setXOssServerSideEncryptionKeyId(String xOssServerSideEncryptionKeyId) {
        this.xOssServerSideEncryptionKeyId = xOssServerSideEncryptionKeyId;
        return this;
    }
    public String getXOssServerSideEncryptionKeyId() {
        return this.xOssServerSideEncryptionKeyId;
    }

    public CopyObjectHeaders setXOssStorageClass(String xOssStorageClass) {
        this.xOssStorageClass = xOssStorageClass;
        return this;
    }
    public String getXOssStorageClass() {
        return this.xOssStorageClass;
    }

    public CopyObjectHeaders setXOssTagging(String xOssTagging) {
        this.xOssTagging = xOssTagging;
        return this;
    }
    public String getXOssTagging() {
        return this.xOssTagging;
    }

    public CopyObjectHeaders setXOssTaggingDirective(String xOssTaggingDirective) {
        this.xOssTaggingDirective = xOssTaggingDirective;
        return this;
    }
    public String getXOssTaggingDirective() {
        return this.xOssTaggingDirective;
    }

}
