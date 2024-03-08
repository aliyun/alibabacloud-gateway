// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetBucketLoggingResponseBody extends TeaModel {
    /**
     * <p>The container that stores the information about access log collection. This parameter is returned if access log collection is enabled for the bucket.</p>
     */
    @NameInMap("LoggingEnabled")
    public LoggingEnabled loggingEnabled;

    public static GetBucketLoggingResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketLoggingResponseBody self = new GetBucketLoggingResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketLoggingResponseBody setLoggingEnabled(LoggingEnabled loggingEnabled) {
        this.loggingEnabled = loggingEnabled;
        return this;
    }
    public LoggingEnabled getLoggingEnabled() {
        return this.loggingEnabled;
    }

}
