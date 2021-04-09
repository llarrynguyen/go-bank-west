package app

import (
	"com.github/llarrynguyen/bank_west/app/domain"
	"com.github/llarrynguyen/bank_west/app/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func Start() {
	//wiring
	//ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDB())}
	router := mux.NewRouter()
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)


	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
