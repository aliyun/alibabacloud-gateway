// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVectorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIndexName(v string) *DeleteVectorsRequest
	GetIndexName() *string
	SetKyes(v []*string) *DeleteVectorsRequest
	GetKyes() []*string
}

type DeleteVectorsRequest struct {
	IndexName *string   `json:"indexName,omitempty" xml:"indexName,omitempty"`
	Kyes      []*string `json:"kyes,omitempty" xml:"kyes,omitempty" type:"Repeated"`
}

func (s DeleteVectorsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVectorsRequest) GoString() string {
	return s.String()
}

func (s *DeleteVectorsRequest) GetIndexName() *string {
	return s.IndexName
}

func (s *DeleteVectorsRequest) GetKyes() []*string {
	return s.Kyes
}

func (s *DeleteVectorsRequest) SetIndexName(v string) *DeleteVectorsRequest {
	s.IndexName = &v
	return s
}

func (s *DeleteVectorsRequest) SetKyes(v []*string) *DeleteVectorsRequest {
	s.Kyes = v
	return s
}

func (s *DeleteVectorsRequest) Validate() error {
	return dara.Validate(s)
}
