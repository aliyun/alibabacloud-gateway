// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutUserDefinedLogFieldsConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserDefinedLogFieldsConfiguration(v *UserDefinedLogFieldsConfiguration) *PutUserDefinedLogFieldsConfigRequest
	GetUserDefinedLogFieldsConfiguration() *UserDefinedLogFieldsConfiguration
}

type PutUserDefinedLogFieldsConfigRequest struct {
	// The container for the user-defined logging configuration.
	UserDefinedLogFieldsConfiguration *UserDefinedLogFieldsConfiguration `json:"UserDefinedLogFieldsConfiguration,omitempty" xml:"UserDefinedLogFieldsConfiguration,omitempty"`
}

func (s PutUserDefinedLogFieldsConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s PutUserDefinedLogFieldsConfigRequest) GoString() string {
	return s.String()
}

func (s *PutUserDefinedLogFieldsConfigRequest) GetUserDefinedLogFieldsConfiguration() *UserDefinedLogFieldsConfiguration {
	return s.UserDefinedLogFieldsConfiguration
}

func (s *PutUserDefinedLogFieldsConfigRequest) SetUserDefinedLogFieldsConfiguration(v *UserDefinedLogFieldsConfiguration) *PutUserDefinedLogFieldsConfigRequest {
	s.UserDefinedLogFieldsConfiguration = v
	return s
}

func (s *PutUserDefinedLogFieldsConfigRequest) Validate() error {
	if s.UserDefinedLogFieldsConfiguration != nil {
		if err := s.UserDefinedLogFieldsConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
