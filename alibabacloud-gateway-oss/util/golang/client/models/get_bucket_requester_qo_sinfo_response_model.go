// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketRequesterQoSInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketRequesterQoSInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketRequesterQoSInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketRequesterQoSInfoResponseBody) *GetBucketRequesterQoSInfoResponse
	GetBody() *GetBucketRequesterQoSInfoResponseBody
}

type GetBucketRequesterQoSInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketRequesterQoSInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketRequesterQoSInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketRequesterQoSInfoResponse) GoString() string {
	return s.String()
}

func (s *GetBucketRequesterQoSInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketRequesterQoSInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketRequesterQoSInfoResponse) GetBody() *GetBucketRequesterQoSInfoResponseBody {
	return s.Body
}

func (s *GetBucketRequesterQoSInfoResponse) SetHeaders(v map[string]*string) *GetBucketRequesterQoSInfoResponse {
	s.Headers = v
	return s
}

func (s *GetBucketRequesterQoSInfoResponse) SetStatusCode(v int32) *GetBucketRequesterQoSInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketRequesterQoSInfoResponse) SetBody(v *GetBucketRequesterQoSInfoResponseBody) *GetBucketRequesterQoSInfoResponse {
	s.Body = v
	return s
}

func (s *GetBucketRequesterQoSInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
