// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostDataLakeStorageAdminOperationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PostDataLakeStorageAdminOperationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PostDataLakeStorageAdminOperationResponse
	GetStatusCode() *int32
	SetBody(v string) *PostDataLakeStorageAdminOperationResponse
	GetBody() *string
}

type PostDataLakeStorageAdminOperationResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostDataLakeStorageAdminOperationResponse) String() string {
	return dara.Prettify(s)
}

func (s PostDataLakeStorageAdminOperationResponse) GoString() string {
	return s.String()
}

func (s *PostDataLakeStorageAdminOperationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PostDataLakeStorageAdminOperationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PostDataLakeStorageAdminOperationResponse) GetBody() *string {
	return s.Body
}

func (s *PostDataLakeStorageAdminOperationResponse) SetHeaders(v map[string]*string) *PostDataLakeStorageAdminOperationResponse {
	s.Headers = v
	return s
}

func (s *PostDataLakeStorageAdminOperationResponse) SetStatusCode(v int32) *PostDataLakeStorageAdminOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *PostDataLakeStorageAdminOperationResponse) SetBody(v string) *PostDataLakeStorageAdminOperationResponse {
	s.Body = &v
	return s
}

func (s *PostDataLakeStorageAdminOperationResponse) Validate() error {
	return dara.Validate(s)
}
