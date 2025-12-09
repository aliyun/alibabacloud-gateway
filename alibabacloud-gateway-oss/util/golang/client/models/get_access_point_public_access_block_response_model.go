// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessPointPublicAccessBlockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccessPointPublicAccessBlockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccessPointPublicAccessBlockResponse
	GetStatusCode() *int32
	SetBody(v *GetAccessPointPublicAccessBlockResponseBody) *GetAccessPointPublicAccessBlockResponse
	GetBody() *GetAccessPointPublicAccessBlockResponseBody
}

type GetAccessPointPublicAccessBlockResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccessPointPublicAccessBlockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccessPointPublicAccessBlockResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccessPointPublicAccessBlockResponse) GoString() string {
	return s.String()
}

func (s *GetAccessPointPublicAccessBlockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccessPointPublicAccessBlockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccessPointPublicAccessBlockResponse) GetBody() *GetAccessPointPublicAccessBlockResponseBody {
	return s.Body
}

func (s *GetAccessPointPublicAccessBlockResponse) SetHeaders(v map[string]*string) *GetAccessPointPublicAccessBlockResponse {
	s.Headers = v
	return s
}

func (s *GetAccessPointPublicAccessBlockResponse) SetStatusCode(v int32) *GetAccessPointPublicAccessBlockResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessPointPublicAccessBlockResponse) SetBody(v *GetAccessPointPublicAccessBlockResponseBody) *GetAccessPointPublicAccessBlockResponse {
	s.Body = v
	return s
}

func (s *GetAccessPointPublicAccessBlockResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
