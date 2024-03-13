// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetObjectTaggingResponseBody extends TeaModel {
    /**
     * <p>The container that stores the returned tag of the bucket.</p>
     */
    @NameInMap("Tagging")
    public Tagging tagging;

    public static GetObjectTaggingResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetObjectTaggingResponseBody self = new GetObjectTaggingResponseBody();
        return TeaModel.build(map, self);
    }

    public GetObjectTaggingResponseBody setTagging(Tagging tagging) {
        this.tagging = tagging;
        return this;
    }
    public Tagging getTagging() {
        return this.tagging;
    }

    public static class Tagging extends TeaModel {
        /**
         * <p>The tag set of the target object.</p>
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
