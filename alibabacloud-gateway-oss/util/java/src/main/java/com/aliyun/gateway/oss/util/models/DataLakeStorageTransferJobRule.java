// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class DataLakeStorageTransferJobRule extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>AliyunOSSRole</p>
     */
    @NameInMap("ExecutorRoleId")
    public String executorRoleId;

    /**
     * <strong>example:</strong>
     * <p>log/</p>
     */
    @NameInMap("LogBaseDir")
    public String logBaseDir;

    /**
     * <strong>example:</strong>
     * <p>false</p>
     */
    @NameInMap("NeedVerify")
    public Boolean needVerify;

    @NameInMap("PrefixFilter")
    public PrefixFilter prefixFilter;

    /**
     * <strong>example:</strong>
     * <p>tag</p>
     */
    @NameInMap("Tag")
    public String tag;

    public static DataLakeStorageTransferJobRule build(java.util.Map<String, ?> map) throws Exception {
        DataLakeStorageTransferJobRule self = new DataLakeStorageTransferJobRule();
        return TeaModel.build(map, self);
    }

    public DataLakeStorageTransferJobRule setExecutorRoleId(String executorRoleId) {
        this.executorRoleId = executorRoleId;
        return this;
    }
    public String getExecutorRoleId() {
        return this.executorRoleId;
    }

    public DataLakeStorageTransferJobRule setLogBaseDir(String logBaseDir) {
        this.logBaseDir = logBaseDir;
        return this;
    }
    public String getLogBaseDir() {
        return this.logBaseDir;
    }

    public DataLakeStorageTransferJobRule setNeedVerify(Boolean needVerify) {
        this.needVerify = needVerify;
        return this;
    }
    public Boolean getNeedVerify() {
        return this.needVerify;
    }

    public DataLakeStorageTransferJobRule setPrefixFilter(PrefixFilter prefixFilter) {
        this.prefixFilter = prefixFilter;
        return this;
    }
    public PrefixFilter getPrefixFilter() {
        return this.prefixFilter;
    }

    public DataLakeStorageTransferJobRule setTag(String tag) {
        this.tag = tag;
        return this;
    }
    public String getTag() {
        return this.tag;
    }

    public static class Includes extends TeaModel {
        @NameInMap("Include")
        public java.util.List<String> include;

        public static Includes build(java.util.Map<String, ?> map) throws Exception {
            Includes self = new Includes();
            return TeaModel.build(map, self);
        }

        public Includes setInclude(java.util.List<String> include) {
            this.include = include;
            return this;
        }
        public java.util.List<String> getInclude() {
            return this.include;
        }

    }

    public static class PrefixFilter extends TeaModel {
        @NameInMap("Includes")
        public Includes includes;

        public static PrefixFilter build(java.util.Map<String, ?> map) throws Exception {
            PrefixFilter self = new PrefixFilter();
            return TeaModel.build(map, self);
        }

        public PrefixFilter setIncludes(Includes includes) {
            this.includes = includes;
            return this;
        }
        public Includes getIncludes() {
            return this.includes;
        }

    }

}
