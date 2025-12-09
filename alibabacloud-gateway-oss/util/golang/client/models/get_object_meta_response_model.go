// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetObjectMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetObjectMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetObjectMetaResponse
	GetStatusCode() *int32
}

type GetObjectMetaResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s GetObjectMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetObjectMetaResponse) GoString() string {
	return s.String()
}

func (s *GetObjectMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetObjectMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetObjectMetaResponse) SetHeaders(v map[string]*string) *GetObjectMetaResponse {
	s.Headers = v
	return s
}

func (s *GetObjectMetaResponse) SetStatusCode(v int32) *GetObjectMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetObjectMetaResponse) Validate() error {
	return dara.Validate(s)
}
