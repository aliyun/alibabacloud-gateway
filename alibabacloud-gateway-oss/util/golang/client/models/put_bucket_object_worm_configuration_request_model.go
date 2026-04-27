// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketObjectWormConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PutBucketObjectWormConfigurationRequestBody) *PutBucketObjectWormConfigurationRequest
	GetBody() *PutBucketObjectWormConfigurationRequestBody
}

type PutBucketObjectWormConfigurationRequest struct {
	Body *PutBucketObjectWormConfigurationRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s PutBucketObjectWormConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketObjectWormConfigurationRequest) GoString() string {
	return s.String()
}

func (s *PutBucketObjectWormConfigurationRequest) GetBody() *PutBucketObjectWormConfigurationRequestBody {
	return s.Body
}

func (s *PutBucketObjectWormConfigurationRequest) SetBody(v *PutBucketObjectWormConfigurationRequestBody) *PutBucketObjectWormConfigurationRequest {
	s.Body = v
	return s
}

func (s *PutBucketObjectWormConfigurationRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PutBucketObjectWormConfigurationRequestBody struct {
	ObjectWormConfiguration *ObjectWormConfiguration `json:"ObjectWormConfiguration,omitempty" xml:"ObjectWormConfiguration,omitempty"`
}

func (s PutBucketObjectWormConfigurationRequestBody) String() string {
	return dara.Prettify(s)
}

func (s PutBucketObjectWormConfigurationRequestBody) GoString() string {
	return s.String()
}

func (s *PutBucketObjectWormConfigurationRequestBody) GetObjectWormConfiguration() *ObjectWormConfiguration {
	return s.ObjectWormConfiguration
}

func (s *PutBucketObjectWormConfigurationRequestBody) SetObjectWormConfiguration(v *ObjectWormConfiguration) *PutBucketObjectWormConfigurationRequestBody {
	s.ObjectWormConfiguration = v
	return s
}

func (s *PutBucketObjectWormConfigurationRequestBody) Validate() error {
	if s.ObjectWormConfiguration != nil {
		if err := s.ObjectWormConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
