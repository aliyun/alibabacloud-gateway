// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVectorIndexPut interface {
	dara.Model
	String() string
	GoString() string
	SetDataType(v string) *VectorIndexPut
	GetDataType() *string
	SetDimension(v int64) *VectorIndexPut
	GetDimension() *int64
	SetDistanceMetric(v string) *VectorIndexPut
	GetDistanceMetric() *string
	SetIndexName(v string) *VectorIndexPut
	GetIndexName() *string
	SetMetadata(v *VectorIndexMetaData) *VectorIndexPut
	GetMetadata() *VectorIndexMetaData
}

type VectorIndexPut struct {
	// example:
	//
	// float32
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 100
	Dimension *int64 `json:"dimension,omitempty" xml:"dimension,omitempty"`
	// example:
	//
	// euclidean
	DistanceMetric *string `json:"distanceMetric,omitempty" xml:"distanceMetric,omitempty"`
	// example:
	//
	// test-vector
	IndexName *string              `json:"indexName,omitempty" xml:"indexName,omitempty"`
	Metadata  *VectorIndexMetaData `json:"metadata,omitempty" xml:"metadata,omitempty"`
}

func (s VectorIndexPut) String() string {
	return dara.Prettify(s)
}

func (s VectorIndexPut) GoString() string {
	return s.String()
}

func (s *VectorIndexPut) GetDataType() *string {
	return s.DataType
}

func (s *VectorIndexPut) GetDimension() *int64 {
	return s.Dimension
}

func (s *VectorIndexPut) GetDistanceMetric() *string {
	return s.DistanceMetric
}

func (s *VectorIndexPut) GetIndexName() *string {
	return s.IndexName
}

func (s *VectorIndexPut) GetMetadata() *VectorIndexMetaData {
	return s.Metadata
}

func (s *VectorIndexPut) SetDataType(v string) *VectorIndexPut {
	s.DataType = &v
	return s
}

func (s *VectorIndexPut) SetDimension(v int64) *VectorIndexPut {
	s.Dimension = &v
	return s
}

func (s *VectorIndexPut) SetDistanceMetric(v string) *VectorIndexPut {
	s.DistanceMetric = &v
	return s
}

func (s *VectorIndexPut) SetIndexName(v string) *VectorIndexPut {
	s.IndexName = &v
	return s
}

func (s *VectorIndexPut) SetMetadata(v *VectorIndexMetaData) *VectorIndexPut {
	s.Metadata = v
	return s
}

func (s *VectorIndexPut) Validate() error {
	if s.Metadata != nil {
		if err := s.Metadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}
