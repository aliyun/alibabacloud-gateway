// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class GetObjectAclResponseBody extends TeaModel {
    /**
     * <p>The container that stores the results of the GetObjectACL request.</p>
     */
    @NameInMap("AccessControlPolicy")
    public AccessControlPolicy accessControlPolicy;

    public static GetObjectAclResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetObjectAclResponseBody self = new GetObjectAclResponseBody();
        return TeaModel.build(map, self);
    }

    public GetObjectAclResponseBody setAccessControlPolicy(AccessControlPolicy accessControlPolicy) {
        this.accessControlPolicy = accessControlPolicy;
        return this;
    }
    public AccessControlPolicy getAccessControlPolicy() {
        return this.accessControlPolicy;
    }

    public static class AccessControlList extends TeaModel {
        /**
         * <p>The ACL of the object. Default value: default.</p>
         */
        @NameInMap("Grant")
        public String ACL;

        public static AccessControlList build(java.util.Map<String, ?> map) throws Exception {
            AccessControlList self = new AccessControlList();
            return TeaModel.build(map, self);
        }

        public AccessControlList setACL(String ACL) {
            this.ACL = ACL;
            return this;
        }
        public String getACL() {
            return this.ACL;
        }

    }

    public static class AccessControlPolicy extends TeaModel {
        /**
         * <p>The container that stores the ACL information.</p>
         */
        @NameInMap("AccessControlList")
        public AccessControlList accessControlList;

        /**
         * <p>The container that stores the information about the bucket owner.</p>
         */
        @NameInMap("Owner")
        public Owner owner;

        public static AccessControlPolicy build(java.util.Map<String, ?> map) throws Exception {
            AccessControlPolicy self = new AccessControlPolicy();
            return TeaModel.build(map, self);
        }

        public AccessControlPolicy setAccessControlList(AccessControlList accessControlList) {
            this.accessControlList = accessControlList;
            return this;
        }
        public AccessControlList getAccessControlList() {
            return this.accessControlList;
        }

        public AccessControlPolicy setOwner(Owner owner) {
            this.owner = owner;
            return this;
        }
        public Owner getOwner() {
            return this.owner;
        }

    }

}
