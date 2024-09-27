package com.aliyun.gateway.sls.util.model;

import java.io.Serializable;
import java.util.ArrayList;
import java.util.Date;
import java.util.List;


/**
 * LogItem used to present a log, it contains log time, log source(ip/hostname,
 * e.g), and multiple key/value pairs to present the log content
 *
 * @author sls_dev
 */
public class LogItem implements Serializable {
    private static final long serialVersionUID = -3488075856612935955L;
    private int time;
    private int logTimeNsPart;
    private List<LogContent> contents;

    /**
     * Construct a logItem, the log time is set according to the sys time
     */
    public LogItem() {
        this((int) (new Date().getTime() / 1000));
    }

    /**
     * Construct a logItem with a certain time stamp
     *
     * @param time log time stamp
     */
    public LogItem(int time) {
        this(time, new ArrayList<LogContent>());
    }

    /**
     * Construct a logItem with a certain time stamp and log contents
     *
     * @param time  log time stamp
     * @param contents log contents
     */
    public LogItem(int time, List<LogContent> contents) {
        this(time, -1, contents);
    }

    /**
     * Construct a logItem with a certain time stamp and log contents
     *
     * @param time       log time stamp
     * @param logTimeNsPart log time nano stamp
     * @param contents      log contents
     */
    public LogItem(int time, int logTimeNsPart, List<LogContent> contents) {
        this.time = time;
        this.logTimeNsPart = logTimeNsPart;
        this.contents = contents;
    }

    /**
     * Get log time
     *
     * @return log time
     */
    public int getTime() {
        return time;
    }

    /**
     * Set logTime
     *
     * @param logTime log time
     */
    public void setTime(int logTime) {
        this.time = logTime;
    }

    /**
     * Get log timeNsPart
     *
     * @return log time ns part
     */
    public int getTimeNsPart() {
        return logTimeNsPart;
    }

    public boolean hasTimeNsPart() {
        return logTimeNsPart != -1;
    }

    /**
     * Set logTimeNsPart
     *
     * @param logTimeNsPart log time ns part
     */
    public void setTimeNsPart(int logTimeNsPart) {
        this.logTimeNsPart = logTimeNsPart;
    }

    /**
     * Add a log content key/value pair to the log
     *
     * @param key   log content key
     * @param value log content value
     */
    public void pushBack(String key, String value) {
        contents.add(new LogContent(key, value));
    }

    /**
     * Get log contents
     *
     * @return log contents
     */
    public List<LogContent> getLogContents() {
        return contents;
    }

    /**
     * set log contents
     *
     * @param contents log contents
     */
    public void setLogContents(List<LogContent> contents) {
        this.contents = contents;
    }

}
