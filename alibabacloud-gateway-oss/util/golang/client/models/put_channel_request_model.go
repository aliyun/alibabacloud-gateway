// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannel(v *Channel) *PutChannelRequest
	GetChannel() *Channel
}

type PutChannelRequest struct {
	// Container for storing image processing channel configuration
	Channel *Channel `json:"Channel,omitempty" xml:"Channel,omitempty"`
}

func (s PutChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s PutChannelRequest) GoString() string {
	return s.String()
}

func (s *PutChannelRequest) GetChannel() *Channel {
	return s.Channel
}

func (s *PutChannelRequest) SetChannel(v *Channel) *PutChannelRequest {
	s.Channel = v
	return s
}

func (s *PutChannelRequest) Validate() error {
	if s.Channel != nil {
		if err := s.Channel.Validate(); err != nil {
			return err
		}
	}
	return nil
}
