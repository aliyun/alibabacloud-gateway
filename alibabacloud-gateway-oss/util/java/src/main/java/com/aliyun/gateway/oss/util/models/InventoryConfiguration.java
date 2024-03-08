// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class InventoryConfiguration extends TeaModel {
    @NameInMap("Destination")
    public InventoryDestination destination;

    @NameInMap("Filter")
    public InventoryFilter filter;

    @NameInMap("Id")
    public String id;

    @NameInMap("IncludedObjectVersions")
    public String includedObjectVersions;

    @NameInMap("IsEnabled")
    public Boolean isEnabled;

    @NameInMap("OptionalFields")
    public InventoryConfigurationOptionalFields optionalFields;

    @NameInMap("Schedule")
    public InventorySchedule schedule;

    public static InventoryConfiguration build(java.util.Map<String, ?> map) throws Exception {
        InventoryConfiguration self = new InventoryConfiguration();
        return TeaModel.build(map, self);
    }

    public InventoryConfiguration setDestination(InventoryDestination destination) {
        this.destination = destination;
        return this;
    }
    public InventoryDestination getDestination() {
        return this.destination;
    }

    public InventoryConfiguration setFilter(InventoryFilter filter) {
        this.filter = filter;
        return this;
    }
    public InventoryFilter getFilter() {
        return this.filter;
    }

    public InventoryConfiguration setId(String id) {
        this.id = id;
        return this;
    }
    public String getId() {
        return this.id;
    }

    public InventoryConfiguration setIncludedObjectVersions(String includedObjectVersions) {
        this.includedObjectVersions = includedObjectVersions;
        return this;
    }
    public String getIncludedObjectVersions() {
        return this.includedObjectVersions;
    }

    public InventoryConfiguration setIsEnabled(Boolean isEnabled) {
        this.isEnabled = isEnabled;
        return this;
    }
    public Boolean getIsEnabled() {
        return this.isEnabled;
    }

    public InventoryConfiguration setOptionalFields(InventoryConfigurationOptionalFields optionalFields) {
        this.optionalFields = optionalFields;
        return this;
    }
    public InventoryConfigurationOptionalFields getOptionalFields() {
        return this.optionalFields;
    }

    public InventoryConfiguration setSchedule(InventorySchedule schedule) {
        this.schedule = schedule;
        return this;
    }
    public InventorySchedule getSchedule() {
        return this.schedule;
    }

    public static class InventoryConfigurationOptionalFields extends TeaModel {
        @NameInMap("Field")
        public java.util.List<String> field;

        public static InventoryConfigurationOptionalFields build(java.util.Map<String, ?> map) throws Exception {
            InventoryConfigurationOptionalFields self = new InventoryConfigurationOptionalFields();
            return TeaModel.build(map, self);
        }

        public InventoryConfigurationOptionalFields setField(java.util.List<String> field) {
            this.field = field;
            return this;
        }
        public java.util.List<String> getField() {
            return this.field;
        }

    }

}
