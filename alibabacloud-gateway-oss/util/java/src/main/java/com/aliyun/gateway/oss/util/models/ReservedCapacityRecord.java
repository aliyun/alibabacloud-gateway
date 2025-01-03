// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ReservedCapacityRecord extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>11258999068426240</p>
     */
    @NameInMap("Capacity")
    public Long capacity;

    /**
     * <strong>example:</strong>
     * <p>1733822455</p>
     */
    @NameInMap("CreateTime")
    public Long createTime;

    /**
     * <strong>example:</strong>
     * <p>LRS</p>
     */
    @NameInMap("DataRedundancyType")
    public String dataRedundancyType;

    /**
     * <strong>example:</strong>
     * <p>0</p>
     */
    @NameInMap("DueTime")
    public Long dueTime;

    /**
     * <strong>example:</strong>
     * <p>0</p>
     */
    @NameInMap("ExpansionTime")
    public Long expansionTime;

    /**
     * <strong>example:</strong>
     * <p>0</p>
     */
    @NameInMap("FirstTimeEnabled")
    public Long firstTimeEnabled;

    /**
     * <strong>example:</strong>
     * <p>c99797e7-510d-41e3-ac82-e851c17e1f5c</p>
     */
    @NameInMap("ID")
    public String ID;

    /**
     * <strong>example:</strong>
     * <p>0</p>
     */
    @NameInMap("LastExpansionCapacity")
    public Long lastExpansionCapacity;

    /**
     * <strong>example:</strong>
     * <p>1733822455</p>
     */
    @NameInMap("LastModifyTime")
    public Long lastModifyTime;

    /**
     * <strong>example:</strong>
     * <p>test-rc-01</p>
     */
    @NameInMap("Name")
    public String name;

    @NameInMap("Owner")
    public Owner owner;

    /**
     * <strong>example:</strong>
     * <p>oss-cn-hangzhou</p>
     */
    @NameInMap("Region")
    public String region;

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
    @NameInMap("Version")
    public Long version;

    /**
     * <strong>example:</strong>
     * <p>1</p>
     */
    @NameInMap("Years")
    public Long years;

    public static ReservedCapacityRecord build(java.util.Map<String, ?> map) throws Exception {
        ReservedCapacityRecord self = new ReservedCapacityRecord();
        return TeaModel.build(map, self);
    }

    public ReservedCapacityRecord setCapacity(Long capacity) {
        this.capacity = capacity;
        return this;
    }
    public Long getCapacity() {
        return this.capacity;
    }

    public ReservedCapacityRecord setCreateTime(Long createTime) {
        this.createTime = createTime;
        return this;
    }
    public Long getCreateTime() {
        return this.createTime;
    }

    public ReservedCapacityRecord setDataRedundancyType(String dataRedundancyType) {
        this.dataRedundancyType = dataRedundancyType;
        return this;
    }
    public String getDataRedundancyType() {
        return this.dataRedundancyType;
    }

    public ReservedCapacityRecord setDueTime(Long dueTime) {
        this.dueTime = dueTime;
        return this;
    }
    public Long getDueTime() {
        return this.dueTime;
    }

    public ReservedCapacityRecord setExpansionTime(Long expansionTime) {
        this.expansionTime = expansionTime;
        return this;
    }
    public Long getExpansionTime() {
        return this.expansionTime;
    }

    public ReservedCapacityRecord setFirstTimeEnabled(Long firstTimeEnabled) {
        this.firstTimeEnabled = firstTimeEnabled;
        return this;
    }
    public Long getFirstTimeEnabled() {
        return this.firstTimeEnabled;
    }

    public ReservedCapacityRecord setID(String ID) {
        this.ID = ID;
        return this;
    }
    public String getID() {
        return this.ID;
    }

    public ReservedCapacityRecord setLastExpansionCapacity(Long lastExpansionCapacity) {
        this.lastExpansionCapacity = lastExpansionCapacity;
        return this;
    }
    public Long getLastExpansionCapacity() {
        return this.lastExpansionCapacity;
    }

    public ReservedCapacityRecord setLastModifyTime(Long lastModifyTime) {
        this.lastModifyTime = lastModifyTime;
        return this;
    }
    public Long getLastModifyTime() {
        return this.lastModifyTime;
    }

    public ReservedCapacityRecord setName(String name) {
        this.name = name;
        return this;
    }
    public String getName() {
        return this.name;
    }

    public ReservedCapacityRecord setOwner(Owner owner) {
        this.owner = owner;
        return this;
    }
    public Owner getOwner() {
        return this.owner;
    }

    public ReservedCapacityRecord setRegion(String region) {
        this.region = region;
        return this;
    }
    public String getRegion() {
        return this.region;
    }

    public ReservedCapacityRecord setStatus(String status) {
        this.status = status;
        return this;
    }
    public String getStatus() {
        return this.status;
    }

    public ReservedCapacityRecord setVersion(Long version) {
        this.version = version;
        return this;
    }
    public Long getVersion() {
        return this.version;
    }

    public ReservedCapacityRecord setYears(Long years) {
        this.years = years;
        return this;
    }
    public Long getYears() {
        return this.years;
    }

}
