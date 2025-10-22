// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class LoggingEnabled extends TeaModel {
    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("TargetBucket")
    public String targetBucket;

    @NameInMap("TargetPrefix")
    public String targetPrefix;

    @NameInMap("LoggingRole")
    public String loggingRole;

    public static LoggingEnabled build(java.util.Map<String, ?> map) throws Exception {
        LoggingEnabled self = new LoggingEnabled();
        return TeaModel.build(map, self);
    }

    public LoggingEnabled setTargetBucket(String targetBucket) {
        this.targetBucket = targetBucket;
        return this;
    }
    public String getTargetBucket() {
        return this.targetBucket;
    }

    public LoggingEnabled setTargetPrefix(String targetPrefix) {
        this.targetPrefix = targetPrefix;
        return this;
    }
    public String getTargetPrefix() {
        return this.targetPrefix;
    }

    public LoggingEnabled setLoggingRole(String loggingRole) {
        this.loggingRole = loggingRole;
        return this;
    }
    public String getLoggingRole() {
        return this.loggingRole;
    }

}
