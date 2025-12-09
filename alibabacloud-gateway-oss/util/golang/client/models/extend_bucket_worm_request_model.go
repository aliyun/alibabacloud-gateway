// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExtendBucketWormRequest interface {
  dara.Model
  String() string
  GoString() string
  SetExtendWormConfiguration(v *ExtendWormConfiguration) *ExtendBucketWormRequest
  GetExtendWormConfiguration() *ExtendWormConfiguration 
  SetWormId(v string) *ExtendBucketWormRequest
  GetWormId() *string 
}

type ExtendBucketWormRequest struct {
  // The parameters for WORM extension.
  ExtendWormConfiguration *ExtendWormConfiguration `json:"ExtendWormConfiguration,omitempty" xml:"ExtendWormConfiguration,omitempty"`
  // The ID of the retention policy.
  // 
  // >  If the ID of the retention policy that specifies the number of days for which objects can be retained does not exist, the HTTP status code 404 is returned.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 1666E2CFB2B3418****
  WormId *string `json:"wormId,omitempty" xml:"wormId,omitempty"`
}

func (s ExtendBucketWormRequest) String() string {
  return dara.Prettify(s)
}

func (s ExtendBucketWormRequest) GoString() string {
  return s.String()
}

func (s *ExtendBucketWormRequest) GetExtendWormConfiguration() *ExtendWormConfiguration  {
  return s.ExtendWormConfiguration
}

func (s *ExtendBucketWormRequest) GetWormId() *string  {
  return s.WormId
}

func (s *ExtendBucketWormRequest) SetExtendWormConfiguration(v *ExtendWormConfiguration) *ExtendBucketWormRequest {
  s.ExtendWormConfiguration = v
  return s
}

func (s *ExtendBucketWormRequest) SetWormId(v string) *ExtendBucketWormRequest {
  s.WormId = &v
  return s
}

func (s *ExtendBucketWormRequest) Validate() error {
  if s.ExtendWormConfiguration != nil {
    if err := s.ExtendWormConfiguration.Validate(); err != nil {
      return err
    }
  }
  return nil
}

