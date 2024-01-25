// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetMetaQueryStatusResponseBody extends TeaModel {
    /**
     * <p>The time when the metadata index library was created. The value follows the [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339) standard in the YYYY-MM-DDTHH:mm:ss+TIMEZONE format. YYYY-MM-DD indicates the year, month, and day. T indicates the beginning of the time element. HH:mm:ss indicates the hour, minute, and second. TIMEZONE indicates the time zone.</p>
     */
    @NameInMap("CreateTime")
    public String createTime;

    /**
     * <p>The scan type. Valid values:</p>
     * <br>
     * <p>*   FullScanning</p>
     * <p>*   IncrementalScanning</p>
     */
    @NameInMap("Phase")
    public String phase;

    /**
     * <p>The state of the metadata index library. Valid values:</p>
     * <br>
     * <p>*   Ready: The metadata index library is being prepared after it is created. In this case, the metadata index library cannot be used to query data.</p>
     * <p>*   Stop: The metadata index library is paused.</p>
     * <p>*   Running: The metadata index library is running.</p>
     * <p>*   Retrying: The metadata index library failed to be created and is being created again.</p>
     * <p>*   Failed: The metadata index library failed to be created.</p>
     * <p>*   Deleted: The metadata index library is deleted.</p>
     */
    @NameInMap("State")
    public String state;

    /**
     * <p>The time when the metadata index library was updated. The value follows the [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339) standard in the YYYY-MM-DDTHH:mm:ss+TIMEZONE format. YYYY-MM-DD indicates the year, month, and day. T indicates the beginning of the time element. HH:mm:ss indicates the hour, minute, and second. TIMEZONE indicates the time zone.</p>
     */
    @NameInMap("UpdateTime")
    public String updateTime;

    public static GetMetaQueryStatusResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetMetaQueryStatusResponseBody self = new GetMetaQueryStatusResponseBody();
        return TeaModel.build(map, self);
    }

    public GetMetaQueryStatusResponseBody setCreateTime(String createTime) {
        this.createTime = createTime;
        return this;
    }
    public String getCreateTime() {
        return this.createTime;
    }

    public GetMetaQueryStatusResponseBody setPhase(String phase) {
        this.phase = phase;
        return this;
    }
    public String getPhase() {
        return this.phase;
    }

    public GetMetaQueryStatusResponseBody setState(String state) {
        this.state = state;
        return this;
    }
    public String getState() {
        return this.state;
    }

    public GetMetaQueryStatusResponseBody setUpdateTime(String updateTime) {
        this.updateTime = updateTime;
        return this;
    }
    public String getUpdateTime() {
        return this.updateTime;
    }

}
