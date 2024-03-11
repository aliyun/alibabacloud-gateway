// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class OptionObjectHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>The custom headers to be sent in the actual cross-origin request. You can configure multiple custom headers in a cross-origin request. Custom headers are separated by commas (,). By default, this header is left empty.</p>
     */
    @NameInMap("Access-Control-Request-Headers")
    public String accessControlRequestHeaders;

    /**
     * <p>The method to be used in the actual cross-origin request. You can specify only one Access-Control-Request-Method header in a cross-origin request. By default, this header is left empty.</p>
     */
    @NameInMap("Access-Control-Request-Method")
    public String accessControlRequestMethod;

    /**
     * <p>The origin of the request. It is used to identify a cross-origin request. You can specify only one Origin header in a cross-origin request. By default, this header is left empty.</p>
     */
    @NameInMap("Origin")
    public String origin;

    public static OptionObjectHeaders build(java.util.Map<String, ?> map) throws Exception {
        OptionObjectHeaders self = new OptionObjectHeaders();
        return TeaModel.build(map, self);
    }

    public OptionObjectHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public OptionObjectHeaders setAccessControlRequestHeaders(String accessControlRequestHeaders) {
        this.accessControlRequestHeaders = accessControlRequestHeaders;
        return this;
    }
    public String getAccessControlRequestHeaders() {
        return this.accessControlRequestHeaders;
    }

    public OptionObjectHeaders setAccessControlRequestMethod(String accessControlRequestMethod) {
        this.accessControlRequestMethod = accessControlRequestMethod;
        return this;
    }
    public String getAccessControlRequestMethod() {
        return this.accessControlRequestMethod;
    }

    public OptionObjectHeaders setOrigin(String origin) {
        this.origin = origin;
        return this;
    }
    public String getOrigin() {
        return this.origin;
    }

}
