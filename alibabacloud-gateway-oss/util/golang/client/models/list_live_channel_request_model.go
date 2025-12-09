// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMarker(v string) *ListLiveChannelRequest
	GetMarker() *string
	SetMaxKeys(v int64) *ListLiveChannelRequest
	GetMaxKeys() *int64
	SetPrefix(v string) *ListLiveChannelRequest
	GetPrefix() *string
}

type ListLiveChannelRequest struct {
	// The name of the LiveChannel from which the list operation starts. LiveChannels whose names are alphabetically after the value of the marker parameter are returned.
	//
	// example:
	//
	// channel-1
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of LiveChannels that can be returned for the current request. The value of max-keys cannot exceed 1000.
	//
	// Default value: 100.
	//
	// example:
	//
	// 100
	MaxKeys *int64 `json:"max-keys,omitempty" xml:"max-keys,omitempty"`
	// The prefix that the names of the LiveChannels that you want to return must contain. If you specify a prefix in the request, the specified prefix is included in the response.
	//
	// example:
	//
	// fun/
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
}

func (s ListLiveChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLiveChannelRequest) GoString() string {
	return s.String()
}

func (s *ListLiveChannelRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListLiveChannelRequest) GetMaxKeys() *int64 {
	return s.MaxKeys
}

func (s *ListLiveChannelRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *ListLiveChannelRequest) SetMarker(v string) *ListLiveChannelRequest {
	s.Marker = &v
	return s
}

func (s *ListLiveChannelRequest) SetMaxKeys(v int64) *ListLiveChannelRequest {
	s.MaxKeys = &v
	return s
}

func (s *ListLiveChannelRequest) SetPrefix(v string) *ListLiveChannelRequest {
	s.Prefix = &v
	return s
}

func (s *ListLiveChannelRequest) Validate() error {
	return dara.Validate(s)
}
