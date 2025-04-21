// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class MetaQueryAggregationsResult extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>Size</p>
     */
    @NameInMap("Field")
    public String field;

    @NameInMap("Groups")
    public Groups groups;

    /**
     * <strong>example:</strong>
     * <p>sum</p>
     */
    @NameInMap("Operation")
    public String operation;

    /**
     * <strong>example:</strong>
     * <p>200</p>
     */
    @NameInMap("Value")
    public Double value;

    public static MetaQueryAggregationsResult build(java.util.Map<String, ?> map) throws Exception {
        MetaQueryAggregationsResult self = new MetaQueryAggregationsResult();
        return TeaModel.build(map, self);
    }

    public MetaQueryAggregationsResult setField(String field) {
        this.field = field;
        return this;
    }
    public String getField() {
        return this.field;
    }

    public MetaQueryAggregationsResult setGroups(Groups groups) {
        this.groups = groups;
        return this;
    }
    public Groups getGroups() {
        return this.groups;
    }

    public MetaQueryAggregationsResult setOperation(String operation) {
        this.operation = operation;
        return this;
    }
    public String getOperation() {
        return this.operation;
    }

    public MetaQueryAggregationsResult setValue(Double value) {
        this.value = value;
        return this;
    }
    public Double getValue() {
        return this.value;
    }

    public static class Group extends TeaModel {
        /**
         * <strong>example:</strong>
         * <p>5</p>
         */
        @NameInMap("Count")
        public Long count;

        /**
         * <strong>example:</strong>
         * <p>100</p>
         */
        @NameInMap("Value")
        public String value;

        public static Group build(java.util.Map<String, ?> map) throws Exception {
            Group self = new Group();
            return TeaModel.build(map, self);
        }

        public Group setCount(Long count) {
            this.count = count;
            return this;
        }
        public Long getCount() {
            return this.count;
        }

        public Group setValue(String value) {
            this.value = value;
            return this;
        }
        public String getValue() {
            return this.value;
        }

    }

    public static class Groups extends TeaModel {
        @NameInMap("Group")
        public java.util.List<Group> group;

        public static Groups build(java.util.Map<String, ?> map) throws Exception {
            Groups self = new Groups();
            return TeaModel.build(map, self);
        }

        public Groups setGroup(java.util.List<Group> group) {
            this.group = group;
            return this;
        }
        public java.util.List<Group> getGroup() {
            return this.group;
        }

    }

}
