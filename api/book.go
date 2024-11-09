package handler

import (
	"encoding/json"
	"net/http"

	"github.com/FAN2257/PELATIHAN_KMTETI-GO_FINALPROJECT/src/service"
)

func BookHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		action := r.URL.Query().Get("action")
		switch action {
		case "displayAll":
			data, err := service.GetAllBook()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
			return

		case "displayDetails":
			id := r.URL.Query().Get("id")
			data, err := service.GetBookByID(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
			return
		}
		
	case "PUT":
		id := r.URL.Query().Get("id")
		err := service.UpdateBook(id, r.Body)
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
		json.NewEncoder(w).Encode("book updated successfully")
		return

	case "POST":
		err := service.CreateBook(r.Body)
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
		json.NewEncoder(w).Encode("book created successfully")
		return

	}
}
