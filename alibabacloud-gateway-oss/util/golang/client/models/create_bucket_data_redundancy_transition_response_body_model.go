// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBucketDataRedundancyTransitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBucketDataRedundancyTransition(v *CreateBucketDataRedundancyTransitionResponseBodyBucketDataRedundancyTransition) *CreateBucketDataRedundancyTransitionResponseBody
	GetBucketDataRedundancyTransition() *CreateBucketDataRedundancyTransitionResponseBodyBucketDataRedundancyTransition
}

type CreateBucketDataRedundancyTransitionResponseBody struct {
	// The container in which the redundancy type conversion task is stored.
	BucketDataRedundancyTransition *CreateBucketDataRedundancyTransitionResponseBodyBucketDataRedundancyTransition `json:"BucketDataRedundancyTransition,omitempty" xml:"BucketDataRedundancyTransition,omitempty" type:"Struct"`
}

func (s CreateBucketDataRedundancyTransitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBucketDataRedundancyTransitionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBucketDataRedundancyTransitionResponseBody) GetBucketDataRedundancyTransition() *CreateBucketDataRedundancyTransitionResponseBodyBucketDataRedundancyTransition {
	return s.BucketDataRedundancyTransition
}

func (s *CreateBucketDataRedundancyTransitionResponseBody) SetBucketDataRedundancyTransition(v *CreateBucketDataRedundancyTransitionResponseBodyBucketDataRedundancyTransition) *CreateBucketDataRedundancyTransitionResponseBody {
	s.BucketDataRedundancyTransition = v
	return s
}

func (s *CreateBucketDataRedundancyTransitionResponseBody) Validate() error {
	if s.BucketDataRedundancyTransition != nil {
		if err := s.BucketDataRedundancyTransition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateBucketDataRedundancyTransitionResponseBodyBucketDataRedundancyTransition struct {
	// The ID of the redundancy type conversion task. The ID can be used to view and delete the redundancy type conversion task.
	//
	// example:
	//
	// 4be5beb0f74f490186311b268bf6****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateBucketDataRedundancyTransitionResponseBodyBucketDataRedundancyTransition) String() string {
	return dara.Prettify(s)
}

func (s CreateBucketDataRedundancyTransitionResponseBodyBucketDataRedundancyTransition) GoString() string {
	return s.String()
}

func (s *CreateBucketDataRedundancyTransitionResponseBodyBucketDataRedundancyTransition) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateBucketDataRedundancyTransitionResponseBodyBucketDataRedundancyTransition) SetTaskId(v string) *CreateBucketDataRedundancyTransitionResponseBodyBucketDataRedundancyTransition {
	s.TaskId = &v
	return s
}

func (s *CreateBucketDataRedundancyTransitionResponseBodyBucketDataRedundancyTransition) Validate() error {
	return dara.Validate(s)
}
