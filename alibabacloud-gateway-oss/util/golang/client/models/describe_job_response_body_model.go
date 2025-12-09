// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescribeJobResult(v *DescribeJobResult) *DescribeJobResponseBody
	GetDescribeJobResult() *DescribeJobResult
}

type DescribeJobResponseBody struct {
	DescribeJobResult *DescribeJobResult `json:"DescribeJobResult,omitempty" xml:"DescribeJobResult,omitempty"`
}

func (s DescribeJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJobResponseBody) GetDescribeJobResult() *DescribeJobResult {
	return s.DescribeJobResult
}

func (s *DescribeJobResponseBody) SetDescribeJobResult(v *DescribeJobResult) *DescribeJobResponseBody {
	s.DescribeJobResult = v
	return s
}

func (s *DescribeJobResponseBody) Validate() error {
	if s.DescribeJobResult != nil {
		if err := s.DescribeJobResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}
