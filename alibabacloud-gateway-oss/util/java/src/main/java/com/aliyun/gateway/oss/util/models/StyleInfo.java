// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class StyleInfo extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>image</p>
     */
    @NameInMap("Category")
    public String category;

    /**
     * <strong>example:</strong>
     * <p>image/resize,p_50</p>
     */
    @NameInMap("Content")
    public String content;

    /**
     * <strong>example:</strong>
     * <p>Wed, 20 May 2020 12:07:15 GMT</p>
     */
    @NameInMap("CreateTime")
    public String createTime;

    /**
     * <strong>example:</strong>
     * <p>Wed, 21 May 2020 12:07:15 GMT</p>
     */
    @NameInMap("LastModifyTime")
    public String lastModifyTime;

    /**
     * <strong>example:</strong>
     * <p>first</p>
     */
    @NameInMap("Name")
    public String name;

    public static StyleInfo build(java.util.Map<String, ?> map) throws Exception {
        StyleInfo self = new StyleInfo();
        return TeaModel.build(map, self);
    }

    public StyleInfo setCategory(String category) {
        this.category = category;
        return this;
    }
    public String getCategory() {
        return this.category;
    }

    public StyleInfo setContent(String content) {
        this.content = content;
        return this;
    }
    public String getContent() {
        return this.content;
    }

    public StyleInfo setCreateTime(String createTime) {
        this.createTime = createTime;
        return this;
    }
    public String getCreateTime() {
        return this.createTime;
    }

    public StyleInfo setLastModifyTime(String lastModifyTime) {
        this.lastModifyTime = lastModifyTime;
        return this;
    }
    public String getLastModifyTime() {
        return this.lastModifyTime;
    }

    public StyleInfo setName(String name) {
        this.name = name;
        return this;
    }
    public String getName() {
        return this.name;
    }

}
