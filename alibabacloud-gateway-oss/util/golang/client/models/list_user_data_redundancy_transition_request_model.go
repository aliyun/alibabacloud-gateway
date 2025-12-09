// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserDataRedundancyTransitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContinuationToken(v string) *ListUserDataRedundancyTransitionRequest
	GetContinuationToken() *string
	SetMaxKeys(v int32) *ListUserDataRedundancyTransitionRequest
	GetMaxKeys() *int32
}

type ListUserDataRedundancyTransitionRequest struct {
	// Specifies that the List operation should start from this token.
	//
	// example:
	//
	// abc
	ContinuationToken *string `json:"continuation-token,omitempty" xml:"continuation-token,omitempty"`
	// Limits the maximum number of tasks returned in this request. Value range: 1-100.
	//
	// example:
	//
	// 10
	MaxKeys *int32 `json:"max-keys,omitempty" xml:"max-keys,omitempty"`
}

func (s ListUserDataRedundancyTransitionRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserDataRedundancyTransitionRequest) GoString() string {
	return s.String()
}

func (s *ListUserDataRedundancyTransitionRequest) GetContinuationToken() *string {
	return s.ContinuationToken
}

func (s *ListUserDataRedundancyTransitionRequest) GetMaxKeys() *int32 {
	return s.MaxKeys
}

func (s *ListUserDataRedundancyTransitionRequest) SetContinuationToken(v string) *ListUserDataRedundancyTransitionRequest {
	s.ContinuationToken = &v
	return s
}

func (s *ListUserDataRedundancyTransitionRequest) SetMaxKeys(v int32) *ListUserDataRedundancyTransitionRequest {
	s.MaxKeys = &v
	return s
}

func (s *ListUserDataRedundancyTransitionRequest) Validate() error {
	return dara.Validate(s)
}
