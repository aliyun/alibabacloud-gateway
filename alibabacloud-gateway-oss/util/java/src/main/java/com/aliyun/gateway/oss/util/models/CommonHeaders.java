// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CommonHeaders extends TeaModel {
    @NameInMap("Header")
    public java.util.List<Header> header;

    public static CommonHeaders build(java.util.Map<String, ?> map) throws Exception {
        CommonHeaders self = new CommonHeaders();
        return TeaModel.build(map, self);
    }

    public CommonHeaders setHeader(java.util.List<Header> header) {
        this.header = header;
        return this;
    }
    public java.util.List<Header> getHeader() {
        return this.header;
    }

    public static class Header extends TeaModel {
        /**
         * <strong>example:</strong>
         * <p>X-Content-Type-Options</p>
         */
        @NameInMap("Key")
        public String key;

        /**
         * <strong>example:</strong>
         * <p>nosniff</p>
         */
        @NameInMap("Value")
        public String value;

        public static Header build(java.util.Map<String, ?> map) throws Exception {
            Header self = new Header();
            return TeaModel.build(map, self);
        }

        public Header setKey(String key) {
            this.key = key;
            return this;
        }
        public String getKey() {
            return this.key;
        }

        public Header setValue(String value) {
            this.value = value;
            return this;
        }
        public String getValue() {
            return this.value;
        }

    }

}
