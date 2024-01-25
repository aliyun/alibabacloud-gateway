// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class DeleteMultipleObjectsRequest extends TeaModel {
    /**
     * <p>The container that stores the DeleteMultipleObjects request.</p>
     */
    @NameInMap("body")
    public Delete body;

    /**
     * <p>The encoding type of the object name in the response. The value of the Key parameter is UTF-8 encoded. If the Key parameter includes control characters that are not supported by the XML 1.0 standard, you can specify this header to encode the value of the Key parameter in the response.</p>
     */
    @NameInMap("encoding-type")
    public String encodingType;

    public static DeleteMultipleObjectsRequest build(java.util.Map<String, ?> map) throws Exception {
        DeleteMultipleObjectsRequest self = new DeleteMultipleObjectsRequest();
        return TeaModel.build(map, self);
    }

    public DeleteMultipleObjectsRequest setBody(Delete body) {
        this.body = body;
        return this;
    }
    public Delete getBody() {
        return this.body;
    }

    public DeleteMultipleObjectsRequest setEncodingType(String encodingType) {
        this.encodingType = encodingType;
        return this;
    }
    public String getEncodingType() {
        return this.encodingType;
    }

}
