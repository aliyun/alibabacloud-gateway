// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class DeletedObject extends TeaModel {
    @NameInMap("DeleteMarker")
    public Boolean deleteMarker;

    @NameInMap("DeleteMarkerVersionId")
    public String deleteMarkerVersionId;

    @NameInMap("Key")
    public String key;

    @NameInMap("VersionId")
    public String versionId;

    public static DeletedObject build(java.util.Map<String, ?> map) throws Exception {
        DeletedObject self = new DeletedObject();
        return TeaModel.build(map, self);
    }

    public DeletedObject setDeleteMarker(Boolean deleteMarker) {
        this.deleteMarker = deleteMarker;
        return this;
    }
    public Boolean getDeleteMarker() {
        return this.deleteMarker;
    }

    public DeletedObject setDeleteMarkerVersionId(String deleteMarkerVersionId) {
        this.deleteMarkerVersionId = deleteMarkerVersionId;
        return this;
    }
    public String getDeleteMarkerVersionId() {
        return this.deleteMarkerVersionId;
    }

    public DeletedObject setKey(String key) {
        this.key = key;
        return this;
    }
    public String getKey() {
        return this.key;
    }

    public DeletedObject setVersionId(String versionId) {
        this.versionId = versionId;
        return this;
    }
    public String getVersionId() {
        return this.versionId;
    }

}
