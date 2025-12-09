// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplicationRuleBandwidth interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int64) *ReplicationRuleBandwidth
	GetBandwidth() *int64
	SetID(v string) *ReplicationRuleBandwidth
	GetID() *string
}

type ReplicationRuleBandwidth struct {
	// example:
	//
	// 209715200
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// example:
	//
	// test_replication_1
	ID *string `json:"ID,omitempty" xml:"ID,omitempty"`
}

func (s ReplicationRuleBandwidth) String() string {
	return dara.Prettify(s)
}

func (s ReplicationRuleBandwidth) GoString() string {
	return s.String()
}

func (s *ReplicationRuleBandwidth) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *ReplicationRuleBandwidth) GetID() *string {
	return s.ID
}

func (s *ReplicationRuleBandwidth) SetBandwidth(v int64) *ReplicationRuleBandwidth {
	s.Bandwidth = &v
	return s
}

func (s *ReplicationRuleBandwidth) SetID(v string) *ReplicationRuleBandwidth {
	s.ID = &v
	return s
}

func (s *ReplicationRuleBandwidth) Validate() error {
	return dara.Validate(s)
}
