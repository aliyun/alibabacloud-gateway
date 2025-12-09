// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutAccessPointPolicyForObjectProcessHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PutAccessPointPolicyForObjectProcessHeaders
	GetCommonHeaders() map[string]*string
	SetXOssAccessPointForObjectProcessName(v string) *PutAccessPointPolicyForObjectProcessHeaders
	GetXOssAccessPointForObjectProcessName() *string
}

type PutAccessPointPolicyForObjectProcessHeaders struct {
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

func (s PutAccessPointPolicyForObjectProcessHeaders) String() string {
	return dara.Prettify(s)
}

func (s PutAccessPointPolicyForObjectProcessHeaders) GoString() string {
	return s.String()
}

func (s *PutAccessPointPolicyForObjectProcessHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PutAccessPointPolicyForObjectProcessHeaders) GetXOssAccessPointForObjectProcessName() *string {
	return s.XOssAccessPointForObjectProcessName
}

func (s *PutAccessPointPolicyForObjectProcessHeaders) SetCommonHeaders(v map[string]*string) *PutAccessPointPolicyForObjectProcessHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PutAccessPointPolicyForObjectProcessHeaders) SetXOssAccessPointForObjectProcessName(v string) *PutAccessPointPolicyForObjectProcessHeaders {
	s.XOssAccessPointForObjectProcessName = &v
	return s
}

func (s *PutAccessPointPolicyForObjectProcessHeaders) Validate() error {
	return dara.Validate(s)
}
