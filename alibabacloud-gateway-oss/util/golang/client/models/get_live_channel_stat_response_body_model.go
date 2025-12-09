// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveChannelStatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLiveChannelStat(v *GetLiveChannelStatResponseBodyLiveChannelStat) *GetLiveChannelStatResponseBody
	GetLiveChannelStat() *GetLiveChannelStatResponseBodyLiveChannelStat
}

type GetLiveChannelStatResponseBody struct {
	// The container that stores the returned results of the GetLiveChannelStat request.
	LiveChannelStat *GetLiveChannelStatResponseBodyLiveChannelStat `json:"LiveChannelStat,omitempty" xml:"LiveChannelStat,omitempty" type:"Struct"`
}

func (s GetLiveChannelStatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLiveChannelStatResponseBody) GoString() string {
	return s.String()
}

func (s *GetLiveChannelStatResponseBody) GetLiveChannelStat() *GetLiveChannelStatResponseBodyLiveChannelStat {
	return s.LiveChannelStat
}

func (s *GetLiveChannelStatResponseBody) SetLiveChannelStat(v *GetLiveChannelStatResponseBodyLiveChannelStat) *GetLiveChannelStatResponseBody {
	s.LiveChannelStat = v
	return s
}

func (s *GetLiveChannelStatResponseBody) Validate() error {
	if s.LiveChannelStat != nil {
		if err := s.LiveChannelStat.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLiveChannelStatResponseBodyLiveChannelStat struct {
	// The container that stores audio stream information if Status is set to Live.
	//
	// >Video and audio containers can be returned only if Status is set to Live. However, these two containers may not necessarily be returned if Status is set to Live. For example, if the client has connected to the LiveChannel but no audio or video stream is sent, these two containers are not returned.
	Audio *LiveChannelAudio `json:"Audio,omitempty" xml:"Audio,omitempty"`
	// If Status is set to Live, this element indicates the time when the current client starts to ingest streams. The value of the element is in the ISO 8601 format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2016-08-25T06:25:15.000Z
	ConnectedTime *string `json:"ConnectedTime,omitempty" xml:"ConnectedTime,omitempty"`
	// If Status is set to Live, this element indicates the IP address of the current client that ingests streams.
	//
	// example:
	//
	// 10.1.2.3:47745
	RemoteAddr *string `json:"RemoteAddr,omitempty" xml:"RemoteAddr,omitempty"`
	// The current stream ingestion status of the LiveChannel. Valid value: Disabled、Live、Idle。
	//
	// example:
	//
	// Live
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The container that stores video stream information if Status is set to Live.
	//
	// >Video and audio containers can be returned only if Status is set to Live. However, these two containers may not necessarily be returned if Status is set to Live. For example, if the client has connected to the LiveChannel but no audio or video stream is sent, these two containers are not returned.
	Video *LiveChannelVideo `json:"Video,omitempty" xml:"Video,omitempty"`
}

func (s GetLiveChannelStatResponseBodyLiveChannelStat) String() string {
	return dara.Prettify(s)
}

func (s GetLiveChannelStatResponseBodyLiveChannelStat) GoString() string {
	return s.String()
}

func (s *GetLiveChannelStatResponseBodyLiveChannelStat) GetAudio() *LiveChannelAudio {
	return s.Audio
}

func (s *GetLiveChannelStatResponseBodyLiveChannelStat) GetConnectedTime() *string {
	return s.ConnectedTime
}

func (s *GetLiveChannelStatResponseBodyLiveChannelStat) GetRemoteAddr() *string {
	return s.RemoteAddr
}

func (s *GetLiveChannelStatResponseBodyLiveChannelStat) GetStatus() *string {
	return s.Status
}

func (s *GetLiveChannelStatResponseBodyLiveChannelStat) GetVideo() *LiveChannelVideo {
	return s.Video
}

func (s *GetLiveChannelStatResponseBodyLiveChannelStat) SetAudio(v *LiveChannelAudio) *GetLiveChannelStatResponseBodyLiveChannelStat {
	s.Audio = v
	return s
}

func (s *GetLiveChannelStatResponseBodyLiveChannelStat) SetConnectedTime(v string) *GetLiveChannelStatResponseBodyLiveChannelStat {
	s.ConnectedTime = &v
	return s
}

func (s *GetLiveChannelStatResponseBodyLiveChannelStat) SetRemoteAddr(v string) *GetLiveChannelStatResponseBodyLiveChannelStat {
	s.RemoteAddr = &v
	return s
}

func (s *GetLiveChannelStatResponseBodyLiveChannelStat) SetStatus(v string) *GetLiveChannelStatResponseBodyLiveChannelStat {
	s.Status = &v
	return s
}

func (s *GetLiveChannelStatResponseBodyLiveChannelStat) SetVideo(v *LiveChannelVideo) *GetLiveChannelStatResponseBodyLiveChannelStat {
	s.Video = v
	return s
}

func (s *GetLiveChannelStatResponseBodyLiveChannelStat) Validate() error {
	if s.Audio != nil {
		if err := s.Audio.Validate(); err != nil {
			return err
		}
	}
	if s.Video != nil {
		if err := s.Video.Validate(); err != nil {
			return err
		}
	}
	return nil
}
