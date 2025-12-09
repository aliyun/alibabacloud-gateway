// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketReplicationLocationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetReplicationLocation(v *GetBucketReplicationLocationResponseBodyReplicationLocation) *GetBucketReplicationLocationResponseBody
	GetReplicationLocation() *GetBucketReplicationLocationResponseBodyReplicationLocation
}

type GetBucketReplicationLocationResponseBody struct {
	// The container that stores the region in which the destination bucket can be located.
	ReplicationLocation *GetBucketReplicationLocationResponseBodyReplicationLocation `json:"ReplicationLocation,omitempty" xml:"ReplicationLocation,omitempty" type:"Struct"`
}

func (s GetBucketReplicationLocationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBucketReplicationLocationResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketReplicationLocationResponseBody) GetReplicationLocation() *GetBucketReplicationLocationResponseBodyReplicationLocation {
	return s.ReplicationLocation
}

func (s *GetBucketReplicationLocationResponseBody) SetReplicationLocation(v *GetBucketReplicationLocationResponseBodyReplicationLocation) *GetBucketReplicationLocationResponseBody {
	s.ReplicationLocation = v
	return s
}

func (s *GetBucketReplicationLocationResponseBody) Validate() error {
	if s.ReplicationLocation != nil {
		if err := s.ReplicationLocation.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBucketReplicationLocationResponseBodyReplicationLocation struct {
	// The regions in which the destination bucket can be located.
	Location []*string `json:"Location,omitempty" xml:"Location,omitempty" type:"Repeated"`
	// The container that stores regions in which the RTC can be enabled.
	LocationRTCConstraint *GetBucketReplicationLocationResponseBodyReplicationLocationLocationRTCConstraint `json:"LocationRTCConstraint,omitempty" xml:"LocationRTCConstraint,omitempty" type:"Struct"`
	// The container that stores regions in which the destination bucket can be located with TransferType specified.
	LocationTransferTypeConstraint *GetBucketReplicationLocationResponseBodyReplicationLocationLocationTransferTypeConstraint `json:"LocationTransferTypeConstraint,omitempty" xml:"LocationTransferTypeConstraint,omitempty" type:"Struct"`
}

func (s GetBucketReplicationLocationResponseBodyReplicationLocation) String() string {
	return dara.Prettify(s)
}

func (s GetBucketReplicationLocationResponseBodyReplicationLocation) GoString() string {
	return s.String()
}

func (s *GetBucketReplicationLocationResponseBodyReplicationLocation) GetLocation() []*string {
	return s.Location
}

func (s *GetBucketReplicationLocationResponseBodyReplicationLocation) GetLocationRTCConstraint() *GetBucketReplicationLocationResponseBodyReplicationLocationLocationRTCConstraint {
	return s.LocationRTCConstraint
}

func (s *GetBucketReplicationLocationResponseBodyReplicationLocation) GetLocationTransferTypeConstraint() *GetBucketReplicationLocationResponseBodyReplicationLocationLocationTransferTypeConstraint {
	return s.LocationTransferTypeConstraint
}

func (s *GetBucketReplicationLocationResponseBodyReplicationLocation) SetLocation(v []*string) *GetBucketReplicationLocationResponseBodyReplicationLocation {
	s.Location = v
	return s
}

func (s *GetBucketReplicationLocationResponseBodyReplicationLocation) SetLocationRTCConstraint(v *GetBucketReplicationLocationResponseBodyReplicationLocationLocationRTCConstraint) *GetBucketReplicationLocationResponseBodyReplicationLocation {
	s.LocationRTCConstraint = v
	return s
}

func (s *GetBucketReplicationLocationResponseBodyReplicationLocation) SetLocationTransferTypeConstraint(v *GetBucketReplicationLocationResponseBodyReplicationLocationLocationTransferTypeConstraint) *GetBucketReplicationLocationResponseBodyReplicationLocation {
	s.LocationTransferTypeConstraint = v
	return s
}

func (s *GetBucketReplicationLocationResponseBodyReplicationLocation) Validate() error {
	if s.LocationRTCConstraint != nil {
		if err := s.LocationRTCConstraint.Validate(); err != nil {
			return err
		}
	}
	if s.LocationTransferTypeConstraint != nil {
		if err := s.LocationTransferTypeConstraint.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBucketReplicationLocationResponseBodyReplicationLocationLocationRTCConstraint struct {
	// The regions where RTC is supported.
	Location []*string `json:"Location,omitempty" xml:"Location,omitempty" type:"Repeated"`
}

func (s GetBucketReplicationLocationResponseBodyReplicationLocationLocationRTCConstraint) String() string {
	return dara.Prettify(s)
}

func (s GetBucketReplicationLocationResponseBodyReplicationLocationLocationRTCConstraint) GoString() string {
	return s.String()
}

func (s *GetBucketReplicationLocationResponseBodyReplicationLocationLocationRTCConstraint) GetLocation() []*string {
	return s.Location
}

func (s *GetBucketReplicationLocationResponseBodyReplicationLocationLocationRTCConstraint) SetLocation(v []*string) *GetBucketReplicationLocationResponseBodyReplicationLocationLocationRTCConstraint {
	s.Location = v
	return s
}

func (s *GetBucketReplicationLocationResponseBodyReplicationLocationLocationRTCConstraint) Validate() error {
	return dara.Validate(s)
}

type GetBucketReplicationLocationResponseBodyReplicationLocationLocationTransferTypeConstraint struct {
	// The container that stores regions in which the destination bucket can be located with the TransferType information.
	LocationTransferType []*LocationTransferType `json:"LocationTransferType,omitempty" xml:"LocationTransferType,omitempty" type:"Repeated"`
}

func (s GetBucketReplicationLocationResponseBodyReplicationLocationLocationTransferTypeConstraint) String() string {
	return dara.Prettify(s)
}

func (s GetBucketReplicationLocationResponseBodyReplicationLocationLocationTransferTypeConstraint) GoString() string {
	return s.String()
}

func (s *GetBucketReplicationLocationResponseBodyReplicationLocationLocationTransferTypeConstraint) GetLocationTransferType() []*LocationTransferType {
	return s.LocationTransferType
}

func (s *GetBucketReplicationLocationResponseBodyReplicationLocationLocationTransferTypeConstraint) SetLocationTransferType(v []*LocationTransferType) *GetBucketReplicationLocationResponseBodyReplicationLocationLocationTransferTypeConstraint {
	s.LocationTransferType = v
	return s
}

func (s *GetBucketReplicationLocationResponseBodyReplicationLocationLocationTransferTypeConstraint) Validate() error {
	if s.LocationTransferType != nil {
		for _, item := range s.LocationTransferType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
