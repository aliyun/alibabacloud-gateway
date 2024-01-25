// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class ListBucketInventoryRequest extends TeaModel {
    /**
     * <p>> </p>
     * <br>
     * <p>*   Syntax of requests that contain continuation-token: GET /?inventory\&continuation-token=xxx HTTP/1.1</p>
     * <br>
     * <p>*   Syntax of requests that do not contain continuation-token: `GET /?inventory HTTP/1.1`</p>
     */
    @NameInMap("continuation-token")
    public String continuationToken;

    public static ListBucketInventoryRequest build(java.util.Map<String, ?> map) throws Exception {
        ListBucketInventoryRequest self = new ListBucketInventoryRequest();
        return TeaModel.build(map, self);
    }

    public ListBucketInventoryRequest setContinuationToken(String continuationToken) {
        this.continuationToken = continuationToken;
        return this;
    }
    public String getContinuationToken() {
        return this.continuationToken;
    }

}
