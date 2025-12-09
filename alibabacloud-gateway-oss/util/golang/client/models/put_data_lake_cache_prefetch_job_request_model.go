// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutDataLakeCachePrefetchJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateDataLakeCachePrefetchJob(v *CreateDataLakeCachePrefetchJob) *PutDataLakeCachePrefetchJobRequest
	GetCreateDataLakeCachePrefetchJob() *CreateDataLakeCachePrefetchJob
	SetXOssDatalakeJobId(v string) *PutDataLakeCachePrefetchJobRequest
	GetXOssDatalakeJobId() *string
}

type PutDataLakeCachePrefetchJobRequest struct {
	CreateDataLakeCachePrefetchJob *CreateDataLakeCachePrefetchJob `json:"CreateDataLakeCachePrefetchJob,omitempty" xml:"CreateDataLakeCachePrefetchJob,omitempty"`
	XOssDatalakeJobId              *string                         `json:"x-oss-datalake-job-id,omitempty" xml:"x-oss-datalake-job-id,omitempty"`
}

func (s PutDataLakeCachePrefetchJobRequest) String() string {
	return dara.Prettify(s)
}

func (s PutDataLakeCachePrefetchJobRequest) GoString() string {
	return s.String()
}

func (s *PutDataLakeCachePrefetchJobRequest) GetCreateDataLakeCachePrefetchJob() *CreateDataLakeCachePrefetchJob {
	return s.CreateDataLakeCachePrefetchJob
}

func (s *PutDataLakeCachePrefetchJobRequest) GetXOssDatalakeJobId() *string {
	return s.XOssDatalakeJobId
}

func (s *PutDataLakeCachePrefetchJobRequest) SetCreateDataLakeCachePrefetchJob(v *CreateDataLakeCachePrefetchJob) *PutDataLakeCachePrefetchJobRequest {
	s.CreateDataLakeCachePrefetchJob = v
	return s
}

func (s *PutDataLakeCachePrefetchJobRequest) SetXOssDatalakeJobId(v string) *PutDataLakeCachePrefetchJobRequest {
	s.XOssDatalakeJobId = &v
	return s
}

func (s *PutDataLakeCachePrefetchJobRequest) Validate() error {
	if s.CreateDataLakeCachePrefetchJob != nil {
		if err := s.CreateDataLakeCachePrefetchJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}
