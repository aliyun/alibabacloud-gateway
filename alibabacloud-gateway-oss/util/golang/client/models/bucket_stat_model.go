// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBucketStat interface {
	dara.Model
	String() string
	GoString() string
	SetArchiveObjectCount(v int64) *BucketStat
	GetArchiveObjectCount() *int64
	SetArchiveRealStorage(v int64) *BucketStat
	GetArchiveRealStorage() *int64
	SetArchiveStorage(v int64) *BucketStat
	GetArchiveStorage() *int64
	SetColdArchiveObjectCount(v int64) *BucketStat
	GetColdArchiveObjectCount() *int64
	SetColdArchiveRealStorage(v int64) *BucketStat
	GetColdArchiveRealStorage() *int64
	SetColdArchiveStorage(v int64) *BucketStat
	GetColdArchiveStorage() *int64
	SetDeepColdArchiveObjectCount(v int64) *BucketStat
	GetDeepColdArchiveObjectCount() *int64
	SetDeepColdArchiveRealStorage(v int64) *BucketStat
	GetDeepColdArchiveRealStorage() *int64
	SetDeepColdArchiveStorage(v int64) *BucketStat
	GetDeepColdArchiveStorage() *int64
	SetDeleteMarkerCount(v int64) *BucketStat
	GetDeleteMarkerCount() *int64
	SetInfrequentAccessObjectCount(v int64) *BucketStat
	GetInfrequentAccessObjectCount() *int64
	SetInfrequentAccessRealStorage(v int64) *BucketStat
	GetInfrequentAccessRealStorage() *int64
	SetInfrequentAccessStorage(v int64) *BucketStat
	GetInfrequentAccessStorage() *int64
	SetLastModifiedTime(v int64) *BucketStat
	GetLastModifiedTime() *int64
	SetLiveChannelCount(v int64) *BucketStat
	GetLiveChannelCount() *int64
	SetMultipartPartCount(v int64) *BucketStat
	GetMultipartPartCount() *int64
	SetMultipartUploadCount(v int64) *BucketStat
	GetMultipartUploadCount() *int64
	SetObjectCount(v int64) *BucketStat
	GetObjectCount() *int64
	SetStandardObjectCount(v int64) *BucketStat
	GetStandardObjectCount() *int64
	SetStandardStorage(v int64) *BucketStat
	GetStandardStorage() *int64
	SetStorage(v int64) *BucketStat
	GetStorage() *int64
}

type BucketStat struct {
	// example:
	//
	// 2
	ArchiveObjectCount *int64 `json:"ArchiveObjectCount,omitempty" xml:"ArchiveObjectCount,omitempty"`
	// example:
	//
	// 120
	ArchiveRealStorage *int64 `json:"ArchiveRealStorage,omitempty" xml:"ArchiveRealStorage,omitempty"`
	// example:
	//
	// 120
	ArchiveStorage *int64 `json:"ArchiveStorage,omitempty" xml:"ArchiveStorage,omitempty"`
	// example:
	//
	// 2
	ColdArchiveObjectCount *int64 `json:"ColdArchiveObjectCount,omitempty" xml:"ColdArchiveObjectCount,omitempty"`
	// example:
	//
	// 120
	ColdArchiveRealStorage *int64 `json:"ColdArchiveRealStorage,omitempty" xml:"ColdArchiveRealStorage,omitempty"`
	// example:
	//
	// 120
	ColdArchiveStorage *int64 `json:"ColdArchiveStorage,omitempty" xml:"ColdArchiveStorage,omitempty"`
	// example:
	//
	// 2
	DeepColdArchiveObjectCount *int64 `json:"DeepColdArchiveObjectCount,omitempty" xml:"DeepColdArchiveObjectCount,omitempty"`
	// example:
	//
	// 120
	DeepColdArchiveRealStorage *int64 `json:"DeepColdArchiveRealStorage,omitempty" xml:"DeepColdArchiveRealStorage,omitempty"`
	// example:
	//
	// 120
	DeepColdArchiveStorage *int64 `json:"DeepColdArchiveStorage,omitempty" xml:"DeepColdArchiveStorage,omitempty"`
	// example:
	//
	// 12
	DeleteMarkerCount *int64 `json:"DeleteMarkerCount,omitempty" xml:"DeleteMarkerCount,omitempty"`
	// example:
	//
	// 2
	InfrequentAccessObjectCount *int64 `json:"InfrequentAccessObjectCount,omitempty" xml:"InfrequentAccessObjectCount,omitempty"`
	// example:
	//
	// 120
	InfrequentAccessRealStorage *int64 `json:"InfrequentAccessRealStorage,omitempty" xml:"InfrequentAccessRealStorage,omitempty"`
	// example:
	//
	// 120
	InfrequentAccessStorage *int64 `json:"InfrequentAccessStorage,omitempty" xml:"InfrequentAccessStorage,omitempty"`
	// example:
	//
	// 1709724731
	LastModifiedTime *int64 `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// example:
	//
	// 2
	LiveChannelCount *int64 `json:"LiveChannelCount,omitempty" xml:"LiveChannelCount,omitempty"`
	// example:
	//
	// 128
	MultipartPartCount *int64 `json:"MultipartPartCount,omitempty" xml:"MultipartPartCount,omitempty"`
	// example:
	//
	// 27
	MultipartUploadCount *int64 `json:"MultipartUploadCount,omitempty" xml:"MultipartUploadCount,omitempty"`
	// example:
	//
	// 32
	ObjectCount *int64 `json:"ObjectCount,omitempty" xml:"ObjectCount,omitempty"`
	// example:
	//
	// 18
	StandardObjectCount *int64 `json:"StandardObjectCount,omitempty" xml:"StandardObjectCount,omitempty"`
	// example:
	//
	// 1990
	StandardStorage *int64 `json:"StandardStorage,omitempty" xml:"StandardStorage,omitempty"`
	// example:
	//
	// 1994
	Storage *int64 `json:"Storage,omitempty" xml:"Storage,omitempty"`
}

func (s BucketStat) String() string {
	return dara.Prettify(s)
}

func (s BucketStat) GoString() string {
	return s.String()
}

func (s *BucketStat) GetArchiveObjectCount() *int64 {
	return s.ArchiveObjectCount
}

func (s *BucketStat) GetArchiveRealStorage() *int64 {
	return s.ArchiveRealStorage
}

func (s *BucketStat) GetArchiveStorage() *int64 {
	return s.ArchiveStorage
}

func (s *BucketStat) GetColdArchiveObjectCount() *int64 {
	return s.ColdArchiveObjectCount
}

func (s *BucketStat) GetColdArchiveRealStorage() *int64 {
	return s.ColdArchiveRealStorage
}

func (s *BucketStat) GetColdArchiveStorage() *int64 {
	return s.ColdArchiveStorage
}

func (s *BucketStat) GetDeepColdArchiveObjectCount() *int64 {
	return s.DeepColdArchiveObjectCount
}

func (s *BucketStat) GetDeepColdArchiveRealStorage() *int64 {
	return s.DeepColdArchiveRealStorage
}

func (s *BucketStat) GetDeepColdArchiveStorage() *int64 {
	return s.DeepColdArchiveStorage
}

func (s *BucketStat) GetDeleteMarkerCount() *int64 {
	return s.DeleteMarkerCount
}

func (s *BucketStat) GetInfrequentAccessObjectCount() *int64 {
	return s.InfrequentAccessObjectCount
}

func (s *BucketStat) GetInfrequentAccessRealStorage() *int64 {
	return s.InfrequentAccessRealStorage
}

func (s *BucketStat) GetInfrequentAccessStorage() *int64 {
	return s.InfrequentAccessStorage
}

func (s *BucketStat) GetLastModifiedTime() *int64 {
	return s.LastModifiedTime
}

func (s *BucketStat) GetLiveChannelCount() *int64 {
	return s.LiveChannelCount
}

func (s *BucketStat) GetMultipartPartCount() *int64 {
	return s.MultipartPartCount
}

func (s *BucketStat) GetMultipartUploadCount() *int64 {
	return s.MultipartUploadCount
}

func (s *BucketStat) GetObjectCount() *int64 {
	return s.ObjectCount
}

func (s *BucketStat) GetStandardObjectCount() *int64 {
	return s.StandardObjectCount
}

func (s *BucketStat) GetStandardStorage() *int64 {
	return s.StandardStorage
}

func (s *BucketStat) GetStorage() *int64 {
	return s.Storage
}

func (s *BucketStat) SetArchiveObjectCount(v int64) *BucketStat {
	s.ArchiveObjectCount = &v
	return s
}

func (s *BucketStat) SetArchiveRealStorage(v int64) *BucketStat {
	s.ArchiveRealStorage = &v
	return s
}

func (s *BucketStat) SetArchiveStorage(v int64) *BucketStat {
	s.ArchiveStorage = &v
	return s
}

func (s *BucketStat) SetColdArchiveObjectCount(v int64) *BucketStat {
	s.ColdArchiveObjectCount = &v
	return s
}

func (s *BucketStat) SetColdArchiveRealStorage(v int64) *BucketStat {
	s.ColdArchiveRealStorage = &v
	return s
}

func (s *BucketStat) SetColdArchiveStorage(v int64) *BucketStat {
	s.ColdArchiveStorage = &v
	return s
}

func (s *BucketStat) SetDeepColdArchiveObjectCount(v int64) *BucketStat {
	s.DeepColdArchiveObjectCount = &v
	return s
}

func (s *BucketStat) SetDeepColdArchiveRealStorage(v int64) *BucketStat {
	s.DeepColdArchiveRealStorage = &v
	return s
}

func (s *BucketStat) SetDeepColdArchiveStorage(v int64) *BucketStat {
	s.DeepColdArchiveStorage = &v
	return s
}

func (s *BucketStat) SetDeleteMarkerCount(v int64) *BucketStat {
	s.DeleteMarkerCount = &v
	return s
}

func (s *BucketStat) SetInfrequentAccessObjectCount(v int64) *BucketStat {
	s.InfrequentAccessObjectCount = &v
	return s
}

func (s *BucketStat) SetInfrequentAccessRealStorage(v int64) *BucketStat {
	s.InfrequentAccessRealStorage = &v
	return s
}

func (s *BucketStat) SetInfrequentAccessStorage(v int64) *BucketStat {
	s.InfrequentAccessStorage = &v
	return s
}

func (s *BucketStat) SetLastModifiedTime(v int64) *BucketStat {
	s.LastModifiedTime = &v
	return s
}

func (s *BucketStat) SetLiveChannelCount(v int64) *BucketStat {
	s.LiveChannelCount = &v
	return s
}

func (s *BucketStat) SetMultipartPartCount(v int64) *BucketStat {
	s.MultipartPartCount = &v
	return s
}

func (s *BucketStat) SetMultipartUploadCount(v int64) *BucketStat {
	s.MultipartUploadCount = &v
	return s
}

func (s *BucketStat) SetObjectCount(v int64) *BucketStat {
	s.ObjectCount = &v
	return s
}

func (s *BucketStat) SetStandardObjectCount(v int64) *BucketStat {
	s.StandardObjectCount = &v
	return s
}

func (s *BucketStat) SetStandardStorage(v int64) *BucketStat {
	s.StandardStorage = &v
	return s
}

func (s *BucketStat) SetStorage(v int64) *BucketStat {
	s.Storage = &v
	return s
}

func (s *BucketStat) Validate() error {
	return dara.Validate(s)
}
