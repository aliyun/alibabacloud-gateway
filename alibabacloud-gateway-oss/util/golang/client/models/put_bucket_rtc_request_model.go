// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketRtcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReplicationRule(v *RtcConfiguration) *PutBucketRtcRequest
	GetReplicationRule() *RtcConfiguration
}

type PutBucketRtcRequest struct {
	// The container that stores the RTC configurations.
	ReplicationRule *RtcConfiguration `json:"ReplicationRule,omitempty" xml:"ReplicationRule,omitempty"`
}

func (s PutBucketRtcRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketRtcRequest) GoString() string {
	return s.String()
}

func (s *PutBucketRtcRequest) GetReplicationRule() *RtcConfiguration {
	return s.ReplicationRule
}

func (s *PutBucketRtcRequest) SetReplicationRule(v *RtcConfiguration) *PutBucketRtcRequest {
	s.ReplicationRule = v
	return s
}

func (s *PutBucketRtcRequest) Validate() error {
	if s.ReplicationRule != nil {
		if err := s.ReplicationRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}
