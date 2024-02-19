package main

import (
	"golang-book-management-system/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "gorm.io/driver/mysql"
)

const (
	port = ":9080"
)

func main() {
	r := mux.NewRouter()
	routes.ReigisterBookRoutes(r)

	log.Printf("Starting Server on port %v", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
