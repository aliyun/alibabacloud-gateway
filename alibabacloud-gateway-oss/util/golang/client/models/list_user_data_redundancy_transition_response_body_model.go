// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserDataRedundancyTransitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListBucketDataRedundancyTransition(v *ListUserDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition) *ListUserDataRedundancyTransitionResponseBody
	GetListBucketDataRedundancyTransition() *ListUserDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition
}

type ListUserDataRedundancyTransitionResponseBody struct {
	// Container for listing storage redundancy transition tasks.
	ListBucketDataRedundancyTransition *ListUserDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition `json:"ListBucketDataRedundancyTransition,omitempty" xml:"ListBucketDataRedundancyTransition,omitempty" type:"Struct"`
}

func (s ListUserDataRedundancyTransitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserDataRedundancyTransitionResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserDataRedundancyTransitionResponseBody) GetListBucketDataRedundancyTransition() *ListUserDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition {
	return s.ListBucketDataRedundancyTransition
}

func (s *ListUserDataRedundancyTransitionResponseBody) SetListBucketDataRedundancyTransition(v *ListUserDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition) *ListUserDataRedundancyTransitionResponseBody {
	s.ListBucketDataRedundancyTransition = v
	return s
}

func (s *ListUserDataRedundancyTransitionResponseBody) Validate() error {
	if s.ListBucketDataRedundancyTransition != nil {
		if err := s.ListBucketDataRedundancyTransition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUserDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition struct {
	// Container for storage redundancy transition tasks.
	BucketDataRedundancyTransition []*BucketDataRedundancyTransition `json:"BucketDataRedundancyTransition,omitempty" xml:"BucketDataRedundancyTransition,omitempty" type:"Repeated"`
	// Indicates whether the results in the request are truncated. The values are as follows:
	//
	// true: Indicates that not all results are returned in this request.
	//
	// false: Indicates that all results are returned in this request.
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// Indicates that the current ListUserDataRedundancyTransition request contains subsequent results, and you need to specify NextContinuationToken as continuation-token to continue retrieving the results.
	//
	// example:
	//
	// def
	NextContinuationToken *string `json:"NextContinuationToken,omitempty" xml:"NextContinuationToken,omitempty"`
}

func (s ListUserDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition) String() string {
	return dara.Prettify(s)
}

func (s ListUserDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition) GoString() string {
	return s.String()
}

func (s *ListUserDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition) GetBucketDataRedundancyTransition() []*BucketDataRedundancyTransition {
	return s.BucketDataRedundancyTransition
}

func (s *ListUserDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListUserDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition) GetNextContinuationToken() *string {
	return s.NextContinuationToken
}

func (s *ListUserDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition) SetBucketDataRedundancyTransition(v []*BucketDataRedundancyTransition) *ListUserDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition {
	s.BucketDataRedundancyTransition = v
	return s
}

func (s *ListUserDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition) SetIsTruncated(v bool) *ListUserDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition {
	s.IsTruncated = &v
	return s
}

func (s *ListUserDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition) SetNextContinuationToken(v string) *ListUserDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition {
	s.NextContinuationToken = &v
	return s
}

func (s *ListUserDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition) Validate() error {
	if s.BucketDataRedundancyTransition != nil {
		for _, item := range s.BucketDataRedundancyTransition {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
