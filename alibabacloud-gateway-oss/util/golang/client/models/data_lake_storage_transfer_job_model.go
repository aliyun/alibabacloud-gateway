// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataLakeStorageTransferJob interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *DataLakeStorageTransferJob
	GetBucket() *string
	SetCreateTime(v int64) *DataLakeStorageTransferJob
	GetCreateTime() *int64
	SetHistoryId(v string) *DataLakeStorageTransferJob
	GetHistoryId() *string
	SetId(v string) *DataLakeStorageTransferJob
	GetId() *string
	SetLastModifyTime(v int64) *DataLakeStorageTransferJob
	GetLastModifyTime() *int64
	SetProgressInfo(v *DataLakeStorageTransferJobProgressInfo) *DataLakeStorageTransferJob
	GetProgressInfo() *DataLakeStorageTransferJobProgressInfo
	SetRule(v *DataLakeStorageTransferJobRule) *DataLakeStorageTransferJob
	GetRule() *DataLakeStorageTransferJobRule
	SetStatus(v string) *DataLakeStorageTransferJob
	GetStatus() *string
	SetType(v int32) *DataLakeStorageTransferJob
	GetType() *int32
}

type DataLakeStorageTransferJob struct {
	// example:
	//
	// bucket1
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// example:
	//
	// 1727162332
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// abcdef87e3c04af2af290ec52d123456
	HistoryId *string `json:"HistoryId,omitempty" xml:"HistoryId,omitempty"`
	// example:
	//
	// abcdef87e3c04af2af290ec52d123456
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1727164655
	LastModifyTime *int64                                  `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	ProgressInfo   *DataLakeStorageTransferJobProgressInfo `json:"ProgressInfo,omitempty" xml:"ProgressInfo,omitempty" type:"Struct"`
	Rule           *DataLakeStorageTransferJobRule         `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// example:
	//
	// REPLICATION_JOB_IDLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DataLakeStorageTransferJob) String() string {
	return dara.Prettify(s)
}

func (s DataLakeStorageTransferJob) GoString() string {
	return s.String()
}

func (s *DataLakeStorageTransferJob) GetBucket() *string {
	return s.Bucket
}

func (s *DataLakeStorageTransferJob) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DataLakeStorageTransferJob) GetHistoryId() *string {
	return s.HistoryId
}

func (s *DataLakeStorageTransferJob) GetId() *string {
	return s.Id
}

func (s *DataLakeStorageTransferJob) GetLastModifyTime() *int64 {
	return s.LastModifyTime
}

func (s *DataLakeStorageTransferJob) GetProgressInfo() *DataLakeStorageTransferJobProgressInfo {
	return s.ProgressInfo
}

func (s *DataLakeStorageTransferJob) GetRule() *DataLakeStorageTransferJobRule {
	return s.Rule
}

func (s *DataLakeStorageTransferJob) GetStatus() *string {
	return s.Status
}

func (s *DataLakeStorageTransferJob) GetType() *int32 {
	return s.Type
}

func (s *DataLakeStorageTransferJob) SetBucket(v string) *DataLakeStorageTransferJob {
	s.Bucket = &v
	return s
}

func (s *DataLakeStorageTransferJob) SetCreateTime(v int64) *DataLakeStorageTransferJob {
	s.CreateTime = &v
	return s
}

func (s *DataLakeStorageTransferJob) SetHistoryId(v string) *DataLakeStorageTransferJob {
	s.HistoryId = &v
	return s
}

func (s *DataLakeStorageTransferJob) SetId(v string) *DataLakeStorageTransferJob {
	s.Id = &v
	return s
}

func (s *DataLakeStorageTransferJob) SetLastModifyTime(v int64) *DataLakeStorageTransferJob {
	s.LastModifyTime = &v
	return s
}

func (s *DataLakeStorageTransferJob) SetProgressInfo(v *DataLakeStorageTransferJobProgressInfo) *DataLakeStorageTransferJob {
	s.ProgressInfo = v
	return s
}

func (s *DataLakeStorageTransferJob) SetRule(v *DataLakeStorageTransferJobRule) *DataLakeStorageTransferJob {
	s.Rule = v
	return s
}

func (s *DataLakeStorageTransferJob) SetStatus(v string) *DataLakeStorageTransferJob {
	s.Status = &v
	return s
}

func (s *DataLakeStorageTransferJob) SetType(v int32) *DataLakeStorageTransferJob {
	s.Type = &v
	return s
}

func (s *DataLakeStorageTransferJob) Validate() error {
	if s.ProgressInfo != nil {
		if err := s.ProgressInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Rule != nil {
		if err := s.Rule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DataLakeStorageTransferJobProgressInfo struct {
	// example:
	//
	// 50
	Percent *int64 `json:"Percent,omitempty" xml:"Percent,omitempty"`
}

func (s DataLakeStorageTransferJobProgressInfo) String() string {
	return dara.Prettify(s)
}

func (s DataLakeStorageTransferJobProgressInfo) GoString() string {
	return s.String()
}

func (s *DataLakeStorageTransferJobProgressInfo) GetPercent() *int64 {
	return s.Percent
}

func (s *DataLakeStorageTransferJobProgressInfo) SetPercent(v int64) *DataLakeStorageTransferJobProgressInfo {
	s.Percent = &v
	return s
}

func (s *DataLakeStorageTransferJobProgressInfo) Validate() error {
	return dara.Validate(s)
}
