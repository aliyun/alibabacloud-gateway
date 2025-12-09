// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessPointsResult interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPoints(v *ListAccessPointsResultAccessPoints) *ListAccessPointsResult
	GetAccessPoints() *ListAccessPointsResultAccessPoints
	SetAccountId(v string) *ListAccessPointsResult
	GetAccountId() *string
	SetIsTruncated(v string) *ListAccessPointsResult
	GetIsTruncated() *string
	SetMaxKeys(v int32) *ListAccessPointsResult
	GetMaxKeys() *int32
	SetNextContinuationToken(v string) *ListAccessPointsResult
	GetNextContinuationToken() *string
}

type ListAccessPointsResult struct {
	AccessPoints *ListAccessPointsResultAccessPoints `json:"AccessPoints,omitempty" xml:"AccessPoints,omitempty" type:"Struct"`
	// example:
	//
	// 11489****34218
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// true
	IsTruncated *string `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// example:
	//
	// 10
	MaxKeys *int32 `json:"MaxKeys,omitempty" xml:"MaxKeys,omitempty"`
	// example:
	//
	// ap-test-01
	NextContinuationToken *string `json:"NextContinuationToken,omitempty" xml:"NextContinuationToken,omitempty"`
}

func (s ListAccessPointsResult) String() string {
	return dara.Prettify(s)
}

func (s ListAccessPointsResult) GoString() string {
	return s.String()
}

func (s *ListAccessPointsResult) GetAccessPoints() *ListAccessPointsResultAccessPoints {
	return s.AccessPoints
}

func (s *ListAccessPointsResult) GetAccountId() *string {
	return s.AccountId
}

func (s *ListAccessPointsResult) GetIsTruncated() *string {
	return s.IsTruncated
}

func (s *ListAccessPointsResult) GetMaxKeys() *int32 {
	return s.MaxKeys
}

func (s *ListAccessPointsResult) GetNextContinuationToken() *string {
	return s.NextContinuationToken
}

func (s *ListAccessPointsResult) SetAccessPoints(v *ListAccessPointsResultAccessPoints) *ListAccessPointsResult {
	s.AccessPoints = v
	return s
}

func (s *ListAccessPointsResult) SetAccountId(v string) *ListAccessPointsResult {
	s.AccountId = &v
	return s
}

func (s *ListAccessPointsResult) SetIsTruncated(v string) *ListAccessPointsResult {
	s.IsTruncated = &v
	return s
}

func (s *ListAccessPointsResult) SetMaxKeys(v int32) *ListAccessPointsResult {
	s.MaxKeys = &v
	return s
}

func (s *ListAccessPointsResult) SetNextContinuationToken(v string) *ListAccessPointsResult {
	s.NextContinuationToken = &v
	return s
}

func (s *ListAccessPointsResult) Validate() error {
	if s.AccessPoints != nil {
		if err := s.AccessPoints.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAccessPointsResultAccessPoints struct {
	AccessPoint []*AccessPoint `json:"AccessPoint,omitempty" xml:"AccessPoint,omitempty" type:"Repeated"`
}

func (s ListAccessPointsResultAccessPoints) String() string {
	return dara.Prettify(s)
}

func (s ListAccessPointsResultAccessPoints) GoString() string {
	return s.String()
}

func (s *ListAccessPointsResultAccessPoints) GetAccessPoint() []*AccessPoint {
	return s.AccessPoint
}

func (s *ListAccessPointsResultAccessPoints) SetAccessPoint(v []*AccessPoint) *ListAccessPointsResultAccessPoints {
	s.AccessPoint = v
	return s
}

func (s *ListAccessPointsResultAccessPoints) Validate() error {
	if s.AccessPoint != nil {
		for _, item := range s.AccessPoint {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
