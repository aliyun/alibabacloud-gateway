// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class CompleteMultipartUploadHeaders extends TeaModel {
    @NameInMap("commonHeaders")
    public java.util.Map<String, String> commonHeaders;

    /**
     * <p>Specifies whether to list all parts that are uploaded by using the current upload ID. </p>
     * <br>
     * <p>- If x-oss-complete-all:yes is specified in the request, OSS lists all parts that are uploaded by using the current upload ID, sorts the parts by part number, and then calls the CompleteMultipartUpload operation. OSS cannot detect the parts that are not uploaded or being uploaded when OSS calls the CompleteMultipartUpload operation. Before you call the CompleteMultipartUpload operation, make sure that all parts are uploaded. </p>
     * <p>- If x-oss-complete-all:yes is specified in the request, the request body cannot be specified. Otherwise, an error occurs. </p>
     * <p>- If x-oss-complete-all:yes is specified in the request, the format of the response remains unchanged.</p>
     */
    @NameInMap("x-oss-complete-all")
    public String xOssCompleteAll;

    /**
     * <p>Specifies whether the object with the same object name is overwritten when you call the CompleteMultipartUpload operation. </p>
     * <br>
     * <p>- If the value of x-oss-forbid-overwrite is not specified or set to false, the existing object can be overwritten by the object that has the same name. </p>
     * <p>- If the value of x-oss-forbid-overwrite is set to true, the existing object cannot be overwritten by the object that has the same name. </p>
     * <br>
     * <p>> </p>
     * <p>- The x-oss-forbid-overwrite request header is invalid if versioning is enabled or suspended for the bucket. In this case, the existing object can be overwritten by the object that has the same name when you call the CompleteMultipartUpload operation. </p>
     * <p>- If you specify the x-oss-forbid-overwrite request header, the queries per second (QPS) performance of OSS may be degraded. If you want to configure the x-oss-forbid-overwrite header in a large number of requests (QPS > 1,000), submit a ticket.</p>
     */
    @NameInMap("x-oss-forbid-overwrite")
    public String xOssForbidOverwrite;

    public static CompleteMultipartUploadHeaders build(java.util.Map<String, ?> map) throws Exception {
        CompleteMultipartUploadHeaders self = new CompleteMultipartUploadHeaders();
        return TeaModel.build(map, self);
    }

    public CompleteMultipartUploadHeaders setCommonHeaders(java.util.Map<String, String> commonHeaders) {
        this.commonHeaders = commonHeaders;
        return this;
    }
    public java.util.Map<String, String> getCommonHeaders() {
        return this.commonHeaders;
    }

    public CompleteMultipartUploadHeaders setXOssCompleteAll(String xOssCompleteAll) {
        this.xOssCompleteAll = xOssCompleteAll;
        return this;
    }
    public String getXOssCompleteAll() {
        return this.xOssCompleteAll;
    }

    public CompleteMultipartUploadHeaders setXOssForbidOverwrite(String xOssForbidOverwrite) {
        this.xOssForbidOverwrite = xOssForbidOverwrite;
        return this;
    }
    public String getXOssForbidOverwrite() {
        return this.xOssForbidOverwrite;
    }

}
