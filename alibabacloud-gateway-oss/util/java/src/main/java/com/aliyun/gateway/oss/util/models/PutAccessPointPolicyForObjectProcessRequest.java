// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutAccessPointPolicyForObjectProcessRequest extends TeaModel {
    /**
     * <p>The json format permission policies for an Object FC Access Point.</p>
     * 
     * <strong>example:</strong>
     * <p>{
     *                 &quot;Version&quot;: &quot;1&quot;,
     *                 &quot;Statement&quot;: [{
     *                     &quot;Effect&quot;: &quot;Allow&quot;,
     *                     &quot;Action&quot;: [
     *                         &quot;oss:GetObject&quot;
     *                     ],
     *                     &quot;Principal&quot;: [
     *                         &quot;237210000000000xxxx&quot;
     *                     ],
     *                     &quot;Resource&quot;: [
     *                         &quot;acs:oss:cn-qingdao:1030700000xxxx:accesspointforobjectprocess/fc-test/object/*&quot;
     *                     ]
     *                 }]
     *             }</p>
     */
    @NameInMap("body")
    public String body;

    public static PutAccessPointPolicyForObjectProcessRequest build(java.util.Map<String, ?> map) throws Exception {
        PutAccessPointPolicyForObjectProcessRequest self = new PutAccessPointPolicyForObjectProcessRequest();
        return TeaModel.build(map, self);
    }

    public PutAccessPointPolicyForObjectProcessRequest setBody(String body) {
        this.body = body;
        return this;
    }
    public String getBody() {
        return this.body;
    }

}
