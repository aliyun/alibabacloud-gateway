// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketRefererRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRefererConfiguration(v *RefererConfiguration) *PutBucketRefererRequest
	GetRefererConfiguration() *RefererConfiguration
}

type PutBucketRefererRequest struct {
	// The container that stores the hotlink protection configurations.
	RefererConfiguration *RefererConfiguration `json:"RefererConfiguration,omitempty" xml:"RefererConfiguration,omitempty"`
}

func (s PutBucketRefererRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketRefererRequest) GoString() string {
	return s.String()
}

func (s *PutBucketRefererRequest) GetRefererConfiguration() *RefererConfiguration {
	return s.RefererConfiguration
}

func (s *PutBucketRefererRequest) SetRefererConfiguration(v *RefererConfiguration) *PutBucketRefererRequest {
	s.RefererConfiguration = v
	return s
}

func (s *PutBucketRefererRequest) Validate() error {
	if s.RefererConfiguration != nil {
		if err := s.RefererConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
