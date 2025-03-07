// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CORSRule extends TeaModel {
    @NameInMap("AllowedHeader")
    public java.util.List<String> allowedHeader;

    @NameInMap("AllowedMethod")
    public java.util.List<String> allowedMethod;

    @NameInMap("AllowedOrigin")
    public java.util.List<String> allowedOrigin;

    @NameInMap("ExposeHeader")
    public java.util.List<String> exposeHeader;

    /**
     * <strong>example:</strong>
     * <p>300</p>
     */
    @NameInMap("MaxAgeSeconds")
    public Long maxAgeSeconds;

    public static CORSRule build(java.util.Map<String, ?> map) throws Exception {
        CORSRule self = new CORSRule();
        return TeaModel.build(map, self);
    }

    public CORSRule setAllowedHeader(java.util.List<String> allowedHeader) {
        this.allowedHeader = allowedHeader;
        return this;
    }
    public java.util.List<String> getAllowedHeader() {
        return this.allowedHeader;
    }

    public CORSRule setAllowedMethod(java.util.List<String> allowedMethod) {
        this.allowedMethod = allowedMethod;
        return this;
    }
    public java.util.List<String> getAllowedMethod() {
        return this.allowedMethod;
    }

    public CORSRule setAllowedOrigin(java.util.List<String> allowedOrigin) {
        this.allowedOrigin = allowedOrigin;
        return this;
    }
    public java.util.List<String> getAllowedOrigin() {
        return this.allowedOrigin;
    }

    public CORSRule setExposeHeader(java.util.List<String> exposeHeader) {
        this.exposeHeader = exposeHeader;
        return this;
    }
    public java.util.List<String> getExposeHeader() {
        return this.exposeHeader;
    }

    public CORSRule setMaxAgeSeconds(Long maxAgeSeconds) {
        this.maxAgeSeconds = maxAgeSeconds;
        return this;
    }
    public Long getMaxAgeSeconds() {
        return this.maxAgeSeconds;
    }

}
