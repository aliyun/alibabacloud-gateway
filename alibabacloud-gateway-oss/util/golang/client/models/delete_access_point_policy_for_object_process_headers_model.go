// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessPointPolicyForObjectProcessHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteAccessPointPolicyForObjectProcessHeaders
	GetCommonHeaders() map[string]*string
	SetXOssAccessPointForObjectProcessName(v string) *DeleteAccessPointPolicyForObjectProcessHeaders
	GetXOssAccessPointForObjectProcessName() *string
}

type DeleteAccessPointPolicyForObjectProcessHeaders struct {
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

func (s DeleteAccessPointPolicyForObjectProcessHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessPointPolicyForObjectProcessHeaders) GoString() string {
	return s.String()
}

func (s *DeleteAccessPointPolicyForObjectProcessHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteAccessPointPolicyForObjectProcessHeaders) GetXOssAccessPointForObjectProcessName() *string {
	return s.XOssAccessPointForObjectProcessName
}

func (s *DeleteAccessPointPolicyForObjectProcessHeaders) SetCommonHeaders(v map[string]*string) *DeleteAccessPointPolicyForObjectProcessHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteAccessPointPolicyForObjectProcessHeaders) SetXOssAccessPointForObjectProcessName(v string) *DeleteAccessPointPolicyForObjectProcessHeaders {
	s.XOssAccessPointForObjectProcessName = &v
	return s
}

func (s *DeleteAccessPointPolicyForObjectProcessHeaders) Validate() error {
	return dara.Validate(s)
}
