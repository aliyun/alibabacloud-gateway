// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class PutBucketRedundancyTypeRequest extends TeaModel {
    @NameInMap("Type")
    public String type;

    public static PutBucketRedundancyTypeRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketRedundancyTypeRequest self = new PutBucketRedundancyTypeRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketRedundancyTypeRequest setType(String type) {
        this.type = type;
        return this;
    }
    public String getType() {
        return this.type;
    }

}
