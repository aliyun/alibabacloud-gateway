// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSymlinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSymlinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSymlinkResponse
	GetStatusCode() *int32
}

type GetSymlinkResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s GetSymlinkResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSymlinkResponse) GoString() string {
	return s.String()
}

func (s *GetSymlinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSymlinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSymlinkResponse) SetHeaders(v map[string]*string) *GetSymlinkResponse {
	s.Headers = v
	return s
}

func (s *GetSymlinkResponse) SetStatusCode(v int32) *GetSymlinkResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSymlinkResponse) Validate() error {
	return dara.Validate(s)
}
