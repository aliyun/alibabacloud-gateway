// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeCachePrefetchJobHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXOssDatalakeJobId(v string) *ListDataLakeCachePrefetchJobHistoryRequest
	GetXOssDatalakeJobId() *string
}

type ListDataLakeCachePrefetchJobHistoryRequest struct {
	// This parameter is required.
	XOssDatalakeJobId *string `json:"x-oss-datalake-job-id,omitempty" xml:"x-oss-datalake-job-id,omitempty"`
}

func (s ListDataLakeCachePrefetchJobHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeCachePrefetchJobHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListDataLakeCachePrefetchJobHistoryRequest) GetXOssDatalakeJobId() *string {
	return s.XOssDatalakeJobId
}

func (s *ListDataLakeCachePrefetchJobHistoryRequest) SetXOssDatalakeJobId(v string) *ListDataLakeCachePrefetchJobHistoryRequest {
	s.XOssDatalakeJobId = &v
	return s
}

func (s *ListDataLakeCachePrefetchJobHistoryRequest) Validate() error {
	return dara.Validate(s)
}
