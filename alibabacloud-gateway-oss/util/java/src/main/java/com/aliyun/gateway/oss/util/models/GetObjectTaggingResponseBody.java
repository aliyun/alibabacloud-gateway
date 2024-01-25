// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetObjectTaggingResponseBody extends TeaModel {
    /**
     * <p>The collection of tags.</p>
     */
    @NameInMap("TagSet")
    public TagSet tagSet;

    public static GetObjectTaggingResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetObjectTaggingResponseBody self = new GetObjectTaggingResponseBody();
        return TeaModel.build(map, self);
    }

    public GetObjectTaggingResponseBody setTagSet(TagSet tagSet) {
        this.tagSet = tagSet;
        return this;
    }
    public TagSet getTagSet() {
        return this.tagSet;
    }

}
