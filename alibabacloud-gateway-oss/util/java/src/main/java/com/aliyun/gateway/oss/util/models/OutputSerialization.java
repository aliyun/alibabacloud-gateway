// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class OutputSerialization extends TeaModel {
    @NameInMap("CSV")
    public CSVOutput CSV;

    @NameInMap("EnablePayloadCrc")
    public Boolean enablePayloadCrc;

    @NameInMap("JSON")
    public JSONOutput JSON;

    @NameInMap("KeepAllColumns")
    public Boolean keepAllColumns;

    @NameInMap("OutputHeader")
    public Boolean outputHeader;

    @NameInMap("OutputRawData")
    public Boolean outputRawData;

    public static OutputSerialization build(java.util.Map<String, ?> map) throws Exception {
        OutputSerialization self = new OutputSerialization();
        return TeaModel.build(map, self);
    }

    public OutputSerialization setCSV(CSVOutput CSV) {
        this.CSV = CSV;
        return this;
    }
    public CSVOutput getCSV() {
        return this.CSV;
    }

    public OutputSerialization setEnablePayloadCrc(Boolean enablePayloadCrc) {
        this.enablePayloadCrc = enablePayloadCrc;
        return this;
    }
    public Boolean getEnablePayloadCrc() {
        return this.enablePayloadCrc;
    }

    public OutputSerialization setJSON(JSONOutput JSON) {
        this.JSON = JSON;
        return this;
    }
    public JSONOutput getJSON() {
        return this.JSON;
    }

    public OutputSerialization setKeepAllColumns(Boolean keepAllColumns) {
        this.keepAllColumns = keepAllColumns;
        return this;
    }
    public Boolean getKeepAllColumns() {
        return this.keepAllColumns;
    }

    public OutputSerialization setOutputHeader(Boolean outputHeader) {
        this.outputHeader = outputHeader;
        return this;
    }
    public Boolean getOutputHeader() {
        return this.outputHeader;
    }

    public OutputSerialization setOutputRawData(Boolean outputRawData) {
        this.outputRawData = outputRawData;
        return this;
    }
    public Boolean getOutputRawData() {
        return this.outputRawData;
    }

}
