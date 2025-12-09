// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessPointForObjectProcessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGetAccessPointForObjectProcessResult(v *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult) *GetAccessPointForObjectProcessResponseBody
	GetGetAccessPointForObjectProcessResult() *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult
}

type GetAccessPointForObjectProcessResponseBody struct {
	// The container that stores information about the Object FC Access Point.
	GetAccessPointForObjectProcessResult *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult `json:"GetAccessPointForObjectProcessResult,omitempty" xml:"GetAccessPointForObjectProcessResult,omitempty" type:"Struct"`
}

func (s GetAccessPointForObjectProcessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccessPointForObjectProcessResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessPointForObjectProcessResponseBody) GetGetAccessPointForObjectProcessResult() *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult {
	return s.GetAccessPointForObjectProcessResult
}

func (s *GetAccessPointForObjectProcessResponseBody) SetGetAccessPointForObjectProcessResult(v *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult) *GetAccessPointForObjectProcessResponseBody {
	s.GetAccessPointForObjectProcessResult = v
	return s
}

func (s *GetAccessPointForObjectProcessResponseBody) Validate() error {
	if s.GetAccessPointForObjectProcessResult != nil {
		if err := s.GetAccessPointForObjectProcessResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult struct {
	// The alias of the Object FC Access Point.
	//
	// example:
	//
	// fc-ap-01-3b00521f653d2b3223680ec39dbbe2****-opapalias
	AccessPointForObjectProcessAlias *string `json:"AccessPointForObjectProcessAlias,omitempty" xml:"AccessPointForObjectProcessAlias,omitempty"`
	// The ARN of the Object FC Access Point.
	//
	// example:
	//
	// acs:oss:cn-qingdao:119335441657143:accesspointforobjectprocess/fc-ap-01
	AccessPointForObjectProcessArn *string `json:"AccessPointForObjectProcessArn,omitempty" xml:"AccessPointForObjectProcessArn,omitempty"`
	// The name of the access point.
	//
	// example:
	//
	// ap-01
	AccessPointName *string `json:"AccessPointName,omitempty" xml:"AccessPointName,omitempty"`
	// The name of the Object FC Access Point.
	//
	// example:
	//
	// fc-ap-01
	AccessPointNameForObjectProcess *string `json:"AccessPointNameForObjectProcess,omitempty" xml:"AccessPointNameForObjectProcess,omitempty"`
	// The UID of the Alibaba Cloud account to which the Object FC Access Point belongs.
	//
	// example:
	//
	// 111933544165****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// Whether allow anonymous users to access this FC Access Point.
	//
	// example:
	//
	// false
	AllowAnonymousAccessForObjectProcess *string `json:"AllowAnonymousAccessForObjectProcess,omitempty" xml:"AllowAnonymousAccessForObjectProcess,omitempty"`
	// The time when the Object FC Access Point was created. The value is a timestamp.
	//
	// example:
	//
	// 1626769503
	CreationDate *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	// The container that stores the endpoints of the Object FC Access Point.
	Endpoints *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResultEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Struct"`
	// The container in which the Block Public Access configurations are stored.
	PublicAccessBlockConfiguration *PublicAccessBlockConfiguration `json:"PublicAccessBlockConfiguration,omitempty" xml:"PublicAccessBlockConfiguration,omitempty"`
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

func (s GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult) String() string {
	return dara.Prettify(s)
}

func (s GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult) GoString() string {
	return s.String()
}

func (s *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult) GetAccessPointForObjectProcessAlias() *string {
	return s.AccessPointForObjectProcessAlias
}

func (s *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult) GetAccessPointForObjectProcessArn() *string {
	return s.AccessPointForObjectProcessArn
}

func (s *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult) GetAccessPointName() *string {
	return s.AccessPointName
}

func (s *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult) GetAccessPointNameForObjectProcess() *string {
	return s.AccessPointNameForObjectProcess
}

func (s *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult) GetAccountId() *string {
	return s.AccountId
}

func (s *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult) GetAllowAnonymousAccessForObjectProcess() *string {
	return s.AllowAnonymousAccessForObjectProcess
}

func (s *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult) GetCreationDate() *string {
	return s.CreationDate
}

func (s *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult) GetEndpoints() *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResultEndpoints {
	return s.Endpoints
}

func (s *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult) GetPublicAccessBlockConfiguration() *PublicAccessBlockConfiguration {
	return s.PublicAccessBlockConfiguration
}

func (s *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult) GetStatus() *string {
	return s.Status
}

func (s *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult) SetAccessPointForObjectProcessAlias(v string) *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult {
	s.AccessPointForObjectProcessAlias = &v
	return s
}

func (s *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult) SetAccessPointForObjectProcessArn(v string) *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult {
	s.AccessPointForObjectProcessArn = &v
	return s
}

func (s *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult) SetAccessPointName(v string) *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult {
	s.AccessPointName = &v
	return s
}

func (s *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult) SetAccessPointNameForObjectProcess(v string) *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult {
	s.AccessPointNameForObjectProcess = &v
	return s
}

func (s *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult) SetAccountId(v string) *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult {
	s.AccountId = &v
	return s
}

func (s *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult) SetAllowAnonymousAccessForObjectProcess(v string) *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult {
	s.AllowAnonymousAccessForObjectProcess = &v
	return s
}

func (s *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult) SetCreationDate(v string) *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult {
	s.CreationDate = &v
	return s
}

func (s *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult) SetEndpoints(v *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResultEndpoints) *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult {
	s.Endpoints = v
	return s
}

func (s *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult) SetPublicAccessBlockConfiguration(v *PublicAccessBlockConfiguration) *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult {
	s.PublicAccessBlockConfiguration = v
	return s
}

func (s *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult) SetStatus(v string) *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult {
	s.Status = &v
	return s
}

func (s *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResult) Validate() error {
	if s.Endpoints != nil {
		if err := s.Endpoints.Validate(); err != nil {
			return err
		}
	}
	if s.PublicAccessBlockConfiguration != nil {
		if err := s.PublicAccessBlockConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResultEndpoints struct {
	// The internal endpoint of the Object FC Access Point.
	//
	// example:
	//
	// fc-ap-01-111933544165****.oss-cn-qingdao-internal.oss-object-process.aliyuncs.com
	InternalEndpoint *string `json:"InternalEndpoint,omitempty" xml:"InternalEndpoint,omitempty"`
	// The public endpoint of the Object FC Access Point.
	//
	// example:
	//
	// fc-ap-01-111933544165****.oss-cn-qingdao.oss-object-process.aliyuncs.com
	PublicEndpoint *string `json:"PublicEndpoint,omitempty" xml:"PublicEndpoint,omitempty"`
}

func (s GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResultEndpoints) String() string {
	return dara.Prettify(s)
}

func (s GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResultEndpoints) GoString() string {
	return s.String()
}

func (s *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResultEndpoints) GetInternalEndpoint() *string {
	return s.InternalEndpoint
}

func (s *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResultEndpoints) GetPublicEndpoint() *string {
	return s.PublicEndpoint
}

func (s *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResultEndpoints) SetInternalEndpoint(v string) *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResultEndpoints {
	s.InternalEndpoint = &v
	return s
}

func (s *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResultEndpoints) SetPublicEndpoint(v string) *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResultEndpoints {
	s.PublicEndpoint = &v
	return s
}

func (s *GetAccessPointForObjectProcessResponseBodyGetAccessPointForObjectProcessResultEndpoints) Validate() error {
	return dara.Validate(s)
}
