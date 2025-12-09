// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessPointHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteAccessPointHeaders
	GetCommonHeaders() map[string]*string
	SetXOssAccessPointName(v string) *DeleteAccessPointHeaders
	GetXOssAccessPointName() *string
}

type DeleteAccessPointHeaders struct {
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

func (s DeleteAccessPointHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessPointHeaders) GoString() string {
	return s.String()
}

func (s *DeleteAccessPointHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteAccessPointHeaders) GetXOssAccessPointName() *string {
	return s.XOssAccessPointName
}

func (s *DeleteAccessPointHeaders) SetCommonHeaders(v map[string]*string) *DeleteAccessPointHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteAccessPointHeaders) SetXOssAccessPointName(v string) *DeleteAccessPointHeaders {
	s.XOssAccessPointName = &v
	return s
}

func (s *DeleteAccessPointHeaders) Validate() error {
	return dara.Validate(s)
}
