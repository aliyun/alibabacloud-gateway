// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataLakeStorageTransferJobs interface {
	dara.Model
	String() string
	GoString() string
	SetDataLakeStorageTransferJob(v []*DataLakeStorageTransferJob) *DataLakeStorageTransferJobs
	GetDataLakeStorageTransferJob() []*DataLakeStorageTransferJob
	SetNextMarkerBucket(v string) *DataLakeStorageTransferJobs
	GetNextMarkerBucket() *string
	SetNextMarkerJobId(v string) *DataLakeStorageTransferJobs
	GetNextMarkerJobId() *string
	SetTruncated(v string) *DataLakeStorageTransferJobs
	GetTruncated() *string
}

type DataLakeStorageTransferJobs struct {
	DataLakeStorageTransferJob []*DataLakeStorageTransferJob `json:"DataLakeStorageTransferJob,omitempty" xml:"DataLakeStorageTransferJob,omitempty" type:"Repeated"`
	// example:
	//
	// test-bucket
	NextMarkerBucket *string `json:"NextMarkerBucket,omitempty" xml:"NextMarkerBucket,omitempty"`
	// example:
	//
	// abcdef03370a4b32ac6bc69eb1123456
	NextMarkerJobId *string `json:"NextMarkerJobId,omitempty" xml:"NextMarkerJobId,omitempty"`
	// example:
	//
	// false
	Truncated *string `json:"Truncated,omitempty" xml:"Truncated,omitempty"`
}

func (s DataLakeStorageTransferJobs) String() string {
	return dara.Prettify(s)
}

func (s DataLakeStorageTransferJobs) GoString() string {
	return s.String()
}

func (s *DataLakeStorageTransferJobs) GetDataLakeStorageTransferJob() []*DataLakeStorageTransferJob {
	return s.DataLakeStorageTransferJob
}

func (s *DataLakeStorageTransferJobs) GetNextMarkerBucket() *string {
	return s.NextMarkerBucket
}

func (s *DataLakeStorageTransferJobs) GetNextMarkerJobId() *string {
	return s.NextMarkerJobId
}

func (s *DataLakeStorageTransferJobs) GetTruncated() *string {
	return s.Truncated
}

func (s *DataLakeStorageTransferJobs) SetDataLakeStorageTransferJob(v []*DataLakeStorageTransferJob) *DataLakeStorageTransferJobs {
	s.DataLakeStorageTransferJob = v
	return s
}

func (s *DataLakeStorageTransferJobs) SetNextMarkerBucket(v string) *DataLakeStorageTransferJobs {
	s.NextMarkerBucket = &v
	return s
}

func (s *DataLakeStorageTransferJobs) SetNextMarkerJobId(v string) *DataLakeStorageTransferJobs {
	s.NextMarkerJobId = &v
	return s
}

func (s *DataLakeStorageTransferJobs) SetTruncated(v string) *DataLakeStorageTransferJobs {
	s.Truncated = &v
	return s
}

func (s *DataLakeStorageTransferJobs) Validate() error {
	if s.DataLakeStorageTransferJob != nil {
		for _, item := range s.DataLakeStorageTransferJob {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
