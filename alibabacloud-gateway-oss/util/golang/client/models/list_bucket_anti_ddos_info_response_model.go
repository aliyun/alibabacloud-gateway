// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBucketAntiDDosInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBucketAntiDDosInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBucketAntiDDosInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListBucketAntiDDosInfoResponseBody) *ListBucketAntiDDosInfoResponse
	GetBody() *ListBucketAntiDDosInfoResponseBody
}

type ListBucketAntiDDosInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBucketAntiDDosInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBucketAntiDDosInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBucketAntiDDosInfoResponse) GoString() string {
	return s.String()
}

func (s *ListBucketAntiDDosInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBucketAntiDDosInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBucketAntiDDosInfoResponse) GetBody() *ListBucketAntiDDosInfoResponseBody {
	return s.Body
}

func (s *ListBucketAntiDDosInfoResponse) SetHeaders(v map[string]*string) *ListBucketAntiDDosInfoResponse {
	s.Headers = v
	return s
}

func (s *ListBucketAntiDDosInfoResponse) SetStatusCode(v int32) *ListBucketAntiDDosInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBucketAntiDDosInfoResponse) SetBody(v *ListBucketAntiDDosInfoResponseBody) *ListBucketAntiDDosInfoResponse {
	s.Body = v
	return s
}

func (s *ListBucketAntiDDosInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
