package main

import (
	"encoding/json"
	"fmt"
	"github.com/Kagami/go-face"
	"github.com/gorilla/mux"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"time"
)

const PORT = ":5000"

var (
	spa = spaHandler{
		staticPath: "public",
		indexPath:  "index.html",
	}
)

func recognizeFacesHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "file is more than 10 mb", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image")

	if err != nil {
		http.Error(w, "form is invalid", http.StatusBadRequest)
		return
	}

	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			http.Error(w, "file can`t close", http.StatusInternalServerError)
			return
		}
	}(file)

	img, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "file isn`t read", http.StatusInternalServerError)
		return
	}

	faces, err := recognizeFaces(img)
	if err != nil {
		http.Error(w, "recognize face error", http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(faces)
	if err != nil {
		http.Error(w, "serialize face error", http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprintf(w, "%s", data)
	if err != nil {
		http.Error(w, "file isn`t sent", http.StatusInternalServerError)
		return
	}
}

func main() {
	var err error
	recognizer, err = face.NewRecognizer("models")
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/api/recognizeFaces", recognizeFacesHandler).Methods("POST")
	router.PathPrefix("/").Handler(spa)

	srv := &http.Server{
		Handler:      router,
		Addr:         PORT,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
