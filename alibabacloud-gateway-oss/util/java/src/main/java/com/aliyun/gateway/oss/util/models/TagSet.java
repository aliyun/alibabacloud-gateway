// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class TagSet extends TeaModel {
    @NameInMap("Tag")
    public java.util.List<Tag> tag;

    public static TagSet build(java.util.Map<String, ?> map) throws Exception {
        TagSet self = new TagSet();
        return TeaModel.build(map, self);
    }

    public TagSet setTag(java.util.List<Tag> tag) {
        this.tag = tag;
        return this;
    }
    public java.util.List<Tag> getTag() {
        return this.tag;
    }

}
