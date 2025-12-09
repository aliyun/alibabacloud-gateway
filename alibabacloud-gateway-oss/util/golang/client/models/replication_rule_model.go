// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplicationRule interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *ReplicationRule
	GetAction() *string
	SetDestination(v *ReplicationDestination) *ReplicationRule
	GetDestination() *ReplicationDestination
	SetEncryptionConfiguration(v *ReplicationEncryptionConfiguration) *ReplicationRule
	GetEncryptionConfiguration() *ReplicationEncryptionConfiguration
	SetHistoricalObjectReplication(v string) *ReplicationRule
	GetHistoricalObjectReplication() *string
	SetID(v string) *ReplicationRule
	GetID() *string
	SetPrefixSet(v *ReplicationPrefixSet) *ReplicationRule
	GetPrefixSet() *ReplicationPrefixSet
	SetRTC(v *RTC) *ReplicationRule
	GetRTC() *RTC
	SetSourceSelectionCriteria(v *ReplicationSourceSelectionCriteria) *ReplicationRule
	GetSourceSelectionCriteria() *ReplicationSourceSelectionCriteria
	SetStatus(v string) *ReplicationRule
	GetStatus() *string
	SetSyncRole(v string) *ReplicationRule
	GetSyncRole() *string
}

type ReplicationRule struct {
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
	// doing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// aliyunramrole
	SyncRole *string `json:"SyncRole,omitempty" xml:"SyncRole,omitempty"`
}

func (s ReplicationRule) String() string {
	return dara.Prettify(s)
}

func (s ReplicationRule) GoString() string {
	return s.String()
}

func (s *ReplicationRule) GetAction() *string {
	return s.Action
}

func (s *ReplicationRule) GetDestination() *ReplicationDestination {
	return s.Destination
}

func (s *ReplicationRule) GetEncryptionConfiguration() *ReplicationEncryptionConfiguration {
	return s.EncryptionConfiguration
}

func (s *ReplicationRule) GetHistoricalObjectReplication() *string {
	return s.HistoricalObjectReplication
}

func (s *ReplicationRule) GetID() *string {
	return s.ID
}

func (s *ReplicationRule) GetPrefixSet() *ReplicationPrefixSet {
	return s.PrefixSet
}

func (s *ReplicationRule) GetRTC() *RTC {
	return s.RTC
}

func (s *ReplicationRule) GetSourceSelectionCriteria() *ReplicationSourceSelectionCriteria {
	return s.SourceSelectionCriteria
}

func (s *ReplicationRule) GetStatus() *string {
	return s.Status
}

func (s *ReplicationRule) GetSyncRole() *string {
	return s.SyncRole
}

func (s *ReplicationRule) SetAction(v string) *ReplicationRule {
	s.Action = &v
	return s
}

func (s *ReplicationRule) SetDestination(v *ReplicationDestination) *ReplicationRule {
	s.Destination = v
	return s
}

func (s *ReplicationRule) SetEncryptionConfiguration(v *ReplicationEncryptionConfiguration) *ReplicationRule {
	s.EncryptionConfiguration = v
	return s
}

func (s *ReplicationRule) SetHistoricalObjectReplication(v string) *ReplicationRule {
	s.HistoricalObjectReplication = &v
	return s
}

func (s *ReplicationRule) SetID(v string) *ReplicationRule {
	s.ID = &v
	return s
}

func (s *ReplicationRule) SetPrefixSet(v *ReplicationPrefixSet) *ReplicationRule {
	s.PrefixSet = v
	return s
}

func (s *ReplicationRule) SetRTC(v *RTC) *ReplicationRule {
	s.RTC = v
	return s
}

func (s *ReplicationRule) SetSourceSelectionCriteria(v *ReplicationSourceSelectionCriteria) *ReplicationRule {
	s.SourceSelectionCriteria = v
	return s
}

func (s *ReplicationRule) SetStatus(v string) *ReplicationRule {
	s.Status = &v
	return s
}

func (s *ReplicationRule) SetSyncRole(v string) *ReplicationRule {
	s.SyncRole = &v
	return s
}

func (s *ReplicationRule) Validate() error {
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
