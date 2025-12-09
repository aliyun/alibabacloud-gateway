// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateBucketConfiguration(v *CreateBucketConfiguration) *PutBucketRequest
	GetCreateBucketConfiguration() *CreateBucketConfiguration
}

type PutBucketRequest struct {
	// The container that stores the information about the bucket to be created.
	CreateBucketConfiguration *CreateBucketConfiguration `json:"CreateBucketConfiguration,omitempty" xml:"CreateBucketConfiguration,omitempty"`
}

func (s PutBucketRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketRequest) GoString() string {
	return s.String()
}

func (s *PutBucketRequest) GetCreateBucketConfiguration() *CreateBucketConfiguration {
	return s.CreateBucketConfiguration
}

func (s *PutBucketRequest) SetCreateBucketConfiguration(v *CreateBucketConfiguration) *PutBucketRequest {
	s.CreateBucketConfiguration = v
	return s
}

func (s *PutBucketRequest) Validate() error {
	if s.CreateBucketConfiguration != nil {
		if err := s.CreateBucketConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
