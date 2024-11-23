package com.aliyun.gateway.sls.util.model;

import com.aliyun.gateway.sls.util.Client;
import com.aliyun.gateway.sls.util.exception.LogException;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;

public class LogsResponseBody {
    // init fields
    private final byte[] rawData;
    private final Map<String, String> headers;
    private final int statusCode;
    // parsed fields
    private List<FastLogGroup> logGroups;
    private int count;
    private String requestId;
    private long cursorTime;
    private boolean isEndOfCursor;

    public LogsResponseBody(byte[] rawData, int statusCode, Map<String, String> headers) {
        this.rawData = rawData;
        this.statusCode = statusCode;
        this.headers = headers;
        parseHeaders();
    }

    public List<FastLogGroup> getLogGroups() throws LogException {
        if (logGroups == null) {
            parseFastLogGroupList(rawData);
        }
        return logGroups;
    }

    /**
     * Get number of logGroups
     *
     * @return number of logGroups fetched
     */
    public int getCount() {
        return count;
    }

    public long getCursorTime() {
        return cursorTime;
    }

    public boolean isEndOfCursor() {
        return isEndOfCursor;
    }

    private void parseHeaders() {
        requestId = headers.get("x-request-id");
        isEndOfCursor = "1".equals(headers.get("x-log-end-of-cursor"));
        try {
            count = Integer.parseInt(headers.get("x-log-count"));
            String cursorTimeHeader = headers.get("x-log-cursor-time");
            cursorTime = cursorTimeHeader != null && !cursorTimeHeader.isEmpty() ? Long.parseLong(cursorTimeHeader.trim()) : 0;
        } catch (NumberFormatException e) {
            throw new LogException("ParsePullLogHeaderError", e.getMessage(), requestId);
        }
    }

    /**
     * Parse LogGroupList using fast deserialize method.
     *
     * @param data is LogGroupList bytes
     * @throws LogException if parse fails
     */
    private void parseFastLogGroupList(byte[] data) throws LogException {
        logGroups = new ArrayList<>();
        if (data == null || data.length == 0) {
            return;
        }
        int pos = 0;
        int rawSize = data.length;
        int mode, index;
        while (pos < rawSize) {
            int[] value = Client.decodeVarInt32(data, pos, rawSize);
            if (value[0] == 0) {
                throw new LogException("InitLogGroupsError", "decode varint32 error", requestId, statusCode);
            }
            pos = value[2];
            mode = value[1] & 0x7;
            index = value[1] >> 3;
            if (mode == 0) {
                value = Client.decodeVarInt32(data, pos, rawSize);
                if (value[0] == 0) {
                    throw new LogException("InitLogGroupsError", "decode varint32 error", requestId, statusCode);
                }
                pos = value[2];
            } else if (mode == 1) {
                pos += 8;
            } else if (mode == 2) {
                value = Client.decodeVarInt32(data, pos, rawSize);
                if (value[0] == 0) {
                    throw new LogException("InitLogGroupsError", "decode varint32 error", requestId, statusCode);
                }
                if (index == 1) {
                    logGroups.add(new FastLogGroup(data, value[2], value[1], requestId));
                }
                pos = value[1] + value[2];
            } else if (mode == 5) {
                pos += 4;
            } else {
                throw new LogException("InitLogGroupsError", "mode: " + mode, requestId, statusCode);
            }
        }
        if (pos != rawSize) {
            throw new LogException("InitLogGroupsError", "parse LogGroupList fail", requestId, statusCode);
        }
    }
}
