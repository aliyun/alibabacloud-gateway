// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetObjectRetentionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetObjectRetentionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetObjectRetentionResponse
	GetStatusCode() *int32
	SetBody(v *GetObjectRetentionResponseBody) *GetObjectRetentionResponse
	GetBody() *GetObjectRetentionResponseBody
}

type GetObjectRetentionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetObjectRetentionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetObjectRetentionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetObjectRetentionResponse) GoString() string {
	return s.String()
}

func (s *GetObjectRetentionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetObjectRetentionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetObjectRetentionResponse) GetBody() *GetObjectRetentionResponseBody {
	return s.Body
}

func (s *GetObjectRetentionResponse) SetHeaders(v map[string]*string) *GetObjectRetentionResponse {
	s.Headers = v
	return s
}

func (s *GetObjectRetentionResponse) SetStatusCode(v int32) *GetObjectRetentionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetObjectRetentionResponse) SetBody(v *GetObjectRetentionResponseBody) *GetObjectRetentionResponse {
	s.Body = v
	return s
}

func (s *GetObjectRetentionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
