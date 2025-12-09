// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchOperationListJobsResult interface {
	dara.Model
	String() string
	GoString() string
	SetJobs(v *BatchOperationListJobs) *BatchOperationListJobsResult
	GetJobs() *BatchOperationListJobs
	SetNextToken(v string) *BatchOperationListJobsResult
	GetNextToken() *string
}

type BatchOperationListJobsResult struct {
	Jobs *BatchOperationListJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty"`
	// example:
	//
	// /
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s BatchOperationListJobsResult) String() string {
	return dara.Prettify(s)
}

func (s BatchOperationListJobsResult) GoString() string {
	return s.String()
}

func (s *BatchOperationListJobsResult) GetJobs() *BatchOperationListJobs {
	return s.Jobs
}

func (s *BatchOperationListJobsResult) GetNextToken() *string {
	return s.NextToken
}

func (s *BatchOperationListJobsResult) SetJobs(v *BatchOperationListJobs) *BatchOperationListJobsResult {
	s.Jobs = v
	return s
}

func (s *BatchOperationListJobsResult) SetNextToken(v string) *BatchOperationListJobsResult {
	s.NextToken = &v
	return s
}

func (s *BatchOperationListJobsResult) Validate() error {
	if s.Jobs != nil {
		if err := s.Jobs.Validate(); err != nil {
			return err
		}
	}
	return nil
}
