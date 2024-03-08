// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class PutSymlinkHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>Specifies whether the PutSymlink operation overwrites the object that has the same name as that of the symbolic link you want to create. </p>
     * <p>  - If the value of **x-oss-forbid-overwrite** is not specified or set to **false**, existing objects can be overwritten by objects that have the same names. </p>
     * <p>  - If the value of **x-oss-forbid-overwrite** is set to **true**, existing objects cannot be overwritten by objects that have the same names. </p>
     * <br>
     * <p>If you specify the **x-oss-forbid-overwrite** request header, the queries per second (QPS) performance of OSS is degraded. If you want to use the **x-oss-forbid-overwrite** request header to perform a large number of operations (QPS greater than 1,000), contact technical support. </p>
     * <p>> The **x-oss-forbid-overwrite** request header is invalid when versioning is enabled or suspended for the destination bucket. In this case, the object with the same name can be overwritten.</p>
     */
    @NameInMap("x-oss-forbid-overwrite")
    public String forbidOverwrite;

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
     * <p>The target object to which the symbolic link points. 
</p>
     * <p>The naming conventions for target objects are the same as those for objects.
</p>
     * <p>  - Similar to ObjectName, TargetObjectName must be URL-encoded. 
</p>
     * <p>  - The target object to which a symbolic link points cannot be a symbolic link.</p>
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
