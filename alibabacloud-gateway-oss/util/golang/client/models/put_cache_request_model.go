// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutCacheRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCacheConfiguration(v *CreateCacheConfiguration) *PutCacheRequest
	GetCreateCacheConfiguration() *CreateCacheConfiguration
}

type PutCacheRequest struct {
	CreateCacheConfiguration *CreateCacheConfiguration `json:"CreateCacheConfiguration,omitempty" xml:"CreateCacheConfiguration,omitempty"`
}

func (s PutCacheRequest) String() string {
	return dara.Prettify(s)
}

func (s PutCacheRequest) GoString() string {
	return s.String()
}

func (s *PutCacheRequest) GetCreateCacheConfiguration() *CreateCacheConfiguration {
	return s.CreateCacheConfiguration
}

func (s *PutCacheRequest) SetCreateCacheConfiguration(v *CreateCacheConfiguration) *PutCacheRequest {
	s.CreateCacheConfiguration = v
	return s
}

func (s *PutCacheRequest) Validate() error {
	if s.CreateCacheConfiguration != nil {
		if err := s.CreateCacheConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
