// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class GetBucketAclResponseBody extends TeaModel {
    /**
     * <p>The container that stores the ACL information.</p>
     */
    @NameInMap("AccessControlPolicy")
    public AccessControlPolicy accessControlPolicy;

    public static GetBucketAclResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketAclResponseBody self = new GetBucketAclResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketAclResponseBody setAccessControlPolicy(AccessControlPolicy accessControlPolicy) {
        this.accessControlPolicy = accessControlPolicy;
        return this;
    }
    public AccessControlPolicy getAccessControlPolicy() {
        return this.accessControlPolicy;
    }

    public static class AccessControlList extends TeaModel {
        /**
         * <p>The ACL of the bucket.</p>
         */
        @NameInMap("Grant")
        public String grant;

        public static AccessControlList build(java.util.Map<String, ?> map) throws Exception {
            AccessControlList self = new AccessControlList();
            return TeaModel.build(map, self);
        }

        public AccessControlList setGrant(String grant) {
            this.grant = grant;
            return this;
        }
        public String getGrant() {
            return this.grant;
        }

    }

    public static class AccessControlPolicy extends TeaModel {
        /**
         * <p>The class of the container that stores the ACL information.</p>
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
