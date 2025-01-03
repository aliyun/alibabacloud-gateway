// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketProcessConfiguration extends TeaModel {
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
     * <p>2024-10-18T09:51:54.000Z</p>
     */
    @NameInMap("LastModified")
    public String lastModified;

    /**
     * <strong>example:</strong>
     * <p>Disabled</p>
     */
    @NameInMap("SourceFileProtect")
    public String sourceFileProtect;

    /**
     * <strong>example:</strong>
     * <p>.jpg</p>
     */
    @NameInMap("SourceFileProtectSuffix")
    public String sourceFileProtectSuffix;

    /**
     * <strong>example:</strong>
     * <p>-,_,/,!</p>
     */
    @NameInMap("StyleDelimiters")
    public String styleDelimiters;

    /**
     * <strong>example:</strong>
     * <p>2</p>
     */
    @NameInMap("Version")
    public Integer version;

    public static GetBucketProcessConfiguration build(java.util.Map<String, ?> map) throws Exception {
        GetBucketProcessConfiguration self = new GetBucketProcessConfiguration();
        return TeaModel.build(map, self);
    }

    public GetBucketProcessConfiguration setBucketChannelConfig(BucketChannelConfig bucketChannelConfig) {
        this.bucketChannelConfig = bucketChannelConfig;
        return this;
    }
    public BucketChannelConfig getBucketChannelConfig() {
        return this.bucketChannelConfig;
    }

    public GetBucketProcessConfiguration setCompliedHost(String compliedHost) {
        this.compliedHost = compliedHost;
        return this;
    }
    public String getCompliedHost() {
        return this.compliedHost;
    }

    public GetBucketProcessConfiguration setLastModified(String lastModified) {
        this.lastModified = lastModified;
        return this;
    }
    public String getLastModified() {
        return this.lastModified;
    }

    public GetBucketProcessConfiguration setSourceFileProtect(String sourceFileProtect) {
        this.sourceFileProtect = sourceFileProtect;
        return this;
    }
    public String getSourceFileProtect() {
        return this.sourceFileProtect;
    }

    public GetBucketProcessConfiguration setSourceFileProtectSuffix(String sourceFileProtectSuffix) {
        this.sourceFileProtectSuffix = sourceFileProtectSuffix;
        return this;
    }
    public String getSourceFileProtectSuffix() {
        return this.sourceFileProtectSuffix;
    }

    public GetBucketProcessConfiguration setStyleDelimiters(String styleDelimiters) {
        this.styleDelimiters = styleDelimiters;
        return this;
    }
    public String getStyleDelimiters() {
        return this.styleDelimiters;
    }

    public GetBucketProcessConfiguration setVersion(Integer version) {
        this.version = version;
        return this;
    }
    public Integer getVersion() {
        return this.version;
    }

}
