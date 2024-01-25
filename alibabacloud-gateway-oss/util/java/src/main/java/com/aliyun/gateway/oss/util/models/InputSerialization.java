// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class InputSerialization extends TeaModel {
    @NameInMap("CSV")
    public CSVInput CSV;

    @NameInMap("CompressionType")
    public String compressionType;

    @NameInMap("JSON")
    public JSONInput JSON;

    public static InputSerialization build(java.util.Map<String, ?> map) throws Exception {
        InputSerialization self = new InputSerialization();
        return TeaModel.build(map, self);
    }

    public InputSerialization setCSV(CSVInput CSV) {
        this.CSV = CSV;
        return this;
    }
    public CSVInput getCSV() {
        return this.CSV;
    }

    public InputSerialization setCompressionType(String compressionType) {
        this.compressionType = compressionType;
        return this;
    }
    public String getCompressionType() {
        return this.compressionType;
    }

    public InputSerialization setJSON(JSONInput JSON) {
        this.JSON = JSON;
        return this;
    }
    public JSONInput getJSON() {
        return this.JSON;
    }

}
