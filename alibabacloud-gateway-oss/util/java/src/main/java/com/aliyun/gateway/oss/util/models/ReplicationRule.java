// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ReplicationRule extends TeaModel {
    @NameInMap("Action")
    public String action;

    @NameInMap("Destination")
    public ReplicationDestination destination;

    @NameInMap("EncryptionConfiguration")
    public ReplicationEncryptionConfiguration encryptionConfiguration;

    @NameInMap("HistoricalObjectReplication")
    public String historicalObjectReplication;

    @NameInMap("ID")
    public String ID;

    @NameInMap("PrefixSet")
    public ReplicationPrefixSet prefixSet;

    @NameInMap("RTC")
    public RTC RTC;

    @NameInMap("SourceSelectionCriteria")
    public ReplicationSourceSelectionCriteria sourceSelectionCriteria;

    @NameInMap("Status")
    public String status;

    @NameInMap("SyncRole")
    public String syncRole;

    public static ReplicationRule build(java.util.Map<String, ?> map) throws Exception {
        ReplicationRule self = new ReplicationRule();
        return TeaModel.build(map, self);
    }

    public ReplicationRule setAction(String action) {
        this.action = action;
        return this;
    }
    public String getAction() {
        return this.action;
    }

    public ReplicationRule setDestination(ReplicationDestination destination) {
        this.destination = destination;
        return this;
    }
    public ReplicationDestination getDestination() {
        return this.destination;
    }

    public ReplicationRule setEncryptionConfiguration(ReplicationEncryptionConfiguration encryptionConfiguration) {
        this.encryptionConfiguration = encryptionConfiguration;
        return this;
    }
    public ReplicationEncryptionConfiguration getEncryptionConfiguration() {
        return this.encryptionConfiguration;
    }

    public ReplicationRule setHistoricalObjectReplication(String historicalObjectReplication) {
        this.historicalObjectReplication = historicalObjectReplication;
        return this;
    }
    public String getHistoricalObjectReplication() {
        return this.historicalObjectReplication;
    }

    public ReplicationRule setID(String ID) {
        this.ID = ID;
        return this;
    }
    public String getID() {
        return this.ID;
    }

    public ReplicationRule setPrefixSet(ReplicationPrefixSet prefixSet) {
        this.prefixSet = prefixSet;
        return this;
    }
    public ReplicationPrefixSet getPrefixSet() {
        return this.prefixSet;
    }

    public ReplicationRule setRTC(RTC RTC) {
        this.RTC = RTC;
        return this;
    }
    public RTC getRTC() {
        return this.RTC;
    }

    public ReplicationRule setSourceSelectionCriteria(ReplicationSourceSelectionCriteria sourceSelectionCriteria) {
        this.sourceSelectionCriteria = sourceSelectionCriteria;
        return this;
    }
    public ReplicationSourceSelectionCriteria getSourceSelectionCriteria() {
        return this.sourceSelectionCriteria;
    }

    public ReplicationRule setStatus(String status) {
        this.status = status;
        return this;
    }
    public String getStatus() {
        return this.status;
    }

    public ReplicationRule setSyncRole(String syncRole) {
        this.syncRole = syncRole;
        return this;
    }
    public String getSyncRole() {
        return this.syncRole;
    }

}
