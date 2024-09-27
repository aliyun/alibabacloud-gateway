package com.aliyun.gateway.sls.util.model;

import java.io.Serializable;

/**
 * LogContent is a simple data structure used in @LogItem, it presents a
 * key/value pair in @logItem.
 *
 * @author sls_dev
 *
 */
public class LogContent implements Serializable {
    private static final long serialVersionUID = 6042186396863898096L;
    private final String key;
    private final String value;

    /**
     * Construct a log content pair
     *
     * @param key
     *            log content key
     * @param value
     *            log content value
     */
    public LogContent(String key, String value) {
        this.key = key;
        this.value = value;
    }


    /**
     * Get log content key
     *
     * @return log content key
     */
    public String getKey() {
        return key;
    }

    /**
     * Get log content value
     *
     * @return log content value
     */
    public String getValue() {
        return value;
    }
}

