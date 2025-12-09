// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListLiveChannelResult(v *ListLiveChannelResponseBodyListLiveChannelResult) *ListLiveChannelResponseBody
	GetListLiveChannelResult() *ListLiveChannelResponseBodyListLiveChannelResult
}

type ListLiveChannelResponseBody struct {
	// The container that stores the results of the ListLiveChannel request.
	ListLiveChannelResult *ListLiveChannelResponseBodyListLiveChannelResult `json:"ListLiveChannelResult,omitempty" xml:"ListLiveChannelResult,omitempty" type:"Struct"`
}

func (s ListLiveChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLiveChannelResponseBody) GoString() string {
	return s.String()
}

func (s *ListLiveChannelResponseBody) GetListLiveChannelResult() *ListLiveChannelResponseBodyListLiveChannelResult {
	return s.ListLiveChannelResult
}

func (s *ListLiveChannelResponseBody) SetListLiveChannelResult(v *ListLiveChannelResponseBodyListLiveChannelResult) *ListLiveChannelResponseBody {
	s.ListLiveChannelResult = v
	return s
}

func (s *ListLiveChannelResponseBody) Validate() error {
	if s.ListLiveChannelResult != nil {
		if err := s.ListLiveChannelResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListLiveChannelResponseBodyListLiveChannelResult struct {
	// Indicates whether all results are returned.
	//
	// - true: All results are returned.
	//
	// - false: Not all results are returned.
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The container that stores the information about each returned LiveChannel.
	LiveChannels []*LiveChannel `json:"LiveChannel,omitempty" xml:"LiveChannel,omitempty" type:"Repeated"`
	// The name of the LiveChannel after which the ListLiveChannel operation starts.
	//
	// example:
	//
	// new
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The maximum number of returned LiveChannels in the response.
	//
	// example:
	//
	// 20
	MaxKeys *int64 `json:"MaxKeys,omitempty" xml:"MaxKeys,omitempty"`
	// If not all results are returned, the NextMarker parameter is included in the response to indicate the Marker value of the next request.
	//
	// example:
	//
	// channel-0
	NextMarker *string `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	// The prefix that the names of the returned LiveChannels contain.
	//
	// example:
	//
	// Channel
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
}

func (s ListLiveChannelResponseBodyListLiveChannelResult) String() string {
	return dara.Prettify(s)
}

func (s ListLiveChannelResponseBodyListLiveChannelResult) GoString() string {
	return s.String()
}

func (s *ListLiveChannelResponseBodyListLiveChannelResult) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListLiveChannelResponseBodyListLiveChannelResult) GetLiveChannels() []*LiveChannel {
	return s.LiveChannels
}

func (s *ListLiveChannelResponseBodyListLiveChannelResult) GetMarker() *string {
	return s.Marker
}

func (s *ListLiveChannelResponseBodyListLiveChannelResult) GetMaxKeys() *int64 {
	return s.MaxKeys
}

func (s *ListLiveChannelResponseBodyListLiveChannelResult) GetNextMarker() *string {
	return s.NextMarker
}

func (s *ListLiveChannelResponseBodyListLiveChannelResult) GetPrefix() *string {
	return s.Prefix
}

func (s *ListLiveChannelResponseBodyListLiveChannelResult) SetIsTruncated(v bool) *ListLiveChannelResponseBodyListLiveChannelResult {
	s.IsTruncated = &v
	return s
}

func (s *ListLiveChannelResponseBodyListLiveChannelResult) SetLiveChannels(v []*LiveChannel) *ListLiveChannelResponseBodyListLiveChannelResult {
	s.LiveChannels = v
	return s
}

func (s *ListLiveChannelResponseBodyListLiveChannelResult) SetMarker(v string) *ListLiveChannelResponseBodyListLiveChannelResult {
	s.Marker = &v
	return s
}

func (s *ListLiveChannelResponseBodyListLiveChannelResult) SetMaxKeys(v int64) *ListLiveChannelResponseBodyListLiveChannelResult {
	s.MaxKeys = &v
	return s
}

func (s *ListLiveChannelResponseBodyListLiveChannelResult) SetNextMarker(v string) *ListLiveChannelResponseBodyListLiveChannelResult {
	s.NextMarker = &v
	return s
}

func (s *ListLiveChannelResponseBodyListLiveChannelResult) SetPrefix(v string) *ListLiveChannelResponseBodyListLiveChannelResult {
	s.Prefix = &v
	return s
}

func (s *ListLiveChannelResponseBodyListLiveChannelResult) Validate() error {
	if s.LiveChannels != nil {
		for _, item := range s.LiveChannels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
