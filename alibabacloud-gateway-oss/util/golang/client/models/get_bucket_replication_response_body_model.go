// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketReplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetReplicationConfiguration(v *GetBucketReplicationResponseBodyReplicationConfiguration) *GetBucketReplicationResponseBody
	GetReplicationConfiguration() *GetBucketReplicationResponseBodyReplicationConfiguration
}

type GetBucketReplicationResponseBody struct {
	// The container that stores data replication configurations.
	ReplicationConfiguration *GetBucketReplicationResponseBodyReplicationConfiguration `json:"ReplicationConfiguration,omitempty" xml:"ReplicationConfiguration,omitempty" type:"Struct"`
}

func (s GetBucketReplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketReplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketReplicationResponseBody) GetReplicationConfiguration() *GetBucketReplicationResponseBodyReplicationConfiguration {
	return s.ReplicationConfiguration
}

func (s *GetBucketReplicationResponseBody) SetReplicationConfiguration(v *GetBucketReplicationResponseBodyReplicationConfiguration) *GetBucketReplicationResponseBody {
	s.ReplicationConfiguration = v
	return s
}

func (s *GetBucketReplicationResponseBody) Validate() error {
	if s.ReplicationConfiguration != nil {
		if err := s.ReplicationConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBucketReplicationResponseBodyReplicationConfiguration struct {
	// The container that stores the data replication rules.
	Rule []*ReplicationRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s GetBucketReplicationResponseBodyReplicationConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetBucketReplicationResponseBodyReplicationConfiguration) GoString() string {
	return s.String()
}

func (s *GetBucketReplicationResponseBodyReplicationConfiguration) GetRule() []*ReplicationRule {
	return s.Rule
}

func (s *GetBucketReplicationResponseBodyReplicationConfiguration) SetRule(v []*ReplicationRule) *GetBucketReplicationResponseBodyReplicationConfiguration {
	s.Rule = v
	return s
}

func (s *GetBucketReplicationResponseBodyReplicationConfiguration) Validate() error {
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
