// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketDataRedundancyTransitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBucketDataRedundancyTransitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBucketDataRedundancyTransitionResponse
	GetStatusCode() *int32
}

type DeleteBucketDataRedundancyTransitionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteBucketDataRedundancyTransitionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketDataRedundancyTransitionResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketDataRedundancyTransitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBucketDataRedundancyTransitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBucketDataRedundancyTransitionResponse) SetHeaders(v map[string]*string) *DeleteBucketDataRedundancyTransitionResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketDataRedundancyTransitionResponse) SetStatusCode(v int32) *DeleteBucketDataRedundancyTransitionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBucketDataRedundancyTransitionResponse) Validate() error {
	return dara.Validate(s)
}
