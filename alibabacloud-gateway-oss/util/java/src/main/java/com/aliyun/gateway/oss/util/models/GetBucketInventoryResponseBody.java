// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetBucketInventoryResponseBody extends TeaModel {
    /**
     * <p>The container that stores the information about the exported inventories.</p>
     */
    @NameInMap("Destination")
    public InventoryDestination destination;

    /**
     * <p>The prefix that is used to filter objects. Only objects whose names contain the specified prefix are included in the inventories.</p>
     */
    @NameInMap("Filter")
    public InventoryFilter filter;

    /**
     * <p>The name of the inventory. The name is globally unique in the bucket.</p>
     */
    @NameInMap("Id")
    public String id;

    /**
     * <p>Indicates whether the version information about objects is included in inventories.</p>
     * <br>
     * <p>Valid values:</p>
     * <br>
     * <p>*   All: All versions of the objects are exported.</p>
     * <p>*   Current: Only the current versions of the objects are exported.</p>
     */
    @NameInMap("IncludedObjectVersions")
    public String includedObjectVersions;

    /**
     * <p>Indicates whether the inventory feature is enabled. Valid values:</p>
     * <br>
     * <p>*   true: The inventory feature is enabled.</p>
     * <p>*   false: No inventory is generated.</p>
     */
    @NameInMap("IsEnabled")
    public Boolean isEnabled;

    /**
     * <p>The configurations that you specified to be included in an exported inventory.</p>
     */
    @NameInMap("OptionalFields")
    public GetBucketInventoryResponseBodyOptionalFields optionalFields;

    /**
     * <p>The container that stores the information about the frequency at which inventories are exported.</p>
     */
    @NameInMap("Schedule")
    public InventorySchedule schedule;

    public static GetBucketInventoryResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketInventoryResponseBody self = new GetBucketInventoryResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketInventoryResponseBody setDestination(InventoryDestination destination) {
        this.destination = destination;
        return this;
    }
    public InventoryDestination getDestination() {
        return this.destination;
    }

    public GetBucketInventoryResponseBody setFilter(InventoryFilter filter) {
        this.filter = filter;
        return this;
    }
    public InventoryFilter getFilter() {
        return this.filter;
    }

    public GetBucketInventoryResponseBody setId(String id) {
        this.id = id;
        return this;
    }
    public String getId() {
        return this.id;
    }

    public GetBucketInventoryResponseBody setIncludedObjectVersions(String includedObjectVersions) {
        this.includedObjectVersions = includedObjectVersions;
        return this;
    }
    public String getIncludedObjectVersions() {
        return this.includedObjectVersions;
    }

    public GetBucketInventoryResponseBody setIsEnabled(Boolean isEnabled) {
        this.isEnabled = isEnabled;
        return this;
    }
    public Boolean getIsEnabled() {
        return this.isEnabled;
    }

    public GetBucketInventoryResponseBody setOptionalFields(GetBucketInventoryResponseBodyOptionalFields optionalFields) {
        this.optionalFields = optionalFields;
        return this;
    }
    public GetBucketInventoryResponseBodyOptionalFields getOptionalFields() {
        return this.optionalFields;
    }

    public GetBucketInventoryResponseBody setSchedule(InventorySchedule schedule) {
        this.schedule = schedule;
        return this;
    }
    public InventorySchedule getSchedule() {
        return this.schedule;
    }

    public static class GetBucketInventoryResponseBodyOptionalFields extends TeaModel {
        /**
         * <p>The configurations that are included in an exported inventory. Valid values: Size, LastModifiedDate, ETag, StorageClass, IsMultipartUploaded, and EncryptionStatus.</p>
         */
        @NameInMap("Field")
        public java.util.List<String> field;

        public static GetBucketInventoryResponseBodyOptionalFields build(java.util.Map<String, ?> map) throws Exception {
            GetBucketInventoryResponseBodyOptionalFields self = new GetBucketInventoryResponseBodyOptionalFields();
            return TeaModel.build(map, self);
        }

        public GetBucketInventoryResponseBodyOptionalFields setField(java.util.List<String> field) {
            this.field = field;
            return this;
        }
        public java.util.List<String> getField() {
            return this.field;
        }

    }

}
