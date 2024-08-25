package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type PostRequest struct {
	Data []string `json:"data"`
}

type PostResponse struct {
	IsSuccess                bool     `json:"is_success"`
	UserID                   string   `json:"user_id"`
	Email                    string   `json:"email"`
	RollNumber               string   `json:"roll_number"`
	Numbers                  []string `json:"numbers"`
	Alphabets                []string `json:"alphabets"`
	HighestLowercaseAlphabet []string `json:"highest_lowercase_alphabet"`
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	var req PostRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	numbers := []string{}
	alphabets := []string{}
	var highestLowercaseAlphabet string

	for _, item := range req.Data {
		if _, err := strconv.Atoi(item); err == nil {
			numbers = append(numbers, item)
		} else if len(item) == 1 && ((item >= "a" && item <= "z") || (item >= "A" && item <= "Z")) {
			alphabets = append(alphabets, item)
			if item >= "a" && item <= "z" {
				if highestLowercaseAlphabet == "" || item > highestLowercaseAlphabet {
					highestLowercaseAlphabet = item
				}
			}
		}
	}

	var highestLowercaseAlphabetArray []string
	if highestLowercaseAlphabet != "" {
		highestLowercaseAlphabetArray = []string{highestLowercaseAlphabet}
	}

	res := PostResponse{
		IsSuccess:                true,
		UserID:                   "john_doe_17091999",
		Email:                    "john@xyz.com",
		RollNumber:               "ABCD123",
		Numbers:                  numbers,
		Alphabets:                alphabets,
		HighestLowercaseAlphabet: highestLowercaseAlphabetArray,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	res := map[string]int{"operation_code": 1}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/bfhl", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlePost(w, r)
		} else if r.Method == http.MethodGet {
			handleGet(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
