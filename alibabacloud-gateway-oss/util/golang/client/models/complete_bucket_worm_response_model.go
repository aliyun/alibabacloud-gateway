// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteBucketWormResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CompleteBucketWormResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CompleteBucketWormResponse
	GetStatusCode() *int32
}

type CompleteBucketWormResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CompleteBucketWormResponse) String() string {
	return dara.Prettify(s)
}

func (s CompleteBucketWormResponse) GoString() string {
	return s.String()
}

func (s *CompleteBucketWormResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CompleteBucketWormResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CompleteBucketWormResponse) SetHeaders(v map[string]*string) *CompleteBucketWormResponse {
	s.Headers = v
	return s
}

func (s *CompleteBucketWormResponse) SetStatusCode(v int32) *CompleteBucketWormResponse {
	s.StatusCode = &v
	return s
}

func (s *CompleteBucketWormResponse) Validate() error {
	return dara.Validate(s)
}
