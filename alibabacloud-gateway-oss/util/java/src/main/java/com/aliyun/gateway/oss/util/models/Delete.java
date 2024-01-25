// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class Delete extends TeaModel {
    @NameInMap("Object")
    public java.util.List<ObjectIdentifier> object;

    @NameInMap("Quiet")
    public Boolean quiet;

    public static Delete build(java.util.Map<String, ?> map) throws Exception {
        Delete self = new Delete();
        return TeaModel.build(map, self);
    }

    public Delete setObject(java.util.List<ObjectIdentifier> object) {
        this.object = object;
        return this;
    }
    public java.util.List<ObjectIdentifier> getObject() {
        return this.object;
    }

    public Delete setQuiet(Boolean quiet) {
        this.quiet = quiet;
        return this;
    }
    public Boolean getQuiet() {
        return this.quiet;
    }

}
