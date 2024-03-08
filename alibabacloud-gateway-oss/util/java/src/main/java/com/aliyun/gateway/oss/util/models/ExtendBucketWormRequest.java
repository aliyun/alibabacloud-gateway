// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class ExtendBucketWormRequest extends TeaModel {
    /**
     * <p>The container that stores the information about the retention policy.</p>
     */
    @NameInMap("body")
    public ExtendWormConfiguration body;

    /**
     * <p>The ID of the retention policy.</p>
     * <br>
     * <p>>  If the ID of the retention policy that specifies the number of days for which objects can be retained does not exist, the HTTP status code 404 is returned.</p>
     */
    @NameInMap("wormId")
    public String wormId;

    public static ExtendBucketWormRequest build(java.util.Map<String, ?> map) throws Exception {
        ExtendBucketWormRequest self = new ExtendBucketWormRequest();
        return TeaModel.build(map, self);
    }

    public ExtendBucketWormRequest setBody(ExtendWormConfiguration body) {
        this.body = body;
        return this;
    }
    public ExtendWormConfiguration getBody() {
        return this.body;
    }

    public ExtendBucketWormRequest setWormId(String wormId) {
        this.wormId = wormId;
        return this;
    }
    public String getWormId() {
        return this.wormId;
    }

}
