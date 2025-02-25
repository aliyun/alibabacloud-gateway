// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutSymlinkHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>Specifies whether the PutSymlink operation overwrites the object that has the same name as that of the symbolic link you want to create. </p>
     * <ul>
     * <li>If the value of <strong>x-oss-forbid-overwrite</strong> is not specified or set to <strong>false</strong>, existing objects can be overwritten by objects that have the same names. </li>
     * <li>If the value of <strong>x-oss-forbid-overwrite</strong> is set to <strong>true</strong>, existing objects cannot be overwritten by objects that have the same names.</li>
     * </ul>
     * <p>If you specify the <strong>x-oss-forbid-overwrite</strong> request header, the queries per second (QPS) performance of OSS is degraded. If you want to use the <strong>x-oss-forbid-overwrite</strong> request header to perform a large number of operations (QPS greater than 1,000), contact technical support. </p>
     * <blockquote>
     * <p>The <strong>x-oss-forbid-overwrite</strong> request header is invalid when versioning is enabled or suspended for the destination bucket. In this case, the object with the same name can be overwritten.</p>
     * </blockquote>
     * 
     * <strong>example:</strong>
     * <p>true</p>
     */
    @NameInMap("x-oss-forbid-overwrite")
    public String forbidOverwrite;

    /**
     * <p>The access control list (ACL) of the object. Default value: default. 
     * Valid values:</p>
     * <ul>
     * <li>default: The ACL of the object is the same as that of the bucket in which the object is stored. </li>
     * <li>private: The ACL of the object is private. Only the owner of the object and authorized users can read and write this object. </li>
     * <li>public-read: The ACL of the object is public-read. Only the owner of the object and authorized users can read and write this object. Other users can only read the object. Exercise caution when you set the object ACL to this value. </li>
     * <li>public-read-write: The ACL of the object is public-read-write. All users can read and write this object. Exercise caution when you set the object ACL to this value.</li>
     * </ul>
     * <p>For more information about the ACL, see <strong><a href="https://help.aliyun.com/document_detail/100676.html">ACL</a></strong>.</p>
     */
    @NameInMap("x-oss-object-acl")
    public String acl;

    /**
     * <p>The storage class of the bucket. Default value: Standard.  Valid values:</p>
     * <ul>
     * <li>Standard</li>
     * <li>IA</li>
     * <li>Archive</li>
     * <li>ColdArchive</li>
     * </ul>
     */
    @NameInMap("x-oss-storage-class")
    public String storageClass;

    /**
     * <p>The target object to which the symbolic link points. 
     * The naming conventions for target objects are the same as those for objects.</p>
     * <ul>
     * <li>Similar to ObjectName, TargetObjectName must be URL-encoded. </li>
     * <li>The target object to which a symbolic link points cannot be a symbolic link.</li>
     * </ul>
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>oss.jpg</p>
     */
    @NameInMap("x-oss-symlink-target")
    public String symlinkTargetKey;

    public static PutSymlinkHeaders build(java.util.Map<String, ?> map) throws Exception {
        PutSymlinkHeaders self = new PutSymlinkHeaders();
        return TeaModel.build(map, self);
    }

    public PutSymlinkHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public PutSymlinkHeaders setForbidOverwrite(String forbidOverwrite) {
        this.forbidOverwrite = forbidOverwrite;
        return this;
    }
    public String getForbidOverwrite() {
        return this.forbidOverwrite;
    }

    public PutSymlinkHeaders setAcl(String acl) {
        this.acl = acl;
        return this;
    }
    public String getAcl() {
        return this.acl;
    }

    public PutSymlinkHeaders setStorageClass(String storageClass) {
        this.storageClass = storageClass;
        return this;
    }
    public String getStorageClass() {
        return this.storageClass;
    }

    public PutSymlinkHeaders setSymlinkTargetKey(String symlinkTargetKey) {
        this.symlinkTargetKey = symlinkTargetKey;
        return this;
    }
    public String getSymlinkTargetKey() {
        return this.symlinkTargetKey;
    }

}
