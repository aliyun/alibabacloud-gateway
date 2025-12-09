// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketDataAcceleratorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataAccelerator(v *DataAccelerator) *GetBucketDataAcceleratorResponseBody
	GetDataAccelerator() *DataAccelerator
}

type GetBucketDataAcceleratorResponseBody struct {
	DataAccelerator *DataAccelerator `json:"DataAccelerator,omitempty" xml:"DataAccelerator,omitempty"`
}

func (s GetBucketDataAcceleratorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketDataAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketDataAcceleratorResponseBody) GetDataAccelerator() *DataAccelerator {
	return s.DataAccelerator
}

func (s *GetBucketDataAcceleratorResponseBody) SetDataAccelerator(v *DataAccelerator) *GetBucketDataAcceleratorResponseBody {
	s.DataAccelerator = v
	return s
}

func (s *GetBucketDataAcceleratorResponseBody) Validate() error {
	if s.DataAccelerator != nil {
		if err := s.DataAccelerator.Validate(); err != nil {
			return err
		}
	}
	return nil
}
