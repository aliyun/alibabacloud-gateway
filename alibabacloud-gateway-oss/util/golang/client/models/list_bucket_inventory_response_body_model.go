// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBucketInventoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListInventoryConfigurationsResult(v *ListBucketInventoryResponseBodyListInventoryConfigurationsResult) *ListBucketInventoryResponseBody
	GetListInventoryConfigurationsResult() *ListBucketInventoryResponseBodyListInventoryConfigurationsResult
}

type ListBucketInventoryResponseBody struct {
	// The container that stores inventory configuration list.
	ListInventoryConfigurationsResult *ListBucketInventoryResponseBodyListInventoryConfigurationsResult `json:"ListInventoryConfigurationsResult,omitempty" xml:"ListInventoryConfigurationsResult,omitempty" type:"Struct"`
}

func (s ListBucketInventoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBucketInventoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListBucketInventoryResponseBody) GetListInventoryConfigurationsResult() *ListBucketInventoryResponseBodyListInventoryConfigurationsResult {
	return s.ListInventoryConfigurationsResult
}

func (s *ListBucketInventoryResponseBody) SetListInventoryConfigurationsResult(v *ListBucketInventoryResponseBodyListInventoryConfigurationsResult) *ListBucketInventoryResponseBody {
	s.ListInventoryConfigurationsResult = v
	return s
}

func (s *ListBucketInventoryResponseBody) Validate() error {
	if s.ListInventoryConfigurationsResult != nil {
		if err := s.ListInventoryConfigurationsResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListBucketInventoryResponseBodyListInventoryConfigurationsResult struct {
	// The container that stores inventory configurations.
	InventoryConfigurations []*InventoryConfiguration `json:"InventoryConfiguration,omitempty" xml:"InventoryConfiguration,omitempty" type:"Repeated"`
	// Specifies whether to list all inventory tasks configured for the bucket.
	//
	// Valid values: true and false
	//
	// - The value of false indicates that all inventory tasks configured for the bucket are listed.
	//
	// - The value of true indicates that not all inventory tasks configured for the bucket are listed. To list the next page of inventory configurations, set the continuation-token parameter in the next request to the value of the NextContinuationToken header in the response to the current request.
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// If the value of IsTruncated in the response is true and value of this header is not null, set the continuation-token parameter in the next request to the value of this header.
	//
	// example:
	//
	// DFSadfe**
	NextContinuationToken *string `json:"NextContinuationToken,omitempty" xml:"NextContinuationToken,omitempty"`
}

func (s ListBucketInventoryResponseBodyListInventoryConfigurationsResult) String() string {
	return dara.Prettify(s)
}

func (s ListBucketInventoryResponseBodyListInventoryConfigurationsResult) GoString() string {
	return s.String()
}

func (s *ListBucketInventoryResponseBodyListInventoryConfigurationsResult) GetInventoryConfigurations() []*InventoryConfiguration {
	return s.InventoryConfigurations
}

func (s *ListBucketInventoryResponseBodyListInventoryConfigurationsResult) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListBucketInventoryResponseBodyListInventoryConfigurationsResult) GetNextContinuationToken() *string {
	return s.NextContinuationToken
}

func (s *ListBucketInventoryResponseBodyListInventoryConfigurationsResult) SetInventoryConfigurations(v []*InventoryConfiguration) *ListBucketInventoryResponseBodyListInventoryConfigurationsResult {
	s.InventoryConfigurations = v
	return s
}

func (s *ListBucketInventoryResponseBodyListInventoryConfigurationsResult) SetIsTruncated(v bool) *ListBucketInventoryResponseBodyListInventoryConfigurationsResult {
	s.IsTruncated = &v
	return s
}

func (s *ListBucketInventoryResponseBodyListInventoryConfigurationsResult) SetNextContinuationToken(v string) *ListBucketInventoryResponseBodyListInventoryConfigurationsResult {
	s.NextContinuationToken = &v
	return s
}

func (s *ListBucketInventoryResponseBodyListInventoryConfigurationsResult) Validate() error {
	if s.InventoryConfigurations != nil {
		for _, item := range s.InventoryConfigurations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
