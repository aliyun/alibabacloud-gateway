// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferAccelerationConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *TransferAccelerationConfiguration
	GetEnabled() *bool
}

type TransferAccelerationConfiguration struct {
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s TransferAccelerationConfiguration) String() string {
	return dara.Prettify(s)
}

func (s TransferAccelerationConfiguration) GoString() string {
	return s.String()
}

func (s *TransferAccelerationConfiguration) GetEnabled() *bool {
	return s.Enabled
}

func (s *TransferAccelerationConfiguration) SetEnabled(v bool) *TransferAccelerationConfiguration {
	s.Enabled = &v
	return s
}

func (s *TransferAccelerationConfiguration) Validate() error {
	return dara.Validate(s)
}
