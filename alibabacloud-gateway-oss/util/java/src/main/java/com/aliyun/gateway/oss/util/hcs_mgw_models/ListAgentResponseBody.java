// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.hcs_mgw_models;

import com.aliyun.tea.*;

public class ListAgentResponseBody extends TeaModel {
    /**
     * <p>The details of the agents.</p>
     */
    @NameInMap("ImportAgentList")
    public ListAgentResp importAgentList;

    public static ListAgentResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListAgentResponseBody self = new ListAgentResponseBody();
        return TeaModel.build(map, self);
    }

    public ListAgentResponseBody setImportAgentList(ListAgentResp importAgentList) {
        this.importAgentList = importAgentList;
        return this;
    }
    public ListAgentResp getImportAgentList() {
        return this.importAgentList;
    }

}
