// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketCallbackPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketCallbackPolicy(v *CallbackPolicy) *PutBucketCallbackPolicyRequest
	GetBucketCallbackPolicy() *CallbackPolicy
}

type PutBucketCallbackPolicyRequest struct {
	BucketCallbackPolicy *CallbackPolicy `json:"BucketCallbackPolicy,omitempty" xml:"BucketCallbackPolicy,omitempty"`
}

func (s PutBucketCallbackPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketCallbackPolicyRequest) GoString() string {
	return s.String()
}

func (s *PutBucketCallbackPolicyRequest) GetBucketCallbackPolicy() *CallbackPolicy {
	return s.BucketCallbackPolicy
}

func (s *PutBucketCallbackPolicyRequest) SetBucketCallbackPolicy(v *CallbackPolicy) *PutBucketCallbackPolicyRequest {
	s.BucketCallbackPolicy = v
	return s
}

func (s *PutBucketCallbackPolicyRequest) Validate() error {
	if s.BucketCallbackPolicy != nil {
		if err := s.BucketCallbackPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}
