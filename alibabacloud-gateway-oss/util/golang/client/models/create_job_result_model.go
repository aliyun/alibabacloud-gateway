// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobResult interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *CreateJobResult
	GetJobId() *string
}

type CreateJobResult struct {
	// This parameter is required.
	//
	// example:
	//
	// IDBkODc3M2RlZjgyNjQ0NDViZGV****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s CreateJobResult) String() string {
	return dara.Prettify(s)
}

func (s CreateJobResult) GoString() string {
	return s.String()
}

func (s *CreateJobResult) GetJobId() *string {
	return s.JobId
}

func (s *CreateJobResult) SetJobId(v string) *CreateJobResult {
	s.JobId = &v
	return s
}

func (s *CreateJobResult) Validate() error {
	return dara.Validate(s)
}
