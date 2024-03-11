// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class GetBucketTagsResponseBody extends TeaModel {
    /**
     * <p>The container that stores the returned tags of the bucket.</p>
     * <p>> If no tags are configured for the bucket, an XML message body is returned in which the Tagging element is empty.</p>
     */
    @NameInMap("Tagging")
    public Tagging tagging;

    public static GetBucketTagsResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketTagsResponseBody self = new GetBucketTagsResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketTagsResponseBody setTagging(Tagging tagging) {
        this.tagging = tagging;
        return this;
    }
    public Tagging getTagging() {
        return this.tagging;
    }

    public static class Tagging extends TeaModel {
        /**
         * <p>The container that stores the returned tags of the bucket.</p>
         */
        @NameInMap("TagSet")
        public TagSet tagSet;

        public static Tagging build(java.util.Map<String, ?> map) throws Exception {
            Tagging self = new Tagging();
            return TeaModel.build(map, self);
        }

        public Tagging setTagSet(TagSet tagSet) {
            this.tagSet = tagSet;
            return this;
        }
        public TagSet getTagSet() {
            return this.tagSet;
        }

    }

}
