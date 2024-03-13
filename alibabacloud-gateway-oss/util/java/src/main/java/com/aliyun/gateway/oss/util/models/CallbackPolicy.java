// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class CallbackPolicy extends TeaModel {
    @NameInMap("PolicyItem")
    public java.util.List<PolicyItem> policyItem;

    public static CallbackPolicy build(java.util.Map<String, ?> map) throws Exception {
        CallbackPolicy self = new CallbackPolicy();
        return TeaModel.build(map, self);
    }

    public CallbackPolicy setPolicyItem(java.util.List<PolicyItem> policyItem) {
        this.policyItem = policyItem;
        return this;
    }
    public java.util.List<PolicyItem> getPolicyItem() {
        return this.policyItem;
    }

    public static class PolicyItem extends TeaModel {
        @NameInMap("Callback")
        public String callback;

        @NameInMap("CallbackVar")
        public String callbackVar;

        @NameInMap("PolicyName")
        public String policyName;

        public static PolicyItem build(java.util.Map<String, ?> map) throws Exception {
            PolicyItem self = new PolicyItem();
            return TeaModel.build(map, self);
        }

        public PolicyItem setCallback(String callback) {
            this.callback = callback;
            return this;
        }
        public String getCallback() {
            return this.callback;
        }

        public PolicyItem setCallbackVar(String callbackVar) {
            this.callbackVar = callbackVar;
            return this;
        }
        public String getCallbackVar() {
            return this.callbackVar;
        }

        public PolicyItem setPolicyName(String policyName) {
            this.policyName = policyName;
            return this;
        }
        public String getPolicyName() {
            return this.policyName;
        }

    }

}
