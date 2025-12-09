// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVectorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v interface{}) *QueryVectorsRequest
	GetFilter() interface{}
	SetIndexName(v string) *QueryVectorsRequest
	GetIndexName() *string
	SetQueryVector(v *VectorData) *QueryVectorsRequest
	GetQueryVector() *VectorData
	SetReturnDistance(v bool) *QueryVectorsRequest
	GetReturnDistance() *bool
	SetReturnMetadata(v bool) *QueryVectorsRequest
	GetReturnMetadata() *bool
	SetTopK(v int64) *QueryVectorsRequest
	GetTopK() *int64
}

type QueryVectorsRequest struct {
	Filter         interface{} `json:"filter,omitempty" xml:"filter,omitempty"`
	IndexName      *string     `json:"indexName,omitempty" xml:"indexName,omitempty"`
	QueryVector    *VectorData `json:"queryVector,omitempty" xml:"queryVector,omitempty"`
	ReturnDistance *bool       `json:"returnDistance,omitempty" xml:"returnDistance,omitempty"`
	ReturnMetadata *bool       `json:"returnMetadata,omitempty" xml:"returnMetadata,omitempty"`
	TopK           *int64      `json:"topK,omitempty" xml:"topK,omitempty"`
}

func (s QueryVectorsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryVectorsRequest) GoString() string {
	return s.String()
}

func (s *QueryVectorsRequest) GetFilter() interface{} {
	return s.Filter
}

func (s *QueryVectorsRequest) GetIndexName() *string {
	return s.IndexName
}

func (s *QueryVectorsRequest) GetQueryVector() *VectorData {
	return s.QueryVector
}

func (s *QueryVectorsRequest) GetReturnDistance() *bool {
	return s.ReturnDistance
}

func (s *QueryVectorsRequest) GetReturnMetadata() *bool {
	return s.ReturnMetadata
}

func (s *QueryVectorsRequest) GetTopK() *int64 {
	return s.TopK
}

func (s *QueryVectorsRequest) SetFilter(v interface{}) *QueryVectorsRequest {
	s.Filter = v
	return s
}

func (s *QueryVectorsRequest) SetIndexName(v string) *QueryVectorsRequest {
	s.IndexName = &v
	return s
}

func (s *QueryVectorsRequest) SetQueryVector(v *VectorData) *QueryVectorsRequest {
	s.QueryVector = v
	return s
}

func (s *QueryVectorsRequest) SetReturnDistance(v bool) *QueryVectorsRequest {
	s.ReturnDistance = &v
	return s
}

func (s *QueryVectorsRequest) SetReturnMetadata(v bool) *QueryVectorsRequest {
	s.ReturnMetadata = &v
	return s
}

func (s *QueryVectorsRequest) SetTopK(v int64) *QueryVectorsRequest {
	s.TopK = &v
	return s
}

func (s *QueryVectorsRequest) Validate() error {
	if s.QueryVector != nil {
		if err := s.QueryVector.Validate(); err != nil {
			return err
		}
	}
	return nil
}
