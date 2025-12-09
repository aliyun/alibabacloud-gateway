// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostDataLakeStorageFileOperationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PostDataLakeStorageFileOperationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PostDataLakeStorageFileOperationResponse
	GetStatusCode() *int32
	SetBody(v string) *PostDataLakeStorageFileOperationResponse
	GetBody() *string
}

type PostDataLakeStorageFileOperationResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostDataLakeStorageFileOperationResponse) String() string {
	return dara.Prettify(s)
}

func (s PostDataLakeStorageFileOperationResponse) GoString() string {
	return s.String()
}

func (s *PostDataLakeStorageFileOperationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PostDataLakeStorageFileOperationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PostDataLakeStorageFileOperationResponse) GetBody() *string {
	return s.Body
}

func (s *PostDataLakeStorageFileOperationResponse) SetHeaders(v map[string]*string) *PostDataLakeStorageFileOperationResponse {
	s.Headers = v
	return s
}

func (s *PostDataLakeStorageFileOperationResponse) SetStatusCode(v int32) *PostDataLakeStorageFileOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *PostDataLakeStorageFileOperationResponse) SetBody(v string) *PostDataLakeStorageFileOperationResponse {
	s.Body = &v
	return s
}

func (s *PostDataLakeStorageFileOperationResponse) Validate() error {
	return dara.Validate(s)
}
