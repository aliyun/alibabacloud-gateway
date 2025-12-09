// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBucketTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBucketTagsResponse
	GetStatusCode() *int32
}

type DeleteBucketTagsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteBucketTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketTagsResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBucketTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBucketTagsResponse) SetHeaders(v map[string]*string) *DeleteBucketTagsResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketTagsResponse) SetStatusCode(v int32) *DeleteBucketTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBucketTagsResponse) Validate() error {
	return dara.Validate(s)
}
