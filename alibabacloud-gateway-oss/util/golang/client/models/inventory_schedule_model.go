// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInventorySchedule interface {
	dara.Model
	String() string
	GoString() string
	SetFrequency(v string) *InventorySchedule
	GetFrequency() *string
}

type InventorySchedule struct {
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
}

func (s InventorySchedule) String() string {
	return dara.Prettify(s)
}

func (s InventorySchedule) GoString() string {
	return s.String()
}

func (s *InventorySchedule) GetFrequency() *string {
	return s.Frequency
}

func (s *InventorySchedule) SetFrequency(v string) *InventorySchedule {
	s.Frequency = &v
	return s
}

func (s *InventorySchedule) Validate() error {
	return dara.Validate(s)
}
