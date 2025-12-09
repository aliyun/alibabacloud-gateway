// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketTagsResponse
	GetStatusCode() *int32
}

type PutBucketTagsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketTagsResponse) GoString() string {
	return s.String()
}

func (s *PutBucketTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketTagsResponse) SetHeaders(v map[string]*string) *PutBucketTagsResponse {
	s.Headers = v
	return s
}

func (s *PutBucketTagsResponse) SetStatusCode(v int32) *PutBucketTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketTagsResponse) Validate() error {
	return dara.Validate(s)
}
