// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class NotificationConfiguration extends TeaModel {
    @NameInMap("TopicConfiguration")
    public java.util.List<TopicConfiguration> topicConfiguration;

    public static NotificationConfiguration build(java.util.Map<String, ?> map) throws Exception {
        NotificationConfiguration self = new NotificationConfiguration();
        return TeaModel.build(map, self);
    }

    public NotificationConfiguration setTopicConfiguration(java.util.List<TopicConfiguration> topicConfiguration) {
        this.topicConfiguration = topicConfiguration;
        return this;
    }
    public java.util.List<TopicConfiguration> getTopicConfiguration() {
        return this.topicConfiguration;
    }

    public static class TopicConfiguration extends TeaModel {
        /**
         * <strong>example:</strong>
         * <p>test-topic-1</p>
         */
        @NameInMap("Id")
        public String id;

        public static TopicConfiguration build(java.util.Map<String, ?> map) throws Exception {
            TopicConfiguration self = new TopicConfiguration();
            return TeaModel.build(map, self);
        }

        public TopicConfiguration setId(String id) {
            this.id = id;
            return this;
        }
        public String getId() {
            return this.id;
        }

    }

}
