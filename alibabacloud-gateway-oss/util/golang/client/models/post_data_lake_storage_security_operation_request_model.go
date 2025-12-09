// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostDataLakeStorageSecurityOperationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *PostDataLakeStorageSecurityOperationRequest
	GetBody() *string
}

type PostDataLakeStorageSecurityOperationRequest struct {
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostDataLakeStorageSecurityOperationRequest) String() string {
	return dara.Prettify(s)
}

func (s PostDataLakeStorageSecurityOperationRequest) GoString() string {
	return s.String()
}

func (s *PostDataLakeStorageSecurityOperationRequest) GetBody() *string {
	return s.Body
}

func (s *PostDataLakeStorageSecurityOperationRequest) SetBody(v string) *PostDataLakeStorageSecurityOperationRequest {
	s.Body = &v
	return s
}

func (s *PostDataLakeStorageSecurityOperationRequest) Validate() error {
	return dara.Validate(s)
}
