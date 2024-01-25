// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetBucketTagsResponseBody extends TeaModel {
    /**
     * <p>The container that stores the returned tags of the bucket.</p>
     */
    @NameInMap("TagSet")
    public TagSet tagSet;

    public static GetBucketTagsResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketTagsResponseBody self = new GetBucketTagsResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketTagsResponseBody setTagSet(TagSet tagSet) {
        this.tagSet = tagSet;
        return this;
    }
    public TagSet getTagSet() {
        return this.tagSet;
    }

}
