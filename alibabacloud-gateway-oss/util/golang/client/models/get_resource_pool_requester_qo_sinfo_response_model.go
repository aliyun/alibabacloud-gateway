// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourcePoolRequesterQoSInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourcePoolRequesterQoSInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourcePoolRequesterQoSInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetResourcePoolRequesterQoSInfoResponseBody) *GetResourcePoolRequesterQoSInfoResponse
	GetBody() *GetResourcePoolRequesterQoSInfoResponseBody
}

type GetResourcePoolRequesterQoSInfoResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourcePoolRequesterQoSInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourcePoolRequesterQoSInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourcePoolRequesterQoSInfoResponse) GoString() string {
	return s.String()
}

func (s *GetResourcePoolRequesterQoSInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourcePoolRequesterQoSInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourcePoolRequesterQoSInfoResponse) GetBody() *GetResourcePoolRequesterQoSInfoResponseBody {
	return s.Body
}

func (s *GetResourcePoolRequesterQoSInfoResponse) SetHeaders(v map[string]*string) *GetResourcePoolRequesterQoSInfoResponse {
	s.Headers = v
	return s
}

func (s *GetResourcePoolRequesterQoSInfoResponse) SetStatusCode(v int32) *GetResourcePoolRequesterQoSInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourcePoolRequesterQoSInfoResponse) SetBody(v *GetResourcePoolRequesterQoSInfoResponseBody) *GetResourcePoolRequesterQoSInfoResponse {
	s.Body = v
	return s
}

func (s *GetResourcePoolRequesterQoSInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
