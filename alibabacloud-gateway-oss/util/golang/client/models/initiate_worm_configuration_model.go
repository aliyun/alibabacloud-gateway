// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitiateWormConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetRetentionPeriodInDays(v int32) *InitiateWormConfiguration
	GetRetentionPeriodInDays() *int32
}

type InitiateWormConfiguration struct {
	// This parameter is required.
	RetentionPeriodInDays *int32 `json:"RetentionPeriodInDays,omitempty" xml:"RetentionPeriodInDays,omitempty"`
}

func (s InitiateWormConfiguration) String() string {
	return dara.Prettify(s)
}

func (s InitiateWormConfiguration) GoString() string {
	return s.String()
}

func (s *InitiateWormConfiguration) GetRetentionPeriodInDays() *int32 {
	return s.RetentionPeriodInDays
}

func (s *InitiateWormConfiguration) SetRetentionPeriodInDays(v int32) *InitiateWormConfiguration {
	s.RetentionPeriodInDays = &v
	return s
}

func (s *InitiateWormConfiguration) Validate() error {
	return dara.Validate(s)
}
