// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessPointHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetAccessPointHeaders
	GetCommonHeaders() map[string]*string
	SetXOssAccessPointName(v string) *GetAccessPointHeaders
	GetXOssAccessPointName() *string
}

type GetAccessPointHeaders struct {
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

func (s GetAccessPointHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetAccessPointHeaders) GoString() string {
	return s.String()
}

func (s *GetAccessPointHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetAccessPointHeaders) GetXOssAccessPointName() *string {
	return s.XOssAccessPointName
}

func (s *GetAccessPointHeaders) SetCommonHeaders(v map[string]*string) *GetAccessPointHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAccessPointHeaders) SetXOssAccessPointName(v string) *GetAccessPointHeaders {
	s.XOssAccessPointName = &v
	return s
}

func (s *GetAccessPointHeaders) Validate() error {
	return dara.Validate(s)
}
