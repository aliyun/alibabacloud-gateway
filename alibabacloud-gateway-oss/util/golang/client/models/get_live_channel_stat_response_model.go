// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveChannelStatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLiveChannelStatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLiveChannelStatResponse
	GetStatusCode() *int32
	SetBody(v *GetLiveChannelStatResponseBody) *GetLiveChannelStatResponse
	GetBody() *GetLiveChannelStatResponseBody
}

type GetLiveChannelStatResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLiveChannelStatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLiveChannelStatResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLiveChannelStatResponse) GoString() string {
	return s.String()
}

func (s *GetLiveChannelStatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLiveChannelStatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLiveChannelStatResponse) GetBody() *GetLiveChannelStatResponseBody {
	return s.Body
}

func (s *GetLiveChannelStatResponse) SetHeaders(v map[string]*string) *GetLiveChannelStatResponse {
	s.Headers = v
	return s
}

func (s *GetLiveChannelStatResponse) SetStatusCode(v int32) *GetLiveChannelStatResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLiveChannelStatResponse) SetBody(v *GetLiveChannelStatResponseBody) *GetLiveChannelStatResponse {
	s.Body = v
	return s
}

func (s *GetLiveChannelStatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
