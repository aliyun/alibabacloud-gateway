// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketRedundancyTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataRedundancyTypeConfiguration(v *PutBucketRedundancyTypeRequestDataRedundancyTypeConfiguration) *PutBucketRedundancyTypeRequest
	GetDataRedundancyTypeConfiguration() *PutBucketRedundancyTypeRequestDataRedundancyTypeConfiguration
}

type PutBucketRedundancyTypeRequest struct {
	DataRedundancyTypeConfiguration *PutBucketRedundancyTypeRequestDataRedundancyTypeConfiguration `json:"DataRedundancyTypeConfiguration,omitempty" xml:"DataRedundancyTypeConfiguration,omitempty" type:"Struct"`
}

func (s PutBucketRedundancyTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s PutBucketRedundancyTypeRequest) GoString() string {
	return s.String()
}

func (s *PutBucketRedundancyTypeRequest) GetDataRedundancyTypeConfiguration() *PutBucketRedundancyTypeRequestDataRedundancyTypeConfiguration {
	return s.DataRedundancyTypeConfiguration
}

func (s *PutBucketRedundancyTypeRequest) SetDataRedundancyTypeConfiguration(v *PutBucketRedundancyTypeRequestDataRedundancyTypeConfiguration) *PutBucketRedundancyTypeRequest {
	s.DataRedundancyTypeConfiguration = v
	return s
}

func (s *PutBucketRedundancyTypeRequest) Validate() error {
	if s.DataRedundancyTypeConfiguration != nil {
		if err := s.DataRedundancyTypeConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PutBucketRedundancyTypeRequestDataRedundancyTypeConfiguration struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PutBucketRedundancyTypeRequestDataRedundancyTypeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s PutBucketRedundancyTypeRequestDataRedundancyTypeConfiguration) GoString() string {
	return s.String()
}

func (s *PutBucketRedundancyTypeRequestDataRedundancyTypeConfiguration) GetType() *string {
	return s.Type
}

func (s *PutBucketRedundancyTypeRequestDataRedundancyTypeConfiguration) SetType(v string) *PutBucketRedundancyTypeRequestDataRedundancyTypeConfiguration {
	s.Type = &v
	return s
}

func (s *PutBucketRedundancyTypeRequestDataRedundancyTypeConfiguration) Validate() error {
	return dara.Validate(s)
}
