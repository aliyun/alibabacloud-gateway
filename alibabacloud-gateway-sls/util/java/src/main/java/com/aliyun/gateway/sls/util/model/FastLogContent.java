package com.aliyun.gateway.sls.util.model;

import com.aliyun.gateway.sls.util.Client;

import java.io.UnsupportedEncodingException;

public class FastLogContent {

    private final byte[] rawBytes;
    // [beginOffset, endOffset)
    private final int beginOffset;
    private final int endOffset;
    private int keyOffset = -1;
    private int keyLength = -1;
    private int valueOffset = -1;
    private int valueLength = -1;

    public FastLogContent(byte[] rawBytes, int offset, int length) {
        this.rawBytes = rawBytes;
        this.beginOffset = offset;
        this.endOffset = offset + length;
        if (!parse()) {
            this.keyOffset = -1;
            this.keyLength = -1;
            this.valueOffset = -1;
            this.valueLength = -1;
        }
    }

    private boolean parse() {
        int pos = this.beginOffset;
        int index, mode;
        while (pos < this.endOffset) {
            int[] value = Client.decodeVarInt32(this.rawBytes, pos, this.endOffset);
            if (value[0] == 0) {
                return false;
            }
            mode = value[1] & 0x7;
            index = value[1] >> 3;
            pos = value[2];
            if (mode == 0) {
                value = Client.decodeVarInt32(this.rawBytes, pos, this.endOffset);
                if (value[0] == 0) {
                    return false;
                }
                pos = value[2];
            } else if (mode == 1) {
                pos += 8;
            } else if (mode == 2) {
                value = Client.decodeVarInt32(this.rawBytes, pos, this.endOffset);
                if (value[0] == 0) {
                    return false;
                }
                pos = value[2] + value[1];
                if (index == 1) {
                    keyOffset = value[2];
                    keyLength = value[1];
                } else if (index == 2) {
                    valueOffset = value[2];
                    valueLength = value[1];
                }
            } else if (mode == 5) {
                pos += 4;
            } else {
                return false;
            }
        }
        return (keyOffset != -1 && valueOffset != -1 && pos == this.endOffset);
    }

    public String getKey() {
        return decodeString(keyOffset, keyLength);
    }

    public String getKey(final String charset) throws UnsupportedEncodingException {
        return decodeString(keyOffset, keyLength, charset);
    }

    public String getValue() {
        return decodeString(valueOffset, valueLength);
    }

    public String getValue(final String charset) throws UnsupportedEncodingException {
        return decodeString(valueOffset, valueLength, charset);
    }

    private String decodeString(int offset, int length) {
        return offset < 0 ? null : new String(rawBytes, offset, length);
    }

    private String decodeString(int offset, int length, String charset) throws UnsupportedEncodingException {
        return offset < 0 ? null : new String(this.rawBytes, offset, length, charset);
    }
}
