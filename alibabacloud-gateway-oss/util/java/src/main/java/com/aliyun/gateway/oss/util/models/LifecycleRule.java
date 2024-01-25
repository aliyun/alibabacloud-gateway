// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class LifecycleRule extends TeaModel {
    @NameInMap("AbortMultipartUpload")
    public LifecycleRuleAbortMultipartUpload abortMultipartUpload;

    @NameInMap("Expiration")
    public LifecycleRuleExpiration expiration;

    @NameInMap("Filter")
    public LifecycleRuleFilter filter;

    @NameInMap("ID")
    public String ID;

    @NameInMap("NoncurrentVersionExpiration")
    public LifecycleRuleNoncurrentVersionExpiration noncurrentVersionExpiration;

    @NameInMap("NoncurrentVersionTransition")
    public java.util.List<LifecycleRuleNoncurrentVersionTransition> noncurrentVersionTransition;

    @NameInMap("Prefix")
    public String prefix;

    @NameInMap("Status")
    public String status;

    @NameInMap("Tag")
    public java.util.List<Tag> tag;

    @NameInMap("Transition")
    public java.util.List<LifecycleRuleTransition> transition;

    public static LifecycleRule build(java.util.Map<String, ?> map) throws Exception {
        LifecycleRule self = new LifecycleRule();
        return TeaModel.build(map, self);
    }

    public LifecycleRule setAbortMultipartUpload(LifecycleRuleAbortMultipartUpload abortMultipartUpload) {
        this.abortMultipartUpload = abortMultipartUpload;
        return this;
    }
    public LifecycleRuleAbortMultipartUpload getAbortMultipartUpload() {
        return this.abortMultipartUpload;
    }

    public LifecycleRule setExpiration(LifecycleRuleExpiration expiration) {
        this.expiration = expiration;
        return this;
    }
    public LifecycleRuleExpiration getExpiration() {
        return this.expiration;
    }

    public LifecycleRule setFilter(LifecycleRuleFilter filter) {
        this.filter = filter;
        return this;
    }
    public LifecycleRuleFilter getFilter() {
        return this.filter;
    }

    public LifecycleRule setID(String ID) {
        this.ID = ID;
        return this;
    }
    public String getID() {
        return this.ID;
    }

    public LifecycleRule setNoncurrentVersionExpiration(LifecycleRuleNoncurrentVersionExpiration noncurrentVersionExpiration) {
        this.noncurrentVersionExpiration = noncurrentVersionExpiration;
        return this;
    }
    public LifecycleRuleNoncurrentVersionExpiration getNoncurrentVersionExpiration() {
        return this.noncurrentVersionExpiration;
    }

    public LifecycleRule setNoncurrentVersionTransition(java.util.List<LifecycleRuleNoncurrentVersionTransition> noncurrentVersionTransition) {
        this.noncurrentVersionTransition = noncurrentVersionTransition;
        return this;
    }
    public java.util.List<LifecycleRuleNoncurrentVersionTransition> getNoncurrentVersionTransition() {
        return this.noncurrentVersionTransition;
    }

    public LifecycleRule setPrefix(String prefix) {
        this.prefix = prefix;
        return this;
    }
    public String getPrefix() {
        return this.prefix;
    }

    public LifecycleRule setStatus(String status) {
        this.status = status;
        return this;
    }
    public String getStatus() {
        return this.status;
    }

    public LifecycleRule setTag(java.util.List<Tag> tag) {
        this.tag = tag;
        return this;
    }
    public java.util.List<Tag> getTag() {
        return this.tag;
    }

    public LifecycleRule setTransition(java.util.List<LifecycleRuleTransition> transition) {
        this.transition = transition;
        return this;
    }
    public java.util.List<LifecycleRuleTransition> getTransition() {
        return this.transition;
    }

    public static class LifecycleRuleAbortMultipartUpload extends TeaModel {
        @NameInMap("CreatedBeforeDate")
        public String createdBeforeDate;

        @NameInMap("Days")
        public Integer days;

        public static LifecycleRuleAbortMultipartUpload build(java.util.Map<String, ?> map) throws Exception {
            LifecycleRuleAbortMultipartUpload self = new LifecycleRuleAbortMultipartUpload();
            return TeaModel.build(map, self);
        }

        public LifecycleRuleAbortMultipartUpload setCreatedBeforeDate(String createdBeforeDate) {
            this.createdBeforeDate = createdBeforeDate;
            return this;
        }
        public String getCreatedBeforeDate() {
            return this.createdBeforeDate;
        }

        public LifecycleRuleAbortMultipartUpload setDays(Integer days) {
            this.days = days;
            return this;
        }
        public Integer getDays() {
            return this.days;
        }

    }

    public static class LifecycleRuleExpiration extends TeaModel {
        @NameInMap("CreatedBeforeDate")
        public String createdBeforeDate;

        @NameInMap("Days")
        public Integer days;

        @NameInMap("ExpiredObjectDeleteMarker")
        public Boolean expiredObjectDeleteMarker;

        public static LifecycleRuleExpiration build(java.util.Map<String, ?> map) throws Exception {
            LifecycleRuleExpiration self = new LifecycleRuleExpiration();
            return TeaModel.build(map, self);
        }

        public LifecycleRuleExpiration setCreatedBeforeDate(String createdBeforeDate) {
            this.createdBeforeDate = createdBeforeDate;
            return this;
        }
        public String getCreatedBeforeDate() {
            return this.createdBeforeDate;
        }

        public LifecycleRuleExpiration setDays(Integer days) {
            this.days = days;
            return this;
        }
        public Integer getDays() {
            return this.days;
        }

        public LifecycleRuleExpiration setExpiredObjectDeleteMarker(Boolean expiredObjectDeleteMarker) {
            this.expiredObjectDeleteMarker = expiredObjectDeleteMarker;
            return this;
        }
        public Boolean getExpiredObjectDeleteMarker() {
            return this.expiredObjectDeleteMarker;
        }

    }

    public static class LifecycleRuleFilterNot extends TeaModel {
        @NameInMap("Prefix")
        public String prefix;

        @NameInMap("Tag")
        public Tag tag;

        public static LifecycleRuleFilterNot build(java.util.Map<String, ?> map) throws Exception {
            LifecycleRuleFilterNot self = new LifecycleRuleFilterNot();
            return TeaModel.build(map, self);
        }

        public LifecycleRuleFilterNot setPrefix(String prefix) {
            this.prefix = prefix;
            return this;
        }
        public String getPrefix() {
            return this.prefix;
        }

        public LifecycleRuleFilterNot setTag(Tag tag) {
            this.tag = tag;
            return this;
        }
        public Tag getTag() {
            return this.tag;
        }

    }

    public static class LifecycleRuleFilter extends TeaModel {
        @NameInMap("Not")
        public LifecycleRuleFilterNot not;

        public static LifecycleRuleFilter build(java.util.Map<String, ?> map) throws Exception {
            LifecycleRuleFilter self = new LifecycleRuleFilter();
            return TeaModel.build(map, self);
        }

        public LifecycleRuleFilter setNot(LifecycleRuleFilterNot not) {
            this.not = not;
            return this;
        }
        public LifecycleRuleFilterNot getNot() {
            return this.not;
        }

    }

    public static class LifecycleRuleNoncurrentVersionExpiration extends TeaModel {
        @NameInMap("NoncurrentDays")
        public Integer noncurrentDays;

        public static LifecycleRuleNoncurrentVersionExpiration build(java.util.Map<String, ?> map) throws Exception {
            LifecycleRuleNoncurrentVersionExpiration self = new LifecycleRuleNoncurrentVersionExpiration();
            return TeaModel.build(map, self);
        }

        public LifecycleRuleNoncurrentVersionExpiration setNoncurrentDays(Integer noncurrentDays) {
            this.noncurrentDays = noncurrentDays;
            return this;
        }
        public Integer getNoncurrentDays() {
            return this.noncurrentDays;
        }

    }

    public static class LifecycleRuleNoncurrentVersionTransition extends TeaModel {
        @NameInMap("AllowSmallFile")
        public Boolean allowSmallFile;

        @NameInMap("IsAccessTime")
        public Boolean isAccessTime;

        @NameInMap("NoncurrentDays")
        public Integer noncurrentDays;

        @NameInMap("ReturnToStdWhenVisit")
        public Boolean returnToStdWhenVisit;

        @NameInMap("StorageClass")
        public String storageClass;

        public static LifecycleRuleNoncurrentVersionTransition build(java.util.Map<String, ?> map) throws Exception {
            LifecycleRuleNoncurrentVersionTransition self = new LifecycleRuleNoncurrentVersionTransition();
            return TeaModel.build(map, self);
        }

        public LifecycleRuleNoncurrentVersionTransition setAllowSmallFile(Boolean allowSmallFile) {
            this.allowSmallFile = allowSmallFile;
            return this;
        }
        public Boolean getAllowSmallFile() {
            return this.allowSmallFile;
        }

        public LifecycleRuleNoncurrentVersionTransition setIsAccessTime(Boolean isAccessTime) {
            this.isAccessTime = isAccessTime;
            return this;
        }
        public Boolean getIsAccessTime() {
            return this.isAccessTime;
        }

        public LifecycleRuleNoncurrentVersionTransition setNoncurrentDays(Integer noncurrentDays) {
            this.noncurrentDays = noncurrentDays;
            return this;
        }
        public Integer getNoncurrentDays() {
            return this.noncurrentDays;
        }

        public LifecycleRuleNoncurrentVersionTransition setReturnToStdWhenVisit(Boolean returnToStdWhenVisit) {
            this.returnToStdWhenVisit = returnToStdWhenVisit;
            return this;
        }
        public Boolean getReturnToStdWhenVisit() {
            return this.returnToStdWhenVisit;
        }

        public LifecycleRuleNoncurrentVersionTransition setStorageClass(String storageClass) {
            this.storageClass = storageClass;
            return this;
        }
        public String getStorageClass() {
            return this.storageClass;
        }

    }

    public static class LifecycleRuleTransition extends TeaModel {
        @NameInMap("AllowSmallFile")
        public Boolean allowSmallFile;

        @NameInMap("CreatedBeforeDate")
        public String createdBeforeDate;

        @NameInMap("Days")
        public Integer days;

        @NameInMap("IsAccessTime")
        public Boolean isAccessTime;

        @NameInMap("ReturnToStdWhenVisit")
        public Boolean returnToStdWhenVisit;

        @NameInMap("StorageClass")
        public String storageClass;

        public static LifecycleRuleTransition build(java.util.Map<String, ?> map) throws Exception {
            LifecycleRuleTransition self = new LifecycleRuleTransition();
            return TeaModel.build(map, self);
        }

        public LifecycleRuleTransition setAllowSmallFile(Boolean allowSmallFile) {
            this.allowSmallFile = allowSmallFile;
            return this;
        }
        public Boolean getAllowSmallFile() {
            return this.allowSmallFile;
        }

        public LifecycleRuleTransition setCreatedBeforeDate(String createdBeforeDate) {
            this.createdBeforeDate = createdBeforeDate;
            return this;
        }
        public String getCreatedBeforeDate() {
            return this.createdBeforeDate;
        }

        public LifecycleRuleTransition setDays(Integer days) {
            this.days = days;
            return this;
        }
        public Integer getDays() {
            return this.days;
        }

        public LifecycleRuleTransition setIsAccessTime(Boolean isAccessTime) {
            this.isAccessTime = isAccessTime;
            return this;
        }
        public Boolean getIsAccessTime() {
            return this.isAccessTime;
        }

        public LifecycleRuleTransition setReturnToStdWhenVisit(Boolean returnToStdWhenVisit) {
            this.returnToStdWhenVisit = returnToStdWhenVisit;
            return this;
        }
        public Boolean getReturnToStdWhenVisit() {
            return this.returnToStdWhenVisit;
        }

        public LifecycleRuleTransition setStorageClass(String storageClass) {
            this.storageClass = storageClass;
            return this;
        }
        public String getStorageClass() {
            return this.storageClass;
        }

    }

}
