// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketAccessMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketAccessMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketAccessMonitorResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketAccessMonitorResponseBody) *GetBucketAccessMonitorResponse
	GetBody() *GetBucketAccessMonitorResponseBody
}

type GetBucketAccessMonitorResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketAccessMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketAccessMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketAccessMonitorResponse) GoString() string {
	return s.String()
}

func (s *GetBucketAccessMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketAccessMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketAccessMonitorResponse) GetBody() *GetBucketAccessMonitorResponseBody {
	return s.Body
}

func (s *GetBucketAccessMonitorResponse) SetHeaders(v map[string]*string) *GetBucketAccessMonitorResponse {
	s.Headers = v
	return s
}

func (s *GetBucketAccessMonitorResponse) SetStatusCode(v int32) *GetBucketAccessMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketAccessMonitorResponse) SetBody(v *GetBucketAccessMonitorResponseBody) *GetBucketAccessMonitorResponse {
	s.Body = v
	return s
}

func (s *GetBucketAccessMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
