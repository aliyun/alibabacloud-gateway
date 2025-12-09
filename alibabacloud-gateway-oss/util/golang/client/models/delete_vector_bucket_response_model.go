// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVectorBucketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVectorBucketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVectorBucketResponse
	GetStatusCode() *int32
}

type DeleteVectorBucketResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteVectorBucketResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVectorBucketResponse) GoString() string {
	return s.String()
}

func (s *DeleteVectorBucketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVectorBucketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVectorBucketResponse) SetHeaders(v map[string]*string) *DeleteVectorBucketResponse {
	s.Headers = v
	return s
}

func (s *DeleteVectorBucketResponse) SetStatusCode(v int32) *DeleteVectorBucketResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVectorBucketResponse) Validate() error {
	return dara.Validate(s)
}
