// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommonHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v []*CommonHeadersHeader) *CommonHeaders
	GetHeader() []*CommonHeadersHeader
}

type CommonHeaders struct {
	Header []*CommonHeadersHeader `json:"Header,omitempty" xml:"Header,omitempty" type:"Repeated"`
}

func (s CommonHeaders) String() string {
	return dara.Prettify(s)
}

func (s CommonHeaders) GoString() string {
	return s.String()
}

func (s *CommonHeaders) GetHeader() []*CommonHeadersHeader {
	return s.Header
}

func (s *CommonHeaders) SetHeader(v []*CommonHeadersHeader) *CommonHeaders {
	s.Header = v
	return s
}

func (s *CommonHeaders) Validate() error {
	if s.Header != nil {
		for _, item := range s.Header {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CommonHeadersHeader struct {
	// example:
	//
	// X-Content-Type-Options
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// nosniff
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CommonHeadersHeader) String() string {
	return dara.Prettify(s)
}

func (s CommonHeadersHeader) GoString() string {
	return s.String()
}

func (s *CommonHeadersHeader) GetKey() *string {
	return s.Key
}

func (s *CommonHeadersHeader) GetValue() *string {
	return s.Value
}

func (s *CommonHeadersHeader) SetKey(v string) *CommonHeadersHeader {
	s.Key = &v
	return s
}

func (s *CommonHeadersHeader) SetValue(v string) *CommonHeadersHeader {
	s.Value = &v
	return s
}

func (s *CommonHeadersHeader) Validate() error {
	return dara.Validate(s)
}
