// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetLiveChannelInfoResponseBody extends TeaModel {
    /**
     * <p>The description of the LiveChannel.</p>
     */
    @NameInMap("Description")
    public String description;

    /**
     * <p>The status of the LiveChannel.</p>
     * <p><br>Valid values:</p>
     * <br>
     * <p>- enabled: indicates that the LiveChannel is enabled.</p>
     * <p>- disabled: indicates that the LiveChannel is disabled.</p>
     */
    @NameInMap("Status")
    public String status;

    /**
     * <p>The container that stores the configurations used by the LiveChannel to store uploaded data.</p>
     */
    @NameInMap("Target")
    public LiveChannelTarget target;

    public static GetLiveChannelInfoResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetLiveChannelInfoResponseBody self = new GetLiveChannelInfoResponseBody();
        return TeaModel.build(map, self);
    }

    public GetLiveChannelInfoResponseBody setDescription(String description) {
        this.description = description;
        return this;
    }
    public String getDescription() {
        return this.description;
    }

    public GetLiveChannelInfoResponseBody setStatus(String status) {
        this.status = status;
        return this;
    }
    public String getStatus() {
        return this.status;
    }

    public GetLiveChannelInfoResponseBody setTarget(LiveChannelTarget target) {
        this.target = target;
        return this;
    }
    public LiveChannelTarget getTarget() {
        return this.target;
    }

}
