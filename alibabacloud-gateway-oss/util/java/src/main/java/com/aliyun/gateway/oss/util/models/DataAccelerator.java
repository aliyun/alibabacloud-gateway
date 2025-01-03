// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class DataAccelerator extends TeaModel {
    @NameInMap("BasicInfomation")
    public BasicInfomation basicInfomation;

    /**
     * <strong>example:</strong>
     * <p>test-acc-bucket</p>
     */
    @NameInMap("BucketName")
    public String bucketName;

    /**
     * <strong>example:</strong>
     * <p>test-acc-bucket_data-acc</p>
     */
    @NameInMap("Name")
    public String name;

    public static DataAccelerator build(java.util.Map<String, ?> map) throws Exception {
        DataAccelerator self = new DataAccelerator();
        return TeaModel.build(map, self);
    }

    public DataAccelerator setBasicInfomation(BasicInfomation basicInfomation) {
        this.basicInfomation = basicInfomation;
        return this;
    }
    public BasicInfomation getBasicInfomation() {
        return this.basicInfomation;
    }

    public DataAccelerator setBucketName(String bucketName) {
        this.bucketName = bucketName;
        return this;
    }
    public String getBucketName() {
        return this.bucketName;
    }

    public DataAccelerator setName(String name) {
        this.name = name;
        return this;
    }
    public String getName() {
        return this.name;
    }

    public static class BasicInfomation extends TeaModel {
        @NameInMap("AcceleratePaths")
        public AcceleratePaths acceleratePaths;

        /**
         * <strong>example:</strong>
         * <p>1731394193189</p>
         */
        @NameInMap("CreationDate")
        public String creationDate;

        /**
         * <strong>example:</strong>
         * <p>1731394193189</p>
         */
        @NameInMap("FlowCap")
        public String flowCap;

        /**
         * <strong>example:</strong>
         * <p>normal</p>
         */
        @NameInMap("FlowCapFlag")
        public Boolean flowCapFlag;

        /**
         * <strong>example:</strong>
         * <p>1024000</p>
         */
        @NameInMap("Quota")
        public String quota;

        /**
         * <strong>example:</strong>
         * <p>v2</p>
         */
        @NameInMap("Version")
        public String version;

        public static BasicInfomation build(java.util.Map<String, ?> map) throws Exception {
            BasicInfomation self = new BasicInfomation();
            return TeaModel.build(map, self);
        }

        public BasicInfomation setAcceleratePaths(AcceleratePaths acceleratePaths) {
            this.acceleratePaths = acceleratePaths;
            return this;
        }
        public AcceleratePaths getAcceleratePaths() {
            return this.acceleratePaths;
        }

        public BasicInfomation setCreationDate(String creationDate) {
            this.creationDate = creationDate;
            return this;
        }
        public String getCreationDate() {
            return this.creationDate;
        }

        public BasicInfomation setFlowCap(String flowCap) {
            this.flowCap = flowCap;
            return this;
        }
        public String getFlowCap() {
            return this.flowCap;
        }

        public BasicInfomation setFlowCapFlag(Boolean flowCapFlag) {
            this.flowCapFlag = flowCapFlag;
            return this;
        }
        public Boolean getFlowCapFlag() {
            return this.flowCapFlag;
        }

        public BasicInfomation setQuota(String quota) {
            this.quota = quota;
            return this;
        }
        public String getQuota() {
            return this.quota;
        }

        public BasicInfomation setVersion(String version) {
            this.version = version;
            return this;
        }
        public String getVersion() {
            return this.version;
        }

    }

}
