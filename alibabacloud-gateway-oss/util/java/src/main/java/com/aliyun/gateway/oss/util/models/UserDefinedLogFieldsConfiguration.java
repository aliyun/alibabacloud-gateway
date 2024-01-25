// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class UserDefinedLogFieldsConfiguration extends TeaModel {
    @NameInMap("HeaderSet")
    public UserDefinedLogFieldsConfigurationHeaderSet headerSet;

    @NameInMap("ParamSet")
    public UserDefinedLogFieldsConfigurationParamSet paramSet;

    public static UserDefinedLogFieldsConfiguration build(java.util.Map<String, ?> map) throws Exception {
        UserDefinedLogFieldsConfiguration self = new UserDefinedLogFieldsConfiguration();
        return TeaModel.build(map, self);
    }

    public UserDefinedLogFieldsConfiguration setHeaderSet(UserDefinedLogFieldsConfigurationHeaderSet headerSet) {
        this.headerSet = headerSet;
        return this;
    }
    public UserDefinedLogFieldsConfigurationHeaderSet getHeaderSet() {
        return this.headerSet;
    }

    public UserDefinedLogFieldsConfiguration setParamSet(UserDefinedLogFieldsConfigurationParamSet paramSet) {
        this.paramSet = paramSet;
        return this;
    }
    public UserDefinedLogFieldsConfigurationParamSet getParamSet() {
        return this.paramSet;
    }

    public static class UserDefinedLogFieldsConfigurationHeaderSet extends TeaModel {
        @NameInMap("header")
        public java.util.List<String> header;

        public static UserDefinedLogFieldsConfigurationHeaderSet build(java.util.Map<String, ?> map) throws Exception {
            UserDefinedLogFieldsConfigurationHeaderSet self = new UserDefinedLogFieldsConfigurationHeaderSet();
            return TeaModel.build(map, self);
        }

        public UserDefinedLogFieldsConfigurationHeaderSet setHeader(java.util.List<String> header) {
            this.header = header;
            return this;
        }
        public java.util.List<String> getHeader() {
            return this.header;
        }

    }

    public static class UserDefinedLogFieldsConfigurationParamSet extends TeaModel {
        @NameInMap("parameter")
        public java.util.List<String> parameter;

        public static UserDefinedLogFieldsConfigurationParamSet build(java.util.Map<String, ?> map) throws Exception {
            UserDefinedLogFieldsConfigurationParamSet self = new UserDefinedLogFieldsConfigurationParamSet();
            return TeaModel.build(map, self);
        }

        public UserDefinedLogFieldsConfigurationParamSet setParameter(java.util.List<String> parameter) {
            this.parameter = parameter;
            return this;
        }
        public java.util.List<String> getParameter() {
            return this.parameter;
        }

    }

}
