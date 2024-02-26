// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class PutObjectHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>Specifies whether the object that is uploaded by calling the PutObject operation overwrites the existing object that has the same name.  When versioning is enabled or suspended for the bucket to which you want to upload the object, the **x-oss-forbid-overwrite** header does not take effect. In this case, the object that is uploaded by calling the PutObject operation overwrites the existing object that has the same name. 
</p>
     * <p>  - If you do not specify the **x-oss-forbid-overwrite** header or set the **x-oss-forbid-overwrite** header to **false**, the object that is uploaded by calling the PutObject operation overwrites the existing object that has the same name. 
</p>
     * <p>  - If the value of **x-oss-forbid-overwrite** is set to **true**, existing objects cannot be overwritten by objects that have the same names. 
</p>
     * <p>
</p>
     * <p>If you specify the **x-oss-forbid-overwrite** request header, the queries per second (QPS) performance of OSS is degraded. If you want to use the **x-oss-forbid-overwrite** request header to perform a large number of operations (QPS greater than 1,000), contact technical support. 
</p>
     * <p>Default value: **false**.</p>
     */
    @NameInMap("x-oss-forbid-overwrite")
    public Boolean forbidOverwrite;

    /**
     * <p>If the PutObject request contains a parameter prefixed with **x-oss-meta-***, the parameter is considered as the user metadata. Example: `x-oss-meta-location`. You can specify multiple similar headers for an object. However, the user metadata of an object cannot exceed 8 KB in size. 
</p>
     * <p>
</p>
     * <p>The names of user metadata headers can contain letters, digits, and hyphens (-). Uppercase letters are converted to lowercase letters. Other characters such as underscores (_) are not supported.</p>
     */
    @NameInMap("x-oss-meta-*")
    public java.util.Map<String, String> metaData;

    /**
     * <p>The access control list (ACL) of the object. Default value: default. 
</p>
     * <p>Valid values:
</p>
     * <p>
</p>
     * <p>- default: The ACL of the object is the same as that of the bucket in which the object is stored. 
</p>
     * <p>- private: The ACL of the object is private. Only the owner of the object and authorized users can read and write this object. 
</p>
     * <p>- public-read: The ACL of the object is public-read. Only the owner of the object and authorized users can read and write this object. Other users can only read the object. Exercise caution when you set the object ACL to this value. 
</p>
     * <p>- public-read-write: The ACL of the object is public-read-write. All users can read and write this object. Exercise caution when you set the object ACL to this value. 
</p>
     * <p>
</p>
     * <p>For more information about the ACL, see **[ACL](~~100676~~)**.</p>
     */
    @NameInMap("x-oss-object-acl")
    public String acl;

    /**
     * <p>The encryption method on the server side when an object is created. 
</p>
     * <p>
</p>
     * <p>Valid values: **AES256**, **KMS**, and **SM4**.
</p>
     * <p>
</p>
     * <p>If you specify the header, the header is returned in the response. OSS uses the method that is specified by this header to encrypt the uploaded object. When you download the encrypted object, the **x-oss-server-side-encryption** header is included in the response and the header value is set to the algorithm that is used to encrypt the object.</p>
     */
    @NameInMap("x-oss-server-side-data-encryption")
    public String sseDataEncryption;

    /**
     * <p>The method that is used to encrypt the object on the OSS server when the object is created. </p>
     * <br>
     * <p>Valid values: **AES256**, **KMS**, and **SM4****.</p>
     * <br>
     * <p>If you specify the header, the header is returned in the response. OSS uses the method that is specified by this header to encrypt the uploaded object. When you download the encrypted object, the **x-oss-server-side-encryption** header is included in the response and the header value is set to the algorithm that is used to encrypt the object.</p>
     */
    @NameInMap("x-oss-server-side-encryption")
    public String serverSideEncryption;

    /**
     * <p>The ID of the customer master key (CMK) managed by Key Management Service (KMS). 
</p>
     * <p>This header is valid only when the **x-oss-server-side-encryption** header is set to KMS.</p>
     */
    @NameInMap("x-oss-server-side-encryption-key-id")
    public String sseKeyId;

    /**
     * <p>The storage class of the bucket. Default value: Standard.  Valid values:
</p>
     * <p>
</p>
     * <p>- Standard
</p>
     * <p>- IA
</p>
     * <p>- Archive
</p>
     * <p>- ColdArchive</p>
     */
    @NameInMap("x-oss-storage-class")
    public String storageClass;

    /**
     * <p>The tag of the object. You can configure multiple tags for the object. Example: TagA=A&TagB=B. 
</p>
     * <p>> The key and value of a tag must be URL-encoded. If a tag does not contain an equal sign (=), the value of the tag is considered an empty string.</p>
     */
    @NameInMap("x-oss-tagging")
    public String tagging;

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

    public PutObjectHeaders setForbidOverwrite(Boolean forbidOverwrite) {
        this.forbidOverwrite = forbidOverwrite;
        return this;
    }
    public Boolean getForbidOverwrite() {
        return this.forbidOverwrite;
    }

    public PutObjectHeaders setMetaData(java.util.Map<String, String> metaData) {
        this.metaData = metaData;
        return this;
    }
    public java.util.Map<String, String> getMetaData() {
        return this.metaData;
    }

    public PutObjectHeaders setAcl(String acl) {
        this.acl = acl;
        return this;
    }
    public String getAcl() {
        return this.acl;
    }

    public PutObjectHeaders setSseDataEncryption(String sseDataEncryption) {
        this.sseDataEncryption = sseDataEncryption;
        return this;
    }
    public String getSseDataEncryption() {
        return this.sseDataEncryption;
    }

    public PutObjectHeaders setServerSideEncryption(String serverSideEncryption) {
        this.serverSideEncryption = serverSideEncryption;
        return this;
    }
    public String getServerSideEncryption() {
        return this.serverSideEncryption;
    }

    public PutObjectHeaders setSseKeyId(String sseKeyId) {
        this.sseKeyId = sseKeyId;
        return this;
    }
    public String getSseKeyId() {
        return this.sseKeyId;
    }

    public PutObjectHeaders setStorageClass(String storageClass) {
        this.storageClass = storageClass;
        return this;
    }
    public String getStorageClass() {
        return this.storageClass;
    }

    public PutObjectHeaders setTagging(String tagging) {
        this.tagging = tagging;
        return this;
    }
    public String getTagging() {
        return this.tagging;
    }

}
