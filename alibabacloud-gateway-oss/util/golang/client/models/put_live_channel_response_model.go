// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutLiveChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutLiveChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutLiveChannelResponse
	GetStatusCode() *int32
	SetBody(v *PutLiveChannelResponseBody) *PutLiveChannelResponse
	GetBody() *PutLiveChannelResponseBody
}

type PutLiveChannelResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutLiveChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutLiveChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s PutLiveChannelResponse) GoString() string {
	return s.String()
}

func (s *PutLiveChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutLiveChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutLiveChannelResponse) GetBody() *PutLiveChannelResponseBody {
	return s.Body
}

func (s *PutLiveChannelResponse) SetHeaders(v map[string]*string) *PutLiveChannelResponse {
	s.Headers = v
	return s
}

func (s *PutLiveChannelResponse) SetStatusCode(v int32) *PutLiveChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *PutLiveChannelResponse) SetBody(v *PutLiveChannelResponseBody) *PutLiveChannelResponse {
	s.Body = v
	return s
}

func (s *PutLiveChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
