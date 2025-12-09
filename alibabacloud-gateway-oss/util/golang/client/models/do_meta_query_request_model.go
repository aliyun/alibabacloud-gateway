// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDoMetaQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMetaQuery(v *MetaQuery) *DoMetaQueryRequest
	GetMetaQuery() *MetaQuery
	SetMode(v string) *DoMetaQueryRequest
	GetMode() *string
}

type DoMetaQueryRequest struct {
	// The containerfor query conditions.
	MetaQuery *MetaQuery `json:"MetaQuery,omitempty" xml:"MetaQuery,omitempty"`
	Mode      *string    `json:"mode,omitempty" xml:"mode,omitempty"`
}

func (s DoMetaQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s DoMetaQueryRequest) GoString() string {
	return s.String()
}

func (s *DoMetaQueryRequest) GetMetaQuery() *MetaQuery {
	return s.MetaQuery
}

func (s *DoMetaQueryRequest) GetMode() *string {
	return s.Mode
}

func (s *DoMetaQueryRequest) SetMetaQuery(v *MetaQuery) *DoMetaQueryRequest {
	s.MetaQuery = v
	return s
}

func (s *DoMetaQueryRequest) SetMode(v string) *DoMetaQueryRequest {
	s.Mode = &v
	return s
}

func (s *DoMetaQueryRequest) Validate() error {
	if s.MetaQuery != nil {
		if err := s.MetaQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}
