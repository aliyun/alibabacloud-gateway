// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ReservedCapacityCreateConfiguration extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>11258999068426240</p>
     */
    @NameInMap("Capacity")
    public Long capacity;

    /**
     * <strong>example:</strong>
     * <p>LRS</p>
     */
    @NameInMap("DataRedundancyType")
    public String dataRedundancyType;

    /**
     * <strong>example:</strong>
     * <p>test-rc-01</p>
     */
    @NameInMap("Name")
    public String name;

    /**
     * <strong>example:</strong>
     * <p>oss-cn-hangzhou</p>
     */
    @NameInMap("Region")
    public String region;

    /**
     * <strong>example:</strong>
     * <p>1</p>
     */
    @NameInMap("Years")
    public Long years;

    public static ReservedCapacityCreateConfiguration build(java.util.Map<String, ?> map) throws Exception {
        ReservedCapacityCreateConfiguration self = new ReservedCapacityCreateConfiguration();
        return TeaModel.build(map, self);
    }

    public ReservedCapacityCreateConfiguration setCapacity(Long capacity) {
        this.capacity = capacity;
        return this;
    }
    public Long getCapacity() {
        return this.capacity;
    }

    public ReservedCapacityCreateConfiguration setDataRedundancyType(String dataRedundancyType) {
        this.dataRedundancyType = dataRedundancyType;
        return this;
    }
    public String getDataRedundancyType() {
        return this.dataRedundancyType;
    }

    public ReservedCapacityCreateConfiguration setName(String name) {
        this.name = name;
        return this;
    }
    public String getName() {
        return this.name;
    }

    public ReservedCapacityCreateConfiguration setRegion(String region) {
        this.region = region;
        return this;
    }
    public String getRegion() {
        return this.region;
    }

    public ReservedCapacityCreateConfiguration setYears(Long years) {
        this.years = years;
        return this;
    }
    public Long getYears() {
        return this.years;
    }

}
