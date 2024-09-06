package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type NumbersRequest struct {
	Numbers []int32 `json:"numbers"`
}

type URLRequest struct {
	URL string `json:"url"`
}

func sumNumbers(w http.ResponseWriter, r *http.Request) {
	var req NumbersRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var sum int32
	for _, num := range req.Numbers {
		sum += num
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"result": %d}`, sum)))
}

func processURL(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Could not read request body", http.StatusBadRequest)
		return
	}

	url := string(body)
	log.Printf("Received URL: %s", url)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"message": "URL received: %s"}`, url)))
}

func main() {
	http.HandleFunc("/sum", sumNumbers)
	http.HandleFunc("/url", processURL)

	fmt.Println("Server is listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
