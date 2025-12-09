// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetObjectTaggingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetObjectTaggingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetObjectTaggingResponse
	GetStatusCode() *int32
	SetBody(v *GetObjectTaggingResponseBody) *GetObjectTaggingResponse
	GetBody() *GetObjectTaggingResponseBody
}

type GetObjectTaggingResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetObjectTaggingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetObjectTaggingResponse) String() string {
	return dara.Prettify(s)
}

func (s GetObjectTaggingResponse) GoString() string {
	return s.String()
}

func (s *GetObjectTaggingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetObjectTaggingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetObjectTaggingResponse) GetBody() *GetObjectTaggingResponseBody {
	return s.Body
}

func (s *GetObjectTaggingResponse) SetHeaders(v map[string]*string) *GetObjectTaggingResponse {
	s.Headers = v
	return s
}

func (s *GetObjectTaggingResponse) SetStatusCode(v int32) *GetObjectTaggingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetObjectTaggingResponse) SetBody(v *GetObjectTaggingResponseBody) *GetObjectTaggingResponse {
	s.Body = v
	return s
}

func (s *GetObjectTaggingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
