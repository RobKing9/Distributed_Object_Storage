package main

import (
	"Distributed_Storage_Deploy/dataServer/heartbeat"
	"Distributed_Storage_Deploy/dataServer/locate"
	"Distributed_Storage_Deploy/dataServer/objects"
	"Distributed_Storage_Deploy/dataServer/temp"
	"log"
	"net/http"
	"os"
)

func main() {
	locate.CollectObjects()
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
