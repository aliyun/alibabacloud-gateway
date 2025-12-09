// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbortBucketWormResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AbortBucketWormResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AbortBucketWormResponse
	GetStatusCode() *int32
}

type AbortBucketWormResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s AbortBucketWormResponse) String() string {
	return dara.Prettify(s)
}

func (s AbortBucketWormResponse) GoString() string {
	return s.String()
}

func (s *AbortBucketWormResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AbortBucketWormResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AbortBucketWormResponse) SetHeaders(v map[string]*string) *AbortBucketWormResponse {
	s.Headers = v
	return s
}

func (s *AbortBucketWormResponse) SetStatusCode(v int32) *AbortBucketWormResponse {
	s.StatusCode = &v
	return s
}

func (s *AbortBucketWormResponse) Validate() error {
	return dara.Validate(s)
}
