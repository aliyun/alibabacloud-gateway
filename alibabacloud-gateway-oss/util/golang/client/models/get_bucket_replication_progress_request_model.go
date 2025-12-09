// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketReplicationProgressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRuleId(v string) *GetBucketReplicationProgressRequest
	GetRuleId() *string
}

type GetBucketReplicationProgressRequest struct {
	// The ID of the data replication rule. You can call the GetBucketReplication operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_replication_1
	RuleId *string `json:"rule-id,omitempty" xml:"rule-id,omitempty"`
}

func (s GetBucketReplicationProgressRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBucketReplicationProgressRequest) GoString() string {
	return s.String()
}

func (s *GetBucketReplicationProgressRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *GetBucketReplicationProgressRequest) SetRuleId(v string) *GetBucketReplicationProgressRequest {
	s.RuleId = &v
	return s
}

func (s *GetBucketReplicationProgressRequest) Validate() error {
	return dara.Validate(s)
}
