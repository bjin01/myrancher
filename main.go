package main

import (
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Login struct {
	rancherURL string
	token      string
	clusterID  string
	projectID  string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	mylogin := Login{
		rancherURL: os.Getenv("RANCHER_SERVER"),
		token:      os.Getenv("RANCHER_TOKEN"),
		clusterID:  os.Getenv("RANCHER_CLUSTER_ID"),
		projectID:  os.Getenv("RANCHER_PROJECT_ID"),
	}
	mycluster := mylogin.getresty()
	var finaloutput string
	if len(mycluster.Data) != 0 {

		finaloutput = "<h1>Rancher Managed Cluster Information!</h1><h2 style=color:green;>peer review app</h2>"

		for _, b := range mycluster.Data {

			p1 := "<p>ClusterID: " + b.Cid + "</p>"
			p2 := "<p>ClusterName: " + b.Cname + "</p>"
			p3 := "<p>Provider: " + b.Cprovider + "</p>"
			p4 := "<p>Number of nodes: " + strconv.Itoa(b.Cnodes) + "</p>"
			p5 := "<p>Status: " + b.Cstate + "</p>"
			if strings.Contains(b.Cstate, "error") {
				p5 = "<p style=color:red;>Status: " + "<b>" + b.Cstate + "&#128078;</b></p>"
			}
			if strings.Contains(b.Cstate, "active") {
				p5 = "<p style=color:green;>Status: " + "<b>" + b.Cstate + "&#128077;</b></p>"
			}
			p6 := "<p>--------------------</p>"
			finaloutput += p1 + p2 + p3 + p4 + p5 + p6
		}
		//finaloutput += finaloutput
	}
	w.Write([]byte(finaloutput))

}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":3000", mux)

}
