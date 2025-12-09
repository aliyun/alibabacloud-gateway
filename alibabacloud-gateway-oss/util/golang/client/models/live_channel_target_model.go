// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLiveChannelTarget interface {
	dara.Model
	String() string
	GoString() string
	SetFragCount(v int64) *LiveChannelTarget
	GetFragCount() *int64
	SetFragDuration(v int64) *LiveChannelTarget
	GetFragDuration() *int64
	SetPlaylistName(v string) *LiveChannelTarget
	GetPlaylistName() *string
	SetType(v string) *LiveChannelTarget
	GetType() *string
}

type LiveChannelTarget struct {
	FragCount    *int64  `json:"FragCount,omitempty" xml:"FragCount,omitempty"`
	FragDuration *int64  `json:"FragDuration,omitempty" xml:"FragDuration,omitempty"`
	PlaylistName *string `json:"PlaylistName,omitempty" xml:"PlaylistName,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s LiveChannelTarget) String() string {
	return dara.Prettify(s)
}

func (s LiveChannelTarget) GoString() string {
	return s.String()
}

func (s *LiveChannelTarget) GetFragCount() *int64 {
	return s.FragCount
}

func (s *LiveChannelTarget) GetFragDuration() *int64 {
	return s.FragDuration
}

func (s *LiveChannelTarget) GetPlaylistName() *string {
	return s.PlaylistName
}

func (s *LiveChannelTarget) GetType() *string {
	return s.Type
}

func (s *LiveChannelTarget) SetFragCount(v int64) *LiveChannelTarget {
	s.FragCount = &v
	return s
}

func (s *LiveChannelTarget) SetFragDuration(v int64) *LiveChannelTarget {
	s.FragDuration = &v
	return s
}

func (s *LiveChannelTarget) SetPlaylistName(v string) *LiveChannelTarget {
	s.PlaylistName = &v
	return s
}

func (s *LiveChannelTarget) SetType(v string) *LiveChannelTarget {
	s.Type = &v
	return s
}

func (s *LiveChannelTarget) Validate() error {
	return dara.Validate(s)
}
