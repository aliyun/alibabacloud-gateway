// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class WriteGetObjectResponseHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    @NameInMap("Content-Length")
    public Long contentLength;

    @NameInMap("x-oss-fwd-header-Accept-Ranges")
    public String xOssFwdHeaderAcceptRanges;

    @NameInMap("x-oss-fwd-header-Cache-Control")
    public String xOssFwdHeaderCacheControl;

    @NameInMap("x-oss-fwd-header-Content-Disposition")
    public String xOssFwdHeaderContentDisposition;

    @NameInMap("x-oss-fwd-header-Content-Encoding")
    public String xOssFwdHeaderContentEncoding;

    @NameInMap("x-oss-fwd-header-Content-Language")
    public String xOssFwdHeaderContentLanguage;

    @NameInMap("x-oss-fwd-header-Content-Range")
    public String xOssFwdHeaderContentRange;

    @NameInMap("x-oss-fwd-header-Content-Type")
    public String xOssFwdHeaderContentType;

    @NameInMap("x-oss-fwd-header-ETag")
    public String xOssFwdHeaderETag;

    @NameInMap("x-oss-fwd-header-Expires")
    public String xOssFwdHeaderExpires;

    @NameInMap("x-oss-fwd-header-Last-Modified")
    public String xOssFwdHeaderLastModified;

    @NameInMap("x-oss-fwd-status")
    public String xOssFwdStatus;

    @NameInMap("x-oss-request-route")
    public String xOssRequestRoute;

    @NameInMap("x-oss-request-token")
    public String xOssRequestToken;

    public static WriteGetObjectResponseHeaders build(java.util.Map<String, ?> map) throws Exception {
        WriteGetObjectResponseHeaders self = new WriteGetObjectResponseHeaders();
        return TeaModel.build(map, self);
    }

    public WriteGetObjectResponseHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public WriteGetObjectResponseHeaders setContentLength(Long contentLength) {
        this.contentLength = contentLength;
        return this;
    }
    public Long getContentLength() {
        return this.contentLength;
    }

    public WriteGetObjectResponseHeaders setXOssFwdHeaderAcceptRanges(String xOssFwdHeaderAcceptRanges) {
        this.xOssFwdHeaderAcceptRanges = xOssFwdHeaderAcceptRanges;
        return this;
    }
    public String getXOssFwdHeaderAcceptRanges() {
        return this.xOssFwdHeaderAcceptRanges;
    }

    public WriteGetObjectResponseHeaders setXOssFwdHeaderCacheControl(String xOssFwdHeaderCacheControl) {
        this.xOssFwdHeaderCacheControl = xOssFwdHeaderCacheControl;
        return this;
    }
    public String getXOssFwdHeaderCacheControl() {
        return this.xOssFwdHeaderCacheControl;
    }

    public WriteGetObjectResponseHeaders setXOssFwdHeaderContentDisposition(String xOssFwdHeaderContentDisposition) {
        this.xOssFwdHeaderContentDisposition = xOssFwdHeaderContentDisposition;
        return this;
    }
    public String getXOssFwdHeaderContentDisposition() {
        return this.xOssFwdHeaderContentDisposition;
    }

    public WriteGetObjectResponseHeaders setXOssFwdHeaderContentEncoding(String xOssFwdHeaderContentEncoding) {
        this.xOssFwdHeaderContentEncoding = xOssFwdHeaderContentEncoding;
        return this;
    }
    public String getXOssFwdHeaderContentEncoding() {
        return this.xOssFwdHeaderContentEncoding;
    }

    public WriteGetObjectResponseHeaders setXOssFwdHeaderContentLanguage(String xOssFwdHeaderContentLanguage) {
        this.xOssFwdHeaderContentLanguage = xOssFwdHeaderContentLanguage;
        return this;
    }
    public String getXOssFwdHeaderContentLanguage() {
        return this.xOssFwdHeaderContentLanguage;
    }

    public WriteGetObjectResponseHeaders setXOssFwdHeaderContentRange(String xOssFwdHeaderContentRange) {
        this.xOssFwdHeaderContentRange = xOssFwdHeaderContentRange;
        return this;
    }
    public String getXOssFwdHeaderContentRange() {
        return this.xOssFwdHeaderContentRange;
    }

    public WriteGetObjectResponseHeaders setXOssFwdHeaderContentType(String xOssFwdHeaderContentType) {
        this.xOssFwdHeaderContentType = xOssFwdHeaderContentType;
        return this;
    }
    public String getXOssFwdHeaderContentType() {
        return this.xOssFwdHeaderContentType;
    }

    public WriteGetObjectResponseHeaders setXOssFwdHeaderETag(String xOssFwdHeaderETag) {
        this.xOssFwdHeaderETag = xOssFwdHeaderETag;
        return this;
    }
    public String getXOssFwdHeaderETag() {
        return this.xOssFwdHeaderETag;
    }

    public WriteGetObjectResponseHeaders setXOssFwdHeaderExpires(String xOssFwdHeaderExpires) {
        this.xOssFwdHeaderExpires = xOssFwdHeaderExpires;
        return this;
    }
    public String getXOssFwdHeaderExpires() {
        return this.xOssFwdHeaderExpires;
    }

    public WriteGetObjectResponseHeaders setXOssFwdHeaderLastModified(String xOssFwdHeaderLastModified) {
        this.xOssFwdHeaderLastModified = xOssFwdHeaderLastModified;
        return this;
    }
    public String getXOssFwdHeaderLastModified() {
        return this.xOssFwdHeaderLastModified;
    }

    public WriteGetObjectResponseHeaders setXOssFwdStatus(String xOssFwdStatus) {
        this.xOssFwdStatus = xOssFwdStatus;
        return this;
    }
    public String getXOssFwdStatus() {
        return this.xOssFwdStatus;
    }

    public WriteGetObjectResponseHeaders setXOssRequestRoute(String xOssRequestRoute) {
        this.xOssRequestRoute = xOssRequestRoute;
        return this;
    }
    public String getXOssRequestRoute() {
        return this.xOssRequestRoute;
    }

    public WriteGetObjectResponseHeaders setXOssRequestToken(String xOssRequestToken) {
        this.xOssRequestToken = xOssRequestToken;
        return this;
    }
    public String getXOssRequestToken() {
        return this.xOssRequestToken;
    }

}
