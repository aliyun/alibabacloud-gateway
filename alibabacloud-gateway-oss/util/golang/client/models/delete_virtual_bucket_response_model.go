// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVirtualBucketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVirtualBucketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVirtualBucketResponse
	GetStatusCode() *int32
}

type DeleteVirtualBucketResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteVirtualBucketResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVirtualBucketResponse) GoString() string {
	return s.String()
}

func (s *DeleteVirtualBucketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVirtualBucketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVirtualBucketResponse) SetHeaders(v map[string]*string) *DeleteVirtualBucketResponse {
	s.Headers = v
	return s
}

func (s *DeleteVirtualBucketResponse) SetStatusCode(v int32) *DeleteVirtualBucketResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVirtualBucketResponse) Validate() error {
	return dara.Validate(s)
}
