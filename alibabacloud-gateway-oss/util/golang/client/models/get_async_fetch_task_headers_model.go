// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsyncFetchTaskHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetAsyncFetchTaskHeaders
	GetCommonHeaders() map[string]*string
	SetXOssTaskId(v string) *GetAsyncFetchTaskHeaders
	GetXOssTaskId() *string
}

type GetAsyncFetchTaskHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	XOssTaskId *string `json:"x-oss-task-id,omitempty" xml:"x-oss-task-id,omitempty"`
}

func (s GetAsyncFetchTaskHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncFetchTaskHeaders) GoString() string {
	return s.String()
}

func (s *GetAsyncFetchTaskHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetAsyncFetchTaskHeaders) GetXOssTaskId() *string {
	return s.XOssTaskId
}

func (s *GetAsyncFetchTaskHeaders) SetCommonHeaders(v map[string]*string) *GetAsyncFetchTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAsyncFetchTaskHeaders) SetXOssTaskId(v string) *GetAsyncFetchTaskHeaders {
	s.XOssTaskId = &v
	return s
}

func (s *GetAsyncFetchTaskHeaders) Validate() error {
	return dara.Validate(s)
}
