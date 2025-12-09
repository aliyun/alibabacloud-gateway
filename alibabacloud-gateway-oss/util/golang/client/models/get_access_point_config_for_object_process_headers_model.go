// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessPointConfigForObjectProcessHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetAccessPointConfigForObjectProcessHeaders
	GetCommonHeaders() map[string]*string
	SetXOssAccessPointForObjectProcessName(v string) *GetAccessPointConfigForObjectProcessHeaders
	GetXOssAccessPointForObjectProcessName() *string
}

type GetAccessPointConfigForObjectProcessHeaders struct {
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

func (s GetAccessPointConfigForObjectProcessHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetAccessPointConfigForObjectProcessHeaders) GoString() string {
	return s.String()
}

func (s *GetAccessPointConfigForObjectProcessHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetAccessPointConfigForObjectProcessHeaders) GetXOssAccessPointForObjectProcessName() *string {
	return s.XOssAccessPointForObjectProcessName
}

func (s *GetAccessPointConfigForObjectProcessHeaders) SetCommonHeaders(v map[string]*string) *GetAccessPointConfigForObjectProcessHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAccessPointConfigForObjectProcessHeaders) SetXOssAccessPointForObjectProcessName(v string) *GetAccessPointConfigForObjectProcessHeaders {
	s.XOssAccessPointForObjectProcessName = &v
	return s
}

func (s *GetAccessPointConfigForObjectProcessHeaders) Validate() error {
	return dara.Validate(s)
}
