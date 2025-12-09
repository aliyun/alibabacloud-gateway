// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDataLakeCachePrefetchJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXOssDatalakeJobId(v string) *StopDataLakeCachePrefetchJobRequest
	GetXOssDatalakeJobId() *string
}

type StopDataLakeCachePrefetchJobRequest struct {
	// This parameter is required.
	XOssDatalakeJobId *string `json:"x-oss-datalake-job-id,omitempty" xml:"x-oss-datalake-job-id,omitempty"`
}

func (s StopDataLakeCachePrefetchJobRequest) String() string {
	return dara.Prettify(s)
}

func (s StopDataLakeCachePrefetchJobRequest) GoString() string {
	return s.String()
}

func (s *StopDataLakeCachePrefetchJobRequest) GetXOssDatalakeJobId() *string {
	return s.XOssDatalakeJobId
}

func (s *StopDataLakeCachePrefetchJobRequest) SetXOssDatalakeJobId(v string) *StopDataLakeCachePrefetchJobRequest {
	s.XOssDatalakeJobId = &v
	return s
}

func (s *StopDataLakeCachePrefetchJobRequest) Validate() error {
	return dara.Validate(s)
}
