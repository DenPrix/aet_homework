package main

import (
	"aet_homework/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func GetContact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	result, err := utils.Database.Query(utils.GetData(), params["id"])
	if err != nil {
		log.Println(err)
	}
	defer result.Close()
	var contact utils.Contact
	for result.Next() {
		err := result.Scan(&contact.Id, &contact.Name, &contact.Phone, &contact.Email)
		if err != nil {
			log.Println(err)
		}
	}
	json.NewEncoder(w).Encode(contact)
}

func CreateContact(w http.ResponseWriter, r *http.Request) {
	stmt := utils.PrepareDatabase(utils.InsertData)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	newContact := utils.Contact{}
	json.Unmarshal(body, &newContact)

	lastId := utils.ExecDatabase(newContact, stmt)
	fmt.Fprintf(w, "{\"Id\":%v}", lastId)
}

func UpdateContact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stmt := utils.PrepareDatabase(utils.UpdateData)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	updContact := utils.Contact{}
	json.Unmarshal(body, &updContact)

	utils.ExecDatabaseWithId(updContact, stmt, params["id"])
	fmt.Fprintf(w, "{\"Id\":%v}", params["id"])
}

func DeleteContact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stmt := utils.PrepareDatabase(utils.DeleteData)
	utils.ExecDatabaseIdOnly(params["id"], stmt)
	fmt.Fprintf(w, "{\"Id\":%v}", params["id"])
}
