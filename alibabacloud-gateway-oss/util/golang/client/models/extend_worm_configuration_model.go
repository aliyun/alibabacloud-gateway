// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExtendWormConfiguration interface {
  dara.Model
  String() string
  GoString() string
  SetRetentionPeriodInDays(v int32) *ExtendWormConfiguration
  GetRetentionPeriodInDays() *int32 
}

type ExtendWormConfiguration struct {
  RetentionPeriodInDays *int32 `json:"RetentionPeriodInDays,omitempty" xml:"RetentionPeriodInDays,omitempty"`
}

func (s ExtendWormConfiguration) String() string {
  return dara.Prettify(s)
}

func (s ExtendWormConfiguration) GoString() string {
  return s.String()
}

func (s *ExtendWormConfiguration) GetRetentionPeriodInDays() *int32  {
  return s.RetentionPeriodInDays
}

func (s *ExtendWormConfiguration) SetRetentionPeriodInDays(v int32) *ExtendWormConfiguration {
  s.RetentionPeriodInDays = &v
  return s
}

func (s *ExtendWormConfiguration) Validate() error {
  return dara.Validate(s)
}

