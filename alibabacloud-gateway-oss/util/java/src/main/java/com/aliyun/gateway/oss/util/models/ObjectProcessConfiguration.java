// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ObjectProcessConfiguration extends TeaModel {
    @NameInMap("AllowedFeatures")
    public AllowedFeatures allowedFeatures;

    @NameInMap("TransformationConfigurations")
    public TransformationConfigurations transformationConfigurations;

    public static ObjectProcessConfiguration build(java.util.Map<String, ?> map) throws Exception {
        ObjectProcessConfiguration self = new ObjectProcessConfiguration();
        return TeaModel.build(map, self);
    }

    public ObjectProcessConfiguration setAllowedFeatures(AllowedFeatures allowedFeatures) {
        this.allowedFeatures = allowedFeatures;
        return this;
    }
    public AllowedFeatures getAllowedFeatures() {
        return this.allowedFeatures;
    }

    public ObjectProcessConfiguration setTransformationConfigurations(TransformationConfigurations transformationConfigurations) {
        this.transformationConfigurations = transformationConfigurations;
        return this;
    }
    public TransformationConfigurations getTransformationConfigurations() {
        return this.transformationConfigurations;
    }

    public static class AllowedFeatures extends TeaModel {
        @NameInMap("AllowedFeature")
        public java.util.List<String> allowedFeature;

        public static AllowedFeatures build(java.util.Map<String, ?> map) throws Exception {
            AllowedFeatures self = new AllowedFeatures();
            return TeaModel.build(map, self);
        }

        public AllowedFeatures setAllowedFeature(java.util.List<String> allowedFeature) {
            this.allowedFeature = allowedFeature;
            return this;
        }
        public java.util.List<String> getAllowedFeature() {
            return this.allowedFeature;
        }

    }

    public static class Actions extends TeaModel {
        @NameInMap("Action")
        public java.util.List<String> action;

        public static Actions build(java.util.Map<String, ?> map) throws Exception {
            Actions self = new Actions();
            return TeaModel.build(map, self);
        }

        public Actions setAction(java.util.List<String> action) {
            this.action = action;
            return this;
        }
        public java.util.List<String> getAction() {
            return this.action;
        }

    }

    public static class CustomForwardHeaders extends TeaModel {
        @NameInMap("CustomForwardHeader")
        public java.util.List<String> customForwardHeader;

        public static CustomForwardHeaders build(java.util.Map<String, ?> map) throws Exception {
            CustomForwardHeaders self = new CustomForwardHeaders();
            return TeaModel.build(map, self);
        }

        public CustomForwardHeaders setCustomForwardHeader(java.util.List<String> customForwardHeader) {
            this.customForwardHeader = customForwardHeader;
            return this;
        }
        public java.util.List<String> getCustomForwardHeader() {
            return this.customForwardHeader;
        }

    }

    public static class AdditionalFeatures extends TeaModel {
        @NameInMap("CustomForwardHeaders")
        public CustomForwardHeaders customForwardHeaders;

        public static AdditionalFeatures build(java.util.Map<String, ?> map) throws Exception {
            AdditionalFeatures self = new AdditionalFeatures();
            return TeaModel.build(map, self);
        }

        public AdditionalFeatures setCustomForwardHeaders(CustomForwardHeaders customForwardHeaders) {
            this.customForwardHeaders = customForwardHeaders;
            return this;
        }
        public CustomForwardHeaders getCustomForwardHeaders() {
            return this.customForwardHeaders;
        }

    }

    public static class FunctionCompute extends TeaModel {
        @NameInMap("FunctionArn")
        public String functionArn;

        @NameInMap("FunctionAssumeRoleArn")
        public String functionAssumeRoleArn;

        public static FunctionCompute build(java.util.Map<String, ?> map) throws Exception {
            FunctionCompute self = new FunctionCompute();
            return TeaModel.build(map, self);
        }

        public FunctionCompute setFunctionArn(String functionArn) {
            this.functionArn = functionArn;
            return this;
        }
        public String getFunctionArn() {
            return this.functionArn;
        }

        public FunctionCompute setFunctionAssumeRoleArn(String functionAssumeRoleArn) {
            this.functionAssumeRoleArn = functionAssumeRoleArn;
            return this;
        }
        public String getFunctionAssumeRoleArn() {
            return this.functionAssumeRoleArn;
        }

    }

    public static class ContentTransformation extends TeaModel {
        @NameInMap("AdditionalFeatures")
        public AdditionalFeatures additionalFeatures;

        @NameInMap("FunctionCompute")
        public FunctionCompute functionCompute;

        public static ContentTransformation build(java.util.Map<String, ?> map) throws Exception {
            ContentTransformation self = new ContentTransformation();
            return TeaModel.build(map, self);
        }

        public ContentTransformation setAdditionalFeatures(AdditionalFeatures additionalFeatures) {
            this.additionalFeatures = additionalFeatures;
            return this;
        }
        public AdditionalFeatures getAdditionalFeatures() {
            return this.additionalFeatures;
        }

        public ContentTransformation setFunctionCompute(FunctionCompute functionCompute) {
            this.functionCompute = functionCompute;
            return this;
        }
        public FunctionCompute getFunctionCompute() {
            return this.functionCompute;
        }

    }

    public static class TransformationConfiguration extends TeaModel {
        @NameInMap("Actions")
        public Actions actions;

        @NameInMap("ContentTransformation")
        public ContentTransformation contentTransformation;

        public static TransformationConfiguration build(java.util.Map<String, ?> map) throws Exception {
            TransformationConfiguration self = new TransformationConfiguration();
            return TeaModel.build(map, self);
        }

        public TransformationConfiguration setActions(Actions actions) {
            this.actions = actions;
            return this;
        }
        public Actions getActions() {
            return this.actions;
        }

        public TransformationConfiguration setContentTransformation(ContentTransformation contentTransformation) {
            this.contentTransformation = contentTransformation;
            return this;
        }
        public ContentTransformation getContentTransformation() {
            return this.contentTransformation;
        }

    }

    public static class TransformationConfigurations extends TeaModel {
        @NameInMap("TransformationConfiguration")
        public java.util.List<TransformationConfiguration> transformationConfiguration;

        public static TransformationConfigurations build(java.util.Map<String, ?> map) throws Exception {
            TransformationConfigurations self = new TransformationConfigurations();
            return TeaModel.build(map, self);
        }

        public TransformationConfigurations setTransformationConfiguration(java.util.List<TransformationConfiguration> transformationConfiguration) {
            this.transformationConfiguration = transformationConfiguration;
            return this;
        }
        public java.util.List<TransformationConfiguration> getTransformationConfiguration() {
            return this.transformationConfiguration;
        }

    }

}
