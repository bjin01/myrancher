# myrancher

This program makes rest api calls to rancher platform and queries certain cluster data. It shows how easy it is using rest api of rancher getting some data using golang.

Before running the binary one needs to export those environment variables in order to provide cluster authentication data.
```
export RANCHER_SERVER="https://borancher.my.cloud/v3"
export RANCHER_TOKEN="token-stvff:95phk8djrxfwgbcsdf7g5z2wx5s54zqdbhctfd2q6tqvgrt4hhrgxnr"
export RANCHER_SERVER="borancher.my.cloud"
export RANCHER_SERVER="https://borancher.my.cloud"
export RANCHER_CLUSTER_ID="local"
```
