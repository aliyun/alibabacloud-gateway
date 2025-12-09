// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutReplicationRule interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *PutReplicationRule
	GetAction() *string
	SetDestination(v *ReplicationDestination) *PutReplicationRule
	GetDestination() *ReplicationDestination
	SetEncryptionConfiguration(v *ReplicationEncryptionConfiguration) *PutReplicationRule
	GetEncryptionConfiguration() *ReplicationEncryptionConfiguration
	SetHistoricalObjectReplication(v string) *PutReplicationRule
	GetHistoricalObjectReplication() *string
	SetID(v string) *PutReplicationRule
	GetID() *string
	SetPrefixSet(v *ReplicationPrefixSet) *PutReplicationRule
	GetPrefixSet() *ReplicationPrefixSet
	SetRTC(v *RTC) *PutReplicationRule
	GetRTC() *RTC
	SetSourceSelectionCriteria(v *ReplicationSourceSelectionCriteria) *PutReplicationRule
	GetSourceSelectionCriteria() *ReplicationSourceSelectionCriteria
	SetSyncRole(v string) *PutReplicationRule
	GetSyncRole() *string
}

type PutReplicationRule struct {
	// example:
	//
	// ALL
	Action                  *string                             `json:"Action,omitempty" xml:"Action,omitempty"`
	Destination             *ReplicationDestination             `json:"Destination,omitempty" xml:"Destination,omitempty"`
	EncryptionConfiguration *ReplicationEncryptionConfiguration `json:"EncryptionConfiguration,omitempty" xml:"EncryptionConfiguration,omitempty"`
	// example:
	//
	// disabled
	HistoricalObjectReplication *string `json:"HistoricalObjectReplication,omitempty" xml:"HistoricalObjectReplication,omitempty"`
	// example:
	//
	// first-repl-rule
	ID                      *string                             `json:"ID,omitempty" xml:"ID,omitempty"`
	PrefixSet               *ReplicationPrefixSet               `json:"PrefixSet,omitempty" xml:"PrefixSet,omitempty"`
	RTC                     *RTC                                `json:"RTC,omitempty" xml:"RTC,omitempty"`
	SourceSelectionCriteria *ReplicationSourceSelectionCriteria `json:"SourceSelectionCriteria,omitempty" xml:"SourceSelectionCriteria,omitempty"`
	// example:
	//
	// aliyunramrole
	SyncRole *string `json:"SyncRole,omitempty" xml:"SyncRole,omitempty"`
}

func (s PutReplicationRule) String() string {
	return dara.Prettify(s)
}

func (s PutReplicationRule) GoString() string {
	return s.String()
}

func (s *PutReplicationRule) GetAction() *string {
	return s.Action
}

func (s *PutReplicationRule) GetDestination() *ReplicationDestination {
	return s.Destination
}

func (s *PutReplicationRule) GetEncryptionConfiguration() *ReplicationEncryptionConfiguration {
	return s.EncryptionConfiguration
}

func (s *PutReplicationRule) GetHistoricalObjectReplication() *string {
	return s.HistoricalObjectReplication
}

func (s *PutReplicationRule) GetID() *string {
	return s.ID
}

func (s *PutReplicationRule) GetPrefixSet() *ReplicationPrefixSet {
	return s.PrefixSet
}

func (s *PutReplicationRule) GetRTC() *RTC {
	return s.RTC
}

func (s *PutReplicationRule) GetSourceSelectionCriteria() *ReplicationSourceSelectionCriteria {
	return s.SourceSelectionCriteria
}

func (s *PutReplicationRule) GetSyncRole() *string {
	return s.SyncRole
}

func (s *PutReplicationRule) SetAction(v string) *PutReplicationRule {
	s.Action = &v
	return s
}

func (s *PutReplicationRule) SetDestination(v *ReplicationDestination) *PutReplicationRule {
	s.Destination = v
	return s
}

func (s *PutReplicationRule) SetEncryptionConfiguration(v *ReplicationEncryptionConfiguration) *PutReplicationRule {
	s.EncryptionConfiguration = v
	return s
}

func (s *PutReplicationRule) SetHistoricalObjectReplication(v string) *PutReplicationRule {
	s.HistoricalObjectReplication = &v
	return s
}

func (s *PutReplicationRule) SetID(v string) *PutReplicationRule {
	s.ID = &v
	return s
}

func (s *PutReplicationRule) SetPrefixSet(v *ReplicationPrefixSet) *PutReplicationRule {
	s.PrefixSet = v
	return s
}

func (s *PutReplicationRule) SetRTC(v *RTC) *PutReplicationRule {
	s.RTC = v
	return s
}

func (s *PutReplicationRule) SetSourceSelectionCriteria(v *ReplicationSourceSelectionCriteria) *PutReplicationRule {
	s.SourceSelectionCriteria = v
	return s
}

func (s *PutReplicationRule) SetSyncRole(v string) *PutReplicationRule {
	s.SyncRole = &v
	return s
}

func (s *PutReplicationRule) Validate() error {
	if s.Destination != nil {
		if err := s.Destination.Validate(); err != nil {
			return err
		}
	}
	if s.EncryptionConfiguration != nil {
		if err := s.EncryptionConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.PrefixSet != nil {
		if err := s.PrefixSet.Validate(); err != nil {
			return err
		}
	}
	if s.RTC != nil {
		if err := s.RTC.Validate(); err != nil {
			return err
		}
	}
	if s.SourceSelectionCriteria != nil {
		if err := s.SourceSelectionCriteria.Validate(); err != nil {
			return err
		}
	}
	return nil
}
