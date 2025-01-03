// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetResourcePoolInfoResp extends TeaModel {
    /**
     * <p>Use the UTC time format: yyyy-MM-ddTHH:mmZ</p>
     * 
     * <strong>example:</strong>
     * <p>2024-07-24T08:42:32.000Z</p>
     */
    @NameInMap("CreateTime")
    public String createTime;

    /**
     * <strong>example:</strong>
     * <p>rp-for-ai</p>
     */
    @NameInMap("Name")
    public String name;

    /**
     * <strong>example:</strong>
     * <p>1234567890987</p>
     */
    @NameInMap("Owner")
    public String owner;

    @NameInMap("QosConfiguration")
    public QoSConfiguration qosConfiguration;

    /**
     * <strong>example:</strong>
     * <p>oss-cn-shanghai</p>
     */
    @NameInMap("Region")
    public String region;

    public static GetResourcePoolInfoResp build(java.util.Map<String, ?> map) throws Exception {
        GetResourcePoolInfoResp self = new GetResourcePoolInfoResp();
        return TeaModel.build(map, self);
    }

    public GetResourcePoolInfoResp setCreateTime(String createTime) {
        this.createTime = createTime;
        return this;
    }
    public String getCreateTime() {
        return this.createTime;
    }

    public GetResourcePoolInfoResp setName(String name) {
        this.name = name;
        return this;
    }
    public String getName() {
        return this.name;
    }

    public GetResourcePoolInfoResp setOwner(String owner) {
        this.owner = owner;
        return this;
    }
    public String getOwner() {
        return this.owner;
    }

    public GetResourcePoolInfoResp setQosConfiguration(QoSConfiguration qosConfiguration) {
        this.qosConfiguration = qosConfiguration;
        return this;
    }
    public QoSConfiguration getQosConfiguration() {
        return this.qosConfiguration;
    }

    public GetResourcePoolInfoResp setRegion(String region) {
        this.region = region;
        return this;
    }
    public String getRegion() {
        return this.region;
    }

}
