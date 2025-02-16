package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func timeHandler(writer http.ResponseWriter, req *http.Request) {
	response := struct {
		Time string `json:"time"`
	}{
		Time: time.Now().Format(time.RFC3339),
	}

	writer.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		http.Error(writer, "JSON encoding error", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/time", timeHandler)

	port := "8795"
	fmt.Printf("The server is running on http://localhost:%s/time\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
