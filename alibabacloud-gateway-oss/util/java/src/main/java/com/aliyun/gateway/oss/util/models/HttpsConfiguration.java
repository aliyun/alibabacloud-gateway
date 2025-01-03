// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class HttpsConfiguration extends TeaModel {
    @NameInMap("TLS")
    public TLS TLS;

    public static HttpsConfiguration build(java.util.Map<String, ?> map) throws Exception {
        HttpsConfiguration self = new HttpsConfiguration();
        return TeaModel.build(map, self);
    }

    public HttpsConfiguration setTLS(TLS TLS) {
        this.TLS = TLS;
        return this;
    }
    public TLS getTLS() {
        return this.TLS;
    }

    public static class TLS extends TeaModel {
        /**
         * <p>This parameter is required.</p>
         * 
         * <strong>example:</strong>
         * <p>true</p>
         */
        @NameInMap("Enable")
        public Boolean enable;

        @NameInMap("TLSVersion")
        public java.util.List<String> TLSVersion;

        public static TLS build(java.util.Map<String, ?> map) throws Exception {
            TLS self = new TLS();
            return TeaModel.build(map, self);
        }

        public TLS setEnable(Boolean enable) {
            this.enable = enable;
            return this;
        }
        public Boolean getEnable() {
            return this.enable;
        }

        public TLS setTLSVersion(java.util.List<String> TLSVersion) {
            this.TLSVersion = TLSVersion;
            return this;
        }
        public java.util.List<String> getTLSVersion() {
            return this.TLSVersion;
        }

    }

}
