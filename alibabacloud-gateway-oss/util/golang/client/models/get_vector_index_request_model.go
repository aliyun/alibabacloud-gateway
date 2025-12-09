// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVectorIndexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIndexName(v string) *GetVectorIndexRequest
	GetIndexName() *string
}

type GetVectorIndexRequest struct {
	IndexName *string `json:"indexName,omitempty" xml:"indexName,omitempty"`
}

func (s GetVectorIndexRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVectorIndexRequest) GoString() string {
	return s.String()
}

func (s *GetVectorIndexRequest) GetIndexName() *string {
	return s.IndexName
}

func (s *GetVectorIndexRequest) SetIndexName(v string) *GetVectorIndexRequest {
	s.IndexName = &v
	return s
}

func (s *GetVectorIndexRequest) Validate() error {
	return dara.Validate(s)
}
