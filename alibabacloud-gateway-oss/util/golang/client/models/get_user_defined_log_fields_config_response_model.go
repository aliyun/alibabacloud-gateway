// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserDefinedLogFieldsConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserDefinedLogFieldsConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserDefinedLogFieldsConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetUserDefinedLogFieldsConfigResponseBody) *GetUserDefinedLogFieldsConfigResponse
	GetBody() *GetUserDefinedLogFieldsConfigResponseBody
}

type GetUserDefinedLogFieldsConfigResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserDefinedLogFieldsConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserDefinedLogFieldsConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserDefinedLogFieldsConfigResponse) GoString() string {
	return s.String()
}

func (s *GetUserDefinedLogFieldsConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserDefinedLogFieldsConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserDefinedLogFieldsConfigResponse) GetBody() *GetUserDefinedLogFieldsConfigResponseBody {
	return s.Body
}

func (s *GetUserDefinedLogFieldsConfigResponse) SetHeaders(v map[string]*string) *GetUserDefinedLogFieldsConfigResponse {
	s.Headers = v
	return s
}

func (s *GetUserDefinedLogFieldsConfigResponse) SetStatusCode(v int32) *GetUserDefinedLogFieldsConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserDefinedLogFieldsConfigResponse) SetBody(v *GetUserDefinedLogFieldsConfigResponseBody) *GetUserDefinedLogFieldsConfigResponse {
	s.Body = v
	return s
}

func (s *GetUserDefinedLogFieldsConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
