// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ReservedCapacityUpdateConfiguration extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>11258999068426240</p>
     */
    @NameInMap("Capacity")
    public Long capacity;

    /**
     * <strong>example:</strong>
     * <p>Init</p>
     */
    @NameInMap("Status")
    public String status;

    /**
     * <strong>example:</strong>
     * <p>1</p>
     */
    @NameInMap("Years")
    public Long years;

    public static ReservedCapacityUpdateConfiguration build(java.util.Map<String, ?> map) throws Exception {
        ReservedCapacityUpdateConfiguration self = new ReservedCapacityUpdateConfiguration();
        return TeaModel.build(map, self);
    }

    public ReservedCapacityUpdateConfiguration setCapacity(Long capacity) {
        this.capacity = capacity;
        return this;
    }
    public Long getCapacity() {
        return this.capacity;
    }

    public ReservedCapacityUpdateConfiguration setStatus(String status) {
        this.status = status;
        return this;
    }
    public String getStatus() {
        return this.status;
    }

    public ReservedCapacityUpdateConfiguration setYears(Long years) {
        this.years = years;
        return this;
    }
    public Long getYears() {
        return this.years;
    }

}
