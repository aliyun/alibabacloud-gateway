// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplicationEncryptionConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetReplicaKmsKeyID(v string) *ReplicationEncryptionConfiguration
	GetReplicaKmsKeyID() *string
}

type ReplicationEncryptionConfiguration struct {
	// example:
	//
	// c4d49f85-ee30-426b-a5ed-95e9139dxxxxx
	ReplicaKmsKeyID *string `json:"ReplicaKmsKeyID,omitempty" xml:"ReplicaKmsKeyID,omitempty"`
}

func (s ReplicationEncryptionConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ReplicationEncryptionConfiguration) GoString() string {
	return s.String()
}

func (s *ReplicationEncryptionConfiguration) GetReplicaKmsKeyID() *string {
	return s.ReplicaKmsKeyID
}

func (s *ReplicationEncryptionConfiguration) SetReplicaKmsKeyID(v string) *ReplicationEncryptionConfiguration {
	s.ReplicaKmsKeyID = &v
	return s
}

func (s *ReplicationEncryptionConfiguration) Validate() error {
	return dara.Validate(s)
}
