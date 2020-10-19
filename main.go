package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Participant struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
	RSVP  string `json:"RSVP"`
}

type Meeting struct {
	Id                string `json:"Id"`
	Title             string `json:"Title"`
	Participants      int    `json:"Participants"`
	StartTime         string `json:"StartTime"`
	EndTime           string `json:"EndTime"`
	CreationTimeStamp string `json:"CreationTimeStamp"`
}

//var Articles []Meeting
var obj Meeting

func allMeetings(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(obj)

}

func returnSingleMeeting(w http.ResponseWriter, r *http.Request) {
	var temp string

	fmt.Println("Enter the meeting ID: ")
	fmt.Scanln(&temp)

	if obj.Id == temp {
		json.NewEncoder(w).Encode(obj)
	}

}

func schedulemeeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Schedule a meeting:- ")

	fmt.Println("ID: ")
	fmt.Scanln(&obj.Id)

	fmt.Println("Title: ")
	fmt.Scanln(&obj.Title)

	fmt.Println("Participants: ")
	fmt.Scanln(&obj.Participants)

	fmt.Println("Start Time: ")
	fmt.Scanln(&obj.StartTime)

	fmt.Println("End Time: ")
	fmt.Scanln(&obj.EndTime)

	fmt.Println("Creation Timestamp: ")
	fmt.Scanln(&obj.CreationTimeStamp)

	fmt.Println("Endpoint Hit: schedulemeeting")

	json.NewEncoder(w).Encode(obj)

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the APPOINTY TASK!")

	fmt.Println("Endpoint Hit: homepage")
}

func handleRequests() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/meetings", allMeetings)
	http.HandleFunc("/schedule", schedulemeeting)
	http.HandleFunc("/meeting/return", returnSingleMeeting)

	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {

	handleRequests()
}
