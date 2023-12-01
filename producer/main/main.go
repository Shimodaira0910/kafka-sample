package main

import (
	"fmt"
	"kafka/producer/api"
	"net/http"
)

func main(){
	api := api.API{}
	http.HandleFunc("/post", api.Post)
	fmt.Println("Server is listening on :8002")
	http.ListenAndServe(":8002", nil)
}

