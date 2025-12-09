// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetaQuery interface {
	dara.Model
	String() string
	GoString() string
	SetAggregations(v *MetaQueryAggregations) *MetaQuery
	GetAggregations() *MetaQueryAggregations
	SetMaxResults(v int64) *MetaQuery
	GetMaxResults() *int64
	SetMediaTypes(v *MetaQueryMediaTypes) *MetaQuery
	GetMediaTypes() *MetaQueryMediaTypes
	SetNextToken(v string) *MetaQuery
	GetNextToken() *string
	SetOrder(v string) *MetaQuery
	GetOrder() *string
	SetQuery(v string) *MetaQuery
	GetQuery() *string
	SetSimpleQuery(v string) *MetaQuery
	GetSimpleQuery() *string
	SetSort(v string) *MetaQuery
	GetSort() *string
}

type MetaQuery struct {
	Aggregations *MetaQueryAggregations `json:"Aggregations,omitempty" xml:"Aggregations,omitempty" type:"Struct"`
	// example:
	//
	// 100
	MaxResults *int64               `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	MediaTypes *MetaQueryMediaTypes `json:"MediaTypes,omitempty" xml:"MediaTypes,omitempty" type:"Struct"`
	// example:
	//
	// MTIzNDU2Nzg6aW1tdGVzdDpleGFtcGxlYnVja2V0OmRhdGFzZXQwMDE6b3NzOi8vZXhhbXBsZWJ1Y2tldC9zYW1wbGVvYmplY3QxLmpw****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Order     *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// {"Field": "Size","Value": "1048576","Operation": "gt"}
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// {"Operation":"gt", "Field": "Size", "Value": "30"}
	SimpleQuery *string `json:"SimpleQuery,omitempty" xml:"SimpleQuery,omitempty"`
	// example:
	//
	// Size
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
}

func (s MetaQuery) String() string {
	return dara.Prettify(s)
}

func (s MetaQuery) GoString() string {
	return s.String()
}

func (s *MetaQuery) GetAggregations() *MetaQueryAggregations {
	return s.Aggregations
}

func (s *MetaQuery) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *MetaQuery) GetMediaTypes() *MetaQueryMediaTypes {
	return s.MediaTypes
}

func (s *MetaQuery) GetNextToken() *string {
	return s.NextToken
}

func (s *MetaQuery) GetOrder() *string {
	return s.Order
}

func (s *MetaQuery) GetQuery() *string {
	return s.Query
}

func (s *MetaQuery) GetSimpleQuery() *string {
	return s.SimpleQuery
}

func (s *MetaQuery) GetSort() *string {
	return s.Sort
}

func (s *MetaQuery) SetAggregations(v *MetaQueryAggregations) *MetaQuery {
	s.Aggregations = v
	return s
}

func (s *MetaQuery) SetMaxResults(v int64) *MetaQuery {
	s.MaxResults = &v
	return s
}

func (s *MetaQuery) SetMediaTypes(v *MetaQueryMediaTypes) *MetaQuery {
	s.MediaTypes = v
	return s
}

func (s *MetaQuery) SetNextToken(v string) *MetaQuery {
	s.NextToken = &v
	return s
}

func (s *MetaQuery) SetOrder(v string) *MetaQuery {
	s.Order = &v
	return s
}

func (s *MetaQuery) SetQuery(v string) *MetaQuery {
	s.Query = &v
	return s
}

func (s *MetaQuery) SetSimpleQuery(v string) *MetaQuery {
	s.SimpleQuery = &v
	return s
}

func (s *MetaQuery) SetSort(v string) *MetaQuery {
	s.Sort = &v
	return s
}

func (s *MetaQuery) Validate() error {
	if s.Aggregations != nil {
		if err := s.Aggregations.Validate(); err != nil {
			return err
		}
	}
	if s.MediaTypes != nil {
		if err := s.MediaTypes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MetaQueryAggregations struct {
	Aggregation []*MetaQueryAggregation `json:"Aggregation,omitempty" xml:"Aggregation,omitempty" type:"Repeated"`
}

func (s MetaQueryAggregations) String() string {
	return dara.Prettify(s)
}

func (s MetaQueryAggregations) GoString() string {
	return s.String()
}

func (s *MetaQueryAggregations) GetAggregation() []*MetaQueryAggregation {
	return s.Aggregation
}

func (s *MetaQueryAggregations) SetAggregation(v []*MetaQueryAggregation) *MetaQueryAggregations {
	s.Aggregation = v
	return s
}

func (s *MetaQueryAggregations) Validate() error {
	if s.Aggregation != nil {
		for _, item := range s.Aggregation {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MetaQueryMediaTypes struct {
	MediaType []*string `json:"MediaType,omitempty" xml:"MediaType,omitempty" type:"Repeated"`
}

func (s MetaQueryMediaTypes) String() string {
	return dara.Prettify(s)
}

func (s MetaQueryMediaTypes) GoString() string {
	return s.String()
}

func (s *MetaQueryMediaTypes) GetMediaType() []*string {
	return s.MediaType
}

func (s *MetaQueryMediaTypes) SetMediaType(v []*string) *MetaQueryMediaTypes {
	s.MediaType = v
	return s
}

func (s *MetaQueryMediaTypes) Validate() error {
	return dara.Validate(s)
}
