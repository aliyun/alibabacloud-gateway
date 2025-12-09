// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutChannelResponse
	GetStatusCode() *int32
}

type PutChannelResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s PutChannelResponse) GoString() string {
	return s.String()
}

func (s *PutChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutChannelResponse) SetHeaders(v map[string]*string) *PutChannelResponse {
	s.Headers = v
	return s
}

func (s *PutChannelResponse) SetStatusCode(v int32) *PutChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *PutChannelResponse) Validate() error {
	return dara.Validate(s)
}
