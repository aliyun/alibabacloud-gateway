package com.aliyun.gateway.sls.util.model;


import com.aliyun.tea.utils.StringUtils;

public class PullLogsRequest {
    private String cursor;
    private String endCursor;
    private int count;
    private String query;

    // GET /logstores/{logStore}/shards/{shardId}?type=log&cursor={cursor}&endCursor={endCursor}&count=${count}&query={query}
    public PullLogsRequest(String cursor, String endCursor, int count, String query) {
        this.cursor = cursor;
        this.endCursor = endCursor;
        this.count = count;
        this.query = query;
        if (StringUtils.isEmpty(cursor)) {
            throw new LogException("InvalidParameter", "The specified parameter cursor is missing.", "");
        }
    }

    public String getCursor() {
        return cursor;
    }

    public void setCursor(String cursor) {
        this.cursor = cursor;
    }

    public String getEndCursor() {
        return endCursor;
    }

    public void setEndCursor(String endCursor) {
        this.endCursor = endCursor;
    }

    public int getCount() {
        return count;
    }

    public void setCount(int count) {
        this.count = count;
    }

    public String getQuery() {
        return query;
    }

    public void setQuery(String query) {
        this.query = query;
    }

}
