// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetAccessPointForObjectProcessHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    @NameInMap("x-oss-access-point-for-object-process-name")
    public String xOssAccessPointForObjectProcessName;

    public static GetAccessPointForObjectProcessHeaders build(java.util.Map<String, ?> map) throws Exception {
        GetAccessPointForObjectProcessHeaders self = new GetAccessPointForObjectProcessHeaders();
        return TeaModel.build(map, self);
    }

    public GetAccessPointForObjectProcessHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public GetAccessPointForObjectProcessHeaders setXOssAccessPointForObjectProcessName(String xOssAccessPointForObjectProcessName) {
        this.xOssAccessPointForObjectProcessName = xOssAccessPointForObjectProcessName;
        return this;
    }
    public String getXOssAccessPointForObjectProcessName() {
        return this.xOssAccessPointForObjectProcessName;
    }

}
