// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class InventoryOSSBucketDestination extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>100000000000000</p>
     */
    @NameInMap("AccountId")
    public String accountId;

    /**
     * <strong>example:</strong>
     * <p>acs:oss:::bucket_0001</p>
     */
    @NameInMap("Bucket")
    public String bucket;

    @NameInMap("Encryption")
    public InventoryEncryption encryption;

    @NameInMap("Format")
    public String format;

    /**
     * <strong>example:</strong>
     * <p>prefix1/</p>
     */
    @NameInMap("Prefix")
    public String prefix;

    /**
     * <strong>example:</strong>
     * <p>acs:ram::100000000000000:role/AliyunOSSRole</p>
     */
    @NameInMap("RoleArn")
    public String roleArn;

    public static InventoryOSSBucketDestination build(java.util.Map<String, ?> map) throws Exception {
        InventoryOSSBucketDestination self = new InventoryOSSBucketDestination();
        return TeaModel.build(map, self);
    }

    public InventoryOSSBucketDestination setAccountId(String accountId) {
        this.accountId = accountId;
        return this;
    }
    public String getAccountId() {
        return this.accountId;
    }

    public InventoryOSSBucketDestination setBucket(String bucket) {
        this.bucket = bucket;
        return this;
    }
    public String getBucket() {
        return this.bucket;
    }

    public InventoryOSSBucketDestination setEncryption(InventoryEncryption encryption) {
        this.encryption = encryption;
        return this;
    }
    public InventoryEncryption getEncryption() {
        return this.encryption;
    }

    public InventoryOSSBucketDestination setFormat(String format) {
        this.format = format;
        return this;
    }
    public String getFormat() {
        return this.format;
    }

    public InventoryOSSBucketDestination setPrefix(String prefix) {
        this.prefix = prefix;
        return this;
    }
    public String getPrefix() {
        return this.prefix;
    }

    public InventoryOSSBucketDestination setRoleArn(String roleArn) {
        this.roleArn = roleArn;
        return this;
    }
    public String getRoleArn() {
        return this.roleArn;
    }

}
