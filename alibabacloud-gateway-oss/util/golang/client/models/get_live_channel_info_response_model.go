// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveChannelInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLiveChannelInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLiveChannelInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetLiveChannelInfoResponseBody) *GetLiveChannelInfoResponse
	GetBody() *GetLiveChannelInfoResponseBody
}

type GetLiveChannelInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLiveChannelInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLiveChannelInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLiveChannelInfoResponse) GoString() string {
	return s.String()
}

func (s *GetLiveChannelInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLiveChannelInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLiveChannelInfoResponse) GetBody() *GetLiveChannelInfoResponseBody {
	return s.Body
}

func (s *GetLiveChannelInfoResponse) SetHeaders(v map[string]*string) *GetLiveChannelInfoResponse {
	s.Headers = v
	return s
}

func (s *GetLiveChannelInfoResponse) SetStatusCode(v int32) *GetLiveChannelInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLiveChannelInfoResponse) SetBody(v *GetLiveChannelInfoResponseBody) *GetLiveChannelInfoResponse {
	s.Body = v
	return s
}

func (s *GetLiveChannelInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
