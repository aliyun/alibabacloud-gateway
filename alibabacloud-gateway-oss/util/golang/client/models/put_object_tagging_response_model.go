// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutObjectTaggingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutObjectTaggingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutObjectTaggingResponse
	GetStatusCode() *int32
}

type PutObjectTaggingResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutObjectTaggingResponse) String() string {
	return dara.Prettify(s)
}

func (s PutObjectTaggingResponse) GoString() string {
	return s.String()
}

func (s *PutObjectTaggingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutObjectTaggingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutObjectTaggingResponse) SetHeaders(v map[string]*string) *PutObjectTaggingResponse {
	s.Headers = v
	return s
}

func (s *PutObjectTaggingResponse) SetStatusCode(v int32) *PutObjectTaggingResponse {
	s.StatusCode = &v
	return s
}

func (s *PutObjectTaggingResponse) Validate() error {
	return dara.Validate(s)
}
