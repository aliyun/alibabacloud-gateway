// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CreateDataLakeCachePrefetchJob extends TeaModel {
    @NameInMap("Excludes")
    public java.util.List<String> excludes;

    @NameInMap("Includes")
    public java.util.List<String> includes;

    /**
     * <strong>example:</strong>
     * <p>Jobxxx</p>
     */
    @NameInMap("Tag")
    public String tag;

    public static CreateDataLakeCachePrefetchJob build(java.util.Map<String, ?> map) throws Exception {
        CreateDataLakeCachePrefetchJob self = new CreateDataLakeCachePrefetchJob();
        return TeaModel.build(map, self);
    }

    public CreateDataLakeCachePrefetchJob setExcludes(java.util.List<String> excludes) {
        this.excludes = excludes;
        return this;
    }
    public java.util.List<String> getExcludes() {
        return this.excludes;
    }

    public CreateDataLakeCachePrefetchJob setIncludes(java.util.List<String> includes) {
        this.includes = includes;
        return this;
    }
    public java.util.List<String> getIncludes() {
        return this.includes;
    }

    public CreateDataLakeCachePrefetchJob setTag(String tag) {
        this.tag = tag;
        return this;
    }
    public String getTag() {
        return this.tag;
    }

}
