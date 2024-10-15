package com.aliyun.gateway.sls.util.model;

import com.aliyun.tea.utils.StringUtils;

import java.io.Serializable;
import java.util.List;

public class PostLogStoreLogsRequest implements Serializable {
    private static final long serialVersionUID = 7226856831224917838L;
    private String topic;
    private String source;
    private List<LogItem> logItems;
    private List<LogTag> tags = null;
    private String compressType = "lz4";

    // POST /logstores/{logStore}/shards/lb
    // content-type: application/x-protobuf
    // extra header:
    //      x-log-compresstype: {compressType}
    //      x-log-bodyrawsize: {rawSize}
    public PostLogStoreLogsRequest(String topic, String source, List<LogItem> logItems, List<LogTag> tags) {
        this.topic = topic;
        this.source = source;
        this.logItems = logItems;
        this.tags = tags;
    }

    public byte[] serializeToPbBytes() {
        Logs.LogGroup.Builder logs = Logs.LogGroup.newBuilder();
        if (!StringUtils.isEmpty(topic)) {
            logs.setTopic(topic);
        }
        if (!StringUtils.isEmpty(source)) {
            logs.setSource(source);
        }

        if (tags != null && !tags.isEmpty()) {
            for (LogTag tag : tags) {
                Logs.LogTag.Builder tagBuilder = logs.addLogTagsBuilder();
                tagBuilder.setKey(tag.getKey());
                tagBuilder.setValue(tag.getValue());
            }
        }
        for (LogItem item : logItems) {
            Logs.Log.Builder log = logs.addLogsBuilder();
            log.setTime(item.getTime());
            if (item.hasTimeNsPart()) {
                log.setTimeNs(item.getTimeNsPart());
            }

            for (LogContent content : item.getLogContents()) {
                Logs.LogContent.Builder contentBuilder = log.addContentsBuilder();
                contentBuilder.setKey(content.getKey());
                if (content.getValue() == null) {
                    contentBuilder.setValue("");
                } else {
                    contentBuilder.setValue(content.getValue());
                }
            }
        }
        return logs.build().toByteArray();
    }

    public String getCompressType() {
        return compressType;
    }

    public PostLogStoreLogsRequest setCompressType(String compressType) {
        this.compressType = compressType;
        return this;
    }

    /**
     * Get the topic
     *
     * @return the topic
     */
    public String getTopic() {
        return topic;
    }

    /**
     * Set topic value
     *
     * @param topic topic value
     */
    public PostLogStoreLogsRequest setTopic(String topic) {
        this.topic = topic;
        return this;
    }

    /**
     * Get log source
     *
     * @return log source
     */
    public String getSource() {
        return source;
    }

    /**
     * Set log source
     *
     * @param source log source
     */
    public PostLogStoreLogsRequest setSource(String source) {
        this.source = source;
        return this;
    }

    /**
     * Get all the log data
     *
     * @return log data
     */
    public List<LogItem> getLogItems() {
        return logItems;
    }

    /**
     * Get all the tag
     *
     * @return tag
     */
    public List<LogTag> getTags() {
        return tags;
    }

    /**
     * Set the log data , shallow copy is used to set the log data
     *
     * @param logItems log data
     */
    public PostLogStoreLogsRequest setLogItems(List<LogItem> logItems) {
        this.logItems = logItems;
        return this;
    }

    public PostLogStoreLogsRequest setTags(List<LogTag> tags) {
        this.tags = tags;
        return this;
    }
}
