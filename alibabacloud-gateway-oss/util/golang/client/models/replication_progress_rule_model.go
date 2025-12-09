// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplicationProgressRule interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *ReplicationProgressRule
	GetAction() *string
	SetDestination(v *ReplicationDestination) *ReplicationProgressRule
	GetDestination() *ReplicationDestination
	SetHistoricalObjectReplication(v string) *ReplicationProgressRule
	GetHistoricalObjectReplication() *string
	SetID(v string) *ReplicationProgressRule
	GetID() *string
	SetPrefixSet(v *ReplicationPrefixSet) *ReplicationProgressRule
	GetPrefixSet() *ReplicationPrefixSet
	SetProgress(v *ReplicationProgressRuleProgress) *ReplicationProgressRule
	GetProgress() *ReplicationProgressRuleProgress
	SetStatus(v string) *ReplicationProgressRule
	GetStatus() *string
}

type ReplicationProgressRule struct {
	// example:
	//
	// ALL
	Action      *string                 `json:"Action,omitempty" xml:"Action,omitempty"`
	Destination *ReplicationDestination `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// example:
	//
	// disabled
	HistoricalObjectReplication *string `json:"HistoricalObjectReplication,omitempty" xml:"HistoricalObjectReplication,omitempty"`
	// example:
	//
	// replicate001
	ID        *string                          `json:"ID,omitempty" xml:"ID,omitempty"`
	PrefixSet *ReplicationPrefixSet            `json:"PrefixSet,omitempty" xml:"PrefixSet,omitempty"`
	Progress  *ReplicationProgressRuleProgress `json:"Progress,omitempty" xml:"Progress,omitempty" type:"Struct"`
	// example:
	//
	// doing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ReplicationProgressRule) String() string {
	return dara.Prettify(s)
}

func (s ReplicationProgressRule) GoString() string {
	return s.String()
}

func (s *ReplicationProgressRule) GetAction() *string {
	return s.Action
}

func (s *ReplicationProgressRule) GetDestination() *ReplicationDestination {
	return s.Destination
}

func (s *ReplicationProgressRule) GetHistoricalObjectReplication() *string {
	return s.HistoricalObjectReplication
}

func (s *ReplicationProgressRule) GetID() *string {
	return s.ID
}

func (s *ReplicationProgressRule) GetPrefixSet() *ReplicationPrefixSet {
	return s.PrefixSet
}

func (s *ReplicationProgressRule) GetProgress() *ReplicationProgressRuleProgress {
	return s.Progress
}

func (s *ReplicationProgressRule) GetStatus() *string {
	return s.Status
}

func (s *ReplicationProgressRule) SetAction(v string) *ReplicationProgressRule {
	s.Action = &v
	return s
}

func (s *ReplicationProgressRule) SetDestination(v *ReplicationDestination) *ReplicationProgressRule {
	s.Destination = v
	return s
}

func (s *ReplicationProgressRule) SetHistoricalObjectReplication(v string) *ReplicationProgressRule {
	s.HistoricalObjectReplication = &v
	return s
}

func (s *ReplicationProgressRule) SetID(v string) *ReplicationProgressRule {
	s.ID = &v
	return s
}

func (s *ReplicationProgressRule) SetPrefixSet(v *ReplicationPrefixSet) *ReplicationProgressRule {
	s.PrefixSet = v
	return s
}

func (s *ReplicationProgressRule) SetProgress(v *ReplicationProgressRuleProgress) *ReplicationProgressRule {
	s.Progress = v
	return s
}

func (s *ReplicationProgressRule) SetStatus(v string) *ReplicationProgressRule {
	s.Status = &v
	return s
}

func (s *ReplicationProgressRule) Validate() error {
	if s.Destination != nil {
		if err := s.Destination.Validate(); err != nil {
			return err
		}
	}
	if s.PrefixSet != nil {
		if err := s.PrefixSet.Validate(); err != nil {
			return err
		}
	}
	if s.Progress != nil {
		if err := s.Progress.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReplicationProgressRuleProgress struct {
	// example:
	//
	// 0.85
	HistoricalObject *string `json:"HistoricalObject,omitempty" xml:"HistoricalObject,omitempty"`
	// example:
	//
	// Thu, 24 Sep 2015 15:39:18 GMT
	NewObject *string `json:"NewObject,omitempty" xml:"NewObject,omitempty"`
}

func (s ReplicationProgressRuleProgress) String() string {
	return dara.Prettify(s)
}

func (s ReplicationProgressRuleProgress) GoString() string {
	return s.String()
}

func (s *ReplicationProgressRuleProgress) GetHistoricalObject() *string {
	return s.HistoricalObject
}

func (s *ReplicationProgressRuleProgress) GetNewObject() *string {
	return s.NewObject
}

func (s *ReplicationProgressRuleProgress) SetHistoricalObject(v string) *ReplicationProgressRuleProgress {
	s.HistoricalObject = &v
	return s
}

func (s *ReplicationProgressRuleProgress) SetNewObject(v string) *ReplicationProgressRuleProgress {
	s.NewObject = &v
	return s
}

func (s *ReplicationProgressRuleProgress) Validate() error {
	return dara.Validate(s)
}
