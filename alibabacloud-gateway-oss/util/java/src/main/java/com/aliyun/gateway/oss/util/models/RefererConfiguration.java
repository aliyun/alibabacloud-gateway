// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class RefererConfiguration extends TeaModel {
    @NameInMap("AllowEmptyReferer")
    public Boolean allowEmptyReferer;

    @NameInMap("AllowTruncateQueryString")
    public Boolean allowTruncateQueryString;

    @NameInMap("RefererBlacklist")
    public RefererConfigurationRefererBlacklist refererBlacklist;

    @NameInMap("RefererList")
    public RefererConfigurationRefererList refererList;

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

    public RefererConfiguration setRefererBlacklist(RefererConfigurationRefererBlacklist refererBlacklist) {
        this.refererBlacklist = refererBlacklist;
        return this;
    }
    public RefererConfigurationRefererBlacklist getRefererBlacklist() {
        return this.refererBlacklist;
    }

    public RefererConfiguration setRefererList(RefererConfigurationRefererList refererList) {
        this.refererList = refererList;
        return this;
    }
    public RefererConfigurationRefererList getRefererList() {
        return this.refererList;
    }

    public RefererConfiguration setTruncatePath(Boolean truncatePath) {
        this.truncatePath = truncatePath;
        return this;
    }
    public Boolean getTruncatePath() {
        return this.truncatePath;
    }

    public static class RefererConfigurationRefererBlacklist extends TeaModel {
        @NameInMap("Referer")
        public java.util.List<String> referer;

        public static RefererConfigurationRefererBlacklist build(java.util.Map<String, ?> map) throws Exception {
            RefererConfigurationRefererBlacklist self = new RefererConfigurationRefererBlacklist();
            return TeaModel.build(map, self);
        }

        public RefererConfigurationRefererBlacklist setReferer(java.util.List<String> referer) {
            this.referer = referer;
            return this;
        }
        public java.util.List<String> getReferer() {
            return this.referer;
        }

    }

    public static class RefererConfigurationRefererList extends TeaModel {
        @NameInMap("Referer")
        public java.util.List<String> referer;

        public static RefererConfigurationRefererList build(java.util.Map<String, ?> map) throws Exception {
            RefererConfigurationRefererList self = new RefererConfigurationRefererList();
            return TeaModel.build(map, self);
        }

        public RefererConfigurationRefererList setReferer(java.util.List<String> referer) {
            this.referer = referer;
            return this;
        }
        public java.util.List<String> getReferer() {
            return this.referer;
        }

    }

}
