// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessPointPublicAccessBlockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPublicAccessBlockConfiguration(v *PublicAccessBlockConfiguration) *GetAccessPointPublicAccessBlockResponseBody
	GetPublicAccessBlockConfiguration() *PublicAccessBlockConfiguration
}

type GetAccessPointPublicAccessBlockResponseBody struct {
	// The container in which the Block Public Access configurations are stored.
	PublicAccessBlockConfiguration *PublicAccessBlockConfiguration `json:"PublicAccessBlockConfiguration,omitempty" xml:"PublicAccessBlockConfiguration,omitempty"`
}

func (s GetAccessPointPublicAccessBlockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccessPointPublicAccessBlockResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessPointPublicAccessBlockResponseBody) GetPublicAccessBlockConfiguration() *PublicAccessBlockConfiguration {
	return s.PublicAccessBlockConfiguration
}

func (s *GetAccessPointPublicAccessBlockResponseBody) SetPublicAccessBlockConfiguration(v *PublicAccessBlockConfiguration) *GetAccessPointPublicAccessBlockResponseBody {
	s.PublicAccessBlockConfiguration = v
	return s
}

func (s *GetAccessPointPublicAccessBlockResponseBody) Validate() error {
	if s.PublicAccessBlockConfiguration != nil {
		if err := s.PublicAccessBlockConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
