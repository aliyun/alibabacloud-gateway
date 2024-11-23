package com.aliyun.gateway.sls.util.model;

import java.util.List;

public interface ILogItem {
    Integer getTime();

    Integer getTimeNsPart();

    List<ILogContent> getLogContents();
}
