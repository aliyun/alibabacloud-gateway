// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveChannelResponse
	GetStatusCode() *int32
}

type DeleteLiveChannelResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteLiveChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveChannelResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveChannelResponse) SetHeaders(v map[string]*string) *DeleteLiveChannelResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveChannelResponse) SetStatusCode(v int32) *DeleteLiveChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveChannelResponse) Validate() error {
	return dara.Validate(s)
}
