// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveChannelHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLiveChannelHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLiveChannelHistoryResponse
	GetStatusCode() *int32
	SetBody(v *GetLiveChannelHistoryResponseBody) *GetLiveChannelHistoryResponse
	GetBody() *GetLiveChannelHistoryResponseBody
}

type GetLiveChannelHistoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLiveChannelHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLiveChannelHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLiveChannelHistoryResponse) GoString() string {
	return s.String()
}

func (s *GetLiveChannelHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLiveChannelHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLiveChannelHistoryResponse) GetBody() *GetLiveChannelHistoryResponseBody {
	return s.Body
}

func (s *GetLiveChannelHistoryResponse) SetHeaders(v map[string]*string) *GetLiveChannelHistoryResponse {
	s.Headers = v
	return s
}

func (s *GetLiveChannelHistoryResponse) SetStatusCode(v int32) *GetLiveChannelHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLiveChannelHistoryResponse) SetBody(v *GetLiveChannelHistoryResponseBody) *GetLiveChannelHistoryResponse {
	s.Body = v
	return s
}

func (s *GetLiveChannelHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
