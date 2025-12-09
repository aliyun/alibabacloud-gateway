// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataLakeCachePrefetchJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataLakeCachePrefetchJob(v *DataLakeCachePrefetchJob) *GetDataLakeCachePrefetchJobResponseBody
	GetDataLakeCachePrefetchJob() *DataLakeCachePrefetchJob
}

type GetDataLakeCachePrefetchJobResponseBody struct {
	DataLakeCachePrefetchJob *DataLakeCachePrefetchJob `json:"DataLakeCachePrefetchJob,omitempty" xml:"DataLakeCachePrefetchJob,omitempty"`
}

func (s GetDataLakeCachePrefetchJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataLakeCachePrefetchJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataLakeCachePrefetchJobResponseBody) GetDataLakeCachePrefetchJob() *DataLakeCachePrefetchJob {
	return s.DataLakeCachePrefetchJob
}

func (s *GetDataLakeCachePrefetchJobResponseBody) SetDataLakeCachePrefetchJob(v *DataLakeCachePrefetchJob) *GetDataLakeCachePrefetchJobResponseBody {
	s.DataLakeCachePrefetchJob = v
	return s
}

func (s *GetDataLakeCachePrefetchJobResponseBody) Validate() error {
	if s.DataLakeCachePrefetchJob != nil {
		if err := s.DataLakeCachePrefetchJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}
