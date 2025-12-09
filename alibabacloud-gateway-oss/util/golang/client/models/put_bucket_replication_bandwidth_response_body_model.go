// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketReplicationBandwidthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetReplicationRule(v *PutBucketReplicationBandwidthResponseBodyReplicationRule) *PutBucketReplicationBandwidthResponseBody
	GetReplicationRule() *PutBucketReplicationBandwidthResponseBodyReplicationRule
}

type PutBucketReplicationBandwidthResponseBody struct {
	ReplicationRule *PutBucketReplicationBandwidthResponseBodyReplicationRule `json:"ReplicationRule,omitempty" xml:"ReplicationRule,omitempty" type:"Struct"`
}

func (s PutBucketReplicationBandwidthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutBucketReplicationBandwidthResponseBody) GoString() string {
	return s.String()
}

func (s *PutBucketReplicationBandwidthResponseBody) GetReplicationRule() *PutBucketReplicationBandwidthResponseBodyReplicationRule {
	return s.ReplicationRule
}

func (s *PutBucketReplicationBandwidthResponseBody) SetReplicationRule(v *PutBucketReplicationBandwidthResponseBodyReplicationRule) *PutBucketReplicationBandwidthResponseBody {
	s.ReplicationRule = v
	return s
}

func (s *PutBucketReplicationBandwidthResponseBody) Validate() error {
	if s.ReplicationRule != nil {
		if err := s.ReplicationRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PutBucketReplicationBandwidthResponseBodyReplicationRule struct {
	Bandwith *int64  `json:"Bandwith,omitempty" xml:"Bandwith,omitempty"`
	ID       *string `json:"ID,omitempty" xml:"ID,omitempty"`
}

func (s PutBucketReplicationBandwidthResponseBodyReplicationRule) String() string {
	return dara.Prettify(s)
}

func (s PutBucketReplicationBandwidthResponseBodyReplicationRule) GoString() string {
	return s.String()
}

func (s *PutBucketReplicationBandwidthResponseBodyReplicationRule) GetBandwith() *int64 {
	return s.Bandwith
}

func (s *PutBucketReplicationBandwidthResponseBodyReplicationRule) GetID() *string {
	return s.ID
}

func (s *PutBucketReplicationBandwidthResponseBodyReplicationRule) SetBandwith(v int64) *PutBucketReplicationBandwidthResponseBodyReplicationRule {
	s.Bandwith = &v
	return s
}

func (s *PutBucketReplicationBandwidthResponseBodyReplicationRule) SetID(v string) *PutBucketReplicationBandwidthResponseBodyReplicationRule {
	s.ID = &v
	return s
}

func (s *PutBucketReplicationBandwidthResponseBodyReplicationRule) Validate() error {
	return dara.Validate(s)
}
