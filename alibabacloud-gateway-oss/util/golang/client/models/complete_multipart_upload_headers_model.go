// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteMultipartUploadHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CompleteMultipartUploadHeaders
	GetCommonHeaders() map[string]*string
	SetCompleteAll(v string) *CompleteMultipartUploadHeaders
	GetCompleteAll() *string
	SetForbidOverwrite(v string) *CompleteMultipartUploadHeaders
	GetForbidOverwrite() *string
}

type CompleteMultipartUploadHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// Specifies whether to list all parts that are uploaded by using the current upload ID.
	//
	// Valid value: yes.
	//
	// - If x-oss-complete-all is set to yes in the request, OSS lists all parts that are uploaded by using the current upload ID, sorts the parts by part number, and then performs the CompleteMultipartUpload operation. When OSS performs the CompleteMultipartUpload operation, OSS cannot detect the parts that are not uploaded or currently being uploaded. Before you call the CompleteMultipartUpload operation, make sure that all parts are uploaded.
	//
	// - If x-oss-complete-all is specified in the request, the request body cannot be specified. Otherwise, an error occurs.
	//
	// - If x-oss-complete-all is specified in the request, the format of the response remains unchanged.
	//
	// example:
	//
	// yes
	CompleteAll *string `json:"x-oss-complete-all,omitempty" xml:"x-oss-complete-all,omitempty"`
	// Specifieswhethertheobjectwith the sameobjectname is overwritten when you call the CompleteMultipartUpload operation.
	//
	// - If the value of x-oss-forbid-overwrite is not specified or set to false, the existing object can be overwritten by the object that has the same name.
	//
	// - If the value of x-oss-forbid-overwrite is set to true, the existing object cannot be overwritten by the object that has the same name.
	//
	// - The x-oss-forbid-overwrite request header is invalid if versioning is enabled or suspended for the bucket. In this case, the existing object can be overwritten by the object that has the same name when you call the CompleteMultipartUpload operation.
	//
	// - If you specify the x-oss-forbid-overwrite request header, the queries per second (QPS) performance of OSS may be degraded. If you want to configure the x-oss-forbid-overwrite header in a large number of requests (QPS > 1,000), submit a ticket.
	//
	// example:
	//
	// true
	ForbidOverwrite *string `json:"x-oss-forbid-overwrite,omitempty" xml:"x-oss-forbid-overwrite,omitempty"`
}

func (s CompleteMultipartUploadHeaders) String() string {
	return dara.Prettify(s)
}

func (s CompleteMultipartUploadHeaders) GoString() string {
	return s.String()
}

func (s *CompleteMultipartUploadHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CompleteMultipartUploadHeaders) GetCompleteAll() *string {
	return s.CompleteAll
}

func (s *CompleteMultipartUploadHeaders) GetForbidOverwrite() *string {
	return s.ForbidOverwrite
}

func (s *CompleteMultipartUploadHeaders) SetCommonHeaders(v map[string]*string) *CompleteMultipartUploadHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CompleteMultipartUploadHeaders) SetCompleteAll(v string) *CompleteMultipartUploadHeaders {
	s.CompleteAll = &v
	return s
}

func (s *CompleteMultipartUploadHeaders) SetForbidOverwrite(v string) *CompleteMultipartUploadHeaders {
	s.ForbidOverwrite = &v
	return s
}

func (s *CompleteMultipartUploadHeaders) Validate() error {
	return dara.Validate(s)
}
