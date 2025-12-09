// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketCallbackPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBucketCallbackPolicy(v *CallbackPolicy) *GetBucketCallbackPolicyResponseBody
	GetBucketCallbackPolicy() *CallbackPolicy
}

type GetBucketCallbackPolicyResponseBody struct {
	BucketCallbackPolicy *CallbackPolicy `json:"BucketCallbackPolicy,omitempty" xml:"BucketCallbackPolicy,omitempty"`
}

func (s GetBucketCallbackPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketCallbackPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketCallbackPolicyResponseBody) GetBucketCallbackPolicy() *CallbackPolicy {
	return s.BucketCallbackPolicy
}

func (s *GetBucketCallbackPolicyResponseBody) SetBucketCallbackPolicy(v *CallbackPolicy) *GetBucketCallbackPolicyResponseBody {
	s.BucketCallbackPolicy = v
	return s
}

func (s *GetBucketCallbackPolicyResponseBody) Validate() error {
	if s.BucketCallbackPolicy != nil {
		if err := s.BucketCallbackPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}
