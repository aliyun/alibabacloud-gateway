// This file is auto-generated, don't edit it. Thanks.
com.aliyun.gateway.oss.util.models

import com.aliyun.tea.*;

public class BucketInfo extends TeaModel {
    @NameInMap("Bucket")
    public BucketInfoBucket bucket;

    public static BucketInfo build(java.util.Map<String, ?> map) throws Exception {
        BucketInfo self = new BucketInfo();
        return TeaModel.build(map, self);
    }

    public BucketInfo setBucket(BucketInfoBucket bucket) {
        this.bucket = bucket;
        return this;
    }
    public BucketInfoBucket getBucket() {
        return this.bucket;
    }

    public static class BucketInfoBucketServerSideEncryptionRule extends TeaModel {
        @NameInMap("KMSDataEncryption")
        public String KMSDataEncryption;

        @NameInMap("KMSMasterKeyID")
        public String KMSMasterKeyID;

        @NameInMap("SSEAlgorithm")
        public String SSEAlgorithm;

        public static BucketInfoBucketServerSideEncryptionRule build(java.util.Map<String, ?> map) throws Exception {
            BucketInfoBucketServerSideEncryptionRule self = new BucketInfoBucketServerSideEncryptionRule();
            return TeaModel.build(map, self);
        }

        public BucketInfoBucketServerSideEncryptionRule setKMSDataEncryption(String KMSDataEncryption) {
            this.KMSDataEncryption = KMSDataEncryption;
            return this;
        }
        public String getKMSDataEncryption() {
            return this.KMSDataEncryption;
        }

        public BucketInfoBucketServerSideEncryptionRule setKMSMasterKeyID(String KMSMasterKeyID) {
            this.KMSMasterKeyID = KMSMasterKeyID;
            return this;
        }
        public String getKMSMasterKeyID() {
            return this.KMSMasterKeyID;
        }

        public BucketInfoBucketServerSideEncryptionRule setSSEAlgorithm(String SSEAlgorithm) {
            this.SSEAlgorithm = SSEAlgorithm;
            return this;
        }
        public String getSSEAlgorithm() {
            return this.SSEAlgorithm;
        }

    }

    public static class BucketInfoBucket extends TeaModel {
        @NameInMap("AccessControlList")
        public AccessControlList accessControlList;

        @NameInMap("AccessMonitor")
        public String accessMonitor;

        @NameInMap("BucketPolicy")
        public LoggingEnabled bucketPolicy;

        @NameInMap("CreationDate")
        public String creationDate;

        @NameInMap("CrossRegionReplication")
        public String crossRegionReplication;

        @NameInMap("DataRedundancyType")
        public String dataRedundancyType;

        @NameInMap("ExtranetEndpoint")
        public String extranetEndpoint;

        @NameInMap("IntranetEndpoint")
        public String intranetEndpoint;

        @NameInMap("Location")
        public String location;

        @NameInMap("Name")
        public String name;

        @NameInMap("Owner")
        public Owner owner;

        @NameInMap("ResourceGroupId")
        public String resourceGroupId;

        @NameInMap("ServerSideEncryptionRule")
        public BucketInfoBucketServerSideEncryptionRule serverSideEncryptionRule;

        @NameInMap("StorageClass")
        public String storageClass;

        @NameInMap("TransferAcceleration")
        public String transferAcceleration;

        @NameInMap("Versioning")
        public String versioning;

        public static BucketInfoBucket build(java.util.Map<String, ?> map) throws Exception {
            BucketInfoBucket self = new BucketInfoBucket();
            return TeaModel.build(map, self);
        }

        public BucketInfoBucket setAccessControlList(AccessControlList accessControlList) {
            this.accessControlList = accessControlList;
            return this;
        }
        public AccessControlList getAccessControlList() {
            return this.accessControlList;
        }

        public BucketInfoBucket setAccessMonitor(String accessMonitor) {
            this.accessMonitor = accessMonitor;
            return this;
        }
        public String getAccessMonitor() {
            return this.accessMonitor;
        }

        public BucketInfoBucket setBucketPolicy(LoggingEnabled bucketPolicy) {
            this.bucketPolicy = bucketPolicy;
            return this;
        }
        public LoggingEnabled getBucketPolicy() {
            return this.bucketPolicy;
        }

        public BucketInfoBucket setCreationDate(String creationDate) {
            this.creationDate = creationDate;
            return this;
        }
        public String getCreationDate() {
            return this.creationDate;
        }

        public BucketInfoBucket setCrossRegionReplication(String crossRegionReplication) {
            this.crossRegionReplication = crossRegionReplication;
            return this;
        }
        public String getCrossRegionReplication() {
            return this.crossRegionReplication;
        }

        public BucketInfoBucket setDataRedundancyType(String dataRedundancyType) {
            this.dataRedundancyType = dataRedundancyType;
            return this;
        }
        public String getDataRedundancyType() {
            return this.dataRedundancyType;
        }

        public BucketInfoBucket setExtranetEndpoint(String extranetEndpoint) {
            this.extranetEndpoint = extranetEndpoint;
            return this;
        }
        public String getExtranetEndpoint() {
            return this.extranetEndpoint;
        }

        public BucketInfoBucket setIntranetEndpoint(String intranetEndpoint) {
            this.intranetEndpoint = intranetEndpoint;
            return this;
        }
        public String getIntranetEndpoint() {
            return this.intranetEndpoint;
        }

        public BucketInfoBucket setLocation(String location) {
            this.location = location;
            return this;
        }
        public String getLocation() {
            return this.location;
        }

        public BucketInfoBucket setName(String name) {
            this.name = name;
            return this;
        }
        public String getName() {
            return this.name;
        }

        public BucketInfoBucket setOwner(Owner owner) {
            this.owner = owner;
            return this;
        }
        public Owner getOwner() {
            return this.owner;
        }

        public BucketInfoBucket setResourceGroupId(String resourceGroupId) {
            this.resourceGroupId = resourceGroupId;
            return this;
        }
        public String getResourceGroupId() {
            return this.resourceGroupId;
        }

        public BucketInfoBucket setServerSideEncryptionRule(BucketInfoBucketServerSideEncryptionRule serverSideEncryptionRule) {
            this.serverSideEncryptionRule = serverSideEncryptionRule;
            return this;
        }
        public BucketInfoBucketServerSideEncryptionRule getServerSideEncryptionRule() {
            return this.serverSideEncryptionRule;
        }

        public BucketInfoBucket setStorageClass(String storageClass) {
            this.storageClass = storageClass;
            return this;
        }
        public String getStorageClass() {
            return this.storageClass;
        }

        public BucketInfoBucket setTransferAcceleration(String transferAcceleration) {
            this.transferAcceleration = transferAcceleration;
            return this;
        }
        public String getTransferAcceleration() {
            return this.transferAcceleration;
        }

        public BucketInfoBucket setVersioning(String versioning) {
            this.versioning = versioning;
            return this;
        }
        public String getVersioning() {
            return this.versioning;
        }

    }

}
