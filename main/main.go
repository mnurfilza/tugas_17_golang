package main

import (
	"belajargolang/tugas17"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/index", HandlerGetData)
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func HandlerGetData(w http.ResponseWriter, r *http.Request) {
	data, err := tugas17.Tugas17()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	jsonData, _ := json.Marshal(data)
	w.Write(jsonData)
}
