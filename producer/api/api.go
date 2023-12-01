package api

import (
	"encoding/json"
	"fmt"
	"io"
	"kafka/producer/models"
	"log"
	"net/http"
)

type HandleAPI interface{
	Post()
}

type API struct{
	traffic *models.Traffic
}

func (a *API) Post(w http.ResponseWriter, r *http.Request){
	body, err := io.ReadAll(r.Body)
	if err != nil{
		log.Fatalln(err)
	}

	if err := json.Unmarshal(body, &a.traffic); err != nil {
		log.Fatalln(err)
	}
	
	fmt.Println(a.traffic.Type)
}