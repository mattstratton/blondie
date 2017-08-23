package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/context"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {

	// connect to db
	db, err := mgo.Dial("db")
	if err != nil {
		log.Fatal("cannot dial mongo - ", err)
	}
	defer db.Close() //clean up when done

	h := Adapt(http.HandlerFunc(handle), withDB(db))

	http.Handle("/accounts", context.ClearHandler(h)) //TODO: This should handle the slash and not accounts

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

// The Adapter function type allows us to write code that can be run before/after HTTP requests
type Adapter func(http.Handler) http.Handler

func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

func withDB(db *mgo.Session) Adapter {

	// return the Adapter
	return func(h http.Handler) http.Handler {
		// when calling the adapter is should return a new Handler
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// copy database Session
			dbsession := db.Copy()
			defer dbsession.Close() // clean up our Session

			// save it in mux context
			context.Set(r, "database", dbsession)

			// pass to the original Handler
			h.ServeHTTP(w, r)
		})
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleRead(w, r)
	case "POST":
		handleInsert(w, r)
	default:
		http.Error(w, "Not supported", http.StatusMethodNotAllowed)
	}
}

type account struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Username  string        `json:"username" bson:"username"`
	FirstName string        `json:"firstname" bson:"firstname"`
	LastName  string        `json:"lastname" bson:"lastname"`
	Email     string        `json:"email" bson:"email"`
	When      time.Time     `json:"when" bson:"when"`
}

func handleInsert(w http.ResponseWriter, r *http.Request) {
	db := context.Get(r, "database").(*mgo.Session)
	// decode the request body
	var a account
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// give the account a unique ID and set the time
	a.ID = bson.NewObjectId()
	a.When = time.Now()
	// insert it into the database
	if err := db.DB("blondie").C("accounts").Insert(&a); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// redirect to it
	http.Redirect(w, r, "/accounts/"+a.ID.Hex(), http.StatusTemporaryRedirect)
}

func handleRead(w http.ResponseWriter, r *http.Request) {
	db := context.Get(r, "database").(*mgo.Session)
	// load the accounts
	var accounts []*account
	if err := db.DB("blondie").C("accounts").
		Find(nil).Sort("-when").Limit(100).All(&accounts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// write it out
	if err := json.NewEncoder(w).Encode(accounts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
