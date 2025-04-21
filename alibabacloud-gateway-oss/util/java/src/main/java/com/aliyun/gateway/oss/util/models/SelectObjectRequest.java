// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class SelectObjectRequest extends TeaModel {
    /**
     * <p>Container for saving Select request.</p>
     */
    @NameInMap("SelectRequest")
    public SelectRequest selectRequest;

    /**
     * <p>If it is a CSV file, this value should be set to csv/select; if it is a JSON file, it should be set to json/select.</p>
     * <p>This parameter is required.</p>
     * 
     * <strong>example:</strong>
     * <p>csv/select</p>
     */
    @NameInMap("x-oss-process")
    public String xOssProcess;

    public static SelectObjectRequest build(java.util.Map<String, ?> map) throws Exception {
        SelectObjectRequest self = new SelectObjectRequest();
        return TeaModel.build(map, self);
    }

    public SelectObjectRequest setSelectRequest(SelectRequest selectRequest) {
        this.selectRequest = selectRequest;
        return this;
    }
    public SelectRequest getSelectRequest() {
        return this.selectRequest;
    }

    public SelectObjectRequest setXOssProcess(String xOssProcess) {
        this.xOssProcess = xOssProcess;
        return this;
    }
    public String getXOssProcess() {
        return this.xOssProcess;
    }

}
