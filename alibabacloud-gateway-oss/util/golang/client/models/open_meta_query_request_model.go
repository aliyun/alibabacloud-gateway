// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenMetaQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMetaQuery(v *MetaQueryOpenRequest) *OpenMetaQueryRequest
	GetMetaQuery() *MetaQueryOpenRequest
	SetMode(v string) *OpenMetaQueryRequest
	GetMode() *string
	SetRole(v string) *OpenMetaQueryRequest
	GetRole() *string
}

type OpenMetaQueryRequest struct {
	MetaQuery *MetaQueryOpenRequest `json:"MetaQuery,omitempty" xml:"MetaQuery,omitempty"`
	Mode      *string               `json:"mode,omitempty" xml:"mode,omitempty"`
	Role      *string               `json:"role,omitempty" xml:"role,omitempty"`
}

func (s OpenMetaQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenMetaQueryRequest) GoString() string {
	return s.String()
}

func (s *OpenMetaQueryRequest) GetMetaQuery() *MetaQueryOpenRequest {
	return s.MetaQuery
}

func (s *OpenMetaQueryRequest) GetMode() *string {
	return s.Mode
}

func (s *OpenMetaQueryRequest) GetRole() *string {
	return s.Role
}

func (s *OpenMetaQueryRequest) SetMetaQuery(v *MetaQueryOpenRequest) *OpenMetaQueryRequest {
	s.MetaQuery = v
	return s
}

func (s *OpenMetaQueryRequest) SetMode(v string) *OpenMetaQueryRequest {
	s.Mode = &v
	return s
}

func (s *OpenMetaQueryRequest) SetRole(v string) *OpenMetaQueryRequest {
	s.Role = &v
	return s
}

func (s *OpenMetaQueryRequest) Validate() error {
	if s.MetaQuery != nil {
		if err := s.MetaQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}
