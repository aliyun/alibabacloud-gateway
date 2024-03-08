// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class CopyObjectResponseBody extends TeaModel {
    /**
     * <p>The ETag value of the destination object.</p>
     */
    @NameInMap("ETag")
    public String ETag;

    /**
     * <p>The time when the destination object was last modified.</p>
     */
    @NameInMap("LastModified")
    public String lastModified;

    public static CopyObjectResponseBody build(java.util.Map<String, ?> map) throws Exception {
        CopyObjectResponseBody self = new CopyObjectResponseBody();
        return TeaModel.build(map, self);
    }

    public CopyObjectResponseBody setETag(String ETag) {
        this.ETag = ETag;
        return this;
    }
    public String getETag() {
        return this.ETag;
    }

    public CopyObjectResponseBody setLastModified(String lastModified) {
        this.lastModified = lastModified;
        return this;
    }
    public String getLastModified() {
        return this.lastModified;
    }

}
