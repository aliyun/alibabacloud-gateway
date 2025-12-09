// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessPointPolicyHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteAccessPointPolicyHeaders
	GetCommonHeaders() map[string]*string
	SetXOssAccessPointName(v string) *DeleteAccessPointPolicyHeaders
	GetXOssAccessPointName() *string
}

type DeleteAccessPointPolicyHeaders struct {
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

func (s DeleteAccessPointPolicyHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessPointPolicyHeaders) GoString() string {
	return s.String()
}

func (s *DeleteAccessPointPolicyHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteAccessPointPolicyHeaders) GetXOssAccessPointName() *string {
	return s.XOssAccessPointName
}

func (s *DeleteAccessPointPolicyHeaders) SetCommonHeaders(v map[string]*string) *DeleteAccessPointPolicyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteAccessPointPolicyHeaders) SetXOssAccessPointName(v string) *DeleteAccessPointPolicyHeaders {
	s.XOssAccessPointName = &v
	return s
}

func (s *DeleteAccessPointPolicyHeaders) Validate() error {
	return dara.Validate(s)
}
