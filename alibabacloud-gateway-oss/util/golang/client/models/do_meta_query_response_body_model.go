// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDoMetaQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetaQuery(v *MetaQueryResp) *DoMetaQueryResponseBody
	GetMetaQuery() *MetaQueryResp
}

type DoMetaQueryResponseBody struct {
	// The container for the query result.
	MetaQuery *MetaQueryResp `json:"MetaQuery,omitempty" xml:"MetaQuery,omitempty"`
}

func (s DoMetaQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DoMetaQueryResponseBody) GoString() string {
	return s.String()
}

func (s *DoMetaQueryResponseBody) GetMetaQuery() *MetaQueryResp {
	return s.MetaQuery
}

func (s *DoMetaQueryResponseBody) SetMetaQuery(v *MetaQueryResp) *DoMetaQueryResponseBody {
	s.MetaQuery = v
	return s
}

func (s *DoMetaQueryResponseBody) Validate() error {
	if s.MetaQuery != nil {
		if err := s.MetaQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}
