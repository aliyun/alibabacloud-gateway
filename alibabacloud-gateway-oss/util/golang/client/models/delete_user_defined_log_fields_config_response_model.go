// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserDefinedLogFieldsConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUserDefinedLogFieldsConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUserDefinedLogFieldsConfigResponse
	GetStatusCode() *int32
}

type DeleteUserDefinedLogFieldsConfigResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteUserDefinedLogFieldsConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserDefinedLogFieldsConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserDefinedLogFieldsConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUserDefinedLogFieldsConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUserDefinedLogFieldsConfigResponse) SetHeaders(v map[string]*string) *DeleteUserDefinedLogFieldsConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserDefinedLogFieldsConfigResponse) SetStatusCode(v int32) *DeleteUserDefinedLogFieldsConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserDefinedLogFieldsConfigResponse) Validate() error {
	return dara.Validate(s)
}
