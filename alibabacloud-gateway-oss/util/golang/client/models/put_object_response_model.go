// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutObjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutObjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutObjectResponse
	GetStatusCode() *int32
}

type PutObjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutObjectResponse) String() string {
	return dara.Prettify(s)
}

func (s PutObjectResponse) GoString() string {
	return s.String()
}

func (s *PutObjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutObjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutObjectResponse) SetHeaders(v map[string]*string) *PutObjectResponse {
	s.Headers = v
	return s
}

func (s *PutObjectResponse) SetStatusCode(v int32) *PutObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *PutObjectResponse) Validate() error {
	return dara.Validate(s)
}
