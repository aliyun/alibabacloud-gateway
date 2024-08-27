package com.aliyun.gateway.sls.util.model;

import com.aliyun.tea.utils.StringUtils;

import java.io.Serializable;
import java.util.List;
import java.util.UUID;

public class PostLogStoreLogsRequest implements Serializable {
    private static final long serialVersionUID = 7226856831224917838L;
    private String topic;
    private String source;
    private String hashKey;
    private List<LogItem> logItems;
    private List<LogTag> tags = null;
    private String compressType = "lz4";
    private Integer hashRouteKeySeqId;

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

    public void setCompressType(String compressType) {
        this.compressType = compressType;
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
    public void setTopic(String topic) {
        this.topic = topic;
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
    public void setSource(String source) {
        this.source = source;
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
    public void setLogItems(List<LogItem> logItems) {
        this.logItems = logItems;
    }

    public void setTags(List<LogTag> tags) {
        this.tags = tags;
    }

    public String getHashKey() {
        return hashKey;
    }

    public void setHashKey(String mHashKey) {
        this.hashKey = mHashKey;
    }

    public Integer getHashRouteKeySeqId() {
        return hashRouteKeySeqId;
    }

    public void setHashRouteKeySeqId(Integer hashRouteKeySeqId) {
        this.hashRouteKeySeqId = hashRouteKeySeqId;
    }

}
