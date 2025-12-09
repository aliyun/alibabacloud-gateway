// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketRequesterQoSInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBucketRequesterQoSInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBucketRequesterQoSInfoResponse
	GetStatusCode() *int32
}

type DeleteBucketRequesterQoSInfoResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteBucketRequesterQoSInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketRequesterQoSInfoResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketRequesterQoSInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBucketRequesterQoSInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBucketRequesterQoSInfoResponse) SetHeaders(v map[string]*string) *DeleteBucketRequesterQoSInfoResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketRequesterQoSInfoResponse) SetStatusCode(v int32) *DeleteBucketRequesterQoSInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBucketRequesterQoSInfoResponse) Validate() error {
	return dara.Validate(s)
}
