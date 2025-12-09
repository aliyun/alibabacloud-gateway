// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessPointsForObjectProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContinuationToken(v string) *ListAccessPointsForObjectProcessRequest
	GetContinuationToken() *string
	SetMaxKeys(v int64) *ListAccessPointsForObjectProcessRequest
	GetMaxKeys() *int64
}

type ListAccessPointsForObjectProcessRequest struct {
	// The token from which the list operation must start. You can obtain this token from the NextContinuationToken element in the returned result.
	//
	// example:
	//
	// abc
	ContinuationToken *string `json:"continuation-token,omitempty" xml:"continuation-token,omitempty"`
	// The maximum number of Object FC Access Points to return.
	//
	// Valid values: 1 to 1000
	//
	// > If the list cannot be complete at a time due to the configurations of the max-keys element, the NextContinuationToken element is included in the response as the token for the next list.
	//
	// example:
	//
	// 10
	MaxKeys *int64 `json:"max-keys,omitempty" xml:"max-keys,omitempty"`
}

func (s ListAccessPointsForObjectProcessRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAccessPointsForObjectProcessRequest) GoString() string {
	return s.String()
}

func (s *ListAccessPointsForObjectProcessRequest) GetContinuationToken() *string {
	return s.ContinuationToken
}

func (s *ListAccessPointsForObjectProcessRequest) GetMaxKeys() *int64 {
	return s.MaxKeys
}

func (s *ListAccessPointsForObjectProcessRequest) SetContinuationToken(v string) *ListAccessPointsForObjectProcessRequest {
	s.ContinuationToken = &v
	return s
}

func (s *ListAccessPointsForObjectProcessRequest) SetMaxKeys(v int64) *ListAccessPointsForObjectProcessRequest {
	s.MaxKeys = &v
	return s
}

func (s *ListAccessPointsForObjectProcessRequest) Validate() error {
	return dara.Validate(s)
}
