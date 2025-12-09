// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketLoggingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBucketLoggingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBucketLoggingResponse
	GetStatusCode() *int32
}

type DeleteBucketLoggingResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteBucketLoggingResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketLoggingResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketLoggingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBucketLoggingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBucketLoggingResponse) SetHeaders(v map[string]*string) *DeleteBucketLoggingResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketLoggingResponse) SetStatusCode(v int32) *DeleteBucketLoggingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBucketLoggingResponse) Validate() error {
	return dara.Validate(s)
}
