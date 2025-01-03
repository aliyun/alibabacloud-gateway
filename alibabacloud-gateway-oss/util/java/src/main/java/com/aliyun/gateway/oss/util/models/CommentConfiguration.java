// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CommentConfiguration extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>comments</p>
     */
    @NameInMap("Comment")
    public String comment;

    public static CommentConfiguration build(java.util.Map<String, ?> map) throws Exception {
        CommentConfiguration self = new CommentConfiguration();
        return TeaModel.build(map, self);
    }

    public CommentConfiguration setComment(String comment) {
        this.comment = comment;
        return this;
    }
    public String getComment() {
        return this.comment;
    }

}
