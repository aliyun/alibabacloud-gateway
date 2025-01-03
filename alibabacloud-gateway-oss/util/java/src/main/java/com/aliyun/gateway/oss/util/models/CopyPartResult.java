// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CopyPartResult extends TeaModel {
    @NameInMap("ETag")
    public String ETag;

    /**
     * <p>Use the UTC time format: yyyy-MM-ddTHH:mmZ</p>
     */
    @NameInMap("LastModified")
    public String lastModified;

    public static CopyPartResult build(java.util.Map<String, ?> map) throws Exception {
        CopyPartResult self = new CopyPartResult();
        return TeaModel.build(map, self);
    }

    public CopyPartResult setETag(String ETag) {
        this.ETag = ETag;
        return this;
    }
    public String getETag() {
        return this.ETag;
    }

    public CopyPartResult setLastModified(String lastModified) {
        this.lastModified = lastModified;
        return this;
    }
    public String getLastModified() {
        return this.lastModified;
    }

}
