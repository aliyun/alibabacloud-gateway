// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class ListBucketInventoryResponseBody extends TeaModel {
    /**
     * <p>The container that stores inventory configurations.</p>
     */
    @NameInMap("InventoryConfiguration")
    public java.util.List<InventoryConfiguration> inventoryConfiguration;

    /**
     * <p>Indicates whether all inventories of the bucket are listed. Valid values:</p>
     * <br>
     * <p>*   false: All inventories of the bucket are listed.</p>
     * <p>*   true: All inventories of the bucket are not listed. You can specify the value of NextContinuationToken as the value of continuation-token in the next request to retrieve a new page of results.</p>
     */
    @NameInMap("IsTruncated")
    public Boolean isTruncated;

    /**
     * <p>If the value of IsTruncated is true and the value of this parameter is not empty, set continuation-token in the next request to the value of this parameter.</p>
     */
    @NameInMap("NextContinuationToken")
    public String nextContinuationToken;

    public static ListBucketInventoryResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListBucketInventoryResponseBody self = new ListBucketInventoryResponseBody();
        return TeaModel.build(map, self);
    }

    public ListBucketInventoryResponseBody setInventoryConfiguration(java.util.List<InventoryConfiguration> inventoryConfiguration) {
        this.inventoryConfiguration = inventoryConfiguration;
        return this;
    }
    public java.util.List<InventoryConfiguration> getInventoryConfiguration() {
        return this.inventoryConfiguration;
    }

    public ListBucketInventoryResponseBody setIsTruncated(Boolean isTruncated) {
        this.isTruncated = isTruncated;
        return this;
    }
    public Boolean getIsTruncated() {
        return this.isTruncated;
    }

    public ListBucketInventoryResponseBody setNextContinuationToken(String nextContinuationToken) {
        this.nextContinuationToken = nextContinuationToken;
        return this;
    }
    public String getNextContinuationToken() {
        return this.nextContinuationToken;
    }

}
