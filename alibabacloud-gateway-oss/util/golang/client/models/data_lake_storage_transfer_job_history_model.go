// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataLakeStorageTransferJobHistory interface {
	dara.Model
	String() string
	GoString() string
	SetDetailInfo(v *DataLakeStorageTransferJobHistoryDetailInfo) *DataLakeStorageTransferJobHistory
	GetDetailInfo() *DataLakeStorageTransferJobHistoryDetailInfo
	SetEndTime(v int64) *DataLakeStorageTransferJobHistory
	GetEndTime() *int64
	SetId(v string) *DataLakeStorageTransferJobHistory
	GetId() *string
	SetJobId(v string) *DataLakeStorageTransferJobHistory
	GetJobId() *string
	SetStartTime(v int64) *DataLakeStorageTransferJobHistory
	GetStartTime() *int64
	SetStatus(v string) *DataLakeStorageTransferJobHistory
	GetStatus() *string
	SetSucceedCount(v int64) *DataLakeStorageTransferJobHistory
	GetSucceedCount() *int64
	SetTotalCount(v int64) *DataLakeStorageTransferJobHistory
	GetTotalCount() *int64
}

type DataLakeStorageTransferJobHistory struct {
	DetailInfo *DataLakeStorageTransferJobHistoryDetailInfo `json:"DetailInfo,omitempty" xml:"DetailInfo,omitempty" type:"Struct"`
	// example:
	//
	// 1731830653
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// abcdef03370a4b32ac6bc69eb1123456
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 43452
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 1731830653
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// REPLICATION_JOB_SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 12345
	SucceedCount *int64 `json:"SucceedCount,omitempty" xml:"SucceedCount,omitempty"`
	// example:
	//
	// 12
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DataLakeStorageTransferJobHistory) String() string {
	return dara.Prettify(s)
}

func (s DataLakeStorageTransferJobHistory) GoString() string {
	return s.String()
}

func (s *DataLakeStorageTransferJobHistory) GetDetailInfo() *DataLakeStorageTransferJobHistoryDetailInfo {
	return s.DetailInfo
}

func (s *DataLakeStorageTransferJobHistory) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DataLakeStorageTransferJobHistory) GetId() *string {
	return s.Id
}

func (s *DataLakeStorageTransferJobHistory) GetJobId() *string {
	return s.JobId
}

func (s *DataLakeStorageTransferJobHistory) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DataLakeStorageTransferJobHistory) GetStatus() *string {
	return s.Status
}

func (s *DataLakeStorageTransferJobHistory) GetSucceedCount() *int64 {
	return s.SucceedCount
}

func (s *DataLakeStorageTransferJobHistory) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DataLakeStorageTransferJobHistory) SetDetailInfo(v *DataLakeStorageTransferJobHistoryDetailInfo) *DataLakeStorageTransferJobHistory {
	s.DetailInfo = v
	return s
}

func (s *DataLakeStorageTransferJobHistory) SetEndTime(v int64) *DataLakeStorageTransferJobHistory {
	s.EndTime = &v
	return s
}

func (s *DataLakeStorageTransferJobHistory) SetId(v string) *DataLakeStorageTransferJobHistory {
	s.Id = &v
	return s
}

func (s *DataLakeStorageTransferJobHistory) SetJobId(v string) *DataLakeStorageTransferJobHistory {
	s.JobId = &v
	return s
}

func (s *DataLakeStorageTransferJobHistory) SetStartTime(v int64) *DataLakeStorageTransferJobHistory {
	s.StartTime = &v
	return s
}

func (s *DataLakeStorageTransferJobHistory) SetStatus(v string) *DataLakeStorageTransferJobHistory {
	s.Status = &v
	return s
}

func (s *DataLakeStorageTransferJobHistory) SetSucceedCount(v int64) *DataLakeStorageTransferJobHistory {
	s.SucceedCount = &v
	return s
}

func (s *DataLakeStorageTransferJobHistory) SetTotalCount(v int64) *DataLakeStorageTransferJobHistory {
	s.TotalCount = &v
	return s
}

func (s *DataLakeStorageTransferJobHistory) Validate() error {
	if s.DetailInfo != nil {
		if err := s.DetailInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DataLakeStorageTransferJobHistoryDetailInfo struct {
	// example:
	//
	// NotSupported
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 0
	HDFSFailedCount *int64 `json:"HDFSFailedCount,omitempty" xml:"HDFSFailedCount,omitempty"`
	// example:
	//
	// abc/
	HDFSTransferDataDir *string `json:"HDFSTransferDataDir,omitempty" xml:"HDFSTransferDataDir,omitempty"`
	// example:
	//
	// hdfs-error/
	HDFSTransferErrInfoDir *string `json:"HDFSTransferErrInfoDir,omitempty" xml:"HDFSTransferErrInfoDir,omitempty"`
	// example:
	//
	// meta-dta/
	HDFSTransferImportMetaDir *string `json:"HDFSTransferImportMetaDir,omitempty" xml:"HDFSTransferImportMetaDir,omitempty"`
	// example:
	//
	// abcdef03370a4b32ac6bc69eb1123456
	HDFSTransferJobId *string `json:"HDFSTransferJobId,omitempty" xml:"HDFSTransferJobId,omitempty"`
	// example:
	//
	// hdfs-logs/
	LogBaseDir *string `json:"LogBaseDir,omitempty" xml:"LogBaseDir,omitempty"`
	// example:
	//
	// verify-error/
	VerifyErrInfoDir *string `json:"VerifyErrInfoDir,omitempty" xml:"VerifyErrInfoDir,omitempty"`
	// example:
	//
	// VERIFY_SUCCEED
	VerifyStatus *string `json:"VerifyStatus,omitempty" xml:"VerifyStatus,omitempty"`
	// example:
	//
	// 123456
	VerifyTotalCount *int64 `json:"VerifyTotalCount,omitempty" xml:"VerifyTotalCount,omitempty"`
}

func (s DataLakeStorageTransferJobHistoryDetailInfo) String() string {
	return dara.Prettify(s)
}

func (s DataLakeStorageTransferJobHistoryDetailInfo) GoString() string {
	return s.String()
}

func (s *DataLakeStorageTransferJobHistoryDetailInfo) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *DataLakeStorageTransferJobHistoryDetailInfo) GetHDFSFailedCount() *int64 {
	return s.HDFSFailedCount
}

func (s *DataLakeStorageTransferJobHistoryDetailInfo) GetHDFSTransferDataDir() *string {
	return s.HDFSTransferDataDir
}

func (s *DataLakeStorageTransferJobHistoryDetailInfo) GetHDFSTransferErrInfoDir() *string {
	return s.HDFSTransferErrInfoDir
}

func (s *DataLakeStorageTransferJobHistoryDetailInfo) GetHDFSTransferImportMetaDir() *string {
	return s.HDFSTransferImportMetaDir
}

func (s *DataLakeStorageTransferJobHistoryDetailInfo) GetHDFSTransferJobId() *string {
	return s.HDFSTransferJobId
}

func (s *DataLakeStorageTransferJobHistoryDetailInfo) GetLogBaseDir() *string {
	return s.LogBaseDir
}

func (s *DataLakeStorageTransferJobHistoryDetailInfo) GetVerifyErrInfoDir() *string {
	return s.VerifyErrInfoDir
}

func (s *DataLakeStorageTransferJobHistoryDetailInfo) GetVerifyStatus() *string {
	return s.VerifyStatus
}

func (s *DataLakeStorageTransferJobHistoryDetailInfo) GetVerifyTotalCount() *int64 {
	return s.VerifyTotalCount
}

func (s *DataLakeStorageTransferJobHistoryDetailInfo) SetErrorMsg(v string) *DataLakeStorageTransferJobHistoryDetailInfo {
	s.ErrorMsg = &v
	return s
}

func (s *DataLakeStorageTransferJobHistoryDetailInfo) SetHDFSFailedCount(v int64) *DataLakeStorageTransferJobHistoryDetailInfo {
	s.HDFSFailedCount = &v
	return s
}

func (s *DataLakeStorageTransferJobHistoryDetailInfo) SetHDFSTransferDataDir(v string) *DataLakeStorageTransferJobHistoryDetailInfo {
	s.HDFSTransferDataDir = &v
	return s
}

func (s *DataLakeStorageTransferJobHistoryDetailInfo) SetHDFSTransferErrInfoDir(v string) *DataLakeStorageTransferJobHistoryDetailInfo {
	s.HDFSTransferErrInfoDir = &v
	return s
}

func (s *DataLakeStorageTransferJobHistoryDetailInfo) SetHDFSTransferImportMetaDir(v string) *DataLakeStorageTransferJobHistoryDetailInfo {
	s.HDFSTransferImportMetaDir = &v
	return s
}

func (s *DataLakeStorageTransferJobHistoryDetailInfo) SetHDFSTransferJobId(v string) *DataLakeStorageTransferJobHistoryDetailInfo {
	s.HDFSTransferJobId = &v
	return s
}

func (s *DataLakeStorageTransferJobHistoryDetailInfo) SetLogBaseDir(v string) *DataLakeStorageTransferJobHistoryDetailInfo {
	s.LogBaseDir = &v
	return s
}

func (s *DataLakeStorageTransferJobHistoryDetailInfo) SetVerifyErrInfoDir(v string) *DataLakeStorageTransferJobHistoryDetailInfo {
	s.VerifyErrInfoDir = &v
	return s
}

func (s *DataLakeStorageTransferJobHistoryDetailInfo) SetVerifyStatus(v string) *DataLakeStorageTransferJobHistoryDetailInfo {
	s.VerifyStatus = &v
	return s
}

func (s *DataLakeStorageTransferJobHistoryDetailInfo) SetVerifyTotalCount(v int64) *DataLakeStorageTransferJobHistoryDetailInfo {
	s.VerifyTotalCount = &v
	return s
}

func (s *DataLakeStorageTransferJobHistoryDetailInfo) Validate() error {
	return dara.Validate(s)
}
