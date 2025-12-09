// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDays(v int64) *RestoreRequest
	GetDays() *int64
	SetJobParameters(v *RestoreRequestJobParameters) *RestoreRequest
	GetJobParameters() *RestoreRequestJobParameters
}

type RestoreRequest struct {
	Days          *int64                       `json:"Days,omitempty" xml:"Days,omitempty"`
	JobParameters *RestoreRequestJobParameters `json:"JobParameters,omitempty" xml:"JobParameters,omitempty" type:"Struct"`
}

func (s RestoreRequest) String() string {
	return dara.Prettify(s)
}

func (s RestoreRequest) GoString() string {
	return s.String()
}

func (s *RestoreRequest) GetDays() *int64 {
	return s.Days
}

func (s *RestoreRequest) GetJobParameters() *RestoreRequestJobParameters {
	return s.JobParameters
}

func (s *RestoreRequest) SetDays(v int64) *RestoreRequest {
	s.Days = &v
	return s
}

func (s *RestoreRequest) SetJobParameters(v *RestoreRequestJobParameters) *RestoreRequest {
	s.JobParameters = v
	return s
}

func (s *RestoreRequest) Validate() error {
	if s.JobParameters != nil {
		if err := s.JobParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RestoreRequestJobParameters struct {
	Tier *string `json:"Tier,omitempty" xml:"Tier,omitempty"`
}

func (s RestoreRequestJobParameters) String() string {
	return dara.Prettify(s)
}

func (s RestoreRequestJobParameters) GoString() string {
	return s.String()
}

func (s *RestoreRequestJobParameters) GetTier() *string {
	return s.Tier
}

func (s *RestoreRequestJobParameters) SetTier(v string) *RestoreRequestJobParameters {
	s.Tier = &v
	return s
}

func (s *RestoreRequestJobParameters) Validate() error {
	return dara.Validate(s)
}
