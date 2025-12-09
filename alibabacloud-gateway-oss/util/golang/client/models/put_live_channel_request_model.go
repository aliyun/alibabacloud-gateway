// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutLiveChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLiveChannelConfiguration(v *LiveChannelConfiguration) *PutLiveChannelRequest
	GetLiveChannelConfiguration() *LiveChannelConfiguration
}

type PutLiveChannelRequest struct {
	// The container that stores the configurations of the LiveChannel.
	LiveChannelConfiguration *LiveChannelConfiguration `json:"LiveChannelConfiguration,omitempty" xml:"LiveChannelConfiguration,omitempty"`
}

func (s PutLiveChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s PutLiveChannelRequest) GoString() string {
	return s.String()
}

func (s *PutLiveChannelRequest) GetLiveChannelConfiguration() *LiveChannelConfiguration {
	return s.LiveChannelConfiguration
}

func (s *PutLiveChannelRequest) SetLiveChannelConfiguration(v *LiveChannelConfiguration) *PutLiveChannelRequest {
	s.LiveChannelConfiguration = v
	return s
}

func (s *PutLiveChannelRequest) Validate() error {
	if s.LiveChannelConfiguration != nil {
		if err := s.LiveChannelConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
