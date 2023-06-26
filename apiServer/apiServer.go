package main

import (
	"Distributed_Object_Storage/apiServer/heartbeat"
	"Distributed_Object_Storage/apiServer/locate"
	"Distributed_Object_Storage/apiServer/objects"
	"Distributed_Object_Storage/apiServer/temp"
	"Distributed_Object_Storage/apiServer/versions"
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
