// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class ListBucketAntiDDosInfoResponseBody extends TeaModel {
    /**
     * <p>The container that stores the information about Anti-DDoS instances.</p>
     */
    @NameInMap("AntiDDOSConfiguration")
    public java.util.List<BucketAntiDDOSInfo> antiDDOSConfiguration;

    /**
     * <p>Indicates whether all Anti-DDoS instances are returned.</p>
     * <br>
     * <p>*   true: All Anti-DDoS instances are not returned.</p>
     * <p>*   false: All Anti-DDoS instances are returned.</p>
     */
    @NameInMap("IsTruncated")
    public Boolean isTruncated;

    /**
     * <p>The Anti-DDoS instances whose names are alphabetically after the value of marker in the request.</p>
     */
    @NameInMap("Marker")
    public String marker;

    public static ListBucketAntiDDosInfoResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListBucketAntiDDosInfoResponseBody self = new ListBucketAntiDDosInfoResponseBody();
        return TeaModel.build(map, self);
    }

    public ListBucketAntiDDosInfoResponseBody setAntiDDOSConfiguration(java.util.List<BucketAntiDDOSInfo> antiDDOSConfiguration) {
        this.antiDDOSConfiguration = antiDDOSConfiguration;
        return this;
    }
    public java.util.List<BucketAntiDDOSInfo> getAntiDDOSConfiguration() {
        return this.antiDDOSConfiguration;
    }

    public ListBucketAntiDDosInfoResponseBody setIsTruncated(Boolean isTruncated) {
        this.isTruncated = isTruncated;
        return this;
    }
    public Boolean getIsTruncated() {
        return this.isTruncated;
    }

    public ListBucketAntiDDosInfoResponseBody setMarker(String marker) {
        this.marker = marker;
        return this;
    }
    public String getMarker() {
        return this.marker;
    }

}
