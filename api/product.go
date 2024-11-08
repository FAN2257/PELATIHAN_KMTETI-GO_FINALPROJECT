package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/FAN2257/PELATIHAN_KMTETI-GO_FINALPROJECT/src/service"
)

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		action := r.URL.Query().Get("action")
		switch action {
		case "displayAll":
			data, err := service.GetAllProduct()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
			return

		// case "displayDetails":
		// 	id := r.URL.Query().Get("id")
		// 	data, err := service.GetProductDetails(id)
		// 	if err != nil {
		// 		http.Error(w, err.Error(), http.StatusInternalServerError)
		// 		return
		// 	}
		}

	case "POST":
		err := service.CreateProduct(r.Body)
		if err != nil {
			if err.Error() == "bad request" {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("product created successfully")
		return
	default:
		log.Default().Println(http.StatusMethodNotAllowed)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}