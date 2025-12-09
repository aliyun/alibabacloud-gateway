// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseMetaQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloseMetaQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloseMetaQueryResponse
	GetStatusCode() *int32
}

type CloseMetaQueryResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CloseMetaQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s CloseMetaQueryResponse) GoString() string {
	return s.String()
}

func (s *CloseMetaQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloseMetaQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloseMetaQueryResponse) SetHeaders(v map[string]*string) *CloseMetaQueryResponse {
	s.Headers = v
	return s
}

func (s *CloseMetaQueryResponse) SetStatusCode(v int32) *CloseMetaQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseMetaQueryResponse) Validate() error {
	return dara.Validate(s)
}
