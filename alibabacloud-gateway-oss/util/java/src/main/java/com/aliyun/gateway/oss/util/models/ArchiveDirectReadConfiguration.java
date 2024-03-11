// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class ArchiveDirectReadConfiguration extends TeaModel {
    @NameInMap("Enabled")
    public Boolean enabled;

    public static ArchiveDirectReadConfiguration build(java.util.Map<String, ?> map) throws Exception {
        ArchiveDirectReadConfiguration self = new ArchiveDirectReadConfiguration();
        return TeaModel.build(map, self);
    }

    public ArchiveDirectReadConfiguration setEnabled(Boolean enabled) {
        this.enabled = enabled;
        return this;
    }
    public Boolean getEnabled() {
        return this.enabled;
    }

}
