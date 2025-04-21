// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class MetaQueryFile extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p><a href="https://aliyundoc.com">https://aliyundoc.com</a></p>
     */
    @NameInMap("AccessControlAllowOrigin")
    public String accessControlAllowOrigin;

    /**
     * <strong>example:</strong>
     * <p>PUT</p>
     */
    @NameInMap("AccessControlRequestMethod")
    public String accessControlRequestMethod;

    @NameInMap("Addresses")
    public Addresses addresses;

    /**
     * <strong>example:</strong>
     * <p>FirstAlbum</p>
     */
    @NameInMap("Album")
    public String album;

    /**
     * <strong>example:</strong>
     * <p>Jenny</p>
     */
    @NameInMap("AlbumArtist")
    public String albumArtist;

    /**
     * <strong>example:</strong>
     * <p>Jane</p>
     */
    @NameInMap("Artist")
    public String artist;

    @NameInMap("AudioStreams")
    public AudioStreams audioStreams;

    /**
     * <strong>example:</strong>
     * <p>13091201</p>
     */
    @NameInMap("Bitrate")
    public Long bitrate;

    /**
     * <strong>example:</strong>
     * <p>no-cache</p>
     */
    @NameInMap("CacheControl")
    public String cacheControl;

    /**
     * <strong>example:</strong>
     * <p>Jane</p>
     */
    @NameInMap("Composer")
    public String composer;

    /**
     * <strong>example:</strong>
     * <p>attachment; filename =test.jpg</p>
     */
    @NameInMap("ContentDisposition")
    public String contentDisposition;

    /**
     * <strong>example:</strong>
     * <p>UTF-8</p>
     */
    @NameInMap("ContentEncoding")
    public String contentEncoding;

    /**
     * <strong>example:</strong>
     * <p>zh-CN</p>
     */
    @NameInMap("ContentLanguage")
    public String contentLanguage;

    /**
     * <strong>example:</strong>
     * <p>image/jpeg</p>
     */
    @NameInMap("ContentType")
    public String contentType;

    /**
     * <strong>example:</strong>
     * <p>15.263000</p>
     */
    @NameInMap("Duration")
    public Double duration;

    /**
     * <strong>example:</strong>
     * <p>&quot;fba9dede5f27731c9771645a3986****&quot;</p>
     */
    @NameInMap("ETag")
    public String ETag;

    /**
     * <strong>example:</strong>
     * <p>2021-06-29T15:04:05.000000000Z07:00</p>
     */
    @NameInMap("FileModifiedTime")
    public String fileModifiedTime;

    /**
     * <strong>example:</strong>
     * <p>exampleobject.txt</p>
     */
    @NameInMap("Filename")
    public String filename;

    /**
     * <strong>example:</strong>
     * <p>500</p>
     */
    @NameInMap("ImageHeight")
    public Long imageHeight;

    /**
     * <strong>example:</strong>
     * <p>270</p>
     */
    @NameInMap("ImageWidth")
    public Long imageWidth;

    /**
     * <strong>example:</strong>
     * <p>30.134390,120.074997</p>
     */
    @NameInMap("LatLong")
    public String latLong;

    /**
     * <strong>example:</strong>
     * <p>image</p>
     */
    @NameInMap("MediaType")
    public String mediaType;

    /**
     * <strong>example:</strong>
     * <p>4858A48BD1466884</p>
     */
    @NameInMap("OSSCRC64")
    public String OSSCRC64;

    /**
     * <strong>example:</strong>
     * <p>2124-12-01T12:00:00.000Z</p>
     */
    @NameInMap("OSSExpiration")
    public String OSSExpiration;

    /**
     * <strong>example:</strong>
     * <p>Normal</p>
     */
    @NameInMap("OSSObjectType")
    public String OSSObjectType;

    /**
     * <strong>example:</strong>
     * <p>Standard</p>
     */
    @NameInMap("OSSStorageClass")
    public String OSSStorageClass;

    @NameInMap("OSSTagging")
    public OSSTagging OSSTagging;

    /**
     * <strong>example:</strong>
     * <p>2</p>
     */
    @NameInMap("OSSTaggingCount")
    public Long OSSTaggingCount;

    @NameInMap("OSSUserMeta")
    public OSSUserMeta OSSUserMeta;

    /**
     * <strong>example:</strong>
     * <p>default</p>
     */
    @NameInMap("ObjectACL")
    public String objectACL;

    /**
     * <strong>example:</strong>
     * <p>Jane</p>
     */
    @NameInMap("Performer")
    public String performer;

    /**
     * <strong>example:</strong>
     * <p>2021-06-29T14:50:13.011643661+08:00</p>
     */
    @NameInMap("ProduceTime")
    public String produceTime;

    /**
     * <strong>example:</strong>
     * <p>SM4</p>
     */
    @NameInMap("ServerSideDataEncryption")
    public String serverSideDataEncryption;

    /**
     * <strong>example:</strong>
     * <p>AES256</p>
     */
    @NameInMap("ServerSideEncryption")
    public String serverSideEncryption;

    /**
     * <strong>example:</strong>
     * <p>SM4</p>
     */
    @NameInMap("ServerSideEncryptionCustomerAlgorithm")
    public String serverSideEncryptionCustomerAlgorithm;

    /**
     * <strong>example:</strong>
     * <p>9468da86-3509-4f8d-a61e-6eab1eac****</p>
     */
    @NameInMap("ServerSideEncryptionKeyId")
    public String serverSideEncryptionKeyId;

    /**
     * <strong>example:</strong>
     * <p>120</p>
     */
    @NameInMap("Size")
    public Long size;

    @NameInMap("Subtitles")
    public Subtitles subtitles;

    /**
     * <strong>example:</strong>
     * <p>test</p>
     */
    @NameInMap("Title")
    public String title;

    /**
     * <strong>example:</strong>
     * <p>oss://examplebucket/test-object.jpg</p>
     */
    @NameInMap("URI")
    public String URI;

    /**
     * <strong>example:</strong>
     * <p>1920</p>
     */
    @NameInMap("VideoHeight")
    public Long videoHeight;

    @NameInMap("VideoStreams")
    public VideoStreams videoStreams;

    /**
     * <strong>example:</strong>
     * <p>1080</p>
     */
    @NameInMap("VideoWidth")
    public Long videoWidth;

    public static MetaQueryFile build(java.util.Map<String, ?> map) throws Exception {
        MetaQueryFile self = new MetaQueryFile();
        return TeaModel.build(map, self);
    }

    public MetaQueryFile setAccessControlAllowOrigin(String accessControlAllowOrigin) {
        this.accessControlAllowOrigin = accessControlAllowOrigin;
        return this;
    }
    public String getAccessControlAllowOrigin() {
        return this.accessControlAllowOrigin;
    }

    public MetaQueryFile setAccessControlRequestMethod(String accessControlRequestMethod) {
        this.accessControlRequestMethod = accessControlRequestMethod;
        return this;
    }
    public String getAccessControlRequestMethod() {
        return this.accessControlRequestMethod;
    }

    public MetaQueryFile setAddresses(Addresses addresses) {
        this.addresses = addresses;
        return this;
    }
    public Addresses getAddresses() {
        return this.addresses;
    }

    public MetaQueryFile setAlbum(String album) {
        this.album = album;
        return this;
    }
    public String getAlbum() {
        return this.album;
    }

    public MetaQueryFile setAlbumArtist(String albumArtist) {
        this.albumArtist = albumArtist;
        return this;
    }
    public String getAlbumArtist() {
        return this.albumArtist;
    }

    public MetaQueryFile setArtist(String artist) {
        this.artist = artist;
        return this;
    }
    public String getArtist() {
        return this.artist;
    }

    public MetaQueryFile setAudioStreams(AudioStreams audioStreams) {
        this.audioStreams = audioStreams;
        return this;
    }
    public AudioStreams getAudioStreams() {
        return this.audioStreams;
    }

    public MetaQueryFile setBitrate(Long bitrate) {
        this.bitrate = bitrate;
        return this;
    }
    public Long getBitrate() {
        return this.bitrate;
    }

    public MetaQueryFile setCacheControl(String cacheControl) {
        this.cacheControl = cacheControl;
        return this;
    }
    public String getCacheControl() {
        return this.cacheControl;
    }

    public MetaQueryFile setComposer(String composer) {
        this.composer = composer;
        return this;
    }
    public String getComposer() {
        return this.composer;
    }

    public MetaQueryFile setContentDisposition(String contentDisposition) {
        this.contentDisposition = contentDisposition;
        return this;
    }
    public String getContentDisposition() {
        return this.contentDisposition;
    }

    public MetaQueryFile setContentEncoding(String contentEncoding) {
        this.contentEncoding = contentEncoding;
        return this;
    }
    public String getContentEncoding() {
        return this.contentEncoding;
    }

    public MetaQueryFile setContentLanguage(String contentLanguage) {
        this.contentLanguage = contentLanguage;
        return this;
    }
    public String getContentLanguage() {
        return this.contentLanguage;
    }

    public MetaQueryFile setContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    public String getContentType() {
        return this.contentType;
    }

    public MetaQueryFile setDuration(Double duration) {
        this.duration = duration;
        return this;
    }
    public Double getDuration() {
        return this.duration;
    }

    public MetaQueryFile setETag(String ETag) {
        this.ETag = ETag;
        return this;
    }
    public String getETag() {
        return this.ETag;
    }

    public MetaQueryFile setFileModifiedTime(String fileModifiedTime) {
        this.fileModifiedTime = fileModifiedTime;
        return this;
    }
    public String getFileModifiedTime() {
        return this.fileModifiedTime;
    }

    public MetaQueryFile setFilename(String filename) {
        this.filename = filename;
        return this;
    }
    public String getFilename() {
        return this.filename;
    }

    public MetaQueryFile setImageHeight(Long imageHeight) {
        this.imageHeight = imageHeight;
        return this;
    }
    public Long getImageHeight() {
        return this.imageHeight;
    }

    public MetaQueryFile setImageWidth(Long imageWidth) {
        this.imageWidth = imageWidth;
        return this;
    }
    public Long getImageWidth() {
        return this.imageWidth;
    }

    public MetaQueryFile setLatLong(String latLong) {
        this.latLong = latLong;
        return this;
    }
    public String getLatLong() {
        return this.latLong;
    }

    public MetaQueryFile setMediaType(String mediaType) {
        this.mediaType = mediaType;
        return this;
    }
    public String getMediaType() {
        return this.mediaType;
    }

    public MetaQueryFile setOSSCRC64(String OSSCRC64) {
        this.OSSCRC64 = OSSCRC64;
        return this;
    }
    public String getOSSCRC64() {
        return this.OSSCRC64;
    }

    public MetaQueryFile setOSSExpiration(String OSSExpiration) {
        this.OSSExpiration = OSSExpiration;
        return this;
    }
    public String getOSSExpiration() {
        return this.OSSExpiration;
    }

    public MetaQueryFile setOSSObjectType(String OSSObjectType) {
        this.OSSObjectType = OSSObjectType;
        return this;
    }
    public String getOSSObjectType() {
        return this.OSSObjectType;
    }

    public MetaQueryFile setOSSStorageClass(String OSSStorageClass) {
        this.OSSStorageClass = OSSStorageClass;
        return this;
    }
    public String getOSSStorageClass() {
        return this.OSSStorageClass;
    }

    public MetaQueryFile setOSSTagging(OSSTagging OSSTagging) {
        this.OSSTagging = OSSTagging;
        return this;
    }
    public OSSTagging getOSSTagging() {
        return this.OSSTagging;
    }

    public MetaQueryFile setOSSTaggingCount(Long OSSTaggingCount) {
        this.OSSTaggingCount = OSSTaggingCount;
        return this;
    }
    public Long getOSSTaggingCount() {
        return this.OSSTaggingCount;
    }

    public MetaQueryFile setOSSUserMeta(OSSUserMeta OSSUserMeta) {
        this.OSSUserMeta = OSSUserMeta;
        return this;
    }
    public OSSUserMeta getOSSUserMeta() {
        return this.OSSUserMeta;
    }

    public MetaQueryFile setObjectACL(String objectACL) {
        this.objectACL = objectACL;
        return this;
    }
    public String getObjectACL() {
        return this.objectACL;
    }

    public MetaQueryFile setPerformer(String performer) {
        this.performer = performer;
        return this;
    }
    public String getPerformer() {
        return this.performer;
    }

    public MetaQueryFile setProduceTime(String produceTime) {
        this.produceTime = produceTime;
        return this;
    }
    public String getProduceTime() {
        return this.produceTime;
    }

    public MetaQueryFile setServerSideDataEncryption(String serverSideDataEncryption) {
        this.serverSideDataEncryption = serverSideDataEncryption;
        return this;
    }
    public String getServerSideDataEncryption() {
        return this.serverSideDataEncryption;
    }

    public MetaQueryFile setServerSideEncryption(String serverSideEncryption) {
        this.serverSideEncryption = serverSideEncryption;
        return this;
    }
    public String getServerSideEncryption() {
        return this.serverSideEncryption;
    }

    public MetaQueryFile setServerSideEncryptionCustomerAlgorithm(String serverSideEncryptionCustomerAlgorithm) {
        this.serverSideEncryptionCustomerAlgorithm = serverSideEncryptionCustomerAlgorithm;
        return this;
    }
    public String getServerSideEncryptionCustomerAlgorithm() {
        return this.serverSideEncryptionCustomerAlgorithm;
    }

    public MetaQueryFile setServerSideEncryptionKeyId(String serverSideEncryptionKeyId) {
        this.serverSideEncryptionKeyId = serverSideEncryptionKeyId;
        return this;
    }
    public String getServerSideEncryptionKeyId() {
        return this.serverSideEncryptionKeyId;
    }

    public MetaQueryFile setSize(Long size) {
        this.size = size;
        return this;
    }
    public Long getSize() {
        return this.size;
    }

    public MetaQueryFile setSubtitles(Subtitles subtitles) {
        this.subtitles = subtitles;
        return this;
    }
    public Subtitles getSubtitles() {
        return this.subtitles;
    }

    public MetaQueryFile setTitle(String title) {
        this.title = title;
        return this;
    }
    public String getTitle() {
        return this.title;
    }

    public MetaQueryFile setURI(String URI) {
        this.URI = URI;
        return this;
    }
    public String getURI() {
        return this.URI;
    }

    public MetaQueryFile setVideoHeight(Long videoHeight) {
        this.videoHeight = videoHeight;
        return this;
    }
    public Long getVideoHeight() {
        return this.videoHeight;
    }

    public MetaQueryFile setVideoStreams(VideoStreams videoStreams) {
        this.videoStreams = videoStreams;
        return this;
    }
    public VideoStreams getVideoStreams() {
        return this.videoStreams;
    }

    public MetaQueryFile setVideoWidth(Long videoWidth) {
        this.videoWidth = videoWidth;
        return this;
    }
    public Long getVideoWidth() {
        return this.videoWidth;
    }

    public static class Addresses extends TeaModel {
        @NameInMap("Address")
        public MetaQueryRespAddress address;

        public static Addresses build(java.util.Map<String, ?> map) throws Exception {
            Addresses self = new Addresses();
            return TeaModel.build(map, self);
        }

        public Addresses setAddress(MetaQueryRespAddress address) {
            this.address = address;
            return this;
        }
        public MetaQueryRespAddress getAddress() {
            return this.address;
        }

    }

    public static class AudioStreams extends TeaModel {
        @NameInMap("AudioStream")
        public MetaQueryRespAudioStream audioStream;

        public static AudioStreams build(java.util.Map<String, ?> map) throws Exception {
            AudioStreams self = new AudioStreams();
            return TeaModel.build(map, self);
        }

        public AudioStreams setAudioStream(MetaQueryRespAudioStream audioStream) {
            this.audioStream = audioStream;
            return this;
        }
        public MetaQueryRespAudioStream getAudioStream() {
            return this.audioStream;
        }

    }

    public static class OSSTagging extends TeaModel {
        @NameInMap("Tagging")
        public java.util.List<MetaQueryTagging> tagging;

        public static OSSTagging build(java.util.Map<String, ?> map) throws Exception {
            OSSTagging self = new OSSTagging();
            return TeaModel.build(map, self);
        }

        public OSSTagging setTagging(java.util.List<MetaQueryTagging> tagging) {
            this.tagging = tagging;
            return this;
        }
        public java.util.List<MetaQueryTagging> getTagging() {
            return this.tagging;
        }

    }

    public static class OSSUserMeta extends TeaModel {
        @NameInMap("UserMeta")
        public java.util.List<MetaQueryUserMeta> userMeta;

        public static OSSUserMeta build(java.util.Map<String, ?> map) throws Exception {
            OSSUserMeta self = new OSSUserMeta();
            return TeaModel.build(map, self);
        }

        public OSSUserMeta setUserMeta(java.util.List<MetaQueryUserMeta> userMeta) {
            this.userMeta = userMeta;
            return this;
        }
        public java.util.List<MetaQueryUserMeta> getUserMeta() {
            return this.userMeta;
        }

    }

    public static class Subtitles extends TeaModel {
        @NameInMap("Subtitle")
        public MetaQueryRespSubtitle subtitle;

        public static Subtitles build(java.util.Map<String, ?> map) throws Exception {
            Subtitles self = new Subtitles();
            return TeaModel.build(map, self);
        }

        public Subtitles setSubtitle(MetaQueryRespSubtitle subtitle) {
            this.subtitle = subtitle;
            return this;
        }
        public MetaQueryRespSubtitle getSubtitle() {
            return this.subtitle;
        }

    }

    public static class VideoStreams extends TeaModel {
        @NameInMap("VideoStream")
        public MetaQueryRespVideoStream videoStream;

        public static VideoStreams build(java.util.Map<String, ?> map) throws Exception {
            VideoStreams self = new VideoStreams();
            return TeaModel.build(map, self);
        }

        public VideoStreams setVideoStream(MetaQueryRespVideoStream videoStream) {
            this.videoStream = videoStream;
            return this;
        }
        public MetaQueryRespVideoStream getVideoStream() {
            return this.videoStream;
        }

    }

}
