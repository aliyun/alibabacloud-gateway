// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketRtcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketRtcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketRtcResponse
	GetStatusCode() *int32
}

type PutBucketRtcResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketRtcResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketRtcResponse) GoString() string {
	return s.String()
}

func (s *PutBucketRtcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketRtcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketRtcResponse) SetHeaders(v map[string]*string) *PutBucketRtcResponse {
	s.Headers = v
	return s
}

func (s *PutBucketRtcResponse) SetStatusCode(v int32) *PutBucketRtcResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketRtcResponse) Validate() error {
	return dara.Validate(s)
}
