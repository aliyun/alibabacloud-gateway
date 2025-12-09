// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketHttpsConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHttpsConfiguration(v *HttpsConfiguration) *PutBucketHttpsConfigRequest
	GetHttpsConfiguration() *HttpsConfiguration
}

type PutBucketHttpsConfigRequest struct {
	// The container that stores HTTPS configurations.
	HttpsConfiguration *HttpsConfiguration `json:"HttpsConfiguration,omitempty" xml:"HttpsConfiguration,omitempty"`
}

func (s PutBucketHttpsConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketHttpsConfigRequest) GoString() string {
	return s.String()
}

func (s *PutBucketHttpsConfigRequest) GetHttpsConfiguration() *HttpsConfiguration {
	return s.HttpsConfiguration
}

func (s *PutBucketHttpsConfigRequest) SetHttpsConfiguration(v *HttpsConfiguration) *PutBucketHttpsConfigRequest {
	s.HttpsConfiguration = v
	return s
}

func (s *PutBucketHttpsConfigRequest) Validate() error {
	if s.HttpsConfiguration != nil {
		if err := s.HttpsConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
