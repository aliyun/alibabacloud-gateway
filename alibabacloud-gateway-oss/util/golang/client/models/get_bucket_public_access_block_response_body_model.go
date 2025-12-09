// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketPublicAccessBlockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPublicAccessBlockConfiguration(v *PublicAccessBlockConfiguration) *GetBucketPublicAccessBlockResponseBody
	GetPublicAccessBlockConfiguration() *PublicAccessBlockConfiguration
}

type GetBucketPublicAccessBlockResponseBody struct {
	// The container in which the Block Public Access configurations are stored.
	PublicAccessBlockConfiguration *PublicAccessBlockConfiguration `json:"PublicAccessBlockConfiguration,omitempty" xml:"PublicAccessBlockConfiguration,omitempty"`
}

func (s GetBucketPublicAccessBlockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketPublicAccessBlockResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketPublicAccessBlockResponseBody) GetPublicAccessBlockConfiguration() *PublicAccessBlockConfiguration {
	return s.PublicAccessBlockConfiguration
}

func (s *GetBucketPublicAccessBlockResponseBody) SetPublicAccessBlockConfiguration(v *PublicAccessBlockConfiguration) *GetBucketPublicAccessBlockResponseBody {
	s.PublicAccessBlockConfiguration = v
	return s
}

func (s *GetBucketPublicAccessBlockResponseBody) Validate() error {
	if s.PublicAccessBlockConfiguration != nil {
		if err := s.PublicAccessBlockConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
