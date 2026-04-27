// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutObjectLegalHoldResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutObjectLegalHoldResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutObjectLegalHoldResponse
	GetStatusCode() *int32
}

type PutObjectLegalHoldResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutObjectLegalHoldResponse) String() string {
	return dara.Prettify(s)
}

func (s PutObjectLegalHoldResponse) GoString() string {
	return s.String()
}

func (s *PutObjectLegalHoldResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutObjectLegalHoldResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutObjectLegalHoldResponse) SetHeaders(v map[string]*string) *PutObjectLegalHoldResponse {
	s.Headers = v
	return s
}

func (s *PutObjectLegalHoldResponse) SetStatusCode(v int32) *PutObjectLegalHoldResponse {
	s.StatusCode = &v
	return s
}

func (s *PutObjectLegalHoldResponse) Validate() error {
	return dara.Validate(s)
}
