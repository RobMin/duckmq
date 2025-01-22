package reciever

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type MessageRequest struct {
	Id        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
}

func Main() {
	http.HandleFunc("/message", func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.Error(res, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		body, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(res, "Error reading body", http.StatusBadRequest)
			return
		}
		defer req.Body.Close()

		var data MessageRequest
		fmt.Print(string(body))
		err = json.Unmarshal(body, &data)
		if err != nil {
			http.Error(res, fmt.Sprintf("Invalid JSON: %s", err.Error()), http.StatusBadRequest)
			return
		}

		response := fmt.Sprintf("Received: Id=%s, Timestamp=%s, Message=%s", data.Id, data.Timestamp.String(), data.Message)
		res.Write([]byte(response))
	})

	fmt.Println("Server is running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
