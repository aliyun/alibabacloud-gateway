// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPublicAccessBlockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPublicAccessBlockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPublicAccessBlockResponse
	GetStatusCode() *int32
	SetBody(v *GetPublicAccessBlockResponseBody) *GetPublicAccessBlockResponse
	GetBody() *GetPublicAccessBlockResponseBody
}

type GetPublicAccessBlockResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPublicAccessBlockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPublicAccessBlockResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPublicAccessBlockResponse) GoString() string {
	return s.String()
}

func (s *GetPublicAccessBlockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPublicAccessBlockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPublicAccessBlockResponse) GetBody() *GetPublicAccessBlockResponseBody {
	return s.Body
}

func (s *GetPublicAccessBlockResponse) SetHeaders(v map[string]*string) *GetPublicAccessBlockResponse {
	s.Headers = v
	return s
}

func (s *GetPublicAccessBlockResponse) SetStatusCode(v int32) *GetPublicAccessBlockResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPublicAccessBlockResponse) SetBody(v *GetPublicAccessBlockResponseBody) *GetPublicAccessBlockResponse {
	s.Body = v
	return s
}

func (s *GetPublicAccessBlockResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
