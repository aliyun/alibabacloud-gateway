// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.spi.models;

import com.aliyun.tea.*;

public class AttributeMap extends TeaModel {
    @NameInMap("attributes")
    @Validation(required = true)
    public java.util.Map<String, ?> attributes;

    @NameInMap("key")
    @Validation(required = true)
    public java.util.Map<String, String> key;

    public static AttributeMap build(java.util.Map<String, ?> map) {
        AttributeMap self = new AttributeMap();
        return TeaModel.build(map, self);
    }

    public AttributeMap setAttributes(java.util.Map<String, ?> attributes) {
        this.attributes = attributes;
        return this;
    }
    public java.util.Map<String, ?> getAttributes() {
        return this.attributes;
    }

    public AttributeMap setKey(java.util.Map<String, String> key) {
        this.key = key;
        return this;
    }
    public java.util.Map<String, String> getKey() {
        return this.key;
    }

}
