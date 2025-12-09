// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostDataLakeStorageSecurityOperationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PostDataLakeStorageSecurityOperationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PostDataLakeStorageSecurityOperationResponse
	GetStatusCode() *int32
	SetBody(v string) *PostDataLakeStorageSecurityOperationResponse
	GetBody() *string
}

type PostDataLakeStorageSecurityOperationResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostDataLakeStorageSecurityOperationResponse) String() string {
	return dara.Prettify(s)
}

func (s PostDataLakeStorageSecurityOperationResponse) GoString() string {
	return s.String()
}

func (s *PostDataLakeStorageSecurityOperationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PostDataLakeStorageSecurityOperationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PostDataLakeStorageSecurityOperationResponse) GetBody() *string {
	return s.Body
}

func (s *PostDataLakeStorageSecurityOperationResponse) SetHeaders(v map[string]*string) *PostDataLakeStorageSecurityOperationResponse {
	s.Headers = v
	return s
}

func (s *PostDataLakeStorageSecurityOperationResponse) SetStatusCode(v int32) *PostDataLakeStorageSecurityOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *PostDataLakeStorageSecurityOperationResponse) SetBody(v string) *PostDataLakeStorageSecurityOperationResponse {
	s.Body = &v
	return s
}

func (s *PostDataLakeStorageSecurityOperationResponse) Validate() error {
	return dara.Validate(s)
}
