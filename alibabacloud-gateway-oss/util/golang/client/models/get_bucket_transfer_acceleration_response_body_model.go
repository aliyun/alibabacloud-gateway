// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketTransferAccelerationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetTransferAccelerationConfiguration(v *GetBucketTransferAccelerationResponseBodyTransferAccelerationConfiguration) *GetBucketTransferAccelerationResponseBody
	GetTransferAccelerationConfiguration() *GetBucketTransferAccelerationResponseBodyTransferAccelerationConfiguration
}

type GetBucketTransferAccelerationResponseBody struct {
	// The container that stores the transfer acceleration configurations.
	TransferAccelerationConfiguration *GetBucketTransferAccelerationResponseBodyTransferAccelerationConfiguration `json:"TransferAccelerationConfiguration,omitempty" xml:"TransferAccelerationConfiguration,omitempty" type:"Struct"`
}

func (s GetBucketTransferAccelerationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketTransferAccelerationResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketTransferAccelerationResponseBody) GetTransferAccelerationConfiguration() *GetBucketTransferAccelerationResponseBodyTransferAccelerationConfiguration {
	return s.TransferAccelerationConfiguration
}

func (s *GetBucketTransferAccelerationResponseBody) SetTransferAccelerationConfiguration(v *GetBucketTransferAccelerationResponseBodyTransferAccelerationConfiguration) *GetBucketTransferAccelerationResponseBody {
	s.TransferAccelerationConfiguration = v
	return s
}

func (s *GetBucketTransferAccelerationResponseBody) Validate() error {
	if s.TransferAccelerationConfiguration != nil {
		if err := s.TransferAccelerationConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBucketTransferAccelerationResponseBodyTransferAccelerationConfiguration struct {
	// Whether the transfer acceleration is enabled for this bucket.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s GetBucketTransferAccelerationResponseBodyTransferAccelerationConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetBucketTransferAccelerationResponseBodyTransferAccelerationConfiguration) GoString() string {
	return s.String()
}

func (s *GetBucketTransferAccelerationResponseBodyTransferAccelerationConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetBucketTransferAccelerationResponseBodyTransferAccelerationConfiguration) SetEnabled(v bool) *GetBucketTransferAccelerationResponseBodyTransferAccelerationConfiguration {
	s.Enabled = &v
	return s
}

func (s *GetBucketTransferAccelerationResponseBodyTransferAccelerationConfiguration) Validate() error {
	return dara.Validate(s)
}
