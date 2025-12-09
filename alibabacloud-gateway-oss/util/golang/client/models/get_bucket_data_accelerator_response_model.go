// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketDataAcceleratorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketDataAcceleratorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketDataAcceleratorResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketDataAcceleratorResponseBody) *GetBucketDataAcceleratorResponse
	GetBody() *GetBucketDataAcceleratorResponseBody
}

type GetBucketDataAcceleratorResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketDataAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketDataAcceleratorResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketDataAcceleratorResponse) GoString() string {
	return s.String()
}

func (s *GetBucketDataAcceleratorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketDataAcceleratorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketDataAcceleratorResponse) GetBody() *GetBucketDataAcceleratorResponseBody {
	return s.Body
}

func (s *GetBucketDataAcceleratorResponse) SetHeaders(v map[string]*string) *GetBucketDataAcceleratorResponse {
	s.Headers = v
	return s
}

func (s *GetBucketDataAcceleratorResponse) SetStatusCode(v int32) *GetBucketDataAcceleratorResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketDataAcceleratorResponse) SetBody(v *GetBucketDataAcceleratorResponseBody) *GetBucketDataAcceleratorResponse {
	s.Body = v
	return s
}

func (s *GetBucketDataAcceleratorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
