// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketRefererResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRefererConfiguration(v *RefererConfiguration) *GetBucketRefererResponseBody
	GetRefererConfiguration() *RefererConfiguration
}

type GetBucketRefererResponseBody struct {
	// The container that stores the hotlink protection configurations.
	RefererConfiguration *RefererConfiguration `json:"RefererConfiguration,omitempty" xml:"RefererConfiguration,omitempty"`
}

func (s GetBucketRefererResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketRefererResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketRefererResponseBody) GetRefererConfiguration() *RefererConfiguration {
	return s.RefererConfiguration
}

func (s *GetBucketRefererResponseBody) SetRefererConfiguration(v *RefererConfiguration) *GetBucketRefererResponseBody {
	s.RefererConfiguration = v
	return s
}

func (s *GetBucketRefererResponseBody) Validate() error {
	if s.RefererConfiguration != nil {
		if err := s.RefererConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
