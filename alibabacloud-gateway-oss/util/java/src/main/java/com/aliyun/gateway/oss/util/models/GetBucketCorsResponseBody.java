// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetBucketCorsResponseBody extends TeaModel {
    /**
     * <p>The container that stores the CORS rules. Up to 10 rules can be configured for a bucket.</p>
     */
    @NameInMap("CORSRule")
    public java.util.List<CORSRule> CORSRule;

    /**
     * <p>Indicates whether the Vary: Origin header is returned. Valid values:</p>
     * <br>
     * <p>*   **true**: The Vary: Origin header is returned regardless whether the request is a cross-origin request or whether the cross-origin request succeeds.</p>
     * <p>*   **false** (default): The Vary: Origin header is not returned.</p>
     * <br>
     * <p>>  This parameter is valid only when at least one CORS rule is configured.</p>
     */
    @NameInMap("ResponseVary")
    public Boolean responseVary;

    public static GetBucketCorsResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketCorsResponseBody self = new GetBucketCorsResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketCorsResponseBody setCORSRule(java.util.List<CORSRule> CORSRule) {
        this.CORSRule = CORSRule;
        return this;
    }
    public java.util.List<CORSRule> getCORSRule() {
        return this.CORSRule;
    }

    public GetBucketCorsResponseBody setResponseVary(Boolean responseVary) {
        this.responseVary = responseVary;
        return this;
    }
    public Boolean getResponseVary() {
        return this.responseVary;
    }

}
