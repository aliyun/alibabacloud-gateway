// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVectorIndex interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *VectorIndex
	GetCreateTime() *string
	SetDataType(v string) *VectorIndex
	GetDataType() *string
	SetDimension(v int64) *VectorIndex
	GetDimension() *int64
	SetDistanceMetric(v string) *VectorIndex
	GetDistanceMetric() *string
	SetIndexName(v string) *VectorIndex
	GetIndexName() *string
	SetMetadata(v *VectorIndexMetaData) *VectorIndex
	GetMetadata() *VectorIndexMetaData
	SetStatus(v string) *VectorIndex
	GetStatus() *string
	SetVectorBucketName(v string) *VectorIndex
	GetVectorBucketName() *string
}

type VectorIndex struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2025-08-31T10:56:21.000Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
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
	// example:
	//
	// running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// test-vector-bucket
	VectorBucketName *string `json:"vectorBucketName,omitempty" xml:"vectorBucketName,omitempty"`
}

func (s VectorIndex) String() string {
	return dara.Prettify(s)
}

func (s VectorIndex) GoString() string {
	return s.String()
}

func (s *VectorIndex) GetCreateTime() *string {
	return s.CreateTime
}

func (s *VectorIndex) GetDataType() *string {
	return s.DataType
}

func (s *VectorIndex) GetDimension() *int64 {
	return s.Dimension
}

func (s *VectorIndex) GetDistanceMetric() *string {
	return s.DistanceMetric
}

func (s *VectorIndex) GetIndexName() *string {
	return s.IndexName
}

func (s *VectorIndex) GetMetadata() *VectorIndexMetaData {
	return s.Metadata
}

func (s *VectorIndex) GetStatus() *string {
	return s.Status
}

func (s *VectorIndex) GetVectorBucketName() *string {
	return s.VectorBucketName
}

func (s *VectorIndex) SetCreateTime(v string) *VectorIndex {
	s.CreateTime = &v
	return s
}

func (s *VectorIndex) SetDataType(v string) *VectorIndex {
	s.DataType = &v
	return s
}

func (s *VectorIndex) SetDimension(v int64) *VectorIndex {
	s.Dimension = &v
	return s
}

func (s *VectorIndex) SetDistanceMetric(v string) *VectorIndex {
	s.DistanceMetric = &v
	return s
}

func (s *VectorIndex) SetIndexName(v string) *VectorIndex {
	s.IndexName = &v
	return s
}

func (s *VectorIndex) SetMetadata(v *VectorIndexMetaData) *VectorIndex {
	s.Metadata = v
	return s
}

func (s *VectorIndex) SetStatus(v string) *VectorIndex {
	s.Status = &v
	return s
}

func (s *VectorIndex) SetVectorBucketName(v string) *VectorIndex {
	s.VectorBucketName = &v
	return s
}

func (s *VectorIndex) Validate() error {
	if s.Metadata != nil {
		if err := s.Metadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}
