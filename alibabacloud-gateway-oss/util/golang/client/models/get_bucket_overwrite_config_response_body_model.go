// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketOverwriteConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOverwriteConfiguration(v *OverwriteConfiguration) *GetBucketOverwriteConfigResponseBody
	GetOverwriteConfiguration() *OverwriteConfiguration
}

type GetBucketOverwriteConfigResponseBody struct {
	OverwriteConfiguration *OverwriteConfiguration `json:"OverwriteConfiguration,omitempty" xml:"OverwriteConfiguration,omitempty"`
}

func (s GetBucketOverwriteConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketOverwriteConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketOverwriteConfigResponseBody) GetOverwriteConfiguration() *OverwriteConfiguration {
	return s.OverwriteConfiguration
}

func (s *GetBucketOverwriteConfigResponseBody) SetOverwriteConfiguration(v *OverwriteConfiguration) *GetBucketOverwriteConfigResponseBody {
	s.OverwriteConfiguration = v
	return s
}

func (s *GetBucketOverwriteConfigResponseBody) Validate() error {
	if s.OverwriteConfiguration != nil {
		if err := s.OverwriteConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
