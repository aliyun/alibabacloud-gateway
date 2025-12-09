// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChannel(v *GetChannelResult) *GetChannelResponseBody
	GetChannel() *GetChannelResult
}

type GetChannelResponseBody struct {
	Channel *GetChannelResult `json:"channel,omitempty" xml:"channel,omitempty"`
}

func (s GetChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetChannelResponseBody) GoString() string {
	return s.String()
}

func (s *GetChannelResponseBody) GetChannel() *GetChannelResult {
	return s.Channel
}

func (s *GetChannelResponseBody) SetChannel(v *GetChannelResult) *GetChannelResponseBody {
	s.Channel = v
	return s
}

func (s *GetChannelResponseBody) Validate() error {
	if s.Channel != nil {
		if err := s.Channel.Validate(); err != nil {
			return err
		}
	}
	return nil
}
