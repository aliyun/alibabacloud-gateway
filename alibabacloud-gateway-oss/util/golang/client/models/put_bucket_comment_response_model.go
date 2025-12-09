// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketCommentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketCommentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketCommentResponse
	GetStatusCode() *int32
}

type PutBucketCommentResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketCommentResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketCommentResponse) GoString() string {
	return s.String()
}

func (s *PutBucketCommentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketCommentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketCommentResponse) SetHeaders(v map[string]*string) *PutBucketCommentResponse {
	s.Headers = v
	return s
}

func (s *PutBucketCommentResponse) SetStatusCode(v int32) *PutBucketCommentResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketCommentResponse) Validate() error {
	return dara.Validate(s)
}
