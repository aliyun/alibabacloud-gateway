// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class ListUserDataRedundancyTransitionRequest extends TeaModel {
    @NameInMap("continuation-token")
    public String continuationToken;

    @NameInMap("max-keys")
    public Integer maxKeys;

    public static ListUserDataRedundancyTransitionRequest build(java.util.Map<String, ?> map) throws Exception {
        ListUserDataRedundancyTransitionRequest self = new ListUserDataRedundancyTransitionRequest();
        return TeaModel.build(map, self);
    }

    public ListUserDataRedundancyTransitionRequest setContinuationToken(String continuationToken) {
        this.continuationToken = continuationToken;
        return this;
    }
    public String getContinuationToken() {
        return this.continuationToken;
    }

    public ListUserDataRedundancyTransitionRequest setMaxKeys(Integer maxKeys) {
        this.maxKeys = maxKeys;
        return this;
    }
    public Integer getMaxKeys() {
        return this.maxKeys;
    }

}
