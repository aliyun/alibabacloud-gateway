// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetBucketWormResponseBody extends TeaModel {
    /**
     * <p>The time at which the retention policy was created.</p>
     */
    @NameInMap("CreationDate")
    public String creationDate;

    /**
     * <p>The number of days for which objects can be retained.</p>
     */
    @NameInMap("RetentionPeriodInDays")
    public Integer retentionPeriodInDays;

    /**
     * <p>The status of the retention policy. Valid values:</p>
     * <br>
     * <p>*   InProgress: indicates that the retention policy is in the InProgress state. By default, a retention policy is in the InProgress state after it is created. The policy remains in this state for 24 hours.</p>
     * <p>*   Locked: indicates that the retention policy is in the Locked state.</p>
     */
    @NameInMap("State")
    public String state;

    /**
     * <p>The ID of the retention policy.</p>
     * <br>
     * <p>>  If the specified retention policy ID that is used to query the retention policy configurations of the bucket does not exist, OSS returns the 404 error code.</p>
     */
    @NameInMap("WormId")
    public String wormId;

    public static GetBucketWormResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketWormResponseBody self = new GetBucketWormResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketWormResponseBody setCreationDate(String creationDate) {
        this.creationDate = creationDate;
        return this;
    }
    public String getCreationDate() {
        return this.creationDate;
    }

    public GetBucketWormResponseBody setRetentionPeriodInDays(Integer retentionPeriodInDays) {
        this.retentionPeriodInDays = retentionPeriodInDays;
        return this;
    }
    public Integer getRetentionPeriodInDays() {
        return this.retentionPeriodInDays;
    }

    public GetBucketWormResponseBody setState(String state) {
        this.state = state;
        return this;
    }
    public String getState() {
        return this.state;
    }

    public GetBucketWormResponseBody setWormId(String wormId) {
        this.wormId = wormId;
        return this;
    }
    public String getWormId() {
        return this.wormId;
    }

}
