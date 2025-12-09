// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplicationRules interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v []*string) *ReplicationRules
	GetIds() []*string
}

type ReplicationRules struct {
	Ids []*string `json:"ID,omitempty" xml:"ID,omitempty" type:"Repeated"`
}

func (s ReplicationRules) String() string {
	return dara.Prettify(s)
}

func (s ReplicationRules) GoString() string {
	return s.String()
}

func (s *ReplicationRules) GetIds() []*string {
	return s.Ids
}

func (s *ReplicationRules) SetIds(v []*string) *ReplicationRules {
	s.Ids = v
	return s
}

func (s *ReplicationRules) Validate() error {
	return dara.Validate(s)
}
