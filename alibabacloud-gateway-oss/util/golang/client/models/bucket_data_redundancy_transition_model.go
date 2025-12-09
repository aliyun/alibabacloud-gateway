// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBucketDataRedundancyTransition interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *BucketDataRedundancyTransition
	GetBucket() *string
	SetCreateTime(v string) *BucketDataRedundancyTransition
	GetCreateTime() *string
	SetEndTime(v string) *BucketDataRedundancyTransition
	GetEndTime() *string
	SetEstimatedRemainingTime(v string) *BucketDataRedundancyTransition
	GetEstimatedRemainingTime() *string
	SetProcessPercentage(v int32) *BucketDataRedundancyTransition
	GetProcessPercentage() *int32
	SetStartTime(v string) *BucketDataRedundancyTransition
	GetStartTime() *string
	SetStatus(v string) *BucketDataRedundancyTransition
	GetStatus() *string
	SetTaskId(v string) *BucketDataRedundancyTransition
	GetTaskId() *string
}

type BucketDataRedundancyTransition struct {
	// example:
	//
	// example-bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// example:
	//
	// 2023-11-17T09:14:39.000Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2023-11-17T09:14:39.000Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 10
	EstimatedRemainingTime *string `json:"EstimatedRemainingTime,omitempty" xml:"EstimatedRemainingTime,omitempty"`
	// example:
	//
	// 88
	ProcessPercentage *int32 `json:"ProcessPercentage,omitempty" xml:"ProcessPercentage,omitempty"`
	// example:
	//
	// 2023-11-17T09:14:39.000Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Queueing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 751f5243f8ac4ae89f34726534d1****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s BucketDataRedundancyTransition) String() string {
	return dara.Prettify(s)
}

func (s BucketDataRedundancyTransition) GoString() string {
	return s.String()
}

func (s *BucketDataRedundancyTransition) GetBucket() *string {
	return s.Bucket
}

func (s *BucketDataRedundancyTransition) GetCreateTime() *string {
	return s.CreateTime
}

func (s *BucketDataRedundancyTransition) GetEndTime() *string {
	return s.EndTime
}

func (s *BucketDataRedundancyTransition) GetEstimatedRemainingTime() *string {
	return s.EstimatedRemainingTime
}

func (s *BucketDataRedundancyTransition) GetProcessPercentage() *int32 {
	return s.ProcessPercentage
}

func (s *BucketDataRedundancyTransition) GetStartTime() *string {
	return s.StartTime
}

func (s *BucketDataRedundancyTransition) GetStatus() *string {
	return s.Status
}

func (s *BucketDataRedundancyTransition) GetTaskId() *string {
	return s.TaskId
}

func (s *BucketDataRedundancyTransition) SetBucket(v string) *BucketDataRedundancyTransition {
	s.Bucket = &v
	return s
}

func (s *BucketDataRedundancyTransition) SetCreateTime(v string) *BucketDataRedundancyTransition {
	s.CreateTime = &v
	return s
}

func (s *BucketDataRedundancyTransition) SetEndTime(v string) *BucketDataRedundancyTransition {
	s.EndTime = &v
	return s
}

func (s *BucketDataRedundancyTransition) SetEstimatedRemainingTime(v string) *BucketDataRedundancyTransition {
	s.EstimatedRemainingTime = &v
	return s
}

func (s *BucketDataRedundancyTransition) SetProcessPercentage(v int32) *BucketDataRedundancyTransition {
	s.ProcessPercentage = &v
	return s
}

func (s *BucketDataRedundancyTransition) SetStartTime(v string) *BucketDataRedundancyTransition {
	s.StartTime = &v
	return s
}

func (s *BucketDataRedundancyTransition) SetStatus(v string) *BucketDataRedundancyTransition {
	s.Status = &v
	return s
}

func (s *BucketDataRedundancyTransition) SetTaskId(v string) *BucketDataRedundancyTransition {
	s.TaskId = &v
	return s
}

func (s *BucketDataRedundancyTransition) Validate() error {
	return dara.Validate(s)
}
