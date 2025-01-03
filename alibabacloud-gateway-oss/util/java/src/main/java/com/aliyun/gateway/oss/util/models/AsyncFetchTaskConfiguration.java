// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class AsyncFetchTaskConfiguration extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>eyJjYWxsYmFja1VybCI6Ind3dy5hYmMuY29tL2NhbGxiYWNrIiwiY2FsbGJhY2tCb2R5IjoiJHtldGFnfSJ9</p>
     */
    @NameInMap("Callback")
    public String callback;

    /**
     * <strong>example:</strong>
     * <p>v23MlMRM/EgJczOs2yHTcA==</p>
     */
    @NameInMap("ContentMD5")
    public String contentMD5;

    /**
     * <strong>example:</strong>
     * <p><a href="http://www.test.com">www.test.com</a></p>
     */
    @NameInMap("Host")
    public String host;

    /**
     * <strong>example:</strong>
     * <p>True</p>
     */
    @NameInMap("IgnoreSameKey")
    public Boolean ignoreSameKey;

    /**
     * <strong>example:</strong>
     * <p>abc.txt</p>
     */
    @NameInMap("Object")
    public String object;

    /**
     * <strong>example:</strong>
     * <p>Standard</p>
     */
    @NameInMap("StorageClass")
    public String storageClass;

    /**
     * <strong>example:</strong>
     * <p><a href="http://www.test.com/abc.txt">www.test.com/abc.txt</a></p>
     */
    @NameInMap("Url")
    public String url;

    public static AsyncFetchTaskConfiguration build(java.util.Map<String, ?> map) throws Exception {
        AsyncFetchTaskConfiguration self = new AsyncFetchTaskConfiguration();
        return TeaModel.build(map, self);
    }

    public AsyncFetchTaskConfiguration setCallback(String callback) {
        this.callback = callback;
        return this;
    }
    public String getCallback() {
        return this.callback;
    }

    public AsyncFetchTaskConfiguration setContentMD5(String contentMD5) {
        this.contentMD5 = contentMD5;
        return this;
    }
    public String getContentMD5() {
        return this.contentMD5;
    }

    public AsyncFetchTaskConfiguration setHost(String host) {
        this.host = host;
        return this;
    }
    public String getHost() {
        return this.host;
    }

    public AsyncFetchTaskConfiguration setIgnoreSameKey(Boolean ignoreSameKey) {
        this.ignoreSameKey = ignoreSameKey;
        return this;
    }
    public Boolean getIgnoreSameKey() {
        return this.ignoreSameKey;
    }

    public AsyncFetchTaskConfiguration setObject(String object) {
        this.object = object;
        return this;
    }
    public String getObject() {
        return this.object;
    }

    public AsyncFetchTaskConfiguration setStorageClass(String storageClass) {
        this.storageClass = storageClass;
        return this;
    }
    public String getStorageClass() {
        return this.storageClass;
    }

    public AsyncFetchTaskConfiguration setUrl(String url) {
        this.url = url;
        return this;
    }
    public String getUrl() {
        return this.url;
    }

}
