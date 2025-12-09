// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketArchiveDirectReadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArchiveDirectReadConfiguration(v *ArchiveDirectReadConfiguration) *GetBucketArchiveDirectReadResponseBody
	GetArchiveDirectReadConfiguration() *ArchiveDirectReadConfiguration
}

type GetBucketArchiveDirectReadResponseBody struct {
	// The container that stores the configurations for real-time access of Archive objects.
	ArchiveDirectReadConfiguration *ArchiveDirectReadConfiguration `json:"ArchiveDirectReadConfiguration,omitempty" xml:"ArchiveDirectReadConfiguration,omitempty"`
}

func (s GetBucketArchiveDirectReadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketArchiveDirectReadResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketArchiveDirectReadResponseBody) GetArchiveDirectReadConfiguration() *ArchiveDirectReadConfiguration {
	return s.ArchiveDirectReadConfiguration
}

func (s *GetBucketArchiveDirectReadResponseBody) SetArchiveDirectReadConfiguration(v *ArchiveDirectReadConfiguration) *GetBucketArchiveDirectReadResponseBody {
	s.ArchiveDirectReadConfiguration = v
	return s
}

func (s *GetBucketArchiveDirectReadResponseBody) Validate() error {
	if s.ArchiveDirectReadConfiguration != nil {
		if err := s.ArchiveDirectReadConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
