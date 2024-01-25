// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class GetStyleResponseBody extends TeaModel {
    /**
     * <p>The content of the image style.</p>
     */
    @NameInMap("Content")
    public String content;

    /**
     * <p>The time when the image style was created.</p>
     */
    @NameInMap("CreateTime")
    public String createTime;

    /**
     * <p>The time when the image style was last modified.</p>
     */
    @NameInMap("LastModifyTime")
    public String lastModifyTime;

    /**
     * <p>The name of the image style.</p>
     */
    @NameInMap("Name")
    public String name;

    public static GetStyleResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetStyleResponseBody self = new GetStyleResponseBody();
        return TeaModel.build(map, self);
    }

    public GetStyleResponseBody setContent(String content) {
        this.content = content;
        return this;
    }
    public String getContent() {
        return this.content;
    }

    public GetStyleResponseBody setCreateTime(String createTime) {
        this.createTime = createTime;
        return this;
    }
    public String getCreateTime() {
        return this.createTime;
    }

    public GetStyleResponseBody setLastModifyTime(String lastModifyTime) {
        this.lastModifyTime = lastModifyTime;
        return this;
    }
    public String getLastModifyTime() {
        return this.lastModifyTime;
    }

    public GetStyleResponseBody setName(String name) {
        this.name = name;
        return this;
    }
    public String getName() {
        return this.name;
    }

}
