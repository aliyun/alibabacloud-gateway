// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutObjectRetentionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutObjectRetentionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutObjectRetentionResponse
	GetStatusCode() *int32
}

type PutObjectRetentionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutObjectRetentionResponse) String() string {
	return dara.Prettify(s)
}

func (s PutObjectRetentionResponse) GoString() string {
	return s.String()
}

func (s *PutObjectRetentionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutObjectRetentionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutObjectRetentionResponse) SetHeaders(v map[string]*string) *PutObjectRetentionResponse {
	s.Headers = v
	return s
}

func (s *PutObjectRetentionResponse) SetStatusCode(v int32) *PutObjectRetentionResponse {
	s.StatusCode = &v
	return s
}

func (s *PutObjectRetentionResponse) Validate() error {
	return dara.Validate(s)
}
