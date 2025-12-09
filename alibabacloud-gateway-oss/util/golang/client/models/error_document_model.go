// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iErrorDocument interface {
  dara.Model
  String() string
  GoString() string
  SetHttpStatus(v int64) *ErrorDocument
  GetHttpStatus() *int64 
  SetKey(v string) *ErrorDocument
  GetKey() *string 
}

type ErrorDocument struct {
  // example:
  // 
  // 200
  HttpStatus *int64 `json:"HttpStatus,omitempty" xml:"HttpStatus,omitempty"`
  // example:
  // 
  // error.html
  Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s ErrorDocument) String() string {
  return dara.Prettify(s)
}

func (s ErrorDocument) GoString() string {
  return s.String()
}

func (s *ErrorDocument) GetHttpStatus() *int64  {
  return s.HttpStatus
}

func (s *ErrorDocument) GetKey() *string  {
  return s.Key
}

func (s *ErrorDocument) SetHttpStatus(v int64) *ErrorDocument {
  s.HttpStatus = &v
  return s
}

func (s *ErrorDocument) SetKey(v string) *ErrorDocument {
  s.Key = &v
  return s
}

func (s *ErrorDocument) Validate() error {
  return dara.Validate(s)
}

