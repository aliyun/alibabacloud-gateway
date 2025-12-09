// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketQoSInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBucketQoSInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBucketQoSInfoResponse
	GetStatusCode() *int32
}

type DeleteBucketQoSInfoResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteBucketQoSInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketQoSInfoResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketQoSInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBucketQoSInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBucketQoSInfoResponse) SetHeaders(v map[string]*string) *DeleteBucketQoSInfoResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketQoSInfoResponse) SetStatusCode(v int32) *DeleteBucketQoSInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBucketQoSInfoResponse) Validate() error {
	return dara.Validate(s)
}
