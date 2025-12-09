// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketTransferAccelerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketTransferAccelerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketTransferAccelerationResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketTransferAccelerationResponseBody) *GetBucketTransferAccelerationResponse
	GetBody() *GetBucketTransferAccelerationResponseBody
}

type GetBucketTransferAccelerationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketTransferAccelerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketTransferAccelerationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketTransferAccelerationResponse) GoString() string {
	return s.String()
}

func (s *GetBucketTransferAccelerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketTransferAccelerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketTransferAccelerationResponse) GetBody() *GetBucketTransferAccelerationResponseBody {
	return s.Body
}

func (s *GetBucketTransferAccelerationResponse) SetHeaders(v map[string]*string) *GetBucketTransferAccelerationResponse {
	s.Headers = v
	return s
}

func (s *GetBucketTransferAccelerationResponse) SetStatusCode(v int32) *GetBucketTransferAccelerationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketTransferAccelerationResponse) SetBody(v *GetBucketTransferAccelerationResponseBody) *GetBucketTransferAccelerationResponse {
	s.Body = v
	return s
}

func (s *GetBucketTransferAccelerationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
