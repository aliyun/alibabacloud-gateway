// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class CallbackPolicy extends TeaModel {
    @NameInMap("PolicyItem")
    public java.util.List<CallbackPolicyPolicyItem> policyItem;

    public static CallbackPolicy build(java.util.Map<String, ?> map) throws Exception {
        CallbackPolicy self = new CallbackPolicy();
        return TeaModel.build(map, self);
    }

    public CallbackPolicy setPolicyItem(java.util.List<CallbackPolicyPolicyItem> policyItem) {
        this.policyItem = policyItem;
        return this;
    }
    public java.util.List<CallbackPolicyPolicyItem> getPolicyItem() {
        return this.policyItem;
    }

    public static class CallbackPolicyPolicyItem extends TeaModel {
        @NameInMap("Callback")
        public String callback;

        @NameInMap("CallbackVar")
        public String callbackVar;

        @NameInMap("PolicyName")
        public String policyName;

        public static CallbackPolicyPolicyItem build(java.util.Map<String, ?> map) throws Exception {
            CallbackPolicyPolicyItem self = new CallbackPolicyPolicyItem();
            return TeaModel.build(map, self);
        }

        public CallbackPolicyPolicyItem setCallback(String callback) {
            this.callback = callback;
            return this;
        }
        public String getCallback() {
            return this.callback;
        }

        public CallbackPolicyPolicyItem setCallbackVar(String callbackVar) {
            this.callbackVar = callbackVar;
            return this;
        }
        public String getCallbackVar() {
            return this.callbackVar;
        }

        public CallbackPolicyPolicyItem setPolicyName(String policyName) {
            this.policyName = policyName;
            return this;
        }
        public String getPolicyName() {
            return this.policyName;
        }

    }

}
