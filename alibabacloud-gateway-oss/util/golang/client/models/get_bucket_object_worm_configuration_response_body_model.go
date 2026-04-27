// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketObjectWormConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetObjectWormConfiguration(v *ObjectWormConfiguration) *GetBucketObjectWormConfigurationResponseBody
	GetObjectWormConfiguration() *ObjectWormConfiguration
}

type GetBucketObjectWormConfigurationResponseBody struct {
	ObjectWormConfiguration *ObjectWormConfiguration `json:"ObjectWormConfiguration,omitempty" xml:"ObjectWormConfiguration,omitempty"`
}

func (s GetBucketObjectWormConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketObjectWormConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketObjectWormConfigurationResponseBody) GetObjectWormConfiguration() *ObjectWormConfiguration {
	return s.ObjectWormConfiguration
}

func (s *GetBucketObjectWormConfigurationResponseBody) SetObjectWormConfiguration(v *ObjectWormConfiguration) *GetBucketObjectWormConfigurationResponseBody {
	s.ObjectWormConfiguration = v
	return s
}

func (s *GetBucketObjectWormConfigurationResponseBody) Validate() error {
	if s.ObjectWormConfiguration != nil {
		if err := s.ObjectWormConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
