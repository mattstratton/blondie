package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"html/template"
)

var appName = "web"
var accountsURL string

func init() {

	accountsURL = "http://accounts:8080/accounts/"

}

func main() {
	log.Printf("Starting %v\n", appName)
	http.HandleFunc("/", accountHandler)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func accountHandler(w http.ResponseWriter, r *http.Request) {
	accountID := "599d910660e6e100078f2f4b" // TODO: Replace this with dynamic; take it from argument

	safeAccountID := url.QueryEscape(accountID)

	url := fmt.Sprintf("%s/%s", accountsURL, safeAccountID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	defer resp.Body.Close()

	var record Account

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	t, _ := template.ParseFiles("templates/accountDetail.html")
	t.Execute(w, record)

}

// Account represents the format of an account. Helpful, no?
type Account struct {
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}
