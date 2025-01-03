// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CreateDataLakeStorageTransferJob extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>AliyunOSSRole</p>
     */
    @NameInMap("ExecutorRoleId")
    public String executorRoleId;

    @NameInMap("Includes")
    public java.util.List<String> includes;

    /**
     * <strong>example:</strong>
     * <p>log</p>
     */
    @NameInMap("LogBaseDir")
    public String logBaseDir;

    /**
     * <strong>example:</strong>
     * <p>false</p>
     */
    @NameInMap("NeedVerify")
    public Boolean needVerify;

    /**
     * <strong>example:</strong>
     * <p>tag</p>
     */
    @NameInMap("Tag")
    public String tag;

    public static CreateDataLakeStorageTransferJob build(java.util.Map<String, ?> map) throws Exception {
        CreateDataLakeStorageTransferJob self = new CreateDataLakeStorageTransferJob();
        return TeaModel.build(map, self);
    }

    public CreateDataLakeStorageTransferJob setExecutorRoleId(String executorRoleId) {
        this.executorRoleId = executorRoleId;
        return this;
    }
    public String getExecutorRoleId() {
        return this.executorRoleId;
    }

    public CreateDataLakeStorageTransferJob setIncludes(java.util.List<String> includes) {
        this.includes = includes;
        return this;
    }
    public java.util.List<String> getIncludes() {
        return this.includes;
    }

    public CreateDataLakeStorageTransferJob setLogBaseDir(String logBaseDir) {
        this.logBaseDir = logBaseDir;
        return this;
    }
    public String getLogBaseDir() {
        return this.logBaseDir;
    }

    public CreateDataLakeStorageTransferJob setNeedVerify(Boolean needVerify) {
        this.needVerify = needVerify;
        return this;
    }
    public Boolean getNeedVerify() {
        return this.needVerify;
    }

    public CreateDataLakeStorageTransferJob setTag(String tag) {
        this.tag = tag;
        return this;
    }
    public String getTag() {
        return this.tag;
    }

}
