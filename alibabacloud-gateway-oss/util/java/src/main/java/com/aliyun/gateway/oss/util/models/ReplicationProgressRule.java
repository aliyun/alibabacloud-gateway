// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.gateway.oss.util.models;

import com.aliyun.tea.*;

public class ReplicationProgressRule extends TeaModel {
    /**
     * <strong>example:</strong>
     * <p>ALL</p>
     */
    @NameInMap("Action")
    public String action;

    @NameInMap("Destination")
    public ReplicationDestination destination;

    /**
     * <strong>example:</strong>
     * <p>disabled</p>
     */
    @NameInMap("HistoricalObjectReplication")
    public String historicalObjectReplication;

    /**
     * <strong>example:</strong>
     * <p>replicate001</p>
     */
    @NameInMap("ID")
    public String ID;

    @NameInMap("PrefixSet")
    public ReplicationPrefixSet prefixSet;

    @NameInMap("Progress")
    public Progress progress;

    /**
     * <strong>example:</strong>
     * <p>doing</p>
     */
    @NameInMap("Status")
    public String status;

    public static ReplicationProgressRule build(java.util.Map<String, ?> map) throws Exception {
        ReplicationProgressRule self = new ReplicationProgressRule();
        return TeaModel.build(map, self);
    }

    public ReplicationProgressRule setAction(String action) {
        this.action = action;
        return this;
    }
    public String getAction() {
        return this.action;
    }

    public ReplicationProgressRule setDestination(ReplicationDestination destination) {
        this.destination = destination;
        return this;
    }
    public ReplicationDestination getDestination() {
        return this.destination;
    }

    public ReplicationProgressRule setHistoricalObjectReplication(String historicalObjectReplication) {
        this.historicalObjectReplication = historicalObjectReplication;
        return this;
    }
    public String getHistoricalObjectReplication() {
        return this.historicalObjectReplication;
    }

    public ReplicationProgressRule setID(String ID) {
        this.ID = ID;
        return this;
    }
    public String getID() {
        return this.ID;
    }

    public ReplicationProgressRule setPrefixSet(ReplicationPrefixSet prefixSet) {
        this.prefixSet = prefixSet;
        return this;
    }
    public ReplicationPrefixSet getPrefixSet() {
        return this.prefixSet;
    }

    public ReplicationProgressRule setProgress(Progress progress) {
        this.progress = progress;
        return this;
    }
    public Progress getProgress() {
        return this.progress;
    }

    public ReplicationProgressRule setStatus(String status) {
        this.status = status;
        return this;
    }
    public String getStatus() {
        return this.status;
    }

    public static class Progress extends TeaModel {
        /**
         * <strong>example:</strong>
         * <p>0.85</p>
         */
        @NameInMap("HistoricalObject")
        public String historicalObject;

        /**
         * <strong>example:</strong>
         * <p>Thu, 24 Sep 2015 15:39:18 GMT</p>
         */
        @NameInMap("NewObject")
        public String newObject;

        public static Progress build(java.util.Map<String, ?> map) throws Exception {
            Progress self = new Progress();
            return TeaModel.build(map, self);
        }

        public Progress setHistoricalObject(String historicalObject) {
            this.historicalObject = historicalObject;
            return this;
        }
        public String getHistoricalObject() {
            return this.historicalObject;
        }

        public Progress setNewObject(String newObject) {
            this.newObject = newObject;
            return this;
        }
        public String getNewObject() {
            return this.newObject;
        }

    }

}
