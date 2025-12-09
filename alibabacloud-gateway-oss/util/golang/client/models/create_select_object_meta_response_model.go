// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSelectObjectMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSelectObjectMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSelectObjectMetaResponse
	GetStatusCode() *int32
	SetBody(v *SelectMetaStatus) *CreateSelectObjectMetaResponse
	GetBody() *SelectMetaStatus
}

type CreateSelectObjectMetaResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SelectMetaStatus  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSelectObjectMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSelectObjectMetaResponse) GoString() string {
	return s.String()
}

func (s *CreateSelectObjectMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSelectObjectMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSelectObjectMetaResponse) GetBody() *SelectMetaStatus {
	return s.Body
}

func (s *CreateSelectObjectMetaResponse) SetHeaders(v map[string]*string) *CreateSelectObjectMetaResponse {
	s.Headers = v
	return s
}

func (s *CreateSelectObjectMetaResponse) SetStatusCode(v int32) *CreateSelectObjectMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSelectObjectMetaResponse) SetBody(v *SelectMetaStatus) *CreateSelectObjectMetaResponse {
	s.Body = v
	return s
}

func (s *CreateSelectObjectMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
