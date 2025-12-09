// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitiateBucketWormResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitiateBucketWormResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitiateBucketWormResponse
	GetStatusCode() *int32
}

type InitiateBucketWormResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s InitiateBucketWormResponse) String() string {
	return dara.Prettify(s)
}

func (s InitiateBucketWormResponse) GoString() string {
	return s.String()
}

func (s *InitiateBucketWormResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitiateBucketWormResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitiateBucketWormResponse) SetHeaders(v map[string]*string) *InitiateBucketWormResponse {
	s.Headers = v
	return s
}

func (s *InitiateBucketWormResponse) SetStatusCode(v int32) *InitiateBucketWormResponse {
	s.StatusCode = &v
	return s
}

func (s *InitiateBucketWormResponse) Validate() error {
	return dara.Validate(s)
}
