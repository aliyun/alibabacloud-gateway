// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVectorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIndexName(v string) *ListVectorsRequest
	GetIndexName() *string
	SetMaxResults(v int64) *ListVectorsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListVectorsRequest
	GetNextToken() *string
	SetReturnData(v bool) *ListVectorsRequest
	GetReturnData() *bool
	SetReturnMetadata(v bool) *ListVectorsRequest
	GetReturnMetadata() *bool
	SetSegmentCount(v int64) *ListVectorsRequest
	GetSegmentCount() *int64
	SetSegmentIndex(v int64) *ListVectorsRequest
	GetSegmentIndex() *int64
}

type ListVectorsRequest struct {
	IndexName      *string `json:"indexName,omitempty" xml:"indexName,omitempty"`
	MaxResults     *int64  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken      *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ReturnData     *bool   `json:"returnData,omitempty" xml:"returnData,omitempty"`
	ReturnMetadata *bool   `json:"returnMetadata,omitempty" xml:"returnMetadata,omitempty"`
	SegmentCount   *int64  `json:"segmentCount,omitempty" xml:"segmentCount,omitempty"`
	SegmentIndex   *int64  `json:"segmentIndex,omitempty" xml:"segmentIndex,omitempty"`
}

func (s ListVectorsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVectorsRequest) GoString() string {
	return s.String()
}

func (s *ListVectorsRequest) GetIndexName() *string {
	return s.IndexName
}

func (s *ListVectorsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListVectorsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVectorsRequest) GetReturnData() *bool {
	return s.ReturnData
}

func (s *ListVectorsRequest) GetReturnMetadata() *bool {
	return s.ReturnMetadata
}

func (s *ListVectorsRequest) GetSegmentCount() *int64 {
	return s.SegmentCount
}

func (s *ListVectorsRequest) GetSegmentIndex() *int64 {
	return s.SegmentIndex
}

func (s *ListVectorsRequest) SetIndexName(v string) *ListVectorsRequest {
	s.IndexName = &v
	return s
}

func (s *ListVectorsRequest) SetMaxResults(v int64) *ListVectorsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVectorsRequest) SetNextToken(v string) *ListVectorsRequest {
	s.NextToken = &v
	return s
}

func (s *ListVectorsRequest) SetReturnData(v bool) *ListVectorsRequest {
	s.ReturnData = &v
	return s
}

func (s *ListVectorsRequest) SetReturnMetadata(v bool) *ListVectorsRequest {
	s.ReturnMetadata = &v
	return s
}

func (s *ListVectorsRequest) SetSegmentCount(v int64) *ListVectorsRequest {
	s.SegmentCount = &v
	return s
}

func (s *ListVectorsRequest) SetSegmentIndex(v int64) *ListVectorsRequest {
	s.SegmentIndex = &v
	return s
}

func (s *ListVectorsRequest) Validate() error {
	return dara.Validate(s)
}
