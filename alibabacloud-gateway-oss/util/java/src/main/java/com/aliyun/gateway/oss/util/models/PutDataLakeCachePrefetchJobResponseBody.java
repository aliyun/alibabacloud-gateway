// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutDataLakeCachePrefetchJobResponseBody extends TeaModel {
    @NameInMap("DataLakeCachePrefetchJobID")
    public DataLakeCachePrefetchJobID dataLakeCachePrefetchJobID;

    public static PutDataLakeCachePrefetchJobResponseBody build(java.util.Map<String, ?> map) throws Exception {
        PutDataLakeCachePrefetchJobResponseBody self = new PutDataLakeCachePrefetchJobResponseBody();
        return TeaModel.build(map, self);
    }

    public PutDataLakeCachePrefetchJobResponseBody setDataLakeCachePrefetchJobID(DataLakeCachePrefetchJobID dataLakeCachePrefetchJobID) {
        this.dataLakeCachePrefetchJobID = dataLakeCachePrefetchJobID;
        return this;
    }
    public DataLakeCachePrefetchJobID getDataLakeCachePrefetchJobID() {
        return this.dataLakeCachePrefetchJobID;
    }

    public static class DataLakeCachePrefetchJobID extends TeaModel {
        @NameInMap("ID")
        public String ID;

        public static DataLakeCachePrefetchJobID build(java.util.Map<String, ?> map) throws Exception {
            DataLakeCachePrefetchJobID self = new DataLakeCachePrefetchJobID();
            return TeaModel.build(map, self);
        }

        public DataLakeCachePrefetchJobID setID(String ID) {
            this.ID = ID;
            return this;
        }
        public String getID() {
            return this.ID;
        }

    }

}
