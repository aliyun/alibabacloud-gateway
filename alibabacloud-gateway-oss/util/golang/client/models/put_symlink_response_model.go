// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutSymlinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutSymlinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutSymlinkResponse
	GetStatusCode() *int32
}

type PutSymlinkResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutSymlinkResponse) String() string {
	return dara.Prettify(s)
}

func (s PutSymlinkResponse) GoString() string {
	return s.String()
}

func (s *PutSymlinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutSymlinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutSymlinkResponse) SetHeaders(v map[string]*string) *PutSymlinkResponse {
	s.Headers = v
	return s
}

func (s *PutSymlinkResponse) SetStatusCode(v int32) *PutSymlinkResponse {
	s.StatusCode = &v
	return s
}

func (s *PutSymlinkResponse) Validate() error {
	return dara.Validate(s)
}
