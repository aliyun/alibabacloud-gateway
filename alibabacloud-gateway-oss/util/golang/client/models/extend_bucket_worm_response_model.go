// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExtendBucketWormResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExtendBucketWormResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExtendBucketWormResponse
  GetStatusCode() *int32 
}

type ExtendBucketWormResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s ExtendBucketWormResponse) String() string {
  return dara.Prettify(s)
}

func (s ExtendBucketWormResponse) GoString() string {
  return s.String()
}

func (s *ExtendBucketWormResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExtendBucketWormResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExtendBucketWormResponse) SetHeaders(v map[string]*string) *ExtendBucketWormResponse {
  s.Headers = v
  return s
}

func (s *ExtendBucketWormResponse) SetStatusCode(v int32) *ExtendBucketWormResponse {
  s.StatusCode = &v
  return s
}

func (s *ExtendBucketWormResponse) Validate() error {
  return dara.Validate(s)
}

