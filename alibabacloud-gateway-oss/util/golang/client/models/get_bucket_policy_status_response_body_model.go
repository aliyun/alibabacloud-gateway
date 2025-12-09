// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketPolicyStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyStatus(v *GetBucketPolicyStatusResponseBodyPolicyStatus) *GetBucketPolicyStatusResponseBody
	GetPolicyStatus() *GetBucketPolicyStatusResponseBodyPolicyStatus
}

type GetBucketPolicyStatusResponseBody struct {
	// The container that stores public access information.
	PolicyStatus *GetBucketPolicyStatusResponseBodyPolicyStatus `json:"PolicyStatus,omitempty" xml:"PolicyStatus,omitempty" type:"Struct"`
}

func (s GetBucketPolicyStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketPolicyStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketPolicyStatusResponseBody) GetPolicyStatus() *GetBucketPolicyStatusResponseBodyPolicyStatus {
	return s.PolicyStatus
}

func (s *GetBucketPolicyStatusResponseBody) SetPolicyStatus(v *GetBucketPolicyStatusResponseBodyPolicyStatus) *GetBucketPolicyStatusResponseBody {
	s.PolicyStatus = v
	return s
}

func (s *GetBucketPolicyStatusResponseBody) Validate() error {
	if s.PolicyStatus != nil {
		if err := s.PolicyStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBucketPolicyStatusResponseBodyPolicyStatus struct {
	// Indicates whether the current bucket policy allows public access.
	//
	// true
	//
	// false
	//
	// example:
	//
	// true
	IsPublic *bool `json:"IsPublic,omitempty" xml:"IsPublic,omitempty"`
}

func (s GetBucketPolicyStatusResponseBodyPolicyStatus) String() string {
	return dara.Prettify(s)
}

func (s GetBucketPolicyStatusResponseBodyPolicyStatus) GoString() string {
	return s.String()
}

func (s *GetBucketPolicyStatusResponseBodyPolicyStatus) GetIsPublic() *bool {
	return s.IsPublic
}

func (s *GetBucketPolicyStatusResponseBodyPolicyStatus) SetIsPublic(v bool) *GetBucketPolicyStatusResponseBodyPolicyStatus {
	s.IsPublic = &v
	return s
}

func (s *GetBucketPolicyStatusResponseBodyPolicyStatus) Validate() error {
	return dara.Validate(s)
}
