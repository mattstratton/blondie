package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mattstratton/blondie/talks/dbclient"
)

var DBClient dbclient.IMongoClient

func GetTalk(w http.ResponseWriter, r *http.Request) {

	// Read the 'talkID' path parameter from the mux map
	var talkID = mux.Vars(r)["talkID"]

	// Read the account struct BoltDB
	talk, err := DBClient.QueryTalk(talkID)

	// If err, return a 404
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// If found, marshal into JSON, write headers and content
	data, _ := json.Marshal(talk)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
