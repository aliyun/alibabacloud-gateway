// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessPointPolicyForObjectProcessHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetAccessPointPolicyForObjectProcessHeaders
	GetCommonHeaders() map[string]*string
	SetXOssAccessPointForObjectProcessName(v string) *GetAccessPointPolicyForObjectProcessHeaders
	GetXOssAccessPointForObjectProcessName() *string
}

type GetAccessPointPolicyForObjectProcessHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The name of the Object FC Access Point.
	//
	// This parameter is required.
	//
	// example:
	//
	// fc-ap-01
	XOssAccessPointForObjectProcessName *string `json:"x-oss-access-point-for-object-process-name,omitempty" xml:"x-oss-access-point-for-object-process-name,omitempty"`
}

func (s GetAccessPointPolicyForObjectProcessHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetAccessPointPolicyForObjectProcessHeaders) GoString() string {
	return s.String()
}

func (s *GetAccessPointPolicyForObjectProcessHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetAccessPointPolicyForObjectProcessHeaders) GetXOssAccessPointForObjectProcessName() *string {
	return s.XOssAccessPointForObjectProcessName
}

func (s *GetAccessPointPolicyForObjectProcessHeaders) SetCommonHeaders(v map[string]*string) *GetAccessPointPolicyForObjectProcessHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAccessPointPolicyForObjectProcessHeaders) SetXOssAccessPointForObjectProcessName(v string) *GetAccessPointPolicyForObjectProcessHeaders {
	s.XOssAccessPointForObjectProcessName = &v
	return s
}

func (s *GetAccessPointPolicyForObjectProcessHeaders) Validate() error {
	return dara.Validate(s)
}
