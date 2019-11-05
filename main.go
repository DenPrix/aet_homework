package main

import (
	"aet_homework/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func main() {
	utils.InitConfig()
	utils.InitDatabase()
	defer utils.CloseDatabase()
	router := mux.NewRouter()

	apiLink := "/api/" + utils.Config.Endpoint
	apiLinkId := apiLink + "/{id}"

	router.HandleFunc(apiLinkId, GetContact).Methods(http.MethodGet)
	router.HandleFunc(apiLink, CreateContact).Methods(http.MethodPost)
	router.HandleFunc(apiLinkId, UpdateContact).Methods(http.MethodPatch)
	router.HandleFunc(apiLinkId, DeleteContact).Methods(http.MethodDelete)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":"+strconv.Itoa(utils.Config.Port), router)
}
