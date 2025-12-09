// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplicationDestination interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *ReplicationDestination
	GetBucket() *string
	SetLocation(v string) *ReplicationDestination
	GetLocation() *string
	SetTransferType(v string) *ReplicationDestination
	GetTransferType() *string
}

type ReplicationDestination struct {
	// example:
	//
	// destbucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// example:
	//
	// oss-cn-hangzhou
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// internal
	TransferType *string `json:"TransferType,omitempty" xml:"TransferType,omitempty"`
}

func (s ReplicationDestination) String() string {
	return dara.Prettify(s)
}

func (s ReplicationDestination) GoString() string {
	return s.String()
}

func (s *ReplicationDestination) GetBucket() *string {
	return s.Bucket
}

func (s *ReplicationDestination) GetLocation() *string {
	return s.Location
}

func (s *ReplicationDestination) GetTransferType() *string {
	return s.TransferType
}

func (s *ReplicationDestination) SetBucket(v string) *ReplicationDestination {
	s.Bucket = &v
	return s
}

func (s *ReplicationDestination) SetLocation(v string) *ReplicationDestination {
	s.Location = &v
	return s
}

func (s *ReplicationDestination) SetTransferType(v string) *ReplicationDestination {
	s.TransferType = &v
	return s
}

func (s *ReplicationDestination) Validate() error {
	return dara.Validate(s)
}
