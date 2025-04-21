// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CreateSelectObjectMetaRequest extends TeaModel {
    /**
     * <p>The request body used to create select meta.</p>
     */
    @NameInMap("CsvMetaRequest")
    public SelectMetaRequest csvMetaRequest;

    /**
     * <p>Parameters to specify the file formate.</p>
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>csv/meta</p>
     */
    @NameInMap("x-oss-process")
    public String xOssProcess;

    public static CreateSelectObjectMetaRequest build(java.util.Map<String, ?> map) throws Exception {
        CreateSelectObjectMetaRequest self = new CreateSelectObjectMetaRequest();
        return TeaModel.build(map, self);
    }

    public CreateSelectObjectMetaRequest setCsvMetaRequest(SelectMetaRequest csvMetaRequest) {
        this.csvMetaRequest = csvMetaRequest;
        return this;
    }
    public SelectMetaRequest getCsvMetaRequest() {
        return this.csvMetaRequest;
    }

    public CreateSelectObjectMetaRequest setXOssProcess(String xOssProcess) {
        this.xOssProcess = xOssProcess;
        return this;
    }
    public String getXOssProcess() {
        return this.xOssProcess;
    }

}
