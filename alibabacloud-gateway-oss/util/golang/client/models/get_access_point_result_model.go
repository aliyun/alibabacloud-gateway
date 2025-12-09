// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessPointResult interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPointArn(v string) *GetAccessPointResult
	GetAccessPointArn() *string
	SetAccessPointName(v string) *GetAccessPointResult
	GetAccessPointName() *string
	SetAccountId(v string) *GetAccessPointResult
	GetAccountId() *string
	SetAlias(v string) *GetAccessPointResult
	GetAlias() *string
	SetBucket(v string) *GetAccessPointResult
	GetBucket() *string
	SetCreationDate(v string) *GetAccessPointResult
	GetCreationDate() *string
	SetEndpoints(v *GetAccessPointResultEndpoints) *GetAccessPointResult
	GetEndpoints() *GetAccessPointResultEndpoints
	SetNetworkOrigin(v string) *GetAccessPointResult
	GetNetworkOrigin() *string
	SetPublicAccessBlockConfiguration(v *PublicAccessBlockConfiguration) *GetAccessPointResult
	GetPublicAccessBlockConfiguration() *PublicAccessBlockConfiguration
	SetStatus(v string) *GetAccessPointResult
	GetStatus() *string
	SetVpcConfiguration(v *AccessPointVpcConfiguration) *GetAccessPointResult
	GetVpcConfiguration() *AccessPointVpcConfiguration
}

type GetAccessPointResult struct {
	AccessPointArn  *string `json:"AccessPointArn,omitempty" xml:"AccessPointArn,omitempty"`
	AccessPointName *string `json:"AccessPointName,omitempty" xml:"AccessPointName,omitempty"`
	// example:
	//
	// 114******818
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	Alias     *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// example:
	//
	// example-bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// example:
	//
	// Sat, 27 Apr 2024 15:04:14 GMT
	CreationDate                   *string                         `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	Endpoints                      *GetAccessPointResultEndpoints  `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Struct"`
	NetworkOrigin                  *string                         `json:"NetworkOrigin,omitempty" xml:"NetworkOrigin,omitempty"`
	PublicAccessBlockConfiguration *PublicAccessBlockConfiguration `json:"PublicAccessBlockConfiguration,omitempty" xml:"PublicAccessBlockConfiguration,omitempty"`
	Status                         *string                         `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcConfiguration               *AccessPointVpcConfiguration    `json:"VpcConfiguration,omitempty" xml:"VpcConfiguration,omitempty"`
}

func (s GetAccessPointResult) String() string {
	return dara.Prettify(s)
}

func (s GetAccessPointResult) GoString() string {
	return s.String()
}

func (s *GetAccessPointResult) GetAccessPointArn() *string {
	return s.AccessPointArn
}

func (s *GetAccessPointResult) GetAccessPointName() *string {
	return s.AccessPointName
}

func (s *GetAccessPointResult) GetAccountId() *string {
	return s.AccountId
}

func (s *GetAccessPointResult) GetAlias() *string {
	return s.Alias
}

func (s *GetAccessPointResult) GetBucket() *string {
	return s.Bucket
}

func (s *GetAccessPointResult) GetCreationDate() *string {
	return s.CreationDate
}

func (s *GetAccessPointResult) GetEndpoints() *GetAccessPointResultEndpoints {
	return s.Endpoints
}

func (s *GetAccessPointResult) GetNetworkOrigin() *string {
	return s.NetworkOrigin
}

func (s *GetAccessPointResult) GetPublicAccessBlockConfiguration() *PublicAccessBlockConfiguration {
	return s.PublicAccessBlockConfiguration
}

func (s *GetAccessPointResult) GetStatus() *string {
	return s.Status
}

func (s *GetAccessPointResult) GetVpcConfiguration() *AccessPointVpcConfiguration {
	return s.VpcConfiguration
}

func (s *GetAccessPointResult) SetAccessPointArn(v string) *GetAccessPointResult {
	s.AccessPointArn = &v
	return s
}

func (s *GetAccessPointResult) SetAccessPointName(v string) *GetAccessPointResult {
	s.AccessPointName = &v
	return s
}

func (s *GetAccessPointResult) SetAccountId(v string) *GetAccessPointResult {
	s.AccountId = &v
	return s
}

func (s *GetAccessPointResult) SetAlias(v string) *GetAccessPointResult {
	s.Alias = &v
	return s
}

func (s *GetAccessPointResult) SetBucket(v string) *GetAccessPointResult {
	s.Bucket = &v
	return s
}

func (s *GetAccessPointResult) SetCreationDate(v string) *GetAccessPointResult {
	s.CreationDate = &v
	return s
}

func (s *GetAccessPointResult) SetEndpoints(v *GetAccessPointResultEndpoints) *GetAccessPointResult {
	s.Endpoints = v
	return s
}

func (s *GetAccessPointResult) SetNetworkOrigin(v string) *GetAccessPointResult {
	s.NetworkOrigin = &v
	return s
}

func (s *GetAccessPointResult) SetPublicAccessBlockConfiguration(v *PublicAccessBlockConfiguration) *GetAccessPointResult {
	s.PublicAccessBlockConfiguration = v
	return s
}

func (s *GetAccessPointResult) SetStatus(v string) *GetAccessPointResult {
	s.Status = &v
	return s
}

func (s *GetAccessPointResult) SetVpcConfiguration(v *AccessPointVpcConfiguration) *GetAccessPointResult {
	s.VpcConfiguration = v
	return s
}

func (s *GetAccessPointResult) Validate() error {
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
	if s.VpcConfiguration != nil {
		if err := s.VpcConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAccessPointResultEndpoints struct {
	InternalEndpoint *string `json:"InternalEndpoint,omitempty" xml:"InternalEndpoint,omitempty"`
	PublicEndpoint   *string `json:"PublicEndpoint,omitempty" xml:"PublicEndpoint,omitempty"`
}

func (s GetAccessPointResultEndpoints) String() string {
	return dara.Prettify(s)
}

func (s GetAccessPointResultEndpoints) GoString() string {
	return s.String()
}

func (s *GetAccessPointResultEndpoints) GetInternalEndpoint() *string {
	return s.InternalEndpoint
}

func (s *GetAccessPointResultEndpoints) GetPublicEndpoint() *string {
	return s.PublicEndpoint
}

func (s *GetAccessPointResultEndpoints) SetInternalEndpoint(v string) *GetAccessPointResultEndpoints {
	s.InternalEndpoint = &v
	return s
}

func (s *GetAccessPointResultEndpoints) SetPublicEndpoint(v string) *GetAccessPointResultEndpoints {
	s.PublicEndpoint = &v
	return s
}

func (s *GetAccessPointResultEndpoints) Validate() error {
	return dara.Validate(s)
}
