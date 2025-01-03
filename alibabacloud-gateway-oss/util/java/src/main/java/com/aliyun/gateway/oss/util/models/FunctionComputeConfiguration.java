// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class FunctionComputeConfiguration extends TeaModel {
    @NameInMap("Event")
    public java.util.List<String> event;

    @NameInMap("Filter")
    public Filter filter;

    @NameInMap("Function")
    public Function function;

    /**
     * <strong>example:</strong>
     * <p>1</p>
     */
    @NameInMap("ID")
    public String ID;

    public static FunctionComputeConfiguration build(java.util.Map<String, ?> map) throws Exception {
        FunctionComputeConfiguration self = new FunctionComputeConfiguration();
        return TeaModel.build(map, self);
    }

    public FunctionComputeConfiguration setEvent(java.util.List<String> event) {
        this.event = event;
        return this;
    }
    public java.util.List<String> getEvent() {
        return this.event;
    }

    public FunctionComputeConfiguration setFilter(Filter filter) {
        this.filter = filter;
        return this;
    }
    public Filter getFilter() {
        return this.filter;
    }

    public FunctionComputeConfiguration setFunction(Function function) {
        this.function = function;
        return this;
    }
    public Function getFunction() {
        return this.function;
    }

    public FunctionComputeConfiguration setID(String ID) {
        this.ID = ID;
        return this;
    }
    public String getID() {
        return this.ID;
    }

    public static class Key extends TeaModel {
        @NameInMap("Prefix")
        public String prefix;

        @NameInMap("Suffix")
        public String suffix;

        public static Key build(java.util.Map<String, ?> map) throws Exception {
            Key self = new Key();
            return TeaModel.build(map, self);
        }

        public Key setPrefix(String prefix) {
            this.prefix = prefix;
            return this;
        }
        public String getPrefix() {
            return this.prefix;
        }

        public Key setSuffix(String suffix) {
            this.suffix = suffix;
            return this;
        }
        public String getSuffix() {
            return this.suffix;
        }

    }

    public static class Filter extends TeaModel {
        @NameInMap("Key")
        public Key key;

        public static Filter build(java.util.Map<String, ?> map) throws Exception {
            Filter self = new Filter();
            return TeaModel.build(map, self);
        }

        public Filter setKey(Key key) {
            this.key = key;
            return this;
        }
        public Key getKey() {
            return this.key;
        }

    }

    public static class Function extends TeaModel {
        @NameInMap("Arn")
        public String arn;

        @NameInMap("AssumeRole")
        public String assumeRole;

        public static Function build(java.util.Map<String, ?> map) throws Exception {
            Function self = new Function();
            return TeaModel.build(map, self);
        }

        public Function setArn(String arn) {
            this.arn = arn;
            return this;
        }
        public String getArn() {
            return this.arn;
        }

        public Function setAssumeRole(String assumeRole) {
            this.assumeRole = assumeRole;
            return this;
        }
        public String getAssumeRole() {
            return this.assumeRole;
        }

    }

}
