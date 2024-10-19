package main

import (
	"fmt"
	"net/http"
	"github.com/FAN2257/PELATIHAN_KMTETI-GO_Backend/src/handler"
)

func main() {
	fmt.Println("Hello, World!")
	s := &http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/", handler.GetProduct)

	fmt.Println("Server started at port 8080")
	err := s.ListenAndServe()

	if(err != nil){
		fmt.Println(err.Error())
	}
}