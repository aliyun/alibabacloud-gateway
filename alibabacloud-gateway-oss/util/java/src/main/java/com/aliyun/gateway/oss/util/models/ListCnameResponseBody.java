// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class ListCnameResponseBody extends TeaModel {
    /**
     * <p>The name of the bucket to which the CNAME records are mapped.</p>
     */
    @NameInMap("Bucket")
    public String bucket;

    /**
     * <p>The container that stores the information about the CNAME records.</p>
     */
    @NameInMap("Cname")
    public java.util.List<CnameInfo> cname;

    /**
     * <p>The name of the bucket owner.</p>
     */
    @NameInMap("Owner")
    public String owner;

    public static ListCnameResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListCnameResponseBody self = new ListCnameResponseBody();
        return TeaModel.build(map, self);
    }

    public ListCnameResponseBody setBucket(String bucket) {
        this.bucket = bucket;
        return this;
    }
    public String getBucket() {
        return this.bucket;
    }

    public ListCnameResponseBody setCname(java.util.List<CnameInfo> cname) {
        this.cname = cname;
        return this;
    }
    public java.util.List<CnameInfo> getCname() {
        return this.cname;
    }

    public ListCnameResponseBody setOwner(String owner) {
        this.owner = owner;
        return this;
    }
    public String getOwner() {
        return this.owner;
    }

}
