// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReservedCapacityRecordList interface {
	dara.Model
	String() string
	GoString() string
	SetReservedCapacityRecord(v []*ReservedCapacityRecord) *ReservedCapacityRecordList
	GetReservedCapacityRecord() []*ReservedCapacityRecord
}

type ReservedCapacityRecordList struct {
	ReservedCapacityRecord []*ReservedCapacityRecord `json:"ReservedCapacityRecord,omitempty" xml:"ReservedCapacityRecord,omitempty" type:"Repeated"`
}

func (s ReservedCapacityRecordList) String() string {
	return dara.Prettify(s)
}

func (s ReservedCapacityRecordList) GoString() string {
	return s.String()
}

func (s *ReservedCapacityRecordList) GetReservedCapacityRecord() []*ReservedCapacityRecord {
	return s.ReservedCapacityRecord
}

func (s *ReservedCapacityRecordList) SetReservedCapacityRecord(v []*ReservedCapacityRecord) *ReservedCapacityRecordList {
	s.ReservedCapacityRecord = v
	return s
}

func (s *ReservedCapacityRecordList) Validate() error {
	if s.ReservedCapacityRecord != nil {
		for _, item := range s.ReservedCapacityRecord {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
