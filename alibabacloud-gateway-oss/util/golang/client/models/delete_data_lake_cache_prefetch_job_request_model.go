// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataLakeCachePrefetchJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXOssDatalakeJobId(v string) *DeleteDataLakeCachePrefetchJobRequest
	GetXOssDatalakeJobId() *string
}

type DeleteDataLakeCachePrefetchJobRequest struct {
	// This parameter is required.
	XOssDatalakeJobId *string `json:"x-oss-datalake-job-id,omitempty" xml:"x-oss-datalake-job-id,omitempty"`
}

func (s DeleteDataLakeCachePrefetchJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataLakeCachePrefetchJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataLakeCachePrefetchJobRequest) GetXOssDatalakeJobId() *string {
	return s.XOssDatalakeJobId
}

func (s *DeleteDataLakeCachePrefetchJobRequest) SetXOssDatalakeJobId(v string) *DeleteDataLakeCachePrefetchJobRequest {
	s.XOssDatalakeJobId = &v
	return s
}

func (s *DeleteDataLakeCachePrefetchJobRequest) Validate() error {
	return dara.Validate(s)
}
