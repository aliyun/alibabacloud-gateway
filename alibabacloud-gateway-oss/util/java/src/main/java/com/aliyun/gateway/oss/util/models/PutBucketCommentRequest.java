// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutBucketCommentRequest extends TeaModel {
    @NameInMap("CommentConfiguration")
    public CommentConfiguration commentConfiguration;

    public static PutBucketCommentRequest build(java.util.Map<String, ?> map) throws Exception {
        PutBucketCommentRequest self = new PutBucketCommentRequest();
        return TeaModel.build(map, self);
    }

    public PutBucketCommentRequest setCommentConfiguration(CommentConfiguration commentConfiguration) {
        this.commentConfiguration = commentConfiguration;
        return this;
    }
    public CommentConfiguration getCommentConfiguration() {
        return this.commentConfiguration;
    }

}
