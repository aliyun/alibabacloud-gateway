package com.aliyun.gateway.sls.util.model;

import com.aliyun.tea.TeaModel;
import com.aliyun.tea.utils.StringUtils;

import java.util.List;

public abstract class LogGroupBase extends TeaModel implements ILog {

    @Override
    public byte[] serializeToPbBytes() {
        Logs.LogGroup.Builder logs = Logs.LogGroup.newBuilder();
        if (!StringUtils.isEmpty(getTopic())) {
            logs.setTopic(getTopic());
        }
        if (!StringUtils.isEmpty(getSource())) {
            logs.setSource(getSource());
        }

        List<ILogTag> logTags = getLogTags();
        if (logTags != null && !logTags.isEmpty()) {
            for (ILogTag tag : logTags) {
                Logs.LogTag.Builder tagBuilder = logs.addLogTagsBuilder();
                tagBuilder.setKey(tag.getKey());
                tagBuilder.setValue(tag.getValue());
            }
        }
        List<ILogItem> logItems = getLogItems();
        for (ILogItem item : logItems) {
            Logs.Log.Builder log = logs.addLogsBuilder();
            log.setTime(item.getTime());
            if (item.getTimeNsPart() != -1) {
                log.setTimeNs(item.getTimeNsPart());
            }

            for (ILogContent content : item.getLogContents()) {
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

    public abstract String getSource();

    public abstract String getTopic();

    public abstract List<ILogTag> getLogTags();

    public abstract List<ILogItem> getLogItems();

}
