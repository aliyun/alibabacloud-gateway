// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class HeadObjectHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>If the ETag value that is specified in the request matches the ETag value of the object, OSS returns 200 OK and the metadata of the object. Otherwise, OSS returns 412 precondition failed. 
     * Default value: null.</p>
     * 
     * <strong>example:</strong>
     * <p>fba9dede5f27731c9771645a3986****</p>
     */
    @NameInMap("If-Match")
    public String ifMatch;

    /**
     * <p>If the time that is specified in the request is earlier than the time when the object is modified, OSS returns 200 OK and the metadata of the object. Otherwise, OSS returns 304 not modified. 
     * Default value: null.</p>
     * 
     * <strong>example:</strong>
     * <p>Fri, 9 Apr 2021 14:47:53 GMT</p>
     */
    @NameInMap("If-Modified-Since")
    public String ifModifiedSince;

    /**
     * <p>If the ETag value that is specified in the request does not match the ETag value of the object, OSS returns 200 OK and the metadata of the object. Otherwise, OSS returns 304 Not Modified. 
     * Default value: null.</p>
     * 
     * <strong>example:</strong>
     * <p>5B3C1A2E0563E1B002CC607C****</p>
     */
    @NameInMap("If-None-Match")
    public String ifNoneMatch;

    /**
     * <p>If the time that is specified in the request is later than or the same as the time when the object is modified, OSS returns 200 OK and the metadata of the object. Otherwise, OSS returns 412 precondition failed. 
     * Default value: null.</p>
     * 
     * <strong>example:</strong>
     * <p>Fri, 13 Oct 2021 14:47:53 GMT</p>
     */
    @NameInMap("If-Unmodified-Since")
    public String ifUnmodifiedSince;

    public static HeadObjectHeaders build(java.util.Map<String, ?> map) throws Exception {
        HeadObjectHeaders self = new HeadObjectHeaders();
        return TeaModel.build(map, self);
    }

    public HeadObjectHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public HeadObjectHeaders setIfMatch(String ifMatch) {
        this.ifMatch = ifMatch;
        return this;
    }
    public String getIfMatch() {
        return this.ifMatch;
    }

    public HeadObjectHeaders setIfModifiedSince(String ifModifiedSince) {
        this.ifModifiedSince = ifModifiedSince;
        return this;
    }
    public String getIfModifiedSince() {
        return this.ifModifiedSince;
    }

    public HeadObjectHeaders setIfNoneMatch(String ifNoneMatch) {
        this.ifNoneMatch = ifNoneMatch;
        return this;
    }
    public String getIfNoneMatch() {
        return this.ifNoneMatch;
    }

    public HeadObjectHeaders setIfUnmodifiedSince(String ifUnmodifiedSince) {
        this.ifUnmodifiedSince = ifUnmodifiedSince;
        return this;
    }
    public String getIfUnmodifiedSince() {
        return this.ifUnmodifiedSince;
    }

}
