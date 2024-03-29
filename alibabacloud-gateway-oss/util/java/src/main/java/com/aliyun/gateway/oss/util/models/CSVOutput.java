// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CSVOutput extends TeaModel {
    @NameInMap("FieldDelimiter")
    public String fieldDelimiter;

    @NameInMap("RecordDelimiter")
    public String recordDelimiter;

    public static CSVOutput build(java.util.Map<String, ?> map) throws Exception {
        CSVOutput self = new CSVOutput();
        return TeaModel.build(map, self);
    }

    public CSVOutput setFieldDelimiter(String fieldDelimiter) {
        this.fieldDelimiter = fieldDelimiter;
        return this;
    }
    public String getFieldDelimiter() {
        return this.fieldDelimiter;
    }

    public CSVOutput setRecordDelimiter(String recordDelimiter) {
        this.recordDelimiter = recordDelimiter;
        return this;
    }
    public String getRecordDelimiter() {
        return this.recordDelimiter;
    }

}
