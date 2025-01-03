// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class GetBucketRequestPaymentResponseBody extends TeaModel {
    /**
     * <p>Indicates the container for the payer.</p>
     */
    @NameInMap("RequestPaymentConfiguration")
    public RequestPaymentConfiguration requestPaymentConfiguration;

    public static GetBucketRequestPaymentResponseBody build(java.util.Map<String, ?> map) throws Exception {
        GetBucketRequestPaymentResponseBody self = new GetBucketRequestPaymentResponseBody();
        return TeaModel.build(map, self);
    }

    public GetBucketRequestPaymentResponseBody setRequestPaymentConfiguration(RequestPaymentConfiguration requestPaymentConfiguration) {
        this.requestPaymentConfiguration = requestPaymentConfiguration;
        return this;
    }
    public RequestPaymentConfiguration getRequestPaymentConfiguration() {
        return this.requestPaymentConfiguration;
    }

    public static class RequestPaymentConfiguration extends TeaModel {
        /**
         * <p>Indicates who pays the download and request fees.</p>
         * 
         * <strong>example:</strong>
         * <p>Requester</p>
         */
        @NameInMap("Payer")
        public String payer;

        public static RequestPaymentConfiguration build(java.util.Map<String, ?> map) throws Exception {
            RequestPaymentConfiguration self = new RequestPaymentConfiguration();
            return TeaModel.build(map, self);
        }

        public RequestPaymentConfiguration setPayer(String payer) {
            this.payer = payer;
            return this;
        }
        public String getPayer() {
            return this.payer;
        }

    }

}
