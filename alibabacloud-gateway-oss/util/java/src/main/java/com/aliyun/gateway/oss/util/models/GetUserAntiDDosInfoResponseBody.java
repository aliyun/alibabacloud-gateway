// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetUserAntiDDosInfoResponseBody extends TeaModel {
    /**
     * <p>The container that stores the information about the specified Anti-DDoS instance.</p>
     */
    @NameInMap("AntiDDOSConfiguration")
    public java.util.List<UserAntiDDOSInfo> antiDDOSConfiguration;

    public static GetUserAntiDDosInfoResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetUserAntiDDosInfoResponseBody self = new GetUserAntiDDosInfoResponseBody();
        return TeaModel.build(map, self);
    }

    public GetUserAntiDDosInfoResponseBody setAntiDDOSConfiguration(java.util.List<UserAntiDDOSInfo> antiDDOSConfiguration) {
        this.antiDDOSConfiguration = antiDDOSConfiguration;
        return this;
    }
    public java.util.List<UserAntiDDOSInfo> getAntiDDOSConfiguration() {
        return this.antiDDOSConfiguration;
    }

}
