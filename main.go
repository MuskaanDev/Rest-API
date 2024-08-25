package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type PostRequest struct {
	Name      string   `json:"name"`
	DOB       string   `json:"dob"`
	Email     string   `json:"email"`
	RollNo    string   `json:"roll_no"`
	Numbers   []int    `json:"numbers"`
	Alphabets []string `json:"alphabets"`
}

type PostResponse struct {
	Status          bool     `json:"is_success"`
	UserID          string   `json:"user_id"`
	Email           string   `json:"email"`
	RollNo          string   `json:"roll_no"`
	Numbers         []int    `json:"numbers"`
	Alphabets       []string `json:"alphabets"`
	HighestAlphabet string   `json:"highest_alphabet"`
}

func getHighestAlphabet(alphabets []string) string {
	highest := ""
	for _, char := range alphabets {
		if char >= highest {
			highest = char
		}
	}
	return highest
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	var req PostRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID := fmt.Sprintf("%s_%s", strings.ReplaceAll(strings.ToLower(req.Name), " ", "_"), req.DOB)
	highestAlphabet := getHighestAlphabet(req.Alphabets)

	res := PostResponse{
		Status:          true,
		UserID:          userID,
		Email:           req.Email,
		RollNo:          req.RollNo,
		Numbers:         req.Numbers,
		Alphabets:       req.Alphabets,
		HighestAlphabet: highestAlphabet,
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
