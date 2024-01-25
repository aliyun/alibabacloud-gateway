// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class ObjectProcessConfiguration extends TeaModel {
    @NameInMap("AllowedFeatures")
    public ObjectProcessConfigurationAllowedFeatures allowedFeatures;

    @NameInMap("TransformationConfigurations")
    public ObjectProcessConfigurationTransformationConfigurations transformationConfigurations;

    public static ObjectProcessConfiguration build(java.util.Map<String, ?> map) throws Exception {
        ObjectProcessConfiguration self = new ObjectProcessConfiguration();
        return TeaModel.build(map, self);
    }

    public ObjectProcessConfiguration setAllowedFeatures(ObjectProcessConfigurationAllowedFeatures allowedFeatures) {
        this.allowedFeatures = allowedFeatures;
        return this;
    }
    public ObjectProcessConfigurationAllowedFeatures getAllowedFeatures() {
        return this.allowedFeatures;
    }

    public ObjectProcessConfiguration setTransformationConfigurations(ObjectProcessConfigurationTransformationConfigurations transformationConfigurations) {
        this.transformationConfigurations = transformationConfigurations;
        return this;
    }
    public ObjectProcessConfigurationTransformationConfigurations getTransformationConfigurations() {
        return this.transformationConfigurations;
    }

    public static class ObjectProcessConfigurationAllowedFeatures extends TeaModel {
        @NameInMap("AllowedFeature")
        public java.util.List<String> allowedFeature;

        public static ObjectProcessConfigurationAllowedFeatures build(java.util.Map<String, ?> map) throws Exception {
            ObjectProcessConfigurationAllowedFeatures self = new ObjectProcessConfigurationAllowedFeatures();
            return TeaModel.build(map, self);
        }

        public ObjectProcessConfigurationAllowedFeatures setAllowedFeature(java.util.List<String> allowedFeature) {
            this.allowedFeature = allowedFeature;
            return this;
        }
        public java.util.List<String> getAllowedFeature() {
            return this.allowedFeature;
        }

    }

    public static class ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationActions extends TeaModel {
        @NameInMap("Action")
        public java.util.List<String> action;

        public static ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationActions build(java.util.Map<String, ?> map) throws Exception {
            ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationActions self = new ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationActions();
            return TeaModel.build(map, self);
        }

        public ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationActions setAction(java.util.List<String> action) {
            this.action = action;
            return this;
        }
        public java.util.List<String> getAction() {
            return this.action;
        }

    }

    public static class ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationFunctionCompute extends TeaModel {
        @NameInMap("FunctionArn")
        public String functionArn;

        @NameInMap("FunctionAssumeRoleArn")
        public String functionAssumeRoleArn;

        public static ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationFunctionCompute build(java.util.Map<String, ?> map) throws Exception {
            ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationFunctionCompute self = new ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationFunctionCompute();
            return TeaModel.build(map, self);
        }

        public ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationFunctionCompute setFunctionArn(String functionArn) {
            this.functionArn = functionArn;
            return this;
        }
        public String getFunctionArn() {
            return this.functionArn;
        }

        public ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationFunctionCompute setFunctionAssumeRoleArn(String functionAssumeRoleArn) {
            this.functionAssumeRoleArn = functionAssumeRoleArn;
            return this;
        }
        public String getFunctionAssumeRoleArn() {
            return this.functionAssumeRoleArn;
        }

    }

    public static class ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformation extends TeaModel {
        @NameInMap("FunctionCompute")
        public ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationFunctionCompute functionCompute;

        public static ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformation build(java.util.Map<String, ?> map) throws Exception {
            ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformation self = new ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformation();
            return TeaModel.build(map, self);
        }

        public ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformation setFunctionCompute(ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationFunctionCompute functionCompute) {
            this.functionCompute = functionCompute;
            return this;
        }
        public ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformationFunctionCompute getFunctionCompute() {
            return this.functionCompute;
        }

    }

    public static class ObjectProcessConfigurationTransformationConfigurationsTransformationConfiguration extends TeaModel {
        @NameInMap("Actions")
        public ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationActions actions;

        @NameInMap("ContentTransformation")
        public ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformation contentTransformation;

        public static ObjectProcessConfigurationTransformationConfigurationsTransformationConfiguration build(java.util.Map<String, ?> map) throws Exception {
            ObjectProcessConfigurationTransformationConfigurationsTransformationConfiguration self = new ObjectProcessConfigurationTransformationConfigurationsTransformationConfiguration();
            return TeaModel.build(map, self);
        }

        public ObjectProcessConfigurationTransformationConfigurationsTransformationConfiguration setActions(ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationActions actions) {
            this.actions = actions;
            return this;
        }
        public ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationActions getActions() {
            return this.actions;
        }

        public ObjectProcessConfigurationTransformationConfigurationsTransformationConfiguration setContentTransformation(ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformation contentTransformation) {
            this.contentTransformation = contentTransformation;
            return this;
        }
        public ObjectProcessConfigurationTransformationConfigurationsTransformationConfigurationContentTransformation getContentTransformation() {
            return this.contentTransformation;
        }

    }

    public static class ObjectProcessConfigurationTransformationConfigurations extends TeaModel {
        @NameInMap("TransformationConfiguration")
        public java.util.List<ObjectProcessConfigurationTransformationConfigurationsTransformationConfiguration> transformationConfiguration;

        public static ObjectProcessConfigurationTransformationConfigurations build(java.util.Map<String, ?> map) throws Exception {
            ObjectProcessConfigurationTransformationConfigurations self = new ObjectProcessConfigurationTransformationConfigurations();
            return TeaModel.build(map, self);
        }

        public ObjectProcessConfigurationTransformationConfigurations setTransformationConfiguration(java.util.List<ObjectProcessConfigurationTransformationConfigurationsTransformationConfiguration> transformationConfiguration) {
            this.transformationConfiguration = transformationConfiguration;
            return this;
        }
        public java.util.List<ObjectProcessConfigurationTransformationConfigurationsTransformationConfiguration> getTransformationConfiguration() {
            return this.transformationConfiguration;
        }

    }

}
