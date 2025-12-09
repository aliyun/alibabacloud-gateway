// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplicationConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetRule(v *PutReplicationRule) *ReplicationConfiguration
	GetRule() *PutReplicationRule
}

type ReplicationConfiguration struct {
	Rule *PutReplicationRule `json:"Rule,omitempty" xml:"Rule,omitempty"`
}

func (s ReplicationConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ReplicationConfiguration) GoString() string {
	return s.String()
}

func (s *ReplicationConfiguration) GetRule() *PutReplicationRule {
	return s.Rule
}

func (s *ReplicationConfiguration) SetRule(v *PutReplicationRule) *ReplicationConfiguration {
	s.Rule = v
	return s
}

func (s *ReplicationConfiguration) Validate() error {
	if s.Rule != nil {
		if err := s.Rule.Validate(); err != nil {
			return err
		}
	}
	return nil
}
