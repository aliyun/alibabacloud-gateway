// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class SSEOSS extends TeaModel {
    @NameInMap("unused")
    public String unused;

    public static SSEOSS build(java.util.Map<String, ?> map) throws Exception {
        SSEOSS self = new SSEOSS();
        return TeaModel.build(map, self);
    }

    public SSEOSS setUnused(String unused) {
        this.unused = unused;
        return this;
    }
    public String getUnused() {
        return this.unused;
    }

}
