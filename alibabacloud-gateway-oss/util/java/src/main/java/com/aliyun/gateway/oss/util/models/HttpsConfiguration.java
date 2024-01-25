// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class HttpsConfiguration extends TeaModel {
    @NameInMap("TLS")
    public HttpsConfigurationTLS TLS;

    public static HttpsConfiguration build(java.util.Map<String, ?> map) throws Exception {
        HttpsConfiguration self = new HttpsConfiguration();
        return TeaModel.build(map, self);
    }

    public HttpsConfiguration setTLS(HttpsConfigurationTLS TLS) {
        this.TLS = TLS;
        return this;
    }
    public HttpsConfigurationTLS getTLS() {
        return this.TLS;
    }

    public static class HttpsConfigurationTLS extends TeaModel {
        @NameInMap("Enable")
        public Boolean enable;

        @NameInMap("TLSVersion")
        public java.util.List<String> TLSVersion;

        public static HttpsConfigurationTLS build(java.util.Map<String, ?> map) throws Exception {
            HttpsConfigurationTLS self = new HttpsConfigurationTLS();
            return TeaModel.build(map, self);
        }

        public HttpsConfigurationTLS setEnable(Boolean enable) {
            this.enable = enable;
            return this;
        }
        public Boolean getEnable() {
            return this.enable;
        }

        public HttpsConfigurationTLS setTLSVersion(java.util.List<String> TLSVersion) {
            this.TLSVersion = TLSVersion;
            return this;
        }
        public java.util.List<String> getTLSVersion() {
            return this.TLSVersion;
        }

    }

}
