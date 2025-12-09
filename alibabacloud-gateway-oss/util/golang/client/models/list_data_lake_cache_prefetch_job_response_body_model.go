// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeCachePrefetchJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataLakeCachePrefetchJobs(v *ListDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobs) *ListDataLakeCachePrefetchJobResponseBody
	GetDataLakeCachePrefetchJobs() *ListDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobs
}

type ListDataLakeCachePrefetchJobResponseBody struct {
	DataLakeCachePrefetchJobs *ListDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobs `json:"DataLakeCachePrefetchJobs,omitempty" xml:"DataLakeCachePrefetchJobs,omitempty" type:"Struct"`
}

func (s ListDataLakeCachePrefetchJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeCachePrefetchJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataLakeCachePrefetchJobResponseBody) GetDataLakeCachePrefetchJobs() *ListDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobs {
	return s.DataLakeCachePrefetchJobs
}

func (s *ListDataLakeCachePrefetchJobResponseBody) SetDataLakeCachePrefetchJobs(v *ListDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobs) *ListDataLakeCachePrefetchJobResponseBody {
	s.DataLakeCachePrefetchJobs = v
	return s
}

func (s *ListDataLakeCachePrefetchJobResponseBody) Validate() error {
	if s.DataLakeCachePrefetchJobs != nil {
		if err := s.DataLakeCachePrefetchJobs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobs struct {
	DataLakeCachePrefetchJob *DataLakeCachePrefetchJob `json:"DataLakeCachePrefetchJob,omitempty" xml:"DataLakeCachePrefetchJob,omitempty"`
	NextMarkerBucket         *string                   `json:"NextMarkerBucket,omitempty" xml:"NextMarkerBucket,omitempty"`
	NextMarkerJobId          *string                   `json:"NextMarkerJobId,omitempty" xml:"NextMarkerJobId,omitempty"`
	Truncated                *bool                     `json:"Truncated,omitempty" xml:"Truncated,omitempty"`
}

func (s ListDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobs) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobs) GoString() string {
	return s.String()
}

func (s *ListDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobs) GetDataLakeCachePrefetchJob() *DataLakeCachePrefetchJob {
	return s.DataLakeCachePrefetchJob
}

func (s *ListDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobs) GetNextMarkerBucket() *string {
	return s.NextMarkerBucket
}

func (s *ListDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobs) GetNextMarkerJobId() *string {
	return s.NextMarkerJobId
}

func (s *ListDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobs) GetTruncated() *bool {
	return s.Truncated
}

func (s *ListDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobs) SetDataLakeCachePrefetchJob(v *DataLakeCachePrefetchJob) *ListDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobs {
	s.DataLakeCachePrefetchJob = v
	return s
}

func (s *ListDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobs) SetNextMarkerBucket(v string) *ListDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobs {
	s.NextMarkerBucket = &v
	return s
}

func (s *ListDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobs) SetNextMarkerJobId(v string) *ListDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobs {
	s.NextMarkerJobId = &v
	return s
}

func (s *ListDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobs) SetTruncated(v bool) *ListDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobs {
	s.Truncated = &v
	return s
}

func (s *ListDataLakeCachePrefetchJobResponseBodyDataLakeCachePrefetchJobs) Validate() error {
	if s.DataLakeCachePrefetchJob != nil {
		if err := s.DataLakeCachePrefetchJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}
