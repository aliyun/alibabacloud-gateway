// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutCnameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketCnameConfiguration(v *BucketCnameConfiguration) *PutCnameRequest
	GetBucketCnameConfiguration() *BucketCnameConfiguration
}

type PutCnameRequest struct {
	// The container that stores the CNAME record.
	BucketCnameConfiguration *BucketCnameConfiguration `json:"BucketCnameConfiguration,omitempty" xml:"BucketCnameConfiguration,omitempty"`
}

func (s PutCnameRequest) String() string {
	return dara.Prettify(s)
}

func (s PutCnameRequest) GoString() string {
	return s.String()
}

func (s *PutCnameRequest) GetBucketCnameConfiguration() *BucketCnameConfiguration {
	return s.BucketCnameConfiguration
}

func (s *PutCnameRequest) SetBucketCnameConfiguration(v *BucketCnameConfiguration) *PutCnameRequest {
	s.BucketCnameConfiguration = v
	return s
}

func (s *PutCnameRequest) Validate() error {
	if s.BucketCnameConfiguration != nil {
		if err := s.BucketCnameConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
