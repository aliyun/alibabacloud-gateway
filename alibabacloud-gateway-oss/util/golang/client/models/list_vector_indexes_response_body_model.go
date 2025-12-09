// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVectorIndexesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIndexes(v []*VectorIndex) *ListVectorIndexesResponseBody
	GetIndexes() []*VectorIndex
	SetNextToken(v string) *ListVectorIndexesResponseBody
	GetNextToken() *string
}

type ListVectorIndexesResponseBody struct {
	Indexes   []*VectorIndex `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
	NextToken *string        `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListVectorIndexesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVectorIndexesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVectorIndexesResponseBody) GetIndexes() []*VectorIndex {
	return s.Indexes
}

func (s *ListVectorIndexesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVectorIndexesResponseBody) SetIndexes(v []*VectorIndex) *ListVectorIndexesResponseBody {
	s.Indexes = v
	return s
}

func (s *ListVectorIndexesResponseBody) SetNextToken(v string) *ListVectorIndexesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVectorIndexesResponseBody) Validate() error {
	if s.Indexes != nil {
		for _, item := range s.Indexes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
