// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class Tagging extends TeaModel {
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
