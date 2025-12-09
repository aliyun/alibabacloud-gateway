// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaQueryStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMetaQueryStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMetaQueryStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetMetaQueryStatusResponseBody) *GetMetaQueryStatusResponse
	GetBody() *GetMetaQueryStatusResponseBody
}

type GetMetaQueryStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMetaQueryStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMetaQueryStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMetaQueryStatusResponse) GoString() string {
	return s.String()
}

func (s *GetMetaQueryStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMetaQueryStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMetaQueryStatusResponse) GetBody() *GetMetaQueryStatusResponseBody {
	return s.Body
}

func (s *GetMetaQueryStatusResponse) SetHeaders(v map[string]*string) *GetMetaQueryStatusResponse {
	s.Headers = v
	return s
}

func (s *GetMetaQueryStatusResponse) SetStatusCode(v int32) *GetMetaQueryStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetaQueryStatusResponse) SetBody(v *GetMetaQueryStatusResponseBody) *GetMetaQueryStatusResponse {
	s.Body = v
	return s
}

func (s *GetMetaQueryStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
