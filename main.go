package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Login struct {
	rancherURL string
	token      string
	clusterID  string
	projectID  string
}

func checkenv() {
	var buf bytes.Buffer
	ErrorLogger := log.New(&buf, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	if os.Getenv("RANCHER_SERVER") == "" || os.Getenv("RANCHER_TOKEN") == "" {

		ErrorLogger.Print("Environment Variable RANCHER_SERVER and or RANCHER_TOKEN not set.")
		fmt.Print(&buf)
		os.Exit(1)
	}
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
		dt := time.Now()

		finaloutput = "<h1>Rancher Managed Cluster Information!</h1><h2 style=color:green;>peer review app</h2><h3>" + dt.Format("02 Jan 2006 15:04:05") + "</h3>"

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
	checkenv()
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":3000", mux)

}
