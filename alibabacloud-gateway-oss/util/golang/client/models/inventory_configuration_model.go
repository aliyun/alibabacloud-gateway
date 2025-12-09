// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInventoryConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetDestination(v *InventoryDestination) *InventoryConfiguration
	GetDestination() *InventoryDestination
	SetFilter(v *InventoryFilter) *InventoryConfiguration
	GetFilter() *InventoryFilter
	SetId(v string) *InventoryConfiguration
	GetId() *string
	SetIncludedObjectVersions(v string) *InventoryConfiguration
	GetIncludedObjectVersions() *string
	SetIsEnabled(v bool) *InventoryConfiguration
	GetIsEnabled() *bool
	SetOptionalFields(v *InventoryConfigurationOptionalFields) *InventoryConfiguration
	GetOptionalFields() *InventoryConfigurationOptionalFields
	SetSchedule(v *InventorySchedule) *InventoryConfiguration
	GetSchedule() *InventorySchedule
}

type InventoryConfiguration struct {
	Destination *InventoryDestination `json:"Destination,omitempty" xml:"Destination,omitempty"`
	Filter      *InventoryFilter      `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// example:
	//
	// report1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// Current
	IncludedObjectVersions *string `json:"IncludedObjectVersions,omitempty" xml:"IncludedObjectVersions,omitempty"`
	// example:
	//
	// true
	IsEnabled      *bool                                 `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	OptionalFields *InventoryConfigurationOptionalFields `json:"OptionalFields,omitempty" xml:"OptionalFields,omitempty" type:"Struct"`
	Schedule       *InventorySchedule                    `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
}

func (s InventoryConfiguration) String() string {
	return dara.Prettify(s)
}

func (s InventoryConfiguration) GoString() string {
	return s.String()
}

func (s *InventoryConfiguration) GetDestination() *InventoryDestination {
	return s.Destination
}

func (s *InventoryConfiguration) GetFilter() *InventoryFilter {
	return s.Filter
}

func (s *InventoryConfiguration) GetId() *string {
	return s.Id
}

func (s *InventoryConfiguration) GetIncludedObjectVersions() *string {
	return s.IncludedObjectVersions
}

func (s *InventoryConfiguration) GetIsEnabled() *bool {
	return s.IsEnabled
}

func (s *InventoryConfiguration) GetOptionalFields() *InventoryConfigurationOptionalFields {
	return s.OptionalFields
}

func (s *InventoryConfiguration) GetSchedule() *InventorySchedule {
	return s.Schedule
}

func (s *InventoryConfiguration) SetDestination(v *InventoryDestination) *InventoryConfiguration {
	s.Destination = v
	return s
}

func (s *InventoryConfiguration) SetFilter(v *InventoryFilter) *InventoryConfiguration {
	s.Filter = v
	return s
}

func (s *InventoryConfiguration) SetId(v string) *InventoryConfiguration {
	s.Id = &v
	return s
}

func (s *InventoryConfiguration) SetIncludedObjectVersions(v string) *InventoryConfiguration {
	s.IncludedObjectVersions = &v
	return s
}

func (s *InventoryConfiguration) SetIsEnabled(v bool) *InventoryConfiguration {
	s.IsEnabled = &v
	return s
}

func (s *InventoryConfiguration) SetOptionalFields(v *InventoryConfigurationOptionalFields) *InventoryConfiguration {
	s.OptionalFields = v
	return s
}

func (s *InventoryConfiguration) SetSchedule(v *InventorySchedule) *InventoryConfiguration {
	s.Schedule = v
	return s
}

func (s *InventoryConfiguration) Validate() error {
	if s.Destination != nil {
		if err := s.Destination.Validate(); err != nil {
			return err
		}
	}
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			return err
		}
	}
	if s.OptionalFields != nil {
		if err := s.OptionalFields.Validate(); err != nil {
			return err
		}
	}
	if s.Schedule != nil {
		if err := s.Schedule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InventoryConfigurationOptionalFields struct {
	Fields []*string `json:"Field,omitempty" xml:"Field,omitempty" type:"Repeated"`
}

func (s InventoryConfigurationOptionalFields) String() string {
	return dara.Prettify(s)
}

func (s InventoryConfigurationOptionalFields) GoString() string {
	return s.String()
}

func (s *InventoryConfigurationOptionalFields) GetFields() []*string {
	return s.Fields
}

func (s *InventoryConfigurationOptionalFields) SetFields(v []*string) *InventoryConfigurationOptionalFields {
	s.Fields = v
	return s
}

func (s *InventoryConfigurationOptionalFields) Validate() error {
	return dara.Validate(s)
}
