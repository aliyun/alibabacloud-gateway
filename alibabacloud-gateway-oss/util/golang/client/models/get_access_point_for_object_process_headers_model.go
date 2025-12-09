// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessPointForObjectProcessHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetAccessPointForObjectProcessHeaders
	GetCommonHeaders() map[string]*string
	SetXOssAccessPointForObjectProcessName(v string) *GetAccessPointForObjectProcessHeaders
	GetXOssAccessPointForObjectProcessName() *string
}

type GetAccessPointForObjectProcessHeaders struct {
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

func (s GetAccessPointForObjectProcessHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetAccessPointForObjectProcessHeaders) GoString() string {
	return s.String()
}

func (s *GetAccessPointForObjectProcessHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetAccessPointForObjectProcessHeaders) GetXOssAccessPointForObjectProcessName() *string {
	return s.XOssAccessPointForObjectProcessName
}

func (s *GetAccessPointForObjectProcessHeaders) SetCommonHeaders(v map[string]*string) *GetAccessPointForObjectProcessHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAccessPointForObjectProcessHeaders) SetXOssAccessPointForObjectProcessName(v string) *GetAccessPointForObjectProcessHeaders {
	s.XOssAccessPointForObjectProcessName = &v
	return s
}

func (s *GetAccessPointForObjectProcessHeaders) Validate() error {
	return dara.Validate(s)
}
