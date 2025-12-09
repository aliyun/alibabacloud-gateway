// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultipleObjectsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteMultipleObjectsHeaders
	GetCommonHeaders() map[string]*string
	SetContentMd5(v string) *DeleteMultipleObjectsHeaders
	GetContentMd5() *string
}

type DeleteMultipleObjectsHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ohhnqLBJFiKkPSBO1eNaUA==
	ContentMd5 *string `json:"content-md5,omitempty" xml:"content-md5,omitempty"`
}

func (s DeleteMultipleObjectsHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultipleObjectsHeaders) GoString() string {
	return s.String()
}

func (s *DeleteMultipleObjectsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteMultipleObjectsHeaders) GetContentMd5() *string {
	return s.ContentMd5
}

func (s *DeleteMultipleObjectsHeaders) SetCommonHeaders(v map[string]*string) *DeleteMultipleObjectsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteMultipleObjectsHeaders) SetContentMd5(v string) *DeleteMultipleObjectsHeaders {
	s.ContentMd5 = &v
	return s
}

func (s *DeleteMultipleObjectsHeaders) Validate() error {
	return dara.Validate(s)
}
