// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class GetUserAntiDDosInfoResponseBody extends TeaModel {
    /**
     * <p>The container that stores the list of Anti-DDoS instances.</p>
     */
    @NameInMap("AntiDDOSListConfiguration")
    public AntiDDOSListConfiguration antiDDOSListConfiguration;

    public static GetUserAntiDDosInfoResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetUserAntiDDosInfoResponseBody self = new GetUserAntiDDosInfoResponseBody();
        return TeaModel.build(map, self);
    }

    public GetUserAntiDDosInfoResponseBody setAntiDDOSListConfiguration(AntiDDOSListConfiguration antiDDOSListConfiguration) {
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
        public java.util.List<UserAntiDDOSInfo> antiDDOSConfiguration;

        public static AntiDDOSListConfiguration build(java.util.Map<String, ?> map) throws Exception {
            AntiDDOSListConfiguration self = new AntiDDOSListConfiguration();
            return TeaModel.build(map, self);
        }

        public AntiDDOSListConfiguration setAntiDDOSConfiguration(java.util.List<UserAntiDDOSInfo> antiDDOSConfiguration) {
            this.antiDDOSConfiguration = antiDDOSConfiguration;
            return this;
        }
        public java.util.List<UserAntiDDOSInfo> getAntiDDOSConfiguration() {
            return this.antiDDOSConfiguration;
        }

    }

}
