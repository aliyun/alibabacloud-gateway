// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProcessConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBucketProcessConfiguration(v *GetBucketProcessConfiguration) *GetProcessConfigurationResponseBody
	GetBucketProcessConfiguration() *GetBucketProcessConfiguration
}

type GetProcessConfigurationResponseBody struct {
	BucketProcessConfiguration *GetBucketProcessConfiguration `json:"BucketProcessConfiguration,omitempty" xml:"BucketProcessConfiguration,omitempty"`
}

func (s GetProcessConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProcessConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetProcessConfigurationResponseBody) GetBucketProcessConfiguration() *GetBucketProcessConfiguration {
	return s.BucketProcessConfiguration
}

func (s *GetProcessConfigurationResponseBody) SetBucketProcessConfiguration(v *GetBucketProcessConfiguration) *GetProcessConfigurationResponseBody {
	s.BucketProcessConfiguration = v
	return s
}

func (s *GetProcessConfigurationResponseBody) Validate() error {
	if s.BucketProcessConfiguration != nil {
		if err := s.BucketProcessConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
