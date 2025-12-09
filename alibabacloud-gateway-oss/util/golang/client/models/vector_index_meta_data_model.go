// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVectorIndexMetaData interface {
	dara.Model
	String() string
	GoString() string
	SetNonFilterableMetadataKeys(v []*string) *VectorIndexMetaData
	GetNonFilterableMetadataKeys() []*string
}

type VectorIndexMetaData struct {
	NonFilterableMetadataKeys []*string `json:"nonFilterableMetadataKeys,omitempty" xml:"nonFilterableMetadataKeys,omitempty" type:"Repeated"`
}

func (s VectorIndexMetaData) String() string {
	return dara.Prettify(s)
}

func (s VectorIndexMetaData) GoString() string {
	return s.String()
}

func (s *VectorIndexMetaData) GetNonFilterableMetadataKeys() []*string {
	return s.NonFilterableMetadataKeys
}

func (s *VectorIndexMetaData) SetNonFilterableMetadataKeys(v []*string) *VectorIndexMetaData {
	s.NonFilterableMetadataKeys = v
	return s
}

func (s *VectorIndexMetaData) Validate() error {
	return dara.Validate(s)
}
