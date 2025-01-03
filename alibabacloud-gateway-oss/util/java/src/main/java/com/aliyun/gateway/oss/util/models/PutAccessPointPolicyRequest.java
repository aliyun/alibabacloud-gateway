// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class PutAccessPointPolicyRequest extends TeaModel {
    /**
     * <p>The configurations of the access point policy.</p>
     * 
     * <strong>example:</strong>
     * <p>{
     *    &quot;Version&quot;:&quot;1&quot;,
     *    &quot;Statement&quot;:[
     *    {
     *      &quot;Action&quot;:[
     *        &quot;oss:PutObject&quot;,
     *        &quot;oss:GetObject&quot;
     *     ],
     *     &quot;Effect&quot;:&quot;Deny&quot;,
     *     &quot;Principal&quot;:[&quot;27737962156157xxxx&quot;],
     *     &quot;Resource&quot;:[
     *        &quot;acs:oss:ap-southeast-2:111933544165xxxx:accesspoint/$ap-01&quot;,
     *        &quot;acs:oss:ap-southeast-2:111933544165xxxx:accesspoint/$ap-01/object/*&quot;
     *      ]
     *    }
     *   ]
     *  }</p>
     */
    @NameInMap("body")
    public String body;

    public static PutAccessPointPolicyRequest build(java.util.Map<String, ?> map) throws Exception {
        PutAccessPointPolicyRequest self = new PutAccessPointPolicyRequest();
        return TeaModel.build(map, self);
    }

    public PutAccessPointPolicyRequest setBody(String body) {
        this.body = body;
        return this;
    }
    public String getBody() {
        return this.body;
    }

}
