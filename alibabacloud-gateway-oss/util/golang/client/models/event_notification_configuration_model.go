// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEventNotificationConfiguration interface {
  dara.Model
  String() string
  GoString() string
  SetFunctionComputeConfiguration(v []*FunctionComputeConfiguration) *EventNotificationConfiguration
  GetFunctionComputeConfiguration() []*FunctionComputeConfiguration 
}

type EventNotificationConfiguration struct {
  FunctionComputeConfiguration []*FunctionComputeConfiguration `json:"FunctionComputeConfiguration,omitempty" xml:"FunctionComputeConfiguration,omitempty" type:"Repeated"`
}

func (s EventNotificationConfiguration) String() string {
  return dara.Prettify(s)
}

func (s EventNotificationConfiguration) GoString() string {
  return s.String()
}

func (s *EventNotificationConfiguration) GetFunctionComputeConfiguration() []*FunctionComputeConfiguration  {
  return s.FunctionComputeConfiguration
}

func (s *EventNotificationConfiguration) SetFunctionComputeConfiguration(v []*FunctionComputeConfiguration) *EventNotificationConfiguration {
  s.FunctionComputeConfiguration = v
  return s
}

func (s *EventNotificationConfiguration) Validate() error {
  if s.FunctionComputeConfiguration != nil {
    for _, item := range s.FunctionComputeConfiguration {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

