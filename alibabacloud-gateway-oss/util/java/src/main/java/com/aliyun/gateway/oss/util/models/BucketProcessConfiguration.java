// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class BucketProcessConfiguration extends TeaModel {
    @NameInMap("BucketChannelConfig")
    public BucketChannelConfig bucketChannelConfig;

    /**
     * <strong>example:</strong>
     * <p>Img</p>
     */
    @NameInMap("CompliedHost")
    public String compliedHost;

    /**
     * <strong>example:</strong>
     * <p>disabled</p>
     */
    @NameInMap("OssDomainSupportAtProcess")
    public String ossDomainSupportAtProcess;

    /**
     * <strong>example:</strong>
     * <p>disabled</p>
     */
    @NameInMap("SourceFileProtect")
    public String sourceFileProtect;

    /**
     * <strong>example:</strong>
     * <p>gif</p>
     */
    @NameInMap("SourceFileProtectSuffix")
    public String sourceFileProtectSuffix;

    /**
     * <strong>example:</strong>
     * <p>-,|</p>
     */
    @NameInMap("StyleDelimiters")
    public String styleDelimiters;

    public static BucketProcessConfiguration build(java.util.Map<String, ?> map) throws Exception {
        BucketProcessConfiguration self = new BucketProcessConfiguration();
        return TeaModel.build(map, self);
    }

    public BucketProcessConfiguration setBucketChannelConfig(BucketChannelConfig bucketChannelConfig) {
        this.bucketChannelConfig = bucketChannelConfig;
        return this;
    }
    public BucketChannelConfig getBucketChannelConfig() {
        return this.bucketChannelConfig;
    }

    public BucketProcessConfiguration setCompliedHost(String compliedHost) {
        this.compliedHost = compliedHost;
        return this;
    }
    public String getCompliedHost() {
        return this.compliedHost;
    }

    public BucketProcessConfiguration setOssDomainSupportAtProcess(String ossDomainSupportAtProcess) {
        this.ossDomainSupportAtProcess = ossDomainSupportAtProcess;
        return this;
    }
    public String getOssDomainSupportAtProcess() {
        return this.ossDomainSupportAtProcess;
    }

    public BucketProcessConfiguration setSourceFileProtect(String sourceFileProtect) {
        this.sourceFileProtect = sourceFileProtect;
        return this;
    }
    public String getSourceFileProtect() {
        return this.sourceFileProtect;
    }

    public BucketProcessConfiguration setSourceFileProtectSuffix(String sourceFileProtectSuffix) {
        this.sourceFileProtectSuffix = sourceFileProtectSuffix;
        return this;
    }
    public String getSourceFileProtectSuffix() {
        return this.sourceFileProtectSuffix;
    }

    public BucketProcessConfiguration setStyleDelimiters(String styleDelimiters) {
        this.styleDelimiters = styleDelimiters;
        return this;
    }
    public String getStyleDelimiters() {
        return this.styleDelimiters;
    }

}
