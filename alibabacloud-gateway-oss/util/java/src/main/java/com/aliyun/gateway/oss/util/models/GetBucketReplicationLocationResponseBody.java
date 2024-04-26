// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketReplicationLocationResponseBody extends TeaModel {
    /**
     * <p>The container that stores the region in which the destination bucket can be located.</p>
     */
    @NameInMap("ReplicationLocation")
    public ReplicationLocation replicationLocation;

    public static GetBucketReplicationLocationResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketReplicationLocationResponseBody self = new GetBucketReplicationLocationResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketReplicationLocationResponseBody setReplicationLocation(ReplicationLocation replicationLocation) {
        this.replicationLocation = replicationLocation;
        return this;
    }
    public ReplicationLocation getReplicationLocation() {
        return this.replicationLocation;
    }

    public static class LocationRTCConstraint extends TeaModel {
        /**
         * <p>The regions where RTC is supported.</p>
         */
        @NameInMap("Location")
        public java.util.List<String> location;

        public static LocationRTCConstraint build(java.util.Map<String, ?> map) throws Exception {
            LocationRTCConstraint self = new LocationRTCConstraint();
            return TeaModel.build(map, self);
        }

        public LocationRTCConstraint setLocation(java.util.List<String> location) {
            this.location = location;
            return this;
        }
        public java.util.List<String> getLocation() {
            return this.location;
        }

    }

    public static class LocationTransferTypeConstraint extends TeaModel {
        /**
         * <p>The container that stores regions in which the destination bucket can be located with the TransferType information.</p>
         */
        @NameInMap("LocationTransferType")
        public java.util.List<LocationTransferType> locationTransferType;

        public static LocationTransferTypeConstraint build(java.util.Map<String, ?> map) throws Exception {
            LocationTransferTypeConstraint self = new LocationTransferTypeConstraint();
            return TeaModel.build(map, self);
        }

        public LocationTransferTypeConstraint setLocationTransferType(java.util.List<LocationTransferType> locationTransferType) {
            this.locationTransferType = locationTransferType;
            return this;
        }
        public java.util.List<LocationTransferType> getLocationTransferType() {
            return this.locationTransferType;
        }

    }

    public static class ReplicationLocation extends TeaModel {
        /**
         * <p>The regions in which the destination bucket can be located.</p>
         */
        @NameInMap("Location")
        public java.util.List<String> location;

        /**
         * <p>The container that stores regions in which the RTC can be enabled.</p>
         */
        @NameInMap("LocationRTCConstraint")
        public LocationRTCConstraint locationRTCConstraint;

        /**
         * <p>The container that stores regions in which the destination bucket can be located with TransferType specified.</p>
         */
        @NameInMap("LocationTransferTypeConstraint")
        public LocationTransferTypeConstraint locationTransferTypeConstraint;

        public static ReplicationLocation build(java.util.Map<String, ?> map) throws Exception {
            ReplicationLocation self = new ReplicationLocation();
            return TeaModel.build(map, self);
        }

        public ReplicationLocation setLocation(java.util.List<String> location) {
            this.location = location;
            return this;
        }
        public java.util.List<String> getLocation() {
            return this.location;
        }

        public ReplicationLocation setLocationRTCConstraint(LocationRTCConstraint locationRTCConstraint) {
            this.locationRTCConstraint = locationRTCConstraint;
            return this;
        }
        public LocationRTCConstraint getLocationRTCConstraint() {
            return this.locationRTCConstraint;
        }

        public ReplicationLocation setLocationTransferTypeConstraint(LocationTransferTypeConstraint locationTransferTypeConstraint) {
            this.locationTransferTypeConstraint = locationTransferTypeConstraint;
            return this;
        }
        public LocationTransferTypeConstraint getLocationTransferTypeConstraint() {
            return this.locationTransferTypeConstraint;
        }

    }

}
