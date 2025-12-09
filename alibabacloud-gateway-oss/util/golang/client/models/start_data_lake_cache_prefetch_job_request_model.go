// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDataLakeCachePrefetchJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXOssDatalakeJobId(v string) *StartDataLakeCachePrefetchJobRequest
	GetXOssDatalakeJobId() *string
}

type StartDataLakeCachePrefetchJobRequest struct {
	// This parameter is required.
	XOssDatalakeJobId *string `json:"x-oss-datalake-job-id,omitempty" xml:"x-oss-datalake-job-id,omitempty"`
}

func (s StartDataLakeCachePrefetchJobRequest) String() string {
	return dara.Prettify(s)
}

func (s StartDataLakeCachePrefetchJobRequest) GoString() string {
	return s.String()
}

func (s *StartDataLakeCachePrefetchJobRequest) GetXOssDatalakeJobId() *string {
	return s.XOssDatalakeJobId
}

func (s *StartDataLakeCachePrefetchJobRequest) SetXOssDatalakeJobId(v string) *StartDataLakeCachePrefetchJobRequest {
	s.XOssDatalakeJobId = &v
	return s
}

func (s *StartDataLakeCachePrefetchJobRequest) Validate() error {
	return dara.Validate(s)
}
