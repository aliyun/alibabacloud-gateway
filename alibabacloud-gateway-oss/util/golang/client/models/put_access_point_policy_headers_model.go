// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutAccessPointPolicyHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PutAccessPointPolicyHeaders
	GetCommonHeaders() map[string]*string
	SetXOssAccessPointName(v string) *PutAccessPointPolicyHeaders
	GetXOssAccessPointName() *string
}

type PutAccessPointPolicyHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The name of the access point.
	//
	// example:
	//
	// ap-01
	XOssAccessPointName *string `json:"x-oss-access-point-name,omitempty" xml:"x-oss-access-point-name,omitempty"`
}

func (s PutAccessPointPolicyHeaders) String() string {
	return dara.Prettify(s)
}

func (s PutAccessPointPolicyHeaders) GoString() string {
	return s.String()
}

func (s *PutAccessPointPolicyHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PutAccessPointPolicyHeaders) GetXOssAccessPointName() *string {
	return s.XOssAccessPointName
}

func (s *PutAccessPointPolicyHeaders) SetCommonHeaders(v map[string]*string) *PutAccessPointPolicyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PutAccessPointPolicyHeaders) SetXOssAccessPointName(v string) *PutAccessPointPolicyHeaders {
	s.XOssAccessPointName = &v
	return s
}

func (s *PutAccessPointPolicyHeaders) Validate() error {
	return dara.Validate(s)
}
