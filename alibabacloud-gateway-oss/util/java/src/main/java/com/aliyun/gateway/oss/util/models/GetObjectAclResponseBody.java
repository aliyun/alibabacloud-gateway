// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetObjectAclResponseBody extends TeaModel {
    /**
     * <p>The container used to store the ACL information.</p>
     */
    @NameInMap("AccessControlList")
    public GetObjectAclResponseBodyAccessControlList accessControlList;

    /**
     * <p>The container used to store the information about the bucket owner.</p>
     */
    @NameInMap("Owner")
    public Owner owner;

    public static GetObjectAclResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetObjectAclResponseBody self = new GetObjectAclResponseBody();
        return TeaModel.build(map, self);
    }

    public GetObjectAclResponseBody setAccessControlList(GetObjectAclResponseBodyAccessControlList accessControlList) {
        this.accessControlList = accessControlList;
        return this;
    }
    public GetObjectAclResponseBodyAccessControlList getAccessControlList() {
        return this.accessControlList;
    }

    public GetObjectAclResponseBody setOwner(Owner owner) {
        this.owner = owner;
        return this;
    }
    public Owner getOwner() {
        return this.owner;
    }

    public static class GetObjectAclResponseBodyAccessControlList extends TeaModel {
        /**
         * <p>The ACL of the object.</p>
         * <p><br>Valid values: default, private, public-read, and public-read-write.</p>
         */
        @NameInMap("Grant")
        public String grant;

        public static GetObjectAclResponseBodyAccessControlList build(java.util.Map<String, ?> map) throws Exception {
            GetObjectAclResponseBodyAccessControlList self = new GetObjectAclResponseBodyAccessControlList();
            return TeaModel.build(map, self);
        }

        public GetObjectAclResponseBodyAccessControlList setGrant(String grant) {
            this.grant = grant;
            return this;
        }
        public String getGrant() {
            return this.grant;
        }

    }

}
