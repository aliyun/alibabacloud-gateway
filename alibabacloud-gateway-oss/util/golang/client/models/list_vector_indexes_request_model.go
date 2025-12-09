// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVectorIndexesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListVectorIndexesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListVectorIndexesRequest
	GetNextToken() *string
	SetPrefix(v string) *ListVectorIndexesRequest
	GetPrefix() *string
}

type ListVectorIndexesRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Prefix     *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
}

func (s ListVectorIndexesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVectorIndexesRequest) GoString() string {
	return s.String()
}

func (s *ListVectorIndexesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVectorIndexesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVectorIndexesRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *ListVectorIndexesRequest) SetMaxResults(v int32) *ListVectorIndexesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVectorIndexesRequest) SetNextToken(v string) *ListVectorIndexesRequest {
	s.NextToken = &v
	return s
}

func (s *ListVectorIndexesRequest) SetPrefix(v string) *ListVectorIndexesRequest {
	s.Prefix = &v
	return s
}

func (s *ListVectorIndexesRequest) Validate() error {
	return dara.Validate(s)
}
