package main

import (
	"Distributed_Storage_Deploy/apiServer/heartbeat"
	"Distributed_Storage_Deploy/apiServer/locate"
	"Distributed_Storage_Deploy/apiServer/objects"
	"Distributed_Storage_Deploy/apiServer/temp"
	"Distributed_Storage_Deploy/apiServer/versions"
	"log"
	"net/http"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", versions.Handler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
