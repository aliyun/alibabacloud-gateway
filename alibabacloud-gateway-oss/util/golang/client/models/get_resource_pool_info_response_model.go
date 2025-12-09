// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourcePoolInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourcePoolInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourcePoolInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetResourcePoolInfoResponseBody) *GetResourcePoolInfoResponse
	GetBody() *GetResourcePoolInfoResponseBody
}

type GetResourcePoolInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourcePoolInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourcePoolInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourcePoolInfoResponse) GoString() string {
	return s.String()
}

func (s *GetResourcePoolInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourcePoolInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourcePoolInfoResponse) GetBody() *GetResourcePoolInfoResponseBody {
	return s.Body
}

func (s *GetResourcePoolInfoResponse) SetHeaders(v map[string]*string) *GetResourcePoolInfoResponse {
	s.Headers = v
	return s
}

func (s *GetResourcePoolInfoResponse) SetStatusCode(v int32) *GetResourcePoolInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourcePoolInfoResponse) SetBody(v *GetResourcePoolInfoResponseBody) *GetResourcePoolInfoResponse {
	s.Body = v
	return s
}

func (s *GetResourcePoolInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
