package main

import (
	"encoding/json"
	"fmt"

	resty "github.com/go-resty/resty/v2"
)

type Clusterinfo struct {
	Cid                string            `json:"id"`
	Cname              string            `json:"name"`
	Cstate             string            `json:"state"`
	Ctransitioning     string            `json:"transitioning"`
	CtransitionMessage string            `json:"transitioningMessage"`
	Cprovider          string            `json:"provider"`
	Cnodes             int               `json:"nodeCount"`
	Ccreated           string            `json:"created"`
	Ccapacity          map[string]string `json:"capacity"`
	Cversion           map[string]string `json:"version"`
}

type Clusterdata struct {
	Data []Clusterinfo `json:"data"`
}

//test
func (c *Login) getresty() Clusterdata {
	//client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client := resty.New()
	var myclusterdata Clusterdata
	resp, err := client.R().
		SetAuthToken(c.token).
		Get(c.rancherURL + "/v3/clusters/")
	if err != nil {
		fmt.Printf("Failed to query Rancher clusters: %v", err)
	}
	body := string(resp.Body()[:])
	json.Unmarshal([]byte(body), &myclusterdata)
	if len(myclusterdata.Data) != 0 {
		for _, b := range myclusterdata.Data {
			fmt.Printf("ClusterID: %v \nClusterName: %v\n\tProvider: %v Nodes:%v\n\tCreated: %v\n", b.Cid, b.Cname, b.Cprovider, b.Cnodes, b.Ccreated)
			fmt.Printf("Cluster Capacity %v\n", b.Ccapacity)
			fmt.Printf("\tState: %v\n\tMessage: %v\n", b.Cstate, b.CtransitionMessage)
			fmt.Printf("Kubernetes Version: %v\n", b.Cversion["gitVersion"])
			fmt.Println("---")
		}
		return myclusterdata
	}
	return myclusterdata

}
