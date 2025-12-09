// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetObjectGroupIndexHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetObjectGroupIndexHeaders
	GetCommonHeaders() map[string]*string
	SetXOssFileGroup(v string) *GetObjectGroupIndexHeaders
	GetXOssFileGroup() *string
}

type GetObjectGroupIndexHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	XOssFileGroup *string `json:"x-oss-file-group,omitempty" xml:"x-oss-file-group,omitempty"`
}

func (s GetObjectGroupIndexHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetObjectGroupIndexHeaders) GoString() string {
	return s.String()
}

func (s *GetObjectGroupIndexHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetObjectGroupIndexHeaders) GetXOssFileGroup() *string {
	return s.XOssFileGroup
}

func (s *GetObjectGroupIndexHeaders) SetCommonHeaders(v map[string]*string) *GetObjectGroupIndexHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetObjectGroupIndexHeaders) SetXOssFileGroup(v string) *GetObjectGroupIndexHeaders {
	s.XOssFileGroup = &v
	return s
}

func (s *GetObjectGroupIndexHeaders) Validate() error {
	return dara.Validate(s)
}
