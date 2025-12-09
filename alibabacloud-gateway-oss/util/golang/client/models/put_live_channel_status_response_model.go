// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutLiveChannelStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutLiveChannelStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutLiveChannelStatusResponse
	GetStatusCode() *int32
}

type PutLiveChannelStatusResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PutLiveChannelStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s PutLiveChannelStatusResponse) GoString() string {
	return s.String()
}

func (s *PutLiveChannelStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutLiveChannelStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutLiveChannelStatusResponse) SetHeaders(v map[string]*string) *PutLiveChannelStatusResponse {
	s.Headers = v
	return s
}

func (s *PutLiveChannelStatusResponse) SetStatusCode(v int32) *PutLiveChannelStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *PutLiveChannelStatusResponse) Validate() error {
	return dara.Validate(s)
}
