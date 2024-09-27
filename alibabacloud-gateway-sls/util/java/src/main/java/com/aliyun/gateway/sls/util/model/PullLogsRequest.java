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

    public PullLogsRequest setCursor(String cursor) {
        this.cursor = cursor;
        return this;
    }

    public String getEndCursor() {
        return endCursor;
    }

    public PullLogsRequest setEndCursor(String endCursor) {
        this.endCursor = endCursor;
        return this;
    }

    public int getCount() {
        return count;
    }

    public PullLogsRequest setCount(int count) {
        this.count = count;
        return this;
    }

    public String getQuery() {
        return query;
    }

    public PullLogsRequest setQuery(String query) {
        this.query = query;
        return this;
    }

}
