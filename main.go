package main

import (
	"log"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Request processed"))
}

func main() {
	// 注册 Prometheus 指标端点

	// 注册业务处理端点
	http.HandleFunc("/api/process", handleRequest)

	log.Println("Server starting on 0.0.0.0:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
