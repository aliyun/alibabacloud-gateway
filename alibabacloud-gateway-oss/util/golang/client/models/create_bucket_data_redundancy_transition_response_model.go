// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBucketDataRedundancyTransitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBucketDataRedundancyTransitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBucketDataRedundancyTransitionResponse
	GetStatusCode() *int32
	SetBody(v *CreateBucketDataRedundancyTransitionResponseBody) *CreateBucketDataRedundancyTransitionResponse
	GetBody() *CreateBucketDataRedundancyTransitionResponseBody
}

type CreateBucketDataRedundancyTransitionResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBucketDataRedundancyTransitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBucketDataRedundancyTransitionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBucketDataRedundancyTransitionResponse) GoString() string {
	return s.String()
}

func (s *CreateBucketDataRedundancyTransitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBucketDataRedundancyTransitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBucketDataRedundancyTransitionResponse) GetBody() *CreateBucketDataRedundancyTransitionResponseBody {
	return s.Body
}

func (s *CreateBucketDataRedundancyTransitionResponse) SetHeaders(v map[string]*string) *CreateBucketDataRedundancyTransitionResponse {
	s.Headers = v
	return s
}

func (s *CreateBucketDataRedundancyTransitionResponse) SetStatusCode(v int32) *CreateBucketDataRedundancyTransitionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBucketDataRedundancyTransitionResponse) SetBody(v *CreateBucketDataRedundancyTransitionResponseBody) *CreateBucketDataRedundancyTransitionResponse {
	s.Body = v
	return s
}

func (s *CreateBucketDataRedundancyTransitionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
