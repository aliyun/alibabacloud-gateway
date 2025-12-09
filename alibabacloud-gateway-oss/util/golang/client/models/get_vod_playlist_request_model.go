// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVodPlaylistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetVodPlaylistRequest
	GetEndTime() *string
	SetStartTime(v string) *GetVodPlaylistRequest
	GetStartTime() *string
}

type GetVodPlaylistRequest struct {
	// The end time of the time range during which the TS files that you want to query are generated in the Unix timestamp format.
	//
	// > The value of EndTime must be greater than the value of StartTime. The duration between EndTime and StartTime must be less than one day.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1636618271
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The start time of the time range during which the TS files that you want to query are generated in the Unix timestamp format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1636600271
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetVodPlaylistRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVodPlaylistRequest) GoString() string {
	return s.String()
}

func (s *GetVodPlaylistRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetVodPlaylistRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetVodPlaylistRequest) SetEndTime(v string) *GetVodPlaylistRequest {
	s.EndTime = &v
	return s
}

func (s *GetVodPlaylistRequest) SetStartTime(v string) *GetVodPlaylistRequest {
	s.StartTime = &v
	return s
}

func (s *GetVodPlaylistRequest) Validate() error {
	return dara.Validate(s)
}
