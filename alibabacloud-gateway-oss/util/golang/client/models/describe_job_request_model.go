// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchJobId(v string) *DescribeJobRequest
	GetBatchJobId() *string
}

type DescribeJobRequest struct {
	BatchJobId *string `json:"batchJobId,omitempty" xml:"batchJobId,omitempty"`
}

func (s DescribeJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobRequest) GoString() string {
	return s.String()
}

func (s *DescribeJobRequest) GetBatchJobId() *string {
	return s.BatchJobId
}

func (s *DescribeJobRequest) SetBatchJobId(v string) *DescribeJobRequest {
	s.BatchJobId = &v
	return s
}

func (s *DescribeJobRequest) Validate() error {
	return dara.Validate(s)
}
