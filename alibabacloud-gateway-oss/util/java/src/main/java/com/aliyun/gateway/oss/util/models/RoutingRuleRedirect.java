// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class RoutingRuleRedirect extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>true</p>
     */
    @NameInMap("EnableReplacePrefix")
    public Boolean enableReplacePrefix;

    /**
     * <strong>example:</strong>
     * <p>example.com</p>
     */
    @NameInMap("HostName")
    public String hostName;

    /**
     * <strong>example:</strong>
     * <p>302</p>
     */
    @NameInMap("HttpRedirectCode")
    public Long httpRedirectCode;

    /**
     * <strong>example:</strong>
     * <p>true</p>
     */
    @NameInMap("MirrorAllowGetImageInfo")
    public Boolean mirrorAllowGetImageInfo;

    /**
     * <strong>example:</strong>
     * <p>true</p>
     */
    @NameInMap("MirrorAllowHeadObject")
    public Boolean mirrorAllowHeadObject;

    /**
     * <strong>example:</strong>
     * <p>false</p>
     */
    @NameInMap("MirrorAllowVideoSnapshot")
    public Boolean mirrorAllowVideoSnapshot;

    /**
     * <strong>example:</strong>
     * <p>303</p>
     */
    @NameInMap("MirrorAsyncStatus")
    public Long mirrorAsyncStatus;

    @NameInMap("MirrorAuth")
    public MirrorAuth mirrorAuth;

    /**
     * <strong>example:</strong>
     * <p>false</p>
     */
    @NameInMap("MirrorCheckMd5")
    public Boolean mirrorCheckMd5;

    /**
     * <strong>example:</strong>
     * <p>cn-hangzhou</p>
     */
    @NameInMap("MirrorDstRegion")
    public String mirrorDstRegion;

    /**
     * <strong>example:</strong>
     * <p>vpc-test-id</p>
     */
    @NameInMap("MirrorDstSlaveVpcId")
    public String mirrorDstSlaveVpcId;

    /**
     * <strong>example:</strong>
     * <p>vpc-test-id</p>
     */
    @NameInMap("MirrorDstVpcId")
    public String mirrorDstVpcId;

    /**
     * <strong>example:</strong>
     * <p>true</p>
     */
    @NameInMap("MirrorFollowRedirect")
    public Boolean mirrorFollowRedirect;

    @NameInMap("MirrorHeaders")
    public MirrorHeaders mirrorHeaders;

    /**
     * <strong>example:</strong>
     * <p>true</p>
     */
    @NameInMap("MirrorIsExpressTunnel")
    public Boolean mirrorIsExpressTunnel;

    @NameInMap("MirrorMultiAlternates")
    public MirrorMultiAlternates mirrorMultiAlternates;

    /**
     * <strong>example:</strong>
     * <p>false</p>
     */
    @NameInMap("MirrorPassOriginalSlashes")
    public Boolean mirrorPassOriginalSlashes;

    /**
     * <strong>example:</strong>
     * <p>true</p>
     */
    @NameInMap("MirrorPassQueryString")
    public Boolean mirrorPassQueryString;

    /**
     * <strong>example:</strong>
     * <p>false</p>
     */
    @NameInMap("MirrorProxyPass")
    public Boolean mirrorProxyPass;

    @NameInMap("MirrorReturnHeaders")
    public MirrorReturnHeaders mirrorReturnHeaders;

    /**
     * <strong>example:</strong>
     * <p>aliyun-test-role</p>
     */
    @NameInMap("MirrorRole")
    public String mirrorRole;

    /**
     * <strong>example:</strong>
     * <p>true</p>
     */
    @NameInMap("MirrorSNI")
    public Boolean mirrorSNI;

    /**
     * <strong>example:</strong>
     * <p>true</p>
     */
    @NameInMap("MirrorSaveOssMeta")
    public Boolean mirrorSaveOssMeta;

    /**
     * <strong>example:</strong>
     * <p>false</p>
     */
    @NameInMap("MirrorSwitchAllErrors")
    public Boolean mirrorSwitchAllErrors;

    @NameInMap("MirrorTaggings")
    public MirrorTaggings mirrorTaggings;

    /**
     * <strong>example:</strong>
     * <p>test-tunnel-id</p>
     */
    @NameInMap("MirrorTunnelId")
    public String mirrorTunnelId;

    /**
     * <strong>example:</strong>
     * <p><a href="http://example.com/">http://example.com/</a></p>
     */
    @NameInMap("MirrorURL")
    public String mirrorURL;

    /**
     * <strong>example:</strong>
     * <p><a href="https://example.com/hartbeat">https://example.com/hartbeat</a></p>
     */
    @NameInMap("MirrorURLProbe")
    public String mirrorURLProbe;

    /**
     * <strong>example:</strong>
     * <p><a href="https://example.com">https://example.com</a></p>
     */
    @NameInMap("MirrorURLSlave")
    public String mirrorURLSlave;

    /**
     * <strong>example:</strong>
     * <p>false</p>
     */
    @NameInMap("MirrorUserLastModified")
    public Boolean mirrorUserLastModified;

    /**
     * <strong>example:</strong>
     * <p>false</p>
     */
    @NameInMap("MirrorUsingRole")
    public Boolean mirrorUsingRole;

    /**
     * <strong>example:</strong>
     * <p>true</p>
     */
    @NameInMap("PassQueryString")
    public Boolean passQueryString;

    /**
     * <strong>example:</strong>
     * <p>https</p>
     */
    @NameInMap("Protocol")
    public String protocol;

    /**
     * <strong>example:</strong>
     * <p>Mirror</p>
     */
    @NameInMap("RedirectType")
    public String redirectType;

    /**
     * <strong>example:</strong>
     * <p>def/</p>
     */
    @NameInMap("ReplaceKeyPrefixWith")
    public String replaceKeyPrefixWith;

    /**
     * <strong>example:</strong>
     * <p>prefix/${key}.suffix</p>
     */
    @NameInMap("ReplaceKeyWith")
    public String replaceKeyWith;

    /**
     * <strong>example:</strong>
     * <p>false</p>
     */
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

    public RoutingRuleRedirect setMirrorAllowGetImageInfo(Boolean mirrorAllowGetImageInfo) {
        this.mirrorAllowGetImageInfo = mirrorAllowGetImageInfo;
        return this;
    }
    public Boolean getMirrorAllowGetImageInfo() {
        return this.mirrorAllowGetImageInfo;
    }

    public RoutingRuleRedirect setMirrorAllowHeadObject(Boolean mirrorAllowHeadObject) {
        this.mirrorAllowHeadObject = mirrorAllowHeadObject;
        return this;
    }
    public Boolean getMirrorAllowHeadObject() {
        return this.mirrorAllowHeadObject;
    }

    public RoutingRuleRedirect setMirrorAllowVideoSnapshot(Boolean mirrorAllowVideoSnapshot) {
        this.mirrorAllowVideoSnapshot = mirrorAllowVideoSnapshot;
        return this;
    }
    public Boolean getMirrorAllowVideoSnapshot() {
        return this.mirrorAllowVideoSnapshot;
    }

    public RoutingRuleRedirect setMirrorAsyncStatus(Long mirrorAsyncStatus) {
        this.mirrorAsyncStatus = mirrorAsyncStatus;
        return this;
    }
    public Long getMirrorAsyncStatus() {
        return this.mirrorAsyncStatus;
    }

    public RoutingRuleRedirect setMirrorAuth(MirrorAuth mirrorAuth) {
        this.mirrorAuth = mirrorAuth;
        return this;
    }
    public MirrorAuth getMirrorAuth() {
        return this.mirrorAuth;
    }

    public RoutingRuleRedirect setMirrorCheckMd5(Boolean mirrorCheckMd5) {
        this.mirrorCheckMd5 = mirrorCheckMd5;
        return this;
    }
    public Boolean getMirrorCheckMd5() {
        return this.mirrorCheckMd5;
    }

    public RoutingRuleRedirect setMirrorDstRegion(String mirrorDstRegion) {
        this.mirrorDstRegion = mirrorDstRegion;
        return this;
    }
    public String getMirrorDstRegion() {
        return this.mirrorDstRegion;
    }

    public RoutingRuleRedirect setMirrorDstSlaveVpcId(String mirrorDstSlaveVpcId) {
        this.mirrorDstSlaveVpcId = mirrorDstSlaveVpcId;
        return this;
    }
    public String getMirrorDstSlaveVpcId() {
        return this.mirrorDstSlaveVpcId;
    }

    public RoutingRuleRedirect setMirrorDstVpcId(String mirrorDstVpcId) {
        this.mirrorDstVpcId = mirrorDstVpcId;
        return this;
    }
    public String getMirrorDstVpcId() {
        return this.mirrorDstVpcId;
    }

    public RoutingRuleRedirect setMirrorFollowRedirect(Boolean mirrorFollowRedirect) {
        this.mirrorFollowRedirect = mirrorFollowRedirect;
        return this;
    }
    public Boolean getMirrorFollowRedirect() {
        return this.mirrorFollowRedirect;
    }

    public RoutingRuleRedirect setMirrorHeaders(MirrorHeaders mirrorHeaders) {
        this.mirrorHeaders = mirrorHeaders;
        return this;
    }
    public MirrorHeaders getMirrorHeaders() {
        return this.mirrorHeaders;
    }

    public RoutingRuleRedirect setMirrorIsExpressTunnel(Boolean mirrorIsExpressTunnel) {
        this.mirrorIsExpressTunnel = mirrorIsExpressTunnel;
        return this;
    }
    public Boolean getMirrorIsExpressTunnel() {
        return this.mirrorIsExpressTunnel;
    }

    public RoutingRuleRedirect setMirrorMultiAlternates(MirrorMultiAlternates mirrorMultiAlternates) {
        this.mirrorMultiAlternates = mirrorMultiAlternates;
        return this;
    }
    public MirrorMultiAlternates getMirrorMultiAlternates() {
        return this.mirrorMultiAlternates;
    }

    public RoutingRuleRedirect setMirrorPassOriginalSlashes(Boolean mirrorPassOriginalSlashes) {
        this.mirrorPassOriginalSlashes = mirrorPassOriginalSlashes;
        return this;
    }
    public Boolean getMirrorPassOriginalSlashes() {
        return this.mirrorPassOriginalSlashes;
    }

    public RoutingRuleRedirect setMirrorPassQueryString(Boolean mirrorPassQueryString) {
        this.mirrorPassQueryString = mirrorPassQueryString;
        return this;
    }
    public Boolean getMirrorPassQueryString() {
        return this.mirrorPassQueryString;
    }

    public RoutingRuleRedirect setMirrorProxyPass(Boolean mirrorProxyPass) {
        this.mirrorProxyPass = mirrorProxyPass;
        return this;
    }
    public Boolean getMirrorProxyPass() {
        return this.mirrorProxyPass;
    }

    public RoutingRuleRedirect setMirrorReturnHeaders(MirrorReturnHeaders mirrorReturnHeaders) {
        this.mirrorReturnHeaders = mirrorReturnHeaders;
        return this;
    }
    public MirrorReturnHeaders getMirrorReturnHeaders() {
        return this.mirrorReturnHeaders;
    }

    public RoutingRuleRedirect setMirrorRole(String mirrorRole) {
        this.mirrorRole = mirrorRole;
        return this;
    }
    public String getMirrorRole() {
        return this.mirrorRole;
    }

    public RoutingRuleRedirect setMirrorSNI(Boolean mirrorSNI) {
        this.mirrorSNI = mirrorSNI;
        return this;
    }
    public Boolean getMirrorSNI() {
        return this.mirrorSNI;
    }

    public RoutingRuleRedirect setMirrorSaveOssMeta(Boolean mirrorSaveOssMeta) {
        this.mirrorSaveOssMeta = mirrorSaveOssMeta;
        return this;
    }
    public Boolean getMirrorSaveOssMeta() {
        return this.mirrorSaveOssMeta;
    }

    public RoutingRuleRedirect setMirrorSwitchAllErrors(Boolean mirrorSwitchAllErrors) {
        this.mirrorSwitchAllErrors = mirrorSwitchAllErrors;
        return this;
    }
    public Boolean getMirrorSwitchAllErrors() {
        return this.mirrorSwitchAllErrors;
    }

    public RoutingRuleRedirect setMirrorTaggings(MirrorTaggings mirrorTaggings) {
        this.mirrorTaggings = mirrorTaggings;
        return this;
    }
    public MirrorTaggings getMirrorTaggings() {
        return this.mirrorTaggings;
    }

    public RoutingRuleRedirect setMirrorTunnelId(String mirrorTunnelId) {
        this.mirrorTunnelId = mirrorTunnelId;
        return this;
    }
    public String getMirrorTunnelId() {
        return this.mirrorTunnelId;
    }

    public RoutingRuleRedirect setMirrorURL(String mirrorURL) {
        this.mirrorURL = mirrorURL;
        return this;
    }
    public String getMirrorURL() {
        return this.mirrorURL;
    }

    public RoutingRuleRedirect setMirrorURLProbe(String mirrorURLProbe) {
        this.mirrorURLProbe = mirrorURLProbe;
        return this;
    }
    public String getMirrorURLProbe() {
        return this.mirrorURLProbe;
    }

    public RoutingRuleRedirect setMirrorURLSlave(String mirrorURLSlave) {
        this.mirrorURLSlave = mirrorURLSlave;
        return this;
    }
    public String getMirrorURLSlave() {
        return this.mirrorURLSlave;
    }

    public RoutingRuleRedirect setMirrorUserLastModified(Boolean mirrorUserLastModified) {
        this.mirrorUserLastModified = mirrorUserLastModified;
        return this;
    }
    public Boolean getMirrorUserLastModified() {
        return this.mirrorUserLastModified;
    }

    public RoutingRuleRedirect setMirrorUsingRole(Boolean mirrorUsingRole) {
        this.mirrorUsingRole = mirrorUsingRole;
        return this;
    }
    public Boolean getMirrorUsingRole() {
        return this.mirrorUsingRole;
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

    public static class MirrorAuth extends TeaModel {
        /**
         * <strong>example:</strong>
         * <p>TESTAK</p>
         */
        @NameInMap("AccessKeyId")
        public String accessKeyId;

        /**
         * <strong>example:</strong>
         * <p>TESTSK</p>
         */
        @NameInMap("AccessKeySecret")
        public String accessKeySecret;

        /**
         * <strong>example:</strong>
         * <p>S3V4</p>
         */
        @NameInMap("AuthType")
        public String authType;

        /**
         * <strong>example:</strong>
         * <p>ap-southeast-1</p>
         */
        @NameInMap("Region")
        public String region;

        public static MirrorAuth build(java.util.Map<String, ?> map) throws Exception {
            MirrorAuth self = new MirrorAuth();
            return TeaModel.build(map, self);
        }

        public MirrorAuth setAccessKeyId(String accessKeyId) {
            this.accessKeyId = accessKeyId;
            return this;
        }
        public String getAccessKeyId() {
            return this.accessKeyId;
        }

        public MirrorAuth setAccessKeySecret(String accessKeySecret) {
            this.accessKeySecret = accessKeySecret;
            return this;
        }
        public String getAccessKeySecret() {
            return this.accessKeySecret;
        }

        public MirrorAuth setAuthType(String authType) {
            this.authType = authType;
            return this;
        }
        public String getAuthType() {
            return this.authType;
        }

        public MirrorAuth setRegion(String region) {
            this.region = region;
            return this;
        }
        public String getRegion() {
            return this.region;
        }

    }

    public static class Set extends TeaModel {
        /**
         * <strong>example:</strong>
         * <p>set-header</p>
         */
        @NameInMap("Key")
        public String key;

        /**
         * <strong>example:</strong>
         * <p>set-value</p>
         */
        @NameInMap("Value")
        public String value;

        public static Set build(java.util.Map<String, ?> map) throws Exception {
            Set self = new Set();
            return TeaModel.build(map, self);
        }

        public Set setKey(String key) {
            this.key = key;
            return this;
        }
        public String getKey() {
            return this.key;
        }

        public Set setValue(String value) {
            this.value = value;
            return this;
        }
        public String getValue() {
            return this.value;
        }

    }

    public static class MirrorHeaders extends TeaModel {
        @NameInMap("Pass")
        public java.util.List<String> pass;

        /**
         * <strong>example:</strong>
         * <p>false</p>
         */
        @NameInMap("PassAll")
        public Boolean passAll;

        @NameInMap("Remove")
        public java.util.List<String> remove;

        @NameInMap("Set")
        public java.util.List<Set> set;

        public static MirrorHeaders build(java.util.Map<String, ?> map) throws Exception {
            MirrorHeaders self = new MirrorHeaders();
            return TeaModel.build(map, self);
        }

        public MirrorHeaders setPass(java.util.List<String> pass) {
            this.pass = pass;
            return this;
        }
        public java.util.List<String> getPass() {
            return this.pass;
        }

        public MirrorHeaders setPassAll(Boolean passAll) {
            this.passAll = passAll;
            return this;
        }
        public Boolean getPassAll() {
            return this.passAll;
        }

        public MirrorHeaders setRemove(java.util.List<String> remove) {
            this.remove = remove;
            return this;
        }
        public java.util.List<String> getRemove() {
            return this.remove;
        }

        public MirrorHeaders setSet(java.util.List<Set> set) {
            this.set = set;
            return this;
        }
        public java.util.List<Set> getSet() {
            return this.set;
        }

    }

    public static class MirrorMultiAlternate extends TeaModel {
        /**
         * <strong>example:</strong>
         * <p>ap-southeast-1</p>
         */
        @NameInMap("MirrorMultiAlternateDstRegion")
        public String mirrorMultiAlternateDstRegion;

        /**
         * <strong>example:</strong>
         * <p>32</p>
         */
        @NameInMap("MirrorMultiAlternateNumber")
        public Long mirrorMultiAlternateNumber;

        /**
         * <strong>example:</strong>
         * <p><a href="https://test-multi-alter.example.com">https://test-multi-alter.example.com</a></p>
         */
        @NameInMap("MirrorMultiAlternateURL")
        public String mirrorMultiAlternateURL;

        /**
         * <strong>example:</strong>
         * <p>vpc-test-id</p>
         */
        @NameInMap("MirrorMultiAlternateVpcId")
        public String mirrorMultiAlternateVpcId;

        public static MirrorMultiAlternate build(java.util.Map<String, ?> map) throws Exception {
            MirrorMultiAlternate self = new MirrorMultiAlternate();
            return TeaModel.build(map, self);
        }

        public MirrorMultiAlternate setMirrorMultiAlternateDstRegion(String mirrorMultiAlternateDstRegion) {
            this.mirrorMultiAlternateDstRegion = mirrorMultiAlternateDstRegion;
            return this;
        }
        public String getMirrorMultiAlternateDstRegion() {
            return this.mirrorMultiAlternateDstRegion;
        }

        public MirrorMultiAlternate setMirrorMultiAlternateNumber(Long mirrorMultiAlternateNumber) {
            this.mirrorMultiAlternateNumber = mirrorMultiAlternateNumber;
            return this;
        }
        public Long getMirrorMultiAlternateNumber() {
            return this.mirrorMultiAlternateNumber;
        }

        public MirrorMultiAlternate setMirrorMultiAlternateURL(String mirrorMultiAlternateURL) {
            this.mirrorMultiAlternateURL = mirrorMultiAlternateURL;
            return this;
        }
        public String getMirrorMultiAlternateURL() {
            return this.mirrorMultiAlternateURL;
        }

        public MirrorMultiAlternate setMirrorMultiAlternateVpcId(String mirrorMultiAlternateVpcId) {
            this.mirrorMultiAlternateVpcId = mirrorMultiAlternateVpcId;
            return this;
        }
        public String getMirrorMultiAlternateVpcId() {
            return this.mirrorMultiAlternateVpcId;
        }

    }

    public static class MirrorMultiAlternates extends TeaModel {
        @NameInMap("MirrorMultiAlternate")
        public java.util.List<MirrorMultiAlternate> mirrorMultiAlternate;

        public static MirrorMultiAlternates build(java.util.Map<String, ?> map) throws Exception {
            MirrorMultiAlternates self = new MirrorMultiAlternates();
            return TeaModel.build(map, self);
        }

        public MirrorMultiAlternates setMirrorMultiAlternate(java.util.List<MirrorMultiAlternate> mirrorMultiAlternate) {
            this.mirrorMultiAlternate = mirrorMultiAlternate;
            return this;
        }
        public java.util.List<MirrorMultiAlternate> getMirrorMultiAlternate() {
            return this.mirrorMultiAlternate;
        }

    }

    public static class ReturnHeader extends TeaModel {
        /**
         * <strong>example:</strong>
         * <p>test-header</p>
         */
        @NameInMap("Key")
        public String key;

        /**
         * <strong>example:</strong>
         * <p>test-value</p>
         */
        @NameInMap("Value")
        public String value;

        public static ReturnHeader build(java.util.Map<String, ?> map) throws Exception {
            ReturnHeader self = new ReturnHeader();
            return TeaModel.build(map, self);
        }

        public ReturnHeader setKey(String key) {
            this.key = key;
            return this;
        }
        public String getKey() {
            return this.key;
        }

        public ReturnHeader setValue(String value) {
            this.value = value;
            return this;
        }
        public String getValue() {
            return this.value;
        }

    }

    public static class MirrorReturnHeaders extends TeaModel {
        @NameInMap("ReturnHeader")
        public java.util.List<ReturnHeader> returnHeader;

        public static MirrorReturnHeaders build(java.util.Map<String, ?> map) throws Exception {
            MirrorReturnHeaders self = new MirrorReturnHeaders();
            return TeaModel.build(map, self);
        }

        public MirrorReturnHeaders setReturnHeader(java.util.List<ReturnHeader> returnHeader) {
            this.returnHeader = returnHeader;
            return this;
        }
        public java.util.List<ReturnHeader> getReturnHeader() {
            return this.returnHeader;
        }

    }

    public static class Taggings extends TeaModel {
        /**
         * <strong>example:</strong>
         * <p>test-tagging</p>
         */
        @NameInMap("Key")
        public String key;

        /**
         * <strong>example:</strong>
         * <p>test-value</p>
         */
        @NameInMap("Value")
        public String value;

        public static Taggings build(java.util.Map<String, ?> map) throws Exception {
            Taggings self = new Taggings();
            return TeaModel.build(map, self);
        }

        public Taggings setKey(String key) {
            this.key = key;
            return this;
        }
        public String getKey() {
            return this.key;
        }

        public Taggings setValue(String value) {
            this.value = value;
            return this;
        }
        public String getValue() {
            return this.value;
        }

    }

    public static class MirrorTaggings extends TeaModel {
        @NameInMap("Taggings")
        public java.util.List<Taggings> taggings;

        public static MirrorTaggings build(java.util.Map<String, ?> map) throws Exception {
            MirrorTaggings self = new MirrorTaggings();
            return TeaModel.build(map, self);
        }

        public MirrorTaggings setTaggings(java.util.List<Taggings> taggings) {
            this.taggings = taggings;
            return this;
        }
        public java.util.List<Taggings> getTaggings() {
            return this.taggings;
        }

    }

}
