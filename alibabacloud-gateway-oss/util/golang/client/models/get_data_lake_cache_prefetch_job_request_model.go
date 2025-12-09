// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataLakeCachePrefetchJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXOssDatalakeJobId(v string) *GetDataLakeCachePrefetchJobRequest
	GetXOssDatalakeJobId() *string
}

type GetDataLakeCachePrefetchJobRequest struct {
	// This parameter is required.
	XOssDatalakeJobId *string `json:"x-oss-datalake-job-id,omitempty" xml:"x-oss-datalake-job-id,omitempty"`
}

func (s GetDataLakeCachePrefetchJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataLakeCachePrefetchJobRequest) GoString() string {
	return s.String()
}

func (s *GetDataLakeCachePrefetchJobRequest) GetXOssDatalakeJobId() *string {
	return s.XOssDatalakeJobId
}

func (s *GetDataLakeCachePrefetchJobRequest) SetXOssDatalakeJobId(v string) *GetDataLakeCachePrefetchJobRequest {
	s.XOssDatalakeJobId = &v
	return s
}

func (s *GetDataLakeCachePrefetchJobRequest) Validate() error {
	return dara.Validate(s)
}
