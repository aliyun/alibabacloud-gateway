// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenMetaQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenMetaQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenMetaQueryResponse
	GetStatusCode() *int32
}

type OpenMetaQueryResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s OpenMetaQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenMetaQueryResponse) GoString() string {
	return s.String()
}

func (s *OpenMetaQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenMetaQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenMetaQueryResponse) SetHeaders(v map[string]*string) *OpenMetaQueryResponse {
	s.Headers = v
	return s
}

func (s *OpenMetaQueryResponse) SetStatusCode(v int32) *OpenMetaQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenMetaQueryResponse) Validate() error {
	return dara.Validate(s)
}
