// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessPointPublicAccessBlockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXOssAccessPointName(v string) *GetAccessPointPublicAccessBlockRequest
	GetXOssAccessPointName() *string
}

type GetAccessPointPublicAccessBlockRequest struct {
	// The name of the access point.
	//
	// example:
	//
	// ap-01
	XOssAccessPointName *string `json:"x-oss-access-point-name,omitempty" xml:"x-oss-access-point-name,omitempty"`
}

func (s GetAccessPointPublicAccessBlockRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccessPointPublicAccessBlockRequest) GoString() string {
	return s.String()
}

func (s *GetAccessPointPublicAccessBlockRequest) GetXOssAccessPointName() *string {
	return s.XOssAccessPointName
}

func (s *GetAccessPointPublicAccessBlockRequest) SetXOssAccessPointName(v string) *GetAccessPointPublicAccessBlockRequest {
	s.XOssAccessPointName = &v
	return s
}

func (s *GetAccessPointPublicAccessBlockRequest) Validate() error {
	return dara.Validate(s)
}
