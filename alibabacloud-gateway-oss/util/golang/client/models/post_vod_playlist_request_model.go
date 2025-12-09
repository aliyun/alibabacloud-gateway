// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostVodPlaylistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *PostVodPlaylistRequest
	GetEndTime() *string
	SetStartTime(v string) *PostVodPlaylistRequest
	GetStartTime() *string
}

type PostVodPlaylistRequest struct {
	// The end time of the time range during which the TS files that you want to query are generated,
	//
	// which is a Unix timestamp.
	//
	// > The value of EndTime must be later than the value of StartTime. The duration between EndTime and StartTime must be shorter than one day.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1636618271
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The start time of the time range during which the TS files that you want to query are generated, which is a Unix timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1636600271
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s PostVodPlaylistRequest) String() string {
	return dara.Prettify(s)
}

func (s PostVodPlaylistRequest) GoString() string {
	return s.String()
}

func (s *PostVodPlaylistRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *PostVodPlaylistRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *PostVodPlaylistRequest) SetEndTime(v string) *PostVodPlaylistRequest {
	s.EndTime = &v
	return s
}

func (s *PostVodPlaylistRequest) SetStartTime(v string) *PostVodPlaylistRequest {
	s.StartTime = &v
	return s
}

func (s *PostVodPlaylistRequest) Validate() error {
	return dara.Validate(s)
}
