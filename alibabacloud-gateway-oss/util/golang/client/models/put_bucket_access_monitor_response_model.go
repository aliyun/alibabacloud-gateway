// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketAccessMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketAccessMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketAccessMonitorResponse
	GetStatusCode() *int32
}

type PutBucketAccessMonitorResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketAccessMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketAccessMonitorResponse) GoString() string {
	return s.String()
}

func (s *PutBucketAccessMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketAccessMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketAccessMonitorResponse) SetHeaders(v map[string]*string) *PutBucketAccessMonitorResponse {
	s.Headers = v
	return s
}

func (s *PutBucketAccessMonitorResponse) SetStatusCode(v int32) *PutBucketAccessMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketAccessMonitorResponse) Validate() error {
	return dara.Validate(s)
}
