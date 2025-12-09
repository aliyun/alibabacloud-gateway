// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketDataAcceleratorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataAcceleratorConfiguration(v *DataAcceleratorConfiguration) *PutBucketDataAcceleratorRequest
	GetDataAcceleratorConfiguration() *DataAcceleratorConfiguration
}

type PutBucketDataAcceleratorRequest struct {
	DataAcceleratorConfiguration *DataAcceleratorConfiguration `json:"DataAcceleratorConfiguration,omitempty" xml:"DataAcceleratorConfiguration,omitempty"`
}

func (s PutBucketDataAcceleratorRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketDataAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *PutBucketDataAcceleratorRequest) GetDataAcceleratorConfiguration() *DataAcceleratorConfiguration {
	return s.DataAcceleratorConfiguration
}

func (s *PutBucketDataAcceleratorRequest) SetDataAcceleratorConfiguration(v *DataAcceleratorConfiguration) *PutBucketDataAcceleratorRequest {
	s.DataAcceleratorConfiguration = v
	return s
}

func (s *PutBucketDataAcceleratorRequest) Validate() error {
	if s.DataAcceleratorConfiguration != nil {
		if err := s.DataAcceleratorConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
