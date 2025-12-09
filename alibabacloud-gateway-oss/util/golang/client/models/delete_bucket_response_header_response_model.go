// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketResponseHeaderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBucketResponseHeaderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBucketResponseHeaderResponse
	GetStatusCode() *int32
}

type DeleteBucketResponseHeaderResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteBucketResponseHeaderResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketResponseHeaderResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketResponseHeaderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBucketResponseHeaderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBucketResponseHeaderResponse) SetHeaders(v map[string]*string) *DeleteBucketResponseHeaderResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketResponseHeaderResponse) SetStatusCode(v int32) *DeleteBucketResponseHeaderResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBucketResponseHeaderResponse) Validate() error {
	return dara.Validate(s)
}
