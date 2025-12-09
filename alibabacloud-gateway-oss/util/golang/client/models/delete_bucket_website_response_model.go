// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketWebsiteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBucketWebsiteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBucketWebsiteResponse
	GetStatusCode() *int32
}

type DeleteBucketWebsiteResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteBucketWebsiteResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketWebsiteResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketWebsiteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBucketWebsiteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBucketWebsiteResponse) SetHeaders(v map[string]*string) *DeleteBucketWebsiteResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketWebsiteResponse) SetStatusCode(v int32) *DeleteBucketWebsiteResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBucketWebsiteResponse) Validate() error {
	return dara.Validate(s)
}
