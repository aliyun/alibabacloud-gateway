// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRoutingRuleRedirect interface {
	dara.Model
	String() string
	GoString() string
	SetEnableReplacePrefix(v bool) *RoutingRuleRedirect
	GetEnableReplacePrefix() *bool
	SetHostName(v string) *RoutingRuleRedirect
	GetHostName() *string
	SetHttpRedirectCode(v int64) *RoutingRuleRedirect
	GetHttpRedirectCode() *int64
	SetMirrorAllowGetImageInfo(v bool) *RoutingRuleRedirect
	GetMirrorAllowGetImageInfo() *bool
	SetMirrorAllowHeadObject(v bool) *RoutingRuleRedirect
	GetMirrorAllowHeadObject() *bool
	SetMirrorAllowVideoSnapshot(v bool) *RoutingRuleRedirect
	GetMirrorAllowVideoSnapshot() *bool
	SetMirrorAsyncStatus(v int64) *RoutingRuleRedirect
	GetMirrorAsyncStatus() *int64
	SetMirrorAuth(v *RoutingRuleRedirectMirrorAuth) *RoutingRuleRedirect
	GetMirrorAuth() *RoutingRuleRedirectMirrorAuth
	SetMirrorCheckMd5(v bool) *RoutingRuleRedirect
	GetMirrorCheckMd5() *bool
	SetMirrorDstRegion(v string) *RoutingRuleRedirect
	GetMirrorDstRegion() *string
	SetMirrorDstSlaveVpcId(v string) *RoutingRuleRedirect
	GetMirrorDstSlaveVpcId() *string
	SetMirrorDstVpcId(v string) *RoutingRuleRedirect
	GetMirrorDstVpcId() *string
	SetMirrorFollowRedirect(v bool) *RoutingRuleRedirect
	GetMirrorFollowRedirect() *bool
	SetMirrorHeaders(v *RoutingRuleRedirectMirrorHeaders) *RoutingRuleRedirect
	GetMirrorHeaders() *RoutingRuleRedirectMirrorHeaders
	SetMirrorIsExpressTunnel(v bool) *RoutingRuleRedirect
	GetMirrorIsExpressTunnel() *bool
	SetMirrorMultiAlternates(v *RoutingRuleRedirectMirrorMultiAlternates) *RoutingRuleRedirect
	GetMirrorMultiAlternates() *RoutingRuleRedirectMirrorMultiAlternates
	SetMirrorPassOriginalSlashes(v bool) *RoutingRuleRedirect
	GetMirrorPassOriginalSlashes() *bool
	SetMirrorPassQueryString(v bool) *RoutingRuleRedirect
	GetMirrorPassQueryString() *bool
	SetMirrorProxyPass(v bool) *RoutingRuleRedirect
	GetMirrorProxyPass() *bool
	SetMirrorReturnHeaders(v *RoutingRuleRedirectMirrorReturnHeaders) *RoutingRuleRedirect
	GetMirrorReturnHeaders() *RoutingRuleRedirectMirrorReturnHeaders
	SetMirrorRole(v string) *RoutingRuleRedirect
	GetMirrorRole() *string
	SetMirrorSNI(v bool) *RoutingRuleRedirect
	GetMirrorSNI() *bool
	SetMirrorSaveOssMeta(v bool) *RoutingRuleRedirect
	GetMirrorSaveOssMeta() *bool
	SetMirrorSwitchAllErrors(v bool) *RoutingRuleRedirect
	GetMirrorSwitchAllErrors() *bool
	SetMirrorTaggings(v *RoutingRuleRedirectMirrorTaggings) *RoutingRuleRedirect
	GetMirrorTaggings() *RoutingRuleRedirectMirrorTaggings
	SetMirrorTunnelId(v string) *RoutingRuleRedirect
	GetMirrorTunnelId() *string
	SetMirrorURL(v string) *RoutingRuleRedirect
	GetMirrorURL() *string
	SetMirrorURLProbe(v string) *RoutingRuleRedirect
	GetMirrorURLProbe() *string
	SetMirrorURLSlave(v string) *RoutingRuleRedirect
	GetMirrorURLSlave() *string
	SetMirrorUserLastModified(v bool) *RoutingRuleRedirect
	GetMirrorUserLastModified() *bool
	SetMirrorUsingRole(v bool) *RoutingRuleRedirect
	GetMirrorUsingRole() *bool
	SetPassQueryString(v bool) *RoutingRuleRedirect
	GetPassQueryString() *bool
	SetProtocol(v string) *RoutingRuleRedirect
	GetProtocol() *string
	SetRedirectType(v string) *RoutingRuleRedirect
	GetRedirectType() *string
	SetReplaceKeyPrefixWith(v string) *RoutingRuleRedirect
	GetReplaceKeyPrefixWith() *string
	SetReplaceKeyWith(v string) *RoutingRuleRedirect
	GetReplaceKeyWith() *string
	SetTransparentMirrorResponseCodes(v string) *RoutingRuleRedirect
	GetTransparentMirrorResponseCodes() *string
}

type RoutingRuleRedirect struct {
	// example:
	//
	// true
	EnableReplacePrefix *bool `json:"EnableReplacePrefix,omitempty" xml:"EnableReplacePrefix,omitempty"`
	// example:
	//
	// example.com
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// example:
	//
	// 302
	HttpRedirectCode *int64 `json:"HttpRedirectCode,omitempty" xml:"HttpRedirectCode,omitempty"`
	// example:
	//
	// true
	MirrorAllowGetImageInfo *bool `json:"MirrorAllowGetImageInfo,omitempty" xml:"MirrorAllowGetImageInfo,omitempty"`
	// example:
	//
	// true
	MirrorAllowHeadObject *bool `json:"MirrorAllowHeadObject,omitempty" xml:"MirrorAllowHeadObject,omitempty"`
	// example:
	//
	// false
	MirrorAllowVideoSnapshot *bool `json:"MirrorAllowVideoSnapshot,omitempty" xml:"MirrorAllowVideoSnapshot,omitempty"`
	// example:
	//
	// 303
	MirrorAsyncStatus *int64                         `json:"MirrorAsyncStatus,omitempty" xml:"MirrorAsyncStatus,omitempty"`
	MirrorAuth        *RoutingRuleRedirectMirrorAuth `json:"MirrorAuth,omitempty" xml:"MirrorAuth,omitempty" type:"Struct"`
	// example:
	//
	// false
	MirrorCheckMd5 *bool `json:"MirrorCheckMd5,omitempty" xml:"MirrorCheckMd5,omitempty"`
	// example:
	//
	// cn-hangzhou
	MirrorDstRegion *string `json:"MirrorDstRegion,omitempty" xml:"MirrorDstRegion,omitempty"`
	// example:
	//
	// vpc-test-id
	MirrorDstSlaveVpcId *string `json:"MirrorDstSlaveVpcId,omitempty" xml:"MirrorDstSlaveVpcId,omitempty"`
	// example:
	//
	// vpc-test-id
	MirrorDstVpcId *string `json:"MirrorDstVpcId,omitempty" xml:"MirrorDstVpcId,omitempty"`
	// example:
	//
	// true
	MirrorFollowRedirect *bool                             `json:"MirrorFollowRedirect,omitempty" xml:"MirrorFollowRedirect,omitempty"`
	MirrorHeaders        *RoutingRuleRedirectMirrorHeaders `json:"MirrorHeaders,omitempty" xml:"MirrorHeaders,omitempty" type:"Struct"`
	// example:
	//
	// true
	MirrorIsExpressTunnel *bool                                     `json:"MirrorIsExpressTunnel,omitempty" xml:"MirrorIsExpressTunnel,omitempty"`
	MirrorMultiAlternates *RoutingRuleRedirectMirrorMultiAlternates `json:"MirrorMultiAlternates,omitempty" xml:"MirrorMultiAlternates,omitempty" type:"Struct"`
	// example:
	//
	// false
	MirrorPassOriginalSlashes *bool `json:"MirrorPassOriginalSlashes,omitempty" xml:"MirrorPassOriginalSlashes,omitempty"`
	// example:
	//
	// true
	MirrorPassQueryString *bool `json:"MirrorPassQueryString,omitempty" xml:"MirrorPassQueryString,omitempty"`
	// example:
	//
	// false
	MirrorProxyPass     *bool                                   `json:"MirrorProxyPass,omitempty" xml:"MirrorProxyPass,omitempty"`
	MirrorReturnHeaders *RoutingRuleRedirectMirrorReturnHeaders `json:"MirrorReturnHeaders,omitempty" xml:"MirrorReturnHeaders,omitempty" type:"Struct"`
	// example:
	//
	// aliyun-test-role
	MirrorRole *string `json:"MirrorRole,omitempty" xml:"MirrorRole,omitempty"`
	// example:
	//
	// true
	MirrorSNI *bool `json:"MirrorSNI,omitempty" xml:"MirrorSNI,omitempty"`
	// example:
	//
	// true
	MirrorSaveOssMeta *bool `json:"MirrorSaveOssMeta,omitempty" xml:"MirrorSaveOssMeta,omitempty"`
	// example:
	//
	// false
	MirrorSwitchAllErrors *bool                              `json:"MirrorSwitchAllErrors,omitempty" xml:"MirrorSwitchAllErrors,omitempty"`
	MirrorTaggings        *RoutingRuleRedirectMirrorTaggings `json:"MirrorTaggings,omitempty" xml:"MirrorTaggings,omitempty" type:"Struct"`
	// example:
	//
	// test-tunnel-id
	MirrorTunnelId *string `json:"MirrorTunnelId,omitempty" xml:"MirrorTunnelId,omitempty"`
	// example:
	//
	// http://example.com/
	MirrorURL *string `json:"MirrorURL,omitempty" xml:"MirrorURL,omitempty"`
	// example:
	//
	// https://example.com/hartbeat
	MirrorURLProbe *string `json:"MirrorURLProbe,omitempty" xml:"MirrorURLProbe,omitempty"`
	// example:
	//
	// https://example.com
	MirrorURLSlave *string `json:"MirrorURLSlave,omitempty" xml:"MirrorURLSlave,omitempty"`
	// example:
	//
	// false
	MirrorUserLastModified *bool `json:"MirrorUserLastModified,omitempty" xml:"MirrorUserLastModified,omitempty"`
	// example:
	//
	// false
	MirrorUsingRole *bool `json:"MirrorUsingRole,omitempty" xml:"MirrorUsingRole,omitempty"`
	// example:
	//
	// true
	PassQueryString *bool `json:"PassQueryString,omitempty" xml:"PassQueryString,omitempty"`
	// example:
	//
	// https
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// Mirror
	RedirectType *string `json:"RedirectType,omitempty" xml:"RedirectType,omitempty"`
	// example:
	//
	// def/
	ReplaceKeyPrefixWith *string `json:"ReplaceKeyPrefixWith,omitempty" xml:"ReplaceKeyPrefixWith,omitempty"`
	// example:
	//
	// prefix/${key}.suffix
	ReplaceKeyWith *string `json:"ReplaceKeyWith,omitempty" xml:"ReplaceKeyWith,omitempty"`
	// example:
	//
	// false
	TransparentMirrorResponseCodes *string `json:"TransparentMirrorResponseCodes,omitempty" xml:"TransparentMirrorResponseCodes,omitempty"`
}

func (s RoutingRuleRedirect) String() string {
	return dara.Prettify(s)
}

func (s RoutingRuleRedirect) GoString() string {
	return s.String()
}

func (s *RoutingRuleRedirect) GetEnableReplacePrefix() *bool {
	return s.EnableReplacePrefix
}

func (s *RoutingRuleRedirect) GetHostName() *string {
	return s.HostName
}

func (s *RoutingRuleRedirect) GetHttpRedirectCode() *int64 {
	return s.HttpRedirectCode
}

func (s *RoutingRuleRedirect) GetMirrorAllowGetImageInfo() *bool {
	return s.MirrorAllowGetImageInfo
}

func (s *RoutingRuleRedirect) GetMirrorAllowHeadObject() *bool {
	return s.MirrorAllowHeadObject
}

func (s *RoutingRuleRedirect) GetMirrorAllowVideoSnapshot() *bool {
	return s.MirrorAllowVideoSnapshot
}

func (s *RoutingRuleRedirect) GetMirrorAsyncStatus() *int64 {
	return s.MirrorAsyncStatus
}

func (s *RoutingRuleRedirect) GetMirrorAuth() *RoutingRuleRedirectMirrorAuth {
	return s.MirrorAuth
}

func (s *RoutingRuleRedirect) GetMirrorCheckMd5() *bool {
	return s.MirrorCheckMd5
}

func (s *RoutingRuleRedirect) GetMirrorDstRegion() *string {
	return s.MirrorDstRegion
}

func (s *RoutingRuleRedirect) GetMirrorDstSlaveVpcId() *string {
	return s.MirrorDstSlaveVpcId
}

func (s *RoutingRuleRedirect) GetMirrorDstVpcId() *string {
	return s.MirrorDstVpcId
}

func (s *RoutingRuleRedirect) GetMirrorFollowRedirect() *bool {
	return s.MirrorFollowRedirect
}

func (s *RoutingRuleRedirect) GetMirrorHeaders() *RoutingRuleRedirectMirrorHeaders {
	return s.MirrorHeaders
}

func (s *RoutingRuleRedirect) GetMirrorIsExpressTunnel() *bool {
	return s.MirrorIsExpressTunnel
}

func (s *RoutingRuleRedirect) GetMirrorMultiAlternates() *RoutingRuleRedirectMirrorMultiAlternates {
	return s.MirrorMultiAlternates
}

func (s *RoutingRuleRedirect) GetMirrorPassOriginalSlashes() *bool {
	return s.MirrorPassOriginalSlashes
}

func (s *RoutingRuleRedirect) GetMirrorPassQueryString() *bool {
	return s.MirrorPassQueryString
}

func (s *RoutingRuleRedirect) GetMirrorProxyPass() *bool {
	return s.MirrorProxyPass
}

func (s *RoutingRuleRedirect) GetMirrorReturnHeaders() *RoutingRuleRedirectMirrorReturnHeaders {
	return s.MirrorReturnHeaders
}

func (s *RoutingRuleRedirect) GetMirrorRole() *string {
	return s.MirrorRole
}

func (s *RoutingRuleRedirect) GetMirrorSNI() *bool {
	return s.MirrorSNI
}

func (s *RoutingRuleRedirect) GetMirrorSaveOssMeta() *bool {
	return s.MirrorSaveOssMeta
}

func (s *RoutingRuleRedirect) GetMirrorSwitchAllErrors() *bool {
	return s.MirrorSwitchAllErrors
}

func (s *RoutingRuleRedirect) GetMirrorTaggings() *RoutingRuleRedirectMirrorTaggings {
	return s.MirrorTaggings
}

func (s *RoutingRuleRedirect) GetMirrorTunnelId() *string {
	return s.MirrorTunnelId
}

func (s *RoutingRuleRedirect) GetMirrorURL() *string {
	return s.MirrorURL
}

func (s *RoutingRuleRedirect) GetMirrorURLProbe() *string {
	return s.MirrorURLProbe
}

func (s *RoutingRuleRedirect) GetMirrorURLSlave() *string {
	return s.MirrorURLSlave
}

func (s *RoutingRuleRedirect) GetMirrorUserLastModified() *bool {
	return s.MirrorUserLastModified
}

func (s *RoutingRuleRedirect) GetMirrorUsingRole() *bool {
	return s.MirrorUsingRole
}

func (s *RoutingRuleRedirect) GetPassQueryString() *bool {
	return s.PassQueryString
}

func (s *RoutingRuleRedirect) GetProtocol() *string {
	return s.Protocol
}

func (s *RoutingRuleRedirect) GetRedirectType() *string {
	return s.RedirectType
}

func (s *RoutingRuleRedirect) GetReplaceKeyPrefixWith() *string {
	return s.ReplaceKeyPrefixWith
}

func (s *RoutingRuleRedirect) GetReplaceKeyWith() *string {
	return s.ReplaceKeyWith
}

func (s *RoutingRuleRedirect) GetTransparentMirrorResponseCodes() *string {
	return s.TransparentMirrorResponseCodes
}

func (s *RoutingRuleRedirect) SetEnableReplacePrefix(v bool) *RoutingRuleRedirect {
	s.EnableReplacePrefix = &v
	return s
}

func (s *RoutingRuleRedirect) SetHostName(v string) *RoutingRuleRedirect {
	s.HostName = &v
	return s
}

func (s *RoutingRuleRedirect) SetHttpRedirectCode(v int64) *RoutingRuleRedirect {
	s.HttpRedirectCode = &v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorAllowGetImageInfo(v bool) *RoutingRuleRedirect {
	s.MirrorAllowGetImageInfo = &v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorAllowHeadObject(v bool) *RoutingRuleRedirect {
	s.MirrorAllowHeadObject = &v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorAllowVideoSnapshot(v bool) *RoutingRuleRedirect {
	s.MirrorAllowVideoSnapshot = &v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorAsyncStatus(v int64) *RoutingRuleRedirect {
	s.MirrorAsyncStatus = &v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorAuth(v *RoutingRuleRedirectMirrorAuth) *RoutingRuleRedirect {
	s.MirrorAuth = v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorCheckMd5(v bool) *RoutingRuleRedirect {
	s.MirrorCheckMd5 = &v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorDstRegion(v string) *RoutingRuleRedirect {
	s.MirrorDstRegion = &v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorDstSlaveVpcId(v string) *RoutingRuleRedirect {
	s.MirrorDstSlaveVpcId = &v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorDstVpcId(v string) *RoutingRuleRedirect {
	s.MirrorDstVpcId = &v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorFollowRedirect(v bool) *RoutingRuleRedirect {
	s.MirrorFollowRedirect = &v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorHeaders(v *RoutingRuleRedirectMirrorHeaders) *RoutingRuleRedirect {
	s.MirrorHeaders = v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorIsExpressTunnel(v bool) *RoutingRuleRedirect {
	s.MirrorIsExpressTunnel = &v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorMultiAlternates(v *RoutingRuleRedirectMirrorMultiAlternates) *RoutingRuleRedirect {
	s.MirrorMultiAlternates = v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorPassOriginalSlashes(v bool) *RoutingRuleRedirect {
	s.MirrorPassOriginalSlashes = &v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorPassQueryString(v bool) *RoutingRuleRedirect {
	s.MirrorPassQueryString = &v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorProxyPass(v bool) *RoutingRuleRedirect {
	s.MirrorProxyPass = &v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorReturnHeaders(v *RoutingRuleRedirectMirrorReturnHeaders) *RoutingRuleRedirect {
	s.MirrorReturnHeaders = v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorRole(v string) *RoutingRuleRedirect {
	s.MirrorRole = &v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorSNI(v bool) *RoutingRuleRedirect {
	s.MirrorSNI = &v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorSaveOssMeta(v bool) *RoutingRuleRedirect {
	s.MirrorSaveOssMeta = &v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorSwitchAllErrors(v bool) *RoutingRuleRedirect {
	s.MirrorSwitchAllErrors = &v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorTaggings(v *RoutingRuleRedirectMirrorTaggings) *RoutingRuleRedirect {
	s.MirrorTaggings = v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorTunnelId(v string) *RoutingRuleRedirect {
	s.MirrorTunnelId = &v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorURL(v string) *RoutingRuleRedirect {
	s.MirrorURL = &v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorURLProbe(v string) *RoutingRuleRedirect {
	s.MirrorURLProbe = &v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorURLSlave(v string) *RoutingRuleRedirect {
	s.MirrorURLSlave = &v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorUserLastModified(v bool) *RoutingRuleRedirect {
	s.MirrorUserLastModified = &v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorUsingRole(v bool) *RoutingRuleRedirect {
	s.MirrorUsingRole = &v
	return s
}

func (s *RoutingRuleRedirect) SetPassQueryString(v bool) *RoutingRuleRedirect {
	s.PassQueryString = &v
	return s
}

func (s *RoutingRuleRedirect) SetProtocol(v string) *RoutingRuleRedirect {
	s.Protocol = &v
	return s
}

func (s *RoutingRuleRedirect) SetRedirectType(v string) *RoutingRuleRedirect {
	s.RedirectType = &v
	return s
}

func (s *RoutingRuleRedirect) SetReplaceKeyPrefixWith(v string) *RoutingRuleRedirect {
	s.ReplaceKeyPrefixWith = &v
	return s
}

func (s *RoutingRuleRedirect) SetReplaceKeyWith(v string) *RoutingRuleRedirect {
	s.ReplaceKeyWith = &v
	return s
}

func (s *RoutingRuleRedirect) SetTransparentMirrorResponseCodes(v string) *RoutingRuleRedirect {
	s.TransparentMirrorResponseCodes = &v
	return s
}

func (s *RoutingRuleRedirect) Validate() error {
	if s.MirrorAuth != nil {
		if err := s.MirrorAuth.Validate(); err != nil {
			return err
		}
	}
	if s.MirrorHeaders != nil {
		if err := s.MirrorHeaders.Validate(); err != nil {
			return err
		}
	}
	if s.MirrorMultiAlternates != nil {
		if err := s.MirrorMultiAlternates.Validate(); err != nil {
			return err
		}
	}
	if s.MirrorReturnHeaders != nil {
		if err := s.MirrorReturnHeaders.Validate(); err != nil {
			return err
		}
	}
	if s.MirrorTaggings != nil {
		if err := s.MirrorTaggings.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RoutingRuleRedirectMirrorAuth struct {
	// example:
	//
	// TESTAK
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// example:
	//
	// TESTSK
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// example:
	//
	// S3V4
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// example:
	//
	// ap-southeast-1
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s RoutingRuleRedirectMirrorAuth) String() string {
	return dara.Prettify(s)
}

func (s RoutingRuleRedirectMirrorAuth) GoString() string {
	return s.String()
}

func (s *RoutingRuleRedirectMirrorAuth) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *RoutingRuleRedirectMirrorAuth) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *RoutingRuleRedirectMirrorAuth) GetAuthType() *string {
	return s.AuthType
}

func (s *RoutingRuleRedirectMirrorAuth) GetRegion() *string {
	return s.Region
}

func (s *RoutingRuleRedirectMirrorAuth) SetAccessKeyId(v string) *RoutingRuleRedirectMirrorAuth {
	s.AccessKeyId = &v
	return s
}

func (s *RoutingRuleRedirectMirrorAuth) SetAccessKeySecret(v string) *RoutingRuleRedirectMirrorAuth {
	s.AccessKeySecret = &v
	return s
}

func (s *RoutingRuleRedirectMirrorAuth) SetAuthType(v string) *RoutingRuleRedirectMirrorAuth {
	s.AuthType = &v
	return s
}

func (s *RoutingRuleRedirectMirrorAuth) SetRegion(v string) *RoutingRuleRedirectMirrorAuth {
	s.Region = &v
	return s
}

func (s *RoutingRuleRedirectMirrorAuth) Validate() error {
	return dara.Validate(s)
}

type RoutingRuleRedirectMirrorHeaders struct {
	Pass []*string `json:"Pass,omitempty" xml:"Pass,omitempty" type:"Repeated"`
	// example:
	//
	// false
	PassAll *bool                                  `json:"PassAll,omitempty" xml:"PassAll,omitempty"`
	Remove  []*string                              `json:"Remove,omitempty" xml:"Remove,omitempty" type:"Repeated"`
	Set     []*RoutingRuleRedirectMirrorHeadersSet `json:"Set,omitempty" xml:"Set,omitempty" type:"Repeated"`
}

func (s RoutingRuleRedirectMirrorHeaders) String() string {
	return dara.Prettify(s)
}

func (s RoutingRuleRedirectMirrorHeaders) GoString() string {
	return s.String()
}

func (s *RoutingRuleRedirectMirrorHeaders) GetPass() []*string {
	return s.Pass
}

func (s *RoutingRuleRedirectMirrorHeaders) GetPassAll() *bool {
	return s.PassAll
}

func (s *RoutingRuleRedirectMirrorHeaders) GetRemove() []*string {
	return s.Remove
}

func (s *RoutingRuleRedirectMirrorHeaders) GetSet() []*RoutingRuleRedirectMirrorHeadersSet {
	return s.Set
}

func (s *RoutingRuleRedirectMirrorHeaders) SetPass(v []*string) *RoutingRuleRedirectMirrorHeaders {
	s.Pass = v
	return s
}

func (s *RoutingRuleRedirectMirrorHeaders) SetPassAll(v bool) *RoutingRuleRedirectMirrorHeaders {
	s.PassAll = &v
	return s
}

func (s *RoutingRuleRedirectMirrorHeaders) SetRemove(v []*string) *RoutingRuleRedirectMirrorHeaders {
	s.Remove = v
	return s
}

func (s *RoutingRuleRedirectMirrorHeaders) SetSet(v []*RoutingRuleRedirectMirrorHeadersSet) *RoutingRuleRedirectMirrorHeaders {
	s.Set = v
	return s
}

func (s *RoutingRuleRedirectMirrorHeaders) Validate() error {
	if s.Set != nil {
		for _, item := range s.Set {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RoutingRuleRedirectMirrorHeadersSet struct {
	// example:
	//
	// set-header
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// set-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RoutingRuleRedirectMirrorHeadersSet) String() string {
	return dara.Prettify(s)
}

func (s RoutingRuleRedirectMirrorHeadersSet) GoString() string {
	return s.String()
}

func (s *RoutingRuleRedirectMirrorHeadersSet) GetKey() *string {
	return s.Key
}

func (s *RoutingRuleRedirectMirrorHeadersSet) GetValue() *string {
	return s.Value
}

func (s *RoutingRuleRedirectMirrorHeadersSet) SetKey(v string) *RoutingRuleRedirectMirrorHeadersSet {
	s.Key = &v
	return s
}

func (s *RoutingRuleRedirectMirrorHeadersSet) SetValue(v string) *RoutingRuleRedirectMirrorHeadersSet {
	s.Value = &v
	return s
}

func (s *RoutingRuleRedirectMirrorHeadersSet) Validate() error {
	return dara.Validate(s)
}

type RoutingRuleRedirectMirrorMultiAlternates struct {
	MirrorMultiAlternate []*RoutingRuleRedirectMirrorMultiAlternatesMirrorMultiAlternate `json:"MirrorMultiAlternate,omitempty" xml:"MirrorMultiAlternate,omitempty" type:"Repeated"`
}

func (s RoutingRuleRedirectMirrorMultiAlternates) String() string {
	return dara.Prettify(s)
}

func (s RoutingRuleRedirectMirrorMultiAlternates) GoString() string {
	return s.String()
}

func (s *RoutingRuleRedirectMirrorMultiAlternates) GetMirrorMultiAlternate() []*RoutingRuleRedirectMirrorMultiAlternatesMirrorMultiAlternate {
	return s.MirrorMultiAlternate
}

func (s *RoutingRuleRedirectMirrorMultiAlternates) SetMirrorMultiAlternate(v []*RoutingRuleRedirectMirrorMultiAlternatesMirrorMultiAlternate) *RoutingRuleRedirectMirrorMultiAlternates {
	s.MirrorMultiAlternate = v
	return s
}

func (s *RoutingRuleRedirectMirrorMultiAlternates) Validate() error {
	if s.MirrorMultiAlternate != nil {
		for _, item := range s.MirrorMultiAlternate {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RoutingRuleRedirectMirrorMultiAlternatesMirrorMultiAlternate struct {
	// example:
	//
	// ap-southeast-1
	MirrorMultiAlternateDstRegion *string `json:"MirrorMultiAlternateDstRegion,omitempty" xml:"MirrorMultiAlternateDstRegion,omitempty"`
	// example:
	//
	// 32
	MirrorMultiAlternateNumber *int64 `json:"MirrorMultiAlternateNumber,omitempty" xml:"MirrorMultiAlternateNumber,omitempty"`
	// example:
	//
	// https://test-multi-alter.example.com
	MirrorMultiAlternateURL *string `json:"MirrorMultiAlternateURL,omitempty" xml:"MirrorMultiAlternateURL,omitempty"`
	// example:
	//
	// vpc-test-id
	MirrorMultiAlternateVpcId *string `json:"MirrorMultiAlternateVpcId,omitempty" xml:"MirrorMultiAlternateVpcId,omitempty"`
}

func (s RoutingRuleRedirectMirrorMultiAlternatesMirrorMultiAlternate) String() string {
	return dara.Prettify(s)
}

func (s RoutingRuleRedirectMirrorMultiAlternatesMirrorMultiAlternate) GoString() string {
	return s.String()
}

func (s *RoutingRuleRedirectMirrorMultiAlternatesMirrorMultiAlternate) GetMirrorMultiAlternateDstRegion() *string {
	return s.MirrorMultiAlternateDstRegion
}

func (s *RoutingRuleRedirectMirrorMultiAlternatesMirrorMultiAlternate) GetMirrorMultiAlternateNumber() *int64 {
	return s.MirrorMultiAlternateNumber
}

func (s *RoutingRuleRedirectMirrorMultiAlternatesMirrorMultiAlternate) GetMirrorMultiAlternateURL() *string {
	return s.MirrorMultiAlternateURL
}

func (s *RoutingRuleRedirectMirrorMultiAlternatesMirrorMultiAlternate) GetMirrorMultiAlternateVpcId() *string {
	return s.MirrorMultiAlternateVpcId
}

func (s *RoutingRuleRedirectMirrorMultiAlternatesMirrorMultiAlternate) SetMirrorMultiAlternateDstRegion(v string) *RoutingRuleRedirectMirrorMultiAlternatesMirrorMultiAlternate {
	s.MirrorMultiAlternateDstRegion = &v
	return s
}

func (s *RoutingRuleRedirectMirrorMultiAlternatesMirrorMultiAlternate) SetMirrorMultiAlternateNumber(v int64) *RoutingRuleRedirectMirrorMultiAlternatesMirrorMultiAlternate {
	s.MirrorMultiAlternateNumber = &v
	return s
}

func (s *RoutingRuleRedirectMirrorMultiAlternatesMirrorMultiAlternate) SetMirrorMultiAlternateURL(v string) *RoutingRuleRedirectMirrorMultiAlternatesMirrorMultiAlternate {
	s.MirrorMultiAlternateURL = &v
	return s
}

func (s *RoutingRuleRedirectMirrorMultiAlternatesMirrorMultiAlternate) SetMirrorMultiAlternateVpcId(v string) *RoutingRuleRedirectMirrorMultiAlternatesMirrorMultiAlternate {
	s.MirrorMultiAlternateVpcId = &v
	return s
}

func (s *RoutingRuleRedirectMirrorMultiAlternatesMirrorMultiAlternate) Validate() error {
	return dara.Validate(s)
}

type RoutingRuleRedirectMirrorReturnHeaders struct {
	ReturnHeader []*RoutingRuleRedirectMirrorReturnHeadersReturnHeader `json:"ReturnHeader,omitempty" xml:"ReturnHeader,omitempty" type:"Repeated"`
}

func (s RoutingRuleRedirectMirrorReturnHeaders) String() string {
	return dara.Prettify(s)
}

func (s RoutingRuleRedirectMirrorReturnHeaders) GoString() string {
	return s.String()
}

func (s *RoutingRuleRedirectMirrorReturnHeaders) GetReturnHeader() []*RoutingRuleRedirectMirrorReturnHeadersReturnHeader {
	return s.ReturnHeader
}

func (s *RoutingRuleRedirectMirrorReturnHeaders) SetReturnHeader(v []*RoutingRuleRedirectMirrorReturnHeadersReturnHeader) *RoutingRuleRedirectMirrorReturnHeaders {
	s.ReturnHeader = v
	return s
}

func (s *RoutingRuleRedirectMirrorReturnHeaders) Validate() error {
	if s.ReturnHeader != nil {
		for _, item := range s.ReturnHeader {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RoutingRuleRedirectMirrorReturnHeadersReturnHeader struct {
	// example:
	//
	// test-header
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RoutingRuleRedirectMirrorReturnHeadersReturnHeader) String() string {
	return dara.Prettify(s)
}

func (s RoutingRuleRedirectMirrorReturnHeadersReturnHeader) GoString() string {
	return s.String()
}

func (s *RoutingRuleRedirectMirrorReturnHeadersReturnHeader) GetKey() *string {
	return s.Key
}

func (s *RoutingRuleRedirectMirrorReturnHeadersReturnHeader) GetValue() *string {
	return s.Value
}

func (s *RoutingRuleRedirectMirrorReturnHeadersReturnHeader) SetKey(v string) *RoutingRuleRedirectMirrorReturnHeadersReturnHeader {
	s.Key = &v
	return s
}

func (s *RoutingRuleRedirectMirrorReturnHeadersReturnHeader) SetValue(v string) *RoutingRuleRedirectMirrorReturnHeadersReturnHeader {
	s.Value = &v
	return s
}

func (s *RoutingRuleRedirectMirrorReturnHeadersReturnHeader) Validate() error {
	return dara.Validate(s)
}

type RoutingRuleRedirectMirrorTaggings struct {
	Taggings []*RoutingRuleRedirectMirrorTaggingsTaggings `json:"Taggings,omitempty" xml:"Taggings,omitempty" type:"Repeated"`
}

func (s RoutingRuleRedirectMirrorTaggings) String() string {
	return dara.Prettify(s)
}

func (s RoutingRuleRedirectMirrorTaggings) GoString() string {
	return s.String()
}

func (s *RoutingRuleRedirectMirrorTaggings) GetTaggings() []*RoutingRuleRedirectMirrorTaggingsTaggings {
	return s.Taggings
}

func (s *RoutingRuleRedirectMirrorTaggings) SetTaggings(v []*RoutingRuleRedirectMirrorTaggingsTaggings) *RoutingRuleRedirectMirrorTaggings {
	s.Taggings = v
	return s
}

func (s *RoutingRuleRedirectMirrorTaggings) Validate() error {
	if s.Taggings != nil {
		for _, item := range s.Taggings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RoutingRuleRedirectMirrorTaggingsTaggings struct {
	// example:
	//
	// test-tagging
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RoutingRuleRedirectMirrorTaggingsTaggings) String() string {
	return dara.Prettify(s)
}

func (s RoutingRuleRedirectMirrorTaggingsTaggings) GoString() string {
	return s.String()
}

func (s *RoutingRuleRedirectMirrorTaggingsTaggings) GetKey() *string {
	return s.Key
}

func (s *RoutingRuleRedirectMirrorTaggingsTaggings) GetValue() *string {
	return s.Value
}

func (s *RoutingRuleRedirectMirrorTaggingsTaggings) SetKey(v string) *RoutingRuleRedirectMirrorTaggingsTaggings {
	s.Key = &v
	return s
}

func (s *RoutingRuleRedirectMirrorTaggingsTaggings) SetValue(v string) *RoutingRuleRedirectMirrorTaggingsTaggings {
	s.Value = &v
	return s
}

func (s *RoutingRuleRedirectMirrorTaggingsTaggings) Validate() error {
	return dara.Validate(s)
}
