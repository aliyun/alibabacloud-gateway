// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketPublicAccessBlockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPublicAccessBlockConfiguration(v *PublicAccessBlockConfiguration) *PutBucketPublicAccessBlockRequest
	GetPublicAccessBlockConfiguration() *PublicAccessBlockConfiguration
}

type PutBucketPublicAccessBlockRequest struct {
	// The container in which the Block Public Access configurations are stored.
	PublicAccessBlockConfiguration *PublicAccessBlockConfiguration `json:"PublicAccessBlockConfiguration,omitempty" xml:"PublicAccessBlockConfiguration,omitempty"`
}

func (s PutBucketPublicAccessBlockRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketPublicAccessBlockRequest) GoString() string {
	return s.String()
}

func (s *PutBucketPublicAccessBlockRequest) GetPublicAccessBlockConfiguration() *PublicAccessBlockConfiguration {
	return s.PublicAccessBlockConfiguration
}

func (s *PutBucketPublicAccessBlockRequest) SetPublicAccessBlockConfiguration(v *PublicAccessBlockConfiguration) *PutBucketPublicAccessBlockRequest {
	s.PublicAccessBlockConfiguration = v
	return s
}

func (s *PutBucketPublicAccessBlockRequest) Validate() error {
	if s.PublicAccessBlockConfiguration != nil {
		if err := s.PublicAccessBlockConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
