// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveChannelInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLiveChannelConfiguration(v *GetLiveChannelInfoResponseBodyLiveChannelConfiguration) *GetLiveChannelInfoResponseBody
	GetLiveChannelConfiguration() *GetLiveChannelInfoResponseBodyLiveChannelConfiguration
}

type GetLiveChannelInfoResponseBody struct {
	// The container that stores the returned results of the GetLiveChannelInfo request.
	LiveChannelConfiguration *GetLiveChannelInfoResponseBodyLiveChannelConfiguration `json:"LiveChannelConfiguration,omitempty" xml:"LiveChannelConfiguration,omitempty" type:"Struct"`
}

func (s GetLiveChannelInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLiveChannelInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetLiveChannelInfoResponseBody) GetLiveChannelConfiguration() *GetLiveChannelInfoResponseBodyLiveChannelConfiguration {
	return s.LiveChannelConfiguration
}

func (s *GetLiveChannelInfoResponseBody) SetLiveChannelConfiguration(v *GetLiveChannelInfoResponseBodyLiveChannelConfiguration) *GetLiveChannelInfoResponseBody {
	s.LiveChannelConfiguration = v
	return s
}

func (s *GetLiveChannelInfoResponseBody) Validate() error {
	if s.LiveChannelConfiguration != nil {
		if err := s.LiveChannelConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLiveChannelInfoResponseBodyLiveChannelConfiguration struct {
	// The description of the LiveChannel.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The status of the LiveChannel.
	//
	// Valid values:
	//
	// - enabled: indicates that the LiveChannel is enabled.
	//
	// - disabled: indicates that the LiveChannel is disabled.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The container that stores the configurations used by the LiveChannel to store uploaded data.
	//
	// > FragDuration, FragCount, and PlaylistName are returned only when the value of Type is HLS.
	Target *LiveChannelTarget `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s GetLiveChannelInfoResponseBodyLiveChannelConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetLiveChannelInfoResponseBodyLiveChannelConfiguration) GoString() string {
	return s.String()
}

func (s *GetLiveChannelInfoResponseBodyLiveChannelConfiguration) GetDescription() *string {
	return s.Description
}

func (s *GetLiveChannelInfoResponseBodyLiveChannelConfiguration) GetStatus() *string {
	return s.Status
}

func (s *GetLiveChannelInfoResponseBodyLiveChannelConfiguration) GetTarget() *LiveChannelTarget {
	return s.Target
}

func (s *GetLiveChannelInfoResponseBodyLiveChannelConfiguration) SetDescription(v string) *GetLiveChannelInfoResponseBodyLiveChannelConfiguration {
	s.Description = &v
	return s
}

func (s *GetLiveChannelInfoResponseBodyLiveChannelConfiguration) SetStatus(v string) *GetLiveChannelInfoResponseBodyLiveChannelConfiguration {
	s.Status = &v
	return s
}

func (s *GetLiveChannelInfoResponseBodyLiveChannelConfiguration) SetTarget(v *LiveChannelTarget) *GetLiveChannelInfoResponseBodyLiveChannelConfiguration {
	s.Target = v
	return s
}

func (s *GetLiveChannelInfoResponseBodyLiveChannelConfiguration) Validate() error {
	if s.Target != nil {
		if err := s.Target.Validate(); err != nil {
			return err
		}
	}
	return nil
}
