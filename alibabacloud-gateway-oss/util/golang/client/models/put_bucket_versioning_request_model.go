// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketVersioningRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVersioningConfiguration(v *VersioningConfiguration) *PutBucketVersioningRequest
	GetVersioningConfiguration() *VersioningConfiguration
}

type PutBucketVersioningRequest struct {
	// The container that stores the versioning state of the bucket.
	VersioningConfiguration *VersioningConfiguration `json:"VersioningConfiguration,omitempty" xml:"VersioningConfiguration,omitempty"`
}

func (s PutBucketVersioningRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketVersioningRequest) GoString() string {
	return s.String()
}

func (s *PutBucketVersioningRequest) GetVersioningConfiguration() *VersioningConfiguration {
	return s.VersioningConfiguration
}

func (s *PutBucketVersioningRequest) SetVersioningConfiguration(v *VersioningConfiguration) *PutBucketVersioningRequest {
	s.VersioningConfiguration = v
	return s
}

func (s *PutBucketVersioningRequest) Validate() error {
	if s.VersioningConfiguration != nil {
		if err := s.VersioningConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
