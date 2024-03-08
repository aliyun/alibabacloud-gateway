// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class RestoreRequest extends TeaModel {
    @NameInMap("Days")
    public Long days;

    @NameInMap("JobParameters")
    public RestoreRequestJobParameters jobParameters;

    public static RestoreRequest build(java.util.Map<String, ?> map) throws Exception {
        RestoreRequest self = new RestoreRequest();
        return TeaModel.build(map, self);
    }

    public RestoreRequest setDays(Long days) {
        this.days = days;
        return this;
    }
    public Long getDays() {
        return this.days;
    }

    public RestoreRequest setJobParameters(RestoreRequestJobParameters jobParameters) {
        this.jobParameters = jobParameters;
        return this;
    }
    public RestoreRequestJobParameters getJobParameters() {
        return this.jobParameters;
    }

    public static class RestoreRequestJobParameters extends TeaModel {
        @NameInMap("Tier")
        public String tier;

        public static RestoreRequestJobParameters build(java.util.Map<String, ?> map) throws Exception {
            RestoreRequestJobParameters self = new RestoreRequestJobParameters();
            return TeaModel.build(map, self);
        }

        public RestoreRequestJobParameters setTier(String tier) {
            this.tier = tier;
            return this;
        }
        public String getTier() {
            return this.tier;
        }

    }

}
