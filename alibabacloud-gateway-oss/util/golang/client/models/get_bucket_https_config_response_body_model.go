// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketHttpsConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHttpsConfiguration(v *HttpsConfiguration) *GetBucketHttpsConfigResponseBody
	GetHttpsConfiguration() *HttpsConfiguration
}

type GetBucketHttpsConfigResponseBody struct {
	// The container that stores HTTPS configurations.
	HttpsConfiguration *HttpsConfiguration `json:"HttpsConfiguration,omitempty" xml:"HttpsConfiguration,omitempty"`
}

func (s GetBucketHttpsConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketHttpsConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketHttpsConfigResponseBody) GetHttpsConfiguration() *HttpsConfiguration {
	return s.HttpsConfiguration
}

func (s *GetBucketHttpsConfigResponseBody) SetHttpsConfiguration(v *HttpsConfiguration) *GetBucketHttpsConfigResponseBody {
	s.HttpsConfiguration = v
	return s
}

func (s *GetBucketHttpsConfigResponseBody) Validate() error {
	if s.HttpsConfiguration != nil {
		if err := s.HttpsConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
