// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketOverwriteConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOverwriteConfiguration(v *OverwriteConfiguration) *PutBucketOverwriteConfigRequest
	GetOverwriteConfiguration() *OverwriteConfiguration
}

type PutBucketOverwriteConfigRequest struct {
	OverwriteConfiguration *OverwriteConfiguration `json:"OverwriteConfiguration,omitempty" xml:"OverwriteConfiguration,omitempty"`
}

func (s PutBucketOverwriteConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketOverwriteConfigRequest) GoString() string {
	return s.String()
}

func (s *PutBucketOverwriteConfigRequest) GetOverwriteConfiguration() *OverwriteConfiguration {
	return s.OverwriteConfiguration
}

func (s *PutBucketOverwriteConfigRequest) SetOverwriteConfiguration(v *OverwriteConfiguration) *PutBucketOverwriteConfigRequest {
	s.OverwriteConfiguration = v
	return s
}

func (s *PutBucketOverwriteConfigRequest) Validate() error {
	if s.OverwriteConfiguration != nil {
		if err := s.OverwriteConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
