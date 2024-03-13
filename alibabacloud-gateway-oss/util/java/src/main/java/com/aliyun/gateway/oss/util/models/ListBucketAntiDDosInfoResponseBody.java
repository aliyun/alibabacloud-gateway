// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListBucketAntiDDosInfoResponseBody extends TeaModel {
    /**
     * <p>The container that stores the protection list of an Anti-DDoS instance of a bucket.</p>
     */
    @NameInMap("AntiDDOSListConfiguration")
    public AntiDDOSListConfiguration antiDDOSListConfiguration;

    public static ListBucketAntiDDosInfoResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListBucketAntiDDosInfoResponseBody self = new ListBucketAntiDDosInfoResponseBody();
        return TeaModel.build(map, self);
    }

    public ListBucketAntiDDosInfoResponseBody setAntiDDOSListConfiguration(AntiDDOSListConfiguration antiDDOSListConfiguration) {
        this.antiDDOSListConfiguration = antiDDOSListConfiguration;
        return this;
    }
    public AntiDDOSListConfiguration getAntiDDOSListConfiguration() {
        return this.antiDDOSListConfiguration;
    }

    public static class AntiDDOSListConfiguration extends TeaModel {
        /**
         * <p>The container that stores information about the Anti-DDoS instance.</p>
         */
        @NameInMap("AntiDDOSConfiguration")
        public java.util.List<BucketAntiDDOSInfo> antiDDOSConfiguration;

        /**
         * <p>Indicates whether all Anti-DDoS instances are returned.</p>
         * <br>
         * <p>- true: All Anti-DDoS instances are returned.</p>
         * <br>
         * <p>- false: Not all Anti-DDoS instances are returned.</p>
         */
        @NameInMap("IsTruncated")
        public Boolean isTruncated;

        /**
         * <p>The Anti-DDoS instances whose names are alphabetically after the specified marker.</p>
         */
        @NameInMap("Marker")
        public String marker;

        public static AntiDDOSListConfiguration build(java.util.Map<String, ?> map) throws Exception {
            AntiDDOSListConfiguration self = new AntiDDOSListConfiguration();
            return TeaModel.build(map, self);
        }

        public AntiDDOSListConfiguration setAntiDDOSConfiguration(java.util.List<BucketAntiDDOSInfo> antiDDOSConfiguration) {
            this.antiDDOSConfiguration = antiDDOSConfiguration;
            return this;
        }
        public java.util.List<BucketAntiDDOSInfo> getAntiDDOSConfiguration() {
            return this.antiDDOSConfiguration;
        }

        public AntiDDOSListConfiguration setIsTruncated(Boolean isTruncated) {
            this.isTruncated = isTruncated;
            return this;
        }
        public Boolean getIsTruncated() {
            return this.isTruncated;
        }

        public AntiDDOSListConfiguration setMarker(String marker) {
            this.marker = marker;
            return this;
        }
        public String getMarker() {
            return this.marker;
        }

    }

}
