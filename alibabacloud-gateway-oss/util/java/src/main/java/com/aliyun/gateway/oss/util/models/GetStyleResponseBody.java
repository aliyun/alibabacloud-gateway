// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.utils.models;

import com.aliyun.tea.*;

public class GetStyleResponseBody extends TeaModel {
    /**
     * <p>The container that stores the information about the image style.</p>
     */
    @NameInMap("Style")
    public Style style;

    public static GetStyleResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetStyleResponseBody self = new GetStyleResponseBody();
        return TeaModel.build(map, self);
    }

    public GetStyleResponseBody setStyle(Style style) {
        this.style = style;
        return this;
    }
    public Style getStyle() {
        return this.style;
    }

    public static class Style extends TeaModel {
        /**
         * <p>The content of the image style.</p>
         */
        @NameInMap("Content")
        public String content;

        /**
         * <p>The time when the style was created.</p>
         */
        @NameInMap("CreateTime")
        public String createTime;

        /**
         * <p>The time when the image style was modified.</p>
         */
        @NameInMap("LastModifyTime")
        public String lastModifyTime;

        /**
         * <p>The name of the image style.</p>
         */
        @NameInMap("Name")
        public String name;

        public static Style build(java.util.Map<String, ?> map) throws Exception {
            Style self = new Style();
            return TeaModel.build(map, self);
        }

        public Style setContent(String content) {
            this.content = content;
            return this;
        }
        public String getContent() {
            return this.content;
        }

        public Style setCreateTime(String createTime) {
            this.createTime = createTime;
            return this;
        }
        public String getCreateTime() {
            return this.createTime;
        }

        public Style setLastModifyTime(String lastModifyTime) {
            this.lastModifyTime = lastModifyTime;
            return this;
        }
        public String getLastModifyTime() {
            return this.lastModifyTime;
        }

        public Style setName(String name) {
            this.name = name;
            return this;
        }
        public String getName() {
            return this.name;
        }

    }

}
