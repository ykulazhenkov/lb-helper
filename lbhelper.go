package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

const (
	defaultHTTPPort = "8080"
	envHTTPPort     = "HTTP_PORT"
	envNodeName     = "NODE_NAME"
	envNodeIP       = "NODE_IP"
	envPodName      = "POD_NAME"
	envPodNamespace = "POD_NAMESPACE"
	envPodIP        = "POD_IP"
)

type staticPodInfo struct {
	NodeName     string `json:"node_name"`
	NodeIP       string `json:"node_ip"`
	PodName      string `json:"pod_name"`
	PodNamespace string `json:"pod_namespace"`
	PodIP        string `json:"pod_ip"`
}

type respData struct {
	staticPodInfo
	ClientIP  string `json:"client_ip"`
	CallCount int    `json:"call_count"`
}

var counter int

var podInfo staticPodInfo

func h(w http.ResponseWriter, req *http.Request) {
	log.Printf("request from: %s", req.RemoteAddr)
	counter++
	data, err := json.Marshal(respData{
		staticPodInfo: podInfo,
		ClientIP:      req.RemoteAddr,
		CallCount:     counter})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Write(data)
}

func readStaticPodInfo() {
	podInfo = staticPodInfo{
		NodeName:     os.Getenv(envNodeName),
		NodeIP:       os.Getenv(envNodeIP),
		PodName:      os.Getenv(envPodName),
		PodNamespace: os.Getenv(envPodNamespace),
		PodIP:        os.Getenv(envPodIP),
	}
}

func main() {
	readStaticPodInfo()
	log.Printf("%+v\n", podInfo)
	http.HandleFunc("/", h)
	httpPort := os.Getenv(envHTTPPort)
	if httpPort == "" {
		httpPort = defaultHTTPPort
	}
	log.Printf("listen on port %s", httpPort)
	log.Fatal(http.ListenAndServe(":"+httpPort, nil))
}
