// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVectorIndexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIndexName(v string) *DeleteVectorIndexRequest
	GetIndexName() *string
}

type DeleteVectorIndexRequest struct {
	IndexName *string `json:"indexName,omitempty" xml:"indexName,omitempty"`
}

func (s DeleteVectorIndexRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVectorIndexRequest) GoString() string {
	return s.String()
}

func (s *DeleteVectorIndexRequest) GetIndexName() *string {
	return s.IndexName
}

func (s *DeleteVectorIndexRequest) SetIndexName(v string) *DeleteVectorIndexRequest {
	s.IndexName = &v
	return s
}

func (s *DeleteVectorIndexRequest) Validate() error {
	return dara.Validate(s)
}
