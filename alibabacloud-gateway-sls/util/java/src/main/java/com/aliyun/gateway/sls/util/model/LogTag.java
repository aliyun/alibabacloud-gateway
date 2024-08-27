package com.aliyun.gateway.sls.util.model;

import java.io.Serializable;

/**
 * LogContent is a simple data structure used in @LogItem, it presents a
 * key/value pair in @logItem.
 *
 * @author sls_dev
 *
 */
public class LogTag implements Serializable {

    private static final long serialVersionUID = 2417167614162708776L;
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
    public LogTag(String key, String value) {
        this.key = key;
        this.value = value;
    }

    /**
     * Get log tag key
     *
     * @return log tag key
     */
    public String getKey() {
        return this.key;
    }

    /**
     * Get log content value
     *
     * @return log content value
     */
    public String getValue() {
        return this.value;
    }
}

