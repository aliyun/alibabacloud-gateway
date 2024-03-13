// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutLiveChannelResponseBody extends TeaModel {
    /**
     * <p>The container that stores the result of the CreateLiveChannel request.</p>
     */
    @NameInMap("CreateLiveChannelResult")
    public CreateLiveChannelResult createLiveChannelResult;

    public static PutLiveChannelResponseBody build(java.util.Map<String, ?> map) throws Exception {
        PutLiveChannelResponseBody self = new PutLiveChannelResponseBody();
        return TeaModel.build(map, self);
    }

    public PutLiveChannelResponseBody setCreateLiveChannelResult(CreateLiveChannelResult createLiveChannelResult) {
        this.createLiveChannelResult = createLiveChannelResult;
        return this;
    }
    public CreateLiveChannelResult getCreateLiveChannelResult() {
        return this.createLiveChannelResult;
    }

    public static class CreateLiveChannelResult extends TeaModel {
        /**
         * <p>保存播放地址的容器。</p>
         */
        @NameInMap("PlayUrls")
        public LiveChannelPlayUrls playUrls;

        /**
         * <p>保存推流地址的容器。</p>
         */
        @NameInMap("PublishUrls")
        public LiveChannelPublishUrls publishUrls;

        public static CreateLiveChannelResult build(java.util.Map<String, ?> map) throws Exception {
            CreateLiveChannelResult self = new CreateLiveChannelResult();
            return TeaModel.build(map, self);
        }

        public CreateLiveChannelResult setPlayUrls(LiveChannelPlayUrls playUrls) {
            this.playUrls = playUrls;
            return this;
        }
        public LiveChannelPlayUrls getPlayUrls() {
            return this.playUrls;
        }

        public CreateLiveChannelResult setPublishUrls(LiveChannelPublishUrls publishUrls) {
            this.publishUrls = publishUrls;
            return this;
        }
        public LiveChannelPublishUrls getPublishUrls() {
            return this.publishUrls;
        }

    }

}
