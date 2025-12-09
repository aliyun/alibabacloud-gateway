// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObjectHashConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayObjectHash(v bool) *ObjectHashConfiguration
	GetDisplayObjectHash() *bool
	SetObjectHashFunction(v string) *ObjectHashConfiguration
	GetObjectHashFunction() *string
}

type ObjectHashConfiguration struct {
	// example:
	//
	// true
	DisplayObjectHash *bool `json:"DisplayObjectHash,omitempty" xml:"DisplayObjectHash,omitempty"`
	// example:
	//
	// SHA-256
	ObjectHashFunction *string `json:"ObjectHashFunction,omitempty" xml:"ObjectHashFunction,omitempty"`
}

func (s ObjectHashConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ObjectHashConfiguration) GoString() string {
	return s.String()
}

func (s *ObjectHashConfiguration) GetDisplayObjectHash() *bool {
	return s.DisplayObjectHash
}

func (s *ObjectHashConfiguration) GetObjectHashFunction() *string {
	return s.ObjectHashFunction
}

func (s *ObjectHashConfiguration) SetDisplayObjectHash(v bool) *ObjectHashConfiguration {
	s.DisplayObjectHash = &v
	return s
}

func (s *ObjectHashConfiguration) SetObjectHashFunction(v string) *ObjectHashConfiguration {
	s.ObjectHashFunction = &v
	return s
}

func (s *ObjectHashConfiguration) Validate() error {
	return dara.Validate(s)
}
