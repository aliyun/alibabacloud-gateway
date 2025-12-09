// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateJobRequest(v *BatchCreateJobRequest) *CreateJobRequest
	GetCreateJobRequest() *BatchCreateJobRequest
}

type CreateJobRequest struct {
	CreateJobRequest *BatchCreateJobRequest `json:"CreateJobRequest,omitempty" xml:"CreateJobRequest,omitempty"`
}

func (s CreateJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateJobRequest) GoString() string {
	return s.String()
}

func (s *CreateJobRequest) GetCreateJobRequest() *BatchCreateJobRequest {
	return s.CreateJobRequest
}

func (s *CreateJobRequest) SetCreateJobRequest(v *BatchCreateJobRequest) *CreateJobRequest {
	s.CreateJobRequest = v
	return s
}

func (s *CreateJobRequest) Validate() error {
	if s.CreateJobRequest != nil {
		if err := s.CreateJobRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}
