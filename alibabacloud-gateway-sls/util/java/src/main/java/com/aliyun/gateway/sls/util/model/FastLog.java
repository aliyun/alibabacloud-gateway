package com.aliyun.gateway.sls.util.model;

import com.aliyun.gateway.sls.util.Client;

import java.util.ArrayList;
import java.util.List;

public class FastLog {

    private final byte[] rawBytes;
    // [beginOffset, endOffset)
    private final int beginOffset;
    private final int endOffset;
    private final List<FastLogContent> contents;
    private int time = -1;
    private int timeNsPart = 0;

    public FastLog(byte[] rawBytes, int offset, int length) {
        this.rawBytes = rawBytes;
        this.beginOffset = offset;
        this.endOffset = offset + length;
        this.contents = new ArrayList<FastLogContent>();
        if (!parse()) {
            this.contents.clear();
        }
    }

    private boolean parse() {
        int pos = this.beginOffset;
        int mode, index;
        boolean findTime = false;
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
                if (index == 1) {
                    this.time = value[1];
                    findTime = true;
                }
            } else if (mode == 1) {
                pos = value[2] + 8;
            } else if (mode == 2) {
                pos = value[2];
                value = Client.decodeVarInt32(this.rawBytes, pos, this.endOffset);
                if (value[0] == 0) {
                    return false;
                }
                pos = value[2] + value[1];
                if (index == 2) {
                    this.contents.add(new FastLogContent(this.rawBytes, value[2], value[1]));
                }
            } else if (mode == 5) {
                if (index == 4) {
                    timeNsPart = this.rawBytes[value[2]] & 255 | (this.rawBytes[value[2] + 1] & 255) << 8 | (this.rawBytes[value[2] + 2] & 255) << 16 | (this.rawBytes[value[2] + 3] & 255) << 24;
                }
                pos = value[2] + 4;
            } else {
                return false;
            }
        }
        return findTime && (pos == this.endOffset);
    }

    public int getTime() {
        return this.time;
    }

    public int getTimeNsPart() {
        return this.timeNsPart;
    }

    public int getContentsCount() {
        return this.contents.size();
    }

    public List<FastLogContent> getContents() {
        return contents;
    }

    public FastLogContent getContents(int i) {
        if (i < this.contents.size()) {
            return this.contents.get(i);
        } else {
            return null;
        }
    }

}
