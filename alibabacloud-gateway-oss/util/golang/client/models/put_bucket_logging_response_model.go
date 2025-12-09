// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketLoggingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketLoggingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketLoggingResponse
	GetStatusCode() *int32
}

type PutBucketLoggingResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketLoggingResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketLoggingResponse) GoString() string {
	return s.String()
}

func (s *PutBucketLoggingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketLoggingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketLoggingResponse) SetHeaders(v map[string]*string) *PutBucketLoggingResponse {
	s.Headers = v
	return s
}

func (s *PutBucketLoggingResponse) SetStatusCode(v int32) *PutBucketLoggingResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketLoggingResponse) Validate() error {
	return dara.Validate(s)
}
