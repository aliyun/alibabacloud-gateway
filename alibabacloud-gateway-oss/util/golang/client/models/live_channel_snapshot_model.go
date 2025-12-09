// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLiveChannelSnapshot interface {
	dara.Model
	String() string
	GoString() string
	SetDestBucket(v string) *LiveChannelSnapshot
	GetDestBucket() *string
	SetInterval(v int64) *LiveChannelSnapshot
	GetInterval() *int64
	SetNotifyTopic(v string) *LiveChannelSnapshot
	GetNotifyTopic() *string
	SetRoleName(v string) *LiveChannelSnapshot
	GetRoleName() *string
}

type LiveChannelSnapshot struct {
	DestBucket  *string `json:"DestBucket,omitempty" xml:"DestBucket,omitempty"`
	Interval    *int64  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	NotifyTopic *string `json:"NotifyTopic,omitempty" xml:"NotifyTopic,omitempty"`
	RoleName    *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s LiveChannelSnapshot) String() string {
	return dara.Prettify(s)
}

func (s LiveChannelSnapshot) GoString() string {
	return s.String()
}

func (s *LiveChannelSnapshot) GetDestBucket() *string {
	return s.DestBucket
}

func (s *LiveChannelSnapshot) GetInterval() *int64 {
	return s.Interval
}

func (s *LiveChannelSnapshot) GetNotifyTopic() *string {
	return s.NotifyTopic
}

func (s *LiveChannelSnapshot) GetRoleName() *string {
	return s.RoleName
}

func (s *LiveChannelSnapshot) SetDestBucket(v string) *LiveChannelSnapshot {
	s.DestBucket = &v
	return s
}

func (s *LiveChannelSnapshot) SetInterval(v int64) *LiveChannelSnapshot {
	s.Interval = &v
	return s
}

func (s *LiveChannelSnapshot) SetNotifyTopic(v string) *LiveChannelSnapshot {
	s.NotifyTopic = &v
	return s
}

func (s *LiveChannelSnapshot) SetRoleName(v string) *LiveChannelSnapshot {
	s.RoleName = &v
	return s
}

func (s *LiveChannelSnapshot) Validate() error {
	return dara.Validate(s)
}
