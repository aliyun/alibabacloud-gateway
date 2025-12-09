// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataLakeCachePrefetchJob interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *DataLakeCachePrefetchJob
	GetBucket() *string
	SetCreateTime(v int64) *DataLakeCachePrefetchJob
	GetCreateTime() *int64
	SetHistoryId(v string) *DataLakeCachePrefetchJob
	GetHistoryId() *string
	SetId(v string) *DataLakeCachePrefetchJob
	GetId() *string
	SetLastModifyTime(v int64) *DataLakeCachePrefetchJob
	GetLastModifyTime() *int64
	SetRule(v *DataLakeCachePrefetchJobRule) *DataLakeCachePrefetchJob
	GetRule() *DataLakeCachePrefetchJobRule
	SetStatus(v string) *DataLakeCachePrefetchJob
	GetStatus() *string
	SetType(v int32) *DataLakeCachePrefetchJob
	GetType() *int32
}

type DataLakeCachePrefetchJob struct {
	// example:
	//
	// bucket-example
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// example:
	//
	// 1727162332
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 84dcbfa10wdp488211dc6cb287f4d804
	HistoryId *string `json:"HistoryId,omitempty" xml:"HistoryId,omitempty"`
	// example:
	//
	// 84dcbfa10wdp488211dc6cb287f4d804
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1727164655
	LastModifyTime *int64                        `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	Rule           *DataLakeCachePrefetchJobRule `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// example:
	//
	// REPLICATION_JOB_IDLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DataLakeCachePrefetchJob) String() string {
	return dara.Prettify(s)
}

func (s DataLakeCachePrefetchJob) GoString() string {
	return s.String()
}

func (s *DataLakeCachePrefetchJob) GetBucket() *string {
	return s.Bucket
}

func (s *DataLakeCachePrefetchJob) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DataLakeCachePrefetchJob) GetHistoryId() *string {
	return s.HistoryId
}

func (s *DataLakeCachePrefetchJob) GetId() *string {
	return s.Id
}

func (s *DataLakeCachePrefetchJob) GetLastModifyTime() *int64 {
	return s.LastModifyTime
}

func (s *DataLakeCachePrefetchJob) GetRule() *DataLakeCachePrefetchJobRule {
	return s.Rule
}

func (s *DataLakeCachePrefetchJob) GetStatus() *string {
	return s.Status
}

func (s *DataLakeCachePrefetchJob) GetType() *int32 {
	return s.Type
}

func (s *DataLakeCachePrefetchJob) SetBucket(v string) *DataLakeCachePrefetchJob {
	s.Bucket = &v
	return s
}

func (s *DataLakeCachePrefetchJob) SetCreateTime(v int64) *DataLakeCachePrefetchJob {
	s.CreateTime = &v
	return s
}

func (s *DataLakeCachePrefetchJob) SetHistoryId(v string) *DataLakeCachePrefetchJob {
	s.HistoryId = &v
	return s
}

func (s *DataLakeCachePrefetchJob) SetId(v string) *DataLakeCachePrefetchJob {
	s.Id = &v
	return s
}

func (s *DataLakeCachePrefetchJob) SetLastModifyTime(v int64) *DataLakeCachePrefetchJob {
	s.LastModifyTime = &v
	return s
}

func (s *DataLakeCachePrefetchJob) SetRule(v *DataLakeCachePrefetchJobRule) *DataLakeCachePrefetchJob {
	s.Rule = v
	return s
}

func (s *DataLakeCachePrefetchJob) SetStatus(v string) *DataLakeCachePrefetchJob {
	s.Status = &v
	return s
}

func (s *DataLakeCachePrefetchJob) SetType(v int32) *DataLakeCachePrefetchJob {
	s.Type = &v
	return s
}

func (s *DataLakeCachePrefetchJob) Validate() error {
	if s.Rule != nil {
		if err := s.Rule.Validate(); err != nil {
			return err
		}
	}
	return nil
}
