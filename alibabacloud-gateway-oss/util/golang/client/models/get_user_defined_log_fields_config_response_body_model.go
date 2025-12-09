// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserDefinedLogFieldsConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetUserDefinedLogFieldsConfiguration(v *UserDefinedLogFieldsConfiguration) *GetUserDefinedLogFieldsConfigResponseBody
	GetUserDefinedLogFieldsConfiguration() *UserDefinedLogFieldsConfiguration
}

type GetUserDefinedLogFieldsConfigResponseBody struct {
	// The container for the user-defined logging configuration.
	UserDefinedLogFieldsConfiguration *UserDefinedLogFieldsConfiguration `json:"UserDefinedLogFieldsConfiguration,omitempty" xml:"UserDefinedLogFieldsConfiguration,omitempty"`
}

func (s GetUserDefinedLogFieldsConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserDefinedLogFieldsConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserDefinedLogFieldsConfigResponseBody) GetUserDefinedLogFieldsConfiguration() *UserDefinedLogFieldsConfiguration {
	return s.UserDefinedLogFieldsConfiguration
}

func (s *GetUserDefinedLogFieldsConfigResponseBody) SetUserDefinedLogFieldsConfiguration(v *UserDefinedLogFieldsConfiguration) *GetUserDefinedLogFieldsConfigResponseBody {
	s.UserDefinedLogFieldsConfiguration = v
	return s
}

func (s *GetUserDefinedLogFieldsConfigResponseBody) Validate() error {
	if s.UserDefinedLogFieldsConfiguration != nil {
		if err := s.UserDefinedLogFieldsConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
