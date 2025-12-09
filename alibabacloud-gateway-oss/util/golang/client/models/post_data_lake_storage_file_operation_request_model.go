// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostDataLakeStorageFileOperationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *PostDataLakeStorageFileOperationRequest
	GetBody() *string
}

type PostDataLakeStorageFileOperationRequest struct {
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostDataLakeStorageFileOperationRequest) String() string {
	return dara.Prettify(s)
}

func (s PostDataLakeStorageFileOperationRequest) GoString() string {
	return s.String()
}

func (s *PostDataLakeStorageFileOperationRequest) GetBody() *string {
	return s.Body
}

func (s *PostDataLakeStorageFileOperationRequest) SetBody(v string) *PostDataLakeStorageFileOperationRequest {
	s.Body = &v
	return s
}

func (s *PostDataLakeStorageFileOperationRequest) Validate() error {
	return dara.Validate(s)
}
