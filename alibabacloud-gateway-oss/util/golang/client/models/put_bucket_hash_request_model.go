// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketHashRequest interface {
	dara.Model
	String() string
	GoString() string
	SetObjectHashConfiguration(v *ObjectHashConfiguration) *PutBucketHashRequest
	GetObjectHashConfiguration() *ObjectHashConfiguration
}

type PutBucketHashRequest struct {
	// Object Hash Algorithm Configuration
	ObjectHashConfiguration *ObjectHashConfiguration `json:"ObjectHashConfiguration,omitempty" xml:"ObjectHashConfiguration,omitempty"`
}

func (s PutBucketHashRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketHashRequest) GoString() string {
	return s.String()
}

func (s *PutBucketHashRequest) GetObjectHashConfiguration() *ObjectHashConfiguration {
	return s.ObjectHashConfiguration
}

func (s *PutBucketHashRequest) SetObjectHashConfiguration(v *ObjectHashConfiguration) *PutBucketHashRequest {
	s.ObjectHashConfiguration = v
	return s
}

func (s *PutBucketHashRequest) Validate() error {
	if s.ObjectHashConfiguration != nil {
		if err := s.ObjectHashConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
