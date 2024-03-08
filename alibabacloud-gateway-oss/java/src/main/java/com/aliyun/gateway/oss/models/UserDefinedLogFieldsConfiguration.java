// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.models;

import com.aliyun.tea.*;

public class UserDefinedLogFieldsConfiguration extends TeaModel {
    @NameInMap("HeaderSet")
    public HeaderSet headerSet;

    @NameInMap("ParamSet")
    public ParamSet paramSet;

    public static UserDefinedLogFieldsConfiguration build(java.util.Map<String, ?> map) throws Exception {
        UserDefinedLogFieldsConfiguration self = new UserDefinedLogFieldsConfiguration();
        return TeaModel.build(map, self);
    }

    public UserDefinedLogFieldsConfiguration setHeaderSet(HeaderSet headerSet) {
        this.headerSet = headerSet;
        return this;
    }
    public HeaderSet getHeaderSet() {
        return this.headerSet;
    }

    public UserDefinedLogFieldsConfiguration setParamSet(ParamSet paramSet) {
        this.paramSet = paramSet;
        return this;
    }
    public ParamSet getParamSet() {
        return this.paramSet;
    }

    public static class HeaderSet extends TeaModel {
        @NameInMap("header")
        public java.util.List<String> header;

        public static HeaderSet build(java.util.Map<String, ?> map) throws Exception {
            HeaderSet self = new HeaderSet();
            return TeaModel.build(map, self);
        }

        public HeaderSet setHeader(java.util.List<String> header) {
            this.header = header;
            return this;
        }
        public java.util.List<String> getHeader() {
            return this.header;
        }

    }

    public static class ParamSet extends TeaModel {
        @NameInMap("parameter")
        public java.util.List<String> parameter;

        public static ParamSet build(java.util.Map<String, ?> map) throws Exception {
            ParamSet self = new ParamSet();
            return TeaModel.build(map, self);
        }

        public ParamSet setParameter(java.util.List<String> parameter) {
            this.parameter = parameter;
            return this;
        }
        public java.util.List<String> getParameter() {
            return this.parameter;
        }

    }

}
