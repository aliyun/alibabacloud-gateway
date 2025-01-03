// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ListDataLakeCachePrefetchJobHistoryResponseBody extends TeaModel {
    @NameInMap("ListDataLakeCachePrefetchJobHistory")
    public ListDataLakeCachePrefetchJobHistory listDataLakeCachePrefetchJobHistory;

    public static ListDataLakeCachePrefetchJobHistoryResponseBody build(java.util.Map<String, ?> map) throws Exception {
        ListDataLakeCachePrefetchJobHistoryResponseBody self = new ListDataLakeCachePrefetchJobHistoryResponseBody();
        return TeaModel.build(map, self);
    }

    public ListDataLakeCachePrefetchJobHistoryResponseBody setListDataLakeCachePrefetchJobHistory(ListDataLakeCachePrefetchJobHistory listDataLakeCachePrefetchJobHistory) {
        this.listDataLakeCachePrefetchJobHistory = listDataLakeCachePrefetchJobHistory;
        return this;
    }
    public ListDataLakeCachePrefetchJobHistory getListDataLakeCachePrefetchJobHistory() {
        return this.listDataLakeCachePrefetchJobHistory;
    }

}
