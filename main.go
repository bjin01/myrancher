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
	dt := time.Now()
	myhead := `<head>
	<style>
	* {
		background-color: blue;
	}
	
	div.a {
		text-align: center;
	}
	h2 {
		color: blue;
	  }
	.row {
		display: flex;
		}
		
		/* Create two equal columns that sits next to each other */
		.column {
		flex: 50%;
		padding: 10px;
		height: 300px; /* Should be removed. Only for demonstration */
		}
	</style>
	</head>`

	finaloutput = myhead + "<body><div class=\"a\"><h1>Rancher Cluster Information!</h1><h2 style=color:white;>API app</h2><h3>" + dt.Format("02 Jan 2006 15:04:05") + "</h3></div><div class=\"row\">"

	if len(mycluster.Data) != 0 {
		for _, b := range mycluster.Data {
			p1 := "<div class=\"column\"><h2>ClusterName: " + b.Cname + "</h2>"
			p2 := "<p>ClusterID: " + b.Cid + "</p>"
			p3 := "<p>Provider: " + b.Cprovider + "</p>"
			p4 := "<p>Number of nodes: " + strconv.Itoa(b.Cnodes) + "</p>"
			p41 := "<p>Capacity: " + "CPU:" + b.Ccapacity["cpu"] + ", Memory: " + b.Ccapacity["memory"] + ", Pods: " + b.Ccapacity["pods"] + "</p>"
			p42 := "<p>K8s Version: " + b.Cversion["gitVersion"] + "</p>"
			p5 := "<p>Status: " + b.Cstate + "</p>"
			if strings.Contains(b.Cstate, "error") || strings.Contains(b.Cstate, "unavailable") {
				p5 = "<p style=color:red;>Status: " + "<b>" + b.Cstate + "&#128078;</b></p>"
			}
			if strings.Contains(b.Cstate, "active") {
				p5 = "<p style=color:white;>Status: " + "<b>" + b.Cstate + "&#128077;</b></p>"
			}
			//p6 := "<p>--------------------</p></div>"
			finaloutput += p1 + p2 + p3 + p4 + p41 + p42 + p5 + "</div>"
		}
	}
	finaloutput = finaloutput + "</div></body>"
	w.Header().Set("bo", "jin")
	w.Write([]byte(finaloutput))

}

func main() {
	checkenv()
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":3000", mux)

}
