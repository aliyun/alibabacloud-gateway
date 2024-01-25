// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class AppendObjectRequest extends TeaModel {
    /**
     * <p>The request body.</p>
     */
    @NameInMap("body")
    public java.io.InputStream body;

    /**
     * <p>The position from which the AppendObject operation starts.  Each time an AppendObject operation succeeds, the x-oss-next-append-position header is included in the response to specify the position from which the next AppendObject operation starts. The value of position in the first AppendObject operation performed on an object must be 0. The value of position in subsequent AppendObject operations performed on the object is the current length of the object. For example, if the value of position specified in the first AppendObject request is 0 and the value of content-length is 65536, the value of position in the second AppendObject request must be 65536. </p>
     * <br>
     * <p>- If the value of position in the AppendObject request is 0 and the name of the object that you want to append is unique, you can set headers such as x-oss-server-side-encryption in an AppendObject request in the same way as you set in a PutObject request. If you add the x-oss-server-side-encryption header to an AppendObject request, the x-oss-server-side-encryption header is included in the response to the request. If you want to modify metadata, you can call the CopyObject operation. </p>
     * <p>- If you call an AppendObject operation to append a 0 KB object whose position value is valid to an Appendable object, the status of the Appendable object is not changed.</p>
     */
    @NameInMap("position")
    public Long position;

    public static AppendObjectRequest build(java.util.Map<String, ?> map) throws Exception {
        AppendObjectRequest self = new AppendObjectRequest();
        return TeaModel.build(map, self);
    }

    public AppendObjectRequest setBody(java.io.InputStream body) {
        this.body = body;
        return this;
    }
    public java.io.InputStream getBody() {
        return this.body;
    }

    public AppendObjectRequest setPosition(Long position) {
        this.position = position;
        return this;
    }
    public Long getPosition() {
        return this.position;
    }

}
