// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourcePoolBucketGroupQoSInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourcePoolBucketGroupQoSInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourcePoolBucketGroupQoSInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetResourcePoolBucketGroupQoSInfoResponseBody) *GetResourcePoolBucketGroupQoSInfoResponse
	GetBody() *GetResourcePoolBucketGroupQoSInfoResponseBody
}

type GetResourcePoolBucketGroupQoSInfoResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourcePoolBucketGroupQoSInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourcePoolBucketGroupQoSInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourcePoolBucketGroupQoSInfoResponse) GoString() string {
	return s.String()
}

func (s *GetResourcePoolBucketGroupQoSInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourcePoolBucketGroupQoSInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourcePoolBucketGroupQoSInfoResponse) GetBody() *GetResourcePoolBucketGroupQoSInfoResponseBody {
	return s.Body
}

func (s *GetResourcePoolBucketGroupQoSInfoResponse) SetHeaders(v map[string]*string) *GetResourcePoolBucketGroupQoSInfoResponse {
	s.Headers = v
	return s
}

func (s *GetResourcePoolBucketGroupQoSInfoResponse) SetStatusCode(v int32) *GetResourcePoolBucketGroupQoSInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourcePoolBucketGroupQoSInfoResponse) SetBody(v *GetResourcePoolBucketGroupQoSInfoResponseBody) *GetResourcePoolBucketGroupQoSInfoResponse {
	s.Body = v
	return s
}

func (s *GetResourcePoolBucketGroupQoSInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
