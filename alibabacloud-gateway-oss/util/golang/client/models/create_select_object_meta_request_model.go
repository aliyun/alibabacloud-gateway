// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSelectObjectMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCsvMetaRequest(v *SelectMetaRequest) *CreateSelectObjectMetaRequest
	GetCsvMetaRequest() *SelectMetaRequest
	SetXOssProcess(v string) *CreateSelectObjectMetaRequest
	GetXOssProcess() *string
}

type CreateSelectObjectMetaRequest struct {
	// The request body used to create select meta.
	CsvMetaRequest *SelectMetaRequest `json:"CsvMetaRequest,omitempty" xml:"CsvMetaRequest,omitempty"`
	// Parameters to specify the file formate.
	//
	// This parameter is required.
	//
	// example:
	//
	// csv/meta
	XOssProcess *string `json:"x-oss-process,omitempty" xml:"x-oss-process,omitempty"`
}

func (s CreateSelectObjectMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSelectObjectMetaRequest) GoString() string {
	return s.String()
}

func (s *CreateSelectObjectMetaRequest) GetCsvMetaRequest() *SelectMetaRequest {
	return s.CsvMetaRequest
}

func (s *CreateSelectObjectMetaRequest) GetXOssProcess() *string {
	return s.XOssProcess
}

func (s *CreateSelectObjectMetaRequest) SetCsvMetaRequest(v *SelectMetaRequest) *CreateSelectObjectMetaRequest {
	s.CsvMetaRequest = v
	return s
}

func (s *CreateSelectObjectMetaRequest) SetXOssProcess(v string) *CreateSelectObjectMetaRequest {
	s.XOssProcess = &v
	return s
}

func (s *CreateSelectObjectMetaRequest) Validate() error {
	if s.CsvMetaRequest != nil {
		if err := s.CsvMetaRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}
