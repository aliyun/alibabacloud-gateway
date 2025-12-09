// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketCommonHeaderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBucketCommonHeaderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBucketCommonHeaderResponse
	GetStatusCode() *int32
}

type DeleteBucketCommonHeaderResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteBucketCommonHeaderResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketCommonHeaderResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketCommonHeaderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBucketCommonHeaderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBucketCommonHeaderResponse) SetHeaders(v map[string]*string) *DeleteBucketCommonHeaderResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketCommonHeaderResponse) SetStatusCode(v int32) *DeleteBucketCommonHeaderResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBucketCommonHeaderResponse) Validate() error {
	return dara.Validate(s)
}
