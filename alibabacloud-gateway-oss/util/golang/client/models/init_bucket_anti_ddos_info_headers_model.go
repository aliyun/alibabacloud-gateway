// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitBucketAntiDDosInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InitBucketAntiDDosInfoHeaders
	GetCommonHeaders() map[string]*string
	SetDefenderInstance(v string) *InitBucketAntiDDosInfoHeaders
	GetDefenderInstance() *string
	SetDefenderType(v string) *InitBucketAntiDDosInfoHeaders
	GetDefenderType() *string
}

type InitBucketAntiDDosInfoHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The ID of the Anti-DDoS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cbcac8d2-4f75-4d6d-9f2e-c3447f73****
	DefenderInstance *string `json:"x-oss-defender-instance,omitempty" xml:"x-oss-defender-instance,omitempty"`
	// The type of the Anti-DDoS instance. Set the value to AntiDDosPremimum.
	//
	// This parameter is required.
	//
	// example:
	//
	// AntiDDosPremimum
	DefenderType *string `json:"x-oss-defender-type,omitempty" xml:"x-oss-defender-type,omitempty"`
}

func (s InitBucketAntiDDosInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s InitBucketAntiDDosInfoHeaders) GoString() string {
	return s.String()
}

func (s *InitBucketAntiDDosInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InitBucketAntiDDosInfoHeaders) GetDefenderInstance() *string {
	return s.DefenderInstance
}

func (s *InitBucketAntiDDosInfoHeaders) GetDefenderType() *string {
	return s.DefenderType
}

func (s *InitBucketAntiDDosInfoHeaders) SetCommonHeaders(v map[string]*string) *InitBucketAntiDDosInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InitBucketAntiDDosInfoHeaders) SetDefenderInstance(v string) *InitBucketAntiDDosInfoHeaders {
	s.DefenderInstance = &v
	return s
}

func (s *InitBucketAntiDDosInfoHeaders) SetDefenderType(v string) *InitBucketAntiDDosInfoHeaders {
	s.DefenderType = &v
	return s
}

func (s *InitBucketAntiDDosInfoHeaders) Validate() error {
	return dara.Validate(s)
}
