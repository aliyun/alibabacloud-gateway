// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class ListBucketAntiDDosInfoRequest extends TeaModel {
    /**
     * <p>The name of the Anti-DDoS instance from which the list starts. The Anti-DDoS instances whose names are alphabetically after the value of marker are returned.</p>
     * <br>
     * <p>>  You can set marker to an empty string in the first request. If IsTruncated is returned in the response and the value of IsTruncated is true, you must use the value of Marker in the response as the value of marker in the next request.</p>
     */
    @NameInMap("marker")
    public String marker;

    /**
     * <p>The maximum number of Anti-DDoS instances that can be returned.</p>
     * <br>
     * <p>Valid values: 1 to 100.</p>
     * <br>
     * <p>Default value: 100.</p>
     */
    @NameInMap("max-keys")
    public String maxKeys;

    public static ListBucketAntiDDosInfoRequest build(java.util.Map<String, ?> map) throws Exception {
        ListBucketAntiDDosInfoRequest self = new ListBucketAntiDDosInfoRequest();
        return TeaModel.build(map, self);
    }

    public ListBucketAntiDDosInfoRequest setMarker(String marker) {
        this.marker = marker;
        return this;
    }
    public String getMarker() {
        return this.marker;
    }

    public ListBucketAntiDDosInfoRequest setMaxKeys(String maxKeys) {
        this.maxKeys = maxKeys;
        return this;
    }
    public String getMaxKeys() {
        return this.maxKeys;
    }

}
