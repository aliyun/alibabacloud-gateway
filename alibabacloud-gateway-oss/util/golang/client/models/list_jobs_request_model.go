// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchJobStatuses(v string) *ListJobsRequest
	GetBatchJobStatuses() *string
	SetContinuationToken(v string) *ListJobsRequest
	GetContinuationToken() *string
	SetMaxKeys(v int32) *ListJobsRequest
	GetMaxKeys() *int32
}

type ListJobsRequest struct {
	BatchJobStatuses  *string `json:"batchJobStatuses,omitempty" xml:"batchJobStatuses,omitempty"`
	ContinuationToken *string `json:"continuation-token,omitempty" xml:"continuation-token,omitempty"`
	MaxKeys           *int32  `json:"max-keys,omitempty" xml:"max-keys,omitempty"`
}

func (s ListJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobsRequest) GoString() string {
	return s.String()
}

func (s *ListJobsRequest) GetBatchJobStatuses() *string {
	return s.BatchJobStatuses
}

func (s *ListJobsRequest) GetContinuationToken() *string {
	return s.ContinuationToken
}

func (s *ListJobsRequest) GetMaxKeys() *int32 {
	return s.MaxKeys
}

func (s *ListJobsRequest) SetBatchJobStatuses(v string) *ListJobsRequest {
	s.BatchJobStatuses = &v
	return s
}

func (s *ListJobsRequest) SetContinuationToken(v string) *ListJobsRequest {
	s.ContinuationToken = &v
	return s
}

func (s *ListJobsRequest) SetMaxKeys(v int32) *ListJobsRequest {
	s.MaxKeys = &v
	return s
}

func (s *ListJobsRequest) Validate() error {
	return dara.Validate(s)
}
