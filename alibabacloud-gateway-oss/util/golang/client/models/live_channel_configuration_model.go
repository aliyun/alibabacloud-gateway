// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLiveChannelConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *LiveChannelConfiguration
	GetDescription() *string
	SetSnapshot(v *LiveChannelSnapshot) *LiveChannelConfiguration
	GetSnapshot() *LiveChannelSnapshot
	SetStatus(v string) *LiveChannelConfiguration
	GetStatus() *string
	SetTarget(v *LiveChannelTarget) *LiveChannelConfiguration
	GetTarget() *LiveChannelTarget
}

type LiveChannelConfiguration struct {
	Description *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	Snapshot    *LiveChannelSnapshot `json:"Snapshot,omitempty" xml:"Snapshot,omitempty"`
	Status      *string              `json:"Status,omitempty" xml:"Status,omitempty"`
	Target      *LiveChannelTarget   `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s LiveChannelConfiguration) String() string {
	return dara.Prettify(s)
}

func (s LiveChannelConfiguration) GoString() string {
	return s.String()
}

func (s *LiveChannelConfiguration) GetDescription() *string {
	return s.Description
}

func (s *LiveChannelConfiguration) GetSnapshot() *LiveChannelSnapshot {
	return s.Snapshot
}

func (s *LiveChannelConfiguration) GetStatus() *string {
	return s.Status
}

func (s *LiveChannelConfiguration) GetTarget() *LiveChannelTarget {
	return s.Target
}

func (s *LiveChannelConfiguration) SetDescription(v string) *LiveChannelConfiguration {
	s.Description = &v
	return s
}

func (s *LiveChannelConfiguration) SetSnapshot(v *LiveChannelSnapshot) *LiveChannelConfiguration {
	s.Snapshot = v
	return s
}

func (s *LiveChannelConfiguration) SetStatus(v string) *LiveChannelConfiguration {
	s.Status = &v
	return s
}

func (s *LiveChannelConfiguration) SetTarget(v *LiveChannelTarget) *LiveChannelConfiguration {
	s.Target = v
	return s
}

func (s *LiveChannelConfiguration) Validate() error {
	if s.Snapshot != nil {
		if err := s.Snapshot.Validate(); err != nil {
			return err
		}
	}
	if s.Target != nil {
		if err := s.Target.Validate(); err != nil {
			return err
		}
	}
	return nil
}
