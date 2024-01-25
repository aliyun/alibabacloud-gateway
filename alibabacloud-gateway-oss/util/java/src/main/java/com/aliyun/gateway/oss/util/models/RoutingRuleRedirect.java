// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class RoutingRuleRedirect extends TeaModel {
    @NameInMap("EnableReplacePrefix")
    public Boolean enableReplacePrefix;

    @NameInMap("HostName")
    public String hostName;

    @NameInMap("HttpRedirectCode")
    public Long httpRedirectCode;

    @NameInMap("MirrorCheckMd5")
    public Boolean mirrorCheckMd5;

    @NameInMap("MirrorFollowRedirect")
    public Boolean mirrorFollowRedirect;

    @NameInMap("MirrorHeaders")
    public RoutingRuleRedirectMirrorHeaders mirrorHeaders;

    @NameInMap("MirrorPassQueryString")
    public Boolean mirrorPassQueryString;

    @NameInMap("MirrorURL")
    public String mirrorURL;

    @NameInMap("PassQueryString")
    public Boolean passQueryString;

    @NameInMap("Protocol")
    public String protocol;

    @NameInMap("RedirectType")
    public String redirectType;

    @NameInMap("ReplaceKeyPrefixWith")
    public String replaceKeyPrefixWith;

    @NameInMap("ReplaceKeyWith")
    public String replaceKeyWith;

    @NameInMap("TransparentMirrorResponseCodes")
    public String transparentMirrorResponseCodes;

    public static RoutingRuleRedirect build(java.util.Map<String, ?> map) throws Exception {
        RoutingRuleRedirect self = new RoutingRuleRedirect();
        return TeaModel.build(map, self);
    }

    public RoutingRuleRedirect setEnableReplacePrefix(Boolean enableReplacePrefix) {
        this.enableReplacePrefix = enableReplacePrefix;
        return this;
    }
    public Boolean getEnableReplacePrefix() {
        return this.enableReplacePrefix;
    }

    public RoutingRuleRedirect setHostName(String hostName) {
        this.hostName = hostName;
        return this;
    }
    public String getHostName() {
        return this.hostName;
    }

    public RoutingRuleRedirect setHttpRedirectCode(Long httpRedirectCode) {
        this.httpRedirectCode = httpRedirectCode;
        return this;
    }
    public Long getHttpRedirectCode() {
        return this.httpRedirectCode;
    }

    public RoutingRuleRedirect setMirrorCheckMd5(Boolean mirrorCheckMd5) {
        this.mirrorCheckMd5 = mirrorCheckMd5;
        return this;
    }
    public Boolean getMirrorCheckMd5() {
        return this.mirrorCheckMd5;
    }

    public RoutingRuleRedirect setMirrorFollowRedirect(Boolean mirrorFollowRedirect) {
        this.mirrorFollowRedirect = mirrorFollowRedirect;
        return this;
    }
    public Boolean getMirrorFollowRedirect() {
        return this.mirrorFollowRedirect;
    }

    public RoutingRuleRedirect setMirrorHeaders(RoutingRuleRedirectMirrorHeaders mirrorHeaders) {
        this.mirrorHeaders = mirrorHeaders;
        return this;
    }
    public RoutingRuleRedirectMirrorHeaders getMirrorHeaders() {
        return this.mirrorHeaders;
    }

    public RoutingRuleRedirect setMirrorPassQueryString(Boolean mirrorPassQueryString) {
        this.mirrorPassQueryString = mirrorPassQueryString;
        return this;
    }
    public Boolean getMirrorPassQueryString() {
        return this.mirrorPassQueryString;
    }

    public RoutingRuleRedirect setMirrorURL(String mirrorURL) {
        this.mirrorURL = mirrorURL;
        return this;
    }
    public String getMirrorURL() {
        return this.mirrorURL;
    }

    public RoutingRuleRedirect setPassQueryString(Boolean passQueryString) {
        this.passQueryString = passQueryString;
        return this;
    }
    public Boolean getPassQueryString() {
        return this.passQueryString;
    }

    public RoutingRuleRedirect setProtocol(String protocol) {
        this.protocol = protocol;
        return this;
    }
    public String getProtocol() {
        return this.protocol;
    }

    public RoutingRuleRedirect setRedirectType(String redirectType) {
        this.redirectType = redirectType;
        return this;
    }
    public String getRedirectType() {
        return this.redirectType;
    }

    public RoutingRuleRedirect setReplaceKeyPrefixWith(String replaceKeyPrefixWith) {
        this.replaceKeyPrefixWith = replaceKeyPrefixWith;
        return this;
    }
    public String getReplaceKeyPrefixWith() {
        return this.replaceKeyPrefixWith;
    }

    public RoutingRuleRedirect setReplaceKeyWith(String replaceKeyWith) {
        this.replaceKeyWith = replaceKeyWith;
        return this;
    }
    public String getReplaceKeyWith() {
        return this.replaceKeyWith;
    }

    public RoutingRuleRedirect setTransparentMirrorResponseCodes(String transparentMirrorResponseCodes) {
        this.transparentMirrorResponseCodes = transparentMirrorResponseCodes;
        return this;
    }
    public String getTransparentMirrorResponseCodes() {
        return this.transparentMirrorResponseCodes;
    }

    public static class RoutingRuleRedirectMirrorHeadersSet extends TeaModel {
        @NameInMap("Key")
        public String key;

        @NameInMap("Value")
        public String value;

        public static RoutingRuleRedirectMirrorHeadersSet build(java.util.Map<String, ?> map) throws Exception {
            RoutingRuleRedirectMirrorHeadersSet self = new RoutingRuleRedirectMirrorHeadersSet();
            return TeaModel.build(map, self);
        }

        public RoutingRuleRedirectMirrorHeadersSet setKey(String key) {
            this.key = key;
            return this;
        }
        public String getKey() {
            return this.key;
        }

        public RoutingRuleRedirectMirrorHeadersSet setValue(String value) {
            this.value = value;
            return this;
        }
        public String getValue() {
            return this.value;
        }

    }

    public static class RoutingRuleRedirectMirrorHeaders extends TeaModel {
        @NameInMap("Pass")
        public java.util.List<String> pass;

        @NameInMap("PassAll")
        public Boolean passAll;

        @NameInMap("Remove")
        public java.util.List<String> remove;

        @NameInMap("Set")
        public java.util.List<RoutingRuleRedirectMirrorHeadersSet> set;

        public static RoutingRuleRedirectMirrorHeaders build(java.util.Map<String, ?> map) throws Exception {
            RoutingRuleRedirectMirrorHeaders self = new RoutingRuleRedirectMirrorHeaders();
            return TeaModel.build(map, self);
        }

        public RoutingRuleRedirectMirrorHeaders setPass(java.util.List<String> pass) {
            this.pass = pass;
            return this;
        }
        public java.util.List<String> getPass() {
            return this.pass;
        }

        public RoutingRuleRedirectMirrorHeaders setPassAll(Boolean passAll) {
            this.passAll = passAll;
            return this;
        }
        public Boolean getPassAll() {
            return this.passAll;
        }

        public RoutingRuleRedirectMirrorHeaders setRemove(java.util.List<String> remove) {
            this.remove = remove;
            return this;
        }
        public java.util.List<String> getRemove() {
            return this.remove;
        }

        public RoutingRuleRedirectMirrorHeaders setSet(java.util.List<RoutingRuleRedirectMirrorHeadersSet> set) {
            this.set = set;
            return this;
        }
        public java.util.List<RoutingRuleRedirectMirrorHeadersSet> getSet() {
            return this.set;
        }

    }

}
