package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/MidhunRajeevan/notification/model"
	"github.com/go-playground/validator/v10"
)

func Notification(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		// get list of notification
	case http.MethodPost:
		CreateNotification(w, r)
	}
}

func CreateNotification(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error occured while reading the request data")
	}

	var message model.MessageRequest
	err = json.Unmarshal(data, &message)
	if err != nil {
		log.Println("bad request")
		BadRequest(w, "bad request,"+err.Error())
	}
	err = validator.New().Struct(message)
	if err != nil {
		log.Println("bad request")
		BadRequest(w, "bad request,"+err.Error())
	}

	fmt.Println(message)

}
