# -*- coding: utf-8 -*-
# This file is auto-generated, don't edit it. Thanks.
from __future__ import annotations

from typing import List

from alibabacloud_oss20190517 import models as main_models
from darabonba.model import DaraModel

class MetaQueryFile(DaraModel):
    def __init__(
        self,
        access_control_allow_origin: str = None,
        access_control_request_method: str = None,
        addresses: main_models.MetaQueryFileAddresses = None,
        album: str = None,
        album_artist: str = None,
        artist: str = None,
        audio_streams: main_models.MetaQueryFileAudioStreams = None,
        bitrate: int = None,
        cache_control: str = None,
        composer: str = None,
        content_disposition: str = None,
        content_encoding: str = None,
        content_language: str = None,
        content_type: str = None,
        duration: float = None,
        etag: str = None,
        file_modified_time: str = None,
        filename: str = None,
        image_height: int = None,
        image_width: int = None,
        lat_long: str = None,
        media_type: str = None,
        osscrc64: str = None,
        ossexpiration: str = None,
        ossobject_type: str = None,
        ossstorage_class: str = None,
        osstagging: main_models.MetaQueryFileOSSTagging = None,
        osstagging_count: int = None,
        ossuser_meta: main_models.MetaQueryFileOSSUserMeta = None,
        object_acl: str = None,
        performer: str = None,
        produce_time: str = None,
        server_side_data_encryption: str = None,
        server_side_encryption: str = None,
        server_side_encryption_customer_algorithm: str = None,
        server_side_encryption_key_id: str = None,
        size: int = None,
        subtitles: main_models.MetaQueryFileSubtitles = None,
        title: str = None,
        uri: str = None,
        video_height: int = None,
        video_streams: main_models.MetaQueryFileVideoStreams = None,
        video_width: int = None,
    ):
        self.access_control_allow_origin = access_control_allow_origin
        self.access_control_request_method = access_control_request_method
        self.addresses = addresses
        self.album = album
        self.album_artist = album_artist
        self.artist = artist
        self.audio_streams = audio_streams
        self.bitrate = bitrate
        self.cache_control = cache_control
        self.composer = composer
        self.content_disposition = content_disposition
        self.content_encoding = content_encoding
        self.content_language = content_language
        self.content_type = content_type
        self.duration = duration
        self.etag = etag
        self.file_modified_time = file_modified_time
        self.filename = filename
        self.image_height = image_height
        self.image_width = image_width
        self.lat_long = lat_long
        self.media_type = media_type
        self.osscrc64 = osscrc64
        self.ossexpiration = ossexpiration
        self.ossobject_type = ossobject_type
        self.ossstorage_class = ossstorage_class
        self.osstagging = osstagging
        self.osstagging_count = osstagging_count
        self.ossuser_meta = ossuser_meta
        self.object_acl = object_acl
        self.performer = performer
        self.produce_time = produce_time
        self.server_side_data_encryption = server_side_data_encryption
        self.server_side_encryption = server_side_encryption
        self.server_side_encryption_customer_algorithm = server_side_encryption_customer_algorithm
        self.server_side_encryption_key_id = server_side_encryption_key_id
        self.size = size
        self.subtitles = subtitles
        self.title = title
        self.uri = uri
        self.video_height = video_height
        self.video_streams = video_streams
        self.video_width = video_width

    def validate(self):
        if self.addresses:
            self.addresses.validate()
        if self.audio_streams:
            self.audio_streams.validate()
        if self.osstagging:
            self.osstagging.validate()
        if self.ossuser_meta:
            self.ossuser_meta.validate()
        if self.subtitles:
            self.subtitles.validate()
        if self.video_streams:
            self.video_streams.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.access_control_allow_origin is not None:
            result['AccessControlAllowOrigin'] = self.access_control_allow_origin

        if self.access_control_request_method is not None:
            result['AccessControlRequestMethod'] = self.access_control_request_method

        if self.addresses is not None:
            result['Addresses'] = self.addresses.to_map()

        if self.album is not None:
            result['Album'] = self.album

        if self.album_artist is not None:
            result['AlbumArtist'] = self.album_artist

        if self.artist is not None:
            result['Artist'] = self.artist

        if self.audio_streams is not None:
            result['AudioStreams'] = self.audio_streams.to_map()

        if self.bitrate is not None:
            result['Bitrate'] = self.bitrate

        if self.cache_control is not None:
            result['CacheControl'] = self.cache_control

        if self.composer is not None:
            result['Composer'] = self.composer

        if self.content_disposition is not None:
            result['ContentDisposition'] = self.content_disposition

        if self.content_encoding is not None:
            result['ContentEncoding'] = self.content_encoding

        if self.content_language is not None:
            result['ContentLanguage'] = self.content_language

        if self.content_type is not None:
            result['ContentType'] = self.content_type

        if self.duration is not None:
            result['Duration'] = self.duration

        if self.etag is not None:
            result['ETag'] = self.etag

        if self.file_modified_time is not None:
            result['FileModifiedTime'] = self.file_modified_time

        if self.filename is not None:
            result['Filename'] = self.filename

        if self.image_height is not None:
            result['ImageHeight'] = self.image_height

        if self.image_width is not None:
            result['ImageWidth'] = self.image_width

        if self.lat_long is not None:
            result['LatLong'] = self.lat_long

        if self.media_type is not None:
            result['MediaType'] = self.media_type

        if self.osscrc64 is not None:
            result['OSSCRC64'] = self.osscrc64

        if self.ossexpiration is not None:
            result['OSSExpiration'] = self.ossexpiration

        if self.ossobject_type is not None:
            result['OSSObjectType'] = self.ossobject_type

        if self.ossstorage_class is not None:
            result['OSSStorageClass'] = self.ossstorage_class

        if self.osstagging is not None:
            result['OSSTagging'] = self.osstagging.to_map()

        if self.osstagging_count is not None:
            result['OSSTaggingCount'] = self.osstagging_count

        if self.ossuser_meta is not None:
            result['OSSUserMeta'] = self.ossuser_meta.to_map()

        if self.object_acl is not None:
            result['ObjectACL'] = self.object_acl

        if self.performer is not None:
            result['Performer'] = self.performer

        if self.produce_time is not None:
            result['ProduceTime'] = self.produce_time

        if self.server_side_data_encryption is not None:
            result['ServerSideDataEncryption'] = self.server_side_data_encryption

        if self.server_side_encryption is not None:
            result['ServerSideEncryption'] = self.server_side_encryption

        if self.server_side_encryption_customer_algorithm is not None:
            result['ServerSideEncryptionCustomerAlgorithm'] = self.server_side_encryption_customer_algorithm

        if self.server_side_encryption_key_id is not None:
            result['ServerSideEncryptionKeyId'] = self.server_side_encryption_key_id

        if self.size is not None:
            result['Size'] = self.size

        if self.subtitles is not None:
            result['Subtitles'] = self.subtitles.to_map()

        if self.title is not None:
            result['Title'] = self.title

        if self.uri is not None:
            result['URI'] = self.uri

        if self.video_height is not None:
            result['VideoHeight'] = self.video_height

        if self.video_streams is not None:
            result['VideoStreams'] = self.video_streams.to_map()

        if self.video_width is not None:
            result['VideoWidth'] = self.video_width

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AccessControlAllowOrigin') is not None:
            self.access_control_allow_origin = m.get('AccessControlAllowOrigin')

        if m.get('AccessControlRequestMethod') is not None:
            self.access_control_request_method = m.get('AccessControlRequestMethod')

        if m.get('Addresses') is not None:
            temp_model = main_models.MetaQueryFileAddresses()
            self.addresses = temp_model.from_map(m.get('Addresses'))

        if m.get('Album') is not None:
            self.album = m.get('Album')

        if m.get('AlbumArtist') is not None:
            self.album_artist = m.get('AlbumArtist')

        if m.get('Artist') is not None:
            self.artist = m.get('Artist')

        if m.get('AudioStreams') is not None:
            temp_model = main_models.MetaQueryFileAudioStreams()
            self.audio_streams = temp_model.from_map(m.get('AudioStreams'))

        if m.get('Bitrate') is not None:
            self.bitrate = m.get('Bitrate')

        if m.get('CacheControl') is not None:
            self.cache_control = m.get('CacheControl')

        if m.get('Composer') is not None:
            self.composer = m.get('Composer')

        if m.get('ContentDisposition') is not None:
            self.content_disposition = m.get('ContentDisposition')

        if m.get('ContentEncoding') is not None:
            self.content_encoding = m.get('ContentEncoding')

        if m.get('ContentLanguage') is not None:
            self.content_language = m.get('ContentLanguage')

        if m.get('ContentType') is not None:
            self.content_type = m.get('ContentType')

        if m.get('Duration') is not None:
            self.duration = m.get('Duration')

        if m.get('ETag') is not None:
            self.etag = m.get('ETag')

        if m.get('FileModifiedTime') is not None:
            self.file_modified_time = m.get('FileModifiedTime')

        if m.get('Filename') is not None:
            self.filename = m.get('Filename')

        if m.get('ImageHeight') is not None:
            self.image_height = m.get('ImageHeight')

        if m.get('ImageWidth') is not None:
            self.image_width = m.get('ImageWidth')

        if m.get('LatLong') is not None:
            self.lat_long = m.get('LatLong')

        if m.get('MediaType') is not None:
            self.media_type = m.get('MediaType')

        if m.get('OSSCRC64') is not None:
            self.osscrc64 = m.get('OSSCRC64')

        if m.get('OSSExpiration') is not None:
            self.ossexpiration = m.get('OSSExpiration')

        if m.get('OSSObjectType') is not None:
            self.ossobject_type = m.get('OSSObjectType')

        if m.get('OSSStorageClass') is not None:
            self.ossstorage_class = m.get('OSSStorageClass')

        if m.get('OSSTagging') is not None:
            temp_model = main_models.MetaQueryFileOSSTagging()
            self.osstagging = temp_model.from_map(m.get('OSSTagging'))

        if m.get('OSSTaggingCount') is not None:
            self.osstagging_count = m.get('OSSTaggingCount')

        if m.get('OSSUserMeta') is not None:
            temp_model = main_models.MetaQueryFileOSSUserMeta()
            self.ossuser_meta = temp_model.from_map(m.get('OSSUserMeta'))

        if m.get('ObjectACL') is not None:
            self.object_acl = m.get('ObjectACL')

        if m.get('Performer') is not None:
            self.performer = m.get('Performer')

        if m.get('ProduceTime') is not None:
            self.produce_time = m.get('ProduceTime')

        if m.get('ServerSideDataEncryption') is not None:
            self.server_side_data_encryption = m.get('ServerSideDataEncryption')

        if m.get('ServerSideEncryption') is not None:
            self.server_side_encryption = m.get('ServerSideEncryption')

        if m.get('ServerSideEncryptionCustomerAlgorithm') is not None:
            self.server_side_encryption_customer_algorithm = m.get('ServerSideEncryptionCustomerAlgorithm')

        if m.get('ServerSideEncryptionKeyId') is not None:
            self.server_side_encryption_key_id = m.get('ServerSideEncryptionKeyId')

        if m.get('Size') is not None:
            self.size = m.get('Size')

        if m.get('Subtitles') is not None:
            temp_model = main_models.MetaQueryFileSubtitles()
            self.subtitles = temp_model.from_map(m.get('Subtitles'))

        if m.get('Title') is not None:
            self.title = m.get('Title')

        if m.get('URI') is not None:
            self.uri = m.get('URI')

        if m.get('VideoHeight') is not None:
            self.video_height = m.get('VideoHeight')

        if m.get('VideoStreams') is not None:
            temp_model = main_models.MetaQueryFileVideoStreams()
            self.video_streams = temp_model.from_map(m.get('VideoStreams'))

        if m.get('VideoWidth') is not None:
            self.video_width = m.get('VideoWidth')

        return self

class MetaQueryFileVideoStreams(DaraModel):
    def __init__(
        self,
        video_stream: main_models.MetaQueryRespVideoStream = None,
    ):
        self.video_stream = video_stream

    def validate(self):
        if self.video_stream:
            self.video_stream.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.video_stream is not None:
            result['VideoStream'] = self.video_stream.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('VideoStream') is not None:
            temp_model = main_models.MetaQueryRespVideoStream()
            self.video_stream = temp_model.from_map(m.get('VideoStream'))

        return self

class MetaQueryFileSubtitles(DaraModel):
    def __init__(
        self,
        subtitle: main_models.MetaQueryRespSubtitle = None,
    ):
        self.subtitle = subtitle

    def validate(self):
        if self.subtitle:
            self.subtitle.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.subtitle is not None:
            result['Subtitle'] = self.subtitle.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Subtitle') is not None:
            temp_model = main_models.MetaQueryRespSubtitle()
            self.subtitle = temp_model.from_map(m.get('Subtitle'))

        return self

class MetaQueryFileOSSUserMeta(DaraModel):
    def __init__(
        self,
        user_meta: List[main_models.MetaQueryUserMeta] = None,
    ):
        self.user_meta = user_meta

    def validate(self):
        if self.user_meta:
            for v1 in self.user_meta:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['UserMeta'] = []
        if self.user_meta is not None:
            for k1 in self.user_meta:
                result['UserMeta'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.user_meta = []
        if m.get('UserMeta') is not None:
            for k1 in m.get('UserMeta'):
                temp_model = main_models.MetaQueryUserMeta()
                self.user_meta.append(temp_model.from_map(k1))

        return self

class MetaQueryFileOSSTagging(DaraModel):
    def __init__(
        self,
        tagging: List[main_models.MetaQueryTagging] = None,
    ):
        self.tagging = tagging

    def validate(self):
        if self.tagging:
            for v1 in self.tagging:
                 if v1:
                    v1.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        result['Tagging'] = []
        if self.tagging is not None:
            for k1 in self.tagging:
                result['Tagging'].append(k1.to_map() if k1 else None)

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        self.tagging = []
        if m.get('Tagging') is not None:
            for k1 in m.get('Tagging'):
                temp_model = main_models.MetaQueryTagging()
                self.tagging.append(temp_model.from_map(k1))

        return self

class MetaQueryFileAudioStreams(DaraModel):
    def __init__(
        self,
        audio_stream: main_models.MetaQueryRespAudioStream = None,
    ):
        self.audio_stream = audio_stream

    def validate(self):
        if self.audio_stream:
            self.audio_stream.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.audio_stream is not None:
            result['AudioStream'] = self.audio_stream.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('AudioStream') is not None:
            temp_model = main_models.MetaQueryRespAudioStream()
            self.audio_stream = temp_model.from_map(m.get('AudioStream'))

        return self

class MetaQueryFileAddresses(DaraModel):
    def __init__(
        self,
        address: main_models.MetaQueryRespAddress = None,
    ):
        self.address = address

    def validate(self):
        if self.address:
            self.address.validate()

    def to_map(self):
        result = dict()
        _map = super().to_map()
        if _map is not None:
            result = _map
        if self.address is not None:
            result['Address'] = self.address.to_map()

        return result

    def from_map(self, m: dict = None):
        m = m or dict()
        if m.get('Address') is not None:
            temp_model = main_models.MetaQueryRespAddress()
            self.address = temp_model.from_map(m.get('Address'))

        return self

