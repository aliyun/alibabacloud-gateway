// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessPointForObjectProcessHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateAccessPointForObjectProcessHeaders
	GetCommonHeaders() map[string]*string
	SetXOssAccessPointForObjectProcessName(v string) *CreateAccessPointForObjectProcessHeaders
	GetXOssAccessPointForObjectProcessName() *string
}

type CreateAccessPointForObjectProcessHeaders struct {
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

func (s CreateAccessPointForObjectProcessHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessPointForObjectProcessHeaders) GoString() string {
	return s.String()
}

func (s *CreateAccessPointForObjectProcessHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateAccessPointForObjectProcessHeaders) GetXOssAccessPointForObjectProcessName() *string {
	return s.XOssAccessPointForObjectProcessName
}

func (s *CreateAccessPointForObjectProcessHeaders) SetCommonHeaders(v map[string]*string) *CreateAccessPointForObjectProcessHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateAccessPointForObjectProcessHeaders) SetXOssAccessPointForObjectProcessName(v string) *CreateAccessPointForObjectProcessHeaders {
	s.XOssAccessPointForObjectProcessName = &v
	return s
}

func (s *CreateAccessPointForObjectProcessHeaders) Validate() error {
	return dara.Validate(s)
}
