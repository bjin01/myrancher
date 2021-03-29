# myrancher

This program makes rest api calls to rancher platform and queries certain cluster data. It shows how easy it is using rest api of rancher getting some data using golang.

Before running the binary one needs to export those environment variables in order to provide cluster authentication data.
```
export RANCHER_TOKEN="token-stvff:95phk8djrxfwgbcsdf7g5z2wx5s54zqdbhctfd2q6tqvgrt4hhrgxnr"
export RANCHER_SERVER="borancher.my.cloud"
export RANCHER_SERVER="https://borancher.my.cloud"
export RANCHER_CLUSTER_ID="local"
```
Sample result to query managed clusters within rancher:
```
$ go run .
ClusterID: c-kqrqv 
ClusterName: ec2rke-import
        Provider: rke Nodes:3
        State: error
        Message: cluster health check failed: Failed to communicate with API server during namespace check: Get "https://10.43.0.1:443/api/v1/namespaces/kube-system?timeout=45s": context deadline exceeded
---
ClusterID: c-rb942 
ClusterName: 2nd-rke
        Provider: rke Nodes:3
        State: active
        Message: 
---
ClusterID: local 
ClusterName: openstack-rancher
        Provider: rke Nodes:3
        State: active
        Message: 
---
```
