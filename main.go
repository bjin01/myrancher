package main

import (
	"os"
)

type Login struct {
	rancherURL string
	token      string
	clusterID  string
	projectID  string
}

func main() {
	mylogin := Login{
		rancherURL: os.Getenv("RANCHER_SERVER"),
		token:      os.Getenv("RANCHER_TOKEN"),
		clusterID:  os.Getenv("RANCHER_CLUSTER_ID"),
		projectID:  os.Getenv("RANCHER_PROJECT_ID"),
	}
	mylogin.getresty()

}
