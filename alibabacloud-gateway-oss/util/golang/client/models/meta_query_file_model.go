// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetaQueryFile interface {
	dara.Model
	String() string
	GoString() string
	SetAccessControlAllowOrigin(v string) *MetaQueryFile
	GetAccessControlAllowOrigin() *string
	SetAccessControlRequestMethod(v string) *MetaQueryFile
	GetAccessControlRequestMethod() *string
	SetAddresses(v *MetaQueryFileAddresses) *MetaQueryFile
	GetAddresses() *MetaQueryFileAddresses
	SetAlbum(v string) *MetaQueryFile
	GetAlbum() *string
	SetAlbumArtist(v string) *MetaQueryFile
	GetAlbumArtist() *string
	SetArtist(v string) *MetaQueryFile
	GetArtist() *string
	SetAudioStreams(v *MetaQueryFileAudioStreams) *MetaQueryFile
	GetAudioStreams() *MetaQueryFileAudioStreams
	SetBitrate(v int64) *MetaQueryFile
	GetBitrate() *int64
	SetCacheControl(v string) *MetaQueryFile
	GetCacheControl() *string
	SetComposer(v string) *MetaQueryFile
	GetComposer() *string
	SetContentDisposition(v string) *MetaQueryFile
	GetContentDisposition() *string
	SetContentEncoding(v string) *MetaQueryFile
	GetContentEncoding() *string
	SetContentLanguage(v string) *MetaQueryFile
	GetContentLanguage() *string
	SetContentType(v string) *MetaQueryFile
	GetContentType() *string
	SetDuration(v float64) *MetaQueryFile
	GetDuration() *float64
	SetETag(v string) *MetaQueryFile
	GetETag() *string
	SetFileModifiedTime(v string) *MetaQueryFile
	GetFileModifiedTime() *string
	SetFilename(v string) *MetaQueryFile
	GetFilename() *string
	SetImageHeight(v int64) *MetaQueryFile
	GetImageHeight() *int64
	SetImageWidth(v int64) *MetaQueryFile
	GetImageWidth() *int64
	SetLatLong(v string) *MetaQueryFile
	GetLatLong() *string
	SetMediaType(v string) *MetaQueryFile
	GetMediaType() *string
	SetOSSCRC64(v string) *MetaQueryFile
	GetOSSCRC64() *string
	SetOSSExpiration(v string) *MetaQueryFile
	GetOSSExpiration() *string
	SetOSSObjectType(v string) *MetaQueryFile
	GetOSSObjectType() *string
	SetOSSStorageClass(v string) *MetaQueryFile
	GetOSSStorageClass() *string
	SetOSSTagging(v *MetaQueryFileOSSTagging) *MetaQueryFile
	GetOSSTagging() *MetaQueryFileOSSTagging
	SetOSSTaggingCount(v int64) *MetaQueryFile
	GetOSSTaggingCount() *int64
	SetOSSUserMeta(v *MetaQueryFileOSSUserMeta) *MetaQueryFile
	GetOSSUserMeta() *MetaQueryFileOSSUserMeta
	SetObjectACL(v string) *MetaQueryFile
	GetObjectACL() *string
	SetPerformer(v string) *MetaQueryFile
	GetPerformer() *string
	SetProduceTime(v string) *MetaQueryFile
	GetProduceTime() *string
	SetServerSideDataEncryption(v string) *MetaQueryFile
	GetServerSideDataEncryption() *string
	SetServerSideEncryption(v string) *MetaQueryFile
	GetServerSideEncryption() *string
	SetServerSideEncryptionCustomerAlgorithm(v string) *MetaQueryFile
	GetServerSideEncryptionCustomerAlgorithm() *string
	SetServerSideEncryptionKeyId(v string) *MetaQueryFile
	GetServerSideEncryptionKeyId() *string
	SetSize(v int64) *MetaQueryFile
	GetSize() *int64
	SetSubtitles(v *MetaQueryFileSubtitles) *MetaQueryFile
	GetSubtitles() *MetaQueryFileSubtitles
	SetTitle(v string) *MetaQueryFile
	GetTitle() *string
	SetURI(v string) *MetaQueryFile
	GetURI() *string
	SetVideoHeight(v int64) *MetaQueryFile
	GetVideoHeight() *int64
	SetVideoStreams(v *MetaQueryFileVideoStreams) *MetaQueryFile
	GetVideoStreams() *MetaQueryFileVideoStreams
	SetVideoWidth(v int64) *MetaQueryFile
	GetVideoWidth() *int64
}

type MetaQueryFile struct {
	// example:
	//
	// https://aliyundoc.com
	AccessControlAllowOrigin *string `json:"AccessControlAllowOrigin,omitempty" xml:"AccessControlAllowOrigin,omitempty"`
	// example:
	//
	// PUT
	AccessControlRequestMethod *string                 `json:"AccessControlRequestMethod,omitempty" xml:"AccessControlRequestMethod,omitempty"`
	Addresses                  *MetaQueryFileAddresses `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Struct"`
	// example:
	//
	// FirstAlbum
	Album *string `json:"Album,omitempty" xml:"Album,omitempty"`
	// example:
	//
	// Jenny
	AlbumArtist *string `json:"AlbumArtist,omitempty" xml:"AlbumArtist,omitempty"`
	// example:
	//
	// Jane
	Artist       *string                    `json:"Artist,omitempty" xml:"Artist,omitempty"`
	AudioStreams *MetaQueryFileAudioStreams `json:"AudioStreams,omitempty" xml:"AudioStreams,omitempty" type:"Struct"`
	// example:
	//
	// 13091201
	Bitrate *int64 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// example:
	//
	// no-cache
	CacheControl *string `json:"CacheControl,omitempty" xml:"CacheControl,omitempty"`
	// example:
	//
	// Jane
	Composer *string `json:"Composer,omitempty" xml:"Composer,omitempty"`
	// example:
	//
	// attachment; filename =test.jpg
	ContentDisposition *string `json:"ContentDisposition,omitempty" xml:"ContentDisposition,omitempty"`
	// example:
	//
	// UTF-8
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// example:
	//
	// zh-CN
	ContentLanguage *string `json:"ContentLanguage,omitempty" xml:"ContentLanguage,omitempty"`
	// example:
	//
	// image/jpeg
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// example:
	//
	// 15.263000
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// "fba9dede5f27731c9771645a3986****"
	ETag *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	// example:
	//
	// 2021-06-29T15:04:05.000000000Z07:00
	FileModifiedTime *string `json:"FileModifiedTime,omitempty" xml:"FileModifiedTime,omitempty"`
	// example:
	//
	// exampleobject.txt
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// example:
	//
	// 500
	ImageHeight *int64 `json:"ImageHeight,omitempty" xml:"ImageHeight,omitempty"`
	// example:
	//
	// 270
	ImageWidth *int64 `json:"ImageWidth,omitempty" xml:"ImageWidth,omitempty"`
	// example:
	//
	// 30.134390,120.074997
	LatLong *string `json:"LatLong,omitempty" xml:"LatLong,omitempty"`
	// example:
	//
	// image
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// example:
	//
	// 4858A48BD1466884
	OSSCRC64 *string `json:"OSSCRC64,omitempty" xml:"OSSCRC64,omitempty"`
	// example:
	//
	// 2124-12-01T12:00:00.000Z
	OSSExpiration *string `json:"OSSExpiration,omitempty" xml:"OSSExpiration,omitempty"`
	// example:
	//
	// Normal
	OSSObjectType *string `json:"OSSObjectType,omitempty" xml:"OSSObjectType,omitempty"`
	// example:
	//
	// Standard
	OSSStorageClass *string                  `json:"OSSStorageClass,omitempty" xml:"OSSStorageClass,omitempty"`
	OSSTagging      *MetaQueryFileOSSTagging `json:"OSSTagging,omitempty" xml:"OSSTagging,omitempty" type:"Struct"`
	// example:
	//
	// 2
	OSSTaggingCount *int64                    `json:"OSSTaggingCount,omitempty" xml:"OSSTaggingCount,omitempty"`
	OSSUserMeta     *MetaQueryFileOSSUserMeta `json:"OSSUserMeta,omitempty" xml:"OSSUserMeta,omitempty" type:"Struct"`
	// example:
	//
	// default
	ObjectACL *string `json:"ObjectACL,omitempty" xml:"ObjectACL,omitempty"`
	// example:
	//
	// Jane
	Performer *string `json:"Performer,omitempty" xml:"Performer,omitempty"`
	// example:
	//
	// 2021-06-29T14:50:13.011643661+08:00
	ProduceTime *string `json:"ProduceTime,omitempty" xml:"ProduceTime,omitempty"`
	// example:
	//
	// SM4
	ServerSideDataEncryption *string `json:"ServerSideDataEncryption,omitempty" xml:"ServerSideDataEncryption,omitempty"`
	// example:
	//
	// AES256
	ServerSideEncryption *string `json:"ServerSideEncryption,omitempty" xml:"ServerSideEncryption,omitempty"`
	// example:
	//
	// SM4
	ServerSideEncryptionCustomerAlgorithm *string `json:"ServerSideEncryptionCustomerAlgorithm,omitempty" xml:"ServerSideEncryptionCustomerAlgorithm,omitempty"`
	// example:
	//
	// 9468da86-3509-4f8d-a61e-6eab1eac****
	ServerSideEncryptionKeyId *string `json:"ServerSideEncryptionKeyId,omitempty" xml:"ServerSideEncryptionKeyId,omitempty"`
	// example:
	//
	// 120
	Size      *int64                  `json:"Size,omitempty" xml:"Size,omitempty"`
	Subtitles *MetaQueryFileSubtitles `json:"Subtitles,omitempty" xml:"Subtitles,omitempty" type:"Struct"`
	// example:
	//
	// test
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// oss://examplebucket/test-object.jpg
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
	// example:
	//
	// 1920
	VideoHeight  *int64                     `json:"VideoHeight,omitempty" xml:"VideoHeight,omitempty"`
	VideoStreams *MetaQueryFileVideoStreams `json:"VideoStreams,omitempty" xml:"VideoStreams,omitempty" type:"Struct"`
	// example:
	//
	// 1080
	VideoWidth *int64 `json:"VideoWidth,omitempty" xml:"VideoWidth,omitempty"`
}

func (s MetaQueryFile) String() string {
	return dara.Prettify(s)
}

func (s MetaQueryFile) GoString() string {
	return s.String()
}

func (s *MetaQueryFile) GetAccessControlAllowOrigin() *string {
	return s.AccessControlAllowOrigin
}

func (s *MetaQueryFile) GetAccessControlRequestMethod() *string {
	return s.AccessControlRequestMethod
}

func (s *MetaQueryFile) GetAddresses() *MetaQueryFileAddresses {
	return s.Addresses
}

func (s *MetaQueryFile) GetAlbum() *string {
	return s.Album
}

func (s *MetaQueryFile) GetAlbumArtist() *string {
	return s.AlbumArtist
}

func (s *MetaQueryFile) GetArtist() *string {
	return s.Artist
}

func (s *MetaQueryFile) GetAudioStreams() *MetaQueryFileAudioStreams {
	return s.AudioStreams
}

func (s *MetaQueryFile) GetBitrate() *int64 {
	return s.Bitrate
}

func (s *MetaQueryFile) GetCacheControl() *string {
	return s.CacheControl
}

func (s *MetaQueryFile) GetComposer() *string {
	return s.Composer
}

func (s *MetaQueryFile) GetContentDisposition() *string {
	return s.ContentDisposition
}

func (s *MetaQueryFile) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *MetaQueryFile) GetContentLanguage() *string {
	return s.ContentLanguage
}

func (s *MetaQueryFile) GetContentType() *string {
	return s.ContentType
}

func (s *MetaQueryFile) GetDuration() *float64 {
	return s.Duration
}

func (s *MetaQueryFile) GetETag() *string {
	return s.ETag
}

func (s *MetaQueryFile) GetFileModifiedTime() *string {
	return s.FileModifiedTime
}

func (s *MetaQueryFile) GetFilename() *string {
	return s.Filename
}

func (s *MetaQueryFile) GetImageHeight() *int64 {
	return s.ImageHeight
}

func (s *MetaQueryFile) GetImageWidth() *int64 {
	return s.ImageWidth
}

func (s *MetaQueryFile) GetLatLong() *string {
	return s.LatLong
}

func (s *MetaQueryFile) GetMediaType() *string {
	return s.MediaType
}

func (s *MetaQueryFile) GetOSSCRC64() *string {
	return s.OSSCRC64
}

func (s *MetaQueryFile) GetOSSExpiration() *string {
	return s.OSSExpiration
}

func (s *MetaQueryFile) GetOSSObjectType() *string {
	return s.OSSObjectType
}

func (s *MetaQueryFile) GetOSSStorageClass() *string {
	return s.OSSStorageClass
}

func (s *MetaQueryFile) GetOSSTagging() *MetaQueryFileOSSTagging {
	return s.OSSTagging
}

func (s *MetaQueryFile) GetOSSTaggingCount() *int64 {
	return s.OSSTaggingCount
}

func (s *MetaQueryFile) GetOSSUserMeta() *MetaQueryFileOSSUserMeta {
	return s.OSSUserMeta
}

func (s *MetaQueryFile) GetObjectACL() *string {
	return s.ObjectACL
}

func (s *MetaQueryFile) GetPerformer() *string {
	return s.Performer
}

func (s *MetaQueryFile) GetProduceTime() *string {
	return s.ProduceTime
}

func (s *MetaQueryFile) GetServerSideDataEncryption() *string {
	return s.ServerSideDataEncryption
}

func (s *MetaQueryFile) GetServerSideEncryption() *string {
	return s.ServerSideEncryption
}

func (s *MetaQueryFile) GetServerSideEncryptionCustomerAlgorithm() *string {
	return s.ServerSideEncryptionCustomerAlgorithm
}

func (s *MetaQueryFile) GetServerSideEncryptionKeyId() *string {
	return s.ServerSideEncryptionKeyId
}

func (s *MetaQueryFile) GetSize() *int64 {
	return s.Size
}

func (s *MetaQueryFile) GetSubtitles() *MetaQueryFileSubtitles {
	return s.Subtitles
}

func (s *MetaQueryFile) GetTitle() *string {
	return s.Title
}

func (s *MetaQueryFile) GetURI() *string {
	return s.URI
}

func (s *MetaQueryFile) GetVideoHeight() *int64 {
	return s.VideoHeight
}

func (s *MetaQueryFile) GetVideoStreams() *MetaQueryFileVideoStreams {
	return s.VideoStreams
}

func (s *MetaQueryFile) GetVideoWidth() *int64 {
	return s.VideoWidth
}

func (s *MetaQueryFile) SetAccessControlAllowOrigin(v string) *MetaQueryFile {
	s.AccessControlAllowOrigin = &v
	return s
}

func (s *MetaQueryFile) SetAccessControlRequestMethod(v string) *MetaQueryFile {
	s.AccessControlRequestMethod = &v
	return s
}

func (s *MetaQueryFile) SetAddresses(v *MetaQueryFileAddresses) *MetaQueryFile {
	s.Addresses = v
	return s
}

func (s *MetaQueryFile) SetAlbum(v string) *MetaQueryFile {
	s.Album = &v
	return s
}

func (s *MetaQueryFile) SetAlbumArtist(v string) *MetaQueryFile {
	s.AlbumArtist = &v
	return s
}

func (s *MetaQueryFile) SetArtist(v string) *MetaQueryFile {
	s.Artist = &v
	return s
}

func (s *MetaQueryFile) SetAudioStreams(v *MetaQueryFileAudioStreams) *MetaQueryFile {
	s.AudioStreams = v
	return s
}

func (s *MetaQueryFile) SetBitrate(v int64) *MetaQueryFile {
	s.Bitrate = &v
	return s
}

func (s *MetaQueryFile) SetCacheControl(v string) *MetaQueryFile {
	s.CacheControl = &v
	return s
}

func (s *MetaQueryFile) SetComposer(v string) *MetaQueryFile {
	s.Composer = &v
	return s
}

func (s *MetaQueryFile) SetContentDisposition(v string) *MetaQueryFile {
	s.ContentDisposition = &v
	return s
}

func (s *MetaQueryFile) SetContentEncoding(v string) *MetaQueryFile {
	s.ContentEncoding = &v
	return s
}

func (s *MetaQueryFile) SetContentLanguage(v string) *MetaQueryFile {
	s.ContentLanguage = &v
	return s
}

func (s *MetaQueryFile) SetContentType(v string) *MetaQueryFile {
	s.ContentType = &v
	return s
}

func (s *MetaQueryFile) SetDuration(v float64) *MetaQueryFile {
	s.Duration = &v
	return s
}

func (s *MetaQueryFile) SetETag(v string) *MetaQueryFile {
	s.ETag = &v
	return s
}

func (s *MetaQueryFile) SetFileModifiedTime(v string) *MetaQueryFile {
	s.FileModifiedTime = &v
	return s
}

func (s *MetaQueryFile) SetFilename(v string) *MetaQueryFile {
	s.Filename = &v
	return s
}

func (s *MetaQueryFile) SetImageHeight(v int64) *MetaQueryFile {
	s.ImageHeight = &v
	return s
}

func (s *MetaQueryFile) SetImageWidth(v int64) *MetaQueryFile {
	s.ImageWidth = &v
	return s
}

func (s *MetaQueryFile) SetLatLong(v string) *MetaQueryFile {
	s.LatLong = &v
	return s
}

func (s *MetaQueryFile) SetMediaType(v string) *MetaQueryFile {
	s.MediaType = &v
	return s
}

func (s *MetaQueryFile) SetOSSCRC64(v string) *MetaQueryFile {
	s.OSSCRC64 = &v
	return s
}

func (s *MetaQueryFile) SetOSSExpiration(v string) *MetaQueryFile {
	s.OSSExpiration = &v
	return s
}

func (s *MetaQueryFile) SetOSSObjectType(v string) *MetaQueryFile {
	s.OSSObjectType = &v
	return s
}

func (s *MetaQueryFile) SetOSSStorageClass(v string) *MetaQueryFile {
	s.OSSStorageClass = &v
	return s
}

func (s *MetaQueryFile) SetOSSTagging(v *MetaQueryFileOSSTagging) *MetaQueryFile {
	s.OSSTagging = v
	return s
}

func (s *MetaQueryFile) SetOSSTaggingCount(v int64) *MetaQueryFile {
	s.OSSTaggingCount = &v
	return s
}

func (s *MetaQueryFile) SetOSSUserMeta(v *MetaQueryFileOSSUserMeta) *MetaQueryFile {
	s.OSSUserMeta = v
	return s
}

func (s *MetaQueryFile) SetObjectACL(v string) *MetaQueryFile {
	s.ObjectACL = &v
	return s
}

func (s *MetaQueryFile) SetPerformer(v string) *MetaQueryFile {
	s.Performer = &v
	return s
}

func (s *MetaQueryFile) SetProduceTime(v string) *MetaQueryFile {
	s.ProduceTime = &v
	return s
}

func (s *MetaQueryFile) SetServerSideDataEncryption(v string) *MetaQueryFile {
	s.ServerSideDataEncryption = &v
	return s
}

func (s *MetaQueryFile) SetServerSideEncryption(v string) *MetaQueryFile {
	s.ServerSideEncryption = &v
	return s
}

func (s *MetaQueryFile) SetServerSideEncryptionCustomerAlgorithm(v string) *MetaQueryFile {
	s.ServerSideEncryptionCustomerAlgorithm = &v
	return s
}

func (s *MetaQueryFile) SetServerSideEncryptionKeyId(v string) *MetaQueryFile {
	s.ServerSideEncryptionKeyId = &v
	return s
}

func (s *MetaQueryFile) SetSize(v int64) *MetaQueryFile {
	s.Size = &v
	return s
}

func (s *MetaQueryFile) SetSubtitles(v *MetaQueryFileSubtitles) *MetaQueryFile {
	s.Subtitles = v
	return s
}

func (s *MetaQueryFile) SetTitle(v string) *MetaQueryFile {
	s.Title = &v
	return s
}

func (s *MetaQueryFile) SetURI(v string) *MetaQueryFile {
	s.URI = &v
	return s
}

func (s *MetaQueryFile) SetVideoHeight(v int64) *MetaQueryFile {
	s.VideoHeight = &v
	return s
}

func (s *MetaQueryFile) SetVideoStreams(v *MetaQueryFileVideoStreams) *MetaQueryFile {
	s.VideoStreams = v
	return s
}

func (s *MetaQueryFile) SetVideoWidth(v int64) *MetaQueryFile {
	s.VideoWidth = &v
	return s
}

func (s *MetaQueryFile) Validate() error {
	if s.Addresses != nil {
		if err := s.Addresses.Validate(); err != nil {
			return err
		}
	}
	if s.AudioStreams != nil {
		if err := s.AudioStreams.Validate(); err != nil {
			return err
		}
	}
	if s.OSSTagging != nil {
		if err := s.OSSTagging.Validate(); err != nil {
			return err
		}
	}
	if s.OSSUserMeta != nil {
		if err := s.OSSUserMeta.Validate(); err != nil {
			return err
		}
	}
	if s.Subtitles != nil {
		if err := s.Subtitles.Validate(); err != nil {
			return err
		}
	}
	if s.VideoStreams != nil {
		if err := s.VideoStreams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MetaQueryFileAddresses struct {
	Address *MetaQueryRespAddress `json:"Address,omitempty" xml:"Address,omitempty"`
}

func (s MetaQueryFileAddresses) String() string {
	return dara.Prettify(s)
}

func (s MetaQueryFileAddresses) GoString() string {
	return s.String()
}

func (s *MetaQueryFileAddresses) GetAddress() *MetaQueryRespAddress {
	return s.Address
}

func (s *MetaQueryFileAddresses) SetAddress(v *MetaQueryRespAddress) *MetaQueryFileAddresses {
	s.Address = v
	return s
}

func (s *MetaQueryFileAddresses) Validate() error {
	if s.Address != nil {
		if err := s.Address.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MetaQueryFileAudioStreams struct {
	AudioStream *MetaQueryRespAudioStream `json:"AudioStream,omitempty" xml:"AudioStream,omitempty"`
}

func (s MetaQueryFileAudioStreams) String() string {
	return dara.Prettify(s)
}

func (s MetaQueryFileAudioStreams) GoString() string {
	return s.String()
}

func (s *MetaQueryFileAudioStreams) GetAudioStream() *MetaQueryRespAudioStream {
	return s.AudioStream
}

func (s *MetaQueryFileAudioStreams) SetAudioStream(v *MetaQueryRespAudioStream) *MetaQueryFileAudioStreams {
	s.AudioStream = v
	return s
}

func (s *MetaQueryFileAudioStreams) Validate() error {
	if s.AudioStream != nil {
		if err := s.AudioStream.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MetaQueryFileOSSTagging struct {
	Tagging []*MetaQueryTagging `json:"Tagging,omitempty" xml:"Tagging,omitempty" type:"Repeated"`
}

func (s MetaQueryFileOSSTagging) String() string {
	return dara.Prettify(s)
}

func (s MetaQueryFileOSSTagging) GoString() string {
	return s.String()
}

func (s *MetaQueryFileOSSTagging) GetTagging() []*MetaQueryTagging {
	return s.Tagging
}

func (s *MetaQueryFileOSSTagging) SetTagging(v []*MetaQueryTagging) *MetaQueryFileOSSTagging {
	s.Tagging = v
	return s
}

func (s *MetaQueryFileOSSTagging) Validate() error {
	if s.Tagging != nil {
		for _, item := range s.Tagging {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MetaQueryFileOSSUserMeta struct {
	UserMeta []*MetaQueryUserMeta `json:"UserMeta,omitempty" xml:"UserMeta,omitempty" type:"Repeated"`
}

func (s MetaQueryFileOSSUserMeta) String() string {
	return dara.Prettify(s)
}

func (s MetaQueryFileOSSUserMeta) GoString() string {
	return s.String()
}

func (s *MetaQueryFileOSSUserMeta) GetUserMeta() []*MetaQueryUserMeta {
	return s.UserMeta
}

func (s *MetaQueryFileOSSUserMeta) SetUserMeta(v []*MetaQueryUserMeta) *MetaQueryFileOSSUserMeta {
	s.UserMeta = v
	return s
}

func (s *MetaQueryFileOSSUserMeta) Validate() error {
	if s.UserMeta != nil {
		for _, item := range s.UserMeta {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MetaQueryFileSubtitles struct {
	Subtitle *MetaQueryRespSubtitle `json:"Subtitle,omitempty" xml:"Subtitle,omitempty"`
}

func (s MetaQueryFileSubtitles) String() string {
	return dara.Prettify(s)
}

func (s MetaQueryFileSubtitles) GoString() string {
	return s.String()
}

func (s *MetaQueryFileSubtitles) GetSubtitle() *MetaQueryRespSubtitle {
	return s.Subtitle
}

func (s *MetaQueryFileSubtitles) SetSubtitle(v *MetaQueryRespSubtitle) *MetaQueryFileSubtitles {
	s.Subtitle = v
	return s
}

func (s *MetaQueryFileSubtitles) Validate() error {
	if s.Subtitle != nil {
		if err := s.Subtitle.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MetaQueryFileVideoStreams struct {
	VideoStream *MetaQueryRespVideoStream `json:"VideoStream,omitempty" xml:"VideoStream,omitempty"`
}

func (s MetaQueryFileVideoStreams) String() string {
	return dara.Prettify(s)
}

func (s MetaQueryFileVideoStreams) GoString() string {
	return s.String()
}

func (s *MetaQueryFileVideoStreams) GetVideoStream() *MetaQueryRespVideoStream {
	return s.VideoStream
}

func (s *MetaQueryFileVideoStreams) SetVideoStream(v *MetaQueryRespVideoStream) *MetaQueryFileVideoStreams {
	s.VideoStream = v
	return s
}

func (s *MetaQueryFileVideoStreams) Validate() error {
	if s.VideoStream != nil {
		if err := s.VideoStream.Validate(); err != nil {
			return err
		}
	}
	return nil
}
