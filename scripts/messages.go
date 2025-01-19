package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	jsonData, err := json.Marshal(map[string]any{
		"id":        "id-to-generate",
		"timestamp": "2025-01-22T15:04:05Z",
		"message":   "some message",
	})
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://localhost:8080/message", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
