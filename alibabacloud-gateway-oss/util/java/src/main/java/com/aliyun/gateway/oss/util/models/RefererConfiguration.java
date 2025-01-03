// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class RefererConfiguration extends TeaModel {
    /**
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>true</p>
     */
    @NameInMap("AllowEmptyReferer")
    public Boolean allowEmptyReferer;

    /**
     * <strong>example:</strong>
     * <p>false</p>
     */
    @NameInMap("AllowTruncateQueryString")
    public Boolean allowTruncateQueryString;

    @NameInMap("RefererBlacklist")
    public RefererBlacklist refererBlacklist;

    /**
     * <p>This parameter is required.</p>
     */
    @NameInMap("RefererList")
    public RefererList refererList;

    /**
     * <strong>example:</strong>
     * <p>false</p>
     */
    @NameInMap("TruncatePath")
    public Boolean truncatePath;

    public static RefererConfiguration build(java.util.Map<String, ?> map) throws Exception {
        RefererConfiguration self = new RefererConfiguration();
        return TeaModel.build(map, self);
    }

    public RefererConfiguration setAllowEmptyReferer(Boolean allowEmptyReferer) {
        this.allowEmptyReferer = allowEmptyReferer;
        return this;
    }
    public Boolean getAllowEmptyReferer() {
        return this.allowEmptyReferer;
    }

    public RefererConfiguration setAllowTruncateQueryString(Boolean allowTruncateQueryString) {
        this.allowTruncateQueryString = allowTruncateQueryString;
        return this;
    }
    public Boolean getAllowTruncateQueryString() {
        return this.allowTruncateQueryString;
    }

    public RefererConfiguration setRefererBlacklist(RefererBlacklist refererBlacklist) {
        this.refererBlacklist = refererBlacklist;
        return this;
    }
    public RefererBlacklist getRefererBlacklist() {
        return this.refererBlacklist;
    }

    public RefererConfiguration setRefererList(RefererList refererList) {
        this.refererList = refererList;
        return this;
    }
    public RefererList getRefererList() {
        return this.refererList;
    }

    public RefererConfiguration setTruncatePath(Boolean truncatePath) {
        this.truncatePath = truncatePath;
        return this;
    }
    public Boolean getTruncatePath() {
        return this.truncatePath;
    }

    public static class RefererBlacklist extends TeaModel {
        @NameInMap("Referer")
        public java.util.List<String> referer;

        public static RefererBlacklist build(java.util.Map<String, ?> map) throws Exception {
            RefererBlacklist self = new RefererBlacklist();
            return TeaModel.build(map, self);
        }

        public RefererBlacklist setReferer(java.util.List<String> referer) {
            this.referer = referer;
            return this;
        }
        public java.util.List<String> getReferer() {
            return this.referer;
        }

    }

    public static class RefererList extends TeaModel {
        @NameInMap("Referer")
        public java.util.List<String> referer;

        public static RefererList build(java.util.Map<String, ?> map) throws Exception {
            RefererList self = new RefererList();
            return TeaModel.build(map, self);
        }

        public RefererList setReferer(java.util.List<String> referer) {
            this.referer = referer;
            return this;
        }
        public java.util.List<String> getReferer() {
            return this.referer;
        }

    }

}
