// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBucketDataRedundancyTransitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBucketDataRedundancyTransitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBucketDataRedundancyTransitionResponse
	GetStatusCode() *int32
	SetBody(v *ListBucketDataRedundancyTransitionResponseBody) *ListBucketDataRedundancyTransitionResponse
	GetBody() *ListBucketDataRedundancyTransitionResponseBody
}

type ListBucketDataRedundancyTransitionResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBucketDataRedundancyTransitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBucketDataRedundancyTransitionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBucketDataRedundancyTransitionResponse) GoString() string {
	return s.String()
}

func (s *ListBucketDataRedundancyTransitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBucketDataRedundancyTransitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBucketDataRedundancyTransitionResponse) GetBody() *ListBucketDataRedundancyTransitionResponseBody {
	return s.Body
}

func (s *ListBucketDataRedundancyTransitionResponse) SetHeaders(v map[string]*string) *ListBucketDataRedundancyTransitionResponse {
	s.Headers = v
	return s
}

func (s *ListBucketDataRedundancyTransitionResponse) SetStatusCode(v int32) *ListBucketDataRedundancyTransitionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBucketDataRedundancyTransitionResponse) SetBody(v *ListBucketDataRedundancyTransitionResponseBody) *ListBucketDataRedundancyTransitionResponse {
	s.Body = v
	return s
}

func (s *ListBucketDataRedundancyTransitionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
