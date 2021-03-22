#!/bin/bash

curl -X POST localhost:7201/api/v1/database/create -d '{
    "type": "cluster",
    "namespace_name": "default",
    "retention_time": "8h",
    "num_shards": 64,
    "replication_factor": 2,
    "hosts": [
        {
            "id": "m3db-1",
            "isolation_group": "ig-1",
            "zone": "embedded",
            "weight": 100,
            "address": "m3db1",
            "port": 19000
        },
      {
            "id": "m3db-2",
            "isolation_group": "ig-2",
            "zone": "embedded",
            "weight": 100,
            "address": "m3db2",
            "port": 29000
        }

]}'

sleep 2

curl -H "Cluster-Environment-Name: default_env" -X POST http://127.0.0.1:7201/api/v1/services/m3coordinator/placement/init -d '{
    "instances": [                    
        {
            "id": "m3coordinator-1",
            "isolation_group": "ig-1",
            "zone": "embedded",
            "weight": 100,
            "endpoint": "m3coordinator1:17507",
            "hostname": "m3coordinator1",
            "port": 17507
        },
        {
            "id": "m3coordinator-2",
            "isolation_group": "ig-2",
            "zone": "embedded",
            "weight": 100,
            "endpoint": "m3coordinator2:27507",
            "hostname": "m3coordinator2",
            "port": 27507
        }
    ]
}'

sleep 2

curl -H "Cluster-Environment-Name: default_env" -X POST http://127.0.0.1:7201/api/v1/services/m3aggregator/placement/init -d '{
    "num_shards": 64,                
    "replication_factor": 1,
    "instances": [
        {
            "id": "m3aggregator-1",
            "isolation_group": "ig-1",
            "zone": "embedded",
            "weight": 100,
            "endpoint": "m3aggregator1:16000",
            "hostname": "m3aggregator1",
            "port": 16000
        },
        {
            "id": "m3aggregator-2",
            "isolation_group": "ig-2",
            "zone": "embedded",
            "weight": 100,
            "endpoint": "m3aggregator2:26000",
            "hostname": "m3aggregator2",
            "port": 26000
        }
    ]
}'

sleep 2

curl -H "Cluster-Environment-Name: default_env" -H "Topic-Name: aggregator_ingest" -X POST http://localhost:7201/api/v1/topic/init -d '{
    "numberOfShards": 64    
}'

sleep 2

curl -vvvsSf -H "Cluster-Environment-Name: default_env" -H "Topic-Name: aggregator_ingest" -X POST http://127.0.0.1:7201/api/v1/topic -d '{ 
  "consumerService": {      
    "serviceId": {
      "name": "m3aggregator",
      "environment": "default_env",
      "zone": "embedded"
    },
    "consumptionType": "REPLICATED",
    "messageTtlNanos": "300000000000"
  }
}'

sleep 2

curl -H "Cluster-Environment-Name: default_env" -H "Topic-Name: aggregator_metrics" -X POST http://localhost:7201/api/v1/topic/init -d '{
    "numberOfShards": 64    
}'

sleep 2

curl -H "Cluster-Environment-Name: default_env" -H "Topic-Name: aggregator_metrics" -X POST http://localhost:7201/api/v1/topic -d '{
  "consumerService": {      
    "serviceId": {
      "name": "m3coordinator",
      "environment": "default_env",
      "zone": "embedded"
    },
    "consumptionType": "SHARED",
    "messageTtlNanos": "300000000000"
  }
}'

sleep 2

podman-compose -p m3db-cluster -f podman-compose-m3db-cluster.yml restart m3coordinator1
podman-compose -p m3db-cluster -f podman-compose-m3db-cluster.yml restart m3coordinator2
podman-compose -p m3db-cluster -f podman-compose-m3db-cluster.yml restart m3aggregator1
podman-compose -p m3db-cluster -f podman-compose-m3db-cluster.yml restart m3aggregator2

sleep 2

curl -X POST http://127.0.0.1:17201/api/v1/services/m3db/namespace -d '{     
  "name": "test_aggregated",
  "options": {
    "bootstrapEnabled": true,
    "flushEnabled": true,
    "writesToCommitLog": true,
    "cleanupEnabled": true,
    "snapshotEnabled": true,
    "repairEnabled": false,
    "retentionOptions": {
      "retentionPeriodDuration": "48h",
      "blockSizeDuration": "2h",
      "bufferFutureDuration": "10m",
      "bufferPastDuration": "10m",
      "blockDataExpiry": true
    },
    "indexOptions": {
      "enabled": true,
      "blockSizeDuration": "2h"
    },
    "aggregationOptions": {
      "aggregations": [
        {
            "aggregated": true,
            "attributes": {
                "downsampleOptions": {
                    "all": true
                },
                "resolutionDuration": "10m0s"
            }
        }
      ]
    }
  }
}'

sleep 2

curl -X POST http://localhost:7201/api/v1/services/m3db/namespace/ready -d '{
  "name": "default"
}'

sleep 2

curl -X POST http://localhost:7201/api/v1/services/m3db/namespace/ready -d '{
  "name": "test_aggregated"
}'
