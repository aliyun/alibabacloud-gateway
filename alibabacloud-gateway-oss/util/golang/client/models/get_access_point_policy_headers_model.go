// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessPointPolicyHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetAccessPointPolicyHeaders
	GetCommonHeaders() map[string]*string
	SetXOssAccessPointName(v string) *GetAccessPointPolicyHeaders
	GetXOssAccessPointName() *string
}

type GetAccessPointPolicyHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The name of the access point.
	//
	// This parameter is required.
	//
	// example:
	//
	// ap-01
	XOssAccessPointName *string `json:"x-oss-access-point-name,omitempty" xml:"x-oss-access-point-name,omitempty"`
}

func (s GetAccessPointPolicyHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetAccessPointPolicyHeaders) GoString() string {
	return s.String()
}

func (s *GetAccessPointPolicyHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetAccessPointPolicyHeaders) GetXOssAccessPointName() *string {
	return s.XOssAccessPointName
}

func (s *GetAccessPointPolicyHeaders) SetCommonHeaders(v map[string]*string) *GetAccessPointPolicyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAccessPointPolicyHeaders) SetXOssAccessPointName(v string) *GetAccessPointPolicyHeaders {
	s.XOssAccessPointName = &v
	return s
}

func (s *GetAccessPointPolicyHeaders) Validate() error {
	return dara.Validate(s)
}
