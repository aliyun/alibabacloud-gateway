// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBucket interface {
	dara.Model
	String() string
	GoString() string
	SetCreationDate(v string) *Bucket
	GetCreationDate() *string
	SetExtranetEndpoint(v string) *Bucket
	GetExtranetEndpoint() *string
	SetIntranetEndpoint(v string) *Bucket
	GetIntranetEndpoint() *string
	SetLocation(v string) *Bucket
	GetLocation() *string
	SetName(v string) *Bucket
	GetName() *string
	SetRegion(v string) *Bucket
	GetRegion() *string
	SetStorageClass(v string) *Bucket
	GetStorageClass() *string
}

type Bucket struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	CreationDate     *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	ExtranetEndpoint *string `json:"ExtranetEndpoint,omitempty" xml:"ExtranetEndpoint,omitempty"`
	IntranetEndpoint *string `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	Location         *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Region           *string `json:"Region,omitempty" xml:"Region,omitempty"`
	StorageClass     *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
}

func (s Bucket) String() string {
	return dara.Prettify(s)
}

func (s Bucket) GoString() string {
	return s.String()
}

func (s *Bucket) GetCreationDate() *string {
	return s.CreationDate
}

func (s *Bucket) GetExtranetEndpoint() *string {
	return s.ExtranetEndpoint
}

func (s *Bucket) GetIntranetEndpoint() *string {
	return s.IntranetEndpoint
}

func (s *Bucket) GetLocation() *string {
	return s.Location
}

func (s *Bucket) GetName() *string {
	return s.Name
}

func (s *Bucket) GetRegion() *string {
	return s.Region
}

func (s *Bucket) GetStorageClass() *string {
	return s.StorageClass
}

func (s *Bucket) SetCreationDate(v string) *Bucket {
	s.CreationDate = &v
	return s
}

func (s *Bucket) SetExtranetEndpoint(v string) *Bucket {
	s.ExtranetEndpoint = &v
	return s
}

func (s *Bucket) SetIntranetEndpoint(v string) *Bucket {
	s.IntranetEndpoint = &v
	return s
}

func (s *Bucket) SetLocation(v string) *Bucket {
	s.Location = &v
	return s
}

func (s *Bucket) SetName(v string) *Bucket {
	s.Name = &v
	return s
}

func (s *Bucket) SetRegion(v string) *Bucket {
	s.Region = &v
	return s
}

func (s *Bucket) SetStorageClass(v string) *Bucket {
	s.StorageClass = &v
	return s
}

func (s *Bucket) Validate() error {
	return dara.Validate(s)
}
