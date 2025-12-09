// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectObjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSelectRequest(v *SelectRequest) *SelectObjectRequest
	GetSelectRequest() *SelectRequest
	SetXOssProcess(v string) *SelectObjectRequest
	GetXOssProcess() *string
}

type SelectObjectRequest struct {
	// Container for saving Select request.
	SelectRequest *SelectRequest `json:"SelectRequest,omitempty" xml:"SelectRequest,omitempty"`
	// If it is a CSV file, this value should be set to csv/select; if it is a JSON file, it should be set to json/select.
	//
	// This parameter is required.
	//
	// example:
	//
	// csv/select
	XOssProcess *string `json:"x-oss-process,omitempty" xml:"x-oss-process,omitempty"`
}

func (s SelectObjectRequest) String() string {
	return dara.Prettify(s)
}

func (s SelectObjectRequest) GoString() string {
	return s.String()
}

func (s *SelectObjectRequest) GetSelectRequest() *SelectRequest {
	return s.SelectRequest
}

func (s *SelectObjectRequest) GetXOssProcess() *string {
	return s.XOssProcess
}

func (s *SelectObjectRequest) SetSelectRequest(v *SelectRequest) *SelectObjectRequest {
	s.SelectRequest = v
	return s
}

func (s *SelectObjectRequest) SetXOssProcess(v string) *SelectObjectRequest {
	s.XOssProcess = &v
	return s
}

func (s *SelectObjectRequest) Validate() error {
	if s.SelectRequest != nil {
		if err := s.SelectRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}
