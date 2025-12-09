// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessPointsForObjectProcessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListAccessPointsForObjectProcessResult(v *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResult) *ListAccessPointsForObjectProcessResponseBody
	GetListAccessPointsForObjectProcessResult() *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResult
}

type ListAccessPointsForObjectProcessResponseBody struct {
	// The container that stores information about the Object FC Access Points that are returned.
	ListAccessPointsForObjectProcessResult *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResult `json:"ListAccessPointsForObjectProcessResult,omitempty" xml:"ListAccessPointsForObjectProcessResult,omitempty" type:"Struct"`
}

func (s ListAccessPointsForObjectProcessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAccessPointsForObjectProcessResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccessPointsForObjectProcessResponseBody) GetListAccessPointsForObjectProcessResult() *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResult {
	return s.ListAccessPointsForObjectProcessResult
}

func (s *ListAccessPointsForObjectProcessResponseBody) SetListAccessPointsForObjectProcessResult(v *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResult) *ListAccessPointsForObjectProcessResponseBody {
	s.ListAccessPointsForObjectProcessResult = v
	return s
}

func (s *ListAccessPointsForObjectProcessResponseBody) Validate() error {
	if s.ListAccessPointsForObjectProcessResult != nil {
		if err := s.ListAccessPointsForObjectProcessResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResult struct {
	// The container that stores information about all Object FC Access Points.
	AccessPointsForObjectProcess *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcess `json:"AccessPointsForObjectProcess,omitempty" xml:"AccessPointsForObjectProcess,omitempty" type:"Struct"`
	// The UID of the Alibaba Cloud account to which the Object FC Access Points belong.
	//
	// example:
	//
	// 111933544165****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// Indicates whether the returned results are truncated. Valid values:
	//
	// true: indicates that not all results are returned for the request.
	//
	// false: indicates that all results are returned for the request.
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// Indicates that this ListAccessPointsForObjectProcess request contains subsequent results. You need to set the NextContinuationToken element to continuation-token for subsequent results.
	//
	// example:
	//
	// abc
	NextContinuationToken *string `json:"NextContinuationToken,omitempty" xml:"NextContinuationToken,omitempty"`
}

func (s ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResult) String() string {
	return dara.Prettify(s)
}

func (s ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResult) GoString() string {
	return s.String()
}

func (s *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResult) GetAccessPointsForObjectProcess() *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcess {
	return s.AccessPointsForObjectProcess
}

func (s *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResult) GetAccountId() *string {
	return s.AccountId
}

func (s *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResult) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResult) GetNextContinuationToken() *string {
	return s.NextContinuationToken
}

func (s *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResult) SetAccessPointsForObjectProcess(v *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcess) *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResult {
	s.AccessPointsForObjectProcess = v
	return s
}

func (s *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResult) SetAccountId(v string) *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResult {
	s.AccountId = &v
	return s
}

func (s *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResult) SetIsTruncated(v bool) *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResult {
	s.IsTruncated = &v
	return s
}

func (s *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResult) SetNextContinuationToken(v string) *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResult {
	s.NextContinuationToken = &v
	return s
}

func (s *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResult) Validate() error {
	if s.AccessPointsForObjectProcess != nil {
		if err := s.AccessPointsForObjectProcess.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcess struct {
	// The container that stores information about a single Object FC Access Point.
	AccessPointForObjectProcess []*ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcessAccessPointForObjectProcess `json:"AccessPointForObjectProcess,omitempty" xml:"AccessPointForObjectProcess,omitempty" type:"Repeated"`
}

func (s ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcess) String() string {
	return dara.Prettify(s)
}

func (s ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcess) GoString() string {
	return s.String()
}

func (s *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcess) GetAccessPointForObjectProcess() []*ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcessAccessPointForObjectProcess {
	return s.AccessPointForObjectProcess
}

func (s *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcess) SetAccessPointForObjectProcess(v []*ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcessAccessPointForObjectProcess) *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcess {
	s.AccessPointForObjectProcess = v
	return s
}

func (s *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcess) Validate() error {
	if s.AccessPointForObjectProcess != nil {
		for _, item := range s.AccessPointForObjectProcess {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcessAccessPointForObjectProcess struct {
	// The alias of the Object FC Access Point.
	//
	// example:
	//
	// fc-ap-01-3b00521f653d2b3223680ec39dbbe2****-opapalias
	AccessPointForObjectProcessAlias *string `json:"AccessPointForObjectProcessAlias,omitempty" xml:"AccessPointForObjectProcessAlias,omitempty"`
	// The name of the access point.
	//
	// example:
	//
	// fc-01
	AccessPointName *string `json:"AccessPointName,omitempty" xml:"AccessPointName,omitempty"`
	// The name of the Object FC Access Point.
	//
	// example:
	//
	// fc-ap-01
	AccessPointNameForObjectProcess *string `json:"AccessPointNameForObjectProcess,omitempty" xml:"AccessPointNameForObjectProcess,omitempty"`
	// Whether allow anonymous user access this FC Access Point.
	//
	// example:
	//
	// false
	AllowAnonymousAccessForObjectProcess *string `json:"AllowAnonymousAccessForObjectProcess,omitempty" xml:"AllowAnonymousAccessForObjectProcess,omitempty"`
	// The status of the Object FC Access Point. Valid values:
	//
	// enable: The Object FC Access Point is created.
	//
	// disable: The Object FC Access Point is disabled.
	//
	// creating: The Object FC Access Point is being created.
	//
	// deleting: The Object FC Access Point is deleted.
	//
	// example:
	//
	// enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcessAccessPointForObjectProcess) String() string {
	return dara.Prettify(s)
}

func (s ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcessAccessPointForObjectProcess) GoString() string {
	return s.String()
}

func (s *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcessAccessPointForObjectProcess) GetAccessPointForObjectProcessAlias() *string {
	return s.AccessPointForObjectProcessAlias
}

func (s *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcessAccessPointForObjectProcess) GetAccessPointName() *string {
	return s.AccessPointName
}

func (s *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcessAccessPointForObjectProcess) GetAccessPointNameForObjectProcess() *string {
	return s.AccessPointNameForObjectProcess
}

func (s *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcessAccessPointForObjectProcess) GetAllowAnonymousAccessForObjectProcess() *string {
	return s.AllowAnonymousAccessForObjectProcess
}

func (s *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcessAccessPointForObjectProcess) GetStatus() *string {
	return s.Status
}

func (s *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcessAccessPointForObjectProcess) SetAccessPointForObjectProcessAlias(v string) *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcessAccessPointForObjectProcess {
	s.AccessPointForObjectProcessAlias = &v
	return s
}

func (s *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcessAccessPointForObjectProcess) SetAccessPointName(v string) *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcessAccessPointForObjectProcess {
	s.AccessPointName = &v
	return s
}

func (s *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcessAccessPointForObjectProcess) SetAccessPointNameForObjectProcess(v string) *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcessAccessPointForObjectProcess {
	s.AccessPointNameForObjectProcess = &v
	return s
}

func (s *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcessAccessPointForObjectProcess) SetAllowAnonymousAccessForObjectProcess(v string) *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcessAccessPointForObjectProcess {
	s.AllowAnonymousAccessForObjectProcess = &v
	return s
}

func (s *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcessAccessPointForObjectProcess) SetStatus(v string) *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcessAccessPointForObjectProcess {
	s.Status = &v
	return s
}

func (s *ListAccessPointsForObjectProcessResponseBodyListAccessPointsForObjectProcessResultAccessPointsForObjectProcessAccessPointForObjectProcess) Validate() error {
	return dara.Validate(s)
}
