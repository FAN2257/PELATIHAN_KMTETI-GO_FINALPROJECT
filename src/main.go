package main

import (
	"fmt"
	"net/http"

	handler "github.com/FAN2257/PELATIHAN_KMTETI-GO_FINALPROJECT/api"
)

func main() {
	h := http.NewServeMux()

	s := &http.Server{
		Addr:    ":8080",
		Handler: h,
	}

	h.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	h.HandleFunc("/api/book", handler.BookHandler)

	h.HandleFunc("/api/employee", handler.EmployeeHandler)

	fmt.Println("HTTP Server running on port 8080")
	fmt.Println("link: http://localhost:8080")
	defer fmt.Println("HTTP Server stopped")

	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
