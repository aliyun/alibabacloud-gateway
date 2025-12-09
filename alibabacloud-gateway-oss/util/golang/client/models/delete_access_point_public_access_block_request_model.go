// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessPointPublicAccessBlockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXOssAccessPointName(v string) *DeleteAccessPointPublicAccessBlockRequest
	GetXOssAccessPointName() *string
}

type DeleteAccessPointPublicAccessBlockRequest struct {
	// The name of the access point.
	//
	// example:
	//
	// ap-01
	XOssAccessPointName *string `json:"x-oss-access-point-name,omitempty" xml:"x-oss-access-point-name,omitempty"`
}

func (s DeleteAccessPointPublicAccessBlockRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessPointPublicAccessBlockRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessPointPublicAccessBlockRequest) GetXOssAccessPointName() *string {
	return s.XOssAccessPointName
}

func (s *DeleteAccessPointPublicAccessBlockRequest) SetXOssAccessPointName(v string) *DeleteAccessPointPublicAccessBlockRequest {
	s.XOssAccessPointName = &v
	return s
}

func (s *DeleteAccessPointPublicAccessBlockRequest) Validate() error {
	return dara.Validate(s)
}
