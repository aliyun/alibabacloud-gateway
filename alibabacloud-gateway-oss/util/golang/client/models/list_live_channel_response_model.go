// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLiveChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLiveChannelResponse
	GetStatusCode() *int32
	SetBody(v *ListLiveChannelResponseBody) *ListLiveChannelResponse
	GetBody() *ListLiveChannelResponseBody
}

type ListLiveChannelResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLiveChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLiveChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLiveChannelResponse) GoString() string {
	return s.String()
}

func (s *ListLiveChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLiveChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLiveChannelResponse) GetBody() *ListLiveChannelResponseBody {
	return s.Body
}

func (s *ListLiveChannelResponse) SetHeaders(v map[string]*string) *ListLiveChannelResponse {
	s.Headers = v
	return s
}

func (s *ListLiveChannelResponse) SetStatusCode(v int32) *ListLiveChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLiveChannelResponse) SetBody(v *ListLiveChannelResponseBody) *ListLiveChannelResponse {
	s.Body = v
	return s
}

func (s *ListLiveChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
