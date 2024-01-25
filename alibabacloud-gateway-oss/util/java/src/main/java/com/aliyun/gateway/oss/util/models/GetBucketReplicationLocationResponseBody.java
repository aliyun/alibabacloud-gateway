// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetBucketReplicationLocationResponseBody extends TeaModel {
    /**
     * <p>The regions in which available destination buckets reside.</p>
     * <br>
     * <p>>  If available destination buckets reside in multiple regions, multiple values of Location are returned in the response. If no destination bucket is available, the value of Location is null.</p>
     */
    @NameInMap("Location")
    public java.util.List<String> location;

    /**
     * <p>The container that stores the regions in which available destination buckets reside with TransferType specified.</p>
     */
    @NameInMap("LocationTransferTypeConstraint")
    public GetBucketReplicationLocationResponseBodyLocationTransferTypeConstraint locationTransferTypeConstraint;

    public static GetBucketReplicationLocationResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketReplicationLocationResponseBody self = new GetBucketReplicationLocationResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketReplicationLocationResponseBody setLocation(java.util.List<String> location) {
        this.location = location;
        return this;
    }
    public java.util.List<String> getLocation() {
        return this.location;
    }

    public GetBucketReplicationLocationResponseBody setLocationTransferTypeConstraint(GetBucketReplicationLocationResponseBodyLocationTransferTypeConstraint locationTransferTypeConstraint) {
        this.locationTransferTypeConstraint = locationTransferTypeConstraint;
        return this;
    }
    public GetBucketReplicationLocationResponseBodyLocationTransferTypeConstraint getLocationTransferTypeConstraint() {
        return this.locationTransferTypeConstraint;
    }

    public static class GetBucketReplicationLocationResponseBodyLocationTransferTypeConstraint extends TeaModel {
        /**
         * <p>The container that stores the regions in which available destination buckets reside with TransferType specified.</p>
         */
        @NameInMap("LocationTransferType")
        public java.util.List<LocationTransferType> locationTransferType;

        public static GetBucketReplicationLocationResponseBodyLocationTransferTypeConstraint build(java.util.Map<String, ?> map) throws Exception {
            GetBucketReplicationLocationResponseBodyLocationTransferTypeConstraint self = new GetBucketReplicationLocationResponseBodyLocationTransferTypeConstraint();
            return TeaModel.build(map, self);
        }

        public GetBucketReplicationLocationResponseBodyLocationTransferTypeConstraint setLocationTransferType(java.util.List<LocationTransferType> locationTransferType) {
            this.locationTransferType = locationTransferType;
            return this;
        }
        public java.util.List<LocationTransferType> getLocationTransferType() {
            return this.locationTransferType;
        }

    }

}
