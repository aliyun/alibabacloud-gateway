// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.hcs_mgw_models;

import com.aliyun.tea.*;

public class FileTypeFilters extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>fasle</p>
     */
    @NameInMap("ExcludeDir")
    public Boolean excludeDir;

    /**
     * <strong>example:</strong>
     * <p>fasle</p>
     */
    @NameInMap("ExcludeSymlink")
    public Boolean excludeSymlink;

    public static FileTypeFilters build(java.util.Map<String, ?> map) throws Exception {
        FileTypeFilters self = new FileTypeFilters();
        return TeaModel.build(map, self);
    }

    public FileTypeFilters setExcludeDir(Boolean excludeDir) {
        this.excludeDir = excludeDir;
        return this;
    }
    public Boolean getExcludeDir() {
        return this.excludeDir;
    }

    public FileTypeFilters setExcludeSymlink(Boolean excludeSymlink) {
        this.excludeSymlink = excludeSymlink;
        return this;
    }
    public Boolean getExcludeSymlink() {
        return this.excludeSymlink;
    }

}
