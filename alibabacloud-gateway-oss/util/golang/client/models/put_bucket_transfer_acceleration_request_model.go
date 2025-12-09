// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketTransferAccelerationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTransferAccelerationConfiguration(v *TransferAccelerationConfiguration) *PutBucketTransferAccelerationRequest
	GetTransferAccelerationConfiguration() *TransferAccelerationConfiguration
}

type PutBucketTransferAccelerationRequest struct {
	// The container that stores the transfer acceleration configurations.
	TransferAccelerationConfiguration *TransferAccelerationConfiguration `json:"TransferAccelerationConfiguration,omitempty" xml:"TransferAccelerationConfiguration,omitempty"`
}

func (s PutBucketTransferAccelerationRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketTransferAccelerationRequest) GoString() string {
	return s.String()
}

func (s *PutBucketTransferAccelerationRequest) GetTransferAccelerationConfiguration() *TransferAccelerationConfiguration {
	return s.TransferAccelerationConfiguration
}

func (s *PutBucketTransferAccelerationRequest) SetTransferAccelerationConfiguration(v *TransferAccelerationConfiguration) *PutBucketTransferAccelerationRequest {
	s.TransferAccelerationConfiguration = v
	return s
}

func (s *PutBucketTransferAccelerationRequest) Validate() error {
	if s.TransferAccelerationConfiguration != nil {
		if err := s.TransferAccelerationConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
