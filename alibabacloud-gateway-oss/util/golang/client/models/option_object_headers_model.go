// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOptionObjectHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *OptionObjectHeaders
	GetCommonHeaders() map[string]*string
	SetAccessControlRequestHeaders(v string) *OptionObjectHeaders
	GetAccessControlRequestHeaders() *string
	SetAccessControlRequestMethod(v string) *OptionObjectHeaders
	GetAccessControlRequestMethod() *string
	SetOrigin(v string) *OptionObjectHeaders
	GetOrigin() *string
}

type OptionObjectHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The custom headers to be sent in the actual cross-origin request. You can configure multiple custom headers in a cross-origin request. Custom headers are separated by commas (,). By default, this header is left empty.
	//
	// example:
	//
	// x-oss-test1,x-oss-test2
	AccessControlRequestHeaders *string `json:"Access-Control-Request-Headers,omitempty" xml:"Access-Control-Request-Headers,omitempty"`
	// The method to be used in the actual cross-origin request. You can specify only one Access-Control-Request-Method header in a cross-origin request. By default, this header is left empty.
	//
	// example:
	//
	// PUT
	AccessControlRequestMethod *string `json:"Access-Control-Request-Method,omitempty" xml:"Access-Control-Request-Method,omitempty"`
	// The origin of the request. It is used to identify a cross-origin request. You can specify only one Origin header in a cross-origin request. By default, this header is left empty.
	//
	// example:
	//
	// http://www.example.com
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
}

func (s OptionObjectHeaders) String() string {
	return dara.Prettify(s)
}

func (s OptionObjectHeaders) GoString() string {
	return s.String()
}

func (s *OptionObjectHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *OptionObjectHeaders) GetAccessControlRequestHeaders() *string {
	return s.AccessControlRequestHeaders
}

func (s *OptionObjectHeaders) GetAccessControlRequestMethod() *string {
	return s.AccessControlRequestMethod
}

func (s *OptionObjectHeaders) GetOrigin() *string {
	return s.Origin
}

func (s *OptionObjectHeaders) SetCommonHeaders(v map[string]*string) *OptionObjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OptionObjectHeaders) SetAccessControlRequestHeaders(v string) *OptionObjectHeaders {
	s.AccessControlRequestHeaders = &v
	return s
}

func (s *OptionObjectHeaders) SetAccessControlRequestMethod(v string) *OptionObjectHeaders {
	s.AccessControlRequestMethod = &v
	return s
}

func (s *OptionObjectHeaders) SetOrigin(v string) *OptionObjectHeaders {
	s.Origin = &v
	return s
}

func (s *OptionObjectHeaders) Validate() error {
	return dara.Validate(s)
}
