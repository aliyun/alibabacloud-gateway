// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListCnameResponseBody extends TeaModel {
    /**
     * <p>The container that is used to query information about all CNAME records.</p>
     */
    @NameInMap("ListCnameResult")
    public ListCnameResult listCnameResult;

    public static ListCnameResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListCnameResponseBody self = new ListCnameResponseBody();
        return TeaModel.build(map, self);
    }

    public ListCnameResponseBody setListCnameResult(ListCnameResult listCnameResult) {
        this.listCnameResult = listCnameResult;
        return this;
    }
    public ListCnameResult getListCnameResult() {
        return this.listCnameResult;
    }

    public static class ListCnameResult extends TeaModel {
        /**
         * <p>The name of the bucket to which the CNAME records you want to query are mapped.</p>
         */
        @NameInMap("Bucket")
        public String bucket;

        /**
         * <p>The container that is used to store the information about all CNAME records.</p>
         */
        @NameInMap("Cname")
        public java.util.List<CnameInfo> cname;

        /**
         * <p>The name of the bucket owner.</p>
         */
        @NameInMap("Owner")
        public String owner;

        public static ListCnameResult build(java.util.Map<String, ?> map) throws Exception {
            ListCnameResult self = new ListCnameResult();
            return TeaModel.build(map, self);
        }

        public ListCnameResult setBucket(String bucket) {
            this.bucket = bucket;
            return this;
        }
        public String getBucket() {
            return this.bucket;
        }

        public ListCnameResult setCname(java.util.List<CnameInfo> cname) {
            this.cname = cname;
            return this;
        }
        public java.util.List<CnameInfo> getCname() {
            return this.cname;
        }

        public ListCnameResult setOwner(String owner) {
            this.owner = owner;
            return this;
        }
        public String getOwner() {
            return this.owner;
        }

    }

}
