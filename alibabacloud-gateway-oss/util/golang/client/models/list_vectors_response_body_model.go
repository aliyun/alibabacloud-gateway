// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVectorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListVectorsResponseBody
	GetNextToken() *string
	SetVectors(v []*Vector) *ListVectorsResponseBody
	GetVectors() []*Vector
}

type ListVectorsResponseBody struct {
	NextToken *string   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Vectors   []*Vector `json:"vectors,omitempty" xml:"vectors,omitempty" type:"Repeated"`
}

func (s ListVectorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVectorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVectorsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVectorsResponseBody) GetVectors() []*Vector {
	return s.Vectors
}

func (s *ListVectorsResponseBody) SetNextToken(v string) *ListVectorsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVectorsResponseBody) SetVectors(v []*Vector) *ListVectorsResponseBody {
	s.Vectors = v
	return s
}

func (s *ListVectorsResponseBody) Validate() error {
	if s.Vectors != nil {
		for _, item := range s.Vectors {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
