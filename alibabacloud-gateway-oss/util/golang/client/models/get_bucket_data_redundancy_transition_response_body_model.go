// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketDataRedundancyTransitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBucketDataRedundancyTransition(v *BucketDataRedundancyTransition) *GetBucketDataRedundancyTransitionResponseBody
	GetBucketDataRedundancyTransition() *BucketDataRedundancyTransition
}

type GetBucketDataRedundancyTransitionResponseBody struct {
	// The container for a specific redundancy type change task.
	BucketDataRedundancyTransition *BucketDataRedundancyTransition `json:"BucketDataRedundancyTransition,omitempty" xml:"BucketDataRedundancyTransition,omitempty"`
}

func (s GetBucketDataRedundancyTransitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketDataRedundancyTransitionResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketDataRedundancyTransitionResponseBody) GetBucketDataRedundancyTransition() *BucketDataRedundancyTransition {
	return s.BucketDataRedundancyTransition
}

func (s *GetBucketDataRedundancyTransitionResponseBody) SetBucketDataRedundancyTransition(v *BucketDataRedundancyTransition) *GetBucketDataRedundancyTransitionResponseBody {
	s.BucketDataRedundancyTransition = v
	return s
}

func (s *GetBucketDataRedundancyTransitionResponseBody) Validate() error {
	if s.BucketDataRedundancyTransition != nil {
		if err := s.BucketDataRedundancyTransition.Validate(); err != nil {
			return err
		}
	}
	return nil
}
