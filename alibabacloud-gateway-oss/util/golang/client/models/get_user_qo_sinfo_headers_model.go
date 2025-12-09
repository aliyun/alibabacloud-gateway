// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserQoSInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetUserQoSInfoHeaders
	GetCommonHeaders() map[string]*string
	SetXOssReturnDefault(v bool) *GetUserQoSInfoHeaders
	GetXOssReturnDefault() *bool
}

type GetUserQoSInfoHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// false
	XOssReturnDefault *bool `json:"x-oss-return-default,omitempty" xml:"x-oss-return-default,omitempty"`
}

func (s GetUserQoSInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetUserQoSInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetUserQoSInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetUserQoSInfoHeaders) GetXOssReturnDefault() *bool {
	return s.XOssReturnDefault
}

func (s *GetUserQoSInfoHeaders) SetCommonHeaders(v map[string]*string) *GetUserQoSInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserQoSInfoHeaders) SetXOssReturnDefault(v bool) *GetUserQoSInfoHeaders {
	s.XOssReturnDefault = &v
	return s
}

func (s *GetUserQoSInfoHeaders) Validate() error {
	return dara.Validate(s)
}
