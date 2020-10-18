// main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Meeting
type Meeting struct {
	Id         string `json:"Id"`
	Title      string `json:"Title"`
	Partipants string `json:"part"`
	Start      string `json:"start"`
	End        string `json:"end"`
	Creation   string `json:"create"`
}

//Participants

type Participants struct {
	Name  string `json:"Id"`
	Email string `json:"Title"`
	RSVP  string `json:"part"`
}

var Meetings []Meeting
var Participants []Participants

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Schedule a meeting")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllMeetings(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllMeetings")
	json.NewEncoder(w).Encode(Meetings)
}

func returnSingleMeeting(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, meeting := range Meetings {
		if meeting.Id == key {
			json.NewEncoder(w).Encode(meeting)
		}
	}
}

func createNewMeeting(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var meeting Meeting
	json.Unmarshal(reqBody, &meeting)
	Meetings = append(Meetings, meeting)

	json.NewEncoder(w).Encode(meeting)
}

func deleteMeeting(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, meeting := range Meetings {
		if meeting.Id == id {
			Meetings = append(Meetings[:index], Meetings[index+1:]...)
		}
	}

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/meetings", returnAllMeetings)
	myRouter.HandleFunc("/meeting", createNewMeeting).Methods("POST")
	myRouter.HandleFunc("/{id}", deleteMeeting).Methods("DELETE")
	myRouter.HandleFunc("/{id}", returnSingleMeeting)
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {
	Meetings = []Meeting{
		Meeting{Id: "1", Title: "Induction", Partipants: "John", Start: "10:00", End: "11:30", Creation: "8:23"},
		Meeting{Id: "2", Title: "Brainstorming", Partipants: "Mark", Start: "11:40", End: "02:30", Creation: "8:23"},
	}
	Participants = []Meeting{
		Meeting{Name: "John", Email: "Induction", RSVP: "Yes"},
		Meeting{Name: "Mark", Email: "Brainstorming", RSVP: "No"},
	}
	handleRequests()
}
