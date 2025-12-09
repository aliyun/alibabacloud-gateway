// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessPointForObjectProcessHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteAccessPointForObjectProcessHeaders
	GetCommonHeaders() map[string]*string
	SetXOssAccessPointForObjectProcessName(v string) *DeleteAccessPointForObjectProcessHeaders
	GetXOssAccessPointForObjectProcessName() *string
}

type DeleteAccessPointForObjectProcessHeaders struct {
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

func (s DeleteAccessPointForObjectProcessHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessPointForObjectProcessHeaders) GoString() string {
	return s.String()
}

func (s *DeleteAccessPointForObjectProcessHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteAccessPointForObjectProcessHeaders) GetXOssAccessPointForObjectProcessName() *string {
	return s.XOssAccessPointForObjectProcessName
}

func (s *DeleteAccessPointForObjectProcessHeaders) SetCommonHeaders(v map[string]*string) *DeleteAccessPointForObjectProcessHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteAccessPointForObjectProcessHeaders) SetXOssAccessPointForObjectProcessName(v string) *DeleteAccessPointForObjectProcessHeaders {
	s.XOssAccessPointForObjectProcessName = &v
	return s
}

func (s *DeleteAccessPointForObjectProcessHeaders) Validate() error {
	return dara.Validate(s)
}
