// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplicationRuleProgress interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *ReplicationRuleProgress
	GetAction() *string
	SetID(v string) *ReplicationRuleProgress
	GetID() *string
	SetPrefixSet(v *ReplicationPrefixSet) *ReplicationRuleProgress
	GetPrefixSet() *ReplicationPrefixSet
}

type ReplicationRuleProgress struct {
	Action    *string               `json:"Action,omitempty" xml:"Action,omitempty"`
	ID        *string               `json:"ID,omitempty" xml:"ID,omitempty"`
	PrefixSet *ReplicationPrefixSet `json:"PrefixSet,omitempty" xml:"PrefixSet,omitempty"`
}

func (s ReplicationRuleProgress) String() string {
	return dara.Prettify(s)
}

func (s ReplicationRuleProgress) GoString() string {
	return s.String()
}

func (s *ReplicationRuleProgress) GetAction() *string {
	return s.Action
}

func (s *ReplicationRuleProgress) GetID() *string {
	return s.ID
}

func (s *ReplicationRuleProgress) GetPrefixSet() *ReplicationPrefixSet {
	return s.PrefixSet
}

func (s *ReplicationRuleProgress) SetAction(v string) *ReplicationRuleProgress {
	s.Action = &v
	return s
}

func (s *ReplicationRuleProgress) SetID(v string) *ReplicationRuleProgress {
	s.ID = &v
	return s
}

func (s *ReplicationRuleProgress) SetPrefixSet(v *ReplicationPrefixSet) *ReplicationRuleProgress {
	s.PrefixSet = v
	return s
}

func (s *ReplicationRuleProgress) Validate() error {
	if s.PrefixSet != nil {
		if err := s.PrefixSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}
