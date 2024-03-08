// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutObjectHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>Specifies whether the object that is uploaded by calling the PutObject operation overwrites the existing object that has the same name.  When versioning is enabled or suspended for the bucket to which you want to upload the object, the **x-oss-forbid-overwrite** header does not take effect. In this case, the object that is uploaded by calling the PutObject operation overwrites the existing object that has the same name. </p>
     * <p>  - If you do not specify the **x-oss-forbid-overwrite** header or set the **x-oss-forbid-overwrite** header to **false**, the object that is uploaded by calling the PutObject operation overwrites the existing object that has the same name. </p>
     * <p>  - If the value of **x-oss-forbid-overwrite** is set to **true**, existing objects cannot be overwritten by objects that have the same names. </p>
     * <br>
     * <p>If you specify the **x-oss-forbid-overwrite** request header, the queries per second (QPS) performance of OSS is degraded. If you want to use the **x-oss-forbid-overwrite** request header to perform a large number of operations (QPS greater than 1,000), contact technical support. </p>
     * <p>Default value: **false**.</p>
     */
    @NameInMap("x-oss-forbid-overwrite")
    public Boolean xOssForbidOverwrite;

    /**
     * <p>If the PutObject request contains a parameter prefixed with **x-oss-meta-***, the parameter is considered as the user metadata. Example: `x-oss-meta-location`. You can specify multiple similar headers for an object. However, the user metadata of an object cannot exceed 8 KB in size. </p>
     * <br>
     * <p>The names of user metadata headers can contain letters, digits, and hyphens (-). Uppercase letters are converted to lowercase letters. Other characters such as underscores (_) are not supported.</p>
     */
    @NameInMap("x-oss-meta-*")
    public java.util.Map<String, String> xOssMeta;

    /**
     * <p>The access control list (ACL) of the object. Default value: default. </p>
     * <p>Valid values:</p>
     * <br>
     * <p>- default: The ACL of the object is the same as that of the bucket in which the object is stored. </p>
     * <p>- private: The ACL of the object is private. Only the owner of the object and authorized users can read and write this object. </p>
     * <p>- public-read: The ACL of the object is public-read. Only the owner of the object and authorized users can read and write this object. Other users can only read the object. Exercise caution when you set the object ACL to this value. </p>
     * <p>- public-read-write: The ACL of the object is public-read-write. All users can read and write this object. Exercise caution when you set the object ACL to this value. </p>
     * <br>
     * <p>For more information about the ACL, see **[ACL](~~100676~~)**.</p>
     */
    @NameInMap("x-oss-object-acl")
    public String xOssObjectAcl;

    /**
     * <p>The encryption method on the server side when an object is created. </p>
     * <br>
     * <p>Valid values: **AES256**, **KMS**, and **SM4**.</p>
     * <br>
     * <p>If you specify the header, the header is returned in the response. OSS uses the method that is specified by this header to encrypt the uploaded object. When you download the encrypted object, the **x-oss-server-side-encryption** header is included in the response and the header value is set to the algorithm that is used to encrypt the object.</p>
     */
    @NameInMap("x-oss-server-side-data-encryption")
    public String xOssServerSideDataEncryption;

    /**
     * <p>The method that is used to encrypt the object on the OSS server when the object is created. </p>
     * <br>
     * <p>Valid values: **AES256**, **KMS**, and **SM4****.</p>
     * <br>
     * <p>If you specify the header, the header is returned in the response. OSS uses the method that is specified by this header to encrypt the uploaded object. When you download the encrypted object, the **x-oss-server-side-encryption** header is included in the response and the header value is set to the algorithm that is used to encrypt the object.</p>
     */
    @NameInMap("x-oss-server-side-encryption")
    public String xOssServerSideEncryption;

    /**
     * <p>The ID of the customer master key (CMK) managed by Key Management Service (KMS). </p>
     * <p>This header is valid only when the **x-oss-server-side-encryption** header is set to KMS.</p>
     */
    @NameInMap("x-oss-server-side-encryption-key-id")
    public String xOssServerSideEncryptionKeyId;

    /**
     * <p>The storage class of the bucket. Default value: Standard.  Valid values:</p>
     * <br>
     * <p>- Standard</p>
     * <p>- IA</p>
     * <p>- Archive</p>
     * <p>- ColdArchive</p>
     */
    @NameInMap("x-oss-storage-class")
    public String xOssStorageClass;

    /**
     * <p>The tag of the object. You can configure multiple tags for the object. Example: TagA=A&TagB=B. </p>
     * <p>> The key and value of a tag must be URL-encoded. If a tag does not contain an equal sign (=), the value of the tag is considered an empty string.</p>
     */
    @NameInMap("x-oss-tagging")
    public String xOssTagging;

    public static PutObjectHeaders build(java.util.Map<String, ?> map) throws Exception {
        PutObjectHeaders self = new PutObjectHeaders();
        return TeaModel.build(map, self);
    }

    public PutObjectHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public PutObjectHeaders setXOssForbidOverwrite(Boolean xOssForbidOverwrite) {
        this.xOssForbidOverwrite = xOssForbidOverwrite;
        return this;
    }
    public Boolean getXOssForbidOverwrite() {
        return this.xOssForbidOverwrite;
    }

    public PutObjectHeaders setXOssMeta(java.util.Map<String, String> xOssMeta) {
        this.xOssMeta = xOssMeta;
        return this;
    }
    public java.util.Map<String, String> getXOssMeta() {
        return this.xOssMeta;
    }

    public PutObjectHeaders setXOssObjectAcl(String xOssObjectAcl) {
        this.xOssObjectAcl = xOssObjectAcl;
        return this;
    }
    public String getXOssObjectAcl() {
        return this.xOssObjectAcl;
    }

    public PutObjectHeaders setXOssServerSideDataEncryption(String xOssServerSideDataEncryption) {
        this.xOssServerSideDataEncryption = xOssServerSideDataEncryption;
        return this;
    }
    public String getXOssServerSideDataEncryption() {
        return this.xOssServerSideDataEncryption;
    }

    public PutObjectHeaders setXOssServerSideEncryption(String xOssServerSideEncryption) {
        this.xOssServerSideEncryption = xOssServerSideEncryption;
        return this;
    }
    public String getXOssServerSideEncryption() {
        return this.xOssServerSideEncryption;
    }

    public PutObjectHeaders setXOssServerSideEncryptionKeyId(String xOssServerSideEncryptionKeyId) {
        this.xOssServerSideEncryptionKeyId = xOssServerSideEncryptionKeyId;
        return this;
    }
    public String getXOssServerSideEncryptionKeyId() {
        return this.xOssServerSideEncryptionKeyId;
    }

    public PutObjectHeaders setXOssStorageClass(String xOssStorageClass) {
        this.xOssStorageClass = xOssStorageClass;
        return this;
    }
    public String getXOssStorageClass() {
        return this.xOssStorageClass;
    }

    public PutObjectHeaders setXOssTagging(String xOssTagging) {
        this.xOssTagging = xOssTagging;
        return this;
    }
    public String getXOssTagging() {
        return this.xOssTagging;
    }

}
