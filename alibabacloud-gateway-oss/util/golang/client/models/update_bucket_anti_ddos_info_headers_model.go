// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBucketAntiDDosInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateBucketAntiDDosInfoHeaders
	GetCommonHeaders() map[string]*string
	SetDefenderInstance(v string) *UpdateBucketAntiDDosInfoHeaders
	GetDefenderInstance() *string
	SetDefenderStatus(v string) *UpdateBucketAntiDDosInfoHeaders
	GetDefenderStatus() *string
}

type UpdateBucketAntiDDosInfoHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The Anti-DDoS instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cbcac8d2-4f75-4d6d-9f2e-c3447f73****
	DefenderInstance *string `json:"x-oss-defender-instance,omitempty" xml:"x-oss-defender-instance,omitempty"`
	// The new status of the Anti-DDoS instance. Valid values:
	//
	// 	- Init: You must specify the custom domain name that you want to protect.
	//
	// 	- Defending: You can select whether to specify the custom domain name that you want to protect.
	//
	// 	- HaltDefending: You do not need to specify the custom domain name that you want to protect.
	//
	// This parameter is required.
	//
	// example:
	//
	// Init
	DefenderStatus *string `json:"x-oss-defender-status,omitempty" xml:"x-oss-defender-status,omitempty"`
}

func (s UpdateBucketAntiDDosInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateBucketAntiDDosInfoHeaders) GoString() string {
	return s.String()
}

func (s *UpdateBucketAntiDDosInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateBucketAntiDDosInfoHeaders) GetDefenderInstance() *string {
	return s.DefenderInstance
}

func (s *UpdateBucketAntiDDosInfoHeaders) GetDefenderStatus() *string {
	return s.DefenderStatus
}

func (s *UpdateBucketAntiDDosInfoHeaders) SetCommonHeaders(v map[string]*string) *UpdateBucketAntiDDosInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateBucketAntiDDosInfoHeaders) SetDefenderInstance(v string) *UpdateBucketAntiDDosInfoHeaders {
	s.DefenderInstance = &v
	return s
}

func (s *UpdateBucketAntiDDosInfoHeaders) SetDefenderStatus(v string) *UpdateBucketAntiDDosInfoHeaders {
	s.DefenderStatus = &v
	return s
}

func (s *UpdateBucketAntiDDosInfoHeaders) Validate() error {
	return dara.Validate(s)
}
