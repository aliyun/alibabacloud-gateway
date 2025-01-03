// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListDataLakeCachePrefetchJobHistory extends TeaModel {
    @NameInMap("DataLakeCachePrefetchJobHistory")
    public java.util.List<DataLakeCachePrefetchJobHistory> dataLakeCachePrefetchJobHistory;

    public static ListDataLakeCachePrefetchJobHistory build(java.util.Map<String, ?> map) throws Exception {
        ListDataLakeCachePrefetchJobHistory self = new ListDataLakeCachePrefetchJobHistory();
        return TeaModel.build(map, self);
    }

    public ListDataLakeCachePrefetchJobHistory setDataLakeCachePrefetchJobHistory(java.util.List<DataLakeCachePrefetchJobHistory> dataLakeCachePrefetchJobHistory) {
        this.dataLakeCachePrefetchJobHistory = dataLakeCachePrefetchJobHistory;
        return this;
    }
    public java.util.List<DataLakeCachePrefetchJobHistory> getDataLakeCachePrefetchJobHistory() {
        return this.dataLakeCachePrefetchJobHistory;
    }

}
