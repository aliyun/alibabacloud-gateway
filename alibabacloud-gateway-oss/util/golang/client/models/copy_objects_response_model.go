// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyObjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CopyObjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CopyObjectsResponse
	GetStatusCode() *int32
	SetBody(v *CopyObjectsResponseBody) *CopyObjectsResponse
	GetBody() *CopyObjectsResponseBody
}

type CopyObjectsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyObjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyObjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s CopyObjectsResponse) GoString() string {
	return s.String()
}

func (s *CopyObjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CopyObjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CopyObjectsResponse) GetBody() *CopyObjectsResponseBody {
	return s.Body
}

func (s *CopyObjectsResponse) SetHeaders(v map[string]*string) *CopyObjectsResponse {
	s.Headers = v
	return s
}

func (s *CopyObjectsResponse) SetStatusCode(v int32) *CopyObjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyObjectsResponse) SetBody(v *CopyObjectsResponseBody) *CopyObjectsResponse {
	s.Body = v
	return s
}

func (s *CopyObjectsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
