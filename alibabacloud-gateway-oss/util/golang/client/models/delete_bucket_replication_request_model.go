// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketReplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReplicationRules(v *DeleteBucketReplicationRequestReplicationRules) *DeleteBucketReplicationRequest
	GetReplicationRules() *DeleteBucketReplicationRequestReplicationRules
}

type DeleteBucketReplicationRequest struct {
	// The container that is used to store the data replication rule to delete.
	ReplicationRules *DeleteBucketReplicationRequestReplicationRules `json:"ReplicationRules,omitempty" xml:"ReplicationRules,omitempty" type:"Struct"`
}

func (s DeleteBucketReplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketReplicationRequest) GoString() string {
	return s.String()
}

func (s *DeleteBucketReplicationRequest) GetReplicationRules() *DeleteBucketReplicationRequestReplicationRules {
	return s.ReplicationRules
}

func (s *DeleteBucketReplicationRequest) SetReplicationRules(v *DeleteBucketReplicationRequestReplicationRules) *DeleteBucketReplicationRequest {
	s.ReplicationRules = v
	return s
}

func (s *DeleteBucketReplicationRequest) Validate() error {
	if s.ReplicationRules != nil {
		if err := s.ReplicationRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteBucketReplicationRequestReplicationRules struct {
	ID *string `json:"ID,omitempty" xml:"ID,omitempty"`
}

func (s DeleteBucketReplicationRequestReplicationRules) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketReplicationRequestReplicationRules) GoString() string {
	return s.String()
}

func (s *DeleteBucketReplicationRequestReplicationRules) GetID() *string {
	return s.ID
}

func (s *DeleteBucketReplicationRequestReplicationRules) SetID(v string) *DeleteBucketReplicationRequestReplicationRules {
	s.ID = &v
	return s
}

func (s *DeleteBucketReplicationRequestReplicationRules) Validate() error {
	return dara.Validate(s)
}
