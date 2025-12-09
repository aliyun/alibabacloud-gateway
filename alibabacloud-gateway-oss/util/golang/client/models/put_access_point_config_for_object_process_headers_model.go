// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutAccessPointConfigForObjectProcessHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PutAccessPointConfigForObjectProcessHeaders
	GetCommonHeaders() map[string]*string
	SetXOssAccessPointForObjectProcessName(v string) *PutAccessPointConfigForObjectProcessHeaders
	GetXOssAccessPointForObjectProcessName() *string
}

type PutAccessPointConfigForObjectProcessHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The name of the Object FC Access Point. The name of an Object FC Access Point must meet the following requirements:
	//
	// The name cannot exceed 63 characters in length.
	//
	// The name can contain only lowercase letters, digits, and hyphens (-) and cannot start or end with a hyphen (-).
	//
	// The name must be unique in the current region.
	//
	// This parameter is required.
	//
	// example:
	//
	// fc-ap-01
	XOssAccessPointForObjectProcessName *string `json:"x-oss-access-point-for-object-process-name,omitempty" xml:"x-oss-access-point-for-object-process-name,omitempty"`
}

func (s PutAccessPointConfigForObjectProcessHeaders) String() string {
	return dara.Prettify(s)
}

func (s PutAccessPointConfigForObjectProcessHeaders) GoString() string {
	return s.String()
}

func (s *PutAccessPointConfigForObjectProcessHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PutAccessPointConfigForObjectProcessHeaders) GetXOssAccessPointForObjectProcessName() *string {
	return s.XOssAccessPointForObjectProcessName
}

func (s *PutAccessPointConfigForObjectProcessHeaders) SetCommonHeaders(v map[string]*string) *PutAccessPointConfigForObjectProcessHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PutAccessPointConfigForObjectProcessHeaders) SetXOssAccessPointForObjectProcessName(v string) *PutAccessPointConfigForObjectProcessHeaders {
	s.XOssAccessPointForObjectProcessName = &v
	return s
}

func (s *PutAccessPointConfigForObjectProcessHeaders) Validate() error {
	return dara.Validate(s)
}
