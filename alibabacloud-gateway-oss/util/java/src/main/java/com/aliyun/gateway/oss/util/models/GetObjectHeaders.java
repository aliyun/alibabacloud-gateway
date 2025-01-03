// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetObjectHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>The encoding type at the client side. 
     * If you want an object to be returned in the GZIP format, you must include the Accept-Encoding:gzip header in your request. OSS determines whether to return the object compressed in the GZip format based on the Content-Type header and whether the size of the object is larger than or equal to 1 KB.</p>
     * <blockquote>
     * <p>If an object is compressed in the GZip format, the response OSS returns does not include the ETag value of the object. </p>
     * <ul>
     * <li>OSS supports the following Content-Type values to compress the object in the GZip format: text/cache-manifest, text/xml, text/plain, text/css, application/javascript, application/x-javascript, application/rss+xml, application/json, and text/json.</li>
     * </ul>
     * </blockquote>
     * <p>Default value: null</p>
     * 
     * <strong>example:</strong>
     * <p>gzip</p>
     */
    @NameInMap("Accept-Encoding")
    public String acceptEncoding;

    /**
     * <p>If the ETag specified in the request matches the ETag value of the object, OSS transmits the object and returns 200 OK. If the ETag specified in the request does not match the ETag value of the object, OSS returns 412 Precondition Failed. 
     * The ETag value of an object is used to check whether the content of the object has changed. You can check data integrity by using the ETag value. 
     * Default value: null</p>
     * 
     * <strong>example:</strong>
     * <p>fba9dede5f27731c9771645a3986****</p>
     */
    @NameInMap("If-Match")
    public String ifMatch;

    /**
     * <p>If the time specified in this header is earlier than the object modified time or is invalid, OSS returns the object and 200 OK. If the time specified in this header is later than or the same as the object modified time, OSS returns 304 Not Modified. 
     * The time must be in GMT. Example: <code>Fri, 13 Nov 2015 14:47:53 GMT</code>.
     * Default value: null</p>
     * 
     * <strong>example:</strong>
     * <p>Fri, 13 Nov 2015 14:47:53 GMT</p>
     */
    @NameInMap("If-Modified-Since")
    public String ifModifiedSince;

    /**
     * <p>If the ETag specified in the request does not match the ETag value of the object, OSS transmits the object and returns 200 OK. If the ETag specified in the request matches the ETag value of the object, OSS returns 304 Not Modified. 
     * You can specify both the <strong>If-Match</strong> and <strong>If-None-Match</strong> headers in a request. 
     * Default value: null</p>
     * 
     * <strong>example:</strong>
     * <p>5B3C1A2E0563E1B002CC607C****</p>
     */
    @NameInMap("If-None-Match")
    public String ifNoneMatch;

    /**
     * <p>If the time specified in this header is the same as or later than the object modified time, OSS returns the object and 200 OK. If the time specified in this header is earlier than the object modified time, OSS returns 412 Precondition Failed.</p>
     * <p>The time must be in GMT. Example: <code>Fri, 13 Nov 2015 14:47:53 GMT</code>.
     * You can specify both the <strong>If-Modified-Since</strong> and <strong>If-Unmodified-Since</strong> headers in a request. 
     * Default value: null</p>
     * 
     * <strong>example:</strong>
     * <p>Fri, 13 Nov 2015 14:47:53 GMT</p>
     */
    @NameInMap("If-Unmodified-Since")
    public String ifUnmodifiedSince;

    /**
     * <p>The range of data of the object to be returned. </p>
     * <ul>
     * <li>If the value of Range is valid, OSS returns the response that includes the total size of the object and the range of data returned. For example, Content-Range: bytes 0~9/44 indicates that the total size of the object is 44 bytes, and the range of data returned is the first 10 bytes. </li>
     * <li>However, if the value of Range is invalid, the entire object is returned, and the response returned by OSS excludes Content-Range.</li>
     * </ul>
     * <p>Default value: null</p>
     * 
     * <strong>example:</strong>
     * <p>Content-Range: bytes 100-900/344606</p>
     */
    @NameInMap("Range")
    public String range;

    public static GetObjectHeaders build(java.util.Map<String, ?> map) throws Exception {
        GetObjectHeaders self = new GetObjectHeaders();
        return TeaModel.build(map, self);
    }

    public GetObjectHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public GetObjectHeaders setAcceptEncoding(String acceptEncoding) {
        this.acceptEncoding = acceptEncoding;
        return this;
    }
    public String getAcceptEncoding() {
        return this.acceptEncoding;
    }

    public GetObjectHeaders setIfMatch(String ifMatch) {
        this.ifMatch = ifMatch;
        return this;
    }
    public String getIfMatch() {
        return this.ifMatch;
    }

    public GetObjectHeaders setIfModifiedSince(String ifModifiedSince) {
        this.ifModifiedSince = ifModifiedSince;
        return this;
    }
    public String getIfModifiedSince() {
        return this.ifModifiedSince;
    }

    public GetObjectHeaders setIfNoneMatch(String ifNoneMatch) {
        this.ifNoneMatch = ifNoneMatch;
        return this;
    }
    public String getIfNoneMatch() {
        return this.ifNoneMatch;
    }

    public GetObjectHeaders setIfUnmodifiedSince(String ifUnmodifiedSince) {
        this.ifUnmodifiedSince = ifUnmodifiedSince;
        return this;
    }
    public String getIfUnmodifiedSince() {
        return this.ifUnmodifiedSince;
    }

    public GetObjectHeaders setRange(String range) {
        this.range = range;
        return this;
    }
    public String getRange() {
        return this.range;
    }

}
