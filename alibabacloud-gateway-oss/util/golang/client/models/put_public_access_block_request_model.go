// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutPublicAccessBlockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPublicAccessBlockConfiguration(v *PublicAccessBlockConfiguration) *PutPublicAccessBlockRequest
	GetPublicAccessBlockConfiguration() *PublicAccessBlockConfiguration
}

type PutPublicAccessBlockRequest struct {
	// The container in which the Block Public Access configurations are stored.
	PublicAccessBlockConfiguration *PublicAccessBlockConfiguration `json:"PublicAccessBlockConfiguration,omitempty" xml:"PublicAccessBlockConfiguration,omitempty"`
}

func (s PutPublicAccessBlockRequest) String() string {
	return dara.Prettify(s)
}

func (s PutPublicAccessBlockRequest) GoString() string {
	return s.String()
}

func (s *PutPublicAccessBlockRequest) GetPublicAccessBlockConfiguration() *PublicAccessBlockConfiguration {
	return s.PublicAccessBlockConfiguration
}

func (s *PutPublicAccessBlockRequest) SetPublicAccessBlockConfiguration(v *PublicAccessBlockConfiguration) *PutPublicAccessBlockRequest {
	s.PublicAccessBlockConfiguration = v
	return s
}

func (s *PutPublicAccessBlockRequest) Validate() error {
	if s.PublicAccessBlockConfiguration != nil {
		if err := s.PublicAccessBlockConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
