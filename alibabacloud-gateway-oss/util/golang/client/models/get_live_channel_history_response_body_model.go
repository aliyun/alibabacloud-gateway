// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveChannelHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLiveChannelHistory(v *GetLiveChannelHistoryResponseBodyLiveChannelHistory) *GetLiveChannelHistoryResponseBody
	GetLiveChannelHistory() *GetLiveChannelHistoryResponseBodyLiveChannelHistory
}

type GetLiveChannelHistoryResponseBody struct {
	// The container that stores the returned results of the GetLiveChannelHistory request.
	LiveChannelHistory *GetLiveChannelHistoryResponseBodyLiveChannelHistory `json:"LiveChannelHistory,omitempty" xml:"LiveChannelHistory,omitempty" type:"Struct"`
}

func (s GetLiveChannelHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLiveChannelHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetLiveChannelHistoryResponseBody) GetLiveChannelHistory() *GetLiveChannelHistoryResponseBodyLiveChannelHistory {
	return s.LiveChannelHistory
}

func (s *GetLiveChannelHistoryResponseBody) SetLiveChannelHistory(v *GetLiveChannelHistoryResponseBodyLiveChannelHistory) *GetLiveChannelHistoryResponseBody {
	s.LiveChannelHistory = v
	return s
}

func (s *GetLiveChannelHistoryResponseBody) Validate() error {
	if s.LiveChannelHistory != nil {
		if err := s.LiveChannelHistory.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLiveChannelHistoryResponseBodyLiveChannelHistory struct {
	// The container that stores a list of stream pushing records.
	LiveRecord []*LiveRecord `json:"LiveRecord,omitempty" xml:"LiveRecord,omitempty" type:"Repeated"`
}

func (s GetLiveChannelHistoryResponseBodyLiveChannelHistory) String() string {
	return dara.Prettify(s)
}

func (s GetLiveChannelHistoryResponseBodyLiveChannelHistory) GoString() string {
	return s.String()
}

func (s *GetLiveChannelHistoryResponseBodyLiveChannelHistory) GetLiveRecord() []*LiveRecord {
	return s.LiveRecord
}

func (s *GetLiveChannelHistoryResponseBodyLiveChannelHistory) SetLiveRecord(v []*LiveRecord) *GetLiveChannelHistoryResponseBodyLiveChannelHistory {
	s.LiveRecord = v
	return s
}

func (s *GetLiveChannelHistoryResponseBodyLiveChannelHistory) Validate() error {
	if s.LiveRecord != nil {
		for _, item := range s.LiveRecord {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
