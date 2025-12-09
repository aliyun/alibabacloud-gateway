// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostDataLakeStorageAdminOperationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *PostDataLakeStorageAdminOperationRequest
	GetBody() *string
}

type PostDataLakeStorageAdminOperationRequest struct {
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostDataLakeStorageAdminOperationRequest) String() string {
	return dara.Prettify(s)
}

func (s PostDataLakeStorageAdminOperationRequest) GoString() string {
	return s.String()
}

func (s *PostDataLakeStorageAdminOperationRequest) GetBody() *string {
	return s.Body
}

func (s *PostDataLakeStorageAdminOperationRequest) SetBody(v string) *PostDataLakeStorageAdminOperationRequest {
	s.Body = &v
	return s
}

func (s *PostDataLakeStorageAdminOperationRequest) Validate() error {
	return dara.Validate(s)
}
