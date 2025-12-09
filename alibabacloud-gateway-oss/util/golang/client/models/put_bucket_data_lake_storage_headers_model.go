// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutBucketDataLakeStorageHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PutBucketDataLakeStorageHeaders
	GetCommonHeaders() map[string]*string
	SetXOssDlsStatus(v string) *PutBucketDataLakeStorageHeaders
	GetXOssDlsStatus() *string
}

type PutBucketDataLakeStorageHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	XOssDlsStatus *string `json:"x-oss-dls-status,omitempty" xml:"x-oss-dls-status,omitempty"`
}

func (s PutBucketDataLakeStorageHeaders) String() string {
	return dara.Prettify(s)
}

func (s PutBucketDataLakeStorageHeaders) GoString() string {
	return s.String()
}

func (s *PutBucketDataLakeStorageHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PutBucketDataLakeStorageHeaders) GetXOssDlsStatus() *string {
	return s.XOssDlsStatus
}

func (s *PutBucketDataLakeStorageHeaders) SetCommonHeaders(v map[string]*string) *PutBucketDataLakeStorageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PutBucketDataLakeStorageHeaders) SetXOssDlsStatus(v string) *PutBucketDataLakeStorageHeaders {
	s.XOssDlsStatus = &v
	return s
}

func (s *PutBucketDataLakeStorageHeaders) Validate() error {
	return dara.Validate(s)
}
