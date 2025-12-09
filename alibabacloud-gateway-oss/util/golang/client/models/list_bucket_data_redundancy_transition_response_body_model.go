// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBucketDataRedundancyTransitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListBucketDataRedundancyTransition(v *ListBucketDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition) *ListBucketDataRedundancyTransitionResponseBody
	GetListBucketDataRedundancyTransition() *ListBucketDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition
}

type ListBucketDataRedundancyTransitionResponseBody struct {
	// The container for listed redundancy type conversion tasks.
	ListBucketDataRedundancyTransition *ListBucketDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition `json:"ListBucketDataRedundancyTransition,omitempty" xml:"ListBucketDataRedundancyTransition,omitempty" type:"Struct"`
}

func (s ListBucketDataRedundancyTransitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBucketDataRedundancyTransitionResponseBody) GoString() string {
	return s.String()
}

func (s *ListBucketDataRedundancyTransitionResponseBody) GetListBucketDataRedundancyTransition() *ListBucketDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition {
	return s.ListBucketDataRedundancyTransition
}

func (s *ListBucketDataRedundancyTransitionResponseBody) SetListBucketDataRedundancyTransition(v *ListBucketDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition) *ListBucketDataRedundancyTransitionResponseBody {
	s.ListBucketDataRedundancyTransition = v
	return s
}

func (s *ListBucketDataRedundancyTransitionResponseBody) Validate() error {
	if s.ListBucketDataRedundancyTransition != nil {
		if err := s.ListBucketDataRedundancyTransition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListBucketDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition struct {
	// The information about the redundancy type conversion task.
	BucketDataRedundancyTransition *BucketDataRedundancyTransition `json:"BucketDataRedundancyTransition,omitempty" xml:"BucketDataRedundancyTransition,omitempty"`
}

func (s ListBucketDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition) String() string {
	return dara.Prettify(s)
}

func (s ListBucketDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition) GoString() string {
	return s.String()
}

func (s *ListBucketDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition) GetBucketDataRedundancyTransition() *BucketDataRedundancyTransition {
	return s.BucketDataRedundancyTransition
}

func (s *ListBucketDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition) SetBucketDataRedundancyTransition(v *BucketDataRedundancyTransition) *ListBucketDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition {
	s.BucketDataRedundancyTransition = v
	return s
}

func (s *ListBucketDataRedundancyTransitionResponseBodyListBucketDataRedundancyTransition) Validate() error {
	if s.BucketDataRedundancyTransition != nil {
		if err := s.BucketDataRedundancyTransition.Validate(); err != nil {
			return err
		}
	}
	return nil
}
