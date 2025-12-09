// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketArchiveDirectReadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArchiveDirectReadConfiguration(v *ArchiveDirectReadConfiguration) *PutBucketArchiveDirectReadRequest
	GetArchiveDirectReadConfiguration() *ArchiveDirectReadConfiguration
}

type PutBucketArchiveDirectReadRequest struct {
	// The container that stores the configurations for real-time access of Archive objects.
	ArchiveDirectReadConfiguration *ArchiveDirectReadConfiguration `json:"ArchiveDirectReadConfiguration,omitempty" xml:"ArchiveDirectReadConfiguration,omitempty"`
}

func (s PutBucketArchiveDirectReadRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketArchiveDirectReadRequest) GoString() string {
	return s.String()
}

func (s *PutBucketArchiveDirectReadRequest) GetArchiveDirectReadConfiguration() *ArchiveDirectReadConfiguration {
	return s.ArchiveDirectReadConfiguration
}

func (s *PutBucketArchiveDirectReadRequest) SetArchiveDirectReadConfiguration(v *ArchiveDirectReadConfiguration) *PutBucketArchiveDirectReadRequest {
	s.ArchiveDirectReadConfiguration = v
	return s
}

func (s *PutBucketArchiveDirectReadRequest) Validate() error {
	if s.ArchiveDirectReadConfiguration != nil {
		if err := s.ArchiveDirectReadConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
