// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutUserDefinedLogFieldsConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutUserDefinedLogFieldsConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutUserDefinedLogFieldsConfigResponse
	GetStatusCode() *int32
}

type PutUserDefinedLogFieldsConfigResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutUserDefinedLogFieldsConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s PutUserDefinedLogFieldsConfigResponse) GoString() string {
	return s.String()
}

func (s *PutUserDefinedLogFieldsConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutUserDefinedLogFieldsConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutUserDefinedLogFieldsConfigResponse) SetHeaders(v map[string]*string) *PutUserDefinedLogFieldsConfigResponse {
	s.Headers = v
	return s
}

func (s *PutUserDefinedLogFieldsConfigResponse) SetStatusCode(v int32) *PutUserDefinedLogFieldsConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *PutUserDefinedLogFieldsConfigResponse) Validate() error {
	return dara.Validate(s)
}
