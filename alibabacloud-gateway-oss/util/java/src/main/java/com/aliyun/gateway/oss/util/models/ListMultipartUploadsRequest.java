// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class ListMultipartUploadsRequest extends TeaModel {
    /**
     * <p>The character used to group objects by name. Objects whose names contain the same string that ranges from the specified prefix to the delimiter that appears for the first time are grouped as a CommonPrefixes element.</p>
     */
    @NameInMap("delimiter")
    public String delimiter;

    /**
     * <p>The encoding type of the object name in the response. Values of Delimiter, KeyMarker, Prefix, NextKeyMarker, and Key can be encoded in UTF-8. However, the XML 1.0 standard cannot be used to parse control characters such as characters with an American Standard Code for Information Interchange (ASCII) value from 0 to 10. You can set the encoding-type parameter to encode values of Delimiter, KeyMarker, Prefix, NextKeyMarker, and Key in the response.</p>
     * <p><br>Default value: null</p>
     */
    @NameInMap("encoding-type")
    public String encodingType;

    /**
     * <p>This parameter is used together with the upload-id-marker parameter to specify the position from which the returned results start. </p>
     * <p>  - If the upload-id-marker parameter is not specified, OSS returns all multipart upload tasks in which object names are alphabetically after the value of the key-marker parameter. </p>
     * <p>  - If the upload-id-marker parameter is specified, the response includes the following tasks:</p>
     * <p>    - Multipart upload tasks in which object names are alphabetically after the value of the key-marker parameter. </p>
     * <p>    - Multipart upload tasks in which object names are the same as the value of the key-marker parameter but whose upload IDs are greater than the value of the upload-id-marker parameter.</p>
     */
    @NameInMap("key-marker")
    public String keyMarker;

    /**
     * <p>The maximum number of multipart upload tasks that can be returned for the current request. Default value: 1000. Maximum value: 1000.</p>
     */
    @NameInMap("max-uploads")
    public Long maxUploads;

    /**
     * <p>The prefix that the names of the objects that you want to return must contain. If you specify a prefix in the request, the specified prefix is included in the response. </p>
     * <p>>  You can use prefixes to group and manage objects in buckets in the same way you manage a folder in a file system.</p>
     */
    @NameInMap("prefix")
    public String prefix;

    /**
     * <p>The upload ID of the multipart upload task after which the list operation starts. This parameter is used together with the key-marker parameter. </p>
     * <p>  - If the key-marker parameter is not specified, OSS ignores the upload-id-marker parameter. </p>
     * <p>  - If the key-marker parameter is specified, the query result includes the following tasks:</p>
     * <p>    - Multipart upload tasks in which object names are alphabetically after the value of the key-marker parameter. </p>
     * <p>    - Multipart upload tasks in which object names are the same as the value of the key-marker parameter but whose upload IDs are greater than the value of the upload-id-marker parameter.</p>
     */
    @NameInMap("upload-id-marker")
    public String uploadIdMarker;

    public static ListMultipartUploadsRequest build(java.util.Map<String, ?> map) throws Exception {
        ListMultipartUploadsRequest self = new ListMultipartUploadsRequest();
        return TeaModel.build(map, self);
    }

    public ListMultipartUploadsRequest setDelimiter(String delimiter) {
        this.delimiter = delimiter;
        return this;
    }
    public String getDelimiter() {
        return this.delimiter;
    }

    public ListMultipartUploadsRequest setEncodingType(String encodingType) {
        this.encodingType = encodingType;
        return this;
    }
    public String getEncodingType() {
        return this.encodingType;
    }

    public ListMultipartUploadsRequest setKeyMarker(String keyMarker) {
        this.keyMarker = keyMarker;
        return this;
    }
    public String getKeyMarker() {
        return this.keyMarker;
    }

    public ListMultipartUploadsRequest setMaxUploads(Long maxUploads) {
        this.maxUploads = maxUploads;
        return this;
    }
    public Long getMaxUploads() {
        return this.maxUploads;
    }

    public ListMultipartUploadsRequest setPrefix(String prefix) {
        this.prefix = prefix;
        return this;
    }
    public String getPrefix() {
        return this.prefix;
    }

    public ListMultipartUploadsRequest setUploadIdMarker(String uploadIdMarker) {
        this.uploadIdMarker = uploadIdMarker;
        return this;
    }
    public String getUploadIdMarker() {
        return this.uploadIdMarker;
    }

}
