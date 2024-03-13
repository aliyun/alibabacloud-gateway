// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListBucketInventoryResponseBody extends TeaModel {
    /**
     * <p>The container that stores inventory configuration list.</p>
     */
    @NameInMap("ListInventoryConfigurationsResult")
    public ListInventoryConfigurationsResult listInventoryConfigurationsResult;

    public static ListBucketInventoryResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListBucketInventoryResponseBody self = new ListBucketInventoryResponseBody();
        return TeaModel.build(map, self);
    }

    public ListBucketInventoryResponseBody setListInventoryConfigurationsResult(ListInventoryConfigurationsResult listInventoryConfigurationsResult) {
        this.listInventoryConfigurationsResult = listInventoryConfigurationsResult;
        return this;
    }
    public ListInventoryConfigurationsResult getListInventoryConfigurationsResult() {
        return this.listInventoryConfigurationsResult;
    }

    public static class ListInventoryConfigurationsResult extends TeaModel {
        /**
         * <p>The container that stores inventory configurations.</p>
         */
        @NameInMap("InventoryConfiguration")
        public java.util.List<InventoryConfiguration> inventoryConfigurations;

        /**
         * <p>Specifies whether to list all inventory tasks configured for the bucket.</p>
         * <p>Valid values: true and false</p>
         * <p>- The value of false indicates that all inventory tasks configured for the bucket are listed.</p>
         * <p>- The value of true indicates that not all inventory tasks configured for the bucket are listed. To list the next page of inventory configurations, set the continuation-token parameter in the next request to the value of the NextContinuationToken header in the response to the current request.</p>
         */
        @NameInMap("IsTruncated")
        public Boolean isTruncated;

        /**
         * <p>If the value of IsTruncated in the response is true and value of this header is not null, set the continuation-token parameter in the next request to the value of this header.</p>
         */
        @NameInMap("NextContinuationToken")
        public String nextContinuationToken;

        public static ListInventoryConfigurationsResult build(java.util.Map<String, ?> map) throws Exception {
            ListInventoryConfigurationsResult self = new ListInventoryConfigurationsResult();
            return TeaModel.build(map, self);
        }

        public ListInventoryConfigurationsResult setInventoryConfigurations(java.util.List<InventoryConfiguration> inventoryConfigurations) {
            this.inventoryConfigurations = inventoryConfigurations;
            return this;
        }
        public java.util.List<InventoryConfiguration> getInventoryConfigurations() {
            return this.inventoryConfigurations;
        }

        public ListInventoryConfigurationsResult setIsTruncated(Boolean isTruncated) {
            this.isTruncated = isTruncated;
            return this;
        }
        public Boolean getIsTruncated() {
            return this.isTruncated;
        }

        public ListInventoryConfigurationsResult setNextContinuationToken(String nextContinuationToken) {
            this.nextContinuationToken = nextContinuationToken;
            return this;
        }
        public String getNextContinuationToken() {
            return this.nextContinuationToken;
        }

    }

}
