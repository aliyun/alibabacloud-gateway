// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class HttpsConfiguration extends TeaModel {
    @NameInMap("CipherSuite")
    public CipherSuite cipherSuite;

    @NameInMap("TLS")
    public TLS TLS;

    public static HttpsConfiguration build(java.util.Map<String, ?> map) throws Exception {
        HttpsConfiguration self = new HttpsConfiguration();
        return TeaModel.build(map, self);
    }

    public HttpsConfiguration setCipherSuite(CipherSuite cipherSuite) {
        this.cipherSuite = cipherSuite;
        return this;
    }
    public CipherSuite getCipherSuite() {
        return this.cipherSuite;
    }

    public HttpsConfiguration setTLS(TLS TLS) {
        this.TLS = TLS;
        return this;
    }
    public TLS getTLS() {
        return this.TLS;
    }

    public static class CipherSuite extends TeaModel {
        @NameInMap("CustomCipherSuite")
        public java.util.List<String> customCipherSuite;

        /**
         * <strong>example:</strong>
         * <p>true</p>
         */
        @NameInMap("Enable")
        public Boolean enable;

        /**
         * <strong>example:</strong>
         * <p>false</p>
         */
        @NameInMap("StrongCipherSuite")
        public Boolean strongCipherSuite;

        @NameInMap("TLS13CustomCipherSuite")
        public java.util.List<String> TLS13CustomCipherSuite;

        public static CipherSuite build(java.util.Map<String, ?> map) throws Exception {
            CipherSuite self = new CipherSuite();
            return TeaModel.build(map, self);
        }

        public CipherSuite setCustomCipherSuite(java.util.List<String> customCipherSuite) {
            this.customCipherSuite = customCipherSuite;
            return this;
        }
        public java.util.List<String> getCustomCipherSuite() {
            return this.customCipherSuite;
        }

        public CipherSuite setEnable(Boolean enable) {
            this.enable = enable;
            return this;
        }
        public Boolean getEnable() {
            return this.enable;
        }

        public CipherSuite setStrongCipherSuite(Boolean strongCipherSuite) {
            this.strongCipherSuite = strongCipherSuite;
            return this;
        }
        public Boolean getStrongCipherSuite() {
            return this.strongCipherSuite;
        }

        public CipherSuite setTLS13CustomCipherSuite(java.util.List<String> TLS13CustomCipherSuite) {
            this.TLS13CustomCipherSuite = TLS13CustomCipherSuite;
            return this;
        }
        public java.util.List<String> getTLS13CustomCipherSuite() {
            return this.TLS13CustomCipherSuite;
        }

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
