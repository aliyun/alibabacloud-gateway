// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPublicAccessBlockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPublicAccessBlockConfiguration(v *PublicAccessBlockConfiguration) *GetPublicAccessBlockResponseBody
	GetPublicAccessBlockConfiguration() *PublicAccessBlockConfiguration
}

type GetPublicAccessBlockResponseBody struct {
	// The container in which the Block Public Access configurations are stored.
	PublicAccessBlockConfiguration *PublicAccessBlockConfiguration `json:"PublicAccessBlockConfiguration,omitempty" xml:"PublicAccessBlockConfiguration,omitempty"`
}

func (s GetPublicAccessBlockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPublicAccessBlockResponseBody) GoString() string {
	return s.String()
}

func (s *GetPublicAccessBlockResponseBody) GetPublicAccessBlockConfiguration() *PublicAccessBlockConfiguration {
	return s.PublicAccessBlockConfiguration
}

func (s *GetPublicAccessBlockResponseBody) SetPublicAccessBlockConfiguration(v *PublicAccessBlockConfiguration) *GetPublicAccessBlockResponseBody {
	s.PublicAccessBlockConfiguration = v
	return s
}

func (s *GetPublicAccessBlockResponseBody) Validate() error {
	if s.PublicAccessBlockConfiguration != nil {
		if err := s.PublicAccessBlockConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
