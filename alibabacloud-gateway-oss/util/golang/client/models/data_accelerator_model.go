// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataAccelerator interface {
	dara.Model
	String() string
	GoString() string
	SetBasicInfomation(v *DataAcceleratorBasicInfomation) *DataAccelerator
	GetBasicInfomation() *DataAcceleratorBasicInfomation
	SetBucketName(v string) *DataAccelerator
	GetBucketName() *string
	SetName(v string) *DataAccelerator
	GetName() *string
}

type DataAccelerator struct {
	BasicInfomation *DataAcceleratorBasicInfomation `json:"BasicInfomation,omitempty" xml:"BasicInfomation,omitempty" type:"Struct"`
	// example:
	//
	// test-acc-bucket
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// example:
	//
	// test-acc-bucket_data-acc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DataAccelerator) String() string {
	return dara.Prettify(s)
}

func (s DataAccelerator) GoString() string {
	return s.String()
}

func (s *DataAccelerator) GetBasicInfomation() *DataAcceleratorBasicInfomation {
	return s.BasicInfomation
}

func (s *DataAccelerator) GetBucketName() *string {
	return s.BucketName
}

func (s *DataAccelerator) GetName() *string {
	return s.Name
}

func (s *DataAccelerator) SetBasicInfomation(v *DataAcceleratorBasicInfomation) *DataAccelerator {
	s.BasicInfomation = v
	return s
}

func (s *DataAccelerator) SetBucketName(v string) *DataAccelerator {
	s.BucketName = &v
	return s
}

func (s *DataAccelerator) SetName(v string) *DataAccelerator {
	s.Name = &v
	return s
}

func (s *DataAccelerator) Validate() error {
	if s.BasicInfomation != nil {
		if err := s.BasicInfomation.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DataAcceleratorBasicInfomation struct {
	AcceleratePaths *AcceleratePaths `json:"AcceleratePaths,omitempty" xml:"AcceleratePaths,omitempty"`
	// example:
	//
	// cn-hangzhou-a
	AvailableZone *string `json:"AvailableZone,omitempty" xml:"AvailableZone,omitempty"`
	// example:
	//
	// 1731394193189
	CreationDate *int64 `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	// example:
	//
	// 1024000
	Quota *string `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// example:
	//
	// 1731394193189
	QuotaFrozenUtil *int64 `json:"QuotaFrozenUtil,omitempty" xml:"QuotaFrozenUtil,omitempty"`
}

func (s DataAcceleratorBasicInfomation) String() string {
	return dara.Prettify(s)
}

func (s DataAcceleratorBasicInfomation) GoString() string {
	return s.String()
}

func (s *DataAcceleratorBasicInfomation) GetAcceleratePaths() *AcceleratePaths {
	return s.AcceleratePaths
}

func (s *DataAcceleratorBasicInfomation) GetAvailableZone() *string {
	return s.AvailableZone
}

func (s *DataAcceleratorBasicInfomation) GetCreationDate() *int64 {
	return s.CreationDate
}

func (s *DataAcceleratorBasicInfomation) GetQuota() *string {
	return s.Quota
}

func (s *DataAcceleratorBasicInfomation) GetQuotaFrozenUtil() *int64 {
	return s.QuotaFrozenUtil
}

func (s *DataAcceleratorBasicInfomation) SetAcceleratePaths(v *AcceleratePaths) *DataAcceleratorBasicInfomation {
	s.AcceleratePaths = v
	return s
}

func (s *DataAcceleratorBasicInfomation) SetAvailableZone(v string) *DataAcceleratorBasicInfomation {
	s.AvailableZone = &v
	return s
}

func (s *DataAcceleratorBasicInfomation) SetCreationDate(v int64) *DataAcceleratorBasicInfomation {
	s.CreationDate = &v
	return s
}

func (s *DataAcceleratorBasicInfomation) SetQuota(v string) *DataAcceleratorBasicInfomation {
	s.Quota = &v
	return s
}

func (s *DataAcceleratorBasicInfomation) SetQuotaFrozenUtil(v int64) *DataAcceleratorBasicInfomation {
	s.QuotaFrozenUtil = &v
	return s
}

func (s *DataAcceleratorBasicInfomation) Validate() error {
	if s.AcceleratePaths != nil {
		if err := s.AcceleratePaths.Validate(); err != nil {
			return err
		}
	}
	return nil
}
