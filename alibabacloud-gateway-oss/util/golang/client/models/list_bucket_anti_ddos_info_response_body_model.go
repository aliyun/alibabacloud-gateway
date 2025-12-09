// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBucketAntiDDosInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAntiDDOSListConfiguration(v *ListBucketAntiDDosInfoResponseBodyAntiDDOSListConfiguration) *ListBucketAntiDDosInfoResponseBody
	GetAntiDDOSListConfiguration() *ListBucketAntiDDosInfoResponseBodyAntiDDOSListConfiguration
}

type ListBucketAntiDDosInfoResponseBody struct {
	// The container that stores the protection list of an Anti-DDoS instance of a bucket.
	AntiDDOSListConfiguration *ListBucketAntiDDosInfoResponseBodyAntiDDOSListConfiguration `json:"AntiDDOSListConfiguration,omitempty" xml:"AntiDDOSListConfiguration,omitempty" type:"Struct"`
}

func (s ListBucketAntiDDosInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBucketAntiDDosInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListBucketAntiDDosInfoResponseBody) GetAntiDDOSListConfiguration() *ListBucketAntiDDosInfoResponseBodyAntiDDOSListConfiguration {
	return s.AntiDDOSListConfiguration
}

func (s *ListBucketAntiDDosInfoResponseBody) SetAntiDDOSListConfiguration(v *ListBucketAntiDDosInfoResponseBodyAntiDDOSListConfiguration) *ListBucketAntiDDosInfoResponseBody {
	s.AntiDDOSListConfiguration = v
	return s
}

func (s *ListBucketAntiDDosInfoResponseBody) Validate() error {
	if s.AntiDDOSListConfiguration != nil {
		if err := s.AntiDDOSListConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListBucketAntiDDosInfoResponseBodyAntiDDOSListConfiguration struct {
	// The container that stores information about the Anti-DDoS instance.
	AntiDDOSConfiguration []*BucketAntiDDOSInfo `json:"AntiDDOSConfiguration,omitempty" xml:"AntiDDOSConfiguration,omitempty" type:"Repeated"`
	// Indicates whether all Anti-DDoS instances are returned.
	//
	// - true: All Anti-DDoS instances are returned.
	//
	// - false: Not all Anti-DDoS instances are returned.
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The Anti-DDoS instances whose names are alphabetically after the specified marker.
	//
	// example:
	//
	// abc
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
}

func (s ListBucketAntiDDosInfoResponseBodyAntiDDOSListConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ListBucketAntiDDosInfoResponseBodyAntiDDOSListConfiguration) GoString() string {
	return s.String()
}

func (s *ListBucketAntiDDosInfoResponseBodyAntiDDOSListConfiguration) GetAntiDDOSConfiguration() []*BucketAntiDDOSInfo {
	return s.AntiDDOSConfiguration
}

func (s *ListBucketAntiDDosInfoResponseBodyAntiDDOSListConfiguration) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListBucketAntiDDosInfoResponseBodyAntiDDOSListConfiguration) GetMarker() *string {
	return s.Marker
}

func (s *ListBucketAntiDDosInfoResponseBodyAntiDDOSListConfiguration) SetAntiDDOSConfiguration(v []*BucketAntiDDOSInfo) *ListBucketAntiDDosInfoResponseBodyAntiDDOSListConfiguration {
	s.AntiDDOSConfiguration = v
	return s
}

func (s *ListBucketAntiDDosInfoResponseBodyAntiDDOSListConfiguration) SetIsTruncated(v bool) *ListBucketAntiDDosInfoResponseBodyAntiDDOSListConfiguration {
	s.IsTruncated = &v
	return s
}

func (s *ListBucketAntiDDosInfoResponseBodyAntiDDOSListConfiguration) SetMarker(v string) *ListBucketAntiDDosInfoResponseBodyAntiDDOSListConfiguration {
	s.Marker = &v
	return s
}

func (s *ListBucketAntiDDosInfoResponseBodyAntiDDOSListConfiguration) Validate() error {
	if s.AntiDDOSConfiguration != nil {
		for _, item := range s.AntiDDOSConfiguration {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
