// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateJobResult(v *CreateJobResult) *CreateJobResponseBody
	GetCreateJobResult() *CreateJobResult
}

type CreateJobResponseBody struct {
	CreateJobResult *CreateJobResult `json:"CreateJobResult,omitempty" xml:"CreateJobResult,omitempty"`
}

func (s CreateJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobResponseBody) GetCreateJobResult() *CreateJobResult {
	return s.CreateJobResult
}

func (s *CreateJobResponseBody) SetCreateJobResult(v *CreateJobResult) *CreateJobResponseBody {
	s.CreateJobResult = v
	return s
}

func (s *CreateJobResponseBody) Validate() error {
	if s.CreateJobResult != nil {
		if err := s.CreateJobResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}
