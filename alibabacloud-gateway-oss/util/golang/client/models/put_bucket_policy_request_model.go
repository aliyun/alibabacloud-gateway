// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicy(v string) *PutBucketPolicyRequest
	GetPolicy() *string
}

type PutBucketPolicyRequest struct {
	// The request parameters.
	//
	// This parameter is required.
	Policy *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutBucketPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketPolicyRequest) GoString() string {
	return s.String()
}

func (s *PutBucketPolicyRequest) GetPolicy() *string {
	return s.Policy
}

func (s *PutBucketPolicyRequest) SetPolicy(v string) *PutBucketPolicyRequest {
	s.Policy = &v
	return s
}

func (s *PutBucketPolicyRequest) Validate() error {
	return dara.Validate(s)
}
