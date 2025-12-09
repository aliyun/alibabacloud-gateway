// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketReplicationProgressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetReplicationProgress(v *GetBucketReplicationProgressResponseBodyReplicationProgress) *GetBucketReplicationProgressResponseBody
	GetReplicationProgress() *GetBucketReplicationProgressResponseBodyReplicationProgress
}

type GetBucketReplicationProgressResponseBody struct {
	// The container that is used to store the progress of data replication tasks.
	ReplicationProgress *GetBucketReplicationProgressResponseBodyReplicationProgress `json:"ReplicationProgress,omitempty" xml:"ReplicationProgress,omitempty" type:"Struct"`
}

func (s GetBucketReplicationProgressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketReplicationProgressResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketReplicationProgressResponseBody) GetReplicationProgress() *GetBucketReplicationProgressResponseBodyReplicationProgress {
	return s.ReplicationProgress
}

func (s *GetBucketReplicationProgressResponseBody) SetReplicationProgress(v *GetBucketReplicationProgressResponseBodyReplicationProgress) *GetBucketReplicationProgressResponseBody {
	s.ReplicationProgress = v
	return s
}

func (s *GetBucketReplicationProgressResponseBody) Validate() error {
	if s.ReplicationProgress != nil {
		if err := s.ReplicationProgress.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBucketReplicationProgressResponseBodyReplicationProgress struct {
	// The container that stores the progress of the data replication task corresponding to each data replication rule.
	Rule []*ReplicationProgressRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s GetBucketReplicationProgressResponseBodyReplicationProgress) String() string {
	return dara.Prettify(s)
}

func (s GetBucketReplicationProgressResponseBodyReplicationProgress) GoString() string {
	return s.String()
}

func (s *GetBucketReplicationProgressResponseBodyReplicationProgress) GetRule() []*ReplicationProgressRule {
	return s.Rule
}

func (s *GetBucketReplicationProgressResponseBodyReplicationProgress) SetRule(v []*ReplicationProgressRule) *GetBucketReplicationProgressResponseBodyReplicationProgress {
	s.Rule = v
	return s
}

func (s *GetBucketReplicationProgressResponseBodyReplicationProgress) Validate() error {
	if s.Rule != nil {
		for _, item := range s.Rule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
