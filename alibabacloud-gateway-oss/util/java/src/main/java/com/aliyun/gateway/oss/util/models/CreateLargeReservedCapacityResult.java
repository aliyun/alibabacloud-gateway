// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CreateLargeReservedCapacityResult extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>c99797e7-510d-41e3-ac82-e851c17e1f5c</p>
     */
    @NameInMap("ID")
    public String ID;

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

    public static CreateLargeReservedCapacityResult build(java.util.Map<String, ?> map) throws Exception {
        CreateLargeReservedCapacityResult self = new CreateLargeReservedCapacityResult();
        return TeaModel.build(map, self);
    }

    public CreateLargeReservedCapacityResult setID(String ID) {
        this.ID = ID;
        return this;
    }
    public String getID() {
        return this.ID;
    }

    public CreateLargeReservedCapacityResult setName(String name) {
        this.name = name;
        return this;
    }
    public String getName() {
        return this.name;
    }

    public CreateLargeReservedCapacityResult setOwner(Owner owner) {
        this.owner = owner;
        return this;
    }
    public Owner getOwner() {
        return this.owner;
    }

    public CreateLargeReservedCapacityResult setRegion(String region) {
        this.region = region;
        return this;
    }
    public String getRegion() {
        return this.region;
    }

}
