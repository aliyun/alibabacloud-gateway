package com.aliyun.gateway.sls.util.model;

import com.aliyun.gateway.sls.util.Client;
import com.aliyun.gateway.sls.util.exception.LogException;

import java.util.ArrayList;
import java.util.List;

public class FastLogGroup {
    private final byte[] rawBytes;
    // [beginOffset, endOffset)
    private final int beginOffset;
    private final int endOffset;
    private final List<FastLog> logs;
    private final List<FastLogTag> tags;
    private final String requestId;
    private int topicOffset;
    private int sourceOffset;

    public FastLogGroup(byte[] rawBytes, int offset, int length, String requestId) {
        this.rawBytes = rawBytes;
        this.beginOffset = offset;
        this.endOffset = offset + length;
        this.topicOffset = -1;
        this.sourceOffset = -1;
        this.logs = new ArrayList<FastLog>();
        this.tags = new ArrayList<FastLogTag>();
        this.requestId = requestId;
        if (!parse()) {
            throw new LogException("InitFastLogGroupError", "invalid logGroup data", this.requestId);
        }
    }

    private boolean parse() {
        int pos = this.beginOffset;
        int mode, index;
        while (pos < this.endOffset) {
            int[] value = Client.decodeVarInt32(this.rawBytes, pos, this.endOffset);
            if (value[0] == 0) {
                return false;
            }
            mode = value[1] & 0x7;
            index = value[1] >> 3;
            if (mode == 0) {
                pos = value[2];
                value = Client.decodeVarInt32(this.rawBytes, pos, this.endOffset);
                if (value[0] == 0) {
                    return false;
                }
                pos = value[2];
            } else if (mode == 1) {
                pos = value[2] + 8;
            } else if (mode == 2) {
                switch (index) {
                    case 1:
                        //logs
                        break;
                    case 3:
                        this.topicOffset = value[2];
                        break;
                    case 4:
                        this.sourceOffset = value[2];
                        break;
                    case 6:
                        //tags
                        break;
                    default:
                }
                pos = value[2];
                value = Client.decodeVarInt32(this.rawBytes, pos, this.endOffset);
                if (value[0] == 0) {
                    return false;
                }
                pos = value[2] + value[1];
                if (index == 1) {
                    this.logs.add(new FastLog(this.rawBytes, value[2], value[1]));
                } else if (index == 6) {
                    this.tags.add(new FastLogTag(this.rawBytes, value[2], value[1]));
                }
            } else if (mode == 5) {
                pos = value[2] + 4;
            } else {
                return false;
            }
        }
        return (pos == this.endOffset);
    }

    public int getLogTagsCount() {
        return this.tags.size();
    }

    public FastLogTag getLogTags(int i) {
        if (i < this.tags.size()) {
            return this.tags.get(i);
        } else {
            return null;
        }
    }

    public List<FastLog> getLogs() {
        return logs;
    }

    public List<FastLogTag> getTags() {
        return tags;
    }

    public int getLogsCount() {
        return this.logs.size();
    }

    public FastLog getLogs(int i) {
        if (i < this.logs.size()) {
            return this.logs.get(i);
        } else {
            return null;
        }
    }

    public boolean hasTopic() {
        return this.topicOffset >= 0;
    }

    public boolean hasSource() {
        return this.sourceOffset >= 0;
    }
}

