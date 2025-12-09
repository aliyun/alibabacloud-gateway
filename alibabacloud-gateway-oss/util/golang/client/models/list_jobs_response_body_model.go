// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListJobsResult(v *BatchOperationListJobsResult) *ListJobsResponseBody
	GetListJobsResult() *BatchOperationListJobsResult
}

type ListJobsResponseBody struct {
	ListJobsResult *BatchOperationListJobsResult `json:"ListJobsResult,omitempty" xml:"ListJobsResult,omitempty"`
}

func (s ListJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBody) GetListJobsResult() *BatchOperationListJobsResult {
	return s.ListJobsResult
}

func (s *ListJobsResponseBody) SetListJobsResult(v *BatchOperationListJobsResult) *ListJobsResponseBody {
	s.ListJobsResult = v
	return s
}

func (s *ListJobsResponseBody) Validate() error {
	if s.ListJobsResult != nil {
		if err := s.ListJobsResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}
