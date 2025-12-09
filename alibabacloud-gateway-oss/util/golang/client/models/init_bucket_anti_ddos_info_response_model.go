// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitBucketAntiDDosInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitBucketAntiDDosInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitBucketAntiDDosInfoResponse
	GetStatusCode() *int32
}

type InitBucketAntiDDosInfoResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s InitBucketAntiDDosInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s InitBucketAntiDDosInfoResponse) GoString() string {
	return s.String()
}

func (s *InitBucketAntiDDosInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitBucketAntiDDosInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitBucketAntiDDosInfoResponse) SetHeaders(v map[string]*string) *InitBucketAntiDDosInfoResponse {
	s.Headers = v
	return s
}

func (s *InitBucketAntiDDosInfoResponse) SetStatusCode(v int32) *InitBucketAntiDDosInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *InitBucketAntiDDosInfoResponse) Validate() error {
	return dara.Validate(s)
}
