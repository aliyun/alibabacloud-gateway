// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketRedundancyTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutBucketRedundancyTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutBucketRedundancyTypeResponse
	GetStatusCode() *int32
}

type PutBucketRedundancyTypeResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutBucketRedundancyTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s PutBucketRedundancyTypeResponse) GoString() string {
	return s.String()
}

func (s *PutBucketRedundancyTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutBucketRedundancyTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutBucketRedundancyTypeResponse) SetHeaders(v map[string]*string) *PutBucketRedundancyTypeResponse {
	s.Headers = v
	return s
}

func (s *PutBucketRedundancyTypeResponse) SetStatusCode(v int32) *PutBucketRedundancyTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *PutBucketRedundancyTypeResponse) Validate() error {
	return dara.Validate(s)
}
