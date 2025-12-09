// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserAntiDDosInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateUserAntiDDosInfoHeaders
	GetCommonHeaders() map[string]*string
	SetDefenderInstance(v string) *UpdateUserAntiDDosInfoHeaders
	GetDefenderInstance() *string
	SetDefenderStatus(v string) *UpdateUserAntiDDosInfoHeaders
	GetDefenderStatus() *string
}

type UpdateUserAntiDDosInfoHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The Anti-DDoS instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cbcac8d2-4f75-4d6d-9f2e-c3447f73****
	DefenderInstance *string `json:"x-oss-defender-instance,omitempty" xml:"x-oss-defender-instance,omitempty"`
	// The new status of the Anti-DDoS instance. Set the value to HaltDefending, which indicates that the Anti-DDos protection is disabled for a bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// HaltDefending
	DefenderStatus *string `json:"x-oss-defender-status,omitempty" xml:"x-oss-defender-status,omitempty"`
}

func (s UpdateUserAntiDDosInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserAntiDDosInfoHeaders) GoString() string {
	return s.String()
}

func (s *UpdateUserAntiDDosInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateUserAntiDDosInfoHeaders) GetDefenderInstance() *string {
	return s.DefenderInstance
}

func (s *UpdateUserAntiDDosInfoHeaders) GetDefenderStatus() *string {
	return s.DefenderStatus
}

func (s *UpdateUserAntiDDosInfoHeaders) SetCommonHeaders(v map[string]*string) *UpdateUserAntiDDosInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateUserAntiDDosInfoHeaders) SetDefenderInstance(v string) *UpdateUserAntiDDosInfoHeaders {
	s.DefenderInstance = &v
	return s
}

func (s *UpdateUserAntiDDosInfoHeaders) SetDefenderStatus(v string) *UpdateUserAntiDDosInfoHeaders {
	s.DefenderStatus = &v
	return s
}

func (s *UpdateUserAntiDDosInfoHeaders) Validate() error {
	return dara.Validate(s)
}
