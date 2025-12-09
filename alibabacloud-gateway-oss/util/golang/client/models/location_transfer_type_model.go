// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLocationTransferType interface {
	dara.Model
	String() string
	GoString() string
	SetLocation(v string) *LocationTransferType
	GetLocation() *string
	SetTransferTypes(v *LocationTransferTypeTransferTypes) *LocationTransferType
	GetTransferTypes() *LocationTransferTypeTransferTypes
}

type LocationTransferType struct {
	// example:
	//
	// oss-eu-west-1
	Location      *string                            `json:"Location,omitempty" xml:"Location,omitempty"`
	TransferTypes *LocationTransferTypeTransferTypes `json:"TransferTypes,omitempty" xml:"TransferTypes,omitempty" type:"Struct"`
}

func (s LocationTransferType) String() string {
	return dara.Prettify(s)
}

func (s LocationTransferType) GoString() string {
	return s.String()
}

func (s *LocationTransferType) GetLocation() *string {
	return s.Location
}

func (s *LocationTransferType) GetTransferTypes() *LocationTransferTypeTransferTypes {
	return s.TransferTypes
}

func (s *LocationTransferType) SetLocation(v string) *LocationTransferType {
	s.Location = &v
	return s
}

func (s *LocationTransferType) SetTransferTypes(v *LocationTransferTypeTransferTypes) *LocationTransferType {
	s.TransferTypes = v
	return s
}

func (s *LocationTransferType) Validate() error {
	if s.TransferTypes != nil {
		if err := s.TransferTypes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LocationTransferTypeTransferTypes struct {
	Type []*string `json:"Type,omitempty" xml:"Type,omitempty" type:"Repeated"`
}

func (s LocationTransferTypeTransferTypes) String() string {
	return dara.Prettify(s)
}

func (s LocationTransferTypeTransferTypes) GoString() string {
	return s.String()
}

func (s *LocationTransferTypeTransferTypes) GetType() []*string {
	return s.Type
}

func (s *LocationTransferTypeTransferTypes) SetType(v []*string) *LocationTransferTypeTransferTypes {
	s.Type = v
	return s
}

func (s *LocationTransferTypeTransferTypes) Validate() error {
	return dara.Validate(s)
}
