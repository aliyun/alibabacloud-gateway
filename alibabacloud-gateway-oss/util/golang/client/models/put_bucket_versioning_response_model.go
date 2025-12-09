// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketVersioningResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketVersioningResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketVersioningResponse
	GetStatusCode() *int32
}

type PutBucketVersioningResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketVersioningResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketVersioningResponse) GoString() string {
	return s.String()
}

func (s *PutBucketVersioningResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketVersioningResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketVersioningResponse) SetHeaders(v map[string]*string) *PutBucketVersioningResponse {
	s.Headers = v
	return s
}

func (s *PutBucketVersioningResponse) SetStatusCode(v int32) *PutBucketVersioningResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketVersioningResponse) Validate() error {
	return dara.Validate(s)
}
