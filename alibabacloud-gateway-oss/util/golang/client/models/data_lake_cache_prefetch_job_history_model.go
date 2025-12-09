// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataLakeCachePrefetchJobHistory interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DataLakeCachePrefetchJobHistory
	GetEndTime() *int64
	SetId(v string) *DataLakeCachePrefetchJobHistory
	GetId() *string
	SetJobId(v string) *DataLakeCachePrefetchJobHistory
	GetJobId() *string
	SetStartTime(v int64) *DataLakeCachePrefetchJobHistory
	GetStartTime() *int64
	SetStatus(v string) *DataLakeCachePrefetchJobHistory
	GetStatus() *string
	SetSucceedCount(v int64) *DataLakeCachePrefetchJobHistory
	GetSucceedCount() *int64
	SetTotalCount(v int64) *DataLakeCachePrefetchJobHistory
	GetTotalCount() *int64
}

type DataLakeCachePrefetchJobHistory struct {
	// example:
	//
	// 1727176823
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 8w643e67a54c4670b5ed321511234567
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 345sdf60df1329d88482912343ea74g2
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 1727176463
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// REPLICATION_JOB_SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 0
	SucceedCount *int64 `json:"SucceedCount,omitempty" xml:"SucceedCount,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DataLakeCachePrefetchJobHistory) String() string {
	return dara.Prettify(s)
}

func (s DataLakeCachePrefetchJobHistory) GoString() string {
	return s.String()
}

func (s *DataLakeCachePrefetchJobHistory) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DataLakeCachePrefetchJobHistory) GetId() *string {
	return s.Id
}

func (s *DataLakeCachePrefetchJobHistory) GetJobId() *string {
	return s.JobId
}

func (s *DataLakeCachePrefetchJobHistory) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DataLakeCachePrefetchJobHistory) GetStatus() *string {
	return s.Status
}

func (s *DataLakeCachePrefetchJobHistory) GetSucceedCount() *int64 {
	return s.SucceedCount
}

func (s *DataLakeCachePrefetchJobHistory) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DataLakeCachePrefetchJobHistory) SetEndTime(v int64) *DataLakeCachePrefetchJobHistory {
	s.EndTime = &v
	return s
}

func (s *DataLakeCachePrefetchJobHistory) SetId(v string) *DataLakeCachePrefetchJobHistory {
	s.Id = &v
	return s
}

func (s *DataLakeCachePrefetchJobHistory) SetJobId(v string) *DataLakeCachePrefetchJobHistory {
	s.JobId = &v
	return s
}

func (s *DataLakeCachePrefetchJobHistory) SetStartTime(v int64) *DataLakeCachePrefetchJobHistory {
	s.StartTime = &v
	return s
}

func (s *DataLakeCachePrefetchJobHistory) SetStatus(v string) *DataLakeCachePrefetchJobHistory {
	s.Status = &v
	return s
}

func (s *DataLakeCachePrefetchJobHistory) SetSucceedCount(v int64) *DataLakeCachePrefetchJobHistory {
	s.SucceedCount = &v
	return s
}

func (s *DataLakeCachePrefetchJobHistory) SetTotalCount(v int64) *DataLakeCachePrefetchJobHistory {
	s.TotalCount = &v
	return s
}

func (s *DataLakeCachePrefetchJobHistory) Validate() error {
	return dara.Validate(s)
}
