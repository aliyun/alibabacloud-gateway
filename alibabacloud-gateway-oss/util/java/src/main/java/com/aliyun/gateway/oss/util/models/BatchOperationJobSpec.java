// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class BatchOperationJobSpec extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>Bucket,Key,VersionId</p>
     */
    @NameInMap("Fields")
    public String fields;

    /**
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>csv</p>
     */
    @NameInMap("Format")
    public String format;

    public static BatchOperationJobSpec build(java.util.Map<String, ?> map) throws Exception {
        BatchOperationJobSpec self = new BatchOperationJobSpec();
        return TeaModel.build(map, self);
    }

    public BatchOperationJobSpec setFields(String fields) {
        this.fields = fields;
        return this;
    }
    public String getFields() {
        return this.fields;
    }

    public BatchOperationJobSpec setFormat(String format) {
        this.format = format;
        return this;
    }
    public String getFormat() {
        return this.format;
    }

}
