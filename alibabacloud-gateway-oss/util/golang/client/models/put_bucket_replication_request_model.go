// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketReplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReplicationConfiguration(v *ReplicationConfiguration) *PutBucketReplicationRequest
	GetReplicationConfiguration() *ReplicationConfiguration
}

type PutBucketReplicationRequest struct {
	// The container that stores data replication configurations.
	ReplicationConfiguration *ReplicationConfiguration `json:"ReplicationConfiguration,omitempty" xml:"ReplicationConfiguration,omitempty"`
}

func (s PutBucketReplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketReplicationRequest) GoString() string {
	return s.String()
}

func (s *PutBucketReplicationRequest) GetReplicationConfiguration() *ReplicationConfiguration {
	return s.ReplicationConfiguration
}

func (s *PutBucketReplicationRequest) SetReplicationConfiguration(v *ReplicationConfiguration) *PutBucketReplicationRequest {
	s.ReplicationConfiguration = v
	return s
}

func (s *PutBucketReplicationRequest) Validate() error {
	if s.ReplicationConfiguration != nil {
		if err := s.ReplicationConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
