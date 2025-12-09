// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketDataRedundancyTransitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketDataRedundancyTransitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketDataRedundancyTransitionResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketDataRedundancyTransitionResponseBody) *GetBucketDataRedundancyTransitionResponse
	GetBody() *GetBucketDataRedundancyTransitionResponseBody
}

type GetBucketDataRedundancyTransitionResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketDataRedundancyTransitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketDataRedundancyTransitionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketDataRedundancyTransitionResponse) GoString() string {
	return s.String()
}

func (s *GetBucketDataRedundancyTransitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketDataRedundancyTransitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketDataRedundancyTransitionResponse) GetBody() *GetBucketDataRedundancyTransitionResponseBody {
	return s.Body
}

func (s *GetBucketDataRedundancyTransitionResponse) SetHeaders(v map[string]*string) *GetBucketDataRedundancyTransitionResponse {
	s.Headers = v
	return s
}

func (s *GetBucketDataRedundancyTransitionResponse) SetStatusCode(v int32) *GetBucketDataRedundancyTransitionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketDataRedundancyTransitionResponse) SetBody(v *GetBucketDataRedundancyTransitionResponseBody) *GetBucketDataRedundancyTransitionResponse {
	s.Body = v
	return s
}

func (s *GetBucketDataRedundancyTransitionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
