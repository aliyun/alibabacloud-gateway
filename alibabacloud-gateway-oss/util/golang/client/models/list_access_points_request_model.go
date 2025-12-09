// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessPointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContinuationToken(v string) *ListAccessPointsRequest
	GetContinuationToken() *string
	SetMaxKeys(v int64) *ListAccessPointsRequest
	GetMaxKeys() *int64
}

type ListAccessPointsRequest struct {
	// The token from which the listing operation starts. You must specify the value of NextContinuationToken that is obtained from the previous query as the value of continuation-token.
	//
	// example:
	//
	// abc
	ContinuationToken *string `json:"continuation-token,omitempty" xml:"continuation-token,omitempty"`
	// The maximum number of access points that can be returned. Valid values:
	//
	// 	- For user-level access points: (0,1000].
	//
	// 	- For bucket-level access points: (0,100].
	//
	// example:
	//
	// 10
	MaxKeys *int64 `json:"max-keys,omitempty" xml:"max-keys,omitempty"`
}

func (s ListAccessPointsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAccessPointsRequest) GoString() string {
	return s.String()
}

func (s *ListAccessPointsRequest) GetContinuationToken() *string {
	return s.ContinuationToken
}

func (s *ListAccessPointsRequest) GetMaxKeys() *int64 {
	return s.MaxKeys
}

func (s *ListAccessPointsRequest) SetContinuationToken(v string) *ListAccessPointsRequest {
	s.ContinuationToken = &v
	return s
}

func (s *ListAccessPointsRequest) SetMaxKeys(v int64) *ListAccessPointsRequest {
	s.MaxKeys = &v
	return s
}

func (s *ListAccessPointsRequest) Validate() error {
	return dara.Validate(s)
}
