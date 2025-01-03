// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class AppendObjectHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>The web page caching behavior for the object. For more information, see <strong><a href="https://www.ietf.org/rfc/rfc2616.txt">RFC 2616</a></strong>. 
     * Default value: null.</p>
     * 
     * <strong>example:</strong>
     * <p>no-cache</p>
     */
    @NameInMap("Cache-Control")
    public String cacheControl;

    /**
     * <p>The name of the object when the object is downloaded. For more information, see <strong><a href="https://www.ietf.org/rfc/rfc2616.txt">RFC 2616</a></strong>. 
     * Default value: null.</p>
     * 
     * <strong>example:</strong>
     * <p>attachment;filename=oss_download.jpg</p>
     */
    @NameInMap("Content-Disposition")
    public String contentDisposition;

    /**
     * <p>The encoding format of the object content. For more information, see <strong><a href="https://www.ietf.org/rfc/rfc2616.txt">RFC 2616</a></strong>. 
     * Default value: null.</p>
     * 
     * <strong>example:</strong>
     * <p>utf-8</p>
     */
    @NameInMap("Content-Encoding")
    public String contentEncoding;

    /**
     * <p>The Content-MD5 header value is a string calculated by using the MD5 algorithm. The header is used to check whether the content of the received message is the same as that of the sent message. 
     * To obtain the value of the Content-MD5 header, calculate a 128-bit number based on the message content except for the header, and then encode the number in Base64. 
     * Default value: null.
     * Limits: none.</p>
     * 
     * <strong>example:</strong>
     * <p>ohhnqLBJFiKkPSBO1eNaUA==</p>
     */
    @NameInMap("Content-MD5")
    public String contentMD5;

    /**
     * <p>The expiration time. For more information, see <strong><a href="https://www.ietf.org/rfc/rfc2616.txt">RFC 2616</a></strong>. 
     * Default value: null.</p>
     * 
     * <strong>example:</strong>
     * <p>Wed, 08 Jul 2015 16:57:01 GMT</p>
     */
    @NameInMap("Expires")
    public String expires;

    /**
     * <p>You can add parameters whose names are prefixed with x-oss-meta-* when you call the AppendObject operation. These parameters cannot be included in the requests when you append objects to an existing object. Parameters whose names are prefixed with x-oss-meta-* are considered the metadata of the object. 
     * You can configure multiple parameters whose name are prefixed with x-oss-meta- for an object. However, the total size of user metadata cannot exceed 8 KB. 
     * The name of parameters whose name are prefixed with x-oss-meta- can contain hyphens (-), numbers, and letters. Uppercase letters are converted to lowercase letters. Other characters such as underscores (_) are not supported.</p>
     * 
     * <strong>example:</strong>
     * <p>x-oss-meta-location</p>
     */
    @NameInMap("x-oss-meta-*")
    public java.util.Map<String, String> metaData;

    /**
     * <p>The access control list (ACL) of the object. Default value: default.  Valid values:</p>
     * <ul>
     * <li>default: The ACL of the object is the same as that of the bucket in which the object is stored. </li>
     * <li>private: The ACL of the object is private. Only the owner of the object and authorized users can read and write this object. </li>
     * <li>public-read: The ACL of the object is public-read. Only the owner of the object and authorized users can read and write this object. Other users can only read the object. Exercise caution when you set the object ACL to this value. </li>
     * <li>public-read-write: The ACL of the object is public-read-write. All users can read and write this object. Exercise caution when you set the object ACL to this value.</li>
     * </ul>
     * <p>For more information about the ACL, see <a href="https://help.aliyun.com/document_detail/100676.html">ACL</a>.</p>
     * 
     * <strong>example:</strong>
     * <p>default</p>
     */
    @NameInMap("x-oss-object-acl")
    public String acl;

    /**
     * <p>The method used to encrypt objects on the specified OSS server. 
     * Valid values:</p>
     * <ul>
     * <li>AES256: Keys managed by OSS are used for encryption and decryption (SSE-OSS). </li>
     * <li>KMS: Keys managed by Key Management Service (KMS) are used for encryption and decryption. </li>
     * <li>SM4: The SM4 block cipher algorithm is used for encryption and decryption.</li>
     * </ul>
     * 
     * <strong>example:</strong>
     * <p>AES256</p>
     */
    @NameInMap("x-oss-server-side-encryption")
    public String serverSideEncryption;

    /**
     * <p>The storage class of the object that you want to upload. Valid values:</p>
     * <ul>
     * <li>Standard</li>
     * <li>IA</li>
     * <li>Archive
     * If you specify the object storage class when you upload an object, the storage class of the uploaded object is the specified value regardless of the storage class of the bucket to which the object is uploaded. If you set x-oss-storage-class to Standard when you upload an object to an IA bucket, the object is stored as a Standard object. 
     * For more information about storage classes, see the &quot;Overview&quot; topic in Developer Guide.</li>
     * </ul>
     * <blockquote>
     * <p>Notice:  The value that you specify takes effect only when you call the AppendObject operation on an object for the first time.</p>
     * </blockquote>
     * 
     * <strong>example:</strong>
     * <p>Standard</p>
     */
    @NameInMap("x-oss-storage-class")
    public String storageClass;

    public static AppendObjectHeaders build(java.util.Map<String, ?> map) throws Exception {
        AppendObjectHeaders self = new AppendObjectHeaders();
        return TeaModel.build(map, self);
    }

    public AppendObjectHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public AppendObjectHeaders setCacheControl(String cacheControl) {
        this.cacheControl = cacheControl;
        return this;
    }
    public String getCacheControl() {
        return this.cacheControl;
    }

    public AppendObjectHeaders setContentDisposition(String contentDisposition) {
        this.contentDisposition = contentDisposition;
        return this;
    }
    public String getContentDisposition() {
        return this.contentDisposition;
    }

    public AppendObjectHeaders setContentEncoding(String contentEncoding) {
        this.contentEncoding = contentEncoding;
        return this;
    }
    public String getContentEncoding() {
        return this.contentEncoding;
    }

    public AppendObjectHeaders setContentMD5(String contentMD5) {
        this.contentMD5 = contentMD5;
        return this;
    }
    public String getContentMD5() {
        return this.contentMD5;
    }

    public AppendObjectHeaders setExpires(String expires) {
        this.expires = expires;
        return this;
    }
    public String getExpires() {
        return this.expires;
    }

    public AppendObjectHeaders setMetaData(java.util.Map<String, String> metaData) {
        this.metaData = metaData;
        return this;
    }
    public java.util.Map<String, String> getMetaData() {
        return this.metaData;
    }

    public AppendObjectHeaders setAcl(String acl) {
        this.acl = acl;
        return this;
    }
    public String getAcl() {
        return this.acl;
    }

    public AppendObjectHeaders setServerSideEncryption(String serverSideEncryption) {
        this.serverSideEncryption = serverSideEncryption;
        return this;
    }
    public String getServerSideEncryption() {
        return this.serverSideEncryption;
    }

    public AppendObjectHeaders setStorageClass(String storageClass) {
        this.storageClass = storageClass;
        return this;
    }
    public String getStorageClass() {
        return this.storageClass;
    }

}
